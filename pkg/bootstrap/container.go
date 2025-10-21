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

type container struct {
	component
}

// Ensure that container implements Component interface
var _ Component = (*container)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new bootstrap container
func Container(opt ...Opt) *container {
	// Create a container
	root := dom.GetWindow().Document().CreateElement("DIV")

	// Apply options
	if opts, err := NewOpts(ContainerComponent, WithClass("container")); err != nil {
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

	return &container{
		component: component{
			name: ContainerComponent,
			root: root,
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (container *container) Element() Element {
	return container.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (container *container) Append(children ...any) Component {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to root
		container.root.AppendChild(child.(Node))
	}

	// Return the container for chaining
	return container
}
