package bootstrap

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type nav struct {
	component
}

type navbar struct {
	component
}

type navitem struct {
	component
}

type navdropdown struct {
	component
}

type navdivider struct {
	component
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NavBar creates a new bootstrap navbar component (<nav class="navbar">).
// Unlike Nav, NavBar is a more complex component with built-in container support,
// brand, toggler, and collapse functionality for responsive navigation.
//
// Structure: <nav class="navbar"> → <div class="container-fluid">
//
// Expand variants: WithClass("navbar-expand-lg") - expands at lg breakpoint
// Color schemes: WithAttribute("data-bs-theme", "dark") + bg-* utilities
// Placement: WithClass("fixed-top"), WithClass("sticky-top")
//
// Example:
//
//	NavBar(WithClass("navbar-expand-lg"), WithClass("bg-body-tertiary")).
//	    Append(brandElement, navLinks, searchForm)
func NavBar(opt ...Opt) *navbar {
	// Create a nav element with navbar class
	root := dom.GetWindow().Document().CreateElement("NAV")

	// Create container-fluid div
	container := dom.GetWindow().Document().CreateElement("DIV")
	container.SetAttribute("class", "container-fluid")
	root.AppendChild(container)

	// Create a body div (which can be used for collapse content)
	bodydiv := dom.GetWindow().Document().CreateElement("DIV")
	bodydiv.ClassList().Add("collapse", "navbar-collapse")
	container.AppendChild(bodydiv)

	// Create a UL element for nav items
	body := dom.GetWindow().Document().CreateElement("UL")
	body.SetAttribute("class", "navbar-nav")
	bodydiv.AppendChild(body)

	// Apply options
	if opts, err := NewOpts(NavBarComponent, WithClass("navbar", "navbar-expand"), WithAttribute("role", "navigation"), WithTheme(DARK)); err != nil {
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

	return &navbar{
		component: component{
			name: NavBarComponent,
			root: root,
			body: body, // Content goes into the navbar-body (NavItem, NavMenu, NavDivider, NavSpacer)
		},
	}
}

// NavItem creates a navigation link item for use in navbars or navs.
// Creates <ul><a class="nav-link"></a></ul> element with proper Bootstrap styling.
// Can be appended to a navbar or nav component.
//
// Accepts any combination of text strings, components, and nodes as children.
// Examples:
//
//	NavItem("#", "Home")             // Simple active link
//	NavItem("/about","About")        // Regular link
//	NavItem("#",  icon, "Dashboard") // Link with icon
func NavItem(href string, children ...any) *navitem {
	root := dom.GetWindow().Document().CreateElement("LI")
	root.SetAttribute("class", "nav-item")

	// Create link element
	link := dom.GetWindow().Document().CreateElement("A")
	root.AppendChild(link)

	// Set attributes
	link.SetAttribute("class", "nav-link")
	link.SetAttribute("href", href)

	// Append children
	for _, child := range children {
		if component, ok := child.(Component); ok {
			link.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			link.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			link.AppendChild(node)
		}
	}

	return &navitem{
		component: component{
			name: NavItemComponent,
			root: root,
			body: link,
		},
	}
}

// NavDropdown creates a navigation dropdown item for use in navbars or navs.
// Can be appended to a navbar or nav component.
//
// Accepts any combination of text strings, components, and nodes as children.
// Examples:
//
//	NavItem("#", "Home")             // Simple active link
//	NavItem("/about","About")        // Regular link
//	NavItem("#",  icon, "Dashboard") // Link with icon
func NavDropdown(children ...any) *navdropdown {
	root := dom.GetWindow().Document().CreateElement("LI")
	root.ClassList().Add("nav-item", "dropdown")

	// Create link element
	link := dom.GetWindow().Document().CreateElement("A")
	link.ClassList().Add("nav-link", "dropdown-toggle")
	link.SetAttribute("role", "button")
	link.SetAttribute("data-bs-toggle", "dropdown")
	link.SetAttribute("aria-expanded", "false")

	root.AppendChild(link)

	// Set attributes
	link.SetAttribute("href", "#")

	// Create the dropdown list
	dropdown := dom.GetWindow().Document().CreateElement("UL")
	dropdown.ClassList().Add("dropdown-menu")
	root.AppendChild(dropdown)

	// Append children
	for _, child := range children {
		if component, ok := child.(Component); ok {
			link.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			link.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			link.AppendChild(node)
		}
	}

	return &navdropdown{
		component: component{
			name: NavDropdownComponent,
			root: root,
			body: dropdown,
		},
	}
}

// NavDivider creates a divider (horizonal or vertical) for use in nav, nav-bar or nav-dropdown.
func NavDivider() Component {
	li := dom.GetWindow().Document().CreateElement("LI")
	hr := dom.GetWindow().Document().CreateElement("HR")
	hr.SetAttribute("class", "dropdown-divider")
	li.AppendChild(hr)

	return &navdivider{
		component: component{
			name: NavDividerComponent,
			root: li,
			body: li,
		},
	}
}

// Nav creates a new bootstrap navigation component (nav element with class="nav")
// By default creates a <nav> element, but can be customized with options.
// Use Append() with NavItem() to add navigation items.
//
// Variants: WithTabs(), WithPills(), WithUnderline()
// Alignment: WithClass("justify-content-center"), WithClass("justify-content-end")
// Vertical: WithClass("flex-column")
// Fill: WithClass("nav-fill"), WithClass("nav-justified")
//
// Example:
//
//	Nav(WithTabs()).
//	    Append(NavItem("#", true, false, "Home")).
//	    Append(NavItem("#", false, false, "Profile")).
//	    Append(NavItem("#", false, false, "Contact"))
func Nav(opt ...Opt) *nav {
	// Create a nav element
	root := dom.GetWindow().Document().CreateElement("NAV")

	// Apply options
	if opts, err := NewOpts(NavComponent, WithClass("nav")); err != nil {
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

	return &nav{
		component: component{
			name: NavComponent,
			root: root,
			body: root, // For nav, body is the same as root
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - OPTIONS

// WithTabs applies the nav-tabs style (tabbed interface).
func WithTabs() Opt {
	return func(o *opts) error {
		if o.name != NavComponent {
			return ErrBadParameter.Withf("WithTabs can only be used with Nav component, not %q", o.name)
		}
		return WithClass("nav-tabs")(o)
	}
}

// WithPills applies the nav-pills style (pill-shaped items).
func WithPills() Opt {
	return func(o *opts) error {
		if o.name != NavComponent {
			return ErrBadParameter.Withf("WithPills can only be used with Nav component, not %q", o.name)
		}
		return WithClass("nav-pills")(o)
	}
}

// WithUnderline applies the nav-underline style (underlined items).
func WithUnderline() Opt {
	return func(o *opts) error {
		if o.name != NavComponent {
			return ErrBadParameter.Withf("WithUnderline can only be used with Nav component, not %q", o.name)
		}
		return WithClass("nav-underline")(o)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Append appends one or more child components to the navbar body.
func (nav *navbar) Append(children ...any) Component {
	for _, child := range children {
		// We can append NavItem, NavDropdown, NavSpacer, NavDivider to the body
		child, ok := child.(Component)
		if !ok {
			panic(fmt.Errorf("cannot append %T to %s", child, nav.Name()))
		}
		switch name(child.Name()) {
		case NavItemComponent:
			nav.body.AppendChild(child.Element())
		case NavDropdownComponent:
			nav.body.AppendChild(child.Element())
		default:
			panic(fmt.Errorf("cannot append %q to %q", child.Name(), nav.Name()))
		}
	}

	// Return nav for chaining
	return nav
}

// Append appends one or more child components to the navbar body.
func (nav *navdropdown) Append(children ...any) Component {
	for _, child := range children {
		// We can append NavItem, NavDivider to the body
		child, ok := child.(Component)
		if !ok {
			panic(fmt.Errorf("cannot append %T to %s", child, nav.Name()))
		}
		switch child.Name() {
		case string(NavDividerComponent):
			nav.body.AppendChild(child.Element())
		case string(NavItemComponent):
			// Replace nav-link with dropdown-item for dropdown menu items
			elements := child.Element().GetElementsByClassName("nav-link")
			for _, element := range elements {
				element.ClassList().Remove("nav-link")
				element.ClassList().Add("dropdown-item")
			}
			nav.body.AppendChild(child.Element())
		default:
			panic(fmt.Errorf("cannot append %q to %q", child.Name(), nav.Name()))
		}
	}

	// Return nav for chaining
	return nav
}

// Append appends one or more child components to the navdropdown body.
func (nav *navitem) Append(children ...any) Component {
	for _, child := range children {
		// We can append NavItem, NavDropdown, NavSpacer, NavDivider to the body
		child, ok := child.(Component)
		if !ok {
			panic(fmt.Errorf("cannot append %T to %s", child, nav.Name()))
		}
		switch child.Name() {
		case string(NavItemComponent):
			nav.body.AppendChild(child.Element())
		default:
			panic(fmt.Errorf("cannot append %q to %q", child, nav.Name()))
		}
	}

	// Return nav for chaining
	return nav
}

// Disabled adds the disabled class to a navitem, making it appear disabled.
func (navitem *navitem) Disabled() *navitem {
	navitem.root.ClassList().Add("disabled")
	navitem.root.SetAttribute("aria-disabled", "true")
	return navitem
}

// Active adds the active class to a navitem, making it appear active.
func (navitem *navitem) Active() *navitem {
	navitem.root.ClassList().Add("active")
	navitem.root.SetAttribute("aria-current", "page")
	return navitem
}

// Brand adds a navbar brand with custom content (e.g., logo + text).
// Returns the navbar for method chaining.
func (navbar *navbar) Brand(children ...any) *navbar {
	// Container is the first child of the navbar root
	container, ok := navbar.root.FirstChild().(Element)
	if !ok {
		panic("invalid navbar structure")
	}

	// Create brand element
	brand := dom.GetWindow().Document().CreateElement("A")
	brand.SetAttribute("class", "navbar-brand")
	brand.SetAttribute("href", "#")

	// If there is a navbar-brand already, replace it
	existing := container.GetElementsByClassName("navbar-brand")
	if len(existing) == 1 {
		// Replace existing brand element
		container.ReplaceChild(existing[0], brand)
	} else {
		// Set as the first node of the container
		container.InsertBefore(brand, container.FirstChild())
	}

	// Append children
	for _, child := range children {
		if component, ok := child.(Component); ok {
			brand.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			brand.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			brand.AppendChild(node)
		}
	}

	return navbar
}

// Toggler adds a navbar toggler button for responsive collapse.
// Creates the toggler button with the hamburger icon.
// The targetId should match the id of the collapse element.
// Returns the navbar for method chaining.
func (n *navbar) Toggler(targetId string) *navbar {
	button := dom.GetWindow().Document().CreateElement("BUTTON")
	button.SetAttribute("class", "navbar-toggler")
	button.SetAttribute("type", "button")
	button.SetAttribute("data-bs-toggle", "collapse")
	button.SetAttribute("data-bs-target", "#"+targetId)
	button.SetAttribute("aria-controls", targetId)
	button.SetAttribute("aria-expanded", "false")
	button.SetAttribute("aria-label", "Toggle navigation")

	// Add toggler icon
	icon := dom.GetWindow().Document().CreateElement("SPAN")
	icon.SetAttribute("class", "navbar-toggler-icon")
	button.AppendChild(icon)

	n.root.AppendChild(button)
	return n
}

// NavDropdownHeader creates a header in the dropdown menu.
// Creates <li><h6 class="dropdown-header">text</h6></li>
func NavDropdownHeader(text string) Component {
	li := dom.GetWindow().Document().CreateElement("LI")
	h6 := dom.GetWindow().Document().CreateElement("H6")
	h6.SetAttribute("class", "dropdown-header")
	h6.AppendChild(dom.GetWindow().Document().CreateTextNode(text))
	li.AppendChild(h6)

	return &component{
		name: NavComponent,
		root: li,
		body: li,
	}
}
