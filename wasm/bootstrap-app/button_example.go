package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// ButtonExamples returns a container with various button examples
func ButtonExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	// Section heading
	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Button Examples"),
	)

	// Basic button variants
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Button Variants"),
	)

	buttonRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Button(bs.PRIMARY).Append("Primary"),
		bs.Button(bs.SECONDARY).Append("Secondary"),
		bs.Button(bs.SUCCESS).Append("Success"),
		bs.Button(bs.DANGER).Append("Danger"),
		bs.Button(bs.WARNING).Append("Warning"),
		bs.Button(bs.INFO).Append("Info"),
		bs.Button(bs.LIGHT).Append("Light"),
		bs.Button(bs.DARK).Append("Dark"),
	)
	container.Append(buttonRow)

	// Outline buttons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Outline Buttons"),
	)

	outlineRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.OutlineButton(bs.PRIMARY).Append("Primary"),
		bs.OutlineButton(bs.SECONDARY).Append("Secondary"),
		bs.OutlineButton(bs.SUCCESS).Append("Success"),
		bs.OutlineButton(bs.DANGER).Append("Danger"),
		bs.OutlineButton(bs.WARNING).Append("Warning"),
		bs.OutlineButton(bs.INFO).Append("Info"),
		bs.OutlineButton(bs.LIGHT).Append("Light"),
		bs.OutlineButton(bs.DARK).Append("Dark"),
	)
	container.Append(outlineRow)

	// Button sizes
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Button Sizes"),
	)

	sizeRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Button(bs.PRIMARY, bs.WithSize(bs.SizeLarge)).Append("Large Button"),
		bs.Button(bs.PRIMARY).Append("Default Button"),
		bs.Button(bs.PRIMARY, bs.WithSize(bs.SizeSmall)).Append("Small Button"),
	)
	container.Append(sizeRow)

	container.Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 3)).Append("Outline Button Sizes"),
	)

	outlineSizeRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.OutlineButton(bs.SECONDARY, bs.WithSize(bs.SizeLarge)).Append("Large Outline"),
		bs.OutlineButton(bs.SECONDARY).Append("Default Outline"),
		bs.OutlineButton(bs.SECONDARY, bs.WithSize(bs.SizeSmall)).Append("Small Outline"),
	)
	container.Append(outlineSizeRow)

	// Buttons with icons/badges
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Buttons with Badges"),
	)

	badgeButtonRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Button(bs.PRIMARY, bs.WithFlex(bs.CENTER), bs.WithClass("gap-2")).Append(
			"Notifications ",
			bs.Badge(bs.WithColor(bs.LIGHT)).Append("4"),
		),
		bs.Button(bs.SUCCESS, bs.WithFlex(bs.CENTER), bs.WithClass("gap-2")).Append(
			"Messages ",
			bs.PillBadge(bs.WithColor(bs.LIGHT)).Append("12"),
		),
		bs.OutlineButton(bs.DANGER, bs.WithFlex(bs.CENTER), bs.WithClass("gap-2")).Append(
			"Alerts ",
			bs.PillBadge(bs.WithColor(bs.DANGER)).Append("3"),
		),
	)
	container.Append(badgeButtonRow)

	// Button groups
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Button Groups"),
	)

	buttonGroup := bs.ButtonGroup(bs.WithAriaLabel("Basic example")).Append(
		bs.Button(bs.PRIMARY).Append("Left"),
		bs.Button(bs.PRIMARY).Append("Middle"),
		bs.Button(bs.PRIMARY).Append("Right"),
	)
	container.Append(buttonGroup)

	container.Append(bs.Span(bs.WithMargin(bs.START, 3)))

	outlineButtonGroup := bs.ButtonGroup(bs.WithAriaLabel("Outlined example")).Append(
		bs.OutlineButton(bs.SECONDARY).Append("Option 1"),
		bs.OutlineButton(bs.SECONDARY).Append("Option 2"),
		bs.OutlineButton(bs.SECONDARY).Append("Option 3"),
	)
	container.Append(outlineButtonGroup)

	// Toolbar with button groups
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Button Toolbar"),
	)

	toolbar := bs.ButtonToolbar(bs.WithAriaLabel("Toolbar with button groups"), bs.WithClass("gap-2")).Append(
		bs.ButtonGroup(bs.WithAriaLabel("First group"), bs.WithMargin(bs.END, 2)).Append(
			bs.Button(bs.PRIMARY).Append("1"),
			bs.Button(bs.PRIMARY).Append("2"),
			bs.Button(bs.PRIMARY).Append("3"),
			bs.Button(bs.PRIMARY).Append("4"),
		),
		bs.ButtonGroup(bs.WithAriaLabel("Second group"), bs.WithMargin(bs.END, 2)).Append(
			bs.Button(bs.SECONDARY).Append("5"),
			bs.Button(bs.SECONDARY).Append("6"),
			bs.Button(bs.SECONDARY).Append("7"),
		),
		bs.ButtonGroup(bs.WithAriaLabel("Third group")).Append(
			bs.Button(bs.SUCCESS).Append("8"),
		),
	)
	container.Append(toolbar)

	// Button group sizing
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Button Group Sizing"),
	)

	sizingContainer := bs.Container(bs.WithFlex(bs.TOP|bs.BOTTOM), bs.WithClass("gap-3")).Append(
		bs.ButtonGroup(bs.WithAriaLabel("Large button group"), bs.WithSize(bs.SizeLarge)).Append(
			bs.Button(bs.PRIMARY).Append("Left"),
			bs.Button(bs.PRIMARY).Append("Middle"),
			bs.Button(bs.PRIMARY).Append("Right"),
		),
		bs.ButtonGroup(bs.WithAriaLabel("Default button group")).Append(
			bs.Button(bs.PRIMARY).Append("Left"),
			bs.Button(bs.PRIMARY).Append("Middle"),
			bs.Button(bs.PRIMARY).Append("Right"),
		),
		bs.ButtonGroup(bs.WithAriaLabel("Small button group"), bs.WithSize(bs.SizeSmall)).Append(
			bs.Button(bs.PRIMARY).Append("Left"),
			bs.Button(bs.PRIMARY).Append("Middle"),
			bs.Button(bs.PRIMARY).Append("Right"),
		),
	)
	container.Append(sizingContainer)

	// Vertical button group
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Vertical Button Group"),
	)

	verticalGroup := bs.VerticalButtonGroup(bs.WithAriaLabel("Vertical button group")).Append(
		bs.Button(bs.PRIMARY).Append("Button 1"),
		bs.Button(bs.PRIMARY).Append("Button 2"),
		bs.Button(bs.PRIMARY).Append("Button 3"),
		bs.Button(bs.PRIMARY).Append("Button 4"),
	)
	container.Append(verticalGroup)

	// Block buttons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Block Buttons"),
	)

	blockButtonContainer := bs.Container(bs.WithClass("d-grid", "gap-2")).Append(
		bs.Button(bs.PRIMARY).Append("Block Level Button"),
		bs.Button(bs.SECONDARY).Append("Another Block Button"),
	)
	container.Append(blockButtonContainer)

	// Buttons with custom spacing
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Buttons with Custom Spacing"),
	)

	spacingRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("flex-wrap")).Append(
		bs.Button(bs.PRIMARY, bs.WithMargin(bs.END, 2), bs.WithMargin(bs.BOTTOM, 2)).Append("Button 1"),
		bs.Button(bs.SUCCESS, bs.WithMargin(bs.END, 2), bs.WithMargin(bs.BOTTOM, 2)).Append("Button 2"),
		bs.Button(bs.WARNING, bs.WithMargin(bs.END, 2), bs.WithMargin(bs.BOTTOM, 2)).Append("Button 3"),
		bs.Button(bs.DANGER, bs.WithMargin(bs.BOTTOM, 2)).Append("Button 4"),
	)
	container.Append(spacingRow)

	// Buttons with padding
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Buttons with Extra Padding"),
	)

	paddingRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Button(bs.PRIMARY, bs.WithPadding(bs.PaddingAll, 1)).Append("Small Padding"),
		bs.Button(bs.PRIMARY).Append("Default Padding"),
		bs.Button(bs.PRIMARY, bs.WithPadding(bs.PaddingAll, 3)).Append("Large Padding"),
	)
	container.Append(paddingRow)

	// Mixed examples
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Mixed Button Styles"),
	)

	mixedRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Button(bs.SUCCESS, bs.WithSize(bs.SizeLarge), bs.WithClass("shadow")).Append("Large with Shadow"),
		bs.OutlineButton(bs.DANGER, bs.WithSize(bs.SizeSmall), bs.WithClass("text-uppercase")).Append("Small Uppercase"),
		bs.Button(bs.INFO, bs.WithClass("rounded-pill")).Append("Rounded Pill"),
		bs.OutlineButton(bs.WARNING, bs.WithClass("rounded-0")).Append("No Border Radius"),
	)
	container.Append(mixedRow)

	// Button states example (using divs as visual examples)
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Button States (Visual)"),
	)

	statesRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithClass("gap-2", "flex-wrap")).Append(
		bs.Button(bs.PRIMARY).Append("Normal State"),
		bs.Button(bs.PRIMARY, bs.WithClass("active")).Append("Active State"),
		bs.Button(bs.PRIMARY, bs.WithClass("disabled")).Append("Disabled Appearance"),
	)
	container.Append(statesRow)

	return container
}
