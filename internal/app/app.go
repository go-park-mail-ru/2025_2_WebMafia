package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"spotify/internal/handler"
	"spotify/internal/router"
)

type App struct {
	server   *http.Server
	handlers *handler.Handlers
	cfg      *Config
}

func NewApp(cfg *Config) *App {
	handlers := handler.NewHandler()
	muxRouter := router.NewRouter(handlers, cfg.CORS)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      muxRouter,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &App{
		server:   server,
		handlers: handlers,
		cfg:      cfg,
	}
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
