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

func Test_Text_002(t *testing.T) {
	doc := NewWindow().Document()
	tests := []struct {
		data string
		want string
	}{
		{"", ""},
		{" ", " "},
		{"test", "test"},
		{"<test>", "&lt;test&gt;"},
		{"&", "&amp;"},
		{"<!-- test & test -->", "&lt;!-- test &amp; test --&gt;"},
	}
	for _, test := range tests {
		text := doc.CreateTextNode(test.data)
		if text.NodeType() != dom.TEXT_NODE {
			t.Error("Expected TEXT_NODE")
		} else if text.NodeName() != "#text" {
			t.Errorf("Expected #text, got %q", text.NodeName())
		} else if text.Data() != test.data {
			t.Errorf("Expected %q, got %q", test.data, text.Data())
		}
		div := doc.CreateElement("div")
		div.AppendChild(text)
		if div.InnerHTML() != test.want {
			t.Errorf("Expected %q, got %q", test.want, div.InnerHTML())
		}
	}
}
