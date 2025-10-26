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

var _ ViewWithHeaderFooter = (*navbar)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewNavbar = "mvc-bs-navbar"
)

func init() {
	RegisterView(ViewNavbar, newNavbarFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NavBar(opts ...Opt) *navbar {
	opts = append([]Opt{WithClass("navbar", "bg-primary", "navbar-expand-lg")}, opts...)
	navbar := &navbar{NewView(ViewNavbar, "NAV", opts...)}

	// Within the navbar, add a container
	navbar.Append(Div(WithClass("container-fluid")).Append(
		Div(WithClass("collapse", "navbar-collapse")).Append(
			Div(WithClass("navbar-nav")).Append(
				// TODO
				NewView(ViewNavbar, "UL", WithClass("navbar-nav")).Append(
					NewView(ViewNavbar, "LI", WithClass("nav-item")).Append(
						NewView(ViewNavbar, "A", WithClass("nav-link"), WithAttr("href", "#")).Append("Home"),
					),
				),
			),
		),
	))

	return navbar
}

func newNavbarFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "NAV" {
		panic(fmt.Sprintf("newNavbarFromElement: invalid tag name %q", tagName))
	}
	return &navbar{NewViewWithElement(element)}
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
