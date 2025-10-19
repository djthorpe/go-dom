package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Offcanvas struct {
	dom.Element
	header dom.Element
	body   dom.Element
	id     string
}

type OffcanvasPlacement string

////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	OffcanvasPlacementStart  OffcanvasPlacement = "start"
	OffcanvasPlacementEnd    OffcanvasPlacement = "end"
	OffcanvasPlacementTop    OffcanvasPlacement = "top"
	OffcanvasPlacementBottom OffcanvasPlacement = "bottom"
)

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Offcanvas creates a Bootstrap 5 offcanvas component
func (app *App) Offcanvas(id string, placement OffcanvasPlacement) *Offcanvas {
	// Create main offcanvas element
	offcanvas := app.CreateElement("div")
	offcanvas.AddClass("offcanvas")
	offcanvas.AddClass("offcanvas-" + string(placement))
	offcanvas.SetAttribute("tabindex", "-1")
	offcanvas.SetAttribute("id", id)
	offcanvas.SetAttribute("aria-labelledby", id+"Label")

	// Create header
	header := app.CreateElement("div")
	header.AddClass("offcanvas-header")

	// Create body
	body := app.CreateElement("div")
	body.AddClass("offcanvas-body")

	offcanvas.AppendChild(header)
	offcanvas.AppendChild(body)

	return &Offcanvas{
		Element: offcanvas,
		header:  header,
		body:    body,
		id:      id,
	}
}

////////////////////////////////////////////////////////////////////////
// METHODS

// AddTitle adds a title to the offcanvas header
func (o *Offcanvas) AddTitle(title string, app *App) *Offcanvas {
	titleEl := app.CreateElement("h5")
	titleEl.AddClass("offcanvas-title")
	titleEl.SetAttribute("id", o.id+"Label")
	titleEl.AppendChild(app.CreateTextNode(title))
	o.header.AppendChild(titleEl)
	return o
}

// AddCloseButton adds a close button to the offcanvas header
func (o *Offcanvas) AddCloseButton(app *App) *Offcanvas {
	closeBtn := app.CreateElement("button")
	closeBtn.SetAttribute("type", "button")
	closeBtn.AddClass("btn-close")
	closeBtn.SetAttribute("data-bs-dismiss", "offcanvas")
	closeBtn.SetAttribute("aria-label", "Close")
	o.header.AppendChild(closeBtn)
	return o
}

// Body returns the offcanvas body element for adding content
func (o *Offcanvas) Body() dom.Element {
	return o.body
}

// Header returns the offcanvas header element for adding custom content
func (o *Offcanvas) Header() dom.Element {
	return o.header
}

// SetBodyScroll enables/disables body scrolling when offcanvas is open
func (o *Offcanvas) SetBodyScroll(scroll bool) *Offcanvas {
	if scroll {
		o.Element.SetAttribute("data-bs-scroll", "true")
	} else {
		o.Element.SetAttribute("data-bs-scroll", "false")
	}
	return o
}

// SetBackdrop sets the backdrop behavior
// backdrop: true (default), false (no backdrop), or "static" (backdrop doesn't close offcanvas)
func (o *Offcanvas) SetBackdrop(backdrop string) *Offcanvas {
	if backdrop == "false" {
		o.Element.SetAttribute("data-bs-backdrop", "false")
	} else if backdrop == "static" {
		o.Element.SetAttribute("data-bs-backdrop", "static")
	} else {
		o.Element.SetAttribute("data-bs-backdrop", "true")
	}
	return o
}

// SetKeyboard sets whether the offcanvas can be closed with the escape key
func (o *Offcanvas) SetKeyboard(keyboard bool) *Offcanvas {
	if keyboard {
		o.Element.SetAttribute("data-bs-keyboard", "true")
	} else {
		o.Element.SetAttribute("data-bs-keyboard", "false")
	}
	return o
}

// SetDark applies dark theme to the offcanvas
func (o *Offcanvas) SetDark(dark bool) *Offcanvas {
	if dark {
		o.Element.SetAttribute("data-bs-theme", "dark")
	} else {
		o.Element.RemoveClass("text-bg-dark")
		o.Element.SetAttribute("data-bs-theme", "light")
	}
	return o
}

// AddClass adds a CSS class to the offcanvas
func (o *Offcanvas) AddClass(className string) *Offcanvas {
	o.Element.AddClass(className)
	return o
}

// RemoveClass removes a CSS class from the offcanvas
func (o *Offcanvas) RemoveClass(className string) *Offcanvas {
	o.Element.RemoveClass(className)
	return o
}

// AddEventListener adds an event listener to the offcanvas
func (o *Offcanvas) AddEventListener(eventType string, callback func(dom.Node)) *Offcanvas {
	o.Element.AddEventListener(eventType, callback)
	return o
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (o *Offcanvas) String() string {
	return "<bs5-offcanvas>"
}
