package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v5"
	"github.com/moroz/kashanka/config"
	"github.com/moroz/kashanka/internal/sessions"
	"github.com/moroz/kashanka/internal/types"
	"github.com/moroz/kashanka/web/helpers"
)

func SetCustomRequestContext(store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("context", types.NewRequestContext(store))
			return next(c)
		}
	}
}

func ReadSessionCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := helpers.GetRequestContext(c)

		// ignore malformed cookie and silently continue
		cookie, err := c.Cookie(config.SessionCookieName)
		if err != nil {
			return next(c)
		}

		if err := ctx.DecodeSessionFromCookie(cookie.Value, config.SessionMaxAge); err != nil {
			slog.Info("failed to decode session from session cookie", "error", err)
		}

		return next(c)
	}
}
