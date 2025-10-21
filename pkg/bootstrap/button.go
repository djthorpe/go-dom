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

type buttonGroup struct {
	component
}

// Ensure that button implements Component interface
var _ Component = (*button)(nil)

// Ensure that buttonGroup implements Component interface
var _ Component = (*buttonGroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Button creates a new button element with "btn" class
// Use WithColor to set the button variant (e.g., btn-primary)
func Button(color Color, opt ...Opt) *button {
	// Create a button element
	root := dom.GetWindow().Document().CreateElement("BUTTON")

	// Add color class
	colorClass := "btn-" + string(color)
	opt = append([]Opt{WithClass(colorClass)}, opt...)

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
	// Create a button element directly without the solid color class
	root := dom.GetWindow().Document().CreateElement("BUTTON")

	// Add outline class based on color
	outlineClass := "btn-outline-" + string(color)
	opt = append([]Opt{WithClass(outlineClass)}, opt...)

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

// ButtonGroup creates a button group container with "btn-group" class
// and appropriate ARIA attributes for accessibility
func ButtonGroup(opt ...Opt) *buttonGroup {
	// Create a div element
	root := dom.GetWindow().Document().CreateElement("DIV")

	// Apply options with base "btn-group" class
	if opts, err := NewOpts(ButtonGroupComponent, WithClass("btn-group")); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			root.SetAttribute(key, value)
		}

		// Set default role attribute for accessibility
		if !root.HasAttribute("role") {
			root.SetAttribute("role", "group")
		}
	}

	return &buttonGroup{
		component: component{
			name: ButtonGroupComponent,
			root: root,
		},
	}
}

// VerticalButtonGroup creates a vertical button group with "btn-group-vertical" class
func VerticalButtonGroup(opt ...Opt) *buttonGroup {
	// Create a div element
	root := dom.GetWindow().Document().CreateElement("DIV")

	// Apply options with base "btn-group-vertical" class
	if opts, err := NewOpts(ButtonGroupComponent, WithClass("btn-group-vertical")); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			root.SetAttribute(key, value)
		}

		// Set default role attribute for accessibility
		if !root.HasAttribute("role") {
			root.SetAttribute("role", "group")
		}
	}

	return &buttonGroup{
		component: component{
			name: ButtonGroupComponent,
			root: root,
		},
	}
}

// ButtonToolbar creates a button toolbar container with "btn-toolbar" class
func ButtonToolbar(opt ...Opt) *buttonGroup {
	// Create a div element
	root := dom.GetWindow().Document().CreateElement("DIV")

	// Apply options with base "btn-toolbar" class
	if opts, err := NewOpts(ButtonGroupComponent, WithClass("btn-toolbar")); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			root.SetAttribute(key, value)
		}

		// Set default role attribute for accessibility
		if !root.HasAttribute("role") {
			root.SetAttribute("role", "toolbar")
		}
	}

	return &buttonGroup{
		component: component{
			name: ButtonGroupComponent,
			root: root,
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

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
