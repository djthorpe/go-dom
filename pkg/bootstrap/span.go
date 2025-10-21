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

type span struct {
	component
}

// Ensure that span implements Component interface
var _ Component = (*span)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Span creates a new span element
func Span(opt ...Opt) *span {
	// Create a span element
	root := dom.GetWindow().Document().CreateElement("SPAN")

	// Apply options
	if opts, err := NewOpts(SpanComponent); err != nil {
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

	return &span{
		component: component{
			name: SpanComponent,
			root: root,
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (span *span) Element() Element {
	return span.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (span *span) Append(children ...any) Component {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to root
		span.root.AppendChild(child.(Node))
	}

	// Return the span for chaining
	return span
}
