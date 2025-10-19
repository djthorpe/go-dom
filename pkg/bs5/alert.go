package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Alert struct {
	dom.Element
	color      ColorVariant
	dismissBtn dom.Element
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Alert creates a Bootstrap 5 alert component
func (app *App) Alert(color ColorVariant, children ...dom.Node) *Alert {
	// Create alert div
	alert := app.CreateElement("div")
	alert.AddClass("alert")
	alert.AddClass("alert-" + string(color))
	alert.SetAttribute("role", "alert")

	// Add children
	for _, child := range children {
		alert.AppendChild(child)
	}

	return &Alert{
		Element: alert,
		color:   color,
	}
}

////////////////////////////////////////////////////////////////////////
// METHODS

// MakeDismissible adds a close button to the alert
func (a *Alert) MakeDismissible() *Alert {
	a.Element.AddClass("alert-dismissible")
	a.Element.AddClass("fade")
	a.Element.AddClass("show")

	// Create close button
	closeBtn := a.Element.OwnerDocument().CreateElement("button")
	closeBtn.AddClass("btn-close")
	closeBtn.SetAttribute("type", "button")
	closeBtn.SetAttribute("data-bs-dismiss", "alert")
	closeBtn.SetAttribute("aria-label", "Close")

	a.dismissBtn = closeBtn
	a.Element.AppendChild(closeBtn)

	return a
}

// AddClass adds a CSS class to the alert
func (a *Alert) AddClass(className string) *Alert {
	a.Element.AddClass(className)
	return a
}

// RemoveClass removes a CSS class from the alert
func (a *Alert) RemoveClass(className string) *Alert {
	a.Element.RemoveClass(className)
	return a
}

// Color returns the alert's color variant
func (a *Alert) Color() ColorVariant {
	return a.color
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (a *Alert) String() string {
	return "<bs5-alert>"
}
