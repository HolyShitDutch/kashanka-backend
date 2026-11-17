package sessions_test

import (
	"testing"

	"github.com/moroz/kashanka/internal/sessions"
	"github.com/stretchr/testify/assert"
)

func roundTripTest[T any](value T) func(t *testing.T) {
	return func(t *testing.T) {
		session := sessions.New()
		session.Set("key", value)
		actual, ok := session.Get[T]("key")
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
		{"inline struct", struct{ test string }{"test"}},
	}

	for _, example := range examples {
		t.Run(example.name, roundTripTest(example.value))
	}
}
