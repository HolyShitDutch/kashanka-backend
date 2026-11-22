package tests

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/internal/crypto"
	"github.com/HolyShitDutch/kashanka/internal/migrate"
	"github.com/bincyber/go-sqlcrypter"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type testDBPool struct {
	admin    *pgxpool.Pool
	base     *pgxpool.Config
	template string
	counter  atomic.Int64
}

var sharedDBPool = sync.OnceValues(newTestDBPool)

func newTestDBPool() (*testDBPool, error) {
	ctx := context.Background()

	crypterer, err := crypto.NewEncryptionProvider(bytes.Repeat([]byte{0xAA}, 32))
	if err != nil {
		return nil, err
	}
	sqlcrypter.Init(crypterer)

	cfg, err := pgxpool.ParseConfig(config.TestDatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("TEST_DATABASE_URL is not a valid URL: %w", err)
	}

	if !strings.Contains(cfg.ConnConfig.Database, "_test") {
		return nil, fmt.Errorf(`Test database name must include "_test", got: %v`, cfg.ConnConfig.Database)
	}

	adminPool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to test database: %w", err)
	}

	d := &testDBPool{
		admin:    adminPool,
		base:     cfg,
		template: cfg.ConnConfig.Database + "_template",
	}
	if err := d.prepareTemplate(ctx); err != nil {
		adminPool.Close()
		return nil, err
	}

	return d, nil
}

func (d *testDBPool) prepareTemplate(ctx context.Context) error {
	return d.withAdvisoryLock(ctx, func(conn *pgxpool.Conn) error {
		exists, err := dbExists(ctx, conn, d.template)
		if err != nil {
			return fmt.Errorf("failed to check database existence: %w", err)
		}

		if !exists {
			if _, err := conn.Exec(ctx, `create database `+quoteIdentifier(d.template)); err != nil {
				return err
			}
		}

		templatePool, err := d.connectToDb(ctx, d.template)
		if err != nil {
			return err
		}
		defer templatePool.Close()

		if _, err := templatePool.Exec(ctx, `drop schema public cascade; create schema public`); err != nil {
			return fmt.Errorf("failed to clean template database: %w", err)
		}

		if err := migrate.Up(ctx, templatePool); err != nil {
			return fmt.Errorf("failed to migrate template database: %w", err)
		}
		return nil
	})
}

func (d *testDBPool) withAdvisoryLock(ctx context.Context, fn func(*pgxpool.Conn) error) error {
	const templateLockKey = 0x12345678

	conn, err := d.admin.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	if _, err := conn.Exec(ctx, "select pg_advisory_lock($1)", templateLockKey); err != nil {
		return err
	}
	defer conn.Exec(context.Background(), "select pg_advisory_unlock($1)", templateLockKey)

	return fn(conn)
}

func quoteIdentifier(s string) string {
	return pgx.Identifier{s}.Sanitize()
}

func Database(t *testing.T) *pgxpool.Pool {
	t.Helper()

	pool, err := sharedDBPool()
	if err != nil {
		t.Fatal(err)
	}

	return pool.clone(t)
}

func (d *testDBPool) clone(t *testing.T) *pgxpool.Pool {
	t.Helper()
	ctx := context.Background()

	name := fmt.Sprintf("%s_%d_%d", d.template, os.Getpid(), d.counter.Add(1))

	err := d.withAdvisoryLock(ctx, func(c *pgxpool.Conn) error {
		return createDatabaseFromTemplate(ctx, c, name, d.template)
	})
	if err != nil {
		t.Fatalf("create test database %s: %v", name, err)
	}

	t.Cleanup(func() {
		if err := dropTestDatabase(context.Background(), d.admin, name); err != nil {
			t.Logf("drop test database %s: %v", name, err)
		}
	})

	pool, err := d.connectToDb(ctx, name)
	if err != nil {
		t.Fatal("connect to test database:", err)
	}
	t.Cleanup(pool.Close)

	return pool
}

func (d *testDBPool) connectToDb(ctx context.Context, name string) (*pgxpool.Pool, error) {
	cfg := d.base.Copy()
	cfg.ConnConfig.Database = name
	return pgxpool.NewWithConfig(ctx, cfg)
}

func dbExists(ctx context.Context, c *pgxpool.Conn, name string) (bool, error) {
	var result bool
	err := c.QueryRow(
		ctx,
		"select exists (select from pg_database where datname = $1)",
		name,
	).Scan(&result)
	return result, err
}

func createDatabaseFromTemplate(ctx context.Context, c *pgxpool.Conn, name, template string) error {
	sql := fmt.Sprintf("create database %s template %s", quoteIdentifier(name), quoteIdentifier(template))
	_, err := c.Exec(ctx, sql)
	return err
}

func dropTestDatabase(ctx context.Context, adminPool *pgxpool.Pool, name string) error {
	sql := fmt.Sprintf("drop database if exists %s with (force)", quoteIdentifier(name))
	_, err := adminPool.Exec(ctx, sql)
	return err
}
