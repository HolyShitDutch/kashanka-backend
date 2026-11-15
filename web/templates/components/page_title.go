package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func PageTitle(title string) Node {
	return H2(
		Class("text-3xl font-bold mb-6"),
		Text(title),
	)
}
