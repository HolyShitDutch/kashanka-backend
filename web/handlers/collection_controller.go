package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v5"
	"github.com/moroz/kashanka/db/queries"
	"github.com/moroz/kashanka/web/helpers"
	"github.com/moroz/kashanka/web/templates/collections"
)

type collectionController struct {
	db queries.DBTX
}

func CollectionController(db queries.DBTX) *collectionController {
	return &collectionController{
		db,
	}
}

func (cc *collectionController) Index(c *echo.Context) error {
	data, err := queries.New(cc.db).ListCollections(c.Request().Context())
	if err != nil {
		slog.Error("Error fetching collections", "error", err)
		return echo.NewHTTPError(500, "Internal Server Error")
	}

	return helpers.Render(c, 200, collections.Index(&collections.IndexProps{
		Collections: data,
	}))
}
