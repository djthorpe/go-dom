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

type icon struct {
	component
}

// Ensure that icon implements Component interface
var _ Component = (*icon)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Icon creates a new Bootstrap Icon element using the icon font.
// The iconName parameter should be the name of the icon without the "bi-" prefix
// (e.g., "heart-fill", "alarm", "bootstrap").
// See https://icons.getbootstrap.com/ for available icons.
//
// Example:
//
//	Icon("heart-fill", WithColor(DANGER), WithClass("fs-3"))
func Icon(iconName string, opt ...Opt) *icon {
	// Create an i element for the icon
	root := dom.GetWindow().Document().CreateElement("I")

	// Add the base Bootstrap Icons class and the specific icon class
	iconClass := "bi-" + iconName
	opt = append([]Opt{WithClass("bi"), WithClass(iconClass)}, opt...)

	// Apply options
	if opts, err := NewOpts(IconComponent); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes (for aria-label, role, etc.)
		for key, value := range opts.attributes {
			root.SetAttribute(key, value)
		}

		// Return the icon component
		return &icon{
			component: component{
				name: IconComponent,
				root: root,
			},
		}
	}
}
