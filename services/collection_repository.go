package services

import (
	"context"

	"github.com/moroz/kashanka/db/queries"
	"github.com/moroz/kashanka/internal/types"
)

type CollectionRepository struct {
	db DB
}

func NewCollectionRepository(db DB) *CollectionRepository {
	return &CollectionRepository{db}
}

func (s *CollectionRepository) GetCollectionWithCardsBySlug(ctx context.Context, slug string) (*types.CollectionDTO, error) {
	collection, err := queries.New(s.db).GetCollectionBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	cards, err := queries.New(s.db).ListCardsForCollection(ctx, collection.ID)
	if err != nil {
		return nil, err
	}

	return &types.CollectionDTO{
		Collection: collection,
		Cards:      cards,
	}, nil
}
