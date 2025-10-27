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
}

// navbar are elements to create a navigation bar
type navitem struct {
	View
}

var _ ViewWithHeaderFooter = (*navbar)(nil)
var _ View = (*navitem)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewNavbar  = "mvc-bs-navbar"
	ViewNavItem = "mvc-bs-navitem"
)

func init() {
	RegisterView(ViewNavbar, newNavbarFromElement)
	RegisterView(ViewNavItem, newNavItemFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NavBar(opts ...Opt) *navbar {
	opts = append([]Opt{WithClass("navbar", "bg-primary", "navbar-expand-lg")}, opts...)
	navbar := &navbar{NewView(ViewNavbar, "NAV", opts...)}

	// The body element is a UL with class navbar-nav
	body := NewView(ViewNavbar, "UL", WithClass("navbar-nav"))

	// Within the navbar, add a container
	navbar.Append(Div(WithClass("container-fluid")).Append(
		Div(WithClass("collapse", "navbar-collapse")).Append(
			Div(WithClass("navbar-nav")).Append(body),
		),
	))

	// Set the body element
	navbar.Body(body.Root())

	// Return the navbar
	return navbar
}

func NavItem(content ...any) *navitem {
	opts := []Opt{WithClass("nav-item")}
	view := &navitem{NewView(ViewNavItem, "LI", opts...)}

	// Create the link element as the body
	body := Link("#", WithClass("nav-link")).Append(content...)
	view.Content(body)
	view.Body(body.Root())

	return view
}

func newNavbarFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "NAV" {
		panic(fmt.Sprintf("newNavbarFromElement: invalid tag name %q", tagName))
	}
	view := &navbar{NewViewWithElement(element)}
	// TODO: set body element
	return view
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

func (navbar *navbar) Header(children ...any) ViewWithHeaderFooter {
	parent, brand := navbar.brand()
	if brand != nil {
		brand.Remove()
	}

	// Create the brand link
	view := Link("#", WithClass("navbar-brand")).Append(children...)
	if firstChild := parent.FirstChild(); firstChild != nil {
		parent.InsertBefore(view.Root(), firstChild)
	} else {
		parent.AppendChild(view.Root())
	}

	// Return the navbar
	return navbar
}

func (navbar *navbar) Footer(children ...any) ViewWithHeaderFooter {
	panic("Footer: not implemented")
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Return the brand element
func (navbar *navbar) brand() (Element, Element) {
	// parent should be <div class="container-fluid">
	parent := navbar.Root().FirstElementChild()
	if parent == nil || parent.TagName() != "DIV" || !parent.ClassList().Contains("container-fluid") {
		panic("navbar.brand: missing parent element")
	}

	// Try and get the brand element
	element := parent.FirstElementChild()
	if element == nil || element.TagName() != "A" || !element.ClassList().Contains("navbar-brand") {
		return parent, nil
	}

	// Return the parent and brand element
	return parent, element
}
