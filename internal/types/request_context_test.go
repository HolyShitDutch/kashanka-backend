package types_test

import (
	"bytes"
	"testing"

	"github.com/moroz/kashanka/internal/sessions"
	"github.com/moroz/kashanka/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestContextSession(t *testing.T) {
	key := bytes.Repeat([]byte{0xAA}, 32)
	store, err := sessions.NewStore(key)
	require.NoError(t, err)
	require.NotNil(t, store)

	ctx := types.NewRequestContext(store)
	ctx.Session().Set("a", 123)
	ctx.Session().Set("b", "456")

	actualA, _ := ctx.Session().Get[int]("a")
	assert.Equal(t, 123, actualA)

	actualB, _ := ctx.Session().Get[string]("b")
	assert.Equal(t, "456", actualB)
}
