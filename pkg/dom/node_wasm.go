//go:build js

package dom

import (
	"fmt"
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type node struct {
	js.Value
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE HELPERS

// toJSValue returns the underlying js.Value from any node type
// This replaces the nodevalue interface v() method
func toJSValue(n dom.Node) js.Value {
	if n == nil {
		return js.Undefined()
	}
	// All node types embed *node which embeds js.Value
	// So we can use a type assertion to access it
	switch v := n.(type) {
	case *node:
		return v.Value
	case *element:
		return v.node.Value
	case *attr:
		return v.node.Value
	case *text:
		return v.node.Value
	case *comment:
		return v.node.Value
	case *doctype:
		return v.node.Value
	case *document:
		return v.node.Value
	default:
		panic("toJSValue: unknown node type")
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// JSValue returns the underlying js.Value for interop with JavaScript
func (n *node) JSValue() js.Value {
	return n.Value
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	// Constructors
	cArray        = js.Global().Get("Array")
	cObject       = js.Global().Get("Object")
	cNode         = js.Global().Get("Node")
	cText         = js.Global().Get("Text")
	cComment      = js.Global().Get("Comment")
	cDocument     = js.Global().Get("HTMLDocument")
	cDocumentType = js.Global().Get("DocumentType")
	cElement      = js.Global().Get("HTMLElement")
	cAttr         = js.Global().Get("Attr")
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewNode returns a new go object wrapping js.Value
func NewNode(v js.Value) dom.Node {
	if v.IsNull() || v.IsUndefined() {
		return nil
	}
	proto := v
	for {
		proto = cObject.Call("getPrototypeOf", proto)
		if proto.IsNull() || proto.IsUndefined() {
			panic(fmt.Sprint("Unknown constructor"))
		}

		// Check if this prototype matches any known types
		// For custom elements, we check the prototype itself, not its constructor
		switch {
		case proto.Equal(cDocument.Get("prototype")):
			return &document{node: &node{v}}
		case proto.Equal(cElement.Get("prototype")):
			return &element{node: &node{v}}
		case proto.Equal(cText.Get("prototype")):
			return &text{node: &node{v}}
		case proto.Equal(cComment.Get("prototype")):
			return &comment{node: &node{v}}
		case proto.Equal(cDocumentType.Get("prototype")):
			return &doctype{node: &node{v}}
		case proto.Equal(cAttr.Get("prototype")):
			return &attr{node: &node{v}}
		case proto.Equal(cNode.Get("prototype")):
			return &node{v}
		} // Also check constructor for compatibility (legacy behavior)
		c := constructor(proto)
		if c.IsNull() || c.IsUndefined() {
			panic("NewNode failed for " + constructor(cObject.Call("getPrototypeOf", v)).Get("name").String())
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *node) String() string {
	str := "<DOMNode"
	if name := this.NodeName(); name != "" {
		str += fmt.Sprintf(" name=%q", name)
	}
	if t := this.NodeType(); t != dom.UNKNOWN_NODE {
		str += fmt.Sprint(" type=", t)
	}
	if parent := this.ParentNode(); parent != nil {
		str += " parent=<DOMNode"
		if name := parent.NodeName(); name != "" {
			str += fmt.Sprintf(" name=%q", name)
		}
		return str + ">"
	}
	for c := this.FirstChild(); c != nil; c = c.NextSibling() {
		str += fmt.Sprint(" child=", c)
	}
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *node) BaseURI() string {
	return this.Get("baseURI").String()
}

func (this *node) ChildNodes() []dom.Node {
	return fromNodeList(this.Get("childNodes"))
}

func (this *node) Contains(other dom.Node) bool {
	return this.Call("contains", toJSValue(other)).Bool()
}

func (this *node) FirstChild() dom.Node {
	return NewNode(this.Get("firstChild"))
}

func (this *node) HasChildNodes() bool {
	return this.Call("hasChildNodes").Bool()
}

func (this *node) IsConnected() bool {
	return this.Get("isConnected").Bool()
}

func (this *node) LastChild() dom.Node {
	return NewNode(this.Get("lastChild"))
}

func (this *node) NextSibling() dom.Node {
	return NewNode(this.Get("nextSibling"))
}

func (this *node) NodeName() string {
	return this.Get("nodeName").String()
}

func (this *node) NodeType() dom.NodeType {
	return dom.NodeType(this.Get("nodeType").Int())
}

func (this *node) OwnerDocument() dom.Document {
	return NewNode(this.Get("ownerDocument")).(dom.Document)
}

func (this *node) ParentNode() dom.Node {
	return NewNode(this.Get("parentNode"))
}

func (this *node) ParentElement() dom.Element {
	return NewNode(this.Get("parentElement")).(dom.Element)
}

func (this *node) PreviousSibling() dom.Node {
	return NewNode(this.Get("previousSibling"))
}

func (this *node) TextContent() string {
	return this.Get("textContent").String()
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *node) Equals(other dom.Node) bool {
	return this.Equal(toJSValue(other))
}

func (this *node) AppendChild(child dom.Node) dom.Node {
	this.Call("appendChild", toJSValue(child))
	return child
}

func (this *node) CloneNode(deep bool) dom.Node {
	return NewNode(this.Call("cloneNode", deep))
}

func (this *node) InsertBefore(child dom.Node, before dom.Node) dom.Node {
	if before == nil {
		return this.AppendChild(child)
	} else {
		this.Call("insertBefore", toJSValue(child), toJSValue(before))
		return child
	}
}

func (this *node) RemoveChild(child dom.Node) {
	this.Call("removeChild", toJSValue(child))
}

func (this *node) ReplaceChild(new, old dom.Node) {
	this.Call("replaceChild", toJSValue(new), toJSValue(old))
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *node) v() js.Value {
	return this.Value
}

func fromNodeList(v js.Value) []dom.Node {
	var result []dom.Node
	for _, v := range nodeListToSlice(v) {
		result = append(result, NewNode(v))
	}
	return result
}

func nodeListToSlice(v js.Value) []js.Value {
	length := v.Get("length").Int()
	result := make([]js.Value, length)
	for i := 0; i < length; i++ {
		result[i] = v.Call("item", i)
	}
	return result
}

func constructor(v js.Value) js.Value {
	if v.IsNull() || v.IsUndefined() {
		return v
	} else {
		return v.Get("constructor")
	}
}
