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

type nav struct {
	component
}

type navItem struct {
	component
}

type navbar struct {
	component
	container Element
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

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

	// Apply options
	if opts, err := NewOpts(NavComponent, WithClass("navbar")); err != nil {
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
			name: NavComponent,
			root: root,
			body: container, // Content goes into the container-fluid
		},
		container: container,
	}
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithTabs applies the nav-tabs style (tabbed interface).
func WithTabs() Opt {
	return WithClass("nav-tabs")
}

// WithPills applies the nav-pills style (pill-shaped items).
func WithPills() Opt {
	return WithClass("nav-pills")
}

// WithUnderline applies the nav-underline style (underlined items).
func WithUnderline() Opt {
	return WithClass("nav-underline")
}

///////////////////////////////////////////////////////////////////////////////
// NAVBAR METHODS

// Brand adds a navbar brand element (navbar-brand).
// Creates <a class="navbar-brand" href="#">text</a>
// Returns the navbar for method chaining.
func (n *navbar) Brand(text string, href string) *navbar {
	brand := dom.GetWindow().Document().CreateElement("A")
	brand.SetAttribute("class", "navbar-brand")
	brand.SetAttribute("href", href)
	brand.AppendChild(dom.GetWindow().Document().CreateTextNode(text))

	n.container.AppendChild(brand)
	return n
}

// BrandWithContent adds a navbar brand with custom content (e.g., logo + text).
// Creates <a class="navbar-brand" href="#"> and appends children.
// Returns the navbar for method chaining.
func (n *navbar) BrandWithContent(href string, children ...any) *navbar {
	brand := dom.GetWindow().Document().CreateElement("A")
	brand.SetAttribute("class", "navbar-brand")
	brand.SetAttribute("href", href)

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

	n.container.AppendChild(brand)
	return n
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

	n.container.AppendChild(button)
	return n
}

// Collapse creates a collapsible container for navbar content.
// Creates <div class="collapse navbar-collapse" id="targetId">
// Content should be added to this collapse element.
// Returns the collapse element for adding nav items, forms, etc.
func (n *navbar) Collapse(targetId string) Element {
	collapse := dom.GetWindow().Document().CreateElement("DIV")
	collapse.SetAttribute("class", "collapse navbar-collapse")
	collapse.SetAttribute("id", targetId)

	n.container.AppendChild(collapse)
	return collapse
}

// NavContent creates a navbar-nav container for navigation links.
// Creates <ul class="navbar-nav"> or <div class="navbar-nav"> based on useList.
// Returns the nav container element for adding items.
func (n *navbar) NavContent(useList bool, classes ...string) Element {
	var navEl Element
	if useList {
		navEl = dom.GetWindow().Document().CreateElement("UL")
	} else {
		navEl = dom.GetWindow().Document().CreateElement("DIV")
	}

	classList := append([]string{"navbar-nav"}, classes...)
	navEl.SetAttribute("class", strings.Join(classList, " "))

	return navEl
}

// NavItem creates a navigation link item for use in navbars or navs.
// Creates <a class="nav-link"> element with proper Bootstrap styling.
// Can be used inside navbar-nav containers or regular nav components.
// Accepts any combination of text strings, components, and nodes as children.
//
// Examples:
//
//	NavItem("#", true, false, "Home")                 // Simple active link
//	NavItem("/about", false, false, "About")          // Regular link
//	NavItem("#", false, false, icon, "Dashboard")     // Link with icon
//	NavItem("#", true, false, "Profile ", badge)      // Active link with badge
//	NavItem("#", false, true, "Disabled")             // Disabled link
func NavItem(href string, active bool, disabled bool, children ...any) *navItem {
	// Create link element
	link := dom.GetWindow().Document().CreateElement("A")

	// Build classes
	classes := []string{"nav-link"}

	if active {
		classes = append(classes, "active")
		link.SetAttribute("aria-current", "page")
	}

	if disabled {
		classes = append(classes, "disabled")
		link.SetAttribute("aria-disabled", "true")
	}

	// Set attributes
	link.SetAttribute("class", strings.Join(classes, " "))
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

	return &navItem{
		component: component{
			name: NavComponent,
			root: link,
			body: link,
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVDROPDOWN

// NavDropdown creates a navigation dropdown item with a toggleable menu.
// Creates a dropdown structure with proper Bootstrap classes and attributes.
//
// Parameters:
//   - text: The dropdown toggle button text
//   - active: Whether the dropdown is in active state
//   - children: NavDropdownItem components for the dropdown menu
//
// Example usage:
//
//	nav := Nav().
//	    Append(NavItem("#", false, false, "Home")).
//	    Append(NavDropdown("More", false,
//	        NavDropdownItem("#", false, "Action"),
//	        NavDropdownItem("#", false, "Another action"),
//	        NavDropdownDivider(),
//	        NavDropdownItem("#", false, "Something else"),
//	    )).
//	    Append(NavItem("#", false, false, "Contact"))
func NavDropdown(text string, active bool, children ...Component) *navItem {
	// Create dropdown container (li for navbar-nav or div for regular nav)
	container := dom.GetWindow().Document().CreateElement("DIV")
	container.SetAttribute("class", "nav-item dropdown")

	// Create dropdown toggle link
	toggle := dom.GetWindow().Document().CreateElement("A")
	toggleClasses := []string{"nav-link", "dropdown-toggle"}

	if active {
		toggleClasses = append(toggleClasses, "active")
		toggle.SetAttribute("aria-current", "page")
	}

	toggle.SetAttribute("class", strings.Join(toggleClasses, " "))
	toggle.SetAttribute("href", "#")
	toggle.SetAttribute("role", "button")
	toggle.SetAttribute("data-bs-toggle", "dropdown")
	toggle.SetAttribute("aria-expanded", "false")
	toggle.AppendChild(dom.GetWindow().Document().CreateTextNode(text))

	// Create dropdown menu
	menu := dom.GetWindow().Document().CreateElement("UL")
	menu.SetAttribute("class", "dropdown-menu")

	// Add children to dropdown menu
	for _, child := range children {
		menu.AppendChild(child.Element())
	}

	container.AppendChild(toggle)
	container.AppendChild(menu)

	return &navItem{
		component: component{
			name: NavComponent,
			root: container,
			body: menu, // New dropdown items get added to the menu
		},
	}
}

// NavDropdownItem creates a dropdown menu item.
// Creates <li><a class="dropdown-item"> structure.
//
// Parameters:
//   - href: The link URL
//   - active: Whether the item is active
//   - children: Text or other content for the item
//
// Example: NavDropdownItem("/profile", false, "My Profile")
func NavDropdownItem(href string, active bool, children ...any) Component {
	// Create list item
	li := dom.GetWindow().Document().CreateElement("LI")

	// Create dropdown item link
	link := dom.GetWindow().Document().CreateElement("A")
	itemClasses := []string{"dropdown-item"}

	if active {
		itemClasses = append(itemClasses, "active")
		link.SetAttribute("aria-current", "true")
	}

	link.SetAttribute("class", strings.Join(itemClasses, " "))
	link.SetAttribute("href", href)

	// Append children to link
	for _, child := range children {
		if component, ok := child.(Component); ok {
			link.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			link.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			link.AppendChild(node)
		}
	}

	li.AppendChild(link)

	return &component{
		name: NavComponent,
		root: li,
		body: li,
	}
}

// NavDropdownDivider creates a divider in the dropdown menu.
// Creates <li><hr class="dropdown-divider"></li>
func NavDropdownDivider() Component {
	li := dom.GetWindow().Document().CreateElement("LI")
	hr := dom.GetWindow().Document().CreateElement("HR")
	hr.SetAttribute("class", "dropdown-divider")
	li.AppendChild(hr)

	return &component{
		name: NavComponent,
		root: li,
		body: li,
	}
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

///////////////////////////////////////////////////////////////////////////////
// NAVSPACER

// NavSpacer creates a flexible spacer for navigation components that pushes
// subsequent nav items to the opposite end. Commonly used in navbars to
// separate left-aligned items from right-aligned items.
//
// Uses Bootstrap's flex utilities (me-auto or ms-auto) to create flexible spacing.
//
// Parameters:
//   - opt: Optional styling options (WithClass, etc.)
//
// Example usage:
//
//	nav := Nav().
//	    Append(NavItem("#", false, false, "Home")).
//	    Append(NavItem("#", false, false, "About")).
//	    Append(NavSpacer()).
//	    Append(NavItem("#", false, false, "Login"))
func NavSpacer(opt ...Opt) Component {
	// Create a div element with flexible margin
	spacer := dom.GetWindow().Document().CreateElement("DIV")

	// Apply options with default margin-end auto class
	if opts, err := NewOpts(NavComponent, WithClass("me-auto")); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			spacer.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			spacer.SetAttribute(key, value)
		}
	}

	return &component{
		name: NavComponent,
		root: spacer,
		body: spacer,
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVDIVIDER

// NavDivider creates a visual separator for navigation components.
// Creates a vertical line (pipe) or horizontal rule depending on nav orientation.
//
// For horizontal navs, creates a vertical divider using Bootstrap's vr class.
// For vertical navs, creates a horizontal divider using hr element.
//
// Parameters:
//   - vertical: true for vertical divider (default), false for horizontal divider
//   - opt: Optional styling options (WithClass, etc.)
//
// Example usage:
//
//	nav := Nav().
//	    Append(NavItem("#", false, false, "Home")).
//	    Append(NavDivider(true)).  // Vertical divider
//	    Append(NavItem("#", false, false, "About"))
//
//	verticalNav := Nav(WithClass("flex-column")).
//	    Append(NavItem("#", false, false, "Home")).
//	    Append(NavDivider(false)). // Horizontal divider
//	    Append(NavItem("#", false, false, "About"))
func NavDivider(vertical bool, opt ...Opt) Component {
	var defaultClass string
	var divider Element

	if vertical {
		// Vertical divider - use div with Bootstrap's vr class
		divider = dom.GetWindow().Document().CreateElement("DIV")
		defaultClass = "vr"
	} else {
		// Horizontal divider - use hr element
		divider = dom.GetWindow().Document().CreateElement("HR")
		defaultClass = "my-2"
	}

	// Apply options with appropriate default class
	if opts, err := NewOpts(NavComponent, WithClass(defaultClass)); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			divider.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			divider.SetAttribute(key, value)
		}
	}

	return &component{
		name: NavComponent,
		root: divider,
		body: divider,
	}
}
