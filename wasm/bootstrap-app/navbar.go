package main

import (
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// NavBar for the main application
func NavBar() Component {
	return bs.NavBar().Append(
		bs.NavItem("#home", true, false, "Home"),
		bs.NavItem("#nav", true, false, "Navigation").Append(
			bs.NavItem("#navs", true, false, "Navs"),
			bs.NavItem("#navbars", true, false, "Navbars"),
		),
		bs.NavItem("#page", true, false, "Page Elements").Append(
			bs.NavItem("#alerts", true, false, "Spans"),
			bs.NavItem("#alerts", true, false, "Paragraphs"),
			bs.NavItem("#alerts", true, false, "Images"),
			bs.NavItem("#alerts", true, false, "Headings"),
			bs.NavItem("#alerts", true, false, "Blockquotes"),
		),
		bs.NavItem("#decorations", true, false, "Decorations").Append(
			bs.NavItem("#badges", true, false, "Badges"),
			bs.NavItem("#badges", true, false, "Links"),
			bs.NavItem("#badges", true, false, "Buttons"),
			bs.NavItem("#badges", true, false, "Icons"),
			bs.NavItem("#badges", true, false, "Lists"),
			bs.NavItem("#badges", true, false, "Rules"),
		),
		bs.NavItem("#decorations", true, false, "Scaffolding").Append(
			bs.NavItem("#badges", true, false, "Containers"),
			bs.NavItem("#badges", true, false, "Cards"),
			bs.NavItem("#badges", true, false, "Routers"),
			bs.NavItem("#badges", true, false, "Grids"),
			bs.NavItem("#badges", true, false, "Breadcrumbs"),
			bs.NavItem("#badges", true, false, "Pagination"),
			bs.NavItem("#badges", true, false, "Accordions"),
		),
		bs.NavItem("#decorations", true, false, "Windows").Append(
			bs.NavItem("#badges", true, false, "Modals"),
			bs.NavItem("#badges", true, false, "Offcanvas"),
			bs.NavItem("#badges", true, false, "Toasts"),
			bs.NavItem("#badges", true, false, "Tooltips"),
		),
		bs.NavItem("#decorations", true, false, "Forms").Append(
			bs.NavItem("#badges", true, false, "Form"),
			bs.NavItem("#badges", true, false, "Input"),
			bs.NavItem("#badges", true, false, "Textarea"),
			bs.NavItem("#badges", true, false, "Select"),
			bs.NavItem("#badges", true, false, "Radio"),
			bs.NavItem("#badges", true, false, "Checkbox"),
			bs.NavItem("#badges", true, false, "Range"),
			bs.NavItem("#badges", true, false, "Color"),
			bs.NavItem("#badges", true, false, "Date"),
		),
		bs.NavItem("#decorations", true, false, "Alignment").Append(
			bs.NavItem("#badges", true, false, "Flex"),
			bs.NavItem("#badges", true, false, "Grid"),
			bs.NavItem("#badges", true, false, "Spacing"),
			bs.NavItem("#badges", true, false, "Sizing"),
			bs.NavItem("#badges", true, false, "Display"),
		),
	)
}
