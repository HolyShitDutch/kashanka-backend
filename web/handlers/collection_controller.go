package handlers

import (
	"github.com/labstack/echo/v5"
	"github.com/moroz/kashanka/db/queries"
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
	collections, err := queries.New(cc.db).ListCollections(c.Request().Context())

	return nil
}
