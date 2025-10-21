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
	root Element
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
	if opts, err := NewOpts(WithClass("container")); err != nil {
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
		root: root,
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (container *container) Element() Element {
	return container.root
}
