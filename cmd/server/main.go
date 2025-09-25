package main

import (
	"log"
	"spotify/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := app.NewConfig()
	app := app.NewApp(cfg)
	if err := app.Run(); err != nil {
		log.Fatalf("application run failed: %v", err)
	}
}
