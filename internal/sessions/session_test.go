package sessions_test

import (
	"encoding/gob"
	"testing"
	"time"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/HolyShitDutch/kashanka/internal/sessions"
	"github.com/HolyShitDutch/kashanka/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func roundTripTest[T any](value T) func(t *testing.T) {
	return func(t *testing.T) {
		gob.Register(value)

		store := tests.SessionStore(t)

		session := sessions.New()
		session.Set("key", value)
		encrypted, err := store.EncryptSession(session)
		require.NoError(t, err)

		decrypted, err := store.DecryptSession(encrypted, time.Hour)
		require.NoError(t, err)

		actual, ok := decrypted.Get[T]("key")
		assert.True(t, ok)
		assert.Equal(t, value, actual)
	}
}

func TestSessionEncDec(t *testing.T) {
	var examples = []struct {
		name  string
		value any
	}{
		{"int", int(42)},
		{"uint", uint(42)},
		{"negative int", int(-42)},
		{"string", "I am a string."},
		{"[]byte", []byte{1, 2, 3}},
		{"float64", float64(0.1)},
		{"float32", float32(0.2)},
		{"inline struct", struct{ Test string }{"test"}},
		{"webauthn session", &webauthn.SessionData{}},
	}

	for _, example := range examples {
		t.Run(example.name, roundTripTest(example.value))
	}
}
