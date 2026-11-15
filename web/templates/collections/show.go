package collections

import (
	"github.com/moroz/kashanka/db/queries"
	"github.com/moroz/kashanka/internal/types"
	"github.com/moroz/kashanka/web/templates/layout"

	. "maragu.dev/gomponents"

	. "maragu.dev/gomponents/html"
)

type ShowProps struct {
	Collection *types.CollectionDTO
}

func Show(props *ShowProps) Node {
	return layout.Layout(props.Collection.Title, Div(
		H2(Text(props.Collection.Title)),
		Map(props.Collection.Cards, func(row *queries.Card) Node {
			return Article(
				H3(Text(row.Question)),
				P(Text(row.Answers["en"])),
			)
		}),
	))
}
