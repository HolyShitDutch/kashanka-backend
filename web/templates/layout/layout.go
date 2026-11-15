package layout

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Layout(title string, children ...Node) Node {
	return RootLayout(title, Div(
		Class("container mx-auto my-6"),
		Group(children),
	))
}

func RootLayout(title string, children ...Node) Node {
	return HTML(
		Head(
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
			Meta(Name("robots"), Content("noindex, nofollow")),
			Meta(Name("color-scheme"), Content("dark light")),
			TitleEl(Text(title)),
			AssetEntryPoint(),
		),
		Body(Group(children)),
	)
}
