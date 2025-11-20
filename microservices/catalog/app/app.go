package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"spotify/internal/app"
	"spotify/internal/middleware"
	"spotify/internal/server"
	grpcDelivery "spotify/microservices/catalog/delivery/grpc"
	httpDelivery "spotify/microservices/catalog/delivery/http"
	catalogMiddleware "spotify/microservices/catalog/middleware"
	repository "spotify/microservices/catalog/repository/postgres"
	service "spotify/microservices/catalog/service"
	"spotify/pkg/logger"
	"spotify/pkg/postgres"
	pbAuth "spotify/proto/auth"
	pb "spotify/proto/catalog"
	"sync"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	cfg        *app.Config
	logger     logger.Logger
	db         *sql.DB
	httpServer *server.Server
	grpcServer *server.GRPCServer
	authConn   *grpc.ClientConn
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	cfg, err := app.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	appLogger, err := logger.New(cfg.Catalog.Logger.Level, cfg.Catalog.Logger.Mode)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}
	appLogger.Infof("Logger initialized for Catalog service")

	db, err := postgres.New(ctx, cfg.DB)
	if err != nil {
		appLogger.Errorf("failed to connect to db: %v", err)
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	appLogger.Infof("Database connection established")

	repo := repository.New(db)
	catalogService := service.New(repo)

	authConn, err := grpc.NewClient(
		cfg.Catalog.Clients.Auth,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	authClient := pbAuth.NewAuthServiceClient(authConn)
	authMiddleware := catalogMiddleware.NewAuthGrpcMiddleware(authClient)

	httpHandler := httpDelivery.NewHandler(catalogService)
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.RequestLoggerMiddleware(appLogger))
	api.Use(middleware.CORS(cfg.Catalog.HTTP.CORS))

	protected := api.PathPrefix("").Subrouter()
	protected.Use(authMiddleware.Handle)

	httpHandler.RegisterRoutes(api, protected)

	httpServer := server.NewHTTPServer(&cfg.Catalog.HTTP, router, appLogger)

	grpcHandler := grpcDelivery.NewHandler(catalogService)
	grpcServer := server.NewGRPCServer(&cfg.Catalog.GRPC, appLogger, func(s *grpc.Server) {
		pb.RegisterCatalogServiceServer(s, grpcHandler)
	})

	return &App{
		cfg:        cfg,
		logger:     appLogger,
		db:         db,
		httpServer: httpServer,
		grpcServer: grpcServer,
		authConn:   authConn,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	serverErrors := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.httpServer.Run(); err != nil && err != http.ErrServerClosed {
			serverErrors <- fmt.Errorf("http server error: %w", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.grpcServer.Run(); err != nil {
			serverErrors <- fmt.Errorf("grpc server error: %w", err)
		}
	}()

	a.logger.Infof("Catalog microservice is running...")

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server run failed: %w", err)
	case <-ctx.Done():
		a.logger.Infof("shutting down servers due to context cancellation...")
	}

	if err := a.Stop(); err != nil {
		return fmt.Errorf("failed to gracefully stop application: %w", err)
	}

	wg.Wait()
	a.logger.Infof("All servers stopped, application is shutting down.")
	return nil
}

func (a *App) Stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), a.cfg.Catalog.HTTP.ShutdownTimeout)
	defer cancel()

	errHTTP := a.httpServer.Shutdown(shutdownCtx)
	a.grpcServer.Stop()

	errAuth := a.authConn.Close()

	errDB := a.db.Close()

	errLog := a.logger.Sync()
	if errLog != nil {
		log.Printf("ERROR: failed to sync logger: %v", errLog)
	}

	if errHTTP != nil || errDB != nil || errAuth != nil {
		return fmt.Errorf("shutdown errors: http=%v, db=%v, auth_conn=%v", errHTTP, errDB, errAuth)
	}

	a.logger.Infof("Application stopped gracefully.")
	return nil
}
