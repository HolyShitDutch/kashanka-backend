package handlers

import (
	"github.com/labstack/echo/v5"
	"github.com/moroz/kashanka/web/helpers"
	"github.com/moroz/kashanka/web/templates/auth"
)

type userRegistrationController struct{}

func UserRegistrationController() *userRegistrationController {
	return &userRegistrationController{}
}

func (cc *userRegistrationController) New(c *echo.Context) error {
	return helpers.Render(c, 200, auth.SignUp())
}
