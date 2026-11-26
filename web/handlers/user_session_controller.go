package handlers

import (
	"github.com/HolyShitDutch/kashanka/web/helpers"
	"github.com/HolyShitDutch/kashanka/web/templates/auth"
	"github.com/labstack/echo/v5"
)

type userSessionController struct{}

func UserSessionController() *userSessionController {
	return &userSessionController{}
}

func (cc *userSessionController) New(c *echo.Context) error {
	return helpers.Render(c, 200, auth.SignIn())
}
