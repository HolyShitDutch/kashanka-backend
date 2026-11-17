package helpers

import (
	"github.com/labstack/echo/v5"
	"github.com/moroz/kashanka/internal/types"
)

func GetRequestContext(c *echo.Context) *types.RequestContext {
	// panic if context is not set
	return c.Get("context").(*types.RequestContext)
}
