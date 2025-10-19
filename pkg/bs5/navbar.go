package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type NavBrand struct {
	dom.Element
}

type NavItem struct {
	dom.Element
	link dom.Element
}

type NavSpacer struct {
	dom.Element
}

type Nav struct {
	dom.Element
	container dom.Element
	brand     dom.Node
	collapse  dom.Element
	nav       dom.Element
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NavBrand creates a navbar brand link
func (app *App) NavBrand(href string, children ...dom.Node) *NavBrand {
	// Create brand link
	brandLink := app.CreateElement("a")
	brandLink.AddClass("navbar-brand")
	brandLink.AddClass("me-auto") // Keep brand on the left when toggler is visible
	brandLink.SetAttribute("href", href)
	for _, child := range children {
		brandLink.AppendChild(child)
	}

	return &NavBrand{
		Element: brandLink,
	}
}

// NavItem creates a navbar navigation item
func (app *App) NavItem(label string, href string, active bool) *NavItem {
	// Create list item
	li := app.CreateElement("li")
	li.AddClass("nav-item")

	// Create link
	a := app.CreateElement("a")
	a.AddClass("nav-link")
	if active {
		a.AddClass("active")
		a.SetAttribute("aria-current", "page")
	}
	a.SetAttribute("href", href)
	a.AppendChild(app.CreateTextNode(label))

	li.AppendChild(a)

	return &NavItem{
		Element: li,
		link:    a,
	}
}

// NavSpacer creates a spacer element that pushes following nav items to the right
func (app *App) NavSpacer() *NavSpacer {
	// Create an empty li with flex-grow to consume space
	spacer := app.CreateElement("li")
	spacer.SetAttribute("style", "flex-grow: 1;")

	return &NavSpacer{
		Element: spacer,
	}
}

// Nav creates a Bootstrap 5 navbar component
func (app *App) Nav(brand dom.Node, items ...dom.Node) *Nav {
	// Create main nav element
	nav := app.CreateElement("nav")
	nav.AddClass("navbar")
	nav.AddClass("navbar-expand-lg")
	nav.AddClass("bg-body-tertiary")

	// Create container-fluid div
	container := app.CreateElement("div")
	container.AddClass("container-fluid")

	// Create toggler button for mobile (add first for left positioning)
	toggler := app.CreateElement("button")
	toggler.AddClass("navbar-toggler")
	toggler.SetAttribute("type", "button")
	toggler.SetAttribute("data-bs-toggle", "collapse")
	toggler.SetAttribute("data-bs-target", "#navbarNav")
	toggler.SetAttribute("aria-controls", "navbarNav")
	toggler.SetAttribute("aria-expanded", "false")
	toggler.SetAttribute("aria-label", "Toggle navigation")

	// Create toggler icon
	togglerIcon := app.CreateElement("span")
	togglerIcon.AddClass("navbar-toggler-icon")
	toggler.AppendChild(togglerIcon)

	// Add toggler first, then brand (for left-side burger)
	container.AppendChild(toggler)
	container.AppendChild(brand)

	// Create collapsible div
	collapse := app.CreateElement("div")
	collapse.AddClass("collapse")
	collapse.AddClass("navbar-collapse")
	collapse.SetAttribute("id", "navbarNav")

	// Create nav list
	navList := app.CreateElement("ul")
	navList.AddClass("navbar-nav")
	navList.AddClass("w-100") // Full width to enable flex spacing

	// Add nav items
	for _, item := range items {
		navList.AppendChild(item)
	}

	collapse.AppendChild(navList)

	// Assemble the navbar (toggler and brand already added above)
	container.AppendChild(collapse)
	nav.AppendChild(container)

	return &Nav{
		Element:   nav,
		container: container,
		brand:     brand,
		collapse:  collapse,
		nav:       navList,
	}
}

////////////////////////////////////////////////////////////////////////
// METHODS

// AddClass adds a CSS class to the brand
func (b *NavBrand) AddClass(className string) *NavBrand {
	b.Element.AddClass(className)
	return b
}

// RemoveClass removes a CSS class from the brand
func (b *NavBrand) RemoveClass(className string) *NavBrand {
	b.Element.RemoveClass(className)
	return b
}

// AddClass adds a CSS class to the nav item
func (i *NavItem) AddClass(className string) *NavItem {
	i.Element.AddClass(className)
	return i
}

// RemoveClass removes a CSS class from the nav item
func (i *NavItem) RemoveClass(className string) *NavItem {
	i.Element.RemoveClass(className)
	return i
}

// SetActive sets the active state of the nav item
func (i *NavItem) SetActive(active bool) *NavItem {
	if active {
		i.link.AddClass("active")
		i.link.SetAttribute("aria-current", "page")
	} else {
		i.link.RemoveClass("active")
	}
	return i
}

// AddClass adds a CSS class to the spacer
func (s *NavSpacer) AddClass(className string) *NavSpacer {
	s.Element.AddClass(className)
	return s
}

// RemoveClass removes a CSS class from the spacer
func (s *NavSpacer) RemoveClass(className string) *NavSpacer {
	s.Element.RemoveClass(className)
	return s
}

// AddItem adds a navigation item to the navbar (can be used to add items dynamically)
func (n *Nav) AddItem(item dom.Node) *Nav {
	n.nav.AppendChild(item)
	return n
}

// AddClass adds a CSS class to the navbar
func (n *Nav) AddClass(className string) *Nav {
	n.Element.AddClass(className)
	return n
}

// RemoveClass removes a CSS class from the navbar
func (n *Nav) RemoveClass(className string) *Nav {
	n.Element.RemoveClass(className)
	return n
}

// SetColorScheme sets the navbar color scheme
func (n *Nav) SetColorScheme(scheme ColorVariant) *Nav {
	n.Element.SetAttribute("data-bs-theme", string(scheme))
	return n
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (b *NavBrand) String() string {
	return "<bs5-navbrand>"
}

func (i *NavItem) String() string {
	return "<bs5-navitem>"
}

func (s *NavSpacer) String() string {
	return "<bs5-navspacer>"
}

func (n *Nav) String() string {
	return "<bs5-nav>"
}
