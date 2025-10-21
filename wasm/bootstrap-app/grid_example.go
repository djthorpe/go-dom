package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddGridExamples adds grid layout examples with cards to the app
func AddGridExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Example 1: Equal width columns with cards
	card1 := app.Card()
	card1.AddClass("mb-4")
	card1.Header(app.H4(app.CreateTextNode("Equal Width Card Grid")).Element)
	card1Body := card1.Body()

	row1 := app.Row()
	row1.SetGutters(3)

	for i := 1; i <= 3; i++ {
		col := app.Col()
		card := app.Card()
		cardBody := card.Body()
		cardBody.Title("Card " + string(rune(i+48)))
		cardBody.Text("This is an equal-width card in a responsive grid.")
		col.AppendChild(card.Element)
		row1.AppendChild(col.Element)
	}

	card1Body.Element.AppendChild(row1.Element)
	container.AppendChild(card1.Element)

	// Example 2: Responsive columns (stacked on mobile, 2 cols on tablet, 3 on desktop)
	card2 := app.Card()
	card2.AddClass("mb-4")
	card2.Header(app.H4(app.CreateTextNode("Responsive Card Grid")).Element)
	card2Body := card2.Body()

	row2 := app.Row()
	row2.SetGutters(4)

	cardData := []struct {
		title string
		text  string
		color bs5.ColorVariant
		icon  string
	}{
		{"Primary", "This card uses primary styling with a house icon.", bs5.ColorPrimary, "house-fill"},
		{"Success", "This card uses success styling with a check icon.", bs5.ColorSuccess, "check-circle-fill"},
		{"Danger", "This card uses danger styling with an X icon.", bs5.ColorDanger, "x-circle-fill"},
		{"Warning", "This card uses warning styling with a warning icon.", bs5.ColorWarning, "exclamation-triangle-fill"},
		{"Info", "This card uses info styling with an info icon.", bs5.ColorInfo, "info-circle-fill"},
		{"Secondary", "This card uses secondary styling with a star icon.", bs5.ColorSecondary, "star-fill"},
	}

	for _, data := range cardData {
		col := app.Col()
		col.SetSize(12)            // Full width on mobile
		col.SetBreakpoint("md", 6) // 2 columns on medium screens
		col.SetBreakpoint("lg", 4) // 3 columns on large screens

		card := app.Card()
		card.SetBorder(data.color)

		cardHeader := card.Header()
		headerContent := app.CreateElement("span")
		icon := app.Icon(data.icon)
		icon.AddClass("me-2")
		icon.SetColor("text-" + string(data.color))
		headerContent.AppendChild(icon.Element)
		headerContent.AppendChild(app.CreateTextNode(data.title))
		cardHeader.Element.AppendChild(headerContent)

		cardBody := card.Body()
		cardBody.Text(data.text)

		col.AppendChild(card.Element)
		row2.AppendChild(col.Element)
	}

	card2Body.Element.AppendChild(row2.Element)
	container.AppendChild(card2.Element)

	// Example 3: Mixed width columns
	card3 := app.Card()
	card3.AddClass("mb-4")
	card3.Header(app.H4(app.CreateTextNode("Mixed Width Card Grid")).Element)
	card3Body := card3.Body()

	row3 := app.Row()
	row3.SetGutters(3)

	// Wide card (8 columns)
	col1 := app.Col()
	col1.SetSize(12)
	col1.SetBreakpoint("md", 8)
	wideCard := app.Card()
	wideCard.SetBackground(bs5.ColorLight)
	wideCardBody := wideCard.Body()
	wideCardBody.Title("Wide Card")
	wideCardBody.Text("This card takes up 8 columns on medium screens and larger, making it wider than the card next to it.")
	col1.AppendChild(wideCard.Element)
	row3.AppendChild(col1.Element)

	// Narrow card (4 columns)
	col2 := app.Col()
	col2.SetSize(12)
	col2.SetBreakpoint("md", 4)
	narrowCard := app.Card()
	narrowCard.SetBorder(bs5.ColorPrimary)
	narrowCardBody := narrowCard.Body()
	narrowCardBody.Title("Narrow Card")
	narrowCardBody.Text("This card takes up 4 columns on medium screens and larger.")
	col2.AppendChild(narrowCard.Element)
	row3.AppendChild(col2.Element)

	card3Body.Element.AppendChild(row3.Element)
	container.AppendChild(card3.Element)

	// Example 4: Four column grid
	card4 := app.Card()
	card4.AddClass("mb-4")
	card4.Header(app.H4(app.CreateTextNode("Four Column Card Grid")).Element)
	card4Body := card4.Body()

	row4 := app.Row()
	row4.SetGuttersY(3)
	row4.SetGuttersX(3)

	icons := []string{"heart-fill", "star-fill", "bell-fill", "gear-fill",
		"envelope-fill", "calendar-fill", "chat-fill", "trophy-fill"}
	colors := []bs5.ColorVariant{bs5.ColorDanger, bs5.ColorWarning, bs5.ColorInfo, bs5.ColorSecondary,
		bs5.ColorPrimary, bs5.ColorSuccess, bs5.ColorDanger, bs5.ColorWarning}

	for i, iconName := range icons {
		col := app.Col()
		col.SetSize(6)             // 2 columns on mobile
		col.SetBreakpoint("md", 3) // 4 columns on medium screens

		card := app.Card()
		card.AddClass("text-center")
		cardBody := card.Body()

		icon := app.Icon(iconName)
		icon.SetSize("fs-1")
		icon.SetColor("text-" + string(colors[i]))
		cardBody.Element.AppendChild(icon.Element)

		col.AppendChild(card.Element)
		row4.AppendChild(col.Element)
	}

	card4Body.Element.AppendChild(row4.Element)
	container.AppendChild(card4.Element)

	return container.Element
}
