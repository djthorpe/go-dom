package main

import (
	"github.com/djthorpe/go-dom/pkg/bs5"
)

func main() {
	// Create a new bootstrap application
	app := bs5.New("Bootstrap App Demo")

	// Add a container for the content
	container := app.CreateElement("div")
	container.AddClass("container")
	container.AddClass("mt-4")

	// Add title with Bootstrap icon
	title := app.H1()
	bootstrapIcon := app.Icon("bootstrap-fill")
	bootstrapIcon.SetColor("text-primary")
	bootstrapIcon.SetSize("fs-1")
	title.Element.AppendChild(bootstrapIcon.Element)
	title.Element.AppendChild(app.CreateTextNode(" Bootstrap 5 Components"))
	container.AppendChild(title.Element)

	// Add description
	desc := app.CreateElement("p")
	desc.AddClass("lead")
	desc.AppendChild(app.CreateTextNode("Explore Bootstrap 5 components built with Go and WASM"))
	container.AppendChild(desc)

	// Create tabs for different component categories
	tabs := app.Tabs("main-tabs")
	tabs.SetStyle(bs5.TabStyleTabs) // Use classic tab style
	tabs.SetJustified(true)         // Make tabs fill the width
	tabs.AddClass("mt-4")

	// Tab 1: Navigation
	navTab := tabs.AddTab("Navigation", true, app)
	navTab.SetPadding("3")
	navTab.AppendChild(AddNavExamples(app))

	// Tab 2: Alerts
	alertsTab := tabs.AddTab("Alerts", false, app)
	alertsTab.SetPadding("3")
	alertsTab.AppendChild(AddAlertExamples(app))

	// Tab 3: Badges
	badgesTab := tabs.AddTab("Badges", false, app)
	badgesTab.SetPadding("3")
	badgesTab.AppendChild(AddBadgeExamples(app))

	// Tab 4: Buttons
	buttonsTab := tabs.AddTab("Buttons", false, app)
	buttonsTab.SetPadding("3")
	buttonsTab.AppendChild(AddButtonExamples(app))

	// Tab 5: Modals
	modalsTab := tabs.AddTab("Modals", false, app)
	modalsTab.SetPadding("3")
	modalsTab.AppendChild(AddModalExample(app))

	// Tab 6: Accordion
	accordionTab := tabs.AddTab("Accordion", false, app)
	accordionTab.SetPadding("3")
	accordionTab.AppendChild(AddAccordionExample(app))

	// Tab 7: Breadcrumbs
	breadcrumbsTab := tabs.AddTab("Breadcrumbs", false, app)
	breadcrumbsTab.SetPadding("3")
	breadcrumbsTab.AppendChild(AddBreadcrumbExamples(app))

	// Tab 8: Cards
	cardsTab := tabs.AddTab("Cards", false, app)
	cardsTab.SetPadding("3")
	cardsTab.AppendChild(AddCardExamples(app))

	// Tab 9: Pagination
	paginationTab := tabs.AddTab("Pagination", false, app)
	paginationTab.SetPadding("3")
	paginationTab.AppendChild(AddPaginationExamples(app))

	// Tab 10: Tables
	tablesTab := tabs.AddTab("Tables", false, app)
	tablesTab.SetPadding("3")
	tablesTab.AppendChild(AddTableExamples(app))

	// Tab 11: Tabs Demo
	tabsDemoTab := tabs.AddTab("Tabs", false, app)
	tabsDemoTab.SetPadding("3")
	tabsDemoTab.AppendChild(AddTabsExamples(app))

	// Tab 12: Progress
	progressTab := tabs.AddTab("Progress", false, app)
	progressTab.SetPadding("3")
	progressTab.AppendChild(AddProgressExamples(app))

	// Tab 13: Offcanvas
	offcanvasTab := tabs.AddTab("Offcanvas", false, app)
	offcanvasTab.SetPadding("3")
	offcanvasTab.AppendChild(AddOffcanvasExamples(app))

	// Tab 14: Toasts
	toastsTab := tabs.AddTab("Toasts", false, app)
	toastsTab.SetPadding("3")
	toastContainer := app.CreateElement("div")
	AddToastExamples(app, toastContainer)
	toastsTab.AppendChild(toastContainer)

	// Tab 15: Forms
	formsTab := tabs.AddTab("Forms", false, app)
	formsTab.SetPadding("3")
	formContainer := app.CreateElement("div")
	AddFormExamples(app, formContainer)
	formsTab.AppendChild(formContainer)

	// Tab 16: Icons
	iconsTab := tabs.AddTab("Icons", false, app)
	iconsTab.SetPadding("3")
	iconContainer := app.CreateElement("div")
	AddIconExamples(app, iconContainer)
	iconsTab.AppendChild(iconContainer)

	// Tab 17: Grid
	gridTab := tabs.AddTab("Grid", false, app)
	gridTab.SetPadding("3")
	gridTab.AppendChild(AddGridExamples(app))

	// Add tabs to container
	container.AppendChild(tabs.Element)

	// Add container to body
	app.Document.Body().AppendChild(container)

	// Keep the Go program running to handle events
	select {}
}
