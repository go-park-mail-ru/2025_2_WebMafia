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
	userService "spotify/internal/user/service"

	"spotify/internal/middleware"
	"spotify/internal/router"
	"spotify/pkg/jwtmanager"
	"spotify/pkg/logger"
	"spotify/pkg/postgres"
)

type App struct {
	server *http.Server
	cfg    *Config
	db     *sql.DB
	logger logger.Logger
}

func NewApp(cfg *Config) (*App, error) {
	log, err := logger.New(cfg.Logger.Level, cfg.Logger.Mode)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}
	log.Infow("Logger initialized")

	db, err := postgres.New(context.Background(), cfg.DB)
	if err != nil {
		log.Errorw("failed to connect to db", "error", err)
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	log.Infow("Database connection")

	userRepository := userRepo.NewUserRepository(db)
	artistRepository := artistRepo.New(db)
	albumRepository := albumRepo.New(db)
	trackRepository := trackRepo.New(db)

	userSvc := userService.NewUserService(userRepository)
	artistSvc := artistService.New(artistRepository)

	albumSvc := albumService.New(albumRepository, artistSvc)
	trackSvc := trackService.New(trackRepository, albumSvc, artistSvc)

	jwtManager := jwtmanager.NewManager(cfg.JWTSecretKey, cfg.AccessTokenTTL)
	authMiddleware := middleware.NewAuthMiddleware(jwtManager)

	userHandler := userDelivery.NewHandler(userSvc, jwtManager)
	artistHandler := artistDelivery.NewHandler(artistSvc)
	albumHandler := albumDelivery.NewHandler(albumSvc)
	trackHandler := trackDelivery.NewHandler(trackSvc)

	handlers := router.AppHandlers{
		UserHandler:   userHandler,
		ArtistHandler: artistHandler,
		AlbumHandler:  albumHandler,
		TrackHandler:  trackHandler,
	}

	muxRouter := router.NewRouter(log, handlers, authMiddleware, cfg.CORS)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      muxRouter,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
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
		a.logger.Infow("server is starting", "port", a.server.Addr)
		serverErrors <- a.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			a.logger.Errorw("server error", "error", err)
			return fmt.Errorf("server error: %w", err)
		}
	case <-ctx.Done():
		a.logger.Infow("shutting down server gracefully...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
		defer cancel()

		if err := a.server.Shutdown(shutdownCtx); err != nil {
			a.logger.Errorw("server forced to shutdown", "error", err)
			return fmt.Errorf("server forced to shutdown: %w", err)
		}

	}

	a.logger.Infow("server exiting")
	return nil
}
