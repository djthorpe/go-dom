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

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (rule *rule) Element() Element {
	return rule.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Append is provided for consistency but HR elements typically don't have children
func (rule *rule) Append(children ...any) Component {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to root
		rule.root.AppendChild(child.(Node))
	}

	// Return the rule for chaining
	return rule
}
