package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// ImageExamples returns a container with various image examples
func ImageExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	// Section heading
	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Image Examples"),
	)

	// Responsive Images
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Responsive Images"),
	)

	container.Append(
		bs.Para().Append("Images with the ", bs.Span(bs.WithClass("badge", "bg-secondary")).Append("img-fluid"), " class scale nicely to the parent element."),
	)

	responsiveImage := bs.Image("favicon.png",
		bs.WithClass("img-fluid", "rounded", "mb-4"),
		bs.WithAriaLabel("Responsive placeholder image"),
	)

	container.Append(responsiveImage)

	// Image Thumbnails
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Image Thumbnails"),
	)

	container.Append(
		bs.Para().Append("Add ", bs.Span(bs.WithClass("badge", "bg-secondary")).Append("img-thumbnail"), " to give an image a rounded 1px border appearance."),
	)

	thumbnailsRow := bs.Container(bs.WithClass("row", "g-3", "mb-4"))

	thumb1 := bs.Container(bs.WithClass("col-md-3")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("img-thumbnail"),
				bs.WithAriaLabel("Thumbnail 1"),
			),
		)

	thumb2 := bs.Container(bs.WithClass("col-md-3")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("img-thumbnail"),
				bs.WithAriaLabel("Thumbnail 2"),
			),
		)

	thumb3 := bs.Container(bs.WithClass("col-md-3")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("img-thumbnail"),
				bs.WithAriaLabel("Thumbnail 3"),
			),
		)

	thumb4 := bs.Container(bs.WithClass("col-md-3")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("img-thumbnail"),
				bs.WithAriaLabel("Thumbnail 4"),
			),
		)

	thumbnailsRow.Append(thumb1, thumb2, thumb3, thumb4)
	container.Append(thumbnailsRow)

	// Rounded Images
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Rounded Images"),
	)

	container.Append(
		bs.Para().Append("Use border-radius utilities to round image corners."),
	)

	roundedRow := bs.Container(bs.WithClass("row", "g-3", "mb-4", "text-center"))

	rounded := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("rounded", "mb-2", "w-25"),
				bs.WithAriaLabel("Rounded image"),
			),
			bs.Para(bs.WithClass("small", "text-muted")).Append("Rounded"),
		)

	roundedCircle := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("rounded-circle", "mb-2", "w-25"),
				bs.WithAriaLabel("Circle image"),
			),
			bs.Para(bs.WithClass("small", "text-muted")).Append("Rounded Circle"),
		)

	roundedPill := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("rounded-pill", "mb-2", "w-25"),
				bs.WithAriaLabel("Pill shaped image"),
			),
			bs.Para(bs.WithClass("small", "text-muted")).Append("Rounded Pill"),
		)

	roundedRow.Append(rounded, roundedCircle, roundedPill)
	container.Append(roundedRow)

	// Floating Images
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Floating Images"),
	)

	floatContainer := bs.Container(bs.WithClass("clearfix", "mb-4"))

	floatContainer.Append(
		bs.Image("favicon.png",
			bs.WithClass("float-start", "me-3", "rounded"),
			bs.WithAriaLabel("Float start image"),
		),
		bs.Para().Append("This image floats to the start (left in LTR). Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer posuere erat a ante. Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Cras mattis consectetur purus sit amet fermentum."),
	)

	container.Append(floatContainer)

	floatEndContainer := bs.Container(bs.WithClass("clearfix", "mb-4"))

	floatEndContainer.Append(
		bs.Image("favicon.png",
			bs.WithClass("float-end", "ms-3", "rounded"),
			bs.WithAriaLabel("Float end image"),
		),
		bs.Para().Append("This image floats to the end (right in LTR). Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer posuere erat a ante. Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus."),
	)

	container.Append(floatEndContainer)

	// Images in Cards
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Images in Cards"),
	)

	cardsRow := bs.Container(bs.WithClass("row", "g-3", "mb-4"))

	// Card 1
	cardBody1 := bs.Container(bs.WithClass("card-body"))
	cardBody1.Append(
		bs.Heading(5, bs.WithClass("card-title")).Append("Card with image"),
		bs.Para(bs.WithClass("card-text")).Append("This card has an image at the top using card-img-top class."),
		bs.Button(bs.PRIMARY, bs.WithClass("btn-sm")).Append("Learn more"),
	)

	card1 := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card().
				Append(
					bs.Image("favicon.png",
						bs.WithClass("card-img-top"),
						bs.WithAriaLabel("Card image top"),
					),
				).
				Append(cardBody1),
		)

	// Card 2
	cardBody2 := bs.Container(bs.WithClass("card-body"))
	cardBody2.Append(
		bs.Heading(5, bs.WithClass("card-title")).Append("Image card"),
		bs.Para(bs.WithClass("card-text")).Append("Cards can display images in various positions."),
		bs.Button(bs.SUCCESS, bs.WithClass("btn-sm")).Append("View details"),
	)

	card2 := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card().
				Append(
					bs.Image("favicon.png",
						bs.WithClass("card-img-top"),
						bs.WithAriaLabel("Card thumbnail"),
					),
				).
				Append(cardBody2),
		)

	// Card 3
	cardBody3 := bs.Container(bs.WithClass("card-body"))
	cardBody3.Append(
		bs.Heading(5, bs.WithClass("card-title")).Append("Featured card"),
		bs.Para(bs.WithClass("card-text")).Append("Images make cards more visually appealing."),
		bs.Button(bs.DANGER, bs.WithClass("btn-sm")).Append("Explore"),
	)

	card3 := bs.Container(bs.WithClass("col-md-4")).
		Append(
			bs.Card().
				Append(
					bs.Image("favicon.png",
						bs.WithClass("card-img-top"),
						bs.WithAriaLabel("Card cover"),
					),
				).
				Append(cardBody3),
		)

	cardsRow.Append(card1, card2, card3)
	container.Append(cardsRow)

	// Sized Images
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Sized Images"),
	)

	container.Append(
		bs.Para().Append("Use width utilities to control image size."),
	)

	sizedRow := bs.Container(bs.WithClass("row", "g-3", "mb-4", "text-center"))

	w25 := bs.Container(bs.WithClass("col-12")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("w-25", "rounded", "mb-2"),
				bs.WithAriaLabel("25% width image"),
			),
			bs.Para(bs.WithClass("small", "text-muted")).Append("w-25"),
		)

	w50 := bs.Container(bs.WithClass("col-12")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("w-50", "rounded", "mb-2"),
				bs.WithAriaLabel("50% width image"),
			),
			bs.Para(bs.WithClass("small", "text-muted")).Append("w-50"),
		)

	w75 := bs.Container(bs.WithClass("col-12")).
		Append(
			bs.Image("favicon.png",
				bs.WithClass("w-75", "rounded", "mb-2"),
				bs.WithAriaLabel("75% width image"),
			),
			bs.Para(bs.WithClass("small", "text-muted")).Append("w-75"),
		)

	sizedRow.Append(w25, w50, w75)
	container.Append(sizedRow)

	return container
}
