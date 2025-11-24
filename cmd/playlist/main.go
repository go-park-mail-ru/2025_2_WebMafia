package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"spotify/microservices/playlist/app"
)

func main() {
	configPath := parseFlags()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	playlistApp, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("failed to init playlist app: %v", err)
	}

	if err := playlistApp.Run(ctx); err != nil {
		log.Printf("playlist app stopped with error: %v", err)
	}
}

func parseFlags() string {
	var configPath string
	flag.StringVar(&configPath, "f", "config", "path to config directory")
	flag.Parse()
	return configPath
}
