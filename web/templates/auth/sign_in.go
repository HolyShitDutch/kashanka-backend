package auth

import (
	"github.com/HolyShitDutch/kashanka/web/templates/components"
	"github.com/HolyShitDutch/kashanka/web/templates/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func SignIn() Node {
	return layout.AuthLayout("Sign in",
		components.PageTitle("Sign in"),
		Div(
			Class("grid gap-4 mb-6"),
		),
		Hr(Class("my-4")),
		Footer(
			P(
				Text("Don't have an account yet?"),
				Br(),
				A(Href("/sign-up"), Text("Sign up")),
			),
		),
	)
}
