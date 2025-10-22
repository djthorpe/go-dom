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

type selectElement struct {
	component
}

type option struct {
	component
}

// Ensure that selectElement implements Component interface
var _ Component = (*selectElement)(nil)

// Ensure that option implements Component interface
var _ Component = (*option)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Select creates a new select dropdown element with Bootstrap form-select class.
// Use Append() to add Option components to the select.
//
// Example:
//
//	Select().Append(
//	    Option().Append("Choose..."),
//	    Option(WithValue("1")).Append("Option 1"),
//	    Option(WithValue("2")).Append("Option 2"),
//	)
func Select(opt ...Opt) *selectElement {
	// Create a select element
	root := dom.GetWindow().Document().CreateElement("SELECT")

	// Apply options with base "form-select" class
	if opts, err := NewOpts(SelectComponent, WithClass("form-select")); err != nil {
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
	}

	return &selectElement{
		component: component{
			name: SelectComponent,
			root: root,
		},
	}
}

// Option creates a new option element for use within a Select component.
//
// Example:
//
//	Option(WithValue("1"), WithSelected()).Append("First Option")
func Option(opt ...Opt) *option {
	// Create an option element
	root := dom.GetWindow().Document().CreateElement("OPTION")

	// Apply options
	if opts, err := NewOpts(OptionComponent); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list if any
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			root.SetAttribute(key, value)
		}
	}

	return &option{
		component: component{
			name: OptionComponent,
			root: root,
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// OnChange adds a change event listener to the select element.
// The callback receives the Event when the selection changes.
// Returns the select for method chaining.
//
// Example:
//
//	Select().OnChange(func(e Event) {
//	    // Handle selection change
//	}).Append(...)
func (s *selectElement) OnChange(callback func(Event)) *selectElement {
	s.root.AddEventListener("change", callback)
	return s
}

// Value returns the current value of the selected option.
func (s *selectElement) Value() string {
	return s.root.GetAttribute("value")
}

// SetValue sets the selected option by value.
func (s *selectElement) SetValue(value string) *selectElement {
	s.root.SetAttribute("value", value)
	return s
}
