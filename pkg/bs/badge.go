package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type badge struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewBadge = "mvc-bs-badge"
)

func init() {
	RegisterView(ViewBadge, newBadgeFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Badge(opt ...Opt) *badge {
	opt = append([]Opt{WithClass("badge"), WithColor(PRIMARY)}, opt...)
	view := &badge{NewView(ViewBadge, "SPAN", opt...)}
	return view
}

func PillBadge(opt ...Opt) *badge {
	return Badge(append(opt, WithClass("rounded-pill"))...)
}

func newBadgeFromElement(element Element) View {
	if element.TagName() != "SPAN" {
		return nil
	}
	return &badge{NewViewWithElement(element)}
}
