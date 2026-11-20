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

func Passkey() Node {
	return layout.AuthLayout("Sign up with passkey",
		layout.ScriptEntryPoint("src/webauthn.ts"),
		components.PageTitle("Sign up with passkey"),
		Form(
			ID("webauthn-registration"),
			Class("grid gap-4"),
			Label(
				Class("input-field"),
				Span(
					Text("What should we call you? "),
					Span(
						Class("text-muted-fg text-sm"),
						Text("(optional)"),
					),
				),
				Input(
					Class("input"),
					AutoFocus(),
					Type("text"), Name("display_name"),
				),
				Span(
					Class("text-xs leading-tight font-normal"),
					Text("This name will be stored in your passkey manager. If you leave this field empty, the app will generate a random name, e.g. Rabid Raccoon."),
				),
			),
			Button(
				Class("button w-full"),
				Text("Create a passkey"),
			),
		),
	)
}
