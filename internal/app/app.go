package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	albumDelivery "spotify/internal/album/delivery/http"
	albumRepo "spotify/internal/album/repository/postgres"
	albumService "spotify/internal/album/service"
	"syscall"

	artistDelivery "spotify/internal/artist/delivery/http"
	artistRepo "spotify/internal/artist/repository/postgres"
	artistService "spotify/internal/artist/service"

	"spotify/internal/router"
	trackDelivery "spotify/internal/track/delivery/http"
	trackRepo "spotify/internal/track/repository/postgres"
	trackService "spotify/internal/track/service"
	"spotify/pkg/postgres"
)

type App struct {
	server *http.Server
	cfg    *Config
	db     *sql.DB
}

func NewApp(cfg *Config) (*App, error) {
	db, err := postgres.New(context.Background(), cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	artistRepository := artistRepo.New(db)
	albumRepository := albumRepo.New(db)
	trackRepository := trackRepo.New(db)

	artistSvc := artistService.New(artistRepository)
	albumSvc := albumService.New(albumRepository)
	trackSvc := trackService.New(trackRepository)

	artistHandler := artistDelivery.NewHandler(artistSvc)
	albumHandler := albumDelivery.NewHandler(albumSvc)
	trackHandler := trackDelivery.NewHandler(trackSvc)

	// jwtManager := jwtmanager.NewManager(cfg.JWTSecretKey, cfg.AccessTokenTTL)
	muxRouter := router.NewRouter(
		trackHandler,
		artistHandler,
		albumHandler,
		cfg.CORS,
	)

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
	}, nil
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	serverErrors := make(chan error, 1)
	go func() {
		fmt.Println("Server is starting on port ", a.server.Addr)
		serverErrors <- a.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", err)
		}
	case <-ctx.Done():
		fmt.Println("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
		defer cancel()

		if err := a.server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server forced to shutdown: %w", err)
		}

	}

	fmt.Println("Server exiting")
	return nil
}
