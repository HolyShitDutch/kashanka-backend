package webauthn

import (
	"github.com/moroz/kashanka/web/templates/components"
	"github.com/moroz/kashanka/web/templates/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Register() Node {
	return layout.Layout("Sign up", Div(
		components.PageTitle("Sign up"),
		Div(
			Button(
				Class("bg-blue-700 px-4 py-2 text-white font-bold"),
				ID("webauthn-registration-start"),
				Text("Sign up with passkey"),
			),
		),
	))
}
