package bs5

import (
	"fmt"

	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Pagination struct {
	dom.Element
	nav  dom.Element
	list dom.Element
}

type PaginationItem struct {
	dom.Element
	pagination *Pagination
	link       dom.Element
	active     bool
	disabled   bool
}

type PaginationSize string

////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	PaginationSizeSmall  PaginationSize = "sm"
	PaginationSizeMedium PaginationSize = ""
	PaginationSizeLarge  PaginationSize = "lg"
)

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Pagination creates a Bootstrap 5 pagination component
func (app *App) Pagination(ariaLabel string) *Pagination {
	// <nav aria-label="Page navigation">
	nav := app.CreateElement("nav")
	if ariaLabel != "" {
		nav.SetAttribute("aria-label", ariaLabel)
	} else {
		nav.SetAttribute("aria-label", "Page navigation")
	}

	// <ul class="pagination">
	list := app.CreateElement("ul")
	list.AddClass("pagination")

	nav.AppendChild(list)

	return &Pagination{
		Element: nav,
		nav:     nav,
		list:    list,
	}
}

////////////////////////////////////////////////////////////////////////
// PAGINATION METHODS

// AddItem adds a pagination item (link or active page)
func (p *Pagination) AddItem(text string, href string, active bool, disabled bool) *PaginationItem {
	// <li class="page-item">
	li := p.Element.OwnerDocument().CreateElement("li")
	li.AddClass("page-item")

	if active {
		li.AddClass("active")
		li.SetAttribute("aria-current", "page")
	}
	if disabled {
		li.AddClass("disabled")
	}

	// <a class="page-link" href="#">
	link := p.Element.OwnerDocument().CreateElement("a")
	link.AddClass("page-link")
	if href != "" {
		link.SetAttribute("href", href)
	} else {
		link.SetAttribute("href", "#")
	}
	if disabled {
		link.SetAttribute("tabindex", "-1")
		link.SetAttribute("aria-disabled", "true")
	}
	link.AppendChild(p.Element.OwnerDocument().CreateTextNode(text))

	li.AppendChild(link)
	p.list.AppendChild(li)

	return &PaginationItem{
		Element:    li,
		pagination: p,
		link:       link,
		active:     active,
		disabled:   disabled,
	}
}

// AddItemWithNode adds a pagination item with custom content (can include icons)
func (p *Pagination) AddItemWithNode(content dom.Node, href string, active bool, disabled bool) *PaginationItem {
	// <li class="page-item">
	li := p.Element.OwnerDocument().CreateElement("li")
	li.AddClass("page-item")

	if active {
		li.AddClass("active")
		li.SetAttribute("aria-current", "page")
	}
	if disabled {
		li.AddClass("disabled")
	}

	// <a class="page-link" href="#">
	link := p.Element.OwnerDocument().CreateElement("a")
	link.AddClass("page-link")
	if href != "" {
		link.SetAttribute("href", href)
	} else {
		link.SetAttribute("href", "#")
	}
	if disabled {
		link.SetAttribute("tabindex", "-1")
		link.SetAttribute("aria-disabled", "true")
	}
	link.AppendChild(content)

	li.AppendChild(link)
	p.list.AppendChild(li)

	return &PaginationItem{
		Element:    li,
		pagination: p,
		link:       link,
		active:     active,
		disabled:   disabled,
	}
}

// AddPrevious adds a "Previous" pagination item
func (p *Pagination) AddPrevious(href string, disabled bool) *PaginationItem {
	item := p.AddItem("Previous", href, false, disabled)
	return item
}

// AddNext adds a "Next" pagination item
func (p *Pagination) AddNext(href string, disabled bool) *PaginationItem {
	item := p.AddItem("Next", href, false, disabled)
	return item
}

// AddPage adds a numbered page item
func (p *Pagination) AddPage(page int, href string, active bool) *PaginationItem {
	return p.AddItem(fmt.Sprintf("%d", page), href, active, false)
}

// AddEllipsis adds an ellipsis (disabled item with "...")
func (p *Pagination) AddEllipsis() *PaginationItem {
	return p.AddItem("...", "#", false, true)
}

// SetSize sets the pagination size
func (p *Pagination) SetSize(size PaginationSize) *Pagination {
	// Remove existing size classes
	p.list.RemoveClass("pagination-sm")
	p.list.RemoveClass("pagination-lg")

	// Add new size class if not medium (default)
	if size != PaginationSizeMedium {
		p.list.AddClass("pagination-" + string(size))
	}
	return p
}

// SetAlignment sets the pagination alignment
func (p *Pagination) SetAlignment(alignment string) *Pagination {
	// Remove existing alignment classes
	p.list.RemoveClass("justify-content-start")
	p.list.RemoveClass("justify-content-center")
	p.list.RemoveClass("justify-content-end")

	// Add new alignment class
	if alignment != "" && alignment != "start" {
		p.list.AddClass("justify-content-" + alignment)
	}
	return p
}

// AddClass adds a CSS class to the pagination nav
func (p *Pagination) AddClass(className string) *Pagination {
	p.Element.AddClass(className)
	return p
}

// RemoveClass removes a CSS class from the pagination nav
func (p *Pagination) RemoveClass(className string) *Pagination {
	p.Element.RemoveClass(className)
	return p
}

////////////////////////////////////////////////////////////////////////
// PAGINATION ITEM METHODS

// SetActive makes this pagination item active
func (item *PaginationItem) SetActive(active bool) *PaginationItem {
	if active && !item.active {
		item.Element.AddClass("active")
		item.Element.SetAttribute("aria-current", "page")
		item.active = true
	} else if !active && item.active {
		item.Element.RemoveClass("active")
		item.Element.SetAttribute("aria-current", "")
		item.active = false
	}
	return item
}

// SetDisabled makes this pagination item disabled
func (item *PaginationItem) SetDisabled(disabled bool) *PaginationItem {
	if disabled && !item.disabled {
		item.Element.AddClass("disabled")
		item.link.SetAttribute("tabindex", "-1")
		item.link.SetAttribute("aria-disabled", "true")
		item.disabled = true
	} else if !disabled && item.disabled {
		item.Element.RemoveClass("disabled")
		item.link.SetAttribute("tabindex", "")
		item.link.SetAttribute("aria-disabled", "false")
		item.disabled = false
	}
	return item
}

// AddEventListener adds an event listener to the pagination item's link
func (item *PaginationItem) AddEventListener(eventType string, callback func(dom.Node)) *PaginationItem {
	if item.link != nil {
		item.link.AddEventListener(eventType, callback)
	}
	return item
}

// AddClass adds a CSS class to the pagination item
func (item *PaginationItem) AddClass(className string) *PaginationItem {
	item.Element.AddClass(className)
	return item
}

// RemoveClass removes a CSS class from the pagination item
func (item *PaginationItem) RemoveClass(className string) *PaginationItem {
	item.Element.RemoveClass(className)
	return item
}

// IsActive returns whether this item is active
func (item *PaginationItem) IsActive() bool {
	return item.active
}

// IsDisabled returns whether this item is disabled
func (item *PaginationItem) IsDisabled() bool {
	return item.disabled
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (p *Pagination) String() string {
	return "<bs5-pagination>"
}

func (item *PaginationItem) String() string {
	return "<bs5-pagination-item>"
}
