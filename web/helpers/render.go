package helpers

import (
	"bytes"
	"encoding/json/v2"
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

func RenderJSON(c *echo.Context, status int, value any) error {
	binary, err := json.Marshal(value)
	if err != nil {
		slog.Error("Failed to encode response as JSON", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Response().WriteHeader(status)
	_, err = c.Response().Write(binary)
	return err
}
