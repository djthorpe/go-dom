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

type badge struct {
	component
}

// Ensure that badge implements Component interface
var _ Component = (*badge)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Badge creates a new bootstrap badge (span element)
// Default badge uses "badge" class. Use WithBackground to set color variant.
func Badge(opt ...Opt) *badge {
	// Create a badge span element
	root := dom.GetWindow().Document().CreateElement("SPAN")

	// Apply options
	if opts, err := NewOpts(BadgeComponent, WithClass("badge")); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}
	}

	return &badge{
		component: component{
			name: BadgeComponent,
			root: root,
		},
	}
}

// PillBadge creates a new bootstrap rounded-pill badge (span element)
func PillBadge(opt ...Opt) *badge {
	// Append the "rounded-pill" class to options
	opt = append(opt, WithClass("rounded-pill"))
	return Badge(opt...)
}
