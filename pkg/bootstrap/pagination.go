package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type pagination struct {
	component
	nav  Element
	list Element
}

var _ Component = (*pagination)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Pagination creates a new Bootstrap pagination component
// Use Append to add content wrapped in <li class="page-item"> tags
// Use Insert (inherited from Component) to insert content at the top
//
// Example:
//
//	Pagination("Page navigation").
//	    Append("Previous").
//	    Append("1").
//	    Append("2").
//	    Append("Next")
func Pagination(ariaLabel string, opt ...Opt) *pagination {
	doc := dom.GetWindow().Document()

	// Create nav element
	nav := doc.CreateElement("NAV")
	if ariaLabel != "" {
		nav.SetAttribute("aria-label", ariaLabel)
	} else {
		nav.SetAttribute("aria-label", "Page navigation")
	}

	// Create ul element with pagination class
	list := doc.CreateElement("UL")
	list.SetAttribute("class", "pagination")

	nav.AppendChild(list)

	c := newComponent(PaginationComponent, nav)

	// Apply options to nav element
	if err := c.applyTo(nav, opt...); err != nil {
		panic(err)
	}

	return &pagination{
		component: *c,
		nav:       nav,
		list:      list,
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Append wraps content in <li class="page-item"> and appends to the pagination list
// Accepts Component, Element, or string children
func (p *pagination) Append(children ...any) Component {
	doc := p.nav.OwnerDocument()

	for _, child := range children {
		// Create li element
		li := doc.CreateElement("LI")
		li.SetAttribute("class", "page-item")

		// Append child content to li
		if component, ok := child.(Component); ok {
			li.AppendChild(component.Element())
		} else if elem, ok := child.(Element); ok {
			li.AppendChild(elem)
		} else if str, ok := child.(string); ok {
			li.AppendChild(doc.CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			li.AppendChild(node)
		}

		// Append li to list
		p.list.AppendChild(li)
	}

	return p
}
