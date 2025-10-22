package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type FormInput struct {
	dom.Element
}

type FormTextarea struct {
	dom.Element
}

type FormLabel struct {
	dom.Element
}

type FormText struct {
	dom.Element
}

type FormSelect struct {
	dom.Element
}

type FormCheck struct {
	dom.Element
	input dom.Element
	label dom.Element
}

type FormRange struct {
	dom.Element
}

type InputGroup struct {
	dom.Element
}

type InputGroupText struct {
	dom.Element
}

type InputSize string

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	InputSizeSmall  InputSize = "sm"
	InputSizeMedium InputSize = ""
	InputSizeLarge  InputSize = "lg"
)

////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE - INPUT

// FormInput creates a text input with form-control class
func (app *App) FormInput(inputType, id string) *FormInput {
	input := app.CreateElement("input")
	input.AddClass("form-control")
	input.SetAttribute("type", inputType)
	if id != "" {
		input.SetAttribute("id", id)
	}

	return &FormInput{
		Element: input,
	}
}

// FormTextarea creates a textarea with form-control class
func (app *App) FormTextarea(id string, rows int) *FormTextarea {
	textarea := app.CreateElement("textarea")
	textarea.AddClass("form-control")
	if id != "" {
		textarea.SetAttribute("id", id)
	}
	if rows > 0 {
		textarea.SetAttribute("rows", string(rune(rows)))
	}

	return &FormTextarea{
		Element: textarea,
	}
}

// FormLabel creates a label with form-label class
func (app *App) FormLabel(forId string, text dom.Node) *FormLabel {
	label := app.CreateElement("label")
	label.AddClass("form-label")
	if forId != "" {
		label.SetAttribute("for", forId)
	}
	if text != nil {
		label.AppendChild(text)
	}

	return &FormLabel{
		Element: label,
	}
}

// FormText creates help text with form-text class
func (app *App) FormText(id string, text dom.Node) *FormText {
	div := app.CreateElement("div")
	div.AddClass("form-text")
	if id != "" {
		div.SetAttribute("id", id)
	}
	if text != nil {
		div.AppendChild(text)
	}

	return &FormText{
		Element: div,
	}
}

// FormSelect creates a select dropdown with form-select class
func (app *App) FormSelect(id string) *FormSelect {
	select_ := app.CreateElement("select")
	select_.AddClass("form-select")
	if id != "" {
		select_.SetAttribute("id", id)
	}

	return &FormSelect{
		Element: select_,
	}
}

// FormCheck creates a checkbox or radio input with form-check wrapper
func (app *App) FormCheck(inputType, id, label string) *FormCheck {
	wrapper := app.CreateElement("div")
	wrapper.AddClass("form-check")

	input := app.CreateElement("input")
	input.AddClass("form-check-input")
	input.SetAttribute("type", inputType)
	if id != "" {
		input.SetAttribute("id", id)
	}
	wrapper.AppendChild(input)

	var labelEl dom.Element
	if label != "" {
		labelEl = app.CreateElement("label")
		labelEl.AddClass("form-check-label")
		if id != "" {
			labelEl.SetAttribute("for", id)
		}
		labelEl.AppendChild(app.CreateTextNode(label))
		wrapper.AppendChild(labelEl)
	}

	return &FormCheck{
		Element: wrapper,
		input:   input,
		label:   labelEl,
	}
}

// FormRange creates a range input with form-range class
func (app *App) FormRange(id string) *FormRange {
	input := app.CreateElement("input")
	input.AddClass("form-range")
	input.SetAttribute("type", "range")
	if id != "" {
		input.SetAttribute("id", id)
	}

	return &FormRange{
		Element: input,
	}
}

// InputGroup creates an input group wrapper
func (app *App) InputGroup() *InputGroup {
	div := app.CreateElement("div")
	div.AddClass("input-group")

	return &InputGroup{
		Element: div,
	}
}

// InputGroupText creates text addon for input groups
func (app *App) InputGroupText(text dom.Node) *InputGroupText {
	span := app.CreateElement("span")
	span.AddClass("input-group-text")
	if text != nil {
		span.AppendChild(text)
	}

	return &InputGroupText{
		Element: span,
	}
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - FORM INPUT

// SetPlaceholder sets the placeholder text
func (fi *FormInput) SetPlaceholder(placeholder string) *FormInput {
	fi.SetAttribute("placeholder", placeholder)
	return fi
}

// SetValue sets the input value
func (fi *FormInput) SetValue(value string) *FormInput {
	fi.SetAttribute("value", value)
	return fi
}

// SetSize sets the input size
func (fi *FormInput) SetSize(size InputSize) *FormInput {
	if size != InputSizeMedium {
		fi.AddClass("form-control-" + string(size))
	}
	return fi
}

// SetDisabled disables the input
func (fi *FormInput) SetDisabled(disabled bool) *FormInput {
	if disabled {
		fi.SetAttribute("disabled", "")
	}
	return fi
}

// SetReadonly makes the input readonly
func (fi *FormInput) SetReadonly(readonly bool) *FormInput {
	if readonly {
		fi.SetAttribute("readonly", "")
	}
	return fi
}

// SetRequired makes the input required
func (fi *FormInput) SetRequired(required bool) *FormInput {
	if required {
		fi.SetAttribute("required", "")
	}
	return fi
}

// SetAriaDescribedBy sets the aria-describedby attribute
func (fi *FormInput) SetAriaDescribedBy(id string) *FormInput {
	fi.SetAttribute("aria-describedby", id)
	return fi
}

// MakePlaintext converts input to plaintext style
func (fi *FormInput) MakePlaintext() *FormInput {
	fi.RemoveClass("form-control")
	fi.AddClass("form-control-plaintext")
	fi.SetReadonly(true)
	return fi
}

// AddEventListener adds an event listener
func (fi *FormInput) AddEventListener(eventType string, callback func(dom.Event)) *FormInput {
	fi.Element.AddEventListener(eventType, callback)
	return fi
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - FORM TEXTAREA

// SetPlaceholder sets the placeholder text
func (ft *FormTextarea) SetPlaceholder(placeholder string) *FormTextarea {
	ft.SetAttribute("placeholder", placeholder)
	return ft
}

// SetSize sets the textarea size
func (ft *FormTextarea) SetSize(size InputSize) *FormTextarea {
	if size != InputSizeMedium {
		ft.AddClass("form-control-" + string(size))
	}
	return ft
}

// SetDisabled disables the textarea
func (ft *FormTextarea) SetDisabled(disabled bool) *FormTextarea {
	if disabled {
		ft.SetAttribute("disabled", "")
	}
	return ft
}

// SetReadonly makes the textarea readonly
func (ft *FormTextarea) SetReadonly(readonly bool) *FormTextarea {
	if readonly {
		ft.SetAttribute("readonly", "")
	}
	return ft
}

// SetRequired makes the textarea required
func (ft *FormTextarea) SetRequired(required bool) *FormTextarea {
	if required {
		ft.SetAttribute("required", "")
	}
	return ft
}

// AddEventListener adds an event listener
func (ft *FormTextarea) AddEventListener(eventType string, callback func(dom.Event)) *FormTextarea {
	ft.Element.AddEventListener(eventType, callback)
	return ft
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - FORM SELECT

// AddOption adds an option to the select
func (fs *FormSelect) AddOption(value, text string, selected bool) *FormSelect {
	option := fs.OwnerDocument().CreateElement("option")
	option.SetAttribute("value", value)
	option.AppendChild(fs.OwnerDocument().CreateTextNode(text))
	if selected {
		option.SetAttribute("selected", "")
	}
	fs.AppendChild(option)
	return fs
}

// SetSize sets the select size
func (fs *FormSelect) SetSize(size InputSize) *FormSelect {
	if size != InputSizeMedium {
		fs.AddClass("form-select-" + string(size))
	}
	return fs
}

// SetDisabled disables the select
func (fs *FormSelect) SetDisabled(disabled bool) *FormSelect {
	if disabled {
		fs.SetAttribute("disabled", "")
	}
	return fs
}

// SetRequired makes the select required
func (fs *FormSelect) SetRequired(required bool) *FormSelect {
	if required {
		fs.SetAttribute("required", "")
	}
	return fs
}

// SetMultiple allows multiple selections
func (fs *FormSelect) SetMultiple(multiple bool) *FormSelect {
	if multiple {
		fs.SetAttribute("multiple", "")
	}
	return fs
}

// AddEventListener adds an event listener
func (fs *FormSelect) AddEventListener(eventType string, callback func(dom.Event)) *FormSelect {
	fs.Element.AddEventListener(eventType, callback)
	return fs
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - FORM CHECK

// Input returns the input element
func (fc *FormCheck) Input() dom.Element {
	return fc.input
}

// Label returns the label element
func (fc *FormCheck) Label() dom.Element {
	return fc.label
}

// SetChecked sets the checked state
func (fc *FormCheck) SetChecked(checked bool) *FormCheck {
	if checked {
		fc.input.SetAttribute("checked", "")
	}
	return fc
}

// SetDisabled disables the checkbox/radio
func (fc *FormCheck) SetDisabled(disabled bool) *FormCheck {
	if disabled {
		fc.input.SetAttribute("disabled", "")
	}
	return fc
}

// SetRequired makes the checkbox/radio required
func (fc *FormCheck) SetRequired(required bool) *FormCheck {
	if required {
		fc.input.SetAttribute("required", "")
	}
	return fc
}

// SetInline makes the checkbox/radio display inline
func (fc *FormCheck) SetInline(inline bool) *FormCheck {
	if inline {
		fc.AddClass("form-check-inline")
	} else {
		fc.RemoveClass("form-check-inline")
	}
	return fc
}

// MakeSwitch converts checkbox to a switch
func (fc *FormCheck) MakeSwitch() *FormCheck {
	fc.AddClass("form-switch")
	return fc
}

// SetValue sets the input value
func (fc *FormCheck) SetValue(value string) *FormCheck {
	fc.input.SetAttribute("value", value)
	return fc
}

// AddEventListener adds an event listener to the input
func (fc *FormCheck) AddEventListener(eventType string, callback func(dom.Event)) *FormCheck {
	fc.input.AddEventListener(eventType, callback)
	return fc
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - FORM RANGE

// SetMin sets the minimum value
func (fr *FormRange) SetMin(min string) *FormRange {
	fr.SetAttribute("min", min)
	return fr
}

// SetMax sets the maximum value
func (fr *FormRange) SetMax(max string) *FormRange {
	fr.SetAttribute("max", max)
	return fr
}

// SetStep sets the step value
func (fr *FormRange) SetStep(step string) *FormRange {
	fr.SetAttribute("step", step)
	return fr
}

// SetValue sets the current value
func (fr *FormRange) SetValue(value string) *FormRange {
	fr.SetAttribute("value", value)
	return fr
}

// SetDisabled disables the range input
func (fr *FormRange) SetDisabled(disabled bool) *FormRange {
	if disabled {
		fr.SetAttribute("disabled", "")
	}
	return fr
}

// AddEventListener adds an event listener to the range input
func (fr *FormRange) AddEventListener(eventType string, callback func(dom.Node)) *FormRange {
	fr.Element.AddEventListener(eventType, callback)
	return fr
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - INPUT GROUP

// AppendChild adds a child element to the input group
func (ig *InputGroup) AppendChild(child dom.Node) *InputGroup {
	ig.Element.AppendChild(child)
	return ig
}

// SetSize sets the input group size (sm or lg)
func (ig *InputGroup) SetSize(size InputSize) *InputGroup {
	if size != InputSizeMedium {
		ig.AddClass("input-group-" + string(size))
	}
	return ig
}

// Prepend adds an element at the beginning of the input group
func (ig *InputGroup) Prepend(child dom.Node) *InputGroup {
	if ig.Element.FirstChild() != nil {
		ig.Element.InsertBefore(child, ig.Element.FirstChild())
	} else {
		ig.Element.AppendChild(child)
	}
	return ig
}

// Append adds an element at the end of the input group
func (ig *InputGroup) Append(child dom.Node) *InputGroup {
	ig.Element.AppendChild(child)
	return ig
}
