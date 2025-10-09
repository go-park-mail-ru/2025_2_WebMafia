package main

import (
	"log"
	"spotify/internal/app"
)

func main() {
	cfg := app.NewConfig()
	app, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("application init failed: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("application run failed: %v", err)
	}
}
