package main

import (
	"log"
	"spotify/internal/app"
)

func main() {
	configPath := "config"
	app, err := app.NewApp(configPath)
	if err != nil {
		log.Fatalf("application init failed: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("application run failed: %v", err)
	}
}
