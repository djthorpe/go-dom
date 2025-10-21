//go:build js && wasm

package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

func AddIconExamples(app *bs5.App, container dom.Element) {
	// Add heading
	container.AppendChild(app.H3(app.CreateTextNode("Icon Examples")).Element)

	// Add description
	desc := app.CreateElement("p")
	desc.AppendChild(app.CreateTextNode("Bootstrap Icons provide a library of high quality, open source icons. Visit "))
	link := app.CreateElement("a")
	link.SetAttribute("href", "https://icons.getbootstrap.com/")
	link.SetAttribute("target", "_blank")
	link.AppendChild(app.CreateTextNode("Bootstrap Icons"))
	desc.AppendChild(link)
	desc.AppendChild(app.CreateTextNode(" to browse all available icons."))
	container.AppendChild(desc)

	// Example 1: Basic Icons
	basicCard := app.Card()
	basicCard.AddClass("mb-4")
	container.AppendChild(basicCard.Element)

	basicCardHeader := basicCard.Header()
	basicCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Basic Icons")).Element)

	basicCardBody := basicCard.Body()

	// Row of common icons
	iconRow := app.CreateElement("div")
	iconRow.AddClass("d-flex")
	iconRow.AddClass("gap-3")
	iconRow.AddClass("flex-wrap")

	iconNames := []string{"house", "heart", "star", "envelope", "search", "person", "gear", "bell", "calendar", "chat"}
	for _, name := range iconNames {
		iconWrapper := app.CreateElement("div")
		iconWrapper.AddClass("text-center")

		icon := app.Icon(name)
		icon.SetSize("fs-2")
		iconWrapper.AppendChild(icon.Element)

		iconWrapper.AppendChild(app.CreateElement("br"))

		label := app.CreateElement("small")
		label.AppendChild(app.CreateTextNode(name))
		iconWrapper.AppendChild(label)

		iconRow.AppendChild(iconWrapper)
	}

	basicCardBody.AppendChild(iconRow)

	// Example 2: Icon Sizes
	sizesCard := app.Card()
	sizesCard.AddClass("mb-4")
	container.AppendChild(sizesCard.Element)

	sizesCardHeader := sizesCard.Header()
	sizesCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Icon Sizes")).Element)

	sizesCardBody := sizesCard.Body()

	sizesRow := app.CreateElement("div")
	sizesRow.AddClass("d-flex")
	sizesRow.AddClass("align-items-center")
	sizesRow.AddClass("gap-3")

	sizes := []struct {
		class string
		label string
	}{
		{"fs-1", "fs-1 (largest)"},
		{"fs-2", "fs-2"},
		{"fs-3", "fs-3"},
		{"fs-4", "fs-4"},
		{"fs-5", "fs-5"},
		{"fs-6", "fs-6 (smallest)"},
	}

	for _, size := range sizes {
		sizeWrapper := app.CreateElement("div")
		sizeWrapper.AddClass("text-center")

		icon := app.Icon("star-fill")
		icon.SetSize(size.class)
		sizeWrapper.AppendChild(icon.Element)

		sizeWrapper.AppendChild(app.CreateElement("br"))

		label := app.CreateElement("small")
		label.AppendChild(app.CreateTextNode(size.label))
		sizeWrapper.AppendChild(label)

		sizesRow.AppendChild(sizeWrapper)
	}

	sizesCardBody.AppendChild(sizesRow)

	// Example 3: Icon Colors
	colorsCard := app.Card()
	colorsCard.AddClass("mb-4")
	container.AppendChild(colorsCard.Element)

	colorsCardHeader := colorsCard.Header()
	colorsCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Icon Colors")).Element)

	colorsCardBody := colorsCard.Body()

	colorsRow := app.CreateElement("div")
	colorsRow.AddClass("d-flex")
	colorsRow.AddClass("align-items-center")
	colorsRow.AddClass("gap-3")

	colors := []struct {
		class string
		label string
	}{
		{"text-primary", "Primary"},
		{"text-secondary", "Secondary"},
		{"text-success", "Success"},
		{"text-danger", "Danger"},
		{"text-warning", "Warning"},
		{"text-info", "Info"},
		{"text-dark", "Dark"},
		{"text-muted", "Muted"},
	}

	for _, color := range colors {
		colorWrapper := app.CreateElement("div")
		colorWrapper.AddClass("text-center")

		icon := app.Icon("heart-fill")
		icon.SetSize("fs-1")
		icon.SetColor(color.class)
		colorWrapper.AppendChild(icon.Element)

		colorWrapper.AppendChild(app.CreateElement("br"))

		label := app.CreateElement("small")
		label.AppendChild(app.CreateTextNode(color.label))
		colorWrapper.AppendChild(label)

		colorsRow.AppendChild(colorWrapper)
	}

	colorsCardBody.AppendChild(colorsRow)

	// Example 4: Icons in Buttons
	buttonsCard := app.Card()
	buttonsCard.AddClass("mb-4")
	container.AppendChild(buttonsCard.Element)

	buttonsCardHeader := buttonsCard.Header()
	buttonsCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Icons in Buttons")).Element)

	buttonsCardBody := buttonsCard.Body()

	buttonsRow := app.CreateElement("div")
	buttonsRow.AddClass("d-flex")
	buttonsRow.AddClass("gap-2")
	buttonsRow.AddClass("flex-wrap")

	// Button with icon
	btn1 := app.Button(bs5.ColorPrimary, app.Icon("download").Element)
	btn1.Element.AppendChild(app.CreateTextNode(" Download"))
	buttonsRow.AppendChild(btn1.Element)

	// Button with icon after text
	btn2 := app.Button(bs5.ColorSuccess, app.CreateTextNode("Save "))
	btn2.Element.AppendChild(app.Icon("check-circle").Element)
	buttonsRow.AppendChild(btn2.Element)

	// Icon-only button
	btn3 := app.Button(bs5.ColorDanger, app.Icon("trash").Element)
	buttonsRow.AppendChild(btn3.Element)

	// Icon button with outline
	btn4 := app.Button(bs5.ColorSecondary, app.Icon("gear").Element)
	btn4.Element.AppendChild(app.CreateTextNode(" Settings"))
	btn4.SetOutline(true)
	buttonsRow.AppendChild(btn4.Element)

	buttonsCardBody.AppendChild(buttonsRow)

	// Example 5: Icons with Input Groups
	inputGroupCard := app.Card()
	inputGroupCard.AddClass("mb-4")
	container.AppendChild(inputGroupCard.Element)

	inputGroupCardHeader := inputGroupCard.Header()
	inputGroupCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Icons with Input Groups")).Element)

	inputGroupCardBody := inputGroupCard.Body()

	// Search input with icon
	searchLabel := app.FormLabel("", app.CreateTextNode("Search"))
	inputGroupCardBody.AppendChild(searchLabel.Element)

	searchGroup := app.InputGroup()
	searchIcon := app.InputGroupText(app.Icon("search").Element)
	searchGroup.Prepend(searchIcon.Element)
	searchInput := app.FormInput("text", "search")
	searchInput.SetPlaceholder("Search...")
	searchGroup.Append(searchInput.Element)
	inputGroupCardBody.AppendChild(searchGroup.Element)

	inputGroupCardBody.AppendChild(app.CreateElement("br"))

	// Email input with icon
	emailLabel := app.FormLabel("iconEmail", app.CreateTextNode("Email"))
	inputGroupCardBody.AppendChild(emailLabel.Element)

	emailGroup := app.InputGroup()
	emailIcon := app.InputGroupText(app.Icon("envelope").Element)
	emailGroup.Prepend(emailIcon.Element)
	emailInput := app.FormInput("iconEmail", "email")
	emailInput.SetPlaceholder("user@example.com")
	emailGroup.Append(emailInput.Element)
	inputGroupCardBody.AppendChild(emailGroup.Element)

	inputGroupCardBody.AppendChild(app.CreateElement("br"))

	// Calendar input with icon
	calendarLabel := app.FormLabel("", app.CreateTextNode("Date"))
	inputGroupCardBody.AppendChild(calendarLabel.Element)

	calendarGroup := app.InputGroup()
	calendarIcon := app.InputGroupText(app.Icon("calendar").Element)
	calendarGroup.Prepend(calendarIcon.Element)
	calendarInput := app.FormInput("date", "date")
	calendarGroup.Append(calendarInput.Element)
	inputGroupCardBody.AppendChild(calendarGroup.Element)
}
