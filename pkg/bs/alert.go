package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type alert struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewAlert = "mvc-bs-alert"
)

func init() {
	RegisterView(ViewAlert, newAlertFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Alert(opt ...Opt) *alert {
	opt = append([]Opt{WithClass("alert", "alert-primary"), WithAttr("role", "alert")}, opt...)
	view := &alert{NewView(ViewAlert, "DIV", opt...)}
	return view
}

func newAlertFromElement(element Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	return &alert{NewViewWithElement(element)}
}
