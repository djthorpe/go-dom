package main

import (
	. "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
)

func NavExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Navigation Examples"),
	)

	// Create containers for each example
	basicContainer := bs.Container(bs.WithClass("navbar-example"))
	primaryContainer := bs.Container(bs.WithClass("navbar-example", "d-none"))
	darkContainer := bs.Container(bs.WithClass("navbar-example", "d-none"))
	dropdownContainer := bs.Container(bs.WithClass("navbar-example", "d-none"))
	infoContainer := bs.Container(bs.WithClass("navbar-example", "d-none"))
	warningContainer := bs.Container(bs.WithClass("navbar-example", "d-none"))
	dangerContainer := bs.Container(bs.WithClass("navbar-example", "d-none"))
	complexContainer := bs.Container(bs.WithClass("navbar-example", "d-none"))

	// Create a map to easily access containers by name
	containerMap := map[string]Component{
		"basic":    basicContainer,
		"primary":  primaryContainer,
		"dark":     darkContainer,
		"dropdown": dropdownContainer,
		"info":     infoContainer,
		"warning":  warningContainer,
		"danger":   dangerContainer,
		"complex":  complexContainer,
	}

	// Track currently selected example
	var currentExample string = "basic"

	// Helper function to switch examples
	switchToExample := func(exampleName string) func(Event) {
		return func(e Event) {
			e.PreventDefault()

			// Hide current example
			if current, ok := containerMap[currentExample]; ok {
				current.Element().ClassList().Add("d-none")
			}

			// Show selected example
			if selected, ok := containerMap[exampleName]; ok {
				selected.Element().ClassList().Remove("d-none")
				currentExample = exampleName
			}
		}
	}

	// Create buttons for selecting examples
	basicBtn := bs.Button(bs.PRIMARY, bs.WithClass("btn-sm"))
	primaryBtn := bs.OutlineButton(bs.SECONDARY, bs.WithClass("btn-sm"))
	darkBtn := bs.OutlineButton(bs.SECONDARY, bs.WithClass("btn-sm"))
	dropdownBtn := bs.OutlineButton(bs.SECONDARY, bs.WithClass("btn-sm"))
	infoBtn := bs.OutlineButton(bs.SECONDARY, bs.WithClass("btn-sm"))
	warningBtn := bs.OutlineButton(bs.SECONDARY, bs.WithClass("btn-sm"))
	dangerBtn := bs.OutlineButton(bs.SECONDARY, bs.WithClass("btn-sm"))
	complexBtn := bs.OutlineButton(bs.SECONDARY, bs.WithClass("btn-sm"))

	// Add event listeners
	basicBtn.Element().AddEventListener("click", switchToExample("basic"))
	primaryBtn.Element().AddEventListener("click", switchToExample("primary"))
	darkBtn.Element().AddEventListener("click", switchToExample("dark"))
	dropdownBtn.Element().AddEventListener("click", switchToExample("dropdown"))
	infoBtn.Element().AddEventListener("click", switchToExample("info"))
	warningBtn.Element().AddEventListener("click", switchToExample("warning"))
	dangerBtn.Element().AddEventListener("click", switchToExample("danger"))
	complexBtn.Element().AddEventListener("click", switchToExample("complex"))

	// Add button labels
	basicBtn.Append("Basic")
	primaryBtn.Append("Primary")
	darkBtn.Append("Dark")
	dropdownBtn.Append("Dropdown")
	infoBtn.Append("Info")
	warningBtn.Append("Warning")
	dangerBtn.Append("Danger")
	complexBtn.Append("Complex")

	// Create button group for selecting examples
	buttonGroup := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Para(bs.WithClass("mb-2")).Append(
			bs.Span(bs.WithClass("fw-bold")).Append("Select Navbar Style:"),
		),
		bs.ButtonGroup(bs.WithClass("flex-wrap")).Append(
			basicBtn,
			primaryBtn,
			darkBtn,
			dropdownBtn,
			infoBtn,
			warningBtn,
			dangerBtn,
			complexBtn,
		),
	)

	// Basic navbar
	basicContainer.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Basic Navbar"),
		bs.Para(bs.WithClass("text-muted")).Append("A simple light-colored navbar with basic navigation items."),
		bs.NavBar(
			bs.WithColor(bs.LIGHT),
			bs.WithBorder(bs.BOTTOM, bs.SECONDARY),
		).Brand(
			bs.Span().Append("Brand"),
		).Append(
			bs.NavItem("#home", "Home"),
			bs.NavItem("#features", "Features"),
			bs.NavItem("#pricing", "Pricing"),
			bs.NavItem("#about", "About"),
		),
	)

	// Primary navbar
	primaryContainer.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Primary Navbar"),
		bs.Para(bs.WithClass("text-muted")).Append("A navbar with primary color scheme."),
		bs.NavBar(bs.WithColor(bs.PRIMARY)).Brand(
			bs.Span().Append("Company"),
		).Append(
			bs.NavItem("#", "Products"),
			bs.NavItem("#", "Services"),
			bs.NavItem("#", "Contact"),
		),
	)

	// Dark navbar
	darkContainer.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Dark Navbar"),
		bs.Para(bs.WithClass("text-muted")).Append("A dark-themed navbar."),
		bs.NavBar(bs.WithColor(bs.DARK)).Brand(
			bs.Span().Append("DarkBrand"),
		).Append(
			bs.NavItem("#", "Dashboard"),
			bs.NavItem("#", "Reports"),
			bs.NavItem("#", "Settings"),
		),
	)

	// Navbar with dropdowns
	dropdownContainer.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Navbar with Dropdowns"),
		bs.Para(bs.WithClass("text-muted")).Append("A navbar featuring dropdown menus."),
		bs.NavBar(bs.WithColor(bs.LIGHT)).Brand(
			bs.Span().Append("DropBrand"),
		).Append(
			bs.NavItem("#", "Home"),
			bs.NavDropdown("Products").Append(
				bs.NavItem("#product1", "Product 1"),
				bs.NavItem("#product2", "Product 2"),
				bs.NavDivider(),
				bs.NavItem("#all", "All Products"),
			),
			bs.NavDropdown("Services").Append(
				bs.NavItem("#service1", "Consulting"),
				bs.NavItem("#service2", "Development"),
			),
		),
	)

	// Info navbar
	infoContainer.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Info Navbar"),
		bs.Para(bs.WithClass("text-muted")).Append("A navbar with info color scheme."),
		bs.NavBar(bs.WithColor(bs.INFO)).Brand(
			bs.Span().Append("InfoBrand"),
		).Append(
			bs.NavItem("#", "News"),
			bs.NavItem("#", "Updates"),
			bs.NavItem("#", "Archive"),
		),
	)

	// Warning navbar
	warningContainer.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Warning Navbar"),
		bs.Para(bs.WithClass("text-muted")).Append("A navbar with warning color scheme."),
		bs.NavBar(bs.WithColor(bs.WARNING)).Brand(
			bs.Span().Append("WarnBrand"),
		).Append(
			bs.NavItem("#", "Alerts"),
			bs.NavItem("#", "Notifications"),
		),
	)

	// Danger navbar
	dangerContainer.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Danger Navbar"),
		bs.Para(bs.WithClass("text-muted")).Append("A navbar with danger color scheme."),
		bs.NavBar(bs.WithColor(bs.DANGER)).Brand(
			bs.Span().Append("DangerBrand"),
		).Append(
			bs.NavItem("#", "Errors"),
			bs.NavItem("#", "Warnings"),
			bs.NavItem("#", "Critical"),
		),
	)

	// Complex navbar
	complexContainer.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Complex Navbar"),
		bs.Para(bs.WithClass("text-muted")).Append("A comprehensive navbar with multiple dropdowns and items."),
		bs.NavBar(bs.WithColor(bs.SECONDARY)).Brand(
			bs.Span().Append("ComplexApp"),
		).Append(
			bs.NavItem("#dashboard", "Dashboard"),
			bs.NavDropdown("Content").Append(
				bs.NavItem("#posts", "Posts"),
				bs.NavItem("#pages", "Pages"),
				bs.NavItem("#media", "Media"),
				bs.NavDivider(),
				bs.NavItem("#trash", "Trash"),
			),
			bs.NavDropdown("Users").Append(
				bs.NavItem("#all-users", "All Users"),
				bs.NavItem("#add-user", "Add New"),
				bs.NavItem("#profile", "Your Profile"),
			),
			bs.NavDropdown("Settings").Append(
				bs.NavItem("#general", "General"),
				bs.NavItem("#security", "Security"),
				bs.NavItem("#privacy", "Privacy"),
			),
			bs.NavItem("#help", "Help"),
		),
	)

	// Append all parts to main container
	container.Append(
		buttonGroup,
		basicContainer,
		primaryContainer,
		darkContainer,
		dropdownContainer,
		infoContainer,
		warningContainer,
		dangerContainer,
		complexContainer,
	)

	return container
}
