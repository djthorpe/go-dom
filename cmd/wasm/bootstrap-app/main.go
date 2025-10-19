package main

import (
	// Packages
	"github.com/djthorpe/go-dom/pkg/bs5"
)

func main() {
	// Create a new bootstrap application
	app := bs5.New("Bootstrap App Demo")

	// Add navigation
	app.Document.Body().AppendChild(
		app.Nav(
			app.NavBrand("#", app.CreateTextNode("BS5")).Element,
			app.NavItem("Home", "#home", true).Element,
			app.NavItem("Features", "#features", false).Element,
			app.NavItem("Pricing", "#pricing", false).Element,
			app.NavSpacer().Element,
			app.NavItem("About", "#about", false).Element,
		).SetColorScheme(bs5.ColorDark).AddClass("bg-dark").Element,
	)

	// Add alerts
	app.Document.Body().AppendChild(
		app.Alert(
			bs5.ColorSuccess,
			app.CreateTextNode("Success! Your operation completed successfully."),
		).MakeDismissible().Element,
	)
	app.Document.Body().AppendChild(
		app.Alert(
			bs5.ColorWarning,
			app.CreateTextNode("Warning! Please review this information carefully."),
		).Element,
	)
	app.Document.Body().AppendChild(
		app.Alert(
			bs5.ColorDanger,
			app.CreateTextNode("Error! Something went wrong."),
		).MakeDismissible().Element,
	)

	// Add body content
	app.Document.Body().AppendChild(app.Container(
		app.H1(
			app.CreateTextNode("Main Title "),
			app.Badge(bs5.ColorSecondary, app.CreateTextNode("New")).Element,
			app.CreateTextNode(" "),
			app.Badge(bs5.ColorPrimary, app.CreateTextNode("Featured")).Element,
		).Element,
		app.CreateTextNode("Hello, World! "),
		app.Badge(
			bs5.ColorSuccess,
			app.CreateTextNode("Success"),
		).Element,
		app.CreateTextNode(" "),
		app.Badge(
			bs5.ColorDanger,
			app.CreateTextNode("Hot"),
		).Element,
		app.CreateTextNode(" "),
		app.Badge(
			bs5.ColorWarning,
			app.CreateTextNode("Warning"),
		).Element,
		app.CreateTextNode(" "),
		app.Badge(
			bs5.ColorInfo,
			app.CreateTextNode("Info"),
		).Element,
	).Element)
}
