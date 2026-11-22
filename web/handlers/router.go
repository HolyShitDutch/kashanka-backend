package handlers

import (
	"errors"
	"log/slog"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/internal/sessions"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type RouterDeps struct {
	DB           *pgxpool.Pool
	SessionStore sessions.Store
	WebAuthn     *webauthn.WebAuthn
}

func (r *RouterDeps) Validate() bool {
	return r.DB != nil && r.SessionStore != nil
}

func Router(props *RouterDeps) (*echo.Echo, error) {
	if !props.Validate() {
		return nil, errors.New("invalid props")
	}

	if props.WebAuthn == nil {
		w, err := config.WebauthnConfig()
		if err != nil {
			return nil, err
		}
		props.WebAuthn = w
	}

	r := echo.New()
	r.Logger = slog.Default()

	r.Use(middleware.RequestLogger())
	r.Use(SetCustomRequestContext(props.SessionStore))
	r.Use(ReadSessionCookie)

	collections := CollectionController(props.DB)
	r.GET("/", collections.Index)
	r.GET("/collections/:slug", collections.Show)

	registrations := UserRegistrationController()
	r.GET("/sign-up", registrations.New)
	r.GET("/sign-up/passkey", registrations.Passkey)

	webauthn := WebauthnController(&WebauthnControllerDeps{
		DB:       props.DB,
		WebAuthn: props.WebAuthn,
	})
	r.GET("/.well-known/webauthn", webauthn.WellKnown)
	r.POST("/webauthn/register/start", webauthn.RegisterStart)
	r.POST("/webauthn/register/finalize", webauthn.RegisterFinalize)

	if !config.IsProd {
		r.IPExtractor = echo.ExtractIPDirect()
		r.Static("/assets", "assets/public/assets")
	}

	return r, nil
}
