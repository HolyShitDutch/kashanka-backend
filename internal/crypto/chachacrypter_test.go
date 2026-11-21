package crypto_test

import (
	"bytes"
	"testing"

	"github.com/HolyShitDutch/kashanka/internal/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEncryptionProvider(t *testing.T) {
	var examples = []struct {
		keys  [][]byte
		valid bool
	}{
		{keys: [][]byte{bytes.Repeat([]byte{0xAA}, 32)}, valid: true},
		{keys: [][]byte{bytes.Repeat([]byte{0xAA}, 16)}, valid: false},
		{keys: [][]byte{bytes.Repeat([]byte{0xAA}, 33)}, valid: false},
	}

	for _, example := range examples {
		_, err := crypto.NewEncryptionProvider(example.keys...)
		if example.valid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestChachaCrypterRoundTrip(t *testing.T) {
	key := bytes.Repeat([]byte{0xAA}, 32)
	crypter, err := crypto.NewEncryptionProvider(key)
	require.NoError(t, err)

	var ciphertext, decrypted bytes.Buffer

	original := "Did you ever hear the Tragedy of Darth Plagueis the Wise?"
	err = crypter.Encrypt(&ciphertext, bytes.NewBufferString(original))
	assert.NoError(t, err)
	assert.NotEmpty(t, ciphertext.Bytes())

	err = crypter.Decrypt(&decrypted, &ciphertext)
	assert.NoError(t, err)
	assert.Equal(t, original, decrypted.String())
}

func TestChachaCrypterDecryptWithOlderKey(t *testing.T) {
	oldKey := bytes.Repeat([]byte{0x5C}, 32)
	crypter, err := crypto.NewEncryptionProvider(oldKey)
	require.NoError(t, err)

	var ciphertext, decrypted bytes.Buffer

	original := "I have a bad feeling about this..."
	err = crypter.Encrypt(&ciphertext, bytes.NewBufferString(original))
	assert.NoError(t, err)
	assert.NotEmpty(t, ciphertext.Bytes())

	newKey := bytes.Repeat([]byte{0xAA}, 32)
	crypterWithNewKeyOnly, err := crypto.NewEncryptionProvider(newKey)
	require.NoError(t, err)

	err = crypterWithNewKeyOnly.Decrypt(&decrypted, bytes.NewReader(ciphertext.Bytes()))
	assert.Error(t, err)
	assert.Empty(t, decrypted.Bytes())

	crypterWithBothKeys, err := crypto.NewEncryptionProvider(newKey, oldKey)
	require.NoError(t, err)

	err = crypterWithBothKeys.Decrypt(&decrypted, bytes.NewReader(ciphertext.Bytes()))
	assert.NoError(t, err)
	assert.Equal(t, original, decrypted.String())
}
