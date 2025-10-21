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
	body Element // Where content is appended; usually same as root
}

// Ensure that component implements Component interface
var _ Component = (*component)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ContainerComponent   name = "container"
	HeadingComponent     name = "heading"
	BadgeComponent       name = "badge"
	AlertComponent       name = "alert"
	SpanComponent        name = "span"
	ParaComponent        name = "para"
	RuleComponent        name = "rule"
	ButtonComponent      name = "button"
	ButtonGroupComponent name = "button-group"
	IconComponent        name = "icon"
	LinkComponent        name = "link"
	CardComponent        name = "card"
	ImageComponent       name = "image"
	NavComponent         name = "nav"
	ScrollspyComponent   name = "scrollspy"
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
	// Append Component, Element or string children to the body element (or root if no body)
	target := component.body
	if target == nil {
		target = component.root
	}

	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to target
		target.AppendChild(child.(Node))
	}

	// Return the component for chaining
	return component
}
