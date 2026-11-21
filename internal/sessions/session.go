package sessions

import (
	"bytes"
	"encoding/gob"
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
	var out bytes.Buffer
	err := gob.NewEncoder(&out).Encode(c)
	return out.Bytes(), err
}

func decodeSession(payload []byte) (Session, error) {
	var target Session
	in := bytes.NewBuffer(payload)
	err := gob.NewDecoder(in).Decode(&target)
	return target, err
}
