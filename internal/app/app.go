package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"spotify/internal/handler"
	"spotify/internal/router"
	"syscall"
	"time"
)

type App struct {
	server   *http.Server
	handlers *handler.Handlers
}

func New(addr string) *App {
	handlers := handler.NewHandler()
	muxRouter := router.NewRouter(handlers)

	server := &http.Server{
		Addr:    addr,
		Handler: muxRouter,
	}

	return &App{
		server:   server,
		handlers: handlers,
	}
}

func (a *App) Run() {
	go func() {
		fmt.Println("Server is starting on port ", a.server.Addr)
		if err := a.server.ListenAndServe(); err != nil {
			log.Fatalf("could not start server: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")
}
