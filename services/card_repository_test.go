package services_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/moroz/kashanka/internal/tests"
	"github.com/stretchr/testify/assert"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelError)
}

func TestCardRepository_ListCards(t *testing.T) {
	db := tests.Database(t)

	var version string
	err := db.QueryRow(t.Context(), "select version()").Scan(&version)
	assert.NoError(t, err)
	fmt.Println(version)
}
