package bootstrap

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type link struct {
	component
}

// Ensure that link implements Component interface
var _ Component = (*link)(nil)

////////////////////////////////////////////////////////////////////////////////
// CONSTRUCTOR

// Link creates a Bootstrap link (anchor) element with optional href and content
// The href parameter sets the link destination, and options can be used to style the link
//
// Example:
//
//	Link("/home", WithColor(PRIMARY)).Append("Home Page")
//	Link("#section").Append("Jump to Section")
//	Link("https://example.com", WithClass("link-offset-2")).Append("External Link")
func Link(href string, opt ...Opt) *link {
	// Create an anchor element
	root := dom.GetWindow().Document().CreateElement("A")

	// Set href attribute
	if href != "" {
		root.SetAttribute("href", href)
	}

	// Apply options
	if opts, err := NewOpts(LinkComponent); err != nil {
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

		// Return the link component
		return &link{
			component: component{
				name: LinkComponent,
				root: root,
			},
		}
	}
}
