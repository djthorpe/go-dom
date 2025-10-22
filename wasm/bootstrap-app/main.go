package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

func main() {
	// Get the window
	window := dom.GetWindow()

	// Create the bootstrap app
	app := bs.New()

	// Add navbar
	app.Append(NavBar())

	// Create router
	router := bs.Router()

	// Home route
	homeRoute := bs.Route("^$", "^home$")
	homeRoute.Append(
		bs.Heading(1).Append("Bootstrap Components"),
		bs.Para().Append("Welcome to the Bootstrap WASM component library. Use the navigation above to explore different components."),
	)
	router.AddRoute(homeRoute)

	// Navigation route
	navigationRoute := bs.Route("^navs$", "^navbars$")
	navigationRoute.Append(
		bs.Heading(1).Append("Navigation"),
		NavExamples(),
	)
	router.AddRoute(navigationRoute)

	// Page Elements route
	pageElementsRoute := bs.Route("^spans$", "^paragraphs$", "^headings$", "^blockquotes$", "^images$")
	pageElementsRoute.Append(
		bs.Heading(1).Append("Page Elements"),
		ImageExamples(),
	)
	router.AddRoute(pageElementsRoute)

	// Decorations route
	decorationsRoute := bs.Route("^badges$", "^links$", "^buttons$", "^icons$", "^lists$", "^rules$")
	decorationsRoute.Append(
		bs.Heading(1).Append("Decorations"),
		BadgeExamples(),
		ButtonExamples(),
		IconExamples(),
		LinkExamples(),
	)
	router.AddRoute(decorationsRoute)

	// Scaffolding route
	scaffoldingRoute := bs.Route("^containers$", "^cards$", "^routers$", "^grids$", "^breadcrumbs$", "^pagination$", "^accordions$")
	scaffoldingRoute.Append(
		bs.Heading(1).Append("Scaffolding"),
		CardExamples(),
	)
	router.AddRoute(scaffoldingRoute)

	// Windows route
	windowsRoute := bs.Route("^modals$", "^offcanvas$", "^toasts$", "^tooltips$")
	windowsRoute.Append(
		bs.Heading(1).Append("Windows"),
		bs.Para().Append("Modal dialogs, offcanvas panels, toasts, and tooltips."),
	)
	router.AddRoute(windowsRoute)

	// Forms route
	formsRoute := bs.Route("^form$", "^input$", "^textarea$", "^select$", "^radio$", "^checkbox$", "^range$", "^color$", "^date$")
	formsRoute.Append(
		bs.Heading(1).Append("Forms"),
		FormExamples(),
	)
	router.AddRoute(formsRoute)

	// Alignment route
	alignmentRoute := bs.Route("^flex$", "^grid$", "^spacing$", "^sizing$", "^display$")
	alignmentRoute.Append(
		bs.Heading(1).Append("Alignment"),
		bs.Para().Append("Layout and alignment utilities including flex, grid, spacing, and sizing."),
	)
	router.AddRoute(alignmentRoute)

	// Alerts route (for backward compatibility)
	alertsRoute := bs.Route("^alerts$")
	alertsRoute.Append(
		bs.Heading(1).Append("Alerts"),
		AlertExamples(),
	)
	router.AddRoute(alertsRoute)

	// Add router to app
	app.Append(router)

	// Listen for hash changes
	window.AddEventListener("hashchange", func(event Event) {
		hash := window.Location().Hash()
		router.Navigate(hash)
	})

	// Initial navigation based on current hash
	router.Navigate(window.Location().Hash())

	// Run the application
	select {}
}
