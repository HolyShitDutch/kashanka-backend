package helpers

import (
	"bytes"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"maragu.dev/gomponents"
)

func Render(c *echo.Context, status int, template gomponents.Node) error {
	var buf bytes.Buffer
	if err := template.Render(&buf); err != nil {
		slog.Error("Template render failed", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.HTMLBlob(status, buf.Bytes())
}
