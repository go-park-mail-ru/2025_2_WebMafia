package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	albumDelivery "spotify/internal/album/delivery/http"
	albumRepo "spotify/internal/album/repository/postgres"
	albumService "spotify/internal/album/service"

	artistDelivery "spotify/internal/artist/delivery/http"
	artistRepo "spotify/internal/artist/repository/postgres"
	artistService "spotify/internal/artist/service"

	trackDelivery "spotify/internal/track/delivery/http"
	trackRepo "spotify/internal/track/repository/postgres"
	trackService "spotify/internal/track/service"

	userDelivery "spotify/internal/user/delivery/http"
	userRepo "spotify/internal/user/repository/postgres"
	storageRepo "spotify/internal/user/repository/storage"
	userService "spotify/internal/user/service"

	"spotify/internal/middleware"
	"spotify/internal/router"
	"spotify/pkg/csrfmanager"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/logger"
	"spotify/pkg/minio"
	"spotify/pkg/postgres"
)

type App struct {
	server *http.Server
	cfg    *Config
	db     *sql.DB
	logger logger.Logger
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	log, err := logger.New(cfg.App.Logger.Level, cfg.App.Logger.Mode)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}
	log.Infof("Logger initialized")

	pgConfig := postgres.Config{
		Host:            cfg.DB.Host,
		Port:            cfg.DB.Port,
		User:            cfg.DB.User,
		Password:        cfg.DB.Password,
		DBName:          cfg.DB.DBName,
		SSLMode:         cfg.DB.SSLMode,
		MaxOpenConns:    cfg.DB.MaxOpenConns,
		MaxIdleConns:    cfg.DB.MaxIdleConns,
		ConnMaxLifetime: cfg.DB.ConnMaxLifetime,
	}
	db, err := postgres.New(ctx, pgConfig)
	if err != nil {
		log.Errorf("failed to connect to db: %v", err)
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	log.Infof("Database connection")

	minioCfg := minio.Config{
		Endpoint:  cfg.Minio.Endpoint,
		AccessKey: cfg.Minio.AccessKey,
		SecretKey: cfg.Minio.SecretKey,
		UseSSL:    cfg.Minio.UseSSL,
	}
	minioClient, err := minio.New(minioCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init minio: %w", err)
	}
	avatarStorage := storageRepo.NewStorage(minioClient, cfg.Minio.Bucket)

	userRepository := userRepo.NewUserRepository(db)
	artistRepository := artistRepo.New(db)
	albumRepository := albumRepo.New(db)
	trackRepository := trackRepo.New(db)

	userSvc := userService.NewUserService(userRepository, avatarStorage)

	artistSvc := artistService.New(artistRepository, nil)

	albumSvc := albumService.New(albumRepository, artistSvc)
	trackSvc := trackService.New(trackRepository, albumSvc, artistSvc)

	artistSvc.SetTrackService(trackSvc)

	jwtManager := jwtmanager.NewManager(cfg.App.HTTP.Auth.JWT.SecretKey, cfg.App.HTTP.Auth.JWT.AccessTokenTTL)
	authMiddleware := middleware.NewAuthMiddleware(jwtManager)

	csrfManager := csrfmanager.NewManager(cfg.App.HTTP.Auth.CSRF.SecretKey, cfg.App.HTTP.Auth.CSRF.TokenTTL)
	csrfMiddleware := middleware.NewCSRFMiddleware(csrfManager)

	userHandler := userDelivery.NewHandler(userSvc, jwtManager, csrfManager, cfg.App.HTTP.AllowedAvatarTypes)
	artistHandler := artistDelivery.NewHandler(artistSvc)
	albumHandler := albumDelivery.NewHandler(albumSvc)
	trackHandler := trackDelivery.NewHandler(trackSvc)

	handlers := router.AppHandlers{
		UserHandler:   userHandler,
		ArtistHandler: artistHandler,
		AlbumHandler:  albumHandler,
		TrackHandler:  trackHandler,
	}

	CorsConfig := middleware.CORSConfig{
		AllowedOrigins:   cfg.App.HTTP.CORS.AllowedOrigins,
		AllowedMethods:   cfg.App.HTTP.CORS.AllowedMethods,
		AllowedHeaders:   cfg.App.HTTP.CORS.AllowedHeaders,
		AllowCredentials: cfg.App.HTTP.CORS.AllowCredentials,
	}
	muxRouter := router.NewRouter(log, handlers, authMiddleware, csrfMiddleware, CorsConfig)

	server := &http.Server{
		Addr:         ":" + cfg.App.HTTP.Port,
		Handler:      muxRouter,
		ReadTimeout:  cfg.App.HTTP.ReadTimeout,
		WriteTimeout: cfg.App.HTTP.WriteTimeout,
		IdleTimeout:  cfg.App.HTTP.IdleTimeout,
	}

	return &App{
		server: server,
		cfg:    cfg,
		db:     db,
		logger: log,
	}, nil
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		stop()

		if err := a.logger.Sync(); err != nil {
			log.Printf("ERROR: failed to sync logger: %v", err)
		}
	}()

	serverErrors := make(chan error, 1)
	go func() {
		a.logger.Infof("server is starting on port %s", a.server.Addr)
		serverErrors <- a.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			a.logger.Errorf("server error: %v", err)
			return fmt.Errorf("server error: %w", err)
		}
	case <-ctx.Done():
		a.logger.Infof("shutting down server gracefully...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), a.cfg.App.HTTP.ShutdownTimeout)
		defer cancel()

		if err := a.server.Shutdown(shutdownCtx); err != nil {
			a.logger.Errorf("server forced to shutdown: %v", err)
			return fmt.Errorf("server forced to shutdown: %w", err)
		}

	}

	a.logger.Infof("server exiting")
	return nil
}
