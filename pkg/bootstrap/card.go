package bootstrap

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type card struct {
	component
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Card creates a new bootstrap card (div element with class="card")
// The card automatically creates a card-body div where content is appended.
// Use Heading() to add a card-header and Footer() to add a card-footer.
// Content added via Append() goes to the card-body.
//
// Example:
//
//	Card().Heading("Featured").Append("Card content")
//	Card(WithColor(PRIMARY)).Heading("Header").Append("Content").Footer("Footer")
func Card(opt ...Opt) *card {
	// Create a card div element
	root := dom.GetWindow().Document().CreateElement("DIV")

	// Apply options
	if opts, err := NewOpts(CardComponent, WithClass("card")); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			root.SetAttribute(key, value)
		}
	}

	// Create card-body div
	body := dom.GetWindow().Document().CreateElement("DIV")
	body.SetAttribute("class", "card-body")
	root.AppendChild(body)

	return &card{
		component: component{
			name: CardComponent,
			root: root,
			body: body,
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Heading adds a card-header div at the beginning of the card (root element).
// It creates a <div class="card-header"> element and accepts string, Component, or Element children.
// Returns *card to allow method chaining with Heading(), Footer(), and Append().
// If called multiple times, only the last call is used.
func (c *card) Heading(children ...any) *card {
	// Remove existing header if present
	childNodes := c.root.ChildNodes()
	for i, node := range childNodes {
		if elem, ok := node.(Element); ok {
			if elem.GetAttribute("class") == "card-header" {
				c.root.RemoveChild(node)
				break
			}
		}
		// Only check first child
		if i > 0 {
			break
		}
	}

	// Create header element
	header := dom.GetWindow().Document().CreateElement("DIV")
	header.SetAttribute("class", "card-header")

	// Append children to header using Append logic
	for _, child := range children {
		if component, ok := child.(Component); ok {
			header.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			header.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			header.AppendChild(node)
		}
	}

	// Insert at the beginning of root
	childNodes = c.root.ChildNodes()
	if len(childNodes) > 0 {
		c.root.InsertBefore(header, childNodes[0])
	} else {
		c.root.AppendChild(header)
	}

	return c
}

// Footer adds a card-footer div at the end of the card (root element).
// It creates a <div class="card-footer"> element and accepts string, Component, or Element children.
// Returns *card to allow method chaining with Heading(), Footer(), and Append().
// If called multiple times, only the last call is used.
func (c *card) Footer(children ...any) *card {
	// Remove existing footer if present
	childNodes := c.root.ChildNodes()
	for i := len(childNodes) - 1; i >= 0; i-- {
		node := childNodes[i]
		if elem, ok := node.(Element); ok {
			if elem.GetAttribute("class") == "card-footer" {
				c.root.RemoveChild(node)
				break
			}
		}
		// Only check last child
		if i < len(childNodes)-1 {
			break
		}
	}

	// Create footer element
	footer := dom.GetWindow().Document().CreateElement("DIV")
	footer.SetAttribute("class", "card-footer")

	// Append children to footer using Append logic
	for _, child := range children {
		if component, ok := child.(Component); ok {
			footer.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			footer.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			footer.AppendChild(node)
		}
	}

	// Append at the end of root
	c.root.AppendChild(footer)

	return c
}
