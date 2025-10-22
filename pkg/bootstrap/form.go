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

type form struct {
	component
}

// Ensure that form implements Component interface
var _ Component = (*form)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Form creates a new bootstrap form element.
// The form element wraps form controls and provides styling and structure.
// Use Append() to add form controls and other components to the form.
//
// By default, forms have:
// - novalidate attribute (disables browser validation)
// - needs-validation class (Bootstrap custom validation)
//
// To enable browser validation instead, use WithoutValidation() option.
//
// Example:
//
//	Form(
//	    WithAction("/api/login"),
//	    WithMethod("POST"),
//	).OnSubmit(func(e Event) {
//	    e.PreventDefault()
//	    // Validate and handle form submission
//	    // Add "was-validated" class to show feedback
//	}).Append(
//	    Heading(2).Append("Login"),
//	    Para().Append("Please enter your credentials"),
//	)
func Form(opt ...Opt) *form {
	// Create a form element
	root := dom.GetWindow().Document().CreateElement("FORM")

	// Apply options
	if opts, err := NewOpts(ContainerComponent); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Add default Bootstrap validation classes
		opts.classList.Add("needs-validation")

		// Set default novalidate attribute unless browser validation is explicitly enabled
		if enableBrowser, ok := opts.attributes["_enable_browser_validation"]; !ok || enableBrowser != "true" {
			opts.attributes["novalidate"] = ""
		}
		// Remove the marker attribute
		delete(opts.attributes, "_enable_browser_validation")

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

	return &form{
		component: component{
			name: ContainerComponent,
			root: root,
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// OnSubmit adds a submit event listener to the form.
// The callback receives the Event and can call PreventDefault() to stop form submission.
// Returns the form for method chaining.
//
// Example:
//
//	Form().OnSubmit(func(e Event) {
//	    e.PreventDefault()
//	    // Handle form submission
//	})
func (f *form) OnSubmit(callback func(Event)) *form {
	f.root.AddEventListener("submit", callback)
	return f
}
