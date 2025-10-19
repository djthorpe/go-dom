package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddAccordionExample adds an accordion component example to the app
func AddAccordionExample(app *bs5.App) dom.Element {
	return app.Container(
		app.H2(app.CreateTextNode("Accordion")).Element,
		func() dom.Element {
			accordion := app.Accordion("accordionExample")

			// First item - expanded by default
			item1 := accordion.AddItem("Accordion Item #1", true)
			item1.Body(
				app.CreateElement("strong").AppendChild(
					app.CreateTextNode("This is the first item's accordion body."),
				),
				app.CreateTextNode(" It is shown by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions."),
			)

			// Second item - collapsed
			item2 := accordion.AddItem("Accordion Item #2", false)
			item2.Body(
				app.CreateElement("strong").AppendChild(
					app.CreateTextNode("This is the second item's accordion body."),
				),
				app.CreateTextNode(" It is hidden by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions."),
			)

			// Third item - collapsed
			item3 := accordion.AddItem("Accordion Item #3", false)
			item3.Body(
				app.CreateElement("strong").AppendChild(
					app.CreateTextNode("This is the third item's accordion body."),
				),
				app.CreateTextNode(" It is hidden by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions."),
			)

			return accordion.Element
		}(),
	).Element
}
