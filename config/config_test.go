package config_test

import (
	"bytes"
	"testing"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/stretchr/testify/assert"
)

func TestDeriveKey(t *testing.T) {
	t.Parallel()

	secretKeyBase := bytes.Repeat([]byte{0xAA}, 32)
	salt := bytes.Repeat([]byte{0x5C}, 32)
	context := "github.com/HolyShitDutch/kashanka-backend 2026-11-21 15:20:00 TestDeriveKey"

	for _, l := range []int{32, 64, 128, 1024} {
		actual := config.MustDeriveKey(secretKeyBase, salt, context, l)
		assert.Len(t, actual, l)

		var nonZeroFound bool
		for _, b := range actual {
			if b != 0 {
				nonZeroFound = true
				break
			}
		}

		assert.Truef(t, nonZeroFound, "Key of length %v should not be all zeroes", l)
	}
}
