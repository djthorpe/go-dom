package bootstrap

import (
	"strings"
	"testing"

	dom "github.com/djthorpe/go-wasmbuild"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

func TestNav_Basic(t *testing.T) {
	nav := Nav()

	// Check element type
	if tagName := nav.Element().TagName(); tagName != "NAV" {
		t.Errorf("Expected tag name 'NAV', got '%s'", tagName)
	}

	// Check class
	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav") {
		t.Errorf("Expected class to contain 'nav', got '%s'", class)
	}

	// Check no children initially
	children := nav.Element().ChildNodes()
	if len(children) != 0 {
		t.Errorf("Expected 0 children, got %d", len(children))
	}
}

func TestNav_WithTabs(t *testing.T) {
	nav := Nav(WithTabs())

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-tabs") {
		t.Errorf("Expected class to contain 'nav-tabs', got '%s'", class)
	}
}

func TestNav_WithPills(t *testing.T) {
	nav := Nav(WithPills())

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-pills") {
		t.Errorf("Expected class to contain 'nav-pills', got '%s'", class)
	}
}

func TestNav_WithUnderline(t *testing.T) {
	nav := Nav(WithUnderline())

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-underline") {
		t.Errorf("Expected class to contain 'nav-underline', got '%s'", class)
	}
}

func TestNav_WithClass(t *testing.T) {
	nav := Nav(WithClass("justify-content-center"))

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "justify-content-center") {
		t.Errorf("Expected class to contain 'justify-content-center', got '%s'", class)
	}
}

func TestNav_Vertical(t *testing.T) {
	nav := Nav(WithClass("flex-column"))

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "flex-column") {
		t.Errorf("Expected class to contain 'flex-column', got '%s'", class)
	}
}

func TestNav_Fill(t *testing.T) {
	nav := Nav(WithPills(), WithClass("nav-fill"))

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-fill") {
		t.Errorf("Expected class to contain 'nav-fill', got '%s'", class)
	}
	if !strings.Contains(class, "nav-pills") {
		t.Errorf("Expected class to contain 'nav-pills', got '%s'", class)
	}
}

func TestNav_Component(t *testing.T) {
	nav := Nav()

	// Check component name
	if name := nav.Name(); name != "nav" {
		t.Errorf("Expected component name 'nav', got '%s'", name)
	}

	// Check that Element() returns the same as root
	if elem := nav.Element(); elem != nav.root {
		t.Error("Element() should return the root element")
	}
}

func TestNav_WithAttribute(t *testing.T) {
	nav := Nav(WithAttribute("data-bs-theme", "dark"))

	// Check custom attribute
	theme := nav.Element().GetAttribute("data-bs-theme")
	if theme != "dark" {
		t.Errorf("Expected data-bs-theme 'dark', got '%s'", theme)
	}
}

func TestNav_WithMultipleAttributes(t *testing.T) {
	nav := Nav(
		WithAttribute("data-bs-theme", "dark"),
		WithAttribute("role", "navigation"),
		WithAttribute("id", "main-nav"),
	)

	theme := nav.Element().GetAttribute("data-bs-theme")
	if theme != "dark" {
		t.Errorf("Expected data-bs-theme 'dark', got '%s'", theme)
	}

	role := nav.Element().GetAttribute("role")
	if role != "navigation" {
		t.Errorf("Expected role 'navigation', got '%s'", role)
	}

	id := nav.Element().GetAttribute("id")
	if id != "main-nav" {
		t.Errorf("Expected id 'main-nav', got '%s'", id)
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVBAR TESTS

func TestNavBar_Basic(t *testing.T) {
	navbar := NavBar()

	// Check element type
	if tagName := navbar.Element().TagName(); tagName != "NAV" {
		t.Errorf("Expected tag name 'NAV', got '%s'", tagName)
	}

	// Check class
	class := navbar.Element().GetAttribute("class")
	if !strings.Contains(class, "navbar") {
		t.Errorf("Expected class to contain 'navbar', got '%s'", class)
	}

	// Check container-fluid exists
	children := navbar.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (container-fluid), got %d", len(children))
	}

	container := children[0]
	if container.NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected first child to be an element node")
	}

	containerClass := container.(Element).GetAttribute("class")
	if !strings.Contains(containerClass, "container-fluid") {
		t.Errorf("Expected container class to contain 'container-fluid', got '%s'", containerClass)
	}
}

func TestNavBar_WithExpand(t *testing.T) {
	navbar := NavBar(WithClass("navbar-expand-lg"))

	class := navbar.Element().GetAttribute("class")
	if !strings.Contains(class, "navbar-expand-lg") {
		t.Errorf("Expected class to contain 'navbar-expand-lg', got '%s'", class)
	}
}

func TestNavBar_Brand(t *testing.T) {
	navbar := NavBar().Brand("My Brand", "#")

	// Check that brand was added to container
	container := navbar.Element().ChildNodes()[0].(Element)
	brandElements := container.ChildNodes()

	if len(brandElements) != 1 {
		t.Errorf("Expected 1 brand element, got %d", len(brandElements))
	}

	brand := brandElements[0].(Element)
	if brand.TagName() != "A" {
		t.Errorf("Expected brand tag 'A', got '%s'", brand.TagName())
	}

	brandClass := brand.GetAttribute("class")
	if !strings.Contains(brandClass, "navbar-brand") {
		t.Errorf("Expected brand class to contain 'navbar-brand', got '%s'", brandClass)
	}

	// Check text content
	if len(brand.ChildNodes()) != 1 {
		t.Error("Expected brand to have 1 text node")
	}
}

func TestNavBar_Toggler(t *testing.T) {
	navbar := NavBar().Toggler("navbarNav")

	// Check that toggler was added
	container := navbar.Element().ChildNodes()[0].(Element)
	togglerElements := container.ChildNodes()

	if len(togglerElements) != 1 {
		t.Errorf("Expected 1 toggler element, got %d", len(togglerElements))
	}

	toggler := togglerElements[0].(Element)
	if toggler.TagName() != "BUTTON" {
		t.Errorf("Expected toggler tag 'BUTTON', got '%s'", toggler.TagName())
	}

	togglerClass := toggler.GetAttribute("class")
	if !strings.Contains(togglerClass, "navbar-toggler") {
		t.Errorf("Expected toggler class to contain 'navbar-toggler', got '%s'", togglerClass)
	}

	// Check data attributes
	target := toggler.GetAttribute("data-bs-target")
	if target != "#navbarNav" {
		t.Errorf("Expected data-bs-target '#navbarNav', got '%s'", target)
	}

	// Check toggler icon exists
	icon := toggler.ChildNodes()[0].(Element)
	iconClass := icon.GetAttribute("class")
	if !strings.Contains(iconClass, "navbar-toggler-icon") {
		t.Errorf("Expected icon class to contain 'navbar-toggler-icon', got '%s'", iconClass)
	}
}

func TestNavBar_Collapse(t *testing.T) {
	navbar := NavBar()
	collapse := navbar.Collapse("navbarNav")

	// Check collapse element
	if collapse.TagName() != "DIV" {
		t.Errorf("Expected collapse tag 'DIV', got '%s'", collapse.TagName())
	}

	collapseClass := collapse.GetAttribute("class")
	if !strings.Contains(collapseClass, "collapse") {
		t.Errorf("Expected collapse class to contain 'collapse', got '%s'", collapseClass)
	}
	if !strings.Contains(collapseClass, "navbar-collapse") {
		t.Errorf("Expected collapse class to contain 'navbar-collapse', got '%s'", collapseClass)
	}

	collapseId := collapse.GetAttribute("id")
	if collapseId != "navbarNav" {
		t.Errorf("Expected collapse id 'navbarNav', got '%s'", collapseId)
	}
}

func TestNavBar_NavContent(t *testing.T) {
	navbar := NavBar()
	navContent := navbar.NavContent(true, "me-auto")

	// Check nav content element
	if navContent.TagName() != "UL" {
		t.Errorf("Expected nav content tag 'UL', got '%s'", navContent.TagName())
	}

	navClass := navContent.GetAttribute("class")
	if !strings.Contains(navClass, "navbar-nav") {
		t.Errorf("Expected nav class to contain 'navbar-nav', got '%s'", navClass)
	}
	if !strings.Contains(navClass, "me-auto") {
		t.Errorf("Expected nav class to contain 'me-auto', got '%s'", navClass)
	}
}

func TestNavBar_NavContentDiv(t *testing.T) {
	navbar := NavBar()
	navContent := navbar.NavContent(false)

	// Check that div is used instead of ul
	if navContent.TagName() != "DIV" {
		t.Errorf("Expected nav content tag 'DIV', got '%s'", navContent.TagName())
	}

	navClass := navContent.GetAttribute("class")
	if !strings.Contains(navClass, "navbar-nav") {
		t.Errorf("Expected nav class to contain 'navbar-nav', got '%s'", navClass)
	}
}

func TestNavBar_Complete(t *testing.T) {
	// Build a complete navbar
	navbar := NavBar(WithClass("navbar-expand-lg"), WithClass("bg-body-tertiary")).
		Brand("Navbar", "#").
		Toggler("navbarNav")

	collapse := navbar.Collapse("navbarNav")
	navContent := navbar.NavContent(true, "me-auto", "mb-2", "mb-lg-0")
	collapse.AppendChild(navContent)

	// Verify structure
	root := navbar.Element()
	if root.TagName() != "NAV" {
		t.Errorf("Expected root tag 'NAV', got '%s'", root.TagName())
	}

	// Check classes
	rootClass := root.GetAttribute("class")
	if !strings.Contains(rootClass, "navbar") {
		t.Error("Expected navbar class")
	}
	if !strings.Contains(rootClass, "navbar-expand-lg") {
		t.Error("Expected navbar-expand-lg class")
	}
	if !strings.Contains(rootClass, "bg-body-tertiary") {
		t.Error("Expected bg-body-tertiary class")
	}

	// Check container has brand, toggler, and collapse
	container := root.ChildNodes()[0].(Element)
	containerChildren := container.ChildNodes()
	if len(containerChildren) != 3 {
		t.Errorf("Expected 3 children (brand, toggler, collapse), got %d", len(containerChildren))
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVITEM TESTS

func TestNavItem_Basic(t *testing.T) {
	navItem := NavItem("/home", false, false, "Home")

	// Check element type
	if tagName := navItem.Element().TagName(); tagName != "A" {
		t.Errorf("Expected tag name 'A', got '%s'", tagName)
	}

	// Check class
	class := navItem.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-link") {
		t.Errorf("Expected class to contain 'nav-link', got '%s'", class)
	}

	// Check href
	href := navItem.Element().GetAttribute("href")
	if href != "/home" {
		t.Errorf("Expected href '/home', got '%s'", href)
	}

	// Check text content
	children := navItem.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (text node), got %d", len(children))
	}
}

func TestNavItem_Active(t *testing.T) {
	navItem := NavItem("#", true, false, "Home")

	class := navItem.Element().GetAttribute("class")
	if !strings.Contains(class, "active") {
		t.Errorf("Expected class to contain 'active', got '%s'", class)
	}

	ariaCurrent := navItem.Element().GetAttribute("aria-current")
	if ariaCurrent != "page" {
		t.Errorf("Expected aria-current 'page', got '%s'", ariaCurrent)
	}
}

func TestNavItem_Disabled(t *testing.T) {
	navItem := NavItem("#", false, true, "Disabled")

	class := navItem.Element().GetAttribute("class")
	if !strings.Contains(class, "disabled") {
		t.Errorf("Expected class to contain 'disabled', got '%s'", class)
	}

	ariaDisabled := navItem.Element().GetAttribute("aria-disabled")
	if ariaDisabled != "true" {
		t.Errorf("Expected aria-disabled 'true', got '%s'", ariaDisabled)
	}
}

func TestNavItem_WithContent(t *testing.T) {
	icon := Icon("home")
	navItem := NavItem("/dashboard", true, false, icon, " Dashboard")

	// Check active state
	class := navItem.Element().GetAttribute("class")
	if !strings.Contains(class, "active") {
		t.Errorf("Expected class to contain 'active', got '%s'", class)
	}

	// Check children (icon + text)
	children := navItem.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (icon + text), got %d", len(children))
	}

	// Check that first child is the icon element
	if children[0].NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected first child to be an element (icon)")
	}
}

func TestNavItem_Component(t *testing.T) {
	navItem := NavItem("#", false, false, "Test")
	if navItem.component.name != NavComponent {
		t.Errorf("Expected component type %s, got %s", NavComponent, navItem.component.name)
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVSPACER TESTS

func TestNavSpacer_Basic(t *testing.T) {
	spacer := NavSpacer()

	// Check element type
	if tagName := spacer.Element().TagName(); tagName != "DIV" {
		t.Errorf("Expected tag name 'DIV', got '%s'", tagName)
	}

	// Check default class
	class := spacer.Element().GetAttribute("class")
	if !strings.Contains(class, "me-auto") {
		t.Errorf("Expected class to contain 'me-auto', got '%s'", class)
	}
}

func TestNavSpacer_WithCustomClass(t *testing.T) {
	spacer := NavSpacer(WithClass("ms-auto"))

	class := spacer.Element().GetAttribute("class")
	if !strings.Contains(class, "ms-auto") {
		t.Errorf("Expected class to contain 'ms-auto', got '%s'", class)
	}
	// Should still have the default class unless explicitly overridden
	if !strings.Contains(class, "me-auto") {
		t.Errorf("Expected class to contain 'me-auto', got '%s'", class)
	}
}

func TestNavSpacer_ComponentInterface(t *testing.T) {
	spacer := NavSpacer()
	if spacer.Element() == nil {
		t.Error("Expected Element() to return non-nil")
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVDIVIDER TESTS

func TestNavDivider_Vertical(t *testing.T) {
	divider := NavDivider(true)

	// Check element type for vertical divider
	if tagName := divider.Element().TagName(); tagName != "DIV" {
		t.Errorf("Expected tag name 'DIV', got '%s'", tagName)
	}

	// Check vertical divider class
	class := divider.Element().GetAttribute("class")
	if !strings.Contains(class, "vr") {
		t.Errorf("Expected class to contain 'vr', got '%s'", class)
	}
}

func TestNavDivider_Horizontal(t *testing.T) {
	divider := NavDivider(false)

	// Check element type for horizontal divider
	if tagName := divider.Element().TagName(); tagName != "HR" {
		t.Errorf("Expected tag name 'HR', got '%s'", tagName)
	}

	// Check horizontal divider class
	class := divider.Element().GetAttribute("class")
	if !strings.Contains(class, "my-2") {
		t.Errorf("Expected class to contain 'my-2', got '%s'", class)
	}
}

func TestNavDivider_WithCustomClass(t *testing.T) {
	divider := NavDivider(true, WithClass("mx-3"))

	class := divider.Element().GetAttribute("class")
	if !strings.Contains(class, "mx-3") {
		t.Errorf("Expected class to contain 'mx-3', got '%s'", class)
	}
	// Should still have the default class
	if !strings.Contains(class, "vr") {
		t.Errorf("Expected class to contain 'vr', got '%s'", class)
	}
}

func TestNavDivider_ComponentInterface(t *testing.T) {
	divider := NavDivider(true)
	if divider.Element() == nil {
		t.Error("Expected Element() to return non-nil")
	}

	divider2 := NavDivider(false)
	if divider2.Element() == nil {
		t.Error("Expected Element() to return non-nil")
	}
}

///////////////////////////////////////////////////////////////////////////////
// INTEGRATION TESTS

func TestNav_WithSpacerAndDivider(t *testing.T) {
	// Test that NavSpacer and NavDivider work correctly within a Nav component
	nav := Nav().
		Append(NavItem("#", true, false, "Home")).
		Append(NavDivider(true)).
		Append(NavItem("#", false, false, "About")).
		Append(NavSpacer()).
		Append(NavItem("#", false, false, "Login"))

	// Check that nav has the expected number of children
	children := nav.Element().ChildNodes()
	expectedChildren := 5 // NavItem + NavDivider + NavItem + NavSpacer + NavItem
	if len(children) != expectedChildren {
		t.Errorf("Expected %d children, got %d", expectedChildren, len(children))
	}

	// Check specific child types and classes
	// First child should be NavItem (anchor)
	if children[0].NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected first child to be an element")
	}

	// Second child should be NavDivider (div with vr class)
	dividerEl := children[1]
	if dividerEl.NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected second child to be an element")
	}
	if tagName := dividerEl.(dom.Element).TagName(); tagName != "DIV" {
		t.Errorf("Expected divider tag name 'DIV', got '%s'", tagName)
	}

	// Fourth child should be NavSpacer (div with me-auto class)
	spacerEl := children[3]
	if spacerEl.NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected fourth child to be an element")
	}
	if tagName := spacerEl.(dom.Element).TagName(); tagName != "DIV" {
		t.Errorf("Expected spacer tag name 'DIV', got '%s'", tagName)
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVDROPDOWN TESTS

func TestNavDropdown_Basic(t *testing.T) {
	dropdown := NavDropdown("More", false,
		NavDropdownItem("#", false, "Action"),
		NavDropdownItem("#", false, "Another action"),
	)

	// Check element type
	if tagName := dropdown.Element().TagName(); tagName != "DIV" {
		t.Errorf("Expected tag name 'DIV', got '%s'", tagName)
	}

	// Check dropdown container class
	class := dropdown.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-item") || !strings.Contains(class, "dropdown") {
		t.Errorf("Expected class to contain 'nav-item dropdown', got '%s'", class)
	}

	// Check children (toggle + menu)
	children := dropdown.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (toggle + menu), got %d", len(children))
	}

	// Check dropdown toggle
	toggle := children[0].(dom.Element)
	if tagName := toggle.TagName(); tagName != "A" {
		t.Errorf("Expected toggle tag name 'A', got '%s'", tagName)
	}
	toggleClass := toggle.GetAttribute("class")
	if !strings.Contains(toggleClass, "nav-link") || !strings.Contains(toggleClass, "dropdown-toggle") {
		t.Errorf("Expected toggle class to contain 'nav-link dropdown-toggle', got '%s'", toggleClass)
	}

	// Check dropdown menu
	menu := children[1].(dom.Element)
	if tagName := menu.TagName(); tagName != "UL" {
		t.Errorf("Expected menu tag name 'UL', got '%s'", tagName)
	}
	menuClass := menu.GetAttribute("class")
	if !strings.Contains(menuClass, "dropdown-menu") {
		t.Errorf("Expected menu class to contain 'dropdown-menu', got '%s'", menuClass)
	}
}

func TestNavDropdown_Active(t *testing.T) {
	dropdown := NavDropdown("Active", true, NavDropdownItem("#", false, "Item"))

	// Find the toggle element
	children := dropdown.Element().ChildNodes()
	toggle := children[0].(dom.Element)

	// Check active state
	class := toggle.GetAttribute("class")
	if !strings.Contains(class, "active") {
		t.Errorf("Expected class to contain 'active', got '%s'", class)
	}

	ariaCurrent := toggle.GetAttribute("aria-current")
	if ariaCurrent != "page" {
		t.Errorf("Expected aria-current 'page', got '%s'", ariaCurrent)
	}
}

func TestNavDropdownItem_Basic(t *testing.T) {
	item := NavDropdownItem("/profile", false, "My Profile")

	// Check element type
	if tagName := item.Element().TagName(); tagName != "LI" {
		t.Errorf("Expected tag name 'LI', got '%s'", tagName)
	}

	// Check link inside li
	children := item.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (link), got %d", len(children))
	}

	link := children[0].(dom.Element)
	if tagName := link.TagName(); tagName != "A" {
		t.Errorf("Expected link tag name 'A', got '%s'", tagName)
	}

	class := link.GetAttribute("class")
	if !strings.Contains(class, "dropdown-item") {
		t.Errorf("Expected class to contain 'dropdown-item', got '%s'", class)
	}

	href := link.GetAttribute("href")
	if href != "/profile" {
		t.Errorf("Expected href '/profile', got '%s'", href)
	}
}

func TestNavDropdownItem_Active(t *testing.T) {
	item := NavDropdownItem("#", true, "Active Item")

	// Get the link element
	children := item.Element().ChildNodes()
	link := children[0].(dom.Element)

	class := link.GetAttribute("class")
	if !strings.Contains(class, "active") {
		t.Errorf("Expected class to contain 'active', got '%s'", class)
	}

	ariaCurrent := link.GetAttribute("aria-current")
	if ariaCurrent != "true" {
		t.Errorf("Expected aria-current 'true', got '%s'", ariaCurrent)
	}
}

func TestNavDropdownDivider(t *testing.T) {
	divider := NavDropdownDivider()

	// Check element type
	if tagName := divider.Element().TagName(); tagName != "LI" {
		t.Errorf("Expected tag name 'LI', got '%s'", tagName)
	}

	// Check hr inside li
	children := divider.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (hr), got %d", len(children))
	}

	hr := children[0].(dom.Element)
	if tagName := hr.TagName(); tagName != "HR" {
		t.Errorf("Expected hr tag name 'HR', got '%s'", tagName)
	}

	class := hr.GetAttribute("class")
	if !strings.Contains(class, "dropdown-divider") {
		t.Errorf("Expected class to contain 'dropdown-divider', got '%s'", class)
	}
}

func TestNavDropdownHeader(t *testing.T) {
	header := NavDropdownHeader("Settings")

	// Check element type
	if tagName := header.Element().TagName(); tagName != "LI" {
		t.Errorf("Expected tag name 'LI', got '%s'", tagName)
	}

	// Check h6 inside li
	children := header.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (h6), got %d", len(children))
	}

	h6 := children[0].(dom.Element)
	if tagName := h6.TagName(); tagName != "H6" {
		t.Errorf("Expected h6 tag name 'H6', got '%s'", tagName)
	}

	class := h6.GetAttribute("class")
	if !strings.Contains(class, "dropdown-header") {
		t.Errorf("Expected class to contain 'dropdown-header', got '%s'", class)
	}
}

func TestNavDropdown_Integration(t *testing.T) {
	// Test a complete dropdown in a nav
	nav := Nav().
		Append(NavItem("#", true, false, "Home")).
		Append(NavDropdown("More", false,
			NavDropdownHeader("Account"),
			NavDropdownItem("/profile", false, "Profile"),
			NavDropdownItem("/settings", false, "Settings"),
			NavDropdownDivider(),
			NavDropdownHeader("Help"),
			NavDropdownItem("/help", false, "Documentation"),
			NavDropdownItem("/contact", false, "Contact"),
		)).
		Append(NavItem("#", false, false, "About"))

	// Check nav has expected children
	children := nav.Element().ChildNodes()
	expectedChildren := 3 // NavItem + NavDropdown + NavItem
	if len(children) != expectedChildren {
		t.Errorf("Expected %d children, got %d", expectedChildren, len(children))
	}

	// Check that middle child is the dropdown
	dropdownEl := children[1].(dom.Element)
	class := dropdownEl.GetAttribute("class")
	if !strings.Contains(class, "dropdown") {
		t.Errorf("Expected middle child to have dropdown class, got '%s'", class)
	}
}

func TestNavDropdown_InNavBar(t *testing.T) {
	// Test dropdown can be added to navbar nav content
	navbar := NavBar(WithClass("navbar-expand-lg"))
	_ = navbar.Collapse("test") // Create collapse container but don't need to use it
	navContent := navbar.NavContent(false)

	// Create dropdown and add to nav content
	dropdown := NavDropdown("Services", false,
		NavDropdownItem("/web", false, "Web Design"),
		NavDropdownItem("/mobile", false, "Mobile Apps"),
	)

	navContent.AppendChild(dropdown.Element())

	// Verify the dropdown was added correctly
	children := navContent.ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child in nav content, got %d", len(children))
	}

	dropdownEl := children[0].(dom.Element)
	class := dropdownEl.GetAttribute("class")
	if !strings.Contains(class, "nav-item") || !strings.Contains(class, "dropdown") {
		t.Errorf("Expected dropdown class 'nav-item dropdown', got '%s'", class)
	}
}
