package main

import (
	"context"
	"flag"
	"log"
	"spotify/internal/app"
)

func main() {
	configPath := parseFlags()

	ctx := context.Background()

	app, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("application init failed: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("application run failed: %v", err)
	}
}

func parseFlags() string {
	var configPath string
	flag.StringVar(&configPath, "f", "config", "path to config directory")
	flag.Parse()
	return configPath
}
