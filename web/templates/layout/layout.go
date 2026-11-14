package layout

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Layout(title string, children ...Node) Node {
	return HTML(
		Head(
			Title(title),
		),
		Group(children),
	)
}
