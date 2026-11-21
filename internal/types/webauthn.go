package types

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/HolyShitDutch/kashanka/internal/names"
)

type WebAuthnUser struct {
	ID   []byte
	Name string
}

func init() {
	gob.Register(&WebAuthnUser{})
}

var _ webauthn.User = &WebAuthnUser{}

func NewWebAuthnUser(displayName string) *WebAuthnUser {
	id := make([]byte, 64)
	rand.Read(id)

	displayName = strings.TrimSpace(displayName)
	if displayName == "" {
		displayName = names.GenerateName()
	}

	return &WebAuthnUser{
		ID:   id,
		Name: displayName,
	}
}

func (u *WebAuthnUser) WebAuthnID() []byte {
	return bytes.Clone(u.ID)
}

func (u *WebAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return nil
}

func (u *WebAuthnUser) WebAuthnName() string {
	return u.Name
}

func (u *WebAuthnUser) WebAuthnDisplayName() string {
	return u.Name
}
