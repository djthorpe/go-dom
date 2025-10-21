//go:build js && wasm

package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

func AddFormExamples(app *bs5.App, container dom.Element) {
	// Add heading
	container.AppendChild(app.H3(app.CreateTextNode("Form Examples")).Element)

	// Add description
	desc := app.CreateElement("p")
	desc.AppendChild(app.CreateTextNode("Examples and usage guidelines for form control styles, layout options, and custom components."))
	container.AppendChild(desc)

	// Example 1: Basic Form
	basicCard := app.Card()
	basicCard.AddClass("mb-4")
	container.AppendChild(basicCard.Element)

	basicCardHeader := basicCard.Header()
	basicCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Basic Form")).Element)

	basicCardBody := basicCard.Body()

	basicForm := app.CreateElement("form")
	basicCardBody.AppendChild(basicForm)

	// Email input
	emailDiv := app.CreateElement("div")
	emailDiv.AddClass("mb-3")
	basicForm.AppendChild(emailDiv)

	emailLabel := app.FormLabel("exampleInputEmail1", app.CreateTextNode("Email address"))
	emailDiv.AppendChild(emailLabel.Element)

	emailInput := app.FormInput("email", "exampleInputEmail1")
	emailInput.SetPlaceholder("name@example.com")
	emailInput.SetAriaDescribedBy("emailHelp")
	emailDiv.AppendChild(emailInput.Element)

	emailHelp := app.FormText("emailHelp", app.CreateTextNode("We'll never share your email with anyone else."))
	emailDiv.AppendChild(emailHelp.Element)

	// Password input
	passwordDiv := app.CreateElement("div")
	passwordDiv.AddClass("mb-3")
	basicForm.AppendChild(passwordDiv)

	passwordLabel := app.FormLabel("exampleInputPassword1", app.CreateTextNode("Password"))
	passwordDiv.AppendChild(passwordLabel.Element)

	passwordInput := app.FormInput("password", "exampleInputPassword1")
	passwordDiv.AppendChild(passwordInput.Element)

	// Checkbox
	checkDiv := app.CreateElement("div")
	checkDiv.AddClass("mb-3")
	basicForm.AppendChild(checkDiv)

	check := app.FormCheck("checkbox", "exampleCheck1", "Check me out")
	checkDiv.AppendChild(check.Element)

	basicCardFooter := basicCard.Footer()
	submitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	submitBtn.SetAttribute("type", "submit")
	basicCardFooter.Element.AppendChild(submitBtn.Element)

	// Example 2: Input Sizes
	sizesCard := app.Card()
	sizesCard.AddClass("mb-4")
	container.AppendChild(sizesCard.Element)

	sizesCardHeader := sizesCard.Header()
	sizesCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Input Sizes")).Element)

	sizesCardBody := sizesCard.Body()

	largeInput := app.FormInput("text", "")
	largeInput.SetPlaceholder(".form-control-lg").SetSize(bs5.InputSizeLarge)
	sizesCardBody.AppendChild(largeInput.Element)

	app.CreateElement("br")
	sizesCardBody.AppendChild(app.CreateElement("br"))

	defaultInput := app.FormInput("text", "")
	defaultInput.SetPlaceholder("Default input")
	sizesCardBody.AppendChild(defaultInput.Element)

	sizesCardBody.AppendChild(app.CreateElement("br"))

	smallInput := app.FormInput("text", "")
	smallInput.SetPlaceholder(".form-control-sm").SetSize(bs5.InputSizeSmall)
	sizesCardBody.AppendChild(smallInput.Element)

	sizesCardFooter := sizesCard.Footer()
	sizesSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	sizesSubmitBtn.SetAttribute("type", "submit")
	sizesCardFooter.Element.AppendChild(sizesSubmitBtn.Element)

	// Example 3: Textarea
	textareaCard := app.Card()
	textareaCard.AddClass("mb-4")
	container.AppendChild(textareaCard.Element)

	textareaCardHeader := textareaCard.Header()
	textareaCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Textarea")).Element)

	textareaCardBody := textareaCard.Body()

	textareaLabel := app.FormLabel("exampleFormControlTextarea1", app.CreateTextNode("Example textarea"))
	textareaCardBody.AppendChild(textareaLabel.Element)

	textarea := app.FormTextarea("exampleFormControlTextarea1", 3)
	textareaCardBody.AppendChild(textarea.Element)

	textareaCardFooter := textareaCard.Footer()
	textareaSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	textareaSubmitBtn.SetAttribute("type", "submit")
	textareaCardFooter.Element.AppendChild(textareaSubmitBtn.Element)

	// Example 4: Select Menu
	selectCard := app.Card()
	selectCard.AddClass("mb-4")
	container.AppendChild(selectCard.Element)

	selectCardHeader := selectCard.Header()
	selectCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Select Menu")).Element)

	selectCardBody := selectCard.Body()

	selectLabel := app.FormLabel("exampleFormControlSelect1", app.CreateTextNode("Example select"))
	selectCardBody.AppendChild(selectLabel.Element)

	select_ := app.FormSelect("exampleFormControlSelect1")
	select_.AddOption("1", "One", false)
	select_.AddOption("2", "Two", false)
	select_.AddOption("3", "Three", true)
	select_.AddOption("4", "Four", false)
	select_.AddOption("5", "Five", false)
	selectCardBody.AppendChild(select_.Element)

	selectCardFooter := selectCard.Footer()
	selectSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	selectSubmitBtn.SetAttribute("type", "submit")
	selectCardFooter.Element.AppendChild(selectSubmitBtn.Element)

	// Example 5: Disabled and Readonly
	disabledCard := app.Card()
	disabledCard.AddClass("mb-4")
	container.AppendChild(disabledCard.Element)

	disabledCardHeader := disabledCard.Header()
	disabledCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Disabled and Readonly")).Element)

	disabledCardBody := disabledCard.Body()

	disabledInput := app.FormInput("text", "")
	disabledInput.SetPlaceholder("Disabled input").SetDisabled(true)
	disabledCardBody.AppendChild(disabledInput.Element)

	disabledCardBody.AppendChild(app.CreateElement("br"))

	readonlyInput := app.FormInput("text", "")
	readonlyInput.SetValue("Readonly input").SetReadonly(true)
	disabledCardBody.AppendChild(readonlyInput.Element)

	disabledCardFooter := disabledCard.Footer()
	disabledSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	disabledSubmitBtn.SetAttribute("type", "submit")
	disabledCardFooter.Element.AppendChild(disabledSubmitBtn.Element)

	// Example 6: Checkboxes and Radios
	checksCard := app.Card()
	checksCard.AddClass("mb-4")
	container.AppendChild(checksCard.Element)

	checksCardHeader := checksCard.Header()
	checksCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Checkboxes and Radios")).Element)

	checksCardBody := checksCard.Body()

	// Checkboxes
	checksTitle := app.CreateElement("h6")
	checksTitle.AppendChild(app.CreateTextNode("Checkboxes"))
	checksCardBody.AppendChild(checksTitle)

	check1 := app.FormCheck("checkbox", "flexCheckDefault", "Default checkbox")
	checksCardBody.AppendChild(check1.Element)

	check2 := app.FormCheck("checkbox", "flexCheckChecked", "Checked checkbox")
	check2.SetChecked(true)
	checksCardBody.AppendChild(check2.Element)

	check3 := app.FormCheck("checkbox", "flexCheckDisabled", "Disabled checkbox")
	check3.SetDisabled(true)
	checksCardBody.AppendChild(check3.Element)

	checksCardBody.AppendChild(app.CreateElement("br"))

	// Radios
	radiosTitle := app.CreateElement("h6")
	radiosTitle.AppendChild(app.CreateTextNode("Radios"))
	checksCardBody.AppendChild(radiosTitle)

	radio1 := app.FormCheck("radio", "flexRadioDefault1", "Default radio")
	radio1.SetValue("option1")
	radio1.Input().SetAttribute("name", "flexRadioDefault")
	checksCardBody.AppendChild(radio1.Element)

	radio2 := app.FormCheck("radio", "flexRadioDefault2", "Checked radio")
	radio2.SetValue("option2")
	radio2.SetChecked(true)
	radio2.Input().SetAttribute("name", "flexRadioDefault")
	checksCardBody.AppendChild(radio2.Element)

	radio3 := app.FormCheck("radio", "flexRadioDisabled", "Disabled radio")
	radio3.SetDisabled(true)
	radio3.Input().SetAttribute("name", "flexRadioDefault")
	checksCardBody.AppendChild(radio3.Element)

	checksCardFooter := checksCard.Footer()
	checksSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	checksSubmitBtn.SetAttribute("type", "submit")
	checksCardFooter.Element.AppendChild(checksSubmitBtn.Element)

	// Example 7: Switches
	switchesCard := app.Card()
	switchesCard.AddClass("mb-4")
	container.AppendChild(switchesCard.Element)

	switchesCardHeader := switchesCard.Header()
	switchesCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Switches")).Element)

	switchesCardBody := switchesCard.Body()

	switch1 := app.FormCheck("checkbox", "flexSwitchCheckDefault", "Default switch")
	switch1.MakeSwitch()
	switchesCardBody.AppendChild(switch1.Element)

	switch2 := app.FormCheck("checkbox", "flexSwitchCheckChecked", "Checked switch")
	switch2.MakeSwitch().SetChecked(true)
	switchesCardBody.AppendChild(switch2.Element)

	switch3 := app.FormCheck("checkbox", "flexSwitchCheckDisabled", "Disabled switch")
	switch3.MakeSwitch().SetDisabled(true)
	switchesCardBody.AppendChild(switch3.Element)

	switchesCardFooter := switchesCard.Footer()
	switchesSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	switchesSubmitBtn.SetAttribute("type", "submit")
	switchesCardFooter.Element.AppendChild(switchesSubmitBtn.Element)

	// Example 8: Inline Checkboxes and Radios
	inlineCard := app.Card()
	inlineCard.AddClass("mb-4")
	container.AppendChild(inlineCard.Element)

	inlineCardHeader := inlineCard.Header()
	inlineCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Inline Checkboxes and Radios")).Element)

	inlineCardBody := inlineCard.Body()

	inlineCheck1 := app.FormCheck("checkbox", "inlineCheckbox1", "1")
	inlineCheck1.SetInline(true).SetValue("option1")
	inlineCardBody.AppendChild(inlineCheck1.Element)

	inlineCheck2 := app.FormCheck("checkbox", "inlineCheckbox2", "2")
	inlineCheck2.SetInline(true).SetValue("option2")
	inlineCardBody.AppendChild(inlineCheck2.Element)

	inlineCheck3 := app.FormCheck("checkbox", "inlineCheckbox3", "3 (disabled)")
	inlineCheck3.SetInline(true).SetValue("option3").SetDisabled(true)
	inlineCardBody.AppendChild(inlineCheck3.Element)

	inlineCardFooter := inlineCard.Footer()
	inlineSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	inlineSubmitBtn.SetAttribute("type", "submit")
	inlineCardFooter.Element.AppendChild(inlineSubmitBtn.Element)

	// Example 9: File Input
	fileCard := app.Card()
	fileCard.AddClass("mb-4")
	container.AppendChild(fileCard.Element)

	fileCardHeader := fileCard.Header()
	fileCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("File Input")).Element)

	fileCardBody := fileCard.Body()

	fileDiv := app.CreateElement("div")
	fileDiv.AddClass("mb-3")
	fileCardBody.AppendChild(fileDiv)

	fileLabel := app.FormLabel("formFile", app.CreateTextNode("Default file input example"))
	fileDiv.AppendChild(fileLabel.Element)

	fileInput := app.FormInput("file", "formFile")
	fileDiv.AppendChild(fileInput.Element)

	// Multiple file input
	multiFileDiv := app.CreateElement("div")
	multiFileDiv.AddClass("mb-3")
	fileCardBody.AppendChild(multiFileDiv)

	multiFileLabel := app.FormLabel("formFileMultiple", app.CreateTextNode("Multiple files input example"))
	multiFileDiv.AppendChild(multiFileLabel.Element)

	multiFileInput := app.FormInput("file", "formFileMultiple")
	multiFileInput.SetAttribute("multiple", "")
	multiFileDiv.AppendChild(multiFileInput.Element)

	fileCardFooter := fileCard.Footer()
	fileSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	fileSubmitBtn.SetAttribute("type", "submit")
	fileCardFooter.Element.AppendChild(fileSubmitBtn.Element)

	// Example 10: Color Picker
	colorCard := app.Card()
	colorCard.AddClass("mb-4")
	container.AppendChild(colorCard.Element)

	colorCardHeader := colorCard.Header()
	colorCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Color Picker")).Element)

	colorCardBody := colorCard.Body()

	colorLabel := app.FormLabel("exampleColorInput", app.CreateTextNode("Color picker"))
	colorCardBody.AppendChild(colorLabel.Element)

	colorInput := app.FormInput("color", "exampleColorInput")
	colorInput.RemoveClass("form-control")
	colorInput.AddClass("form-control-color")
	colorInput.SetValue("#563d7c")
	colorInput.SetAttribute("title", "Choose your color")
	colorCardBody.AppendChild(colorInput.Element)

	colorCardFooter := colorCard.Footer()
	colorSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	colorSubmitBtn.SetAttribute("type", "submit")
	colorCardFooter.Element.AppendChild(colorSubmitBtn.Element)

	// Example 11: Range
	rangeCard := app.Card()
	rangeCard.AddClass("mb-4")
	container.AppendChild(rangeCard.Element)

	rangeCardHeader := rangeCard.Header()
	rangeCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Range")).Element)

	rangeCardBody := rangeCard.Body()

	// Default range with value display
	defaultRangeLabel := app.FormLabel("customRange1", app.CreateTextNode("Example range"))
	rangeCardBody.AppendChild(defaultRangeLabel.Element)

	defaultRangeWrapper := app.CreateElement("div")
	defaultRangeWrapper.AddClass("d-flex")
	defaultRangeWrapper.AddClass("align-items-center")
	defaultRangeWrapper.AddClass("gap-3")

	defaultRange := app.FormRange("customRange1")
	defaultRange.SetValue("50")
	defaultRange.AddClass("flex-grow-1")
	defaultRangeWrapper.AppendChild(defaultRange.Element)

	defaultRangeValue := app.CreateElement("span")
	defaultRangeValue.AddClass("badge")
	defaultRangeValue.AddClass("bg-secondary")
	defaultRangeValue.SetAttribute("id", "defaultRangeValue")
	defaultRangeValue.AppendChild(app.CreateTextNode("50"))
	defaultRangeWrapper.AppendChild(defaultRangeValue)

	rangeCardBody.AppendChild(defaultRangeWrapper)

	// Set up the value display update
	setupRangeValueDisplay(defaultRange.Element, "defaultRangeValue")

	rangeCardBody.AppendChild(app.CreateElement("br"))

	// Range with min, max, and step with value display
	customRangeLabel := app.FormLabel("customRange2", app.CreateTextNode("Custom range (0-5, step 0.5)"))
	rangeCardBody.AppendChild(customRangeLabel.Element)

	customRangeWrapper := app.CreateElement("div")
	customRangeWrapper.AddClass("d-flex")
	customRangeWrapper.AddClass("align-items-center")
	customRangeWrapper.AddClass("gap-3")

	customRange := app.FormRange("customRange2")
	customRange.SetMin("0").SetMax("5").SetStep("0.5").SetValue("2.5")
	customRange.AddClass("flex-grow-1")
	customRangeWrapper.AppendChild(customRange.Element)

	customRangeValue := app.CreateElement("span")
	customRangeValue.AddClass("badge")
	customRangeValue.AddClass("bg-secondary")
	customRangeValue.SetAttribute("id", "customRangeValue")
	customRangeValue.AppendChild(app.CreateTextNode("2.5"))
	customRangeWrapper.AppendChild(customRangeValue)

	rangeCardBody.AppendChild(customRangeWrapper)

	// Set up the value display update
	setupRangeValueDisplay(customRange.Element, "customRangeValue")

	rangeCardBody.AppendChild(app.CreateElement("br"))

	// Disabled range with value display
	disabledRangeLabel := app.FormLabel("customRange3", app.CreateTextNode("Disabled range"))
	rangeCardBody.AppendChild(disabledRangeLabel.Element)

	disabledRangeWrapper := app.CreateElement("div")
	disabledRangeWrapper.AddClass("d-flex")
	disabledRangeWrapper.AddClass("align-items-center")
	disabledRangeWrapper.AddClass("gap-3")

	disabledRange := app.FormRange("customRange3")
	disabledRange.SetDisabled(true).SetValue("30")
	disabledRange.AddClass("flex-grow-1")
	disabledRangeWrapper.AppendChild(disabledRange.Element)

	disabledRangeValue := app.CreateElement("span")
	disabledRangeValue.AddClass("badge")
	disabledRangeValue.AddClass("bg-secondary")
	disabledRangeValue.AppendChild(app.CreateTextNode("30"))
	disabledRangeWrapper.AppendChild(disabledRangeValue)

	rangeCardBody.AppendChild(disabledRangeWrapper)

	rangeCardFooter := rangeCard.Footer()
	rangeSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	rangeSubmitBtn.SetAttribute("type", "submit")
	rangeCardFooter.Element.AppendChild(rangeSubmitBtn.Element)

	// Example 12: Input Groups
	inputGroupCard := app.Card()
	inputGroupCard.AddClass("mb-4")
	container.AppendChild(inputGroupCard.Element)

	inputGroupCardHeader := inputGroupCard.Header()
	inputGroupCardHeader.Element.AppendChild(app.H4(app.CreateTextNode("Input Groups")).Element)

	inputGroupCardBody := inputGroupCard.Body()

	// Input group with text prepend
	prependLabel := app.FormLabel("", app.CreateTextNode("Username"))
	inputGroupCardBody.AppendChild(prependLabel.Element)

	prependGroup := app.InputGroup()
	prependText := app.InputGroupText(app.CreateTextNode("@"))
	prependGroup.Prepend(prependText.Element)
	prependInput := app.FormInput("text", "username")
	prependInput.SetPlaceholder("Username")
	prependGroup.Append(prependInput.Element)
	inputGroupCardBody.AppendChild(prependGroup.Element)

	inputGroupCardBody.AppendChild(app.CreateElement("br"))

	// Input group with text append
	appendLabel := app.FormLabel("", app.CreateTextNode("Server"))
	inputGroupCardBody.AppendChild(appendLabel.Element)

	appendGroup := app.InputGroup()
	appendInput := app.FormInput("text", "server")
	appendInput.SetPlaceholder("Your server")
	appendGroup.Append(appendInput.Element)
	appendText := app.InputGroupText(app.CreateTextNode(".example.com"))
	appendGroup.Append(appendText.Element)
	inputGroupCardBody.AppendChild(appendGroup.Element)

	inputGroupCardBody.AppendChild(app.CreateElement("br"))

	// Input group with text on both sides
	bothLabel := app.FormLabel("", app.CreateTextNode("Price"))
	inputGroupCardBody.AppendChild(bothLabel.Element)

	bothGroup := app.InputGroup()
	leftText := app.InputGroupText(app.CreateTextNode("$"))
	bothGroup.Prepend(leftText.Element)
	bothInput := app.FormInput("text", "price")
	bothInput.SetPlaceholder("0.00")
	bothGroup.Append(bothInput.Element)
	rightText := app.InputGroupText(app.CreateTextNode(".00"))
	bothGroup.Append(rightText.Element)
	inputGroupCardBody.AppendChild(bothGroup.Element)

	inputGroupCardBody.AppendChild(app.CreateElement("br"))

	// Input group with textarea
	igTextareaLabel := app.FormLabel("", app.CreateTextNode("With textarea"))
	inputGroupCardBody.AppendChild(igTextareaLabel.Element)

	textareaGroup := app.InputGroup()
	textareaGroupText := app.InputGroupText(app.CreateTextNode("Message"))
	textareaGroup.Prepend(textareaGroupText.Element)
	groupTextarea := app.FormTextarea("", 3)
	textareaGroup.Append(groupTextarea.Element)
	inputGroupCardBody.AppendChild(textareaGroup.Element)

	inputGroupCardBody.AppendChild(app.CreateElement("br"))

	// Small input group
	igSmallLabel := app.FormLabel("", app.CreateTextNode("Small"))
	inputGroupCardBody.AppendChild(igSmallLabel.Element)

	smallGroup := app.InputGroup()
	smallGroup.SetSize(bs5.InputSizeSmall)
	smallGroupText := app.InputGroupText(app.CreateTextNode("Small"))
	smallGroup.Prepend(smallGroupText.Element)
	igSmallInput := app.FormInput("text", "")
	smallGroup.Append(igSmallInput.Element)
	inputGroupCardBody.AppendChild(smallGroup.Element)

	inputGroupCardBody.AppendChild(app.CreateElement("br"))

	// Large input group
	igLargeLabel := app.FormLabel("", app.CreateTextNode("Large"))
	inputGroupCardBody.AppendChild(igLargeLabel.Element)

	largeGroup := app.InputGroup()
	largeGroup.SetSize(bs5.InputSizeLarge)
	largeGroupText := app.InputGroupText(app.CreateTextNode("Large"))
	largeGroup.Prepend(largeGroupText.Element)
	igLargeInput := app.FormInput("text", "")
	largeGroup.Append(igLargeInput.Element)
	inputGroupCardBody.AppendChild(largeGroup.Element)

	inputGroupCardFooter := inputGroupCard.Footer()
	inputGroupSubmitBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Submit"))
	inputGroupSubmitBtn.SetAttribute("type", "submit")
	inputGroupCardFooter.Element.AppendChild(inputGroupSubmitBtn.Element)
}
