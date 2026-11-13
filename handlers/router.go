package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Router() http.Handler {
	r := echo.New()

	return r
}
