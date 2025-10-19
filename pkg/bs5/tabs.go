package bs5

import (
	"fmt"

	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

// Tabs represents a Bootstrap tabs navigation and content container
type Tabs struct {
	Element    dom.Element
	nav        *TabNav
	content    *TabContent
	tabCounter int
	id         string
}

// TabNav represents the navigation list for tabs
type TabNav struct {
	Element dom.Element
}

// TabNavItem represents a single tab navigation item
type TabNavItem struct {
	Element dom.Element
	link    dom.Element
	tabID   string
}

// TabContent represents the container for tab panes
type TabContent struct {
	Element dom.Element
}

// TabPane represents a single content pane
type TabPane struct {
	Element dom.Element
	id      string
}

////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Tabs creates a new tabs component with navigation and content areas
func (app *App) Tabs(id string) *Tabs {
	tabs := &Tabs{
		Element:    app.CreateElement("div"),
		tabCounter: 0,
		id:         id,
	}
	tabs.Element.SetAttribute("id", id)

	// Create navigation
	tabs.nav = &TabNav{
		Element: app.CreateElement("ul"),
	}
	tabs.nav.Element.AddClass("nav")
	tabs.nav.Element.AddClass("nav-tabs")
	tabs.nav.Element.SetAttribute("role", "tablist")

	// Create content container
	tabs.content = &TabContent{
		Element: app.CreateElement("div"),
	}
	tabs.content.Element.AddClass("tab-content")
	tabs.content.Element.SetAttribute("id", fmt.Sprintf("%s-content", id))

	// Append nav and content to main container
	tabs.Element.AppendChild(tabs.nav.Element)
	tabs.Element.AppendChild(tabs.content.Element)

	return tabs
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - Tabs

// AddTab adds a new tab with the given title and returns the TabPane for adding content
// If active is true, this tab will be shown by default
func (t *Tabs) AddTab(title string, active bool, app *App) *TabPane {
	t.tabCounter++
	tabID := fmt.Sprintf("%s-tab-%d", t.id, t.tabCounter)
	paneID := fmt.Sprintf("%s-pane-%d", t.id, t.tabCounter)

	// Create nav item
	navItem := &TabNavItem{
		Element: app.CreateElement("li"),
		tabID:   paneID,
	}
	navItem.Element.AddClass("nav-item")
	navItem.Element.SetAttribute("role", "presentation")

	// Create nav link
	navItem.link = app.CreateElement("button")
	navItem.link.AddClass("nav-link")
	if active {
		navItem.link.AddClass("active")
		navItem.link.SetAttribute("aria-selected", "true")
	} else {
		navItem.link.SetAttribute("aria-selected", "false")
	}
	navItem.link.SetAttribute("id", tabID)
	navItem.link.SetAttribute("data-bs-toggle", "tab")
	navItem.link.SetAttribute("data-bs-target", fmt.Sprintf("#%s", paneID))
	navItem.link.SetAttribute("type", "button")
	navItem.link.SetAttribute("role", "tab")
	navItem.link.SetAttribute("aria-controls", paneID)
	navItem.link.AppendChild(app.CreateTextNode(title))

	navItem.Element.AppendChild(navItem.link)
	t.nav.Element.AppendChild(navItem.Element)

	// Create content pane
	pane := &TabPane{
		Element: app.CreateElement("div"),
		id:      paneID,
	}
	pane.Element.AddClass("tab-pane")
	pane.Element.AddClass("fade")
	if active {
		pane.Element.AddClass("show")
		pane.Element.AddClass("active")
	}
	pane.Element.SetAttribute("id", paneID)
	pane.Element.SetAttribute("role", "tabpanel")
	pane.Element.SetAttribute("aria-labelledby", tabID)
	pane.Element.SetAttribute("tabindex", "0")

	t.content.Element.AppendChild(pane.Element)

	return pane
}

// SetStyle sets the style of the tabs navigation
func (t *Tabs) SetStyle(style TabStyle) *Tabs {
	switch style {
	case TabStyleTabs:
		t.nav.Element.RemoveClass("nav-pills")
		t.nav.Element.RemoveClass("nav-underline")
		t.nav.Element.AddClass("nav-tabs")
	case TabStylePills:
		t.nav.Element.RemoveClass("nav-tabs")
		t.nav.Element.RemoveClass("nav-underline")
		t.nav.Element.AddClass("nav-pills")
	case TabStyleUnderline:
		t.nav.Element.RemoveClass("nav-tabs")
		t.nav.Element.RemoveClass("nav-pills")
		t.nav.Element.AddClass("nav-underline")
	}
	return t
}

// SetJustified makes the tabs fill the available width
func (t *Tabs) SetJustified(justified bool) *Tabs {
	if justified {
		t.nav.Element.AddClass("nav-fill")
	} else {
		t.nav.Element.RemoveClass("nav-fill")
	}
	return t
}

// SetVertical arranges tabs vertically (requires custom layout)
func (t *Tabs) SetVertical(vertical bool) *Tabs {
	if vertical {
		t.nav.Element.AddClass("flex-column")
		t.Element.AddClass("d-flex")
	} else {
		t.nav.Element.RemoveClass("flex-column")
		t.Element.RemoveClass("d-flex")
	}
	return t
}

// AddEventListener adds an event listener to the tabs
// Common Bootstrap tab events: show.bs.tab, shown.bs.tab, hide.bs.tab, hidden.bs.tab
func (t *Tabs) AddEventListener(event string, callback func(dom.Node)) *Tabs {
	t.Element.AddEventListener(event, callback)
	return t
}

// AddClass adds a CSS class to the tabs container
func (t *Tabs) AddClass(class string) *Tabs {
	t.Element.AddClass(class)
	return t
}

// RemoveClass removes a CSS class from the tabs container
func (t *Tabs) RemoveClass(class string) *Tabs {
	t.Element.RemoveClass(class)
	return t
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - TabPane

// AppendChild adds a child element to the tab pane
func (p *TabPane) AppendChild(child dom.Element) *TabPane {
	p.Element.AppendChild(child)
	return p
}

// SetActive makes this tab pane active
func (p *TabPane) SetActive(active bool) *TabPane {
	if active {
		p.Element.AddClass("show")
		p.Element.AddClass("active")
	} else {
		p.Element.RemoveClass("show")
		p.Element.RemoveClass("active")
	}
	return p
}

// AddClass adds a CSS class to the tab pane
func (p *TabPane) AddClass(class string) *TabPane {
	p.Element.AddClass(class)
	return p
}

// RemoveClass removes a CSS class from the tab pane
func (p *TabPane) RemoveClass(class string) *TabPane {
	p.Element.RemoveClass(class)
	return p
}

// SetPadding adds padding to the tab pane content
func (p *TabPane) SetPadding(padding string) *TabPane {
	p.Element.AddClass(fmt.Sprintf("p-%s", padding))
	return p
}

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

// TabStyle represents the visual style of tabs
type TabStyle int

const (
	TabStyleTabs TabStyle = iota
	TabStylePills
	TabStyleUnderline
)
