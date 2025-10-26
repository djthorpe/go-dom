package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type button struct {
	View
}

type buttongroup struct {
	View
}

var _ ViewWithState = (*button)(nil)
var _ ViewWithGroupState = (*buttongroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewButton      = "mvc-bs-button"
	ViewButtonGroup = "mvc-bs-buttongroup"

	// The prefix class for outline buttons
	viewOutlineButtonClassPrefix = "btn-outline"
)

func init() {
	RegisterView(ViewButton, newButtonFromElement)
	RegisterView(ViewButtonGroup, newButtonGroupFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Button(opt ...Opt) *button {
	opts := append([]Opt{WithAttr("type", "button"), WithClass("btn"), WithClass("btn-primary")}, opt...)
	view := &button{NewView(ViewButton, "BUTTON", opts...)}
	return view
}

func OutlineButton(opt ...Opt) *button {
	opts := append([]Opt{WithAttr("type", "button"), WithClass("btn"), WithClass("btn-outline-primary"), WithClass(viewOutlineButtonClassPrefix)}, opt...)
	view := &button{NewView(ViewButton, "BUTTON", opts...)}
	return view
}

func ButtonGroup(opt ...Opt) *button {
	opts := append([]Opt{WithAttr("role", "group"), WithClass("btn-group")}, opt...)
	view := &button{NewView(ViewButtonGroup, "DIV", opts...)}
	return view
}

func newButtonFromElement(element Element) View {
	if element.TagName() != "BUTTON" {
		return nil
	}
	return &button{NewViewWithElement(element)}
}

func newButtonGroupFromElement(element Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	return &buttongroup{NewViewWithElement(element)}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Return true if button is disabled
func (b *button) Disabled() bool {
	return b.Root().HasAttribute("disabled")
}

// Return true if button is active
func (b *button) Active() bool {
	return b.Root().ClassList().Contains("active")
}

// Return elements which are active in the button group
func (b *buttongroup) Active() []Element {
	var elements []Element

	// Find active elements
	child := b.Root().FirstElementChild()
	for child != nil {
		if child.ClassList().Contains("active") {
			elements = append(elements, child)
		}
		child = child.NextElementSibling()
	}
	return elements
}

// Return elements which are disabled in the button group
func (b *buttongroup) Disabled() []Element {
	var elements []Element

	// Find disabled elements
	child := b.Root().FirstElementChild()
	for child != nil {
		if child.HasAttribute("disabled") {
			elements = append(elements, child)
		}
		child = child.NextElementSibling()
	}
	return elements
}
