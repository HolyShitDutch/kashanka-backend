package services

import (
	"context"
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/moroz/kashanka/db/queries"
	"github.com/moroz/kashanka/internal/types"
)

type UserService struct {
	db DB
}

func NewUserService(db DB) *UserService {
	return &UserService{db}
}

func (s *UserService) RegisterUserWithCredential(ctx context.Context, userData *types.WebAuthnUser, credential *webauthn.Credential) (*queries.User, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)
}
