package handlers

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type RouterDeps struct {
	DB     *pgxpool.Pool
	Logger *slog.Logger
}

func Router(props *RouterDeps) *echo.Echo {
	r := echo.New()
	r.Logger = props.Logger

	r.Use(middleware.RequestLogger())
	r.Use(SetCustomRequestContext)

	collections := CollectionController(props.DB)
	r.GET("/", collections.Index)
	r.GET("/collections/:slug", collections.Show)

	return r
}
