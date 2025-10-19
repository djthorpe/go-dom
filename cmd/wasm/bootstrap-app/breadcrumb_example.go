package main

import (
	"fmt"

	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddBreadcrumbExamples adds breadcrumb component examples to the app
func AddBreadcrumbExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Example 1: Simple breadcrumb (active only)
	card1 := app.Card()
	card1.AddClass("mb-4")
	card1.Header(app.H4(app.CreateTextNode("Simple Breadcrumb")).Element)
	card1Body := card1.Body()

	bc1 := app.Breadcrumb("breadcrumb1")
	bc1.AddItem("Home", "", true)
	card1Body.Element.AppendChild(bc1.Element)
	container.AppendChild(card1.Element)

	// Example 2: Two-level breadcrumb
	card2 := app.Card()
	card2.AddClass("mb-4")
	card2.Header(app.H4(app.CreateTextNode("Two-Level Breadcrumb")).Element)
	card2Body := card2.Body()

	bc2 := app.Breadcrumb("breadcrumb2")
	bc2.AddItem("Home", "#", false)
	bc2.AddItem("Library", "", true)
	card2Body.Element.AppendChild(bc2.Element)
	container.AppendChild(card2.Element)

	// Example 3: Three-level breadcrumb with click handler
	card3 := app.Card()
	card3.AddClass("mb-4")
	card3.Header(app.H4(app.CreateTextNode("Three-Level Breadcrumb with Click Handler")).Element)
	card3Body := card3.Body()

	bc3 := app.Breadcrumb("breadcrumb3")
	bc3.AddItem("Home", "#home", false).AddEventListener("click", func(target dom.Node) {
		fmt.Println("Breadcrumb: Home clicked")
	})
	bc3.AddItem("Library", "#library", false)
	bc3.AddItem("Data", "", true)
	card3Body.Element.AppendChild(bc3.Element)
	container.AppendChild(card3.Element)

	// Example 4: Custom divider
	card4 := app.Card()
	card4.AddClass("mb-4")
	card4.Header(app.H4(app.CreateTextNode("Custom Divider (»)")).Element)
	card4Body := card4.Body()

	bc4 := app.Breadcrumb("breadcrumb4")
	bc4.SetDivider("»")
	bc4.AddItem("Home", "#", false)
	bc4.AddItem("Settings", "#", false)
	bc4.AddItem("Profile", "", true)
	card4Body.Element.AppendChild(bc4.Element)
	container.AppendChild(card4.Element)

	return container.Element
}
