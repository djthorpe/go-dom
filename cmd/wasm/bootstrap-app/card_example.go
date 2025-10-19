package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddCardExamples adds card component examples to the app
func AddCardExamples(app *bs5.App) dom.Element {
	return app.Container(
		app.H2(app.CreateTextNode("Cards")).Element,

		// Basic card with image
		app.CreateTextNode("Basic card with image:"),
		app.CreateElement("br"),
		app.CreateElement("br"),
		func() dom.Element {
			card := app.Card()
			card.SetWidth("18rem")
			card.Image("/favicon.png", "Card image", bs5.CardImgTop)
			body := card.Body()
			body.Title("Card title")
			body.Text("Some quick example text to build on the card title and make up the bulk of the card's content.")
			btn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Go somewhere"))
			btn.Element.SetAttribute("href", "#")
			body.Element.AppendChild(btn.Element)
			return card.Element
		}(),
		app.CreateElement("br"),
		app.CreateElement("br"),

		// Card with header and footer
		app.CreateTextNode("Card with header and footer:"),
		app.CreateElement("br"),
		app.CreateElement("br"),
		func() dom.Element {
			card := app.Card()
			card.SetWidth("18rem")
			card.Header(app.CreateTextNode("Featured"))
			body := card.Body()
			body.Title("Special title treatment")
			body.Text("With supporting text below as a natural lead-in to additional content.")
			btn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Go somewhere"))
			body.Element.AppendChild(btn.Element)
			card.Footer(app.CreateTextNode("2 days ago")).SetMuted(true)
			return card.Element
		}(),
		app.CreateElement("br"),
		app.CreateElement("br"),

		// Card with subtitle and multiple links
		app.CreateTextNode("Card with subtitle and links:"),
		app.CreateElement("br"),
		app.CreateElement("br"),
		func() dom.Element {
			card := app.Card()
			card.SetWidth("18rem")
			body := card.Body()
			body.Title("Card title")
			body.Subtitle("Card subtitle")
			body.Text("Some quick example text to build on the card title and make up the bulk of the card's content.")
			body.Link("Card link", "#")
			body.Element.AppendChild(app.CreateTextNode(" "))
			body.Link("Another link", "#")
			return card.Element
		}(),
		app.CreateElement("br"),
		app.CreateElement("br"),

		// Colored cards
		app.CreateTextNode("Colored cards:"),
		app.CreateElement("br"),
		app.CreateElement("br"),
		func() dom.Element {
			// Create a container div for horizontal layout
			container := app.CreateElement("div")
			style := container.Style()
			if style != nil {
				style.Set("display", "flex")
				style.Set("gap", "1rem")
				style.Set("flex-wrap", "wrap")
			}

			// Primary card
			card1 := app.Card()
			card1.SetWidth("18rem")
			card1.SetBackground(bs5.ColorPrimary)
			card1.SetTextAlign("center")
			card1.AddClass("text-white")
			header1 := card1.Header(app.CreateTextNode("Header"))
			header1.AddClass("text-white")
			body1 := card1.Body()
			body1.Title("Primary card")
			body1.Text("Some quick example text to build on the card title.")
			container.AppendChild(card1.Element)

			// Success card
			card2 := app.Card()
			card2.SetWidth("18rem")
			card2.SetBackground(bs5.ColorSuccess)
			card2.SetTextAlign("center")
			card2.AddClass("text-white")
			header2 := card2.Header(app.CreateTextNode("Header"))
			header2.AddClass("text-white")
			body2 := card2.Body()
			body2.Title("Success card")
			body2.Text("Some quick example text to build on the card title.")
			container.AppendChild(card2.Element)

			// Danger card
			card3 := app.Card()
			card3.SetWidth("18rem")
			card3.SetBackground(bs5.ColorDanger)
			card3.SetTextAlign("center")
			card3.AddClass("text-white")
			header3 := card3.Header(app.CreateTextNode("Header"))
			header3.AddClass("text-white")
			body3 := card3.Body()
			body3.Title("Danger card")
			body3.Text("Some quick example text to build on the card title.")
			container.AppendChild(card3.Element)

			return container
		}(),
		app.CreateElement("br"),
		app.CreateElement("br"),

		// Card with border
		app.CreateTextNode("Card with border variants:"),
		app.CreateElement("br"),
		app.CreateElement("br"),
		func() dom.Element {
			card := app.Card()
			card.SetWidth("18rem")
			card.SetBorder(bs5.ColorPrimary)
			body := card.Body()
			body.Title("Border card")
			body.Text("This card has a primary border color.")
			return card.Element
		}(),
	).Element
}
