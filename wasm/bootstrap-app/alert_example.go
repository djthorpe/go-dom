package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// AlertExamples returns a container with various alert examples
func AlertExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	// Section heading
	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Alert Examples"),
	)

	// Basic color alerts
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Basic Alerts"),
	)

	container.Append(
		bs.Alert(bs.WithColor(bs.PRIMARY), bs.WithMargin(bs.BOTTOM, 2)).
			Append("A simple primary alert—check it out!"),
	)

	container.Append(
		bs.Alert(bs.WithColor(bs.SECONDARY), bs.WithMargin(bs.BOTTOM, 2)).
			Append("A simple secondary alert—check it out!"),
	)

	container.Append(
		bs.Alert(bs.WithColor(bs.SUCCESS), bs.WithMargin(bs.BOTTOM, 2)).
			Append("A simple success alert—check it out!"),
	)

	container.Append(
		bs.Alert(bs.WithColor(bs.DANGER), bs.WithMargin(bs.BOTTOM, 2)).
			Append("A simple danger alert—check it out!"),
	)

	container.Append(
		bs.Alert(bs.WithColor(bs.WARNING), bs.WithMargin(bs.BOTTOM, 2)).
			Append("A simple warning alert—check it out!"),
	)

	container.Append(
		bs.Alert(bs.WithColor(bs.INFO), bs.WithMargin(bs.BOTTOM, 2)).
			Append("A simple info alert—check it out!"),
	)

	container.Append(
		bs.Alert(bs.WithColor(bs.LIGHT), bs.WithMargin(bs.BOTTOM, 2)).
			Append("A simple light alert—check it out!"),
	)

	container.Append(
		bs.Alert(bs.WithColor(bs.DARK), bs.WithMargin(bs.BOTTOM, 4)).
			Append("A simple dark alert—check it out!"),
	)

	// Dismissible alerts
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Dismissible Alerts"),
	)

	container.Append(
		bs.DismissibleAlert(bs.WithColor(bs.WARNING), bs.WithMargin(bs.BOTTOM, 2)).
			Append("Holy guacamole! You should check in on some of those fields below."),
	)

	container.Append(
		bs.DismissibleAlert(bs.WithColor(bs.DANGER), bs.WithMargin(bs.BOTTOM, 4)).
			Append("Error! Something went wrong with your submission."),
	)

	// Alerts with additional content
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Alert with Additional Content"),
	)

	successAlert := bs.Alert(
		bs.WithColor(bs.SUCCESS),
		bs.WithMargin(bs.BOTTOM, 4),
	)

	// Use the new Heading method instead of manually creating the heading
	successAlert.Heading("Well done!")

	successAlert.Append(
		bs.Para().Append("Aww yeah, you successfully read this important alert message. This example text is going to run a bit longer so that you can see how spacing within an alert works with this kind of content."),
	)

	successAlert.Append(
		bs.Rule(),
	)

	successAlert.Append(
		bs.Para().Append("Whenever you need to, be sure to use margin utilities to keep things nice and tidy."),
	)

	container.Append(successAlert)

	// Alerts with borders
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Alerts with Borders"),
	)

	container.Append(
		bs.Alert(
			bs.WithColor(bs.PRIMARY),
			bs.WithBorder(bs.BorderAll, bs.PRIMARY),
			bs.WithMargin(bs.BOTTOM, 2),
		).Append("Primary alert with border"),
	)

	container.Append(
		bs.Alert(
			bs.WithColor(bs.DANGER),
			bs.WithBorder(bs.BorderAll, bs.DANGER),
			bs.WithMargin(bs.BOTTOM, 2),
		).Append("Danger alert with border"),
	)

	container.Append(
		bs.Alert(
			bs.WithColor(bs.SUCCESS),
			bs.WithBorder(bs.BorderAll, bs.SUCCESS),
			bs.WithPadding(bs.PaddingAll, 4),
			bs.WithMargin(bs.BOTTOM, 4),
		).Append("Success alert with border and extra padding"),
	)

	// Alerts with badges
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Alerts with Badges"),
	)

	infoAlertWithBadge := bs.Alert(
		bs.WithColor(bs.INFO),
		bs.WithMargin(bs.BOTTOM, 2),
		bs.WithFlex(bs.CENTER),
		bs.WithClass("justify-content-between"),
	)
	infoAlertWithBadge.Append(bs.Span().Append("You have new notifications"))
	infoAlertWithBadge.Append(
		bs.PillBadge(bs.WithColor(bs.PRIMARY)).Append("12"),
	)
	container.Append(infoAlertWithBadge)

	warningAlertWithBadge := bs.Alert(
		bs.WithColor(bs.WARNING),
		bs.WithMargin(bs.BOTTOM, 4),
		bs.WithFlex(bs.CENTER),
		bs.WithClass("justify-content-between"),
	)
	warningAlertWithBadge.Append(bs.Span().Append("Pending tasks require attention"))
	warningAlertWithBadge.Append(
		bs.PillBadge(bs.WithColor(bs.DANGER)).Append("3"),
	)
	container.Append(warningAlertWithBadge)

	// Alerts with headings using Heading method
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Alerts with Headings"),
	)

	// Basic alert with heading
	container.Append(
		bs.Alert(
			bs.WithColor(bs.INFO),
			bs.WithMargin(bs.BOTTOM, 2),
		).Heading("Information").Append("This alert uses the Heading method to add a properly styled alert heading."),
	)

	// Alert with heading and icon
	errorAlert := bs.Alert(
		bs.WithColor(bs.DANGER),
		bs.WithMargin(bs.BOTTOM, 2),
	)
	errorAlert.Heading(
		bs.Icon("exclamation-triangle-fill", bs.WithMargin(bs.END, 2)),
		"Error Occurred",
	)
	errorAlert.Append("An error occurred while processing your request. Please try again.")
	container.Append(errorAlert)

	// Dismissible alert with heading
	dismissibleWithHeading := bs.DismissibleAlert(
		bs.WithColor(bs.SUCCESS),
		bs.WithMargin(bs.BOTTOM, 2),
	)
	dismissibleWithHeading.Heading(
		bs.Icon("check-circle-fill", bs.WithMargin(bs.END, 2)),
		"Success!",
	)
	dismissibleWithHeading.Append("Your changes have been saved successfully. You can now close this alert.")
	container.Append(dismissibleWithHeading)

	// Alert with heading using method chaining
	container.Append(
		bs.Alert(
			bs.WithColor(bs.WARNING),
			bs.WithMargin(bs.BOTTOM, 4),
		).Heading("Warning: ").
			Append("Please review your settings before continuing."),
	)

	return container
}
