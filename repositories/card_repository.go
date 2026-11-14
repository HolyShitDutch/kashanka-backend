package repositories

import (
	"context"

	"github.com/moroz/kashanka/db/queries"
)

type CardRepository struct {
	db *queries.DBTX
}

func NewCardRepository(db *queries.DBTX) *CardRepository {
	return &CardRepository{
		db,
	}
}

func (s *CardRepository) ListCards(ctx context.Context) ([]*types.CardDTO, error) {

}
