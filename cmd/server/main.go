package main

import (
	"spotify/internal/app"
)

func main() {
	addr := ":8080"
	app := app.New(addr)
	app.Run()
}
