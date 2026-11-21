package tests

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/internal/sessions"
)

func ClientWithSession(t *testing.T, origin *url.URL, store sessions.Store, payload sessions.Session) *http.Client {
	t.Helper()

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	var cookie string

	if payload != nil {
		sessionCookie, err := store.EncryptSession(payload)
		if err != nil {
			t.Fatalf("Failed to encode session as cookie: %v", err)
		}
		cookie = sessionCookie
	}

	if cookie != "" {
		jar.SetCookies(origin, []*http.Cookie{
			{
				Name:     config.SessionCookieName,
				Value:    cookie,
				HttpOnly: true,
				Secure:   true,
			},
		})
	}

	return &http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}
