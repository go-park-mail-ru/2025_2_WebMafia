package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"spotify/internal/middleware"
	"spotify/internal/server"
	"spotify/pkg/csrfmanager"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/logger"
	"spotify/pkg/minio"
	"spotify/pkg/postgres"
	pb "spotify/proto/auth"
	"sync"

	grpcDelivery "spotify/microservices/auth/delivery/grpc"
	httpDelivery "spotify/microservices/auth/delivery/http"
	userRepo "spotify/microservices/auth/repository/postgres"
	storageRepo "spotify/microservices/auth/repository/storage"
	userService "spotify/microservices/auth/service"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type App struct {
	cfg        *Config
	logger     logger.Logger
	db         *sql.DB
	httpServer *server.Server
	grpcServer *server.GRPCServer
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	appLogger, err := logger.New(cfg.Auth.Logger.Level, cfg.Auth.Logger.Mode)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}
	appLogger.Infof("Logger initialized for Auth service")

	db, err := postgres.New(ctx, cfg.DB)
	if err != nil {
		appLogger.Errorf("failed to connect to db: %v", err)
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	appLogger.Infof("Database connection established")

	minioClient, err := minio.New(cfg.Minio)
	if err != nil {
		return nil, fmt.Errorf("failed to init minio: %w", err)
	}
	appLogger.Infof("Minio connection established")

	avatarStorage := storageRepo.NewStorage(minioClient, "avatars")
	userRepository := userRepo.NewUserRepository(db)
	userSvc := userService.NewUserService(userRepository, avatarStorage)

	jwtManager := jwtmanager.NewManager(cfg.Auth.HTTP.Auth.JWT.SecretKey, cfg.Auth.HTTP.Auth.JWT.AccessTokenTTL)
	csrfManager := csrfmanager.NewManager(cfg.Auth.HTTP.Auth.CSRF.SecretKey, cfg.Auth.HTTP.Auth.CSRF.TokenTTL)

	authMiddleware := middleware.NewAuthMiddleware(jwtManager)
	csrfMiddleware := middleware.NewCSRFMiddleware(csrfManager)

	httpHandler := httpDelivery.NewHandler(userSvc, jwtManager, csrfManager, cfg.Auth.HTTP.AllowedAvatarTypes)

	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.RequestLoggerMiddleware(appLogger))
	api.Use(middleware.CORS(cfg.Auth.HTTP.CORS))

	public := api.PathPrefix("").Subrouter()
	protected := api.PathPrefix("").Subrouter()
	protected.Use(authMiddleware.AuthMiddleware)
	csrfProtected := protected.PathPrefix("").Subrouter()
	csrfProtected.Use(csrfMiddleware.CSRFMiddleware)

	httpHandler.RegisterRoutes(public, protected, csrfProtected)

	httpServer := server.NewHTTPServer(&cfg.Auth.HTTP, router, appLogger)

	grpcHandler := grpcDelivery.NewHandler(jwtManager, csrfManager)
	grpcServer := server.NewGRPCServer(&cfg.Auth.GRPC, appLogger, func(s *grpc.Server) {
		pb.RegisterAuthServiceServer(s, grpcHandler)
	})

	return &App{
		cfg:        cfg,
		logger:     appLogger,
		db:         db,
		httpServer: httpServer,
		grpcServer: grpcServer,
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

	a.logger.Infof("Auth microservice is running...")

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
	shutdownCtx, cancel := context.WithTimeout(context.Background(), a.cfg.Auth.HTTP.ShutdownTimeout)
	defer cancel()

	errHTTP := a.httpServer.Shutdown(shutdownCtx)
	a.grpcServer.Stop()

	errDB := a.db.Close()

	errLog := a.logger.Sync()
	if errLog != nil {
		log.Printf("ERROR: failed to sync logger: %v", errLog)
	}

	if errHTTP != nil || errDB != nil {
		return fmt.Errorf("shutdown errors: http=%v, db=%v", errHTTP, errDB)
	}

	a.logger.Infof("Application stopped gracefully.")
	return nil
}
