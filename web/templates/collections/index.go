package collections

import (
	"fmt"

	"github.com/moroz/kashanka/db/queries"
	"github.com/moroz/kashanka/web/templates/components"
	"github.com/moroz/kashanka/web/templates/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type IndexProps struct {
	Collections []*queries.Collection
}

func Index(props *IndexProps) Node {
	return layout.Layout("Collections",
		components.PageTitle("Collections"),
		Map(props.Collections, func(c *queries.Collection) Node {
			return A(
				Href(fmt.Sprintf("/collections/%s", c.Slug)),
				Div(
					Span(Text(c.Title)),
				),
			)
		}),
	)
}
