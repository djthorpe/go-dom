//go:build js

package dom

import (
	"fmt"
	"syscall/js"

	"github.com/djthorpe/go-dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type node struct {
	js.Value
}

type nodevalue interface {
	dom.Node
	v() js.Value
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	// Constructors
	cArray    = js.Global().Get("Array")
	cObject   = js.Global().Get("Object")
	cDocument = js.Global().Get("HTMLDocument")
	cElement  = js.Global().Get("HTMLElement")
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

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
		switch c := constructor(proto); {
		case c.Equal(cDocument):
			return &document{node: &node{v}}
		case c.Equal(cElement):
			return &element{node: &node{v}}
		case c.Equal(cObject):
			return &node{v}
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *node) BaseURI() string {
	return this.Get("baseURI").String()
}

func (this *node) ChildNodes() []dom.Node {
	return fromNodeList(this.Get("childNodes"))
}

func (this *node) FirstChild() dom.Node {
	return NewNode(this.Get("firstChild"))
}

func (this *node) IsConnected() bool {
	// TODO
	return false
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
	return this.v().Equal(other.(nodevalue).v())
}

func (this *node) AppendChild(child dom.Node) dom.Node {
	this.Call("appendChild", child.(nodevalue).v())
	return child
}

func (this *node) CloneNode() dom.Node {
	// TODO
	return nil
}

func (this *node) Contains(dom.Node) bool {
	// TODO
	return false
}

func (this *node) HasChildNodes() bool {
	// TODO
	return false
}

func (this *node) InsertBefore(dom.Node, dom.Node) dom.Node {
	// TODO
	return nil
}

func (this *node) RemoveChild(dom.Node) {

}

func (this *node) ReplaceChild(dom.Node, dom.Node) {

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
	return v.Get("constructor")
}

func toMap(v js.Value) map[string]interface{} {
	result := make(map[string]interface{})
	if v.Type() != js.TypeObject {
		return nil
	}
	entries := cObject.Call("entries", v)
	fmt.Println("iter=", entries.Get("length"))
	return result
}
