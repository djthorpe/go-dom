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
