//go:build !js

package dom

import (
	"fmt"
	"strings"

	dom "github.com/djthorpe/go-dom"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type node struct {
	document dom.Document
	parent   dom.Node
	name     string
	nodetype dom.NodeType
	children []dom.Node
}

type nodevalue interface {
	dom.Node
	v() *node
}

/////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewNode(doc dom.Document, name string, nodetype dom.NodeType) dom.Node {
	switch nodetype {
	case dom.DOCUMENT_NODE:
		return &document{&node{nil, nil, name, nodetype, nil}, nil}
	case dom.ELEMENT_NODE:
		return &element{&node{doc, nil, name, nodetype, nil}}
	case dom.TEXT_NODE:
		return &node{doc, nil, name, nodetype, nil}
	case dom.COMMENT_NODE:
		return &node{doc, nil, name, nodetype, nil}
	case dom.DOCUMENT_TYPE_NODE:
		return &node{doc, nil, name, nodetype, nil}
	default:
		return &node{doc, nil, name, nodetype, nil}
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *node) String() string {
	str := "<DOMNode"
	if this.name != "" {
		str += fmt.Sprintf(" name=%q", this.name)
	}
	if this.nodetype != dom.UNKNOWN_NODE {
		str += fmt.Sprint(" type=", this.nodetype)
	}
	if this.parent != nil {
		str += " parent=<DOMNode"
		if name := this.parent.NodeName(); name != "" {
			str += fmt.Sprintf(" name=%q", name)
		}
		return str + ">"
	}
	for c := this.FirstChild(); c != nil; c = c.NextSibling() {
		str += fmt.Sprint(" child=", c)
	}
	return str + ">"
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *node) BaseURI() string {
	// TODO
	return "TODO"
}

func (this *node) Equals(other dom.Node) bool {
	return this.v() == other.(nodevalue).v()
}

func (this *node) ChildNodes() []dom.Node {
	result := make([]dom.Node, len(this.children))
	for i := range this.children {
		result[i] = this.children[i]
	}
	return result
}

func (this *node) FirstChild() dom.Node {
	if len(this.children) > 0 {
		return this.children[0]
	} else {
		return nil
	}
}

func (this *node) IsConnected() bool {
	return this.parent != nil
}

func (this *node) LastChild() dom.Node {
	last := len(this.children) - 1
	if last >= 0 {
		return this.children[last]
	} else {
		return nil
	}
}

func (this *node) NextSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return nextSibling(this.parent, this)
	}
}

func (this *node) NodeName() string {
	return strings.ToUpper(this.name)
}

func (this *node) NodeType() dom.NodeType {
	return this.nodetype
}

func (this *node) OwnerDocument() dom.Document {
	return this.document
}

func (this *node) ParentNode() dom.Node {
	return this.parent
}

func (this *node) ParentElement() dom.Element {
	if this.parent != nil && this.parent.NodeType() == dom.ELEMENT_NODE {
		return this.parent.(dom.Element)
	} else {
		return nil
	}
}

func (this *node) PreviousSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return previousSibling(this.parent, this)
	}
}

func (this *node) TextContent() string {
	// TODO
	return "TODO"
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *node) AppendChild(child dom.Node) dom.Node {
	this.appendchild(child, this)
	return child
}

func (this *node) CloneNode() dom.Node {
	clone := NewNode(this.document, this.name, this.nodetype).(nodevalue)
	clone.v().children = make([]dom.Node, len(this.children))
	for i := range this.children {
		child := this.children[i].CloneNode()
		child.(*node).parent = clone
	}
	return clone
}

func (this *node) Contains(child dom.Node) bool {
	for _, c := range this.children {
		if c == child {
			return true
		}
	}
	for _, c := range this.children {
		if c.Contains(child) {
			return true
		}
	}
	return false
}

func (this *node) HasChildNodes() bool {
	return len(this.children) > 0
}

func (this *node) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	return nil
	/*
		if ref == nil {
			// newNode is inserted at the end of parentNode's child nodes.
		}
		switch ref := ref.(type) {
		case *node:
			if ref.parent == this {
				this.removechild(child)
				child.parent = nil
			}
		case *element:
		default:
			panic("InsertBefore: not a *node")
		}
		// TODO
	*/
}

func (this *node) RemoveChild(child dom.Node) {
	switch child := child.(type) {
	case *node:
		if child.parent == this {
			this.removechild(child)
			child.parent = nil
		}
	case *element:
		if child.parent != nil {
			this.removechild(child)
			child.parent = nil
		}
	default:
		panic("RemoveChild: not a *node")
	}
}

func (this *node) ReplaceChild(dom.Node, dom.Node) {
	// TODO
}

/////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *node) v() *node {
	return this
}

// Append child node and set parent node
func (this *node) appendchild(child, parent dom.Node) {
	switch child := child.(type) {
	case *node:
		if child.parent != nil {
			child.parent.RemoveChild(child)
		}
		child.parent = parent
		this.children = append(this.children, child)
	case *element:
		if child.parent != nil {
			child.parent.RemoveChild(child)
		}
		child.parent = parent
		this.children = append(this.children, child)
	default:
		panic("AppendChild: not a *node")
	}
}

// Return next child node
func nextSibling(parent, child dom.Node) dom.Node {
	switch parent := parent.(type) {
	case *node:
		for i, c := range parent.children {
			if c != child {
				continue
			}
			if i < len(parent.children)-1 {
				return parent.children[i+1]
			} else {
				return nil
			}
		}
	case *element:
		for i, c := range parent.children {
			if c != child {
				continue
			}
			if i < len(parent.children)-1 {
				return parent.children[i+1]
			} else {
				return nil
			}
		}
	default:
		panic("NextSibling: not a *node")
	}
	return nil
}

// Return previous child node
func previousSibling(parent, child dom.Node) dom.Node {
	switch parent := parent.(type) {
	case *node:
		for i, c := range parent.children {
			if c != child {
				continue
			}
			if i > 0 {
				return parent.children[i-1]
			} else {
				return nil
			}
		}
	case *element:
		for i, c := range parent.children {
			if c != child {
				continue
			}
			if i > 0 {
				return parent.children[i-1]
			} else {
				return nil
			}
		}
	default:
		panic("PreviousSibling: not a *node")
	}
	return nil
}

// Remove a child node
func (this *node) removechild(child dom.Node) {
	for i, c := range this.children {
		if c != child {
			continue
		}
		this.children = append(this.children[:i], this.children[i+1:]...)
		return
	}
}
