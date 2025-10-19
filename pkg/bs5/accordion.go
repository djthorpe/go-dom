package bs5

import (
	"fmt"

	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Accordion struct {
	dom.Element
	id         string
	items      []*AccordionItem
	itemCount  int
	alwaysOpen bool
}

type AccordionItem struct {
	dom.Element
	accordion *Accordion
	header    *AccordionHeader
	collapse  *AccordionCollapse
	body      *AccordionBody
	id        string
	index     int
}

type AccordionHeader struct {
	dom.Element
	button dom.Element
	item   *AccordionItem
}

type AccordionCollapse struct {
	dom.Element
	item *AccordionItem
}

type AccordionBody struct {
	dom.Element
	item *AccordionItem
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Accordion creates a Bootstrap 5 accordion component
func (app *App) Accordion(id string) *Accordion {
	// <div class="accordion" id="accordionExample">
	accordion := app.CreateElement("div")
	accordion.AddClass("accordion")
	accordion.SetAttribute("id", id)

	return &Accordion{
		Element:   accordion,
		id:        id,
		items:     make([]*AccordionItem, 0),
		itemCount: 0,
	}
}

////////////////////////////////////////////////////////////////////////
// ACCORDION METHODS

// SetFlush removes the default background-color, borders, and rounded corners
func (a *Accordion) SetFlush(flush bool) *Accordion {
	if flush {
		a.Element.AddClass("accordion-flush")
	} else {
		a.Element.RemoveClass("accordion-flush")
	}
	return a
}

// SetAlwaysOpen allows multiple items to remain open at the same time
func (a *Accordion) SetAlwaysOpen(alwaysOpen bool) *Accordion {
	a.alwaysOpen = alwaysOpen
	// Update data-bs-parent on existing items
	for _, item := range a.items {
		if item.collapse != nil {
			if alwaysOpen {
				// Remove parent reference to allow multiple open items
				item.collapse.Element.SetAttribute("data-bs-parent", "")
			} else {
				// Set parent reference to enforce single open item
				item.collapse.Element.SetAttribute("data-bs-parent", "#"+a.id)
			}
		}
	}
	return a
}

// AddItem creates and adds a new accordion item
func (a *Accordion) AddItem(title string, expanded bool) *AccordionItem {
	// Generate unique ID for this item
	itemID := fmt.Sprintf("%s-item-%d", a.id, a.itemCount)
	collapseID := fmt.Sprintf("%s-collapse-%d", a.id, a.itemCount)

	// <div class="accordion-item">
	itemEl := a.Element.OwnerDocument().CreateElement("div")
	itemEl.AddClass("accordion-item")

	item := &AccordionItem{
		Element:   itemEl,
		accordion: a,
		id:        itemID,
		index:     a.itemCount,
	}

	// Create header
	item.header = item.createHeader(title, collapseID, expanded)
	itemEl.AppendChild(item.header.Element)

	// Create collapse container
	item.collapse = item.createCollapse(collapseID, expanded)
	itemEl.AppendChild(item.collapse.Element)

	// Add to accordion
	a.Element.AppendChild(itemEl)
	a.items = append(a.items, item)
	a.itemCount++

	return item
}

// AddClass adds a CSS class to the accordion
func (a *Accordion) AddClass(className string) *Accordion {
	a.Element.AddClass(className)
	return a
}

// RemoveClass removes a CSS class from the accordion
func (a *Accordion) RemoveClass(className string) *Accordion {
	a.Element.RemoveClass(className)
	return a
}

// ID returns the accordion's ID
func (a *Accordion) ID() string {
	return a.id
}

////////////////////////////////////////////////////////////////////////
// ACCORDION ITEM METHODS

// createHeader creates the accordion header with button
func (item *AccordionItem) createHeader(title, collapseID string, expanded bool) *AccordionHeader {
	// <h2 class="accordion-header">
	header := item.Element.OwnerDocument().CreateElement("h2")
	header.AddClass("accordion-header")

	// <button class="accordion-button" type="button" ...>
	button := item.Element.OwnerDocument().CreateElement("button")
	button.AddClass("accordion-button")
	if !expanded {
		button.AddClass("collapsed")
	}
	button.SetAttribute("type", "button")
	button.SetAttribute("data-bs-toggle", "collapse")
	button.SetAttribute("data-bs-target", "#"+collapseID)
	button.SetAttribute("aria-expanded", fmt.Sprintf("%t", expanded))
	button.SetAttribute("aria-controls", collapseID)
	button.AppendChild(item.Element.OwnerDocument().CreateTextNode(title))

	header.AppendChild(button)

	return &AccordionHeader{
		Element: header,
		button:  button,
		item:    item,
	}
}

// createCollapse creates the collapsible content container
func (item *AccordionItem) createCollapse(collapseID string, expanded bool) *AccordionCollapse {
	// <div id="collapseOne" class="accordion-collapse collapse show" data-bs-parent="#accordionExample">
	collapse := item.Element.OwnerDocument().CreateElement("div")
	collapse.SetAttribute("id", collapseID)
	collapse.AddClass("accordion-collapse")
	collapse.AddClass("collapse")
	if expanded {
		collapse.AddClass("show")
	}
	if !item.accordion.alwaysOpen {
		collapse.SetAttribute("data-bs-parent", "#"+item.accordion.id)
	}

	return &AccordionCollapse{
		Element: collapse,
		item:    item,
	}
}

// Body creates or returns the accordion body
func (item *AccordionItem) Body(children ...dom.Node) *AccordionBody {
	if item.body == nil {
		// <div class="accordion-body">
		body := item.Element.OwnerDocument().CreateElement("div")
		body.AddClass("accordion-body")

		// Add children
		for _, child := range children {
			body.AppendChild(child)
		}

		item.collapse.Element.AppendChild(body)
		item.body = &AccordionBody{
			Element: body,
			item:    item,
		}
	}
	return item.body
}

// SetExpanded sets whether this item is expanded
func (item *AccordionItem) SetExpanded(expanded bool) *AccordionItem {
	if expanded {
		item.collapse.Element.AddClass("show")
		item.header.button.RemoveClass("collapsed")
		item.header.button.SetAttribute("aria-expanded", "true")
	} else {
		item.collapse.Element.RemoveClass("show")
		item.header.button.AddClass("collapsed")
		item.header.button.SetAttribute("aria-expanded", "false")
	}
	return item
}

// AddClass adds a CSS class to the accordion item
func (item *AccordionItem) AddClass(className string) *AccordionItem {
	item.Element.AddClass(className)
	return item
}

// RemoveClass removes a CSS class from the accordion item
func (item *AccordionItem) RemoveClass(className string) *AccordionItem {
	item.Element.RemoveClass(className)
	return item
}

////////////////////////////////////////////////////////////////////////
// ACCORDION HEADER METHODS

// SetTitle updates the button text
func (h *AccordionHeader) SetTitle(title string) *AccordionHeader {
	// Clear existing text and set new title
	h.button.SetAttribute("textContent", title)
	return h
}

// AddClass adds a CSS class to the accordion header
func (h *AccordionHeader) AddClass(className string) *AccordionHeader {
	h.Element.AddClass(className)
	return h
}

// RemoveClass removes a CSS class from the accordion header
func (h *AccordionHeader) RemoveClass(className string) *AccordionHeader {
	h.Element.RemoveClass(className)
	return h
}

////////////////////////////////////////////////////////////////////////
// ACCORDION BODY METHODS

// AddClass adds a CSS class to the accordion body
func (b *AccordionBody) AddClass(className string) *AccordionBody {
	b.Element.AddClass(className)
	return b
}

// RemoveClass removes a CSS class from the accordion body
func (b *AccordionBody) RemoveClass(className string) *AccordionBody {
	b.Element.RemoveClass(className)
	return b
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (a *Accordion) String() string {
	return "<bs5-accordion>"
}

func (item *AccordionItem) String() string {
	return "<bs5-accordion-item>"
}

func (h *AccordionHeader) String() string {
	return "<bs5-accordion-header>"
}

func (c *AccordionCollapse) String() string {
	return "<bs5-accordion-collapse>"
}

func (b *AccordionBody) String() string {
	return "<bs5-accordion-body>"
}
