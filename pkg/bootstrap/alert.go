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
// METHODS

// Heading adds an alert heading at the beginning of the alert.
// It creates a <h4 class="alert-heading"> element and accepts string, Component, or Element children.
// Returns the alert for method chaining.
func (a *alert) Heading(children ...any) *alert {
	// Create h4 element with alert-heading class
	heading := dom.GetWindow().Document().CreateElement("H4")
	heading.SetAttribute("class", "alert-heading")

	// Append children to the heading
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to heading
		heading.AppendChild(child.(Node))
	}

	// Insert the heading at the beginning of the alert
	// If the alert has a close button (dismissible), insert after it
	// Otherwise, insert as first child
	firstChild := a.root.FirstChild()
	if firstChild != nil {
		// Check if first child is a close button
		if elem, ok := firstChild.(Element); ok {
			if elem.TagName() == "BUTTON" && elem.ClassList().Contains("btn-close") {
				// Insert after the close button
				if nextSibling := firstChild.NextSibling(); nextSibling != nil {
					a.root.InsertBefore(heading, nextSibling)
				} else {
					a.root.AppendChild(heading)
				}
			} else {
				// Insert as first child
				a.root.InsertBefore(heading, firstChild)
			}
		} else {
			// Insert as first child
			a.root.InsertBefore(heading, firstChild)
		}
	} else {
		// No children, just append
		a.root.AppendChild(heading)
	}

	return a
}
