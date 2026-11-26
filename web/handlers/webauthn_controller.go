package handlers

import (
	"log/slog"
	"net/http"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/internal/requests"
	"github.com/HolyShitDutch/kashanka/internal/types"
	"github.com/HolyShitDutch/kashanka/services"
	"github.com/HolyShitDutch/kashanka/web/helpers"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/labstack/echo/v5"
)

type webauthnController struct {
	db       services.DB
	webauthn *webauthn.WebAuthn
}

type WebauthnControllerDeps struct {
	DB       services.DB
	WebAuthn *webauthn.WebAuthn
}

func WebauthnController(deps *WebauthnControllerDeps) *webauthnController {
	return &webauthnController{
		db:       deps.DB,
		webauthn: deps.WebAuthn,
	}
}

// GET /.well-known/webauthn
func (cc *webauthnController) WellKnown(c *echo.Context) error {
	return helpers.RenderJSON(c, 200, map[string][]string{
		"origins": cc.webauthn.Config.RPOrigins,
	})
}

// POST /webauthn/register/start
func (cc *webauthnController) RegisterStart(c *echo.Context) error {
	ctx := helpers.GetRequestContext(c)
	var params requests.WebauthnRegisterStartRequest
	if err := c.Bind(&params); err != nil {
		slog.Error("Failed to parse request", "error", err)
		return echo.ErrBadRequest
	}

	user := types.NewWebAuthnUser(params.DisplayName)

	var (
		creation    *protocol.CredentialCreation
		sessionData *webauthn.SessionData
		err         error
	)

	opts := []webauthn.RegistrationOption{
		webauthn.WithResidentKeyRequirement(protocol.ResidentKeyRequirementPreferred),
	}

	if creation, sessionData, err = cc.webauthn.BeginMediatedRegistration(user, protocol.MediationDefault, opts...); err != nil {
		slog.Error("Failed to begin registration", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	ctx.Session().Set("registration_challenge", sessionData)
	ctx.Session().Set("new_user", user)
	if err := ctx.SaveSession(c.Response()); err != nil {
		slog.Error("Failed to save session data", "error", err)
		return echo.ErrInternalServerError
	}

	return helpers.RenderJSON(c, 200, creation)
}

func (cc *webauthnController) RegisterFinalize(c *echo.Context) error {
	ctx := helpers.GetRequestContext(c)

	sessionData, ok := ctx.Session().Get[*webauthn.SessionData]("registration_challenge")
	if !ok {
		slog.Error("WebAuthn attestation failed: Failed to decode WebAuthn session data", "details", "challenge")
		return echo.ErrBadRequest
	}

	userData, ok := ctx.Session().Get[*types.WebAuthnUser]("new_user")
	if !ok {
		slog.Error("WebAuthn attestation failed: Failed to decode WebAuthn session data")
		return echo.ErrBadRequest
	}

	pcc, err := protocol.ParseCredentialCreationResponseBody(c.Request().Body)
	if err != nil {
		slog.Error("WebAuthn attestation failed: Failed to parse credential creation response", "error", err)
		return echo.ErrBadRequest
	}

	credential, err := cc.webauthn.CreateCredential(userData, *sessionData, pcc)
	if err != nil {
		slog.Error("WebAuthn attestation failed", "error", err)
		return echo.ErrBadRequest
	}

	tx, err := cc.db.Begin(c.Request().Context())
	if err != nil {
		slog.Error("Failed to begin transaction in RegisterFinalize", "error", err)
		return echo.ErrInternalServerError
	}
	defer tx.Rollback(c.Request().Context())

	user, err := services.NewUserService(tx).RegisterUserWithCredential(c.Request().Context(), cc.webauthn, userData, credential)
	if err != nil {
		return err
	}

	token, err := services.NewUserTokenService(tx).CreateAccessTokenForUser(c.Request().Context(), user)
	if err != nil {
		return err
	}

	if err := tx.Commit(c.Request().Context()); err != nil {
		slog.Error("Failed to store session", "error", err)
		return echo.ErrInternalServerError
	}

	ctx.Session().Delete("registration_challenge", "new_user")
	ctx.Session().Set(config.SessionAccessTokenKey, token.PlaintextToken)
	ctx.SaveSession(c.Response())

	return helpers.RenderJSON(c, 200, map[string]any{
		"data": user,
	})
}
