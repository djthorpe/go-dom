package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Breadcrumb struct {
	dom.Element
	nav  dom.Element
	list dom.Element
}

type BreadcrumbItem struct {
	dom.Element
	breadcrumb *Breadcrumb
	link       dom.Element
	active     bool
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Breadcrumb creates a Bootstrap 5 breadcrumb component
func (app *App) Breadcrumb(ariaLabel string) *Breadcrumb {
	// <nav aria-label="breadcrumb">
	nav := app.CreateElement("nav")
	if ariaLabel != "" {
		nav.SetAttribute("aria-label", ariaLabel)
	} else {
		nav.SetAttribute("aria-label", "breadcrumb")
	}

	// <ol class="breadcrumb">
	list := app.CreateElement("ol")
	list.AddClass("breadcrumb")

	nav.AppendChild(list)

	return &Breadcrumb{
		Element: nav,
		nav:     nav,
		list:    list,
	}
}

////////////////////////////////////////////////////////////////////////
// BREADCRUMB METHODS

// AddItem adds a breadcrumb item (link or active text)
func (b *Breadcrumb) AddItem(text string, href string, active bool) *BreadcrumbItem {
	// <li class="breadcrumb-item">
	li := b.Element.OwnerDocument().CreateElement("li")
	li.AddClass("breadcrumb-item")

	var link dom.Element

	if active {
		// Active item (no link)
		// <li class="breadcrumb-item active" aria-current="page">Text</li>
		li.AddClass("active")
		li.SetAttribute("aria-current", "page")
		li.AppendChild(b.Element.OwnerDocument().CreateTextNode(text))
	} else {
		// Link item
		// <li class="breadcrumb-item"><a href="#">Text</a></li>
		link = b.Element.OwnerDocument().CreateElement("a")
		if href != "" {
			link.SetAttribute("href", href)
		} else {
			link.SetAttribute("href", "#")
		}
		link.AppendChild(b.Element.OwnerDocument().CreateTextNode(text))
		li.AppendChild(link)
	}

	b.list.AppendChild(li)

	return &BreadcrumbItem{
		Element:    li,
		breadcrumb: b,
		link:       link,
		active:     active,
	}
}

// SetDivider changes the breadcrumb divider character using CSS
func (b *Breadcrumb) SetDivider(divider string) *Breadcrumb {
	// Use CSS custom property to change divider
	// Set the style attribute directly with the CSS custom property
	b.list.SetAttribute("style", "--bs-breadcrumb-divider: '"+divider+"';")
	return b
}

// AddClass adds a CSS class to the breadcrumb nav
func (b *Breadcrumb) AddClass(className string) *Breadcrumb {
	b.Element.AddClass(className)
	return b
}

// RemoveClass removes a CSS class from the breadcrumb nav
func (b *Breadcrumb) RemoveClass(className string) *Breadcrumb {
	b.Element.RemoveClass(className)
	return b
}

////////////////////////////////////////////////////////////////////////
// BREADCRUMB ITEM METHODS

// SetActive makes this breadcrumb item active (removes link if present)
func (item *BreadcrumbItem) SetActive(active bool) *BreadcrumbItem {
	if active && !item.active {
		// Make active: add class, set aria-current, remove link
		item.Element.AddClass("active")
		item.Element.SetAttribute("aria-current", "page")
		if item.link != nil {
			// Replace link with text node
			text := item.link.TextContent()
			item.Element.RemoveChild(item.link)
			item.Element.AppendChild(
				item.Element.OwnerDocument().CreateTextNode(text),
			)
			item.link = nil
		}
		item.active = true
	} else if !active && item.active {
		// Make inactive: remove class, remove aria-current, add link
		item.Element.RemoveClass("active")
		item.Element.SetAttribute("aria-current", "")
		// Note: Would need to recreate link with href
		item.active = false
	}
	return item
}

// AddEventListener adds an event listener to the breadcrumb item's link
func (item *BreadcrumbItem) AddEventListener(eventType string, callback func(dom.Node)) *BreadcrumbItem {
	if item.link != nil {
		item.link.AddEventListener(eventType, callback)
	}
	return item
}

// AddClass adds a CSS class to the breadcrumb item
func (item *BreadcrumbItem) AddClass(className string) *BreadcrumbItem {
	item.Element.AddClass(className)
	return item
}

// RemoveClass removes a CSS class from the breadcrumb item
func (item *BreadcrumbItem) RemoveClass(className string) *BreadcrumbItem {
	item.Element.RemoveClass(className)
	return item
}

// IsActive returns whether this item is active
func (item *BreadcrumbItem) IsActive() bool {
	return item.active
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (b *Breadcrumb) String() string {
	return "<bs5-breadcrumb>"
}

func (item *BreadcrumbItem) String() string {
	return "<bs5-breadcrumb-item>"
}
