package types

import (
	"bytes"
	"crypto/rand"

	"github.com/go-webauthn/webauthn/webauthn"
)

type webAuthnUser struct {
	id          []byte
	displayName string
}

var _ webauthn.User = &webAuthnUser{}

func NewWebAuthnUser() webauthn.User {
	id := make([]byte, 64)
	rand.Read(id)
	return &webAuthnUser{
		id: id, displayName: "Test User",
	}
}

func (u *webAuthnUser) WebAuthnID() []byte {
	return bytes.Clone(u.id)
}

func (u *webAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return nil
}

func (u *webAuthnUser) WebAuthnName() string {
	return u.displayName
}

func (u *webAuthnUser) WebAuthnDisplayName() string {
	return u.displayName
}
