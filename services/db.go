package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/moroz/kashanka/db/queries"
)

type DB interface {
	queries.DBTX
	Begin(ctx context.Context) (pgx.Tx, error)
}
