package bootstrap

import (

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type name string

type component struct {
	name name
	root Element
}

// Ensure that component implements Component interface
var _ Component = (*component)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ContainerComponent name = "container"
	HeadingComponent   name = "heading"
	BadgeComponent     name = "badge"
)

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (component *component) Name() string {
	return string(component.name)
}

func (component *component) Element() Element {
	return component.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (component *component) Append(children ...any) Component {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to root
		component.root.AppendChild(child.(Node))
	}

	// Return the component for chaining
	return component
}
