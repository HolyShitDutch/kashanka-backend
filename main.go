package main

import (
	"log/slog"

	"github.com/moroz/kashanka/config"
	"github.com/moroz/kashanka/handlers"
)

func main() {
	r := handlers.Router()

	slog.SetDefault(r.Logger)

	if err := r.Start(config.AppHostPort); err != nil {
		r.Logger.Error("Failed to start server", "error", err)
	}
}
