package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"spotify/microservices/catalog/app"
)

func main() {
	configPath := parseFlags()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	catalog, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("application init failed: %v", err)
	}

	if err := catalog.Run(ctx); err != nil {
		log.Printf("application run failed: %v", err)
	}
}

func parseFlags() string {
	var configPath string
	flag.StringVar(&configPath, "f", "config", "path to config directory")
	flag.Parse()
	return configPath
}
