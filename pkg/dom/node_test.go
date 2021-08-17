//go:build !js
// +build !js

package dom_test

import (
	"testing"

	// Modules
	dom "github.com/djthorpe/go-dom"
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func Test_Node_001(t *testing.T) {
	node := NewNode(nil, "a", dom.ELEMENT_NODE)
	if node.NodeType() != dom.ELEMENT_NODE {
		t.Error("NodeType() failed")
	}
	if node.NodeName() != "A" {
		t.Error("NodeName() failed")
	}
}

func Test_Node_002(t *testing.T) {
	node := NewNode(nil, "a", dom.ELEMENT_NODE)
	b := NewNode(nil, "b", dom.ELEMENT_NODE)
	c := NewNode(nil, "c", dom.ELEMENT_NODE)
	if node.HasChildNodes() != false {
		t.Error("HasChildNodes() failed")
	}
	if b.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	if c.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	node.AppendChild(b)
	if node.FirstChild() != b {
		t.Error("FirstChild() failed")
	}
	if node.LastChild() != b {
		t.Error("LastChild() failed")
	}
	if node.HasChildNodes() != true {
		t.Error("HasChildNodes() failed")
	}
	if b.ParentNode() != node {
		t.Error("ParentNode() failed")
	}
	if b.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if b.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	node.AppendChild(c)
	if node.FirstChild() != b {
		t.Error("FirstChild() failed")
	}
	if node.LastChild() != c {
		t.Error("LastChild() failed")
	}
	if c.ParentNode() != node {
		t.Error("ParentNode() failed")
	}
	if b.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if b.NextSibling() != c {
		t.Error("NextSibling() failed")
	}
	if c.PreviousSibling() != b {
		t.Error("PreviousSibling() failed")
	}
	if c.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	node.RemoveChild(b)
	if node.FirstChild() != c {
		t.Error("FirstChild() failed")
	}
	if node.LastChild() != c {
		t.Error("LastChild() failed")
	}
	if b.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	if c.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if c.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	node.RemoveChild(c)
	if node.FirstChild() != nil {
		t.Error("FirstChild() failed")
	}
	if node.LastChild() != nil {
		t.Error("LastChild() failed")
	}
	if node.HasChildNodes() != false {
		t.Error("HasChildNodes() failed")
	}
	if c.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
}