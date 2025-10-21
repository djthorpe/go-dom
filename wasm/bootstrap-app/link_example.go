package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// LinkExamples returns a container with various link utility examples
func LinkExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	// Section heading
	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Link Utility Examples"),
	)

	// Colored Links
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Colored Links"),
	)

	coloredLinksContainer := bs.Container(bs.WithClass("d-flex", "flex-column", "gap-2"))
	coloredLinksContainer.Append(
		bs.Para().Append(bs.Link("#", bs.WithColor(bs.PRIMARY)).Append("Primary link")),
		bs.Para().Append(bs.Link("#", bs.WithColor(bs.SECONDARY)).Append("Secondary link")),
		bs.Para().Append(bs.Link("#", bs.WithColor(bs.SUCCESS)).Append("Success link")),
		bs.Para().Append(bs.Link("#", bs.WithColor(bs.DANGER)).Append("Danger link")),
		bs.Para().Append(bs.Link("#", bs.WithColor(bs.WARNING)).Append("Warning link")),
		bs.Para().Append(bs.Link("#", bs.WithColor(bs.INFO)).Append("Info link")),
		bs.Para().Append(bs.Link("#", bs.WithColor(bs.LIGHT)).Append("Light link")),
		bs.Para().Append(bs.Link("#", bs.WithColor(bs.DARK)).Append("Dark link")),
	)
	container.Append(coloredLinksContainer)

	// Links with Icons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Links with Icons"),
	)

	iconLinksContainer := bs.Container(bs.WithClass("d-flex", "flex-column", "gap-2"))
	iconLinksContainer.Append(
		bs.Para().Append(
			bs.Link("#", bs.WithColor(bs.PRIMARY)).Append(
				bs.Icon("house-door", bs.WithMargin(bs.END, 2)),
				"Home",
			),
		),
		bs.Para().Append(
			bs.Link("#", bs.WithColor(bs.SUCCESS)).Append(
				bs.Icon("download", bs.WithMargin(bs.END, 2)),
				"Download",
			),
		),
		bs.Para().Append(
			bs.Link("#", bs.WithColor(bs.INFO)).Append(
				bs.Icon("info-circle", bs.WithMargin(bs.END, 2)),
				"Learn more",
			),
		),
		bs.Para().Append(
			bs.Link("#", bs.WithColor(bs.DANGER)).Append(
				bs.Icon("exclamation-triangle", bs.WithMargin(bs.END, 2)),
				"Warning",
			),
		),
		bs.Para().Append(
			bs.Link("#", bs.WithColor(bs.PRIMARY)).Append(
				"Next page ",
				bs.Icon("arrow-right", bs.WithMargin(bs.START, 2)),
			),
		),
	)
	container.Append(iconLinksContainer)

	// External Links
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("External Links"),
	)

	externalLinksContainer := bs.Container(bs.WithClass("d-flex", "flex-column", "gap-2"))

	// Create external links and set target/rel after creation
	bootstrapLink := bs.Link("https://getbootstrap.com", bs.WithColor(bs.PRIMARY)).
		Append(
			"Bootstrap Documentation ",
			bs.Icon("box-arrow-up-right", bs.WithMargin(bs.START, 1)),
		)
	bootstrapLink.Element().SetAttribute("target", "_blank")
	bootstrapLink.Element().SetAttribute("rel", "noopener noreferrer")

	githubLink := bs.Link("https://github.com", bs.WithColor(bs.DARK)).
		Append(
			"GitHub ",
			bs.Icon("github", bs.WithMargin(bs.START, 1)),
		)
	githubLink.Element().SetAttribute("target", "_blank")
	githubLink.Element().SetAttribute("rel", "noopener noreferrer")

	externalLinksContainer.Append(
		bs.Para().Append(bootstrapLink),
		bs.Para().Append(githubLink),
	)
	container.Append(externalLinksContainer)

	// Bottom spacing
	container.Append(
		bs.Container(bs.WithMargin(bs.BOTTOM, 5)),
	)

	return container
}
