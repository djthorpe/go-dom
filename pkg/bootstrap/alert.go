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

type alert struct {
	component
}

// Ensure that alert implements Component interface
var _ Component = (*alert)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Alert creates a new bootstrap alert (div element with role="alert")
// Default alert uses "alert" class. Use WithColor to set color variant.
func Alert(opt ...Opt) *alert {
	// Create an alert div element
	root := dom.GetWindow().Document().CreateElement("DIV")

	// Apply options
	if opts, err := NewOpts(AlertComponent, WithClass("alert")); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set role attribute for accessibility
		root.SetAttribute("role", "alert")
	}

	return &alert{
		component: component{
			name: AlertComponent,
			root: root,
		},
	}
}

// DismissibleAlert creates a new dismissible bootstrap alert with fade and show classes
// and automatically adds a close button
func DismissibleAlert(opt ...Opt) *alert {
	// Append the dismissible classes to options
	opt = append(opt, WithClass("alert-dismissible", "fade", "show"))
	alert := Alert(opt...)

	// Create and append close button
	closeButton := dom.GetWindow().Document().CreateElement("BUTTON")
	closeButton.SetAttribute("type", "button")
	closeButton.SetAttribute("class", "btn-close")
	closeButton.SetAttribute("data-bs-dismiss", "alert")
	closeButton.SetAttribute("aria-label", "Close")

	alert.root.AppendChild(closeButton)

	return alert
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (alert *alert) Element() Element {
	return alert.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (alert *alert) Append(children ...any) Component {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to root
		alert.root.AppendChild(child.(Node))
	}

	// Return the alert for chaining
	return alert
}
