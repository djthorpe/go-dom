package bs

import (
	"fmt"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// navbar are elements to create a navigation bar
type navbar struct {
	View
	brand  Element // Caption
	toggle Element // Toggle control
}

// navbar are elements to create a navigation bar
type navitem struct {
	View
}

var _ ViewWithCaption = (*navbar)(nil)
var _ View = (*navitem)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewNavbar      = "mvc-bs-navbar"
	ViewNavbarBrand = "mvc-bs-navbrand"
	ViewNavItem     = "mvc-bs-navitem"
)

func init() {
	RegisterView(ViewNavbar, newNavbarFromElement)
	RegisterView(ViewNavbarBrand, newNavBrandFromElement)
	RegisterView(ViewNavItem, newNavItemFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NavBar(opts ...Opt) *navbar {
	opts = append([]Opt{WithClass("navbar", "bg-primary", "navbar-expand")}, opts...)
	navbar := &navbar{
		View: NewView(ViewNavbar, "NAV", opts...).Body(
			Tag("DIV", WithClass("container-fluid")),
		),
		brand:  Tag("SPAN"), // TODO: We should not use SPAN here
		toggle: Tag("SPAN"), // TODO: We should not use SPAN here
	}

	// The ID for this navbar
	collapseId := "navbar-items"

	// If we have WithSize() then add a toggle control
	if navbar.shouldToggle() {
		navbar.Toggle(collapseId)
	}

	// Create the body view
	body := Tag("UL", WithClass("navbar-nav", "me-auto", "mb-2", "mb-lg-0"))

	// Append brand, toggle and body to navbar
	navbar.View.Append(
		navbar.brand,
		navbar.toggle,
		Tag("DIV", WithID(collapseId), WithClass("collapse", "navbar-collapse")).Append(body),
	)

	// Set the body of the navbar, where items are added
	navbar.Body(body)

	// Return the navbar
	return navbar
}

func NavItem(content ...any) *navitem {
	opts := []Opt{WithClass("nav-item")}
	return &navitem{
		NewView(
			ViewNavItem, "LI", opts...,
		).Body(
			Tag("A", WithClass("nav-link"), WithAttr("href", "#")).Append(content...),
		),
	}
}

func newNavbarFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "NAV" {
		panic(fmt.Sprintf("newNavbarFromElement: invalid tag name %q", tagName))
	}
	view := &navbar{NewViewWithElement(element), nil, nil}
	// TODO: set body, toggle and brand elements
	return view
}

func newNavBrandFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "A" {
		panic(fmt.Sprintf("newNavBrandFromElement: invalid tag name %q", tagName))
	}
	// TODO: set body
	return NewViewWithElement(element)
}

func newNavItemFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "LI" {
		panic(fmt.Sprintf("newNavItemFromElement: invalid tag name %q", tagName))
	}
	return &navitem{NewViewWithElement(element)}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (n *navbar) Content(children ...any) View {
	n.View.Content(children...)
	return n
}

func (n *navbar) Append(children ...any) View {
	n.View.Append(children...)
	return n
}

func (n *navbar) Caption(content ...any) ViewWithCaption {
	caption := NewView(ViewNavbarBrand, "A", WithClass("navbar-brand"), WithAttr("href", "#")).Content(content...)
	n.brand.ReplaceWith(caption.Root())
	n.brand = caption.Root()
	return n
}

func (n *navbar) Toggle(id string, content ...any) ViewWithCaption {
	toggle := Tag("BUTTON",
		WithClass("navbar-toggler"),
		WithAttr("type", "button"),
		WithAttr("data-bs-toggle", "collapse"),
		WithAttr("data-bs-target", "#"+id),
		WithAttr("aria-controls", id),
		WithAttr("aria-expanded", "false"),
		WithAttr("aria-label", "Toggle navigation"),
	).Append(
		Tag("SPAN", WithClass("navbar-toggler-icon")),
	)
	n.toggle.ReplaceWith(toggle.Element)
	n.toggle = toggle
	return n
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (n *navbar) shouldToggle() bool {
	return !n.Root().ClassList().Contains("navbar-expand")
}
