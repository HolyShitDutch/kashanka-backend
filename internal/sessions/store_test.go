package sessions_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/moroz/kashanka/internal/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStore(t *testing.T) {
	var examples = []struct {
		key   []byte
		valid bool
	}{
		{bytes.Repeat([]byte{0xAA}, 32), true},
		{bytes.Repeat([]byte{0xFF}, 32), true},
		{bytes.Repeat([]byte{0x00}, 32), false},
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

func TestSessionStoreEncDecCookie(t *testing.T) {
	key := bytes.Repeat([]byte{0xAA}, 32)
	store, err := sessions.NewStore(key)
	require.NoError(t, err)
	require.NotNil(t, store)

	payload := []byte(`{"test":123}`)
	cookie, err := store.EncryptCookie(payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, cookie)

	decrypted, err := store.DecryptCookie(cookie, time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, payload, decrypted)
}

func TestSessionStoreEncDecSession(t *testing.T) {
	key := bytes.Repeat([]byte{0xAA}, 32)
	store, err := sessions.NewStore(key)
	require.NoError(t, err)
	require.NotNil(t, store)

	value := []byte{1, 2, 3}

	session := sessions.New()
	session.Set("test", value)

	cookie, err := store.EncryptSession(session)
	assert.NoError(t, err)
	assert.NotEmpty(t, cookie)

	decrypted, err := store.DecryptSession(cookie, time.Hour)
	assert.NoError(t, err)
	actual, ok := decrypted.Get[[]byte]("test")
	assert.Truef(t, ok, `Expected "test" to be present within the decoded session payload`)
	assert.Equal(t, value, actual)
}
