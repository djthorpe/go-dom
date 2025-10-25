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

type navDropdown struct {
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
func Nav(opt ...Opt) *nav {
	c := newComponent(NavComponent, dom.GetWindow().Document().CreateElement("NAV"))

	if err := c.applyTo(c.root, append([]Opt{WithClass("nav")}, opt...)...); err != nil {
		panic(err)
	}

	c.body = c.root // For nav, body is the same as root
	return &nav{
		component: *c,
	}
}

// NavBar creates a new bootstrap navbar component (<nav class="navbar">).
// Unlike Nav, NavBar is a more complex component with built-in container support,
// brand, toggler, and collapse functionality for responsive navigation.
// The navbar automatically creates a toggler button and collapsible body.
func NavBar(opt ...Opt) *navbar {
	c := newComponent(NavBarComponent, dom.GetWindow().Document().CreateElement("NAV"))

	// Create container-fluid div
	container := dom.GetWindow().Document().CreateElement("DIV")
	container.SetAttribute("class", "container-fluid")
	c.root.AppendChild(container)

	// Apply options
	if err := c.applyTo(c.root, append([]Opt{WithClass("navbar")}, opt...)...); err != nil {
		panic(err)
	}

	// Create toggler button using Button component with navbar-toggler class
	togglerIcon := dom.GetWindow().Document().CreateElement("SPAN")
	togglerIcon.SetAttribute("class", "navbar-toggler-icon")

	toggler := Button(Color(""),
		WithClass("navbar-toggler"),
		WithAttribute("data-bs-toggle", "collapse"),
		WithAttribute("data-bs-target", "#navbarNav"),
		WithAttribute("aria-controls", "navbarNav"),
		WithAttribute("aria-expanded", "false"),
		WithAttribute("aria-label", "Toggle navigation"),
	).Append(togglerIcon)

	container.AppendChild(toggler.Element())

	// Create collapsible body
	collapse := dom.GetWindow().Document().CreateElement("DIV")
	collapse.SetAttribute("class", "collapse navbar-collapse")
	collapse.SetAttribute("id", "navbarNav")

	// Create navbar-nav list inside collapse
	navList := dom.GetWindow().Document().CreateElement("UL")
	navList.SetAttribute("class", "navbar-nav")
	collapse.AppendChild(navList)

	container.AppendChild(collapse)

	c.body = navList // Content goes into the navbar-nav list
	return &navbar{
		component: *c,
		container: container,
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVBAR METHODS

// Header adds a navbar brand element at the top of the navbar.
// Creates <a class="navbar-brand" href="#"> and appends children.
// Accepts any combination of text strings, components, and nodes as children.
// Returns the navbar for method chaining.
//
// Examples:
//
//	navbar.Header("My Brand")                          // Simple text brand
//	navbar.Header(Icon("house-fill"), " Home")         // Brand with icon
//	navbar.Header(Image("logo.png"), " Company Name")  // Brand with logo
func (n *navbar) Header(children ...any) *navbar {
	brand := dom.GetWindow().Document().CreateElement("A")
	brand.SetAttribute("class", "navbar-brand")
	brand.SetAttribute("href", "#")

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

	// Insert at the top of the container
	if n.container.FirstChild() != nil {
		n.container.InsertBefore(brand, n.container.FirstChild())
	} else {
		n.container.AppendChild(brand)
	}

	return n
}

// Append adds navigation items to the navbar body.
// Only accepts NavItem, NavSpacer, and NavDivider components.
// Returns Component for interface compatibility.
func (n *navbar) Append(children ...any) Component {
	for _, child := range children {
		switch c := child.(type) {
		case *navItem:
			n.body.AppendChild(c.Element())
		case *navDropdown:
			n.body.AppendChild(c.Element())
		case Component:
			// Check if it's a NavSpacer or NavDivider by checking component name
			elem := c.Element()
			compName := elem.GetAttribute("data-component")
			if compName == string(NavSpacerComponent) || compName == string(NavDividerComponent) {
				n.body.AppendChild(elem)
			} else {
				panic("navbar.Append only accepts NavItem, NavDropdown, NavSpacer, and NavDivider components")
			}
		default:
			panic("navbar.Append only accepts NavItem, NavDropdown, NavSpacer, and NavDivider components")
		}
	}
	return n
}

// Insert adds navigation items to the top of the navbar body.
// Only accepts NavItem, NavDropdown, NavSpacer, and NavDivider components.
// Returns Component for interface compatibility.
func (n *navbar) Insert(children ...any) Component {
	for _, child := range children {
		switch c := child.(type) {
		case *navItem:
			if n.body.FirstChild() != nil {
				n.body.InsertBefore(c.Element(), n.body.FirstChild())
			} else {
				n.body.AppendChild(c.Element())
			}
		case *navDropdown:
			if n.body.FirstChild() != nil {
				n.body.InsertBefore(c.Element(), n.body.FirstChild())
			} else {
				n.body.AppendChild(c.Element())
			}
		case Component:
			// Check if it's a NavSpacer or NavDivider by checking component name
			elem := c.Element()
			compName := elem.GetAttribute("data-component")
			if compName == string(NavSpacerComponent) || compName == string(NavDividerComponent) {
				if n.body.FirstChild() != nil {
					n.body.InsertBefore(elem, n.body.FirstChild())
				} else {
					n.body.AppendChild(elem)
				}
			} else {
				panic("navbar.Insert only accepts NavItem, NavDropdown, NavSpacer, and NavDivider components")
			}
		default:
			panic("navbar.Insert only accepts NavItem, NavDropdown, NavSpacer, and NavDivider components")
		}
	}
	return n
}

// NavItem creates a new bootstrap navigation link item (<a> element with class="nav-link").
// The nav item starts in a non-active, non-disabled state. Use SetActive() and SetDisabled()
// methods to change the state after creation.
//
// Parameters:
//   - href: Link destination URL
//   - children: Content for the link (strings, Components, Nodes)
//
// Example usage:
//
//	NavItem("#", "Home").SetActive(true)                    // Active link
//	NavItem("#", "About")                                    // Regular link
//	NavItem("#", "Disabled").SetDisabled(true)              // Disabled link
//	NavItem("#", icon, "Dashboard")                          // Link with icon
//	NavItem("#", "Profile ", badge)                          // Link with badge
//
// NavItem creates a navigation link item.
// When used in a regular Nav or NavBar, creates: <a class="nav-link" href="...">
// When used in a NavDropdown, creates: <li><a class="dropdown-item" href="..."></a></li>
// The context is automatically detected when the item is appended to a dropdown.
//
// Parameters:
//   - href: The link URL
//   - children: Text, Components, or Nodes to append to the link
//
// Example:
//
//	nav := Nav().
//	    Append(NavItem("#", "Home")).
//	    Append(NavItem("#", "About"))
func NavItem(href string, children ...any) *navItem {
	return NavItemWithOpts(href, children...)
}

// NavItemWithOpts creates a navigation link item with options.
// Separated from NavItem to maintain backwards compatibility while adding options support.
func NavItemWithOpts(href string, children ...any) *navItem {
	// Separate options from content
	var opts []Opt
	var content []any

	for _, child := range children {
		if opt, ok := child.(Opt); ok {
			opts = append(opts, opt)
		} else {
			content = append(content, child)
		}
	}

	// Create link element
	link := dom.GetWindow().Document().CreateElement("A")
	link.SetAttribute("href", href)

	// Create component with nav-link class by default
	// This will be adjusted to dropdown-item when appended to a dropdown
	c := newComponent(NavItemComponent, link)

	// Apply options with nav-link class prepended
	if err := c.applyTo(link, append([]Opt{WithClass("nav-link")}, opts...)...); err != nil {
		panic(err)
	}

	// Append children
	for _, child := range content {
		if component, ok := child.(Component); ok {
			link.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			link.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			link.AppendChild(node)
		}
	}

	c.body = link
	return &navItem{
		component: *c,
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVITEM METHODS

// Active gets or sets the active state of the nav item.
// When called with no arguments, returns the current active state.
// When called with a bool argument, sets the active state and returns the new value.
func (n *navItem) Active(active ...bool) bool {
	if len(active) == 0 {
		// Getter: return current state
		link := n.getLinkElement()
		class := link.GetAttribute("class")
		return strings.Contains(class, "active")
	}
	// Setter: update state and return it
	n.updateClasses(active[0], n.Disabled())
	return active[0]
}

// Disabled gets or sets the disabled state of the nav item.
// When called with no arguments, returns the current disabled state.
// When called with a bool argument, sets the disabled state and returns the new value.
func (n *navItem) Disabled(disabled ...bool) bool {
	if len(disabled) == 0 {
		// Getter: return current state
		link := n.getLinkElement()
		class := link.GetAttribute("class")
		return strings.Contains(class, "disabled")
	}
	// Setter: update state and return it
	n.updateClasses(n.Active(), disabled[0])
	return disabled[0]
}

// getLinkElement returns the <a> element for this nav item
func (n *navItem) getLinkElement() Element {
	link := n.root

	// Check if this is a dropdown (root is div.dropdown)
	if strings.Contains(link.GetAttribute("class"), "dropdown") {
		// For dropdowns, find the first child <a> element
		for child := link.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "A" {
				return elem
			}
		}
	} else if link.TagName() == "LI" {
		// For dropdown items, find the <a> child
		for child := link.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "A" {
				return elem
			}
		}
	}

	return link
}

// updateClasses updates the nav-link or dropdown-item classes based on active and disabled state
func (n *navItem) updateClasses(active bool, disabled bool) {
	// Get the root element - for simple nav items it's the <a> tag,
	// for dropdowns it's the container div, for dropdown items it's the <li>
	link := n.root

	// Check if this is a dropdown (root is div.dropdown)
	if strings.Contains(link.GetAttribute("class"), "dropdown") {
		// For dropdowns, find the first child <a> element
		for child := link.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "A" {
				link = elem
				break
			}
		}
	} else if link.TagName() == "LI" {
		// For dropdown items, find the <a> child
		for child := link.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "A" {
				link = elem
				break
			}
		}
	}

	// Determine if this is a nav-link or dropdown-item
	isDropdownItem := strings.Contains(link.GetAttribute("class"), "dropdown-item")

	// Build class list
	var classes []string
	if isDropdownItem {
		classes = []string{"dropdown-item"}
	} else {
		classes = []string{"nav-link"}
		// For nav dropdowns, preserve dropdown-toggle
		if strings.Contains(link.GetAttribute("class"), "dropdown-toggle") {
			classes = append(classes, "dropdown-toggle")
		}
	}

	if active {
		classes = append(classes, "active")
		if isDropdownItem {
			link.SetAttribute("aria-current", "true")
		} else {
			link.SetAttribute("aria-current", "page")
		}
	} else {
		link.RemoveAttribute("aria-current")
	}

	if disabled {
		classes = append(classes, "disabled")
		link.SetAttribute("aria-disabled", "true")
	} else {
		link.RemoveAttribute("aria-disabled")
	}

	link.SetAttribute("class", strings.Join(classes, " "))
}

///////////////////////////////////////////////////////////////////////////////
// NAVDROPDOWN

// NavDropdown creates a navigation dropdown item with a toggleable menu.
// Creates a dropdown structure with proper Bootstrap classes and attributes.
// The dropdown starts in a non-active state. Use Active() to change it.
//
// Parameters:
//   - text: The dropdown toggle button text
//   - children: NavItem components for the dropdown menu (will be converted to dropdown items)
//
// Example usage:
//
//	nav := Nav().
//	    Append(NavItem("#", "Home")).
//	    Append(NavDropdown("More",
//	        NavItem("#", "Action"),
//	        NavItem("#", "Another action"),
//	        NavDropdownDivider(),
//	        NavItem("#", "Something else"),
//	    )).
//	    Append(NavItem("#", "Contact"))
func NavDropdown(text string, children ...Component) *navDropdown {
	// Create dropdown container (li for navbar-nav or div for regular nav)
	container := dom.GetWindow().Document().CreateElement("DIV")
	container.SetAttribute("class", "nav-item dropdown")

	// Create dropdown toggle link
	toggle := dom.GetWindow().Document().CreateElement("A")
	toggle.SetAttribute("class", "nav-link dropdown-toggle")
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
		elem := child.Element()

		// Convert NavItem to dropdown format
		if elem.TagName() == "A" && elem.ClassList().Contains("nav-link") {
			// This is a NavItem - wrap it in <li> and change class to dropdown-item
			li := dom.GetWindow().Document().CreateElement("LI")

			// Change class from nav-link to dropdown-item
			elem.ClassList().Remove("nav-link")
			elem.ClassList().Add("dropdown-item")

			li.AppendChild(elem)
			menu.AppendChild(li)
		} else {
			// Already in dropdown format or is a divider/header
			menu.AppendChild(elem)
		}
	}

	container.AppendChild(toggle)
	container.AppendChild(menu)

	c := newComponent(NavDropdownComponent, container)
	c.body = menu // New dropdown items get added to the menu
	return &navDropdown{
		component: *c,
	}
}

// Active gets or sets the active state of NavItems within a NavDropdown.
// When called with no arguments or -1, returns array of active item indices.
// When called with an index, sets that item as active and returns all active indices.
func (n *navDropdown) Active(index ...int) []int {
	// Find the menu element (UL)
	var menu Element
	for child := n.root.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "UL" {
			menu = elem
			break
		}
	}
	if menu == nil {
		return []int{}
	}

	// If setting active state
	if len(index) > 0 && index[0] >= 0 {
		idx := 0
		for child := menu.FirstChild(); child != nil; child = child.NextSibling() {
			if li, ok := child.(Element); ok && li.TagName() == "LI" {
				// Find the <a> tag within the <li>
				for linkChild := li.FirstChild(); linkChild != nil; linkChild = linkChild.NextSibling() {
					if link, ok := linkChild.(Element); ok && link.TagName() == "A" {
						if idx == index[0] {
							// Set this item as active
							link.ClassList().Add("active")
							link.SetAttribute("aria-current", "page")
						}
						break
					}
				}
				idx++
			}
		}
	}

	// Return array of active indices
	activeIndices := []int{}
	idx := 0
	for child := menu.FirstChild(); child != nil; child = child.NextSibling() {
		if li, ok := child.(Element); ok && li.TagName() == "LI" {
			// Find the <a> tag within the <li>
			for linkChild := li.FirstChild(); linkChild != nil; linkChild = linkChild.NextSibling() {
				if link, ok := linkChild.(Element); ok && link.TagName() == "A" {
					if link.ClassList().Contains("active") {
						activeIndices = append(activeIndices, idx)
					}
					break
				}
			}
			idx++
		}
	}
	return activeIndices
}

// Disabled gets or sets the disabled state of NavItems within a NavDropdown.
// When called with no arguments or -1, returns array of disabled item indices.
// When called with an index, sets that item as disabled and returns all disabled indices.
func (n *navDropdown) Disabled(index ...int) []int {
	// Find the menu element (UL)
	var menu Element
	for child := n.root.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "UL" {
			menu = elem
			break
		}
	}
	if menu == nil {
		return []int{}
	}

	// If setting disabled state
	if len(index) > 0 && index[0] >= 0 {
		idx := 0
		for child := menu.FirstChild(); child != nil; child = child.NextSibling() {
			if li, ok := child.(Element); ok && li.TagName() == "LI" {
				// Find the <a> tag within the <li>
				for linkChild := li.FirstChild(); linkChild != nil; linkChild = linkChild.NextSibling() {
					if link, ok := linkChild.(Element); ok && link.TagName() == "A" {
						if idx == index[0] {
							// Set this item as disabled
							link.ClassList().Add("disabled")
							link.SetAttribute("aria-disabled", "true")
							link.SetAttribute("tabindex", "-1")
						}
						break
					}
				}
				idx++
			}
		}
	}

	// Return array of disabled indices
	disabledIndices := []int{}
	idx := 0
	for child := menu.FirstChild(); child != nil; child = child.NextSibling() {
		if li, ok := child.(Element); ok && li.TagName() == "LI" {
			// Find the <a> tag within the <li>
			for linkChild := li.FirstChild(); linkChild != nil; linkChild = linkChild.NextSibling() {
				if link, ok := linkChild.(Element); ok && link.TagName() == "A" {
					if link.ClassList().Contains("disabled") {
						disabledIndices = append(disabledIndices, idx)
					}
					break
				}
			}
			idx++
		}
	}
	return disabledIndices
}

// NavDropdownDivider creates a divider in the dropdown menu.
// Creates <li><hr class="dropdown-divider"></li>
func NavDropdownDivider(opt ...Opt) Component {
	li := dom.GetWindow().Document().CreateElement("LI")
	hr := dom.GetWindow().Document().CreateElement("HR")

	c := newComponent(NavItemComponent, li)

	// Apply options to the hr element with default class
	if err := c.applyTo(hr, append([]Opt{WithClass("dropdown-divider")}, opt...)...); err != nil {
		panic(err)
	}

	li.AppendChild(hr)
	c.body = li
	return c
}

// NavDropdownHeader creates a header in the dropdown menu.
// Creates <li><h6 class="dropdown-header">text</h6></li>
func NavDropdownHeader(text string, opt ...Opt) Component {
	li := dom.GetWindow().Document().CreateElement("LI")
	h6 := dom.GetWindow().Document().CreateElement("H6")
	h6.AppendChild(dom.GetWindow().Document().CreateTextNode(text))

	c := newComponent(NavItemComponent, li)

	// Apply options to the h6 element with default class
	if err := c.applyTo(h6, append([]Opt{WithClass("dropdown-header")}, opt...)...); err != nil {
		panic(err)
	}

	li.AppendChild(h6)
	c.body = li
	return c
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
	c := newComponent(NavSpacerComponent, dom.GetWindow().Document().CreateElement("DIV"))

	if err := c.applyTo(c.root, append([]Opt{WithClass("me-auto")}, opt...)...); err != nil {
		panic(err)
	}

	c.body = c.root
	return c
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
	var tag string

	if vertical {
		// Vertical divider - use div with Bootstrap's vr class
		tag = "DIV"
		defaultClass = "vr"
	} else {
		// Horizontal divider - use hr element
		tag = "HR"
		defaultClass = "my-2"
	}

	c := newComponent(NavDividerComponent, dom.GetWindow().Document().CreateElement(tag))

	if err := c.applyTo(c.root, append([]Opt{WithClass(defaultClass)}, opt...)...); err != nil {
		panic(err)
	}

	c.body = c.root
	return c
}
