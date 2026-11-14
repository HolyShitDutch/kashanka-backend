package repositories_test

import (
	"fmt"
	"log/slog"
	"net/url"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moroz/kashanka/config"
	"github.com/moroz/kashanka/internal/migrate"
	"github.com/stretchr/testify/assert"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelError)
}

func freshDatabase(t *testing.T) *pgxpool.Pool {
	t.Helper()

	parsed, err := url.Parse(config.TestDatabaseUrl)
	if err != nil {
		t.Fatal("TEST_DATABASE_URL is not a valid URL:", err)
	}

	dbName := strings.TrimLeft(parsed.Path, "/")

	if !strings.Contains(dbName, "_test") {
		t.Fatalf(`Test database name must include "_test", got: %v`, dbName)
	}

	pool, err := pgxpool.New(t.Context(), config.TestDatabaseUrl)
	if err != nil {
		t.Fatal("Failed to connect to test database", err)
	}

	if _, err := pool.Exec(t.Context(), `drop schema public cascade; create schema public;`); err != nil {
		t.Fatal(err)
	}

	if err := migrate.Up(t.Context(), pool); err != nil {
		t.Fatalf("Failed to migrate test database: %s", err)
	}

	return pool
}

func TestCardRepository_ListCards(t *testing.T) {
	db := freshDatabase(t)

	var version string
	err := db.QueryRow(t.Context(), "select version()").Scan(&version)
	assert.NoError(t, err)
	fmt.Println(version)
}
