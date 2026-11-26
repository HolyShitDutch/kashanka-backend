package tests

import (
	"context"

	"github.com/HolyShitDutch/kashanka/db/queries"
	"github.com/HolyShitDutch/kashanka/internal/names"
	"github.com/HolyShitDutch/kashanka/services"
	"github.com/bincyber/go-sqlcrypter"
)

func User(db services.DB, ctx context.Context, overrides ...func(*queries.InsertUserParams) error) (*queries.User, error) {
	baseParams := &queries.InsertUserParams{
		DisplayName: sqlcrypter.NewEncryptedBytes(names.GenerateName()),
	}

	for _, fn := range overrides {
		if err := fn(baseParams); err != nil {
			return nil, err
		}
	}

	return queries.New(db).InsertUser(ctx, baseParams)
}
