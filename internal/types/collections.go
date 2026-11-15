package types

import (
	"github.com/moroz/kashanka/db/queries"
)

type CollectionDTO struct {
	*queries.Collection
	Cards []*queries.Card
}
