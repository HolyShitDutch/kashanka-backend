package handlers

import (
	"log/slog"
	"net/http"

	"github.com/HolyShitDutch/kashanka/config"
	"github.com/HolyShitDutch/kashanka/internal/sessions"
	"github.com/HolyShitDutch/kashanka/internal/types"
	"github.com/HolyShitDutch/kashanka/services"
	"github.com/HolyShitDutch/kashanka/web/helpers"
	"github.com/labstack/echo/v5"
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
			slog.Debug("failed to decode session from session cookie", "error", err)
		}

		return next(c)
	}
}

func FetchUserFromSession(db services.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			ctx := helpers.GetRequestContext(c)

			userToken, ok := ctx.Session().Get[[]byte](config.SessionAccessTokenKey)
			if !ok {
				return next(c)
			}

			user, _ := services.NewUserTokenService(db).GetUserByAccessToken(c.Request().Context(), userToken)
			ctx.User = user

			return next(c)
		}
	}
}

func RequireAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := helpers.GetRequestContext(c)

		if ctx.User != nil {
			return next(c)
		}

		return c.Redirect(http.StatusSeeOther, "/sign-in")
	}
}

func RedirectIfAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := helpers.GetRequestContext(c)

		if ctx.User == nil {
			return next(c)
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}
}
