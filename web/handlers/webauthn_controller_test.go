package handlers_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/descope/virtualwebauthn"
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

	rp := tests.RelyingParty()
	authenticator := virtualwebauthn.NewAuthenticatorWithOptions(virtualwebauthn.AuthenticatorOptions{
		Transports: []virtualwebauthn.Transport{virtualwebauthn.TransportUSB, virtualwebauthn.TransportInternal},
	})

	credential := virtualwebauthn.NewCredential(virtualwebauthn.KeyTypeEC2)

	parsedAttestationOptions, err := virtualwebauthn.ParseAttestationOptions(string(response))
	require.NoError(t, err)

	attestationResponse := virtualwebauthn.CreateAttestationResponse(rp, authenticator, credential, *parsedAttestationOptions)
	fmt.Println(attestationResponse)
}
