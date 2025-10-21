package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// IconExamples returns a container with various Bootstrap icon examples
func IconExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	// Section heading
	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Bootstrap Icons Examples"),
	)

	// Basic Icons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Basic Icons"),
	)

	basicRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithMargin(bs.BOTTOM, 4))
	basicRow.Append(bs.Icon("heart-fill", bs.WithMargin(bs.END, 2)))
	basicRow.Append(bs.Icon("star-fill", bs.WithMargin(bs.END, 2)))
	basicRow.Append(bs.Icon("alarm", bs.WithMargin(bs.END, 2)))
	basicRow.Append(bs.Icon("bootstrap", bs.WithMargin(bs.END, 2)))
	basicRow.Append(bs.Icon("github", bs.WithMargin(bs.END, 2)))
	basicRow.Append(bs.Icon("bell", bs.WithMargin(bs.END, 2)))
	basicRow.Append(bs.Icon("calendar", bs.WithMargin(bs.END, 2)))
	basicRow.Append(bs.Icon("chat-dots", bs.WithMargin(bs.END, 2)))
	container.Append(basicRow)

	// Colored Icons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Colored Icons"),
	)

	coloredRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithMargin(bs.BOTTOM, 4))
	coloredRow.Append(bs.Icon("heart-fill", bs.WithColor(bs.PRIMARY), bs.WithMargin(bs.END, 2)))
	coloredRow.Append(bs.Icon("star-fill", bs.WithColor(bs.WARNING), bs.WithMargin(bs.END, 2)))
	coloredRow.Append(bs.Icon("check-circle-fill", bs.WithColor(bs.SUCCESS), bs.WithMargin(bs.END, 2)))
	coloredRow.Append(bs.Icon("x-circle-fill", bs.WithColor(bs.DANGER), bs.WithMargin(bs.END, 2)))
	coloredRow.Append(bs.Icon("info-circle-fill", bs.WithColor(bs.INFO), bs.WithMargin(bs.END, 2)))
	coloredRow.Append(bs.Icon("exclamation-triangle-fill", bs.WithColor(bs.WARNING), bs.WithMargin(bs.END, 2)))
	coloredRow.Append(bs.Icon("shield-fill-check", bs.WithColor(bs.SUCCESS), bs.WithMargin(bs.END, 2)))
	coloredRow.Append(bs.Icon("cloud-fill", bs.WithColor(bs.INFO), bs.WithMargin(bs.END, 2)))
	container.Append(coloredRow)

	// Sized Icons (using Bootstrap font-size utilities)
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Sized Icons"),
	)

	sizesRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithMargin(bs.BOTTOM, 4))
	sizesRow.Append(bs.Icon("emoji-smile", bs.WithClass("fs-6"), bs.WithMargin(bs.END, 2)))
	sizesRow.Append(bs.Icon("emoji-smile", bs.WithClass("fs-5"), bs.WithMargin(bs.END, 2)))
	sizesRow.Append(bs.Icon("emoji-smile", bs.WithClass("fs-4"), bs.WithMargin(bs.END, 2)))
	sizesRow.Append(bs.Icon("emoji-smile", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 2)))
	sizesRow.Append(bs.Icon("emoji-smile", bs.WithClass("fs-2"), bs.WithMargin(bs.END, 2)))
	sizesRow.Append(bs.Icon("emoji-smile", bs.WithClass("fs-1"), bs.WithMargin(bs.END, 2)))
	container.Append(sizesRow)

	// Icons in Buttons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Icons in Buttons"),
	)

	buttonsRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithMargin(bs.BOTTOM, 4))

	btn1 := bs.Button(bs.PRIMARY, bs.WithMargin(bs.END, 2))
	btn1.Append(bs.Icon("download", bs.WithMargin(bs.END, 2)))
	btn1.Append("Download")
	buttonsRow.Append(btn1)

	btn2 := bs.Button(bs.SUCCESS, bs.WithMargin(bs.END, 2))
	btn2.Append(bs.Icon("check-circle", bs.WithMargin(bs.END, 2)))
	btn2.Append("Confirm")
	buttonsRow.Append(btn2)

	btn3 := bs.Button(bs.DANGER, bs.WithMargin(bs.END, 2))
	btn3.Append(bs.Icon("trash", bs.WithMargin(bs.END, 2)))
	btn3.Append("Delete")
	buttonsRow.Append(btn3)

	btn4 := bs.Button(bs.INFO, bs.WithMargin(bs.END, 2))
	btn4.Append("Save ")
	btn4.Append(bs.Icon("save", bs.WithMargin(bs.START, 2)))
	buttonsRow.Append(btn4)

	container.Append(buttonsRow)

	// Icons in Alerts
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Icons in Alerts"),
	)

	successAlert := bs.Alert(bs.WithColor(bs.SUCCESS), bs.WithMargin(bs.BOTTOM, 2), bs.WithFlex(bs.CENTER))
	successAlert.Append(bs.Icon("check-circle-fill", bs.WithMargin(bs.END, 2)))
	successAlert.Append("Your changes have been saved successfully!")
	container.Append(successAlert)

	warningAlert := bs.Alert(bs.WithColor(bs.WARNING), bs.WithMargin(bs.BOTTOM, 2), bs.WithFlex(bs.CENTER))
	warningAlert.Append(bs.Icon("exclamation-triangle-fill", bs.WithMargin(bs.END, 2)))
	warningAlert.Append("Please review your input before proceeding.")
	container.Append(warningAlert)

	dangerAlert := bs.Alert(bs.WithColor(bs.DANGER), bs.WithMargin(bs.BOTTOM, 2), bs.WithFlex(bs.CENTER))
	dangerAlert.Append(bs.Icon("x-circle-fill", bs.WithMargin(bs.END, 2)))
	dangerAlert.Append("An error occurred while processing your request.")
	container.Append(dangerAlert)

	infoAlert := bs.Alert(bs.WithColor(bs.INFO), bs.WithMargin(bs.BOTTOM, 4), bs.WithFlex(bs.CENTER))
	infoAlert.Append(bs.Icon("info-circle-fill", bs.WithMargin(bs.END, 2)))
	infoAlert.Append("For more information, please visit our help center.")
	container.Append(infoAlert)

	// Icons with Badges
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Icons with Badges"),
	)

	badgesRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithMargin(bs.BOTTOM, 4))

	notifContainer := bs.Span(bs.WithMargin(bs.END, 3), bs.WithClass("position-relative"))
	notifContainer.Append(bs.Icon("bell", bs.WithClass("fs-2")))
	notifBadge := bs.Badge(bs.WithColor(bs.DANGER), bs.WithClass("rounded-pill"), bs.WithClass("position-absolute"), bs.WithClass("top-0"), bs.WithClass("start-100"), bs.WithClass("translate-middle"))
	notifBadge.Append("5")
	notifContainer.Append(notifBadge)
	badgesRow.Append(notifContainer)

	cartContainer := bs.Span(bs.WithMargin(bs.END, 3), bs.WithClass("position-relative"))
	cartContainer.Append(bs.Icon("cart", bs.WithClass("fs-2")))
	cartBadge := bs.Badge(bs.WithColor(bs.PRIMARY), bs.WithClass("rounded-pill"), bs.WithClass("position-absolute"), bs.WithClass("top-0"), bs.WithClass("start-100"), bs.WithClass("translate-middle"))
	cartBadge.Append("12")
	cartContainer.Append(cartBadge)
	badgesRow.Append(cartContainer)

	mailContainer := bs.Span(bs.WithMargin(bs.END, 3), bs.WithClass("position-relative"))
	mailContainer.Append(bs.Icon("envelope", bs.WithClass("fs-2")))
	mailBadge := bs.Badge(bs.WithColor(bs.SUCCESS), bs.WithClass("rounded-pill"), bs.WithClass("position-absolute"), bs.WithClass("top-0"), bs.WithClass("start-100"), bs.WithClass("translate-middle"))
	mailBadge.Append("3")
	mailContainer.Append(mailBadge)
	badgesRow.Append(mailContainer)

	container.Append(badgesRow)

	// Social Media Icons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Social Media Icons"),
	)

	socialRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithMargin(bs.BOTTOM, 4))
	socialRow.Append(bs.Icon("github", bs.WithClass("fs-2"), bs.WithMargin(bs.END, 3)))
	socialRow.Append(bs.Icon("twitter", bs.WithClass("fs-2"), bs.WithColor(bs.INFO), bs.WithMargin(bs.END, 3)))
	socialRow.Append(bs.Icon("facebook", bs.WithClass("fs-2"), bs.WithColor(bs.PRIMARY), bs.WithMargin(bs.END, 3)))
	socialRow.Append(bs.Icon("linkedin", bs.WithClass("fs-2"), bs.WithColor(bs.PRIMARY), bs.WithMargin(bs.END, 3)))
	socialRow.Append(bs.Icon("instagram", bs.WithClass("fs-2"), bs.WithColor(bs.DANGER), bs.WithMargin(bs.END, 3)))
	socialRow.Append(bs.Icon("youtube", bs.WithClass("fs-2"), bs.WithColor(bs.DANGER), bs.WithMargin(bs.END, 3)))
	container.Append(socialRow)

	// Action Icons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Action Icons"),
	)

	actionsRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithMargin(bs.BOTTOM, 4))
	actionsRow.Append(bs.Icon("search", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 3)))
	actionsRow.Append(bs.Icon("filter", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 3)))
	actionsRow.Append(bs.Icon("gear", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 3)))
	actionsRow.Append(bs.Icon("pencil-square", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 3)))
	actionsRow.Append(bs.Icon("trash3", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 3)))
	actionsRow.Append(bs.Icon("file-earmark-plus", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 3)))
	actionsRow.Append(bs.Icon("upload", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 3)))
	actionsRow.Append(bs.Icon("download", bs.WithClass("fs-3"), bs.WithMargin(bs.END, 3)))
	container.Append(actionsRow)

	// Icon Button Toolbar
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Icon Button Toolbar"),
	)

	// Button toolbar with icon-only buttons
	toolbar := bs.ButtonToolbar(bs.WithMargin(bs.BOTTOM, 3), bs.WithAriaLabel("Toolbar with icon buttons"))

	// First group - editing tools
	editGroup := bs.ButtonGroup(bs.WithMargin(bs.END, 2), bs.WithAriaLabel("Edit tools"))
	editBtn1 := bs.Button(bs.PRIMARY, bs.WithAriaLabel("Bold"))
	editBtn1.Append(bs.Icon("type-bold"))
	editGroup.Append(editBtn1)

	editBtn2 := bs.Button(bs.PRIMARY, bs.WithAriaLabel("Italic"))
	editBtn2.Append(bs.Icon("type-italic"))
	editGroup.Append(editBtn2)

	editBtn3 := bs.Button(bs.PRIMARY, bs.WithAriaLabel("Underline"))
	editBtn3.Append(bs.Icon("type-underline"))
	editGroup.Append(editBtn3)

	toolbar.Append(editGroup)

	// Second group - alignment tools
	alignGroup := bs.ButtonGroup(bs.WithMargin(bs.END, 2), bs.WithAriaLabel("Alignment tools"))
	alignBtn1 := bs.Button(bs.SECONDARY, bs.WithAriaLabel("Align left"))
	alignBtn1.Append(bs.Icon("text-left"))
	alignGroup.Append(alignBtn1)

	alignBtn2 := bs.Button(bs.SECONDARY, bs.WithAriaLabel("Align center"))
	alignBtn2.Append(bs.Icon("text-center"))
	alignGroup.Append(alignBtn2)

	alignBtn3 := bs.Button(bs.SECONDARY, bs.WithAriaLabel("Align right"))
	alignBtn3.Append(bs.Icon("text-right"))
	alignGroup.Append(alignBtn3)

	toolbar.Append(alignGroup)

	// Third group - actions
	actionGroup := bs.ButtonGroup(bs.WithAriaLabel("Actions"))
	actionBtn1 := bs.Button(bs.SUCCESS, bs.WithAriaLabel("Save"))
	actionBtn1.Append(bs.Icon("save"))
	actionGroup.Append(actionBtn1)

	actionBtn2 := bs.Button(bs.WARNING, bs.WithAriaLabel("Edit"))
	actionBtn2.Append(bs.Icon("pencil"))
	actionGroup.Append(actionBtn2)

	actionBtn3 := bs.Button(bs.DANGER, bs.WithAriaLabel("Delete"))
	actionBtn3.Append(bs.Icon("trash"))
	actionGroup.Append(actionBtn3)

	toolbar.Append(actionGroup)

	container.Append(toolbar)

	// Outline button toolbar with icons
	container.Append(
		bs.Heading(5, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 3)).Append("Outline Icon Button Toolbar"),
	)

	outlineToolbar := bs.ButtonToolbar(bs.WithMargin(bs.BOTTOM, 4), bs.WithAriaLabel("Outline toolbar with icon buttons"))

	// Media controls
	mediaGroup := bs.ButtonGroup(bs.WithMargin(bs.END, 2), bs.WithAriaLabel("Media controls"))
	mediaBtn1 := bs.OutlineButton(bs.PRIMARY, bs.WithAriaLabel("Previous"))
	mediaBtn1.Append(bs.Icon("skip-backward-fill"))
	mediaGroup.Append(mediaBtn1)

	mediaBtn2 := bs.OutlineButton(bs.PRIMARY, bs.WithAriaLabel("Play"))
	mediaBtn2.Append(bs.Icon("play-fill"))
	mediaGroup.Append(mediaBtn2)

	mediaBtn3 := bs.OutlineButton(bs.PRIMARY, bs.WithAriaLabel("Pause"))
	mediaBtn3.Append(bs.Icon("pause-fill"))
	mediaGroup.Append(mediaBtn3)

	mediaBtn4 := bs.OutlineButton(bs.PRIMARY, bs.WithAriaLabel("Next"))
	mediaBtn4.Append(bs.Icon("skip-forward-fill"))
	mediaGroup.Append(mediaBtn4)

	outlineToolbar.Append(mediaGroup)

	// Navigation controls
	navGroup := bs.ButtonGroup(bs.WithAriaLabel("Navigation"))
	navBtn1 := bs.OutlineButton(bs.SECONDARY, bs.WithAriaLabel("First page"))
	navBtn1.Append(bs.Icon("chevron-bar-left"))
	navGroup.Append(navBtn1)

	navBtn2 := bs.OutlineButton(bs.SECONDARY, bs.WithAriaLabel("Previous page"))
	navBtn2.Append(bs.Icon("chevron-left"))
	navGroup.Append(navBtn2)

	navBtn3 := bs.OutlineButton(bs.SECONDARY, bs.WithAriaLabel("Next page"))
	navBtn3.Append(bs.Icon("chevron-right"))
	navGroup.Append(navBtn3)

	navBtn4 := bs.OutlineButton(bs.SECONDARY, bs.WithAriaLabel("Last page"))
	navBtn4.Append(bs.Icon("chevron-bar-right"))
	navGroup.Append(navBtn4)

	outlineToolbar.Append(navGroup)

	container.Append(outlineToolbar)

	// Weather Icons
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM|bs.TOP, 3)).Append("Weather Icons"),
	)

	weatherRow := bs.Container(bs.WithFlex(bs.CENTER), bs.WithMargin(bs.BOTTOM, 4))
	weatherRow.Append(bs.Icon("sun-fill", bs.WithClass("fs-2"), bs.WithColor(bs.WARNING), bs.WithMargin(bs.END, 3)))
	weatherRow.Append(bs.Icon("cloud-fill", bs.WithClass("fs-2"), bs.WithColor(bs.INFO), bs.WithMargin(bs.END, 3)))
	weatherRow.Append(bs.Icon("cloud-rain-fill", bs.WithClass("fs-2"), bs.WithColor(bs.PRIMARY), bs.WithMargin(bs.END, 3)))
	weatherRow.Append(bs.Icon("cloud-lightning-fill", bs.WithClass("fs-2"), bs.WithColor(bs.WARNING), bs.WithMargin(bs.END, 3)))
	weatherRow.Append(bs.Icon("snow", bs.WithClass("fs-2"), bs.WithColor(bs.INFO), bs.WithMargin(bs.END, 3)))
	weatherRow.Append(bs.Icon("moon-stars-fill", bs.WithClass("fs-2"), bs.WithColor(bs.DARK), bs.WithMargin(bs.END, 3)))
	container.Append(weatherRow)

	return container
}
