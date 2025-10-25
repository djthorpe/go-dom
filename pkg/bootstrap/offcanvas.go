package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type offcanvas struct {
	component
	header Element
	body   Element
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Offcanvas creates a new bootstrap offcanvas component.
// An offcanvas is a sidebar component that can be toggled via JavaScript to appear
// from the left (start), right (end), top, or bottom edge of the viewport.
//
// Use WithID() to set the ID (required for toggling).
// Use WithPosition() to set the position (START, END, TOP, BOTTOM).
// The offcanvas automatically creates a header and body section.
//
// Example:
//
//	Offcanvas(
//	    WithID("myMenu"),
//	    WithPosition(START),
//	    WithTheme(DARK),
//	).Header("Menu").Append("Content")
func Offcanvas(opt ...Opt) *offcanvas {
	// Create a new component
	c := newComponent(OffcanvasComponent, dom.GetWindow().Document().CreateElement("DIV"))

	// Apply options with offcanvas class and tabindex
	if err := c.applyTo(c.root, append(opt, WithClass("offcanvas"), WithAttribute("tabindex", "-1"))...); err != nil {
		panic(err)
	}

	// Handle position attribute - convert data-position to offcanvas-{position} class
	if position := c.root.GetAttribute("data-position"); position != "" {
		c.root.ClassList().Add("offcanvas-" + position)
		c.root.RemoveAttribute("data-position")
	}

	// Create header
	header := dom.GetWindow().Document().CreateElement("DIV")
	header.SetAttribute("class", "offcanvas-header")

	// Create body
	body := dom.GetWindow().Document().CreateElement("DIV")
	body.SetAttribute("class", "offcanvas-body")

	// Append header and body to root
	c.root.AppendChild(header)
	c.root.AppendChild(body)

	c.body = body

	return &offcanvas{
		component: *c,
		header:    header,
		body:      body,
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Header adds content to the offcanvas header.
// Accepts string, Component, or Element children.
// Returns *offcanvas to allow method chaining.
// Multiple calls append additional content to the header.
//
// Example:
//
//	offcanvas.Header(Icon("menu"), Heading("Navigation", 5))
func (o *offcanvas) Header(children ...any) *offcanvas {
	// Append children to header
	for _, child := range children {
		if component, ok := child.(Component); ok {
			o.header.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			o.header.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			o.header.AppendChild(node)
		}
	}

	return o
}

// Show makes the offcanvas visible by adding the "show" class.
// Returns *offcanvas to allow method chaining.
//
// Example:
//
//	offcanvas.Show().Header("Menu")
func (o *offcanvas) Show() *offcanvas {
	o.root.ClassList().Add("show")
	return o
}

// Hide makes the offcanvas hidden by removing the "show" class.
// Returns *offcanvas to allow method chaining.
//
// Example:
//
//	offcanvas.Hide()
func (o *offcanvas) Hide() *offcanvas {
	o.root.ClassList().Remove("show")
	return o
}

// Toggle toggles the visibility of the offcanvas by toggling the "show" class.
// Returns *offcanvas to allow method chaining.
//
// Example:
//
//	offcanvas.Toggle()
func (o *offcanvas) Toggle() *offcanvas {
	o.root.ClassList().Toggle("show")
	return o
}
