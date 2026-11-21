package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/internal/sessions"
	"github.com/HolyShitDutch/kashanka/web/handlers"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := pgxpool.New(context.Background(), config.DatabaseUrl)
	if err != nil {
		logger.Error("Failed to connect to the database", "error", err)
		os.Exit(1)
	}

	sessionStore, err := sessions.NewStore(config.SessionEncryptionKey)
	if err != nil {
		logger.Error("Failed to initialize session store", "error", err)
		os.Exit(1)
	}

	r, err := handlers.Router(&handlers.RouterDeps{
		DB:           db,
		Logger:       logger,
		SessionStore: sessionStore,
	})
	if err != nil {
		panic(err)
	}

	if err := r.Start(config.AppHostPort); err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}
