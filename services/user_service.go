package services

import (
	"context"
	"encoding/json/v2"
	"fmt"

	"github.com/HolyShitDutch/kashanka/db/queries"
	"github.com/HolyShitDutch/kashanka/internal/types"
	"github.com/bincyber/go-sqlcrypter"
	"github.com/go-webauthn/webauthn/webauthn"
)

type UserService struct {
	db DB
}

func NewUserService(db DB) *UserService {
	return &UserService{db}
}

func (s *UserService) RegisterUserWithCredential(ctx context.Context, config *webauthn.WebAuthn, userData *types.WebAuthnUser, credential *webauthn.Credential) (*queries.User, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	user, err := queries.New(s.db).InsertUser(ctx, &queries.InsertUserParams{
		DisplayName: sqlcrypter.NewEncryptedBytes(userData.Name),
		WebauthnID:  userData.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("RegisterUserWithCredential: %w", err)
	}

	credentialJson, err := json.Marshal(credential)

	_, err = queries.New(s.db).InsertWebAuthnCredential(ctx, &queries.InsertWebAuthnCredentialParams{
		RpID:           config.Config.RPID,
		UserID:         user.ID,
		Kid:            credential.ID,
		CredentialJSON: sqlcrypter.NewEncryptedBytes(string(credentialJson)),
	})

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return user, nil
}
