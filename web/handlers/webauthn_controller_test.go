package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/moroz/kashanka/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebauthnController_RegisterStart(t *testing.T) {
	db := tests.Database(t)
	store := tests.SessionStore(t)
	router := tests.Router(t, db, store)
	srv := httptest.NewServer(router)
	defer srv.Close()

	origin, err := url.Parse(srv.URL)
	require.NoError(t, err)

	client := tests.ClientWithSession(t, origin, store, nil)

	req, _ := http.NewRequest("GET", srv.URL+"/webauthn/register/start", nil)
	resp, err := client.Do(req)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)

	response, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.NotEmpty(t, response)
}
