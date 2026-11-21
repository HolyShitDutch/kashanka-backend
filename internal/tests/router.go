package tests

import (
	"log/slog"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/moroz/kashanka/internal/sessions"
	"github.com/moroz/kashanka/web/handlers"
)

func Router(t *testing.T, db *pgxpool.Pool, store sessions.Store) *echo.Echo {
	t.Helper()

	logger := slog.New(slog.DiscardHandler)

	router, err := handlers.Router(&handlers.RouterDeps{
		DB:           db,
		Logger:       logger,
		SessionStore: store,
	})
	if err != nil {
		t.Fatal(err)
	}

	return router
}
