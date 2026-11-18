package tests

import (
	"bytes"
	"testing"

	"github.com/moroz/kashanka/internal/sessions"
)

func SessionStore(t *testing.T) sessions.Store {
	t.Helper()
	key := bytes.Repeat([]byte{0xAA}, 32)
	store, err := sessions.NewStore(key)
	if err != nil {
		t.Fatal(err)
	}
	return store
}
