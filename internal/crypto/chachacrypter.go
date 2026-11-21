package crypto

import (
	"crypto/cipher"
	"fmt"
	"io"

	"github.com/bincyber/go-sqlcrypter"
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/chacha20poly1305"
)

// ChachaCrypter implements the sqlcrypter.Crypterer interface. The implementation in this file is mostly a blatant copy of https://github.com/bincyber/go-sqlcrypter/blob/master/providers/aesgcm/aes.go
type ChachaCrypter struct {
	keys []cipher.AEAD
}

var _ sqlcrypter.Crypterer = (*ChachaCrypter)(nil)

func NewEncryptionProvider(keys ...[]byte) (sqlcrypter.Crypterer, error) {
	var aeads []cipher.AEAD

	if len(keys) == 0 {
		return nil, fmt.Errorf("at least one key is required")
	}

	for i, key := range keys {
		aead, err := chacha20poly1305.NewX(key)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize an AEAD at index %v: %w", i, err)
		}

		aeads = append(aeads, aead)
	}

	return &ChachaCrypter{
		keys: aeads,
	}, nil
}

func (c *ChachaCrypter) currentKey() cipher.AEAD {
	return c.keys[0]
}

func (c *ChachaCrypter) Encrypt(w io.Writer, r io.Reader) error {
	src, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read from io.Reader: %w", err)
	}

	nonce, err := sqlcrypter.GenerateBytes(chacha20.NonceSizeX)
	if err != nil {
		return fmt.Errorf("failed to generate %v-byte random nonce: %w", chacha20.NonceSizeX, err)
	}

	padded, err := Pad(src, CalculatePaddingLength(len(src)))
	if err != nil {
		return err // unlikely
	}

	// appends ciphertext and authentication tag after the nonce
	result := c.currentKey().Seal(nonce, nonce, padded, nil)

	if _, err := w.Write(result); err != nil {
		return err
	}

	return nil
}

func (c *ChachaCrypter) Decrypt(w io.Writer, r io.Reader) error {
	src, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read from io.Reader: %w", err)
	}

	minLen := c.currentKey().Overhead() + chacha20.NonceSizeX
	if len(src) < minLen {
		return fmt.Errorf("input is corrupt: expected at least %v bytes of input, got %v", minLen, len(src))
	}

	// First 24 bytes is the nonce, followed by the ciphertext
	nonce := src[:chacha20.NonceSizeX]
	ciphertext := src[chacha20.NonceSizeX:]

	var (
		decrypted []byte
		success   bool
	)

	for _, key := range c.keys {
		decrypted, err = key.Open(nil, nonce, ciphertext, nil)
		if err == nil {
			success = true
			break
		}
	}

	if !success {
		return fmt.Errorf("failed to decrypt message: %w", err)
	}

	unpadded, err := Unpad(decrypted)
	if err != nil {
		return fmt.Errorf("failed to unpad message: %w", err)
	}
	if _, err := w.Write(unpadded); err != nil {
		return fmt.Errorf("failed to write decrypted text: %w", err)
	}

	return nil
}
