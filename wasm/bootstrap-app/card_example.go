package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// CardExamples returns a container with various card examples
func CardExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	// Section heading
	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Card Examples"),
	)

	// Basic Card
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Basic Card"),
	)

	basicCard := bs.Card(bs.WithClass("mb-4")).
		Append(bs.Heading(5).Append("Card Title")).
		Append(bs.Para().Append("This is a basic card with some content. Cards are flexible content containers."))

	container.Append(basicCard)

	// Card with Header
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Card with Header"),
	)

	cardWithHeader := bs.Card(bs.WithClass("mb-4")).
		Heading("Featured").
		Append(bs.Heading(5).Append("Special title treatment")).
		Append(bs.Para().Append("With supporting text below as a natural lead-in to additional content.")).
		Append(bs.Button(bs.PRIMARY).Append("Go somewhere"))

	container.Append(cardWithHeader)

	// Card with Footer
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Card with Footer"),
	)

	cardWithFooter := bs.Card(bs.WithClass("mb-4"))
	cardWithFooter.Append(bs.Heading(5).Append("Card Title"))
	cardWithFooter.Append(bs.Para().Append("Some quick example text to build on the card title and make up the bulk of the card's content."))
	cardWithFooter.Footer("Last updated 3 mins ago")

	container.Append(cardWithFooter)

	// Card with Header and Footer
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Card with Header and Footer"),
	)

	cardComplete := bs.Card(bs.WithClass("mb-4")).Heading("Quote")
	cardComplete.Append(bs.Para().Append("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer posuere erat a ante."))
	cardComplete.Footer("Someone famous in ", bs.Link("#").Append("Source Title"))

	container.Append(cardComplete)

	// Colored Cards
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Colored Cards"),
	)

	coloredCardsRow := bs.Container(bs.WithClass("row", "g-3", "mb-4"))

	// Primary Card
	primaryCard := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card(bs.WithColor(bs.PRIMARY)).
				Heading("Primary Card").
				Append(bs.Para().Append("This card uses the primary color scheme.")),
		)

	// Success Card
	successCard := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card(bs.WithColor(bs.SUCCESS)).
				Heading("Success Card").
				Append(bs.Para().Append("This card uses the success color scheme.")),
		)

	// Danger Card
	dangerCard := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card(bs.WithColor(bs.DANGER)).
				Heading("Danger Card").
				Append(bs.Para().Append("This card uses the danger color scheme.")),
		)

	coloredCardsRow.Append(primaryCard, successCard, dangerCard)
	container.Append(coloredCardsRow)

	// Cards with Icons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Cards with Icons"),
	)

	iconCardsRow := bs.Container(bs.WithClass("row", "g-3", "mb-4"))

	// Star Card
	starCard := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card(bs.WithColor(bs.WARNING)).
				Heading(
					bs.Icon("star-fill", bs.WithMargin(bs.END, 2)),
					"Featured",
				).
				Append(bs.Para().Append("This is a featured card with a star icon in the header.")),
		)

	// Info Card
	infoCard := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card(bs.WithColor(bs.INFO)).
				Heading(
					bs.Icon("info-circle-fill", bs.WithMargin(bs.END, 2)),
					"Information",
				).
				Append(bs.Para().Append("This card displays important information.")),
		)

	// Settings Card
	settingsCard := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card(bs.WithColor(bs.SECONDARY)).
				Heading(
					bs.Icon("gear-fill", bs.WithMargin(bs.END, 2)),
					"Settings",
				).
				Append(bs.Para().Append("Configure your preferences here.")),
		)

	iconCardsRow.Append(starCard, infoCard, settingsCard)
	container.Append(iconCardsRow)

	// Card with List Group
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Card with Multiple Elements"),
	)

	complexCard := bs.Card(bs.WithClass("mb-4")).
		Heading(
			bs.Icon("list-ul", bs.WithMargin(bs.END, 2)),
			"Card with List",
		)
	complexCard.Append(bs.Para().Append("Here's a card with various types of content:"))
	complexCard.Append(
		bs.Container(bs.WithClass("list-group", "list-group-flush")).
			Append(
				bs.Container(bs.WithClass("list-group-item")).Append("An item"),
				bs.Container(bs.WithClass("list-group-item")).Append("A second item"),
				bs.Container(bs.WithClass("list-group-item")).Append("A third item"),
			),
	)
	complexCard.Append(
		bs.Container(bs.WithClass("d-flex", "gap-2", "mt-3")).
			Append(
				bs.Button(bs.PRIMARY, bs.WithClass("btn-sm")).Append("Save"),
				bs.Button(bs.SECONDARY, bs.WithClass("btn-sm")).Append("Cancel"),
			),
	)
	complexCard.Footer("Card footer with action buttons")

	container.Append(complexCard)

	// Text-centered Card
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Centered Card"),
	)

	centeredCard := bs.Card(bs.WithClass("text-center", "mb-4")).Heading("Centered Card")
	centeredCard.Append(bs.Heading(5).Append("Special Title"))
	centeredCard.Append(bs.Para().Append("This card has centered text content."))
	centeredCard.Append(bs.Button(bs.PRIMARY).Append("Action"))
	centeredCard.Footer("2 days ago")

	container.Append(centeredCard)

	return container
}
