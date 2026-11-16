package types

import (
	"time"

	"github.com/moroz/kashanka/internal/sessions"
)

type RequestContext struct {
	session      sessions.Session
	sessionStore sessions.Store
}

func NewRequestContext(store sessions.Store) *RequestContext {
	return &RequestContext{
		session:      make(sessions.Session),
		sessionStore: store,
	}
}

func (c *RequestContext) Session() sessions.Session {
	return c.session
}

func (c *RequestContext) DecodeSessionFromCookie(cookie string, maxAge time.Duration) error {
	session, err := c.sessionStore.DecryptSession(cookie, maxAge)
	if err != nil {
		return err
	}
	c.session = session
	return nil
}
