package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/labstack/echo/v5"
	"github.com/moroz/kashanka/config"
	"github.com/moroz/kashanka/internal/requests"
	"github.com/moroz/kashanka/internal/types"
	"github.com/moroz/kashanka/services"
	"github.com/moroz/kashanka/web/helpers"
)

type webauthnController struct {
	db       services.DB
	webauthn *webauthn.WebAuthn
}

type WebauthnControllerDeps struct {
	DB        services.DB
	RPID      string
	RPOrigins []string
}

func WebauthnController(deps *WebauthnControllerDeps) *webauthnController {
	w, _ := webauthn.New(&webauthn.Config{
		RPID:          deps.RPID,
		RPDisplayName: config.RPDisplayName,
		RPOrigins:     config.RPOrigins,
	})

	return &webauthnController{
		db:       deps.DB,
		webauthn: w,
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
	ctx.SaveSession(c.Response())
	return helpers.RenderJSON(c, 200, creation)
}
