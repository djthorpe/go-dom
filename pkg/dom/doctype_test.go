package dom_test

import (
	"testing"

	// Modules
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func Test_Doctype_001(t *testing.T) {
	doc := NewWindow().Document()
	if doc.Doctype() == nil {
		t.Error("Doctype() returned nil")
	} else {
		t.Log(doc.Doctype())
	}
}
