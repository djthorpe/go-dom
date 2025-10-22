package main

import (
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// NavBar for the main application
func NavBar() Component {
	return bs.NavBar(
		bs.WithColor(bs.PRIMARY),
		bs.WithBorder(bs.BOTTOM, bs.BLACK),
	).Brand(
		bs.Icon("bootstrap-fill", bs.WithMargin(bs.START|bs.END, 2)),
	).Append(
		bs.NavDropdown("Navigation").Append(
			bs.NavItem("#navs", "Navs"),
			bs.NavItem("#navbars", "Navbars"),
		),
		bs.NavDropdown("Page Elements").Append(
			bs.NavItem("#spans", "Spans"),
			bs.NavItem("#paragraphs", "Paragraphs"),
			bs.NavItem("#headings", "Headings"),
			bs.NavItem("#blockquotes", "Blockquotes"),
			bs.NavDivider(),
			bs.NavItem("#images", "Images"),
		),
		bs.NavDropdown("Decorations").Append(
			bs.NavItem("#badges", "Badges"),
			bs.NavItem("#links", "Links"),
			bs.NavItem("#buttons", "Buttons"),
			bs.NavItem("#icons", "Icons"),
			bs.NavItem("#lists", "Lists"),
			bs.NavItem("#rules", "Rules"),
		),
		bs.NavDropdown("Scaffolding").Append(
			bs.NavItem("#containers", "Containers"),
			bs.NavItem("#cards", "Cards"),
			bs.NavItem("#routers", "Routers"),
			bs.NavItem("#grids", "Grids"),
			bs.NavItem("#breadcrumbs", "Breadcrumbs"),
			bs.NavItem("#pagination", "Pagination"),
			bs.NavItem("#accordions", "Accordions"),
		),
		bs.NavDropdown("Windows").Append(
			bs.NavItem("#modals", "Modals"),
			bs.NavItem("#offcanvas", "Offcanvas"),
			bs.NavItem("#toasts", "Toasts"),
			bs.NavItem("#tooltips", "Tooltips"),
		),
		bs.NavDropdown("Forms").Append(
			bs.NavItem("#form", "Form"),
			bs.NavItem("#input", "Input"),
			bs.NavItem("#textarea", "Textarea"),
			bs.NavItem("#select", "Select"),
			bs.NavItem("#radio", "Radio"),
			bs.NavItem("#checkbox", "Checkbox"),
			bs.NavItem("#range", "Range"),
			bs.NavItem("#color", "Color"),
			bs.NavItem("#date", "Date"),
		),
		bs.NavDropdown("Alignment").Append(
			bs.NavItem("#flex", "Flex"),
			bs.NavItem("#grid", "Grid"),
			bs.NavItem("#spacing", "Spacing"),
			bs.NavItem("#sizing", "Sizing"),
			bs.NavItem("#display", "Display"),
		),
	)
}
