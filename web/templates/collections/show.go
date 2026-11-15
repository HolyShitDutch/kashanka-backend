package collections

import (
	"github.com/moroz/kashanka/db/queries"
	"github.com/moroz/kashanka/internal/types"
	"github.com/moroz/kashanka/web/templates/components"
	"github.com/moroz/kashanka/web/templates/layout"

	. "maragu.dev/gomponents"

	. "maragu.dev/gomponents/html"
)

type ShowProps struct {
	Collection *types.CollectionDTO
}

func Show(props *ShowProps) Node {
	return layout.Layout(props.Collection.Title,
		A(Href("/"), Text("<< Back to list")),
		components.PageTitle("Collections"),
		Div(
			Class("grid grid-cols-3 gap-6"),
			Map(props.Collection.Cards, func(row *queries.Card) Node {
				return Article(
					Class("p-4 border"),
					H3(
						Class("font-bold"),
						Text(row.Question),
					),
					P(Text(row.Answers["en"])),
				)
			}),
		),
	)
}
