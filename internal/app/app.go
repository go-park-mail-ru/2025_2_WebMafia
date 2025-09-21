package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"spotify/internal/config"
	"spotify/internal/handler"
	"spotify/internal/router"
)

type App struct {
	server          *http.Server
	handlers        *handler.Handlers
	shutdownTimeout time.Duration
}

func New(cfg *config.Config) *App {
	handlers := handler.NewHandler()
	muxRouter := router.NewRouter(handlers)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      muxRouter,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return &App{
		server:          server,
		handlers:        handlers,
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (a *App) Run() error {
	serverErrors := make(chan error, 1)
	go func() {
		fmt.Println("Server is starting on port ", a.server.Addr)
		serverErrors <- a.server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", err)
		}
	case <-quit:
		fmt.Println("Shutting down server...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), a.shutdownTimeout)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	fmt.Println("Server exiting")
	return nil
}
