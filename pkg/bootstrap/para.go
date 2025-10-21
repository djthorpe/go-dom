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

type para struct {
	component
}

// Ensure that para implements Component interface
var _ Component = (*para)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Para creates a new paragraph (P) element
func Para(opt ...Opt) *para {
	// Create a paragraph element
	root := dom.GetWindow().Document().CreateElement("P")

	// Apply options
	if opts, err := NewOpts(ParaComponent); err != nil {
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

	return &para{
		component: component{
			name: ParaComponent,
			root: root,
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (para *para) Element() Element {
	return para.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (para *para) Append(children ...any) Component {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to root
		para.root.AppendChild(child.(Node))
	}

	// Return the para for chaining
	return para
}
