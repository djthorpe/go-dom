package mvc

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// div is a simple div view
type div struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewDiv = "mvc-div"
)

func init() {
	RegisterView(ViewDiv, newDivFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Div(opts ...Opt) *div {
	view := &div{NewView(ViewDiv, "DIV", opts...)}
	return view
}

func newDivFromElement(element Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	return &div{NewViewWithElement(element)}
}
