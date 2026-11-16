package sessions

import (
	"github.com/fxamacker/cbor/v2"
)

type Session map[string]any

func New() Session {
	return make(Session)
}

func (c Session) Get[T any](key string) (T, bool) {
	value, ok := c[key].(T)
	return value, ok
}

func (c Session) Set(key string, value any) {
	c[key] = value
}

func (c Session) Serialize() ([]byte, error) {
	return cbor.Marshal(c)
}

func decodeSession(payload []byte) (Session, error) {
	var target Session
	err := cbor.Unmarshal(payload, &target)
	return target, err
}
