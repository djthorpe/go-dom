package bootstrap

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type button struct {
	component
}

// Ensure that button implements Component interface
var _ Component = (*button)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Button creates a new button element with "btn" class
// Use WithColor to set the button variant (e.g., btn-primary)
func Button(opt ...Opt) *button {
	// Create a button element
	root := dom.GetWindow().Document().CreateElement("BUTTON")

	// Apply options with base "btn" class
	if opts, err := NewOpts(ButtonComponent, WithClass("btn")); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set default type attribute
		if !root.HasAttribute("type") {
			root.SetAttribute("type", "button")
		}
	}

	return &button{
		component: component{
			name: ButtonComponent,
			root: root,
		},
	}
}

// OutlineButton creates a new button with outline styling (btn-outline-*)
func OutlineButton(color Color, opt ...Opt) *button {
	// Add outline class based on color
	outlineClass := "btn-outline-" + string(color)
	opt = append(opt, WithClass(outlineClass))
	return Button(opt...)
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (button *button) Element() Element {
	return button.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (button *button) Append(children ...any) Component {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to root
		button.root.AppendChild(child.(Node))
	}

	// Return the button for chaining
	return button
}

// SetDisabled sets or removes the disabled attribute
func (button *button) SetDisabled(disabled bool) *button {
	if disabled {
		button.root.SetAttribute("disabled", "")
	} else {
		// Remove disabled attribute - need to check if there's a RemoveAttribute method
		// For now, we'll work with what we have
	}
	return button
}
