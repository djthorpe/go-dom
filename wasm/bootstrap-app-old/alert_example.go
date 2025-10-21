package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddAlertExamples adds alert component examples to the app
func AddAlertExamples(app *bs5.App) dom.Element {
	container := app.CreateElement("div")

	container.AppendChild(
		app.H3(app.CreateTextNode("Alert Examples")).Element,
	)

	// Success alert with icon
	successContent := app.CreateElement("span")
	successIcon := app.Icon("check-circle-fill")
	successIcon.AddClass("me-2")
	successContent.AppendChild(successIcon.Element)
	successContent.AppendChild(app.CreateTextNode("Success! Your operation completed successfully."))
	container.AppendChild(
		app.Alert(
			bs5.ColorSuccess,
			successContent,
		).MakeDismissible().Element,
	)

	// Warning alert with icon
	warningContent := app.CreateElement("span")
	warningIcon := app.Icon("exclamation-triangle-fill")
	warningIcon.AddClass("me-2")
	warningContent.AppendChild(warningIcon.Element)
	warningContent.AppendChild(app.CreateTextNode("Warning! Please review this information carefully."))
	container.AppendChild(
		app.Alert(
			bs5.ColorWarning,
			warningContent,
		).Element,
	)

	// Error/Danger alert with icon
	dangerContent := app.CreateElement("span")
	dangerIcon := app.Icon("x-circle-fill")
	dangerIcon.AddClass("me-2")
	dangerContent.AppendChild(dangerIcon.Element)
	dangerContent.AppendChild(app.CreateTextNode("Error! Something went wrong."))
	container.AppendChild(
		app.Alert(
			bs5.ColorDanger,
			dangerContent,
		).MakeDismissible().Element,
	)

	// Info alert with icon
	infoContent := app.CreateElement("span")
	infoIcon := app.Icon("info-circle-fill")
	infoIcon.AddClass("me-2")
	infoContent.AppendChild(infoIcon.Element)
	infoContent.AppendChild(app.CreateTextNode("Info! Here's some helpful information."))
	container.AppendChild(
		app.Alert(
			bs5.ColorInfo,
			infoContent,
		).Element,
	)

	// Primary alert with icon
	primaryContent := app.CreateElement("span")
	primaryIcon := app.Icon("star-fill")
	primaryIcon.AddClass("me-2")
	primaryContent.AppendChild(primaryIcon.Element)
	primaryContent.AppendChild(app.CreateTextNode("Primary alert - the main call to action."))
	container.AppendChild(
		app.Alert(
			bs5.ColorPrimary,
			primaryContent,
		).MakeDismissible().Element,
	)

	return container
}
