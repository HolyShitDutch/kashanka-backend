package tests

import (
	"testing"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/descope/virtualwebauthn"
	"github.com/go-webauthn/webauthn/webauthn"
)

func WebAuthnConfig(t *testing.T) *webauthn.WebAuthn {
	w, err := webauthn.New(&webauthn.Config{
		RPID:          config.RPID,
		RPDisplayName: config.RPDisplayName,
		RPOrigins:     config.RPOrigins,
	})
	if err != nil {
		t.Fatal(err)
	}
	return w
}

func RelyingParty() virtualwebauthn.RelyingParty {
	return virtualwebauthn.RelyingParty{
		ID:     config.RPID,
		Name:   config.RPDisplayName,
		Origin: config.RPOrigins[0],
	}
}

func Authenticator() virtualwebauthn.Authenticator {
	return virtualwebauthn.NewAuthenticatorWithOptions(virtualwebauthn.AuthenticatorOptions{
		Transports: []virtualwebauthn.Transport{virtualwebauthn.TransportUSB, virtualwebauthn.TransportInternal},
	})
}

func Credential() virtualwebauthn.Credential {
	return virtualwebauthn.NewCredential(virtualwebauthn.KeyTypeEC2)
}
