package dom_test

import (
	"testing"

	// Modules
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func Test_Document_001(t *testing.T) {
	if doc := NewWindow().Document(); doc == nil {
		t.Error("Unexpected nil returned from NewWindow().Document()")
	} else if body := doc.Body(); body == nil {
		t.Error("Unexpected nil returned from Document.Body()")
	} else if elem := doc.CreateElement("div"); elem == nil {
		t.Error("Unexpected nil returned from doc.CreateElement")
	} else if elem2 := body.AppendChild(elem); elem2 != elem {
		t.Error("Unexpected return from AppendChild")
	}
}
