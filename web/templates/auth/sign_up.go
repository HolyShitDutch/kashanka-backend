package auth

import (
	"github.com/moroz/kashanka/web/templates/components"
	"github.com/moroz/kashanka/web/templates/icons"
	"github.com/moroz/kashanka/web/templates/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func SignUp() Node {
	return layout.AuthLayout("Sign up",
		components.PageTitle("Sign up"),
		Div(
			Class("grid gap-4 mb-6"),
			A(
				Href("/sign-up/passkey"),
				Class("button w-full"),
				icons.Icon(&icons.IconProps{
					Name:    "lucide-key",
					ViewBox: "0 0 24 24",
				}),
				Text("Sign up with passkey"),
			),
			A(
				Href("/oauth/google/start"),
				Class("button w-full"),
				icons.Icon(&icons.IconProps{
					Name:    "google",
					ViewBox: "0 0 24 24",
				}),
				Text("Sign up with Google"),
			),
		),
		Hr(Class("my-4")),
		Footer(
			P(
				Text("Already have an account?"),
				Br(),
				A(Href("/sign-in"), Text("Sign in")),
			),
		),
	)
}
