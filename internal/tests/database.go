package tests

import (
	"bytes"
	"net/url"
	"strings"
	"testing"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/internal/crypto"
	"github.com/HolyShitDutch/kashanka/internal/migrate"
	"github.com/bincyber/go-sqlcrypter"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Database(t *testing.T) *pgxpool.Pool {
	t.Helper()

	crypterer, err := crypto.NewEncryptionProvider(bytes.Repeat([]byte{0xAA}, 32))
	if err != nil {
		t.Fatal(err)
	}
	sqlcrypter.Init(crypterer)

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
