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

type rule struct {
	component
}

// Ensure that rule implements Component interface
var _ Component = (*rule)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Rule creates a new horizontal rule (HR) element
func Rule(opt ...Opt) *rule {
	// Create a horizontal rule element
	root := dom.GetWindow().Document().CreateElement("HR")

	// Apply options
	if opts, err := NewOpts(RuleComponent); err != nil {
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

	return &rule{
		component: component{
			name: RuleComponent,
			root: root,
		},
	}
}

// VerticalRule creates a new vertical rule (DIV element with "vr" class)
func VerticalRule(opt ...Opt) *rule {
	// Create a div element for vertical rule
	root := dom.GetWindow().Document().CreateElement("DIV")

	// Apply options with "vr" class
	if opts, err := NewOpts(RuleComponent, WithClass("vr")); err != nil {
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

	return &rule{
		component: component{
			name: RuleComponent,
			root: root,
		},
	}
}
