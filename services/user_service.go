package services

import (
	"context"
	"fmt"

	"github.com/HolyShitDutch/kashanka/db/queries"
	"github.com/HolyShitDutch/kashanka/internal/types"
	"github.com/go-webauthn/webauthn/webauthn"
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
