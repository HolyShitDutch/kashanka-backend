package sessions_test

import (
	"bytes"
	"testing"

	"github.com/moroz/kashanka/internal/sessions"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	var examples = []struct {
		key   []byte
		valid bool
	}{
		{bytes.Repeat([]byte{0xAA}, 32), true},
		{bytes.Repeat([]byte{0x00}, 32), true},
		{bytes.Repeat([]byte{0xFF}, 32), true},
		{bytes.Repeat([]byte{0xAA}, 33), false},
		{bytes.Repeat([]byte{0xAA}, 64), false},
		{bytes.Repeat([]byte{0xAA}, 121), false},
		{[]byte{}, false},
		{nil, false},
	}

	for _, example := range examples {
		actual, err := sessions.NewStore(example.key)

		if example.valid {
			assert.NoError(t, err)
			assert.NotNil(t, actual)
		} else {
			assert.Error(t, err)
			assert.Nil(t, actual)
		}
	}
}
