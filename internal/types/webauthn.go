package types

import (
	"bytes"
	"crypto/rand"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/moroz/kashanka/internal/names"
)

type webAuthnUser struct {
	id   []byte
	name string
}

var _ webauthn.User = &webAuthnUser{}

func NewWebAuthnUser(displayName string) webauthn.User {
	id := make([]byte, 64)
	rand.Read(id)

	displayName = strings.TrimSpace(displayName)
	if displayName == "" {
		displayName = names.GenerateName()
	}

	return &webAuthnUser{
		id:   id,
		name: displayName,
	}
}

func (u *webAuthnUser) WebAuthnID() []byte {
	return bytes.Clone(u.id)
}

func (u *webAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return nil
}

func (u *webAuthnUser) WebAuthnName() string {
	return u.name
}

func (u *webAuthnUser) WebAuthnDisplayName() string {
	return u.name
}
