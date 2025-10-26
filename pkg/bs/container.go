package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// containers are elements to wrap any content
type container struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewContainer = "mvc-bs-container"
)

func init() {
	RegisterView(ViewContainer, newContainerFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Container(opt ...Opt) *container {
	view := &container{NewView(ViewContainer, "DIV", append([]Opt{WithClass("container")}, opt...)...)}
	return view
}

func newContainerFromElement(element Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	return &container{NewViewWithElement(element)}
}
