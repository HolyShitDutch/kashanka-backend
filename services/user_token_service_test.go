package services_test

import (
	"testing"

	"github.com/HolyShitDutch/kashanka/internal/tests"
	"github.com/HolyShitDutch/kashanka/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserTokenService(t *testing.T) {
	db := tests.Database(t)

	user, err := tests.User(db, t.Context())
	require.NoError(t, err)

	srv := services.NewUserTokenService(db)

	token, err := srv.CreateAccessTokenForUser(t.Context(), user)
	require.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, token.PlaintextToken, token.Token)

	actual, err := srv.GetUserByAccessToken(t.Context(), token.PlaintextToken)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, actual.ID)
}
