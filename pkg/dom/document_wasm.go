//go:build js

package dom

import (
	"fmt"
	"syscall/js"

	dom "github.com/djthorpe/go-dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type document struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *document) String() string {
	str := "<DOMDocument"
	str += fmt.Sprint(" ", this.node)
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *document) Body() dom.Element {
	return NewNode(this.Get("body")).(dom.Element)
}

func (this *document) Doctype() dom.DocumentType {
	return NewNode(this.Get("doctype")).(dom.DocumentType)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *document) CreateElement(name string) dom.Element {
	return NewNode(this.Call("createElement", name)).(dom.Element)
}

func (this *document) CreateComment(data string) dom.Comment {
	return NewNode(this.Call("createComment", data)).(dom.Comment)
}

func (this *document) CreateTextNode(data string) dom.Text {
	return NewNode(this.Call("createTextNode", data)).(dom.Text)
}

/////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *document) v() js.Value {
	return this.node.v()
}
