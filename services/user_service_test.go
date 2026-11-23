package services_test

import (
	"bytes"
	"encoding/json/v2"
	"testing"

	"github.com/HolyShitDutch/kashanka/db/queries"
	"github.com/HolyShitDutch/kashanka/internal/names"
	"github.com/HolyShitDutch/kashanka/internal/tests"
	"github.com/HolyShitDutch/kashanka/internal/types"
	"github.com/HolyShitDutch/kashanka/services"
	"github.com/descope/virtualwebauthn"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_RegisterUserWithCredential(t *testing.T) {
	t.Parallel()

	db := tests.Database(t)
	srv := services.NewUserService(db)

	userCount, err := queries.New(db).CountUsers(t.Context())
	require.NoError(t, err)
	assert.Zero(t, userCount)

	userName := names.GenerateName()
	webauthnUser := types.NewWebAuthnUser(userName)

	w := tests.WebAuthnConfig(t)

	creation, sessionData, err := w.BeginMediatedRegistration(webauthnUser, protocol.MediationDefault)
	require.NoError(t, err)

	rp := tests.RelyingParty()
	authenticator := tests.Authenticator()
	credential := tests.Credential()

	creationJSON, err := json.Marshal(creation)
	require.NoError(t, err)

	parsedAttestationOptions, err := virtualwebauthn.ParseAttestationOptions(string(creationJSON))

	attestationResponse := virtualwebauthn.CreateAttestationResponse(rp, authenticator, credential, *parsedAttestationOptions)

	pcc, err := protocol.ParseCredentialCreationResponseBody(bytes.NewBufferString(attestationResponse))

	createCredentialResult, err := w.CreateCredential(webauthnUser, *sessionData, pcc)
	assert.NotNil(t, createCredentialResult)

	user, err := srv.RegisterUserWithCredential(t.Context(), w, webauthnUser, createCredentialResult)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userName, user.DisplayName.Plaintext())
}
