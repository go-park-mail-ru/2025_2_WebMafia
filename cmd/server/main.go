package main

import (
	"log"
	"spotify/internal/app"
	"spotify/internal/config"
)

func main() {
	cfg := config.New()
	app := app.New(cfg)
	if err := app.Run(); err != nil {
		log.Fatalf("application run failed: %v", err)
	}
}
