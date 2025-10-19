package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Button struct {
	dom.Element
	color ColorVariant
}

type ButtonDropdown struct {
	dom.Element
	button       dom.Element
	dropdownMenu dom.Element
}

type ButtonDropdownItem struct {
	dom.Element
}

type ButtonDropdownDivider struct {
	dom.Element
}

type ButtonSize string

////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	ButtonSizeSmall  ButtonSize = "sm"
	ButtonSizeMedium ButtonSize = ""
	ButtonSizeLarge  ButtonSize = "lg"
)

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Button creates a Bootstrap 5 button component
func (app *App) Button(color ColorVariant, children ...dom.Node) *Button {
	// Create button element
	btn := app.CreateElement("button")
	btn.AddClass("btn")
	btn.AddClass("btn-" + string(color))
	btn.SetAttribute("type", "button")

	// Add children
	for _, child := range children {
		btn.AppendChild(child)
	}

	return &Button{
		Element: btn,
		color:   color,
	}
}

// ButtonDropdown creates a button with a dropdown menu
func (app *App) ButtonDropdown(label string, color ColorVariant, items ...dom.Node) *ButtonDropdown {
	// Create wrapper div
	div := app.CreateElement("div")
	div.AddClass("dropdown")

	// Create button element
	btn := app.CreateElement("button")
	btn.AddClass("btn")
	btn.AddClass("btn-" + string(color))
	btn.AddClass("dropdown-toggle")
	btn.SetAttribute("type", "button")
	btn.SetAttribute("data-bs-toggle", "dropdown")
	btn.SetAttribute("aria-expanded", "false")
	btn.AppendChild(app.CreateTextNode(label))

	// Create dropdown menu
	ul := app.CreateElement("ul")
	ul.AddClass("dropdown-menu")

	// Add items to dropdown menu
	for _, item := range items {
		ul.AppendChild(item)
	}

	div.AppendChild(btn)
	div.AppendChild(ul)

	return &ButtonDropdown{
		Element:      div,
		button:       btn,
		dropdownMenu: ul,
	}
}

// ButtonDropdownItem creates a dropdown menu item for button dropdowns
func (app *App) ButtonDropdownItem(label string, href string) *ButtonDropdownItem {
	// Create list item
	li := app.CreateElement("li")

	// Create link
	a := app.CreateElement("a")
	a.AddClass("dropdown-item")
	a.SetAttribute("href", href)
	a.AppendChild(app.CreateTextNode(label))

	li.AppendChild(a)

	return &ButtonDropdownItem{
		Element: li,
	}
}

// ButtonDropdownDivider creates a dropdown divider for button dropdowns
func (app *App) ButtonDropdownDivider() *ButtonDropdownDivider {
	// Create list item
	li := app.CreateElement("li")

	// Create hr element
	hr := app.CreateElement("hr")
	hr.AddClass("dropdown-divider")

	li.AppendChild(hr)

	return &ButtonDropdownDivider{
		Element: li,
	}
}

////////////////////////////////////////////////////////////////////////
// METHODS

// SetSize sets the button size
func (b *Button) SetSize(size ButtonSize) *Button {
	// Remove existing size classes
	b.Element.RemoveClass("btn-sm")
	b.Element.RemoveClass("btn-lg")

	// Add new size class if not medium (default)
	if size != ButtonSizeMedium {
		b.Element.AddClass("btn-" + string(size))
	}
	return b
}

// SetOutline changes the button to outline style
func (b *Button) SetOutline(outline bool) *Button {
	if outline {
		// Remove solid style, add outline
		b.Element.RemoveClass("btn-" + string(b.color))
		b.Element.AddClass("btn-outline-" + string(b.color))
	} else {
		// Remove outline, add solid
		b.Element.RemoveClass("btn-outline-" + string(b.color))
		b.Element.AddClass("btn-" + string(b.color))
	}
	return b
}

// SetDisabled sets the disabled state
func (b *Button) SetDisabled(disabled bool) *Button {
	if disabled {
		b.Element.SetAttribute("disabled", "")
	} else {
		// Note: We can't remove attributes yet, would need RemoveAttribute method
		b.Element.SetAttribute("disabled", "false")
	}
	return b
}

// AddClass adds a CSS class to the button
func (b *Button) AddClass(className string) *Button {
	b.Element.AddClass(className)
	return b
}

// RemoveClass removes a CSS class from the button
func (b *Button) RemoveClass(className string) *Button {
	b.Element.RemoveClass(className)
	return b
}

// AddEventListener adds an event listener to the button
func (b *Button) AddEventListener(eventType string, callback func(dom.Node)) *Button {
	b.Element.AddEventListener(eventType, callback)
	return b
}

// Color returns the button's color variant
func (b *Button) Color() ColorVariant {
	return b.color
}

// AddClass adds a CSS class to the button dropdown
func (d *ButtonDropdown) AddClass(className string) *ButtonDropdown {
	d.Element.AddClass(className)
	return d
}

// RemoveClass removes a CSS class from the button dropdown
func (d *ButtonDropdown) RemoveClass(className string) *ButtonDropdown {
	d.Element.RemoveClass(className)
	return d
}

// AddItem adds a dropdown item to the dropdown menu
func (d *ButtonDropdown) AddItem(item dom.Node) *ButtonDropdown {
	d.dropdownMenu.AppendChild(item)
	return d
}

// SetSize sets the button size for the dropdown button
func (d *ButtonDropdown) SetSize(size ButtonSize) *ButtonDropdown {
	// Remove existing size classes
	d.button.RemoveClass("btn-sm")
	d.button.RemoveClass("btn-lg")

	// Add new size class if not medium (default)
	if size != ButtonSizeMedium {
		d.button.AddClass("btn-" + string(size))
	}
	return d
}

// AddClass adds a CSS class to the button dropdown item
func (i *ButtonDropdownItem) AddClass(className string) *ButtonDropdownItem {
	i.Element.AddClass(className)
	return i
}

// RemoveClass removes a CSS class from the button dropdown item
func (i *ButtonDropdownItem) RemoveClass(className string) *ButtonDropdownItem {
	i.Element.RemoveClass(className)
	return i
}

// AddEventListener adds an event listener to the button dropdown item
func (i *ButtonDropdownItem) AddEventListener(eventType string, callback func(dom.Node)) *ButtonDropdownItem {
	i.Element.AddEventListener(eventType, callback)
	return i
}

// AddClass adds a CSS class to the button dropdown divider
func (d *ButtonDropdownDivider) AddClass(className string) *ButtonDropdownDivider {
	d.Element.AddClass(className)
	return d
}

// RemoveClass removes a CSS class from the button dropdown divider
func (d *ButtonDropdownDivider) RemoveClass(className string) *ButtonDropdownDivider {
	d.Element.RemoveClass(className)
	return d
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (b *Button) String() string {
	return "<bs5-button>"
}

func (d *ButtonDropdown) String() string {
	return "<bs5-buttondropdown>"
}

func (i *ButtonDropdownItem) String() string {
	return "<bs5-buttondropdownitem>"
}

func (d *ButtonDropdownDivider) String() string {
	return "<bs5-buttondropdowndivider>"
}
