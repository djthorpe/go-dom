package main

import (
	"fmt"

	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddNavExamples adds navigation bar examples to the app
func AddNavExamples(app *bs5.App) dom.Element {
	container := app.CreateElement("div")

	// Example 1: Dark navbar with icon
	darkCard := app.Card()
	darkCard.AddClass("mb-4")
	darkCard.Header(app.H4(app.CreateTextNode("Dark Navbar with Icon")).Element)
	darkCardBody := darkCard.Body()

	descP1 := app.CreateElement("p")
	descP1.AppendChild(app.CreateTextNode("A dark themed navbar with icons and dropdown menu."))
	darkCardBody.Element.AppendChild(descP1)

	// Create brand with icon
	darkBrandContent := app.CreateElement("span")
	darkBrandContent.AppendChild(app.Icon("bootstrap-fill").Element)
	darkBrandContent.AppendChild(app.CreateTextNode(" Brand"))

	darkNav := app.Nav(
		app.NavBrand("#", darkBrandContent).Element,
		app.NavItem("Home", "#home", true).AddEventListener("click", func(target dom.Node) {
			fmt.Println("Home clicked:", target)
		}).Element,
		app.NavItem("Features", "#features", false).Element,
		app.NavDropdown("More",
			app.NavDropdownItem("Action", "#action").Element,
			app.NavDropdownItem("Another action", "#another").Element,
			app.NavDropdownDivider().Element,
			app.NavDropdownItem("Something else", "#else").Element,
		).Element,
		app.NavItem("Pricing", "#pricing", false).Element,
		app.NavSpacer().Element,
		app.NavItem("About", "#about", false).Element,
	).SetColorScheme(bs5.ColorDark).AddClass("bg-dark")
	darkCardBody.Element.AppendChild(darkNav.Element)
	container.AppendChild(darkCard.Element)

	// Example 2: Light navbar with icon
	lightCard := app.Card()
	lightCard.AddClass("mb-4")
	lightCard.Header(app.H4(app.CreateTextNode("Light Navbar with Icon")).Element)
	lightCardBody := lightCard.Body()

	descP2 := app.CreateElement("p")
	descP2.AppendChild(app.CreateTextNode("A light themed navbar perfect for bright designs."))
	lightCardBody.Element.AppendChild(descP2)

	// Create brand with icon
	lightBrandContent := app.CreateElement("span")
	lightBrandContent.AppendChild(app.Icon("building").Element)
	lightBrandContent.AppendChild(app.CreateTextNode(" Company"))

	lightNav := app.Nav(
		app.NavBrand("#", lightBrandContent).Element,
		app.NavItem("Products", "#products", true).Element,
		app.NavItem("Services", "#services", false).Element,
		app.NavItem("Contact", "#contact", false).Element,
	).SetColorScheme(bs5.ColorLight).AddClass("bg-light")
	lightCardBody.Element.AppendChild(lightNav.Element)
	container.AppendChild(lightCard.Element)

	// Example 3: Primary colored navbar with icon
	primaryCard := app.Card()
	primaryCard.AddClass("mb-4")
	primaryCard.Header(app.H4(app.CreateTextNode("Primary Colored Navbar with Icon")).Element)
	primaryCardBody := primaryCard.Body()

	descP3 := app.CreateElement("p")
	descP3.AppendChild(app.CreateTextNode("A navbar with Bootstrap's primary blue color."))
	primaryCardBody.Element.AppendChild(descP3)

	// Create brand with icon
	primaryBrandContent := app.CreateElement("span")
	primaryBrandContent.AppendChild(app.Icon("app-indicator").Element)
	primaryBrandContent.AppendChild(app.CreateTextNode(" MyApp"))

	primaryNav := app.Nav(
		app.NavBrand("#", primaryBrandContent).Element,
		app.NavItem("Dashboard", "#dashboard", true).Element,
		app.NavItem("Reports", "#reports", false).Element,
		app.NavItem("Settings", "#settings", false).Element,
		app.NavSpacer().Element,
		app.NavItem("Logout", "#logout", false).Element,
	)
	primaryNav.Element.RemoveClass("bg-body-tertiary")
	primaryNav.SetColorScheme(bs5.ColorDark).AddClass("bg-primary")
	primaryCardBody.Element.AppendChild(primaryNav.Element)
	container.AppendChild(primaryCard.Element)

	return container
}
