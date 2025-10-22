package main

import (
	"strings"

	. "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

// NavBarExamples demonstrates various Bootstrap navbar components
func NavBarExamples() Component {
	container := bs.Container(bs.WithClass("my-5"))

	container.Append(
		bs.Heading(2, bs.WithClass("mb-4")).Append("NavBar Examples"),
		dom.GetWindow().Document().CreateTextNode("Bootstrap navbar components with responsive behavior, branding, and navigation."),
	)

	// Basic Navbar
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Basic NavBar"),
	)

	basicNavbar := bs.NavBar(bs.WithClass("bg-body-tertiary", "border", "rounded"))
	basicNavbar.Brand("Navbar", "#")

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(basicNavbar))

	// Navbar with Brand and Links
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("NavBar with Navigation"),
	)

	navWithLinks := bs.NavBar(bs.WithClass("navbar-expand-lg", "bg-body-tertiary", "border", "rounded"))
	navWithLinks.Brand("Navbar", "#").Toggler("navbarNav")

	collapse := navWithLinks.Collapse("navbarNav")
	navContent := navWithLinks.NavContent(false, "me-auto")

	// Add navigation links
	navContent.AppendChild(bs.NavItem("#", true, false, "Home").Element())
	navContent.AppendChild(bs.NavItem("#", false, false, "Link").Element())
	navContent.AppendChild(bs.NavItem("#", false, true, "Disabled").Element())

	collapse.AppendChild(navContent)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(navWithLinks))

	// Navbar with Search Form
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("NavBar with Search"),
	)

	navWithSearch := bs.NavBar(bs.WithClass("navbar-expand-lg", "bg-body-tertiary", "border", "rounded"))
	navWithSearch.Brand("Navbar", "#").Toggler("navbarSearch")

	searchCollapse := navWithSearch.Collapse("navbarSearch")
	searchNavContent := navWithSearch.NavContent(false, "me-auto", "mb-2", "mb-lg-0")

	// Add navigation links
	searchNavContent.AppendChild(bs.NavItem("#", true, false, "Home").Element())
	searchNavContent.AppendChild(bs.NavItem("#", false, false, "Features").Element())
	searchNavContent.AppendChild(bs.NavItem("#", false, false, "Pricing").Element())

	// Add search form
	searchForm := createSearchForm()

	searchCollapse.AppendChild(searchNavContent)
	searchCollapse.AppendChild(searchForm)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(navWithSearch))

	// Dark Navbar
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Dark NavBar"),
	)

	darkNavbar := bs.NavBar(
		bs.WithClass("navbar-expand-lg", "bg-dark", "border", "rounded"),
		bs.WithAttribute("data-bs-theme", "dark"),
	)
	darkNavbar.Brand("Navbar", "#").Toggler("navbarDark")

	darkCollapse := darkNavbar.Collapse("navbarDark")
	darkNavContent := darkNavbar.NavContent(false, "me-auto")

	// Add navigation links
	darkNavContent.AppendChild(bs.NavItem("#", true, false, "Home").Element())
	darkNavContent.AppendChild(bs.NavItem("#", false, false, "Features").Element())
	darkNavContent.AppendChild(bs.NavItem("#", false, false, "About").Element())

	darkCollapse.AppendChild(darkNavContent)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(darkNavbar))

	// Primary Themed Navbar
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Primary NavBar"),
	)

	primaryNavbar := bs.NavBar(
		bs.WithClass("navbar-expand-lg", "bg-primary", "border", "rounded"),
		bs.WithAttribute("data-bs-theme", "dark"),
	)
	primaryNavbar.Brand("Brand", "#").Toggler("navbarPrimary")

	primaryCollapse := primaryNavbar.Collapse("navbarPrimary")
	primaryNavContent := primaryNavbar.NavContent(false, "me-auto")

	// Add navigation links
	primaryNavContent.AppendChild(bs.NavItem("#", true, false, "Dashboard").Element())
	primaryNavContent.AppendChild(bs.NavItem("#", false, false, "Orders").Element())
	primaryNavContent.AppendChild(bs.NavItem("#", false, false, "Products").Element())

	primaryCollapse.AppendChild(primaryNavContent)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(primaryNavbar))

	// Navbar with Icon and Text
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("NavBar with Icon"),
	)

	logoNavbar := bs.NavBar(bs.WithClass("navbar-expand-lg", "bg-body-tertiary", "border", "rounded"))

	// Create brand with icon and text
	logoIcon := bs.Icon("house-fill", bs.WithClass("me-2"))

	logoNavbar.BrandWithContent("#", logoIcon, "Bootstrap")
	logoNavbar.Toggler("navbarLogo")

	logoCollapse := logoNavbar.Collapse("navbarLogo")
	logoNavContent := logoNavbar.NavContent(false, "me-auto")

	// Add navigation links
	logoNavContent.AppendChild(bs.NavItem("#", true, false, "Home").Element())
	logoNavContent.AppendChild(bs.NavItem("#", false, false, "Products").Element())
	logoNavContent.AppendChild(bs.NavItem("#", false, false, "Contact").Element())

	logoCollapse.AppendChild(logoNavContent)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(logoNavbar))

	// Fixed Top Navbar (shown without fixed positioning for demo)
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Fixed Top Style NavBar"),
		dom.GetWindow().Document().CreateTextNode("(Fixed positioning disabled for demo purposes)"),
	)

	fixedNavbar := bs.NavBar(
		bs.WithClass("navbar-expand-lg", "bg-success", "border", "rounded"),
		bs.WithAttribute("data-bs-theme", "dark"),
	)
	fixedNavbar.Brand("Fixed Brand", "#").Toggler("navbarFixed")

	fixedCollapse := fixedNavbar.Collapse("navbarFixed")
	fixedNavContent := fixedNavbar.NavContent(false, "me-auto")

	// Add navigation links
	fixedNavContent.AppendChild(bs.NavItem("#", true, false, "Dashboard").Element())
	fixedNavContent.AppendChild(bs.NavItem("#", false, false, "Analytics").Element())
	fixedNavContent.AppendChild(bs.NavItem("#", false, false, "Settings").Element())

	// Add user dropdown area (simplified)
	userText := dom.GetWindow().Document().CreateElement("SPAN")
	userText.SetAttribute("class", "navbar-text")
	userText.AppendChild(dom.GetWindow().Document().CreateTextNode("Welcome, User"))

	fixedCollapse.AppendChild(fixedNavContent)
	fixedCollapse.AppendChild(userText)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(fixedNavbar))

	// Responsive Breakpoint Examples
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Responsive Breakpoints"),
	)

	breakpointRow := bs.Container(bs.WithClass("row", "g-3"))

	// Always collapsed (no expand class)
	alwaysCollapsed := bs.Container(bs.WithClass("col-md-6")).
		Append(
			bs.Heading(6).Append("Always Collapsed"),
			bs.NavBar(bs.WithClass("bg-info", "border", "rounded"), bs.WithAttribute("data-bs-theme", "dark")).
				Brand("Brand", "#").
				Toggler("alwaysCollapsed"),
		)

	// Expand at medium
	expandMd := bs.Container(bs.WithClass("col-md-6")).
		Append(
			bs.Heading(6).Append("Expand at MD+"),
			bs.NavBar(bs.WithClass("navbar-expand-md", "bg-warning", "border", "rounded")).
				Brand("Brand", "#").
				Toggler("expandMd"),
		)

	breakpointRow.Append(alwaysCollapsed, expandMd)
	container.Append(bs.Container(bs.WithClass("mb-4")).Append(breakpointRow))

	// NavBar with Spacers and Dividers
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("NavBar with Spacers and Dividers"),
	)

	spacerNavbar := bs.NavBar(bs.WithClass("navbar-expand-lg", "bg-body-tertiary", "border", "rounded"))
	spacerNavbar.Brand("MyApp", "#").Toggler("navbarSpacer")

	spacerCollapse := spacerNavbar.Collapse("navbarSpacer")
	spacerNavContent := spacerNavbar.NavContent(false, "me-auto")

	// Add navigation links with spacers and dividers
	spacerNavContent.AppendChild(bs.NavItem("#", true, false, "Home").Element())
	spacerNavContent.AppendChild(bs.NavDivider(true, bs.WithClass("mx-2")).Element())
	spacerNavContent.AppendChild(bs.NavItem("#", false, false, "Products").Element())
	spacerNavContent.AppendChild(bs.NavItem("#", false, false, "Services").Element())
	spacerNavContent.AppendChild(bs.NavSpacer().Element())
	spacerNavContent.AppendChild(bs.NavItem("#", false, false, "Login").Element())
	spacerNavContent.AppendChild(bs.NavDivider(true, bs.WithClass("mx-2")).Element())
	spacerNavContent.AppendChild(bs.NavItem("#", false, false, "Sign Up").Element())

	spacerCollapse.AppendChild(spacerNavContent)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(spacerNavbar))

	// NavBar with Dropdowns
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("NavBar with Dropdowns and Icons"),
	)

	dropdownNavbar := bs.NavBar(bs.WithClass("navbar-expand-lg", "bg-body-tertiary", "border", "rounded"))
	dropdownNavbar.Brand("Company", "#").Toggler("navbarDropdown")

	dropdownCollapse := dropdownNavbar.Collapse("navbarDropdown")
	dropdownNavContent := dropdownNavbar.NavContent(false, "me-auto")

	// Add navigation with dropdowns
	dropdownNavContent.AppendChild(bs.NavItem("#", true, false, "Home").Element())

	// Services dropdown with icons and badges
	servicesDropdown := bs.NavDropdown("Services", false,
		bs.NavDropdownHeader("Web Solutions"),
		createDropdownItemWithIcon("/web-design", false, "palette", "Web Design", ""),
		createDropdownItemWithIcon("/web-dev", false, "code-slash", "Development", "NEW"),
		createDropdownItemWithIcon("/ecommerce", false, "cart", "E-commerce", ""),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Marketing"),
		createDropdownItemWithIcon("/seo", false, "search", "SEO Services", ""),
		createDropdownItemWithIcon("/social", false, "share", "Social Media", "HOT"),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Advanced"),
		createDropdownItemWithIcon("/analytics", false, "graph-up", "Analytics", "PRO"),
		createDropdownItemWithIcon("/automation", false, "gear-fill", "Automation", ""),
	)
	dropdownNavContent.AppendChild(servicesDropdown.Element())

	// Products dropdown with multi-level structure
	productsDropdown := bs.NavDropdown("Products", false,
		bs.NavDropdownHeader("Software"),
		createDropdownItemWithIcon("/cms", false, "laptop", "CMS Platform", ""),
		createDropdownItemWithIcon("/crm", false, "people", "CRM System", "POPULAR"),
		createDropdownItemWithIcon("/erp", false, "building", "ERP Solution", "ENTERPRISE"),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Mobile Apps"),
		createDropdownItemWithIcon("/ios-app", false, "phone", "iOS App", ""),
		createDropdownItemWithIcon("/android-app", false, "phone", "Android App", ""),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Cloud Services"),
		createDropdownItemWithIcon("/hosting", false, "cloud", "Web Hosting", ""),
		createDropdownItemWithIcon("/storage", false, "hdd-stack", "Cloud Storage", "50% OFF"),
		createDropdownItemWithIcon("/backup", false, "shield-check", "Backup Services", ""),
	)
	dropdownNavContent.AppendChild(productsDropdown.Element())

	// Company dropdown with icons
	companyDropdown := bs.NavDropdown("Company", false,
		createDropdownItemWithIcon("/about", false, "info-circle", "About Us", ""),
		createDropdownItemWithIcon("/team", false, "people-fill", "Our Team", ""),
		createDropdownItemWithIcon("/careers", false, "briefcase", "Careers", "3 OPEN"),
		bs.NavDropdownDivider(),
		createDropdownItemWithIcon("/news", false, "newspaper", "News & Updates", ""),
		createDropdownItemWithIcon("/contact", false, "envelope", "Contact", ""),
		createDropdownItemWithIcon("/investors", false, "graph-up-arrow", "Investors", ""),
	)
	dropdownNavContent.AppendChild(companyDropdown.Element())

	dropdownNavContent.AppendChild(bs.NavItem("#", false, false, "Portfolio").Element())

	// Add search form and user dropdown on the right
	searchFormDropdown := createSearchForm()
	dropdownCollapse.AppendChild(dropdownNavContent)
	dropdownCollapse.AppendChild(searchFormDropdown)

	// User account dropdown on the right with icons and badges
	userNavContent := dropdownNavbar.NavContent(false)
	userDropdown := bs.NavDropdown("Account", false,
		bs.NavDropdownHeader("Profile"),
		createDropdownItemWithIcon("/profile", false, "person-circle", "My Profile", ""),
		createDropdownItemWithIcon("/orders", false, "box-seam", "My Orders", "2"),
		createDropdownItemWithIcon("/wishlist", false, "heart", "Wishlist", "5"),
		createDropdownItemWithIcon("/settings", false, "gear", "Settings", ""),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Support"),
		createDropdownItemWithIcon("/help", false, "question-circle", "Help Center", ""),
		createDropdownItemWithIcon("/tickets", false, "chat-dots", "Support Tickets", "1"),
		bs.NavDropdownDivider(),
		createDropdownItemWithIcon("/logout", false, "box-arrow-right", "Sign Out", ""),
	)
	userNavContent.AppendChild(userDropdown.Element())
	dropdownCollapse.AppendChild(userNavContent)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(dropdownNavbar))

	// Advanced NavBar with Multi-Level Menus and Enhanced Styling
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Advanced NavBar - Dashboard Style"),
	)

	advancedNavbar := bs.NavBar(bs.WithClass("navbar-expand-lg", "bg-dark", "border", "rounded"), bs.WithAttribute("data-bs-theme", "dark"))
	advancedNavbar.BrandWithContent("#", bs.Icon("building", bs.WithClass("me-2")), "Dashboard")
	advancedNavbar.Toggler("navbarAdvanced")

	advancedCollapse := advancedNavbar.Collapse("navbarAdvanced")
	advancedNavContent := advancedNavbar.NavContent(false, "me-auto")

	// Dashboard navigation
	advancedNavContent.AppendChild(createNavItemWithIcon("#", true, "house-door", "Dashboard"))

	// Analytics dropdown with detailed structure
	analyticsDropdown := bs.NavDropdown("Analytics", false,
		bs.NavDropdownHeader("Traffic Analysis"),
		createDropdownItemWithIcon("/visitors", false, "people", "Visitor Stats", "LIVE"),
		createDropdownItemWithIcon("/pageviews", false, "eye", "Page Views", ""),
		createDropdownItemWithIcon("/bounce-rate", false, "arrow-repeat", "Bounce Rate", ""),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Performance"),
		createDropdownItemWithIcon("/speed", false, "speedometer2", "Site Speed", ""),
		createDropdownItemWithIcon("/errors", false, "exclamation-triangle", "Error Reports", "12"),
		createDropdownItemWithIcon("/uptime", false, "check-circle", "Uptime Monitor", "99.9%"),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Advanced Reports"),
		createDropdownItemWithIcon("/custom-reports", false, "graph-up-arrow", "Custom Reports", "PRO"),
		createDropdownItemWithIcon("/api-analytics", false, "code-square", "API Analytics", "BETA"),
	)
	advancedNavContent.AppendChild(analyticsDropdown.Element())

	// E-commerce dropdown with categories
	ecommerceDropdown := bs.NavDropdown("E-commerce", false,
		bs.NavDropdownHeader("Order Management"),
		createDropdownItemWithIcon("/orders", false, "cart-check", "Orders", "45"),
		createDropdownItemWithIcon("/pending", false, "hourglass-split", "Pending Orders", "7"),
		createDropdownItemWithIcon("/returns", false, "arrow-left-right", "Returns", "3"),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Inventory"),
		createDropdownItemWithIcon("/products", false, "box-seam", "Products", ""),
		createDropdownItemWithIcon("/low-stock", false, "exclamation-diamond", "Low Stock", "WARNING"),
		createDropdownItemWithIcon("/suppliers", false, "truck", "Suppliers", ""),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Financial"),
		createDropdownItemWithIcon("/revenue", false, "currency-dollar", "Revenue", ""),
		createDropdownItemWithIcon("/invoices", false, "receipt", "Invoices", "23"),
		createDropdownItemWithIcon("/taxes", false, "calculator", "Tax Reports", ""),
	)
	advancedNavContent.AppendChild(ecommerceDropdown.Element())

	// Tools & Settings dropdown
	toolsDropdown := bs.NavDropdown("Tools", false,
		bs.NavDropdownHeader("Content Management"),
		createDropdownItemWithIcon("/pages", false, "file-earmark-text", "Pages", ""),
		createDropdownItemWithIcon("/media", false, "images", "Media Library", ""),
		createDropdownItemWithIcon("/seo", false, "search", "SEO Tools", ""),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("User Management"),
		createDropdownItemWithIcon("/users", false, "people-fill", "Users", "1,234"),
		createDropdownItemWithIcon("/roles", false, "shield-check", "Roles & Permissions", ""),
		createDropdownItemWithIcon("/activity", false, "clock-history", "Activity Log", ""),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("System"),
		createDropdownItemWithIcon("/settings", false, "gear-fill", "Settings", ""),
		createDropdownItemWithIcon("/backup", false, "cloud-upload", "Backup", ""),
		createDropdownItemWithIcon("/updates", false, "arrow-up-circle", "Updates", "NEW"),
	)
	advancedNavContent.AppendChild(toolsDropdown.Element())

	advancedNavContent.AppendChild(createNavItemWithIcon("#", false, "envelope", "Messages"))

	// Right side navigation
	rightNavContent := advancedNavbar.NavContent(false)

	// Notifications dropdown
	notificationsDropdown := bs.NavDropdown("Notifications", false,
		bs.NavDropdownHeader("Recent Activity"),
		createDropdownItemWithIcon("/notif1", false, "bell", "New order received", "2m"),
		createDropdownItemWithIcon("/notif2", false, "person-plus", "New user registered", "5m"),
		createDropdownItemWithIcon("/notif3", false, "exclamation-triangle", "Server alert", "10m"),
		bs.NavDropdownDivider(),
		createDropdownItemWithIcon("/all-notifications", false, "bell-fill", "View All", "25"),
		createDropdownItemWithIcon("/notification-settings", false, "gear", "Settings", ""),
	)
	rightNavContent.AppendChild(notificationsDropdown.Element())

	// Admin account dropdown
	adminDropdown := bs.NavDropdown("Admin", false,
		bs.NavDropdownHeader("Account"),
		createDropdownItemWithIcon("/admin-profile", false, "person-circle", "My Profile", ""),
		createDropdownItemWithIcon("/admin-preferences", false, "sliders", "Preferences", ""),
		createDropdownItemWithIcon("/admin-security", false, "shield-lock", "Security", ""),
		bs.NavDropdownDivider(),
		bs.NavDropdownHeader("Quick Actions"),
		createDropdownItemWithIcon("/switch-account", false, "arrow-left-right", "Switch Account", ""),
		createDropdownItemWithIcon("/admin-help", false, "question-circle", "Help & Support", ""),
		bs.NavDropdownDivider(),
		createDropdownItemWithIcon("/logout", false, "box-arrow-right", "Sign Out", ""),
	)
	rightNavContent.AppendChild(adminDropdown.Element())

	advancedCollapse.AppendChild(advancedNavContent)
	advancedCollapse.AppendChild(rightNavContent)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(advancedNavbar))

	// Scrollspy Example
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Scrollspy Navigation"),
		dom.GetWindow().Document().CreateTextNode("Navigation that automatically updates based on scroll position."),
	)

	// Create scrollspy navigation
	scrollspyNav := bs.Nav(bs.WithClass("nav-pills", "flex-column"), bs.WithAttribute("id", "navbar-scrollspy"))
	scrollspyNav.Append(bs.NavItem("#item-1", false, false, "First Item"))
	scrollspyNav.Append(bs.NavItem("#item-2", false, false, "Second Item"))
	scrollspyNav.Append(bs.NavItem("#item-3", false, false, "Third Item"))
	scrollspyNav.Append(bs.NavItem("#item-4", false, false, "Fourth Item"))

	// Create scrollspy body manually to have proper control
	scrollspyBody := dom.GetWindow().Document().CreateElement("DIV")
	scrollspyBody.SetAttribute("data-bs-spy", "scroll")
	scrollspyBody.SetAttribute("data-bs-target", "#navbar-scrollspy")
	scrollspyBody.SetAttribute("data-bs-root-margin", "0px 0px -40%")
	scrollspyBody.SetAttribute("data-bs-smooth-scroll", "true")
	scrollspyBody.SetAttribute("class", "scrollspy-example bg-body-tertiary p-3 rounded-2")
	scrollspyBody.SetAttribute("tabindex", "0")
	scrollspyBody.SetAttribute("style", "height: 300px; overflow-y: auto;")

	// Add content sections to the scrollspy body
	sections := []struct {
		id      string
		title   string
		content string
	}{
		{"item-1", "First Item", "This is some placeholder content for the scrollspy page. Note that as you scroll down the page, the appropriate navigation link is highlighted. It's repeated throughout the component example. We keep adding some more example copy here to emphasize the scrolling and highlighting."},
		{"item-2", "Second Item", "This is some placeholder content for the scrollspy page. Note that as you scroll down the page, the appropriate navigation link is highlighted. It's repeated throughout the component example. We keep adding some more example copy here to emphasize the scrolling and highlighting."},
		{"item-3", "Third Item", "This is some placeholder content for the scrollspy page. Note that as you scroll down the page, the appropriate navigation link is highlighted. It's repeated throughout the component example. We keep adding some more example copy here to emphasize the scrolling and highlighting."},
		{"item-4", "Fourth Item", "This is some placeholder content for the scrollspy page. Note that as you scroll down the page, the appropriate navigation link is highlighted. It's repeated throughout the component example. We keep adding some more example copy here to emphasize the scrolling and highlighting."},
	}

	for _, section := range sections {
		heading := bs.Heading(4, bs.WithAttribute("id", section.id), bs.WithClass("mt-4"))
		heading.Append(section.title)

		para1 := bs.Para()
		para1.Append(section.content)

		para2 := bs.Para()
		para2.Append("Keep in mind that the JavaScript plugin tries to pick the right element among all that may be visible. Multiple visible elements at the same time may cause some issues.")

		scrollspyBody.AppendChild(heading.Element())
		scrollspyBody.AppendChild(para1.Element())
		scrollspyBody.AppendChild(para2.Element())
	}

	// Add scrollspy in a row layout
	scrollspyRow := bs.Container(bs.WithClass("row"))
	scrollspyRow.Append(
		bs.Container(bs.WithClass("col-4")).Append(scrollspyNav),
		bs.Container(bs.WithClass("col-8")).Append(scrollspyBody),
	)

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(scrollspyRow))

	return container
}

// Helper function to create search form
func createSearchForm() Element {
	form := dom.GetWindow().Document().CreateElement("FORM")
	form.SetAttribute("class", "d-flex")
	form.SetAttribute("role", "search")

	// Search input
	input := dom.GetWindow().Document().CreateElement("INPUT")
	input.SetAttribute("class", "form-control me-2")
	input.SetAttribute("type", "search")
	input.SetAttribute("placeholder", "Search")
	input.SetAttribute("aria-label", "Search")

	// Search button
	button := dom.GetWindow().Document().CreateElement("BUTTON")
	button.SetAttribute("class", "btn btn-outline-success")
	button.SetAttribute("type", "submit")
	button.AppendChild(dom.GetWindow().Document().CreateTextNode("Search"))

	form.AppendChild(input)
	form.AppendChild(button)

	return form
}

// Helper function to create dropdown item with icon and optional badge
func createDropdownItemWithIcon(href string, active bool, iconName, text, badgeText string) Component {
	// Create list item
	li := dom.GetWindow().Document().CreateElement("LI")

	// Create dropdown item link
	link := dom.GetWindow().Document().CreateElement("A")
	itemClasses := []string{"dropdown-item", "d-flex", "align-items-center", "gap-2"}

	if active {
		itemClasses = append(itemClasses, "active")
		link.SetAttribute("aria-current", "true")
	}

	link.SetAttribute("class", strings.Join(itemClasses, " "))
	link.SetAttribute("href", href)

	// Create content container
	contentContainer := dom.GetWindow().Document().CreateElement("SPAN")
	contentContainer.SetAttribute("class", "d-flex align-items-center gap-2 w-100")

	// Add icon if provided
	if iconName != "" {
		iconEl := bs.Icon(iconName, bs.WithClass("text-muted"))
		contentContainer.AppendChild(iconEl.Element())
	}

	// Add text
	textSpan := dom.GetWindow().Document().CreateElement("SPAN")
	textSpan.AppendChild(dom.GetWindow().Document().CreateTextNode(text))
	contentContainer.AppendChild(textSpan)

	// Add badge if provided
	if badgeText != "" {
		badgeColor := bs.SECONDARY
		// Choose badge color based on text content
		switch badgeText {
		case "NEW":
			badgeColor = bs.SUCCESS
		case "HOT":
			badgeColor = bs.DANGER
		case "PRO", "ENTERPRISE":
			badgeColor = bs.WARNING
		case "POPULAR":
			badgeColor = bs.PRIMARY
		}

		// Check if it's a number (like "2", "5", etc.)
		if len(badgeText) <= 3 && badgeText != "NEW" && badgeText != "HOT" && badgeText != "PRO" {
			badgeColor = bs.PRIMARY
		}

		badge := bs.PillBadge(bs.WithColor(badgeColor), bs.WithClass("ms-auto"))
		badge.Append(badgeText)
		contentContainer.AppendChild(badge.Element())
	}

	link.AppendChild(contentContainer)
	li.AppendChild(link)

	return &navDropdownItemComponent{
		element: li,
	}
}

// Simple component wrapper for custom dropdown items
type navDropdownItemComponent struct {
	element Element
}

func (n *navDropdownItemComponent) Element() Element {
	return n.element
}

func (n *navDropdownItemComponent) Name() string {
	return "nav-dropdown-item"
}

func (n *navDropdownItemComponent) Append(children ...any) Component {
	for _, child := range children {
		if component, ok := child.(Component); ok {
			n.element.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			n.element.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			n.element.AppendChild(node)
		}
	}
	return n
}

// Helper function to create nav item with icon
func createNavItemWithIcon(href string, active bool, iconName, text string) Element {
	// Create nav item link
	link := dom.GetWindow().Document().CreateElement("A")
	itemClasses := []string{"nav-link", "d-flex", "align-items-center", "gap-2"}

	if active {
		itemClasses = append(itemClasses, "active")
		link.SetAttribute("aria-current", "page")
	}

	link.SetAttribute("class", strings.Join(itemClasses, " "))
	link.SetAttribute("href", href)

	// Add icon if provided
	if iconName != "" {
		iconEl := bs.Icon(iconName, bs.WithClass("me-2"))
		link.AppendChild(iconEl.Element())
	}

	// Add text
	link.AppendChild(dom.GetWindow().Document().CreateTextNode(text))

	return link
}
