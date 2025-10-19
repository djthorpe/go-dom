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
	doctype := this.Get("doctype")
	if !doctype.Truthy() {
		return nil
	} else {
		return NewNode(doctype).(dom.DocumentType)
	}
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

func (this *document) CreateAttribute(name string) dom.Attr {
	return NewNode(this.Call("createAttribute", name)).(dom.Attr)
}

/////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *document) v() js.Value {
	return this.node.v()
}
