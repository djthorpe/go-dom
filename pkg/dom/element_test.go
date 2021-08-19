package dom_test

import (
	"testing"

	// Modules

	"github.com/djthorpe/go-dom"
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func Test_Element_001(t *testing.T) {
	element := GetWindow().Document().CreateElement("A")
	if element.NodeName() != "A" {
		t.Errorf("Element.NodeName() = %v, want %v", element.NodeName(), "a")
	}
	if element.NodeType() != dom.ELEMENT_NODE {
		t.Errorf("Element.NodeType() = %v, want %v", element.NodeType(), dom.ELEMENT_NODE)
	}
}

func Test_Element_002(t *testing.T) {
	doc := GetWindow().Document()
	parent := doc.CreateElement("A")
	b := doc.CreateElement("CB")
	c := doc.CreateElement("CC")
	if parent.HasChildNodes() != false {
		t.Error("HasChildNodes() failed")
	}
	if b.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	if c.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	parent.AppendChild(b)
	if parent.FirstChild().Equals(b) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(b) == false {
		t.Error("LastChild() failed")
	}
	if parent.HasChildNodes() == false {
		t.Error("HasChildNodes() failed: ", parent)
	}
	if b.ParentNode().Equals(parent) == false {
		t.Error("ParentNode() failed")
	}
	if b.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if b.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.AppendChild(c)
	if b.NextSibling() == nil {
		t.Error("NextSibling() failed")
	}
	if parent.FirstChild().Equals(b) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(c) == false {
		t.Error("LastChild() failed")
	}
	if c.ParentNode().Equals(parent) == false {
		t.Error("ParentNode() failed")
	}
	if b.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if b.NextSibling().Equals(c) == false {
		t.Error("NextSibling() failed")
	}
	if c.PreviousSibling().Equals(b) == false {
		t.Error("PreviousSibling() failed")
	}
	if c.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.RemoveChild(b)
	if parent.FirstChild().Equals(c) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(c) == false {
		t.Error("LastChild() failed")
	}
	if b.ParentNode() != nil {
		t.Error("ParentNode() failed", b.ParentNode())
	}
	if c.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if c.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.RemoveChild(c)
	if parent.FirstChild() != nil {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild() != nil {
		t.Error("LastChild() failed")
	}
	if parent.HasChildNodes() {
		t.Error("HasChildNodes() failed")
	}
	if c.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
}
func Test_Element_003(t *testing.T) {
	doc := GetWindow().Document()
	parent := doc.CreateElement("a")
	if parent.InnerHTML() != "" {
		t.Error("InnerHTML() failed")
	}
	if parent.OuterHTML() != "<a></a>" {
		t.Error("OuterHTML() failed: ", parent.OuterHTML())
	}
}
