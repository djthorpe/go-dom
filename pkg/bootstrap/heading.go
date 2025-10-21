package bootstrap

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type heading struct {
	component
	level int
}

// Ensure that heading implements Component interface
var _ Component = (*heading)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Heading creates a new heading element with the specified level (1-6)
func Heading(level int, opt ...Opt) *heading {
	if level < 1 || level > 6 {
		panic("heading level must be between 1 and 6")
	}

	// Create heading element
	tag := fmt.Sprintf("H%d", level)
	root := dom.GetWindow().Document().CreateElement(tag)

	// Apply options
	if opts, err := NewOpts(HeadingComponent, opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}
	}

	return &heading{
		component: component{
			name: HeadingComponent,
			root: root,
		},
		level: level,
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

// Level returns the heading level (1-6)
func (h *heading) Level() int {
	return h.level
}
