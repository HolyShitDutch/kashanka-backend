package config

import "github.com/go-webauthn/webauthn/webauthn"

func WebauthnConfig(overrides ...func(*webauthn.Config) error) (*webauthn.WebAuthn, error) {
	cfg := &webauthn.Config{
		RPID:          RPID,
		RPDisplayName: RPDisplayName,
		RPOrigins:     RPOrigins,
	}

	for _, fn := range overrides {
		if err := fn(cfg); err != nil {
			return nil, err
		}
	}

	w, err := webauthn.New(cfg)
	if err != nil {
		return nil, err
	}

	return w, nil
}
