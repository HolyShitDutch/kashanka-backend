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

func AuthLayout(title string, children ...Node) Node {
	return RootLayout(title, Div(
		Class("h-screen grid place-items-center"),
		Div(
			Class("grid gap-4 justify-center -mt-12"),
			A(
				Class("text-center flex justify-center items-center"),
				Href("/"),
				Picture(
					Source(SrcSet("/assets/logo-dark.svg"), Attr("media", "(prefers-color-scheme: dark)")),
					Img(
						Class("h-24"),
						SrcSet("/assets/logo-light.svg"),
						Src(""),
					),
				),
			),
			Div(
				Class("card w-100 text-center"),
				Group(children),
			),
		),
	))
}
