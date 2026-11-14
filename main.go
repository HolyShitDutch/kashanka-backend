package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moroz/kashanka/config"
	"github.com/moroz/kashanka/handlers"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := pgxpool.New(context.Background(), config.DatabaseUrl)
	if err != nil {
		logger.Error("Failed to connect to the database", "error", err)
		os.Exit(1)
	}

	r := handlers.Router(&handlers.RouterDeps{
		DB:     db,
		Logger: logger,
	})

	if err := r.Start(config.AppHostPort); err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}
