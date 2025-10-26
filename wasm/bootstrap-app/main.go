package main

import (
	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func main() {
	// Make a new application
	app := mvc.New("Bart App")

	// Create a heading
	app.Append(
		Alerts(),
		Containers(),
		Badges(),
		Rules(),
		Icons(),
	)

	/*
			bs.Container(
				bs.WithBackground(bs.LIGHT),
				bs.WithBorder(bs.Top),
				bs.WithMargin(bs.Top|bs.Bottom, 5),
				bs.WithPadding(bs.All, 5),
			).Append(
				bs.Heading(2).Append("Button Examples"),
				bs.OutlineButton(bs.WithColor(bs.PRIMARY)).Append("Primary Button"),
				bs.Button(bs.WithColor(bs.DANGER), bs.WithSize(bs.SizeLarge), bs.WithMargin(bs.All, 1)).Append("Large Button"),
				bs.Button(bs.WithColor(bs.DANGER), bs.WithSize(bs.SizeSmall), bs.WithMargin(bs.All, 1)).Append("Small Button"),
				bs.Button(bs.WithColor(bs.SUCCESS), bs.WithDisabled(true)).Append("Disabled Button"),
				bs.Heading(2).Append("Button Groups"),
				bs.ButtonGroup().Append(
					bs.Button().Append("Left"),
					bs.Button().Append("Middle"),
					bs.Button().Append("Right"),
				),

				bs.Heading(1, bs.WithColor(bs.PRIMARY)).Append(
					"Heading 1",
					" ",
					bs.Badge(bs.WithColor(bs.LIGHT), bs.WithBorder(bs.Top)).Append("New"),
				),
				bs.Para("Lorum Ipsum"),

				bs.Heading(2, bs.WithColor(bs.SECONDARY)).Append(
					"Heading 2",
					" ",
					bs.PillBadge(bs.WithColor(bs.DARK)).Append("New"),
				),
				bs.Para("Lorum Ipsum"),

				bs.Heading(3, bs.WithColor(bs.SUCCESS)).Append(
					"Heading 3",
					" ",
					bs.Badge(bs.WithColor(bs.DARK)).Append("New"),
				),
				bs.Para("Lorum Ipsum"),

				bs.Heading(4, bs.WithColor(bs.WARNING)).Append("Heading 4"),
				bs.Para(bs.Highlighted("Lorum Ipsum")),

				bs.Heading(5, bs.WithColor(bs.INFO)).Append("Heading 5"),
				bs.Para("Lorum Ipsum"),

				bs.Heading(2).Append("Blockquote Example"),
				bs.Blockquote(bs.Para("A well-known quote, contained in a blockquote element.")),

				bs.Heading(2).Append("Figure Example"),
				bs.Figure().Caption("A caption for the figure").Append(
					bs.RoundedImage("favicon.png"),
				),

				mvc.Div().Append("Hello, World!").AddEventListener("click", func(target Node) {
					if view := mvc.ViewFromNode(target); view != nil {
						fmt.Println("click view:", view.Name(), view)
					}
				}),

				bs.Figure(),
			),
		)

		// Print out the application
		//attr := app.Root().LastElementChild()
		//fmt.Printf("%q\n", attr)
	*/

	// Run the application
	select {}
}
