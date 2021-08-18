package dom_test

import (
	"testing"

	// Modules
	"github.com/djthorpe/go-dom"
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func Test_Text_001(t *testing.T) {
	doc := NewWindow().Document()
	text := doc.CreateTextNode("test")
	if text.NodeType() != dom.TEXT_NODE {
		t.Error("Expected TEXT_NODE")
	} else if text.NodeName() != "#text" {
		t.Errorf("Expected #text, got %q", text.NodeName())
	} else {
		t.Log(text)
	}
}
