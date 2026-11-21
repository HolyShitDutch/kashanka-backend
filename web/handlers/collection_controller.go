package handlers

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/labstack/echo/v5"
	"github.com/HolyShitDutch/kashanka/db/queries"
	"github.com/HolyShitDutch/kashanka/services"
	"github.com/HolyShitDutch/kashanka/web/helpers"
	"github.com/HolyShitDutch/kashanka/web/templates/collections"
)

type collectionController struct {
	db          services.DB
	collections *services.CollectionRepository
}

func CollectionController(db services.DB) *collectionController {
	return &collectionController{
		db:          db,
		collections: services.NewCollectionRepository(db),
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

func (cc *collectionController) Show(c *echo.Context) error {
	data, err := cc.collections.GetCollectionWithCardsBySlug(c.Request().Context(), c.Param("slug"))
	if errors.Is(err, sql.ErrNoRows) {
		return echo.ErrNotFound
	}
	if err != nil {
		slog.Error("Error fetching collections", "error", err)
		return echo.NewHTTPError(500, "Internal Server Error")
	}

	return helpers.Render(c, 200, collections.Show(&collections.ShowProps{
		Collection: data,
	}))
}
