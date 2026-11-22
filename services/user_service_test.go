package services

import (
	"testing"

	"github.com/HolyShitDutch/kashanka/internal/tests"
)

func TestUserService_RegisterUserWithCredential(t *testing.T) {
	t.Parallel()

	db := tests.Database(t)
	_ = db
}
