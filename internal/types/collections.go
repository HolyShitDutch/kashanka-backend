package types

import (
	"github.com/HolyShitDutch/kashanka/db/queries"
)

type CollectionDTO struct {
	*queries.Collection
	Cards []*queries.Card
}
