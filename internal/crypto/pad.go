package crypto

import (
	"errors"
	"math/bits"
)

// Pad appends 0x80 and zero bytes up to size. size must be > len(msg).
func Pad(msg []byte, size int) ([]byte, error) {
	if size <= len(msg) {
		return nil, errors.New("Pad: size too small")
	}
	out := make([]byte, size)
	copy(out, msg)
	out[len(msg)] = 0x80
	return out, nil
}

func Unpad(b []byte) ([]byte, error) {
	for i := len(b) - 1; i >= 0; i-- {
		switch b[i] {
		case 0x00:
			continue
		case 0x80:
			return b[:i], nil
		default:
			return nil, errors.New("invalid padding")
		}
	}
	return nil, errors.New("invalid padding")
}

// CalculatePaddingLength returns the padded length for a plaintext of length L.
// Uses the Padmé algorithm (Anakin's special)
func CalculatePaddingLength(L int) int {
	if L < 2 {
		return L
	}
	E := bits.Len(uint(L)) - 1 // floor(log2 L)
	S := bits.Len(uint(E))     // floor(log2 E) + 1
	mask := (1 << (E - S)) - 1
	return (L + mask) &^ mask
}
