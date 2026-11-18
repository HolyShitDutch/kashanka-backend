package handlers

import (
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/moroz/kashanka/config"
	"github.com/moroz/kashanka/internal/sessions"
)

type RouterDeps struct {
	DB           *pgxpool.Pool
	Logger       *slog.Logger
	SessionStore sessions.Store
}

func (r *RouterDeps) Validate() bool {
	return r.DB != nil && r.Logger != nil && r.SessionStore != nil
}

func Router(props *RouterDeps) (*echo.Echo, error) {
	if !props.Validate() {
		return nil, errors.New("invalid props")
	}

	r := echo.New()
	r.Logger = props.Logger

	r.Use(middleware.RequestLogger())
	r.Use(SetCustomRequestContext(props.SessionStore))
	r.Use(ReadSessionCookie)

	collections := CollectionController(props.DB)
	r.GET("/", collections.Index)
	r.GET("/collections/:slug", collections.Show)

	webauthn := WebauthnController(&WebauthnControllerDeps{
		DB:        props.DB,
		RPID:      config.RPID,
		RPOrigins: config.RPOrigins,
	})
	r.GET("/sign-up", webauthn.Register)
	r.GET("/webauthn/register/start", webauthn.RegisterStart)

	return r, nil
}
