package sessions

import (
	"encoding/gob"

	"github.com/go-webauthn/webauthn/webauthn"
)

func init() {
	gob.Register(&Session{})
	gob.Register(&webauthn.SessionData{})
}
