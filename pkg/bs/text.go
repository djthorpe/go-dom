package bs

import (
	"slices"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type text struct {
	View
}

var _ View = (*text)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewText = "mvc-bs-text"
)

var (
	textTagNames = []string{
		"P",
		"DEL",
		"MARK",
		"SMALL",
		"STRONG",
		"EM",
		"BLOCKQUOTE",
		"CODE",
	}
)

func init() {
	RegisterView(ViewText, newTextFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Para(children ...any) *text {
	view := &text{NewView(ViewText, "P")}
	view.Append(children...)
	return view
}

func Deleted(children ...any) *text {
	view := &text{NewView(ViewText, "DEL")}
	view.Append(children...)
	return view
}

func Highlighted(children ...any) *text {
	view := &text{NewView(ViewText, "MARK")}
	view.Append(children...)
	return view
}

func Small(children ...any) *text {
	view := &text{NewView(ViewText, "SMALL")}
	view.Append(children...)
	return view
}

func Strong(children ...any) *text {
	view := &text{NewView(ViewText, "STRONG")}
	view.Append(children...)
	return view
}

func Em(children ...any) *text {
	view := &text{NewView(ViewText, "EM")}
	view.Append(children...)
	return view
}

func Blockquote(children ...any) *text {
	view := &text{NewView(ViewText, "BLOCKQUOTE", WithClass("blockquote"))}
	view.Append(children...)
	return view
}

func Code(children ...any) *text {
	view := &text{NewView(ViewText, "CODE")}
	view.Append(children...)
	return view
}

func newTextFromElement(element Element) View {
	if !slices.Contains(textTagNames, element.TagName()) {
		return nil
	}
	return &text{NewViewWithElement(element)}
}
