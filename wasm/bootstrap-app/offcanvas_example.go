package main

import (
	"fmt"

	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddOffcanvasExamples adds offcanvas component examples to the app
func AddOffcanvasExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Placement examples
	placementCard := app.Card()
	placementCard.AddClass("mb-4")
	placementCard.Header(app.H4(app.CreateTextNode("Placement")).Element)
	placementCardBody := placementCard.Body()

	descP1 := app.CreateElement("p")
	descP1.AppendChild(app.CreateTextNode("Offcanvas can be placed on different sides of the viewport:"))
	placementCardBody.Element.AppendChild(descP1)

	// Start (left) offcanvas
	startOffcanvas := app.Offcanvas("offcanvasStart", bs5.OffcanvasPlacementStart)
	startOffcanvas.AddTitle("Offcanvas Start", app).AddCloseButton(app)
	startOffcanvas.Body().AppendChild(app.CreateTextNode("This offcanvas slides in from the left side of the screen."))
	app.Document.Body().AppendChild(startOffcanvas.Element)

	startBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Left"))
	startBtn.SetAttribute("data-bs-toggle", "offcanvas")
	startBtn.SetAttribute("data-bs-target", "#offcanvasStart")
	startBtn.AddClass("me-2")
	placementCardBody.Element.AppendChild(startBtn.Element)

	// End (right) offcanvas
	endOffcanvas := app.Offcanvas("offcanvasEnd", bs5.OffcanvasPlacementEnd)
	endOffcanvas.AddTitle("Offcanvas End", app).AddCloseButton(app)
	endOffcanvas.Body().AppendChild(app.CreateTextNode("This offcanvas slides in from the right side of the screen."))
	app.Document.Body().AppendChild(endOffcanvas.Element)

	endBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Right"))
	endBtn.SetAttribute("data-bs-toggle", "offcanvas")
	endBtn.SetAttribute("data-bs-target", "#offcanvasEnd")
	endBtn.AddClass("me-2")
	placementCardBody.Element.AppendChild(endBtn.Element)

	// Top offcanvas
	topOffcanvas := app.Offcanvas("offcanvasTop", bs5.OffcanvasPlacementTop)
	topOffcanvas.AddTitle("Offcanvas Top", app).AddCloseButton(app)
	topOffcanvas.Body().AppendChild(app.CreateTextNode("This offcanvas slides in from the top of the screen."))
	app.Document.Body().AppendChild(topOffcanvas.Element)

	topBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Top"))
	topBtn.SetAttribute("data-bs-toggle", "offcanvas")
	topBtn.SetAttribute("data-bs-target", "#offcanvasTop")
	topBtn.AddClass("me-2")
	placementCardBody.Element.AppendChild(topBtn.Element)

	// Bottom offcanvas
	bottomOffcanvas := app.Offcanvas("offcanvasBottom", bs5.OffcanvasPlacementBottom)
	bottomOffcanvas.AddTitle("Offcanvas Bottom", app).AddCloseButton(app)
	bottomOffcanvas.Body().AppendChild(app.CreateTextNode("This offcanvas slides in from the bottom of the screen."))
	app.Document.Body().AppendChild(bottomOffcanvas.Element)

	bottomBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Bottom"))
	bottomBtn.SetAttribute("data-bs-toggle", "offcanvas")
	bottomBtn.SetAttribute("data-bs-target", "#offcanvasBottom")
	placementCardBody.Element.AppendChild(bottomBtn.Element)

	container.AppendChild(placementCard.Element)

	// Backdrop options
	backdropCard := app.Card()
	backdropCard.AddClass("mb-4")
	backdropCard.Header(app.H4(app.CreateTextNode("Backdrop Options")).Element)
	backdropCardBody := backdropCard.Body()

	// With body scrolling
	scrollOffcanvas := app.Offcanvas("offcanvasScroll", bs5.OffcanvasPlacementStart)
	scrollOffcanvas.AddTitle("Body Scrolling Enabled", app).AddCloseButton(app)
	scrollOffcanvas.SetBodyScroll(true).SetBackdrop("false")
	scrollOffcanvas.Body().AppendChild(app.CreateTextNode("Body scrolling is enabled and there is no backdrop. You can scroll the page while this offcanvas is open."))
	app.Document.Body().AppendChild(scrollOffcanvas.Element)

	scrollBtn := app.Button(bs5.ColorSuccess, app.CreateTextNode("Enable Body Scrolling"))
	scrollBtn.SetAttribute("data-bs-toggle", "offcanvas")
	scrollBtn.SetAttribute("data-bs-target", "#offcanvasScroll")
	scrollBtn.AddClass("me-2")
	backdropCardBody.Element.AppendChild(scrollBtn.Element)

	// With backdrop and scrolling
	bothOffcanvas := app.Offcanvas("offcanvasBoth", bs5.OffcanvasPlacementStart)
	bothOffcanvas.AddTitle("Backdrop with Scrolling", app).AddCloseButton(app)
	bothOffcanvas.SetBodyScroll(true).SetBackdrop("true")
	bothOffcanvas.Body().AppendChild(app.CreateTextNode("Body scrolling is enabled with a visible backdrop."))
	app.Document.Body().AppendChild(bothOffcanvas.Element)

	bothBtn := app.Button(bs5.ColorSuccess, app.CreateTextNode("Scrolling + Backdrop"))
	bothBtn.SetAttribute("data-bs-toggle", "offcanvas")
	bothBtn.SetAttribute("data-bs-target", "#offcanvasBoth")
	bothBtn.AddClass("me-2")
	backdropCardBody.Element.AppendChild(bothBtn.Element)

	// Static backdrop
	staticOffcanvas := app.Offcanvas("offcanvasStatic", bs5.OffcanvasPlacementStart)
	staticOffcanvas.AddTitle("Static Backdrop", app).AddCloseButton(app)
	staticOffcanvas.SetBackdrop("static")
	staticOffcanvas.Body().AppendChild(app.CreateTextNode("I will not close if you click outside of me. Use the close button or ESC key."))
	app.Document.Body().AppendChild(staticOffcanvas.Element)

	staticBtn := app.Button(bs5.ColorWarning, app.CreateTextNode("Static Backdrop"))
	staticBtn.SetAttribute("data-bs-toggle", "offcanvas")
	staticBtn.SetAttribute("data-bs-target", "#offcanvasStatic")
	backdropCardBody.Element.AppendChild(staticBtn.Element)

	container.AppendChild(backdropCard.Element)

	// Dark theme
	darkCard := app.Card()
	darkCard.AddClass("mb-4")
	darkCard.Header(app.H4(app.CreateTextNode("Dark Theme")).Element)
	darkCardBody := darkCard.Body()

	darkOffcanvas := app.Offcanvas("offcanvasDark", bs5.OffcanvasPlacementStart)
	darkOffcanvas.AddTitle("Dark Offcanvas", app).AddCloseButton(app)
	darkOffcanvas.SetDark(true)
	darkOffcanvas.Body().AppendChild(app.CreateTextNode("This offcanvas uses a dark theme."))
	app.Document.Body().AppendChild(darkOffcanvas.Element)

	darkBtn := app.Button(bs5.ColorDark, app.CreateTextNode("Dark Theme"))
	darkBtn.SetAttribute("data-bs-toggle", "offcanvas")
	darkBtn.SetAttribute("data-bs-target", "#offcanvasDark")
	darkCardBody.Element.AppendChild(darkBtn.Element)

	container.AppendChild(darkCard.Element)

	// With content
	contentCard := app.Card()
	contentCard.AddClass("mb-4")
	contentCard.Header(app.H4(app.CreateTextNode("With Content")).Element)
	contentCardBody := contentCard.Body()

	contentOffcanvas := app.Offcanvas("offcanvasContent", bs5.OffcanvasPlacementEnd)
	contentOffcanvas.AddTitle("Shopping Cart", app).AddCloseButton(app)

	// Add some content to the offcanvas body
	bodyContent := app.CreateElement("div")

	// Add a paragraph
	p := app.CreateElement("p")
	p.AppendChild(app.CreateTextNode("Your shopping cart items:"))
	bodyContent.AppendChild(p)

	// Add a list
	ul := app.CreateElement("ul")
	ul.AddClass("list-group")
	for i := 1; i <= 3; i++ {
		li := app.CreateElement("li")
		li.AddClass("list-group-item")
		li.AppendChild(app.CreateTextNode(fmt.Sprintf("Item %d - $%d.99", i, i*10)))
		ul.AppendChild(li)
	}
	bodyContent.AppendChild(ul)

	// Add a button
	checkoutBtn := app.Button(bs5.ColorSuccess, app.CreateTextNode("Checkout"))
	checkoutBtn.AddClass("mt-3").AddClass("w-100")
	checkoutBtn.AddEventListener("click", func(target dom.Node) {
		fmt.Println("Checkout clicked!")
	})
	bodyContent.AppendChild(checkoutBtn.Element)

	contentOffcanvas.Body().AppendChild(bodyContent)
	app.Document.Body().AppendChild(contentOffcanvas.Element)

	contentBtn := app.Button(bs5.ColorInfo, app.CreateTextNode("Shopping Cart"))
	contentBtn.SetAttribute("data-bs-toggle", "offcanvas")
	contentBtn.SetAttribute("data-bs-target", "#offcanvasContent")
	contentCardBody.Element.AppendChild(contentBtn.Element)

	container.AppendChild(contentCard.Element)

	return container.Element
}
