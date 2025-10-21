package main

import (
	"fmt"

	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddButtonExamples adds button component examples to the app
func AddButtonExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Heading
	container.AppendChild(app.H2(app.CreateTextNode("Buttons")).Element)

	// Solid Buttons section with icons
	solidDiv := app.CreateElement("div")
	solidDiv.AddClass("mb-3")
	solidDiv.AppendChild(app.CreateTextNode("Solid Buttons with Icons: "))

	primaryContent := app.CreateElement("span")
	primaryIcon := app.Icon("house-fill")
	primaryIcon.AddClass("me-1")
	primaryContent.AppendChild(primaryIcon.Element)
	primaryContent.AppendChild(app.CreateTextNode("Primary"))
	solidDiv.AppendChild(app.Button(bs5.ColorPrimary, primaryContent).
		AddEventListener("click", func(target dom.Node) {
			fmt.Println("Primary button clicked!")
		}).AddClass("me-2").Element)

	successContent := app.CreateElement("span")
	successIcon := app.Icon("check-circle-fill")
	successIcon.AddClass("me-1")
	successContent.AppendChild(successIcon.Element)
	successContent.AppendChild(app.CreateTextNode("Success"))
	solidDiv.AppendChild(app.Button(bs5.ColorSuccess, successContent).AddClass("me-2").Element)

	dangerContent := app.CreateElement("span")
	dangerIcon := app.Icon("x-circle-fill")
	dangerIcon.AddClass("me-1")
	dangerContent.AppendChild(dangerIcon.Element)
	dangerContent.AppendChild(app.CreateTextNode("Danger"))
	solidDiv.AppendChild(app.Button(bs5.ColorDanger, dangerContent).AddClass("me-2").Element)

	warningContent := app.CreateElement("span")
	warningIcon := app.Icon("exclamation-triangle-fill")
	warningIcon.AddClass("me-1")
	warningContent.AppendChild(warningIcon.Element)
	warningContent.AppendChild(app.CreateTextNode("Warning"))
	solidDiv.AppendChild(app.Button(bs5.ColorWarning, warningContent).AddClass("me-2").Element)

	infoContent := app.CreateElement("span")
	infoIcon := app.Icon("info-circle-fill")
	infoIcon.AddClass("me-1")
	infoContent.AppendChild(infoIcon.Element)
	infoContent.AppendChild(app.CreateTextNode("Info"))
	solidDiv.AppendChild(app.Button(bs5.ColorInfo, infoContent).Element)
	container.AppendChild(solidDiv)

	// Outline Buttons section with icons
	outlineDiv := app.CreateElement("div")
	outlineDiv.AddClass("mb-3")
	outlineDiv.AppendChild(app.CreateTextNode("Outline Buttons with Icons: "))

	outlinePrimaryContent := app.CreateElement("span")
	outlinePrimaryIcon := app.Icon("star")
	outlinePrimaryIcon.AddClass("me-1")
	outlinePrimaryContent.AppendChild(outlinePrimaryIcon.Element)
	outlinePrimaryContent.AppendChild(app.CreateTextNode("Primary"))
	outlineDiv.AppendChild(app.Button(bs5.ColorPrimary, outlinePrimaryContent).SetOutline(true).AddClass("me-2").Element)

	outlineSuccessContent := app.CreateElement("span")
	outlineSuccessIcon := app.Icon("heart")
	outlineSuccessIcon.AddClass("me-1")
	outlineSuccessContent.AppendChild(outlineSuccessIcon.Element)
	outlineSuccessContent.AppendChild(app.CreateTextNode("Success"))
	outlineDiv.AppendChild(app.Button(bs5.ColorSuccess, outlineSuccessContent).SetOutline(true).AddClass("me-2").Element)

	outlineDangerContent := app.CreateElement("span")
	outlineDangerIcon := app.Icon("trash")
	outlineDangerIcon.AddClass("me-1")
	outlineDangerContent.AppendChild(outlineDangerIcon.Element)
	outlineDangerContent.AppendChild(app.CreateTextNode("Danger"))
	outlineDiv.AppendChild(app.Button(bs5.ColorDanger, outlineDangerContent).SetOutline(true).Element)
	container.AppendChild(outlineDiv)

	// Button Sizes section with icons
	sizesDiv := app.CreateElement("div")
	sizesDiv.AddClass("mb-3")
	sizesDiv.AppendChild(app.CreateTextNode("Button Sizes with Icons: "))

	smallContent := app.CreateElement("span")
	smallIcon := app.Icon("download")
	smallIcon.AddClass("me-1")
	smallContent.AppendChild(smallIcon.Element)
	smallContent.AppendChild(app.CreateTextNode("Small"))
	sizesDiv.AppendChild(app.Button(bs5.ColorPrimary, smallContent).SetSize(bs5.ButtonSizeSmall).AddClass("me-2").Element)

	mediumContent := app.CreateElement("span")
	mediumIcon := app.Icon("download")
	mediumIcon.AddClass("me-1")
	mediumContent.AppendChild(mediumIcon.Element)
	mediumContent.AppendChild(app.CreateTextNode("Medium"))
	sizesDiv.AppendChild(app.Button(bs5.ColorPrimary, mediumContent).AddClass("me-2").Element)

	largeContent := app.CreateElement("span")
	largeIcon := app.Icon("download")
	largeIcon.AddClass("me-1")
	largeContent.AppendChild(largeIcon.Element)
	largeContent.AppendChild(app.CreateTextNode("Large"))
	sizesDiv.AppendChild(app.Button(bs5.ColorPrimary, largeContent).SetSize(bs5.ButtonSizeLarge).Element)
	container.AppendChild(sizesDiv)

	// Icon-only buttons section
	iconOnlyDiv := app.CreateElement("div")
	iconOnlyDiv.AddClass("mb-3")
	iconOnlyDiv.AppendChild(app.CreateTextNode("Icon-Only Buttons: "))
	iconOnlyDiv.AppendChild(app.Button(bs5.ColorPrimary, app.Icon("heart-fill").Element).AddClass("me-2").Element)
	iconOnlyDiv.AppendChild(app.Button(bs5.ColorSuccess, app.Icon("check-lg").Element).AddClass("me-2").Element)
	iconOnlyDiv.AppendChild(app.Button(bs5.ColorDanger, app.Icon("x-lg").Element).AddClass("me-2").Element)
	iconOnlyDiv.AppendChild(app.Button(bs5.ColorInfo, app.Icon("gear-fill").Element).AddClass("me-2").Element)
	iconOnlyDiv.AppendChild(app.Button(bs5.ColorWarning, app.Icon("bell-fill").Element).Element)
	container.AppendChild(iconOnlyDiv)

	// Action buttons section
	actionDiv := app.CreateElement("div")
	actionDiv.AddClass("mb-3")
	actionDiv.AppendChild(app.CreateTextNode("Common Action Buttons: "))

	saveContent := app.CreateElement("span")
	saveIcon := app.Icon("save")
	saveIcon.AddClass("me-1")
	saveContent.AppendChild(saveIcon.Element)
	saveContent.AppendChild(app.CreateTextNode("Save"))
	actionDiv.AppendChild(app.Button(bs5.ColorSuccess, saveContent).AddClass("me-2").Element)

	editContent := app.CreateElement("span")
	editIcon := app.Icon("pencil-square")
	editIcon.AddClass("me-1")
	editContent.AppendChild(editIcon.Element)
	editContent.AppendChild(app.CreateTextNode("Edit"))
	actionDiv.AppendChild(app.Button(bs5.ColorPrimary, editContent).AddClass("me-2").Element)

	deleteContent := app.CreateElement("span")
	deleteIcon := app.Icon("trash-fill")
	deleteIcon.AddClass("me-1")
	deleteContent.AppendChild(deleteIcon.Element)
	deleteContent.AppendChild(app.CreateTextNode("Delete"))
	actionDiv.AppendChild(app.Button(bs5.ColorDanger, deleteContent).AddClass("me-2").Element)

	uploadContent := app.CreateElement("span")
	uploadIcon := app.Icon("cloud-upload")
	uploadIcon.AddClass("me-1")
	uploadContent.AppendChild(uploadIcon.Element)
	uploadContent.AppendChild(app.CreateTextNode("Upload"))
	actionDiv.AppendChild(app.Button(bs5.ColorInfo, uploadContent).Element)
	container.AppendChild(actionDiv)

	// Button Dropdowns section
	dropdownDiv := app.CreateElement("div")
	dropdownDiv.AddClass("mb-3")
	dropdownDiv.AppendChild(app.CreateTextNode("Button Dropdowns: "))
	dropdownDiv.AppendChild(app.ButtonDropdown(
		"Dropdown button",
		bs5.ColorSecondary,
		app.ButtonDropdownItem("Action", "#").Element,
		app.ButtonDropdownItem("Another action", "#").Element,
		app.ButtonDropdownItem("Something else here", "#").Element,
	).AddClass("d-inline-block").AddClass("me-2").Element)
	dropdownDiv.AppendChild(app.ButtonDropdown(
		"Primary Dropdown",
		bs5.ColorPrimary,
		app.ButtonDropdownItem("First item", "#").Element,
		app.ButtonDropdownDivider().Element,
		app.ButtonDropdownItem("Second item", "#").Element,
		app.ButtonDropdownItem("Third item", "#").AddEventListener("click", func(target dom.Node) {
			fmt.Println("Third item clicked!")
		}).Element,
	).AddClass("d-inline-block").AddClass("me-2").Element)
	dropdownDiv.AppendChild(app.ButtonDropdown(
		"Danger Dropdown",
		bs5.ColorDanger,
		app.ButtonDropdownItem("Delete", "#").Element,
		app.ButtonDropdownItem("Remove", "#").Element,
	).SetSize(bs5.ButtonSizeSmall).AddClass("d-inline-block").Element)
	container.AppendChild(dropdownDiv)

	return container.Element
}
