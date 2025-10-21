//go:build js && wasm

package main

import (
	"fmt"

	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

func AddToastExamples(app *bs5.App, container dom.Element) {
	// Create a toast container for all toasts (fixed position bottom-right)
	toastContainer := app.ToastContainer()
	toastContainer.SetPosition("bottom-0 end-0").SetPadding("p-3")
	app.Document.Body().AppendChild(toastContainer.Element)

	// Section: Basic Toast
	basicCard := app.Card()
	basicCard.AddClass("mb-4")
	basicCard.Header(app.H4(app.CreateTextNode("Basic Toast")).Element)
	basicCardBody := basicCard.Body()

	descP1 := app.CreateElement("p")
	descP1.AppendChild(app.CreateTextNode("Push notifications to your visitors with a toast, a lightweight and easily customizable alert message."))
	basicCardBody.Element.AppendChild(descP1)

	// Create the toast
	basicToast := app.Toast()
	basicToast.SetAttribute("id", "liveToast1")
	basicToast.AddHeader("Bootstrap", "11 mins ago", app)
	basicToast.AddCloseButton(app)
	basicToast.SetBody(app.CreateTextNode("Hello, world! This is a toast message."), app)
	toastContainer.AddToast(basicToast)

	basicBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Show Basic Toast"))
	basicBtn.AddEventListener("click", func(dom.Node) {
		showToast("liveToast1")
	})
	basicCardBody.Element.AppendChild(basicBtn.Element)
	container.AppendChild(basicCard.Element)

	// Section: Color Schemes
	colorCard := app.Card()
	colorCard.AddClass("mb-4")
	colorCard.Header(app.H4(app.CreateTextNode("Colored Toasts")).Element)
	colorCardBody := colorCard.Body()

	descP2 := app.CreateElement("p")
	descP2.AppendChild(app.CreateTextNode("Toasts with contextual color schemes and icons."))
	colorCardBody.Element.AppendChild(descP2)

	colorButtonsDiv := app.CreateElement("div")
	colorCardBody.Element.AppendChild(colorButtonsDiv)

	colors := []struct {
		variant bs5.ColorVariant
		id      string
		icon    string
	}{
		{bs5.ColorPrimary, "toastPrimary", "info-circle-fill"},
		{bs5.ColorSuccess, "toastSuccess", "check-circle-fill"},
		{bs5.ColorDanger, "toastDanger", "x-circle-fill"},
		{bs5.ColorWarning, "toastWarning", "exclamation-triangle-fill"},
		{bs5.ColorInfo, "toastInfo", "info-circle-fill"},
	}

	for _, c := range colors {
		// Create colored toast
		colorToast := app.Toast()
		colorToast.SetAttribute("id", c.id)
		colorToast.SetColor(c.variant)
		colorToast.AddClass("align-items-center")

		wrapper := app.CreateElement("div")
		wrapper.AddClass("d-flex")

		body := app.CreateElement("div")
		body.AddClass("toast-body")
		body.AddClass("d-flex")
		body.AddClass("align-items-center")

		// Add icon
		icon := app.Icon(c.icon)
		icon.AddClass("me-2")
		body.AppendChild(icon.Element)

		body.AppendChild(app.CreateTextNode(fmt.Sprintf("This is a %s toast message.", c.variant)))
		wrapper.AppendChild(body)

		closeBtn := app.CreateElement("button")
		closeBtn.SetAttribute("type", "button")
		closeBtn.AddClass("btn-close")
		closeBtn.AddClass("btn-close-white")
		closeBtn.AddClass("me-2")
		closeBtn.AddClass("m-auto")
		closeBtn.SetAttribute("data-bs-dismiss", "toast")
		closeBtn.SetAttribute("aria-label", "Close")
		wrapper.AppendChild(closeBtn)

		colorToast.Element.AppendChild(wrapper)
		toastContainer.AddToast(colorToast)

		// Create button to show toast
		btn := app.Button(c.variant, app.CreateTextNode(string(c.variant)+" Toast"))
		btn.AddClass("me-2")
		btn.AddClass("mb-2")
		toastId := c.id // Capture for closure
		btn.AddEventListener("click", func(dom.Node) {
			showToast(toastId)
		})
		colorButtonsDiv.AppendChild(btn.Element)
	}
	container.AppendChild(colorCard.Element)

	// Section: Persistent Toast (no auto-hide)
	persistCard := app.Card()
	persistCard.AddClass("mb-4")
	persistCard.Header(app.H4(app.CreateTextNode("Persistent Toast")).Element)
	persistCardBody := persistCard.Body()

	descP3 := app.CreateElement("p")
	descP3.AppendChild(app.CreateTextNode("This toast will not auto-hide. Click the close button to dismiss."))
	persistCardBody.Element.AppendChild(descP3)

	persistToast := app.Toast()
	persistToast.SetAttribute("id", "persistToast")
	persistToast.SetAutoHide(false)
	persistToast.AddHeader("Important", "now", app)
	persistToast.AddCloseButton(app)
	persistToast.SetBody(app.CreateTextNode("This message requires your attention and won't disappear automatically."), app)
	toastContainer.AddToast(persistToast)

	persistBtn := app.Button(bs5.ColorWarning, app.CreateTextNode("Show Persistent Toast"))
	persistBtn.AddEventListener("click", func(dom.Node) {
		showToast("persistToast")
	})
	persistCardBody.Element.AppendChild(persistBtn.Element)
	container.AppendChild(persistCard.Element)

	// Section: Custom Delay
	delayCard := app.Card()
	delayCard.AddClass("mb-4")
	delayCard.Header(app.H4(app.CreateTextNode("Custom Delay (10 seconds)")).Element)
	delayCardBody := delayCard.Body()

	descP4 := app.CreateElement("p")
	descP4.AppendChild(app.CreateTextNode("This toast will remain visible for 10 seconds before auto-hiding."))
	delayCardBody.Element.AppendChild(descP4)

	delayToast := app.Toast()
	delayToast.SetAttribute("id", "delayToast")
	delayToast.SetDelay(10000)
	delayToast.AddHeader("Long Toast", "now", app)
	delayToast.AddCloseButton(app)
	delayToast.SetBody(app.CreateTextNode("Watch me stay visible for a full 10 seconds!"), app)
	toastContainer.AddToast(delayToast)

	delayBtn := app.Button(bs5.ColorSuccess, app.CreateTextNode("Show 10 Second Toast"))
	delayBtn.AddEventListener("click", func(dom.Node) {
		showToast("delayToast")
	})
	delayCardBody.Element.AppendChild(delayBtn.Element)
	container.AppendChild(delayCard.Element)

	// Section: Simple Toast (without header)
	simpleCard := app.Card()
	simpleCard.AddClass("mb-4")
	simpleCard.Header(app.H4(app.CreateTextNode("Simple Toast (No Header)")).Element)
	simpleCardBody := simpleCard.Body()

	descP5 := app.CreateElement("p")
	descP5.AppendChild(app.CreateTextNode("A minimal toast without a header, perfect for simple notifications."))
	simpleCardBody.Element.AppendChild(descP5)

	simpleToast := app.Toast()
	simpleToast.SetAttribute("id", "simpleToast")
	simpleToast.AddClass("align-items-center")

	simpleWrapper := app.CreateElement("div")
	simpleWrapper.AddClass("d-flex")

	simpleBody := app.CreateElement("div")
	simpleBody.AddClass("toast-body")
	simpleBody.AppendChild(app.CreateTextNode("Hello, world! This is a simple toast message without a header."))
	simpleWrapper.AppendChild(simpleBody)

	simpleCloseBtn := app.CreateElement("button")
	simpleCloseBtn.SetAttribute("type", "button")
	simpleCloseBtn.AddClass("btn-close")
	simpleCloseBtn.AddClass("me-2")
	simpleCloseBtn.AddClass("m-auto")
	simpleCloseBtn.SetAttribute("data-bs-dismiss", "toast")
	simpleCloseBtn.SetAttribute("aria-label", "Close")
	simpleWrapper.AppendChild(simpleCloseBtn)

	simpleToast.Element.AppendChild(simpleWrapper)
	toastContainer.AddToast(simpleToast)

	simpleBtn := app.Button(bs5.ColorSecondary, app.CreateTextNode("Show Simple Toast"))
	simpleBtn.AddEventListener("click", func(dom.Node) {
		showToast("simpleToast")
	})
	simpleCardBody.Element.AppendChild(simpleBtn.Element)
	container.AppendChild(simpleCard.Element)
}
