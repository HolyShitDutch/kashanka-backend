package services

import (
	"context"
	"crypto/rand"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/db/queries"
	"github.com/HolyShitDutch/kashanka/internal/types"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserTokenService struct {
	db DB
}

func NewUserTokenService(db DB) *UserTokenService {
	return &UserTokenService{db}
}

func (s *UserTokenService) CreateAccessTokenForUser(ctx context.Context, user *queries.User) (*types.UserTokenDTO, error) {
	token, err := s.generateToken()
	if err != nil {
		return nil, err
	}

	// access tokens can be stored in plaintext because they are only ever accepted in encrypted
	// cookie form
	userToken, err := queries.New(s.db).InsertUserToken(ctx, &queries.InsertUserTokenParams{
		Context: queries.UserTokenContextAccess,
		Token:   token,
		UserID:  user.ID,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserTokenDTO{
		UserToken:      userToken,
		PlaintextToken: token,
	}, nil
}

func (s *UserTokenService) generateToken() ([]byte, error) {
	out := make([]byte, 32)
	_, err := rand.Read(out)
	return out, err
}

func (s *UserTokenService) GetUserByAccessToken(ctx context.Context, token []byte) (*queries.User, error) {
	return queries.New(s.db).GetUserByAccessToken(ctx, &queries.GetUserByAccessTokenParams{
		Token: token,
		Maxage: pgtype.Interval{
			Valid:        true,
			Microseconds: config.SessionMaxAge.Microseconds(),
		},
	})
}
