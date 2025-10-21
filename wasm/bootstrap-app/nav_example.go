package main

import (
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// NavExamples demonstrates various navigation components
func NavExamples() Component {
	container := bs.Container()

	container.Append(
		bs.Heading(2, bs.WithClass("mb-4")).Append("Navigation Components"),
	)

	// Basic Nav
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3")).Append("Basic Nav"),
	)

	basicNav := bs.Nav(bs.WithClass("border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, false, "Another Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(basicNav))

	// Nav with Tabs
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Tabs"),
	)

	tabsNav := bs.Nav(bs.WithTabs(), bs.WithClass("border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, false, "Another Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(tabsNav))

	// Nav with Pills
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Pills"),
	)

	pillsNav := bs.Nav(bs.WithPills(), bs.WithClass("border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, false, "Another Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(pillsNav))

	// Nav with Underline
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Underline"),
	)

	underlineNav := bs.Nav(bs.WithUnderline(), bs.WithClass("border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, false, "Another Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(underlineNav))

	// Centered Nav
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Centered"),
	)

	centeredNav := bs.Nav(bs.WithClass("justify-content-center", "border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, false, "Another Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(centeredNav))

	// Right-aligned Nav
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Right-aligned"),
	)

	rightNav := bs.Nav(bs.WithClass("justify-content-end", "border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, false, "Another Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(rightNav))

	// Vertical Nav
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Vertical"),
	)

	verticalNav := bs.Nav(bs.WithClass("flex-column", "border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, false, "Another Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(verticalNav))

	// Pills with Fill
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Pills - Fill"),
	)

	fillNav := bs.Nav(bs.WithPills(), bs.WithClass("nav-fill", "border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Much longer nav link")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(fillNav))

	// Pills with Justify
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Pills - Justified"),
	)

	justifyNav := bs.Nav(bs.WithPills(), bs.WithClass("nav-justified", "border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Much longer nav link")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(justifyNav))

	// Nav with Icons
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Nav with Icons"),
	)

	iconNav := bs.Nav(bs.WithPills(), bs.WithClass("border", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, bs.Icon("house-fill", bs.WithMargin(bs.END, 2)), "Home")).
		Append(bs.NavItem("#", false, false, bs.Icon("person-fill", bs.WithMargin(bs.END, 2)), "Profile")).
		Append(bs.NavItem("#", false, false, bs.Icon("envelope-fill", bs.WithMargin(bs.END, 2)), "Messages")).
		Append(bs.NavItem("#", false, false, bs.Icon("gear-fill", bs.WithMargin(bs.END, 2)), "Settings"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(iconNav))

	// Tabs with different colors
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Colored Pills"),
	)

	coloredRow := bs.Container(bs.WithClass("row", "g-3"))

	// Primary pills
	primaryPills := bs.Container(bs.WithClass("col-md-6")).
		Append(
			bs.Nav(bs.WithPills(), bs.WithClass("border", "p-3", "rounded")).
				Append(bs.NavItem("#", true, false, "Primary")).
				Append(bs.NavItem("#", false, false, "Link")).
				Append(bs.NavItem("#", false, false, "Link")),
		)

	// Success pills with vertical layout
	successPills := bs.Container(bs.WithClass("col-md-6")).
		Append(
			bs.Nav(bs.WithPills(), bs.WithClass("flex-column", "border", "p-3", "rounded")).
				Append(bs.NavItem("#", true, false, "Success")).
				Append(bs.NavItem("#", false, false, "Link")).
				Append(bs.NavItem("#", false, false, "Link")),
		)

	coloredRow.Append(primaryPills, successPills)
	container.Append(bs.Container(bs.WithClass("mb-4")).Append(coloredRow))

	// Nav with Background Colors
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Nav with Background Colors"),
	)

	// Light background
	lightBgNav := bs.Nav(bs.WithClass("bg-light", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Home")).
		Append(bs.NavItem("#", false, false, "Features")).
		Append(bs.NavItem("#", false, false, "Pricing")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(lightBgNav))

	// Dark background with data-bs-theme
	darkBgNav := bs.Nav(bs.WithClass("bg-dark", "p-3", "rounded"), bs.WithAttribute("data-bs-theme", "dark")).
		Append(bs.NavItem("#", true, false, "Home")).
		Append(bs.NavItem("#", false, false, "Features")).
		Append(bs.NavItem("#", false, false, "Pricing")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(darkBgNav))

	// Primary background
	primaryBgNav := bs.Nav(bs.WithPills(), bs.WithClass("bg-primary", "p-3", "rounded"), bs.WithAttribute("data-bs-theme", "dark")).
		Append(bs.NavItem("#", true, false, "Home")).
		Append(bs.NavItem("#", false, false, "Features")).
		Append(bs.NavItem("#", false, false, "Pricing")).
		Append(bs.NavItem("#", false, true, "Disabled"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(primaryBgNav))

	// Success background with pills
	successBgNav := bs.Nav(bs.WithPills(), bs.WithClass("bg-success", "p-3", "rounded"), bs.WithAttribute("data-bs-theme", "dark")).
		Append(bs.NavItem("#", true, false, "Active")).
		Append(bs.NavItem("#", false, false, "Link")).
		Append(bs.NavItem("#", false, false, "Another"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(successBgNav))

	// Info background with tabs
	infoBgNav := bs.Nav(bs.WithTabs(), bs.WithClass("bg-info", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Dashboard")).
		Append(bs.NavItem("#", false, false, "Reports")).
		Append(bs.NavItem("#", false, false, "Settings"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(infoBgNav))

	// Warning background
	warningBgNav := bs.Nav(bs.WithClass("bg-warning", "p-3", "rounded")).
		Append(bs.NavItem("#", true, false, "Home")).
		Append(bs.NavItem("#", false, false, "About")).
		Append(bs.NavItem("#", false, false, "Contact"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(warningBgNav))

	// Secondary background with pills
	secondaryBgNav := bs.Nav(bs.WithPills(), bs.WithClass("bg-secondary", "p-3", "rounded"), bs.WithAttribute("data-bs-theme", "dark")).
		Append(bs.NavItem("#", true, false, "Overview")).
		Append(bs.NavItem("#", false, false, "Details")).
		Append(bs.NavItem("#", false, false, "More"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(secondaryBgNav))

	// Nav with Spacers
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Nav with Spacers"),
	)

	// Horizontal nav with spacer
	spacerNav := bs.Nav(bs.WithClass("border", "rounded", "p-2")).
		Append(bs.NavItem("#", false, false, "Home")).
		Append(bs.NavItem("#", false, false, "About")).
		Append(bs.NavSpacer()).
		Append(bs.NavItem("#", false, false, "Login")).
		Append(bs.NavItem("#", false, false, "Signup"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(spacerNav))

	// Pills nav with custom spacer (margin start)
	spacerPillsNav := bs.Nav(bs.WithPills(), bs.WithClass("border", "rounded", "p-2")).
		Append(bs.NavItem("#", true, false, "Dashboard")).
		Append(bs.NavItem("#", false, false, "Analytics")).
		Append(bs.NavSpacer(bs.WithClass("ms-auto"))).
		Append(bs.NavItem("#", false, false, "Settings"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(spacerPillsNav))

	// Nav with Dividers
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Nav with Dividers"),
	)

	// Horizontal nav with vertical dividers
	dividerNav := bs.Nav(bs.WithClass("border", "rounded", "p-2")).
		Append(bs.NavItem("#", true, false, "Home")).
		Append(bs.NavDivider(true)).
		Append(bs.NavItem("#", false, false, "Products")).
		Append(bs.NavDivider(true)).
		Append(bs.NavItem("#", false, false, "Services")).
		Append(bs.NavDivider(true)).
		Append(bs.NavItem("#", false, false, "Contact"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(dividerNav))

	// Vertical nav with horizontal dividers
	verticalDividerNav := bs.Nav(bs.WithClass("flex-column", "border", "rounded", "p-2")).
		Append(bs.NavItem("#", true, false, "Dashboard")).
		Append(bs.NavDivider(false)).
		Append(bs.NavItem("#", false, false, "Reports")).
		Append(bs.NavDivider(false)).
		Append(bs.NavItem("#", false, false, "Analytics")).
		Append(bs.NavDivider(false)).
		Append(bs.NavItem("#", false, false, "Settings"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(verticalDividerNav))

	// Complex nav with spacers and dividers
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Complex Nav with Spacers and Dividers"),
	)

	complexNav := bs.Nav(bs.WithPills(), bs.WithClass("border", "rounded", "p-2")).
		Append(bs.NavItem("#", true, false, "Home")).
		Append(bs.NavDivider(true, bs.WithClass("mx-2"))).
		Append(bs.NavItem("#", false, false, "About")).
		Append(bs.NavSpacer()).
		Append(bs.NavItem("#", false, false, "Profile")).
		Append(bs.NavDivider(true, bs.WithClass("mx-2"))).
		Append(bs.NavItem("#", false, false, "Logout"))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(complexNav))

	// Nav with Dropdowns
	container.Append(
		bs.Heading(4, bs.WithClass("mb-3", "mt-4")).Append("Nav with Dropdowns"),
	)

	// Basic nav with dropdown
	dropdownNav := bs.Nav(bs.WithTabs(), bs.WithClass("border", "rounded", "p-2")).
		Append(bs.NavItem("#", true, false, "Home")).
		Append(bs.NavDropdown("Services", false,
			bs.NavDropdownItem("/web-design", false, "Web Design"),
			bs.NavDropdownItem("/web-dev", false, "Web Development"),
			bs.NavDropdownDivider(),
			bs.NavDropdownItem("/consulting", false, "Consulting"),
		)).
		Append(bs.NavDropdown("Account", false,
			bs.NavDropdownHeader("User"),
			bs.NavDropdownItem("/profile", false, "Profile"),
			bs.NavDropdownItem("/settings", false, "Settings"),
			bs.NavDropdownDivider(),
			bs.NavDropdownHeader("Support"),
			bs.NavDropdownItem("/help", false, "Help Center"),
			bs.NavDropdownItem("/contact", false, "Contact Us"),
		)).
		Append(bs.NavItem("#", false, false, "About"))

	container.Append(bs.Container(bs.WithClass("mb-3")).Append(dropdownNav))

	// Pills nav with dropdown and spacer
	pillsDropdownNav := bs.Nav(bs.WithPills(), bs.WithClass("border", "rounded", "p-2")).
		Append(bs.NavItem("#", false, false, "Dashboard")).
		Append(bs.NavDropdown("Tools", false,
			bs.NavDropdownItem("/analytics", false, "Analytics"),
			bs.NavDropdownItem("/reports", false, "Reports"),
			bs.NavDropdownItem("/export", false, "Export Data"),
		)).
		Append(bs.NavSpacer()).
		Append(bs.NavDropdown("Profile", false,
			bs.NavDropdownItem("/account", false, "Account Settings"),
			bs.NavDropdownDivider(),
			bs.NavDropdownItem("/logout", false, "Sign Out"),
		))

	container.Append(bs.Container(bs.WithClass("mb-4")).Append(pillsDropdownNav))

	return container
}
