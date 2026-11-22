package handlers_test

import (
	"bytes"
	"encoding/json/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/HolyShitDutch/kashanka/db/queries"
	"github.com/HolyShitDutch/kashanka/internal/requests"
	"github.com/HolyShitDutch/kashanka/internal/tests"
	"github.com/HolyShitDutch/kashanka/web/handlers"
	"github.com/descope/virtualwebauthn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebauthnController_RegisterStart(t *testing.T) {
	t.Parallel()

	db := tests.Database(t)
	store := tests.SessionStore(t)
	router, err := handlers.Router(&handlers.RouterDeps{
		DB:           db,
		SessionStore: store,
	})
	require.NoError(t, err)
	srv := httptest.NewServer(router)
	defer srv.Close()

	countBefore, err := queries.New(db).CountUsers(t.Context())
	require.NoError(t, err)
	assert.Zero(t, countBefore)

	origin, err := url.Parse(srv.URL)
	require.NoError(t, err)

	client := tests.ClientWithSession(t, origin, store, nil)

	var body bytes.Buffer
	json.MarshalWrite(&body, &requests.WebauthnRegisterStartRequest{
		DisplayName: "John Smith",
	})
	req, _ := http.NewRequest("POST", srv.URL+"/webauthn/register/start", &body)
	req.Header.Add("Content-Type", "application/json")
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

	req, _ = http.NewRequest("POST", srv.URL+"/webauthn/register/finalize", bytes.NewBufferString(attestationResponse))
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)

	countAfter, err := queries.New(db).CountUsers(t.Context())
	require.NoError(t, err)
	assert.EqualValues(t, 1, countAfter)
}
