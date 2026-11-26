package types

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/db/queries"
	"github.com/HolyShitDutch/kashanka/internal/sessions"
)

type RequestContext struct {
	session      sessions.Session
	sessionStore sessions.Store
	User         *queries.User
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

func (c *RequestContext) SaveSession(w http.ResponseWriter) error {
	cookie, err := c.sessionStore.EncryptSession(c.session)
	if err != nil {
		return fmt.Errorf("failed to encode session cookie: %w", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     config.SessionCookieName,
		Value:    cookie,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
	return nil
}
