package services

import (
	"context"
	"fmt"

	"github.com/moroz/kashanka/db/queries"
	"github.com/moroz/kashanka/internal/types"
)

type CardRepository struct {
	db queries.DBTX
}

func NewCardRepository(db queries.DBTX) *CardRepository {
	return &CardRepository{
		db,
	}
}

func (s *CardRepository) ListCards(ctx context.Context) ([]*types.CardDTO, error) {
	cards, err := queries.New(s.db).ListCards(ctx)
	if err != nil {
		return nil, fmt.Errorf("ListCards: %w", err)
	}

	result := make([]*types.CardDTO, len(cards))
	for i, card := range cards {
		dto, err := types.NewCardDTO(card)
		if err != nil {
			return nil, fmt.Errorf("ListCards: %w", err)
		}
		result[i] = dto
	}

	return result, nil
}
