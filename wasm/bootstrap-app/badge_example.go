package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// BadgeExamples returns a container with various badge examples
func BadgeExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	// Section heading
	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Badge Examples"),
	)

	// Basic badges
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Basic Badges"),
	)

	badgeRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Badge(bs.WithColor(bs.PRIMARY)).Append("Primary"),
		bs.Badge(bs.WithColor(bs.SECONDARY)).Append("Secondary"),
		bs.Badge(bs.WithColor(bs.SUCCESS)).Append("Success"),
		bs.Badge(bs.WithColor(bs.DANGER)).Append("Danger"),
		bs.Badge(bs.WithColor(bs.WARNING)).Append("Warning"),
		bs.Badge(bs.WithColor(bs.INFO)).Append("Info"),
		bs.Badge(bs.WithColor(bs.LIGHT)).Append("Light"),
		bs.Badge(bs.WithColor(bs.DARK)).Append("Dark"),
	)
	container.Append(badgeRow)

	// Pill badges
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Pill Badges"),
	)

	pillRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.PillBadge(bs.WithColor(bs.PRIMARY)).Append("Primary"),
		bs.PillBadge(bs.WithColor(bs.SECONDARY)).Append("Secondary"),
		bs.PillBadge(bs.WithColor(bs.SUCCESS)).Append("Success"),
		bs.PillBadge(bs.WithColor(bs.DANGER)).Append("Danger"),
		bs.PillBadge(bs.WithColor(bs.WARNING)).Append("Warning"),
		bs.PillBadge(bs.WithColor(bs.INFO)).Append("Info"),
		bs.PillBadge(bs.WithColor(bs.LIGHT)).Append("Light"),
		bs.PillBadge(bs.WithColor(bs.DARK)).Append("Dark"),
	)
	container.Append(pillRow)

	// Badges in headings
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Badges in Headings"),
	)

	container.Append(
		bs.Heading(1).Append(
			"Example heading ",
			bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New"),
		),
	)

	container.Append(
		bs.Heading(2).Append(
			"Example heading ",
			bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New"),
		),
	)

	container.Append(
		bs.Heading(3).Append(
			"Example heading ",
			bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New"),
		),
	)

	container.Append(
		bs.Heading(4).Append(
			"Example heading ",
			bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New"),
		),
	)

	container.Append(
		bs.Heading(5).Append(
			"Example heading ",
			bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New"),
		),
	)

	container.Append(
		bs.Heading(6).Append(
			"Example heading ",
			bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New"),
		),
	)

	// Badges in buttons (using spans styled as buttons)
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Badges as Counters"),
	)

	buttonRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Button(bs.PRIMARY, bs.WithFlex(bs.CENTER), bs.WithClass("gap-2")).Append(
			"Notifications ",
			bs.PillBadge(bs.WithColor(bs.LIGHT)).Append("4"),
		),
		bs.Button(bs.SECONDARY, bs.WithFlex(bs.CENTER), bs.WithClass("gap-2")).Append(
			"Messages ",
			bs.PillBadge(bs.WithColor(bs.LIGHT)).Append("7"),
		),
		bs.Button(bs.SUCCESS, bs.WithFlex(bs.CENTER), bs.WithClass("gap-2")).Append(
			"Updates ",
			bs.PillBadge(bs.WithColor(bs.LIGHT)).Append("12"),
		),
	)
	container.Append(buttonRow)

	// Positioned badges
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Positioned Badges"),
	)

	positionedRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-4", "flex-wrap")).Append(
		bs.Span(bs.WithClass("position-relative")).Append(
			bs.Button(bs.PRIMARY).Append("Inbox"),
			bs.Span(bs.WithClass("position-absolute", "top-0", "start-100", "translate-middle")).Append(
				bs.PillBadge(bs.WithColor(bs.DANGER)).Append("99+"),
			),
		),
		bs.Span(bs.WithClass("position-relative")).Append(
			bs.Button(bs.PRIMARY).Append("Profile"),
			bs.Span(bs.WithClass("position-absolute", "top-0", "start-100", "translate-middle", "p-2", "bg-danger", "border", "border-light", "rounded-circle")).Append(
				bs.Span(bs.WithClass("visually-hidden")).Append("New alerts"),
			),
		),
	)
	container.Append(positionedRow)

	// Badges with padding and margins
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Badges with Custom Spacing"),
	)

	spacingRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Badge(bs.WithColor(bs.PRIMARY), bs.WithPadding(bs.PaddingAll, 3)).Append("Large Padding"),
		bs.Badge(bs.WithColor(bs.SUCCESS), bs.WithPadding(bs.PaddingAll, 1)).Append("Small Padding"),
		bs.PillBadge(bs.WithColor(bs.WARNING), bs.WithMargin(bs.START, 2)).Append("With Margin"),
	)
	container.Append(spacingRow)

	// Badges with borders
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Badges with Borders"),
	)

	borderRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Badge(bs.WithColor(bs.PRIMARY), bs.WithBorder(bs.BorderAll, bs.DARK)).Append("Primary"),
		bs.Badge(bs.WithColor(bs.SUCCESS), bs.WithBorder(bs.BorderAll, bs.DARK)).Append("Success"),
		bs.PillBadge(bs.WithColor(bs.WARNING), bs.WithBorder(bs.BorderAll, bs.DARK)).Append("Warning"),
		bs.PillBadge(bs.WithColor(bs.DANGER), bs.WithBorder(bs.BorderAll, bs.LIGHT)).Append("Danger"),
	)
	container.Append(borderRow)

	// List group with badges
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("List with Badges"),
	)

	listGroup := bs.Container(bs.WithClass("list-group", "w-50")).Append(
		bs.Span(bs.WithFlex(bs.CENTER), bs.WithClass("list-group-item", "justify-content-between")).Append(
			"Inbox",
			bs.PillBadge(bs.WithColor(bs.PRIMARY)).Append("14"),
		),
		bs.Span(bs.WithFlex(bs.CENTER), bs.WithClass("list-group-item", "justify-content-between")).Append(
			"Drafts",
			bs.PillBadge(bs.WithColor(bs.SECONDARY)).Append("2"),
		),
		bs.Span(bs.WithFlex(bs.CENTER), bs.WithClass("list-group-item", "justify-content-between")).Append(
			"Sent",
			bs.PillBadge(bs.WithColor(bs.SUCCESS)).Append("128"),
		),
		bs.Span(bs.WithFlex(bs.CENTER), bs.WithClass("list-group-item", "justify-content-between")).Append(
			"Spam",
			bs.PillBadge(bs.WithColor(bs.DANGER)).Append("3"),
		),
	)
	container.Append(listGroup)

	return container
}
