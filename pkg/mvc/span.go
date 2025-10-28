package mvc

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// div is a simple div view
type span struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewSpan = "mvc-span"
)

func init() {
	RegisterView(ViewSpan, newSpanFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Span(opts ...Opt) *span {
	return &span{NewView(ViewSpan, "SPAN", opts...)}
}

func newSpanFromElement(element Element) View {
	if element.TagName() != "SPAN" {
		return nil
	}
	return &span{NewViewWithElement(element)}
}
