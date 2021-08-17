//go:build js

package dom

import (
	"syscall/js"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	*node
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *element) InnerHTML() string {
	return this.Get("innerHTML").String()
}

func (this *element) OuterHTML() string {
	return this.Get("outerHTML").String()
}

/////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *element) v() js.Value {
	return this.node.v()
}

/*
/////////////////////////////////////////////////////////////////////
// LIFECYCLE

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func NewElement(document dom.Document, name string) *element {
	element := GetWindow
	return &element{NewNode(document, name, dom.ELEMENT_NODE)}
}

func (this *element) AppendChild(child dom.Node) dom.Node {
	this.appendchild(child, this)
	return child
}

func (this *element) NextSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return nextSibling(this.parent, this)
	}
}

func (this *element) PreviousSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return previousSibling(this.parent, this)
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *element) String() string {
	str := "<DOMElement"
	str += fmt.Sprint(" ", this.node)
	return str + ">"
}
*/
