package app

import (
	"context"
	"database/sql"
	"fmt"
	"spotify/internal/metrics"
	"spotify/internal/middleware"
	"spotify/internal/server"
	"spotify/microservices/playlist/ai"
	"spotify/pkg/logger"
	"spotify/pkg/minio"
	"spotify/pkg/postgres"

	pbAuth "spotify/proto/auth"

	"google.golang.org/grpc/credentials/insecure"

	httpDelivery "spotify/microservices/playlist/delivery/http"
	repository "spotify/microservices/playlist/repository/postgres"
	storageRepo "spotify/microservices/playlist/repository/storage"
	service "spotify/microservices/playlist/service"

	pbCatalog "spotify/proto/catalog"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type App struct {
	cfg        *Config
	logger     logger.Logger
	db         *sql.DB
	minio      *minio.Client
	httpServer *server.Server
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load playlist config: %w", err)
	}

	appLogger, err := logger.New(cfg.Playlist.Logger.Level, cfg.Playlist.Logger.Mode)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}

	db, err := postgres.New(ctx, cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	if err := prometheus.Register(postgres.NewMonitor(db, cfg.DB.DBName)); err != nil {
		appLogger.Warnf("db metrics already registered or failed: %v", err)
	}

	minioClient, err := minio.New(cfg.Minio)
	if err != nil {
		return nil, fmt.Errorf("failed to init minio: %w", err)
	}

	stor := storageRepo.NewStorage(minioClient, cfg.Playlist.Buckets.Avatars)
	repo := repository.New(db)

	grpcprometheus.EnableClientHandlingTimeHistogram()

	if err := prometheus.Register(grpcprometheus.DefaultClientMetrics); err != nil {
		appLogger.Warnf("grpc client metrics already registered: %v", err)
	}

	catalogConn, err := grpc.NewClient(
		cfg.Playlist.Clients.Catalog,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcprometheus.UnaryClientInterceptor),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to catalog: %w", err)
	}

	catalogClient := pbCatalog.NewCatalogServiceClient(catalogConn)

	aiClient := ai.NewGigaChat(ai.GigaChatConfig{
		AuthKey: cfg.Playlist.AI.AuthKey,
		Model:   cfg.Playlist.AI.Model,
	})
	playlistService := service.New(repo, stor, catalogClient, aiClient)

	mtr := metrics.New("playlist")

	httpHandler := httpDelivery.NewHandler(
		playlistService,
		cfg.Playlist.AllowedAvatarTypes,
	)

	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())

	api := router.PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.RequestLoggerMiddleware(appLogger))
	api.Use(middleware.MetricsMiddleware(mtr))
	api.Use(middleware.CORS(cfg.Playlist.HTTP.CORS))

	authConn, err := grpc.NewClient(
		cfg.Playlist.Clients.Auth,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth: %w", err)
	}
	authClient := pbAuth.NewAuthServiceClient(authConn)
	authMiddleware := middleware.NewAuthGrpcMiddleware(authClient)

	protected := api.PathPrefix("").Subrouter()
	protected.Use(authMiddleware.Handle)

	csrfProtected := protected.PathPrefix("").Subrouter()
	csrfProtected.Use(authMiddleware.Handle)
	public := api.PathPrefix("").Subrouter()

	httpHandler.RegisterRoutes(public, protected, csrfProtected)

	httpServer := server.NewHTTPServer(&cfg.Playlist.HTTP, router, appLogger)

	return &App{
		cfg:        cfg,
		logger:     appLogger,
		db:         db,
		minio:      minioClient,
		httpServer: httpServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- a.httpServer.Run()
	}()

	select {
	case <-ctx.Done():
		shutCtx, cancel := context.WithTimeout(context.Background(), a.cfg.Playlist.HTTP.ShutdownTimeout)
		defer cancel()
		return a.httpServer.Shutdown(shutCtx)
	case err := <-errCh:
		return err
	}
}
