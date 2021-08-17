//go:build !js

package dom

import (
	"fmt"

	dom "github.com/djthorpe/go-dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type document struct {
	*node

	body dom.Element
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewHTMLDocument() *document {
	doc := NewNode(nil, "#document", dom.DOCUMENT_NODE).(*document)
	doc.body = doc.AppendChild(doc.CreateElement("html")).AppendChild(doc.CreateElement("body")).(dom.Element)
	return doc
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *document) Body() dom.Element {
	return this.body
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *document) CreateElement(name string) dom.Element {
	return NewNode(this, name, dom.ELEMENT_NODE).(dom.Element)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *document) v() *node {
	return this.node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *document) String() string {
	str := "<DOMDocument"
	str += fmt.Sprint(" ", this.node)
	return str + ">"
}
