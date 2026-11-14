package handlers

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
)

type RouterDeps struct {
	DB     *pgxpool.Pool
	Logger *slog.Logger
}

func Router(props *RouterDeps) *echo.Echo {
	r := echo.New()
	r.Logger = props.Logger

	return r
}
