package main

import (
	"context"
	"flag"
	"log"
	"spotify/internal/app"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "f", "config", "path to config directory")
	flag.Parse()

	ctx := context.Background()

	app, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("application init failed: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("application run failed: %v", err)
	}
}
