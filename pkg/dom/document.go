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

	doctype dom.DocumentType
	head    dom.Element
	body    dom.Element
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewHTMLDocument(title string) *document {
	doc := NewNode(nil, "#document", dom.DOCUMENT_NODE).(*document)
	html := doc.AppendChild(doc.CreateElement("html"))

	// Append doctype, head, body and title to document
	doc.doctype = doc.AppendChild(NewNode(doc, "html", dom.DOCUMENT_TYPE_NODE)).(dom.DocumentType)
	doc.head = html.AppendChild(doc.CreateElement("head")).(dom.Element)
	doc.body = html.AppendChild(doc.CreateElement("body")).(dom.Element)
	doc.head.AppendChild(doc.CreateTextNode(title))

	// Return the document
	return doc
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *document) Body() dom.Element {
	return this.body
}

func (this *document) Doctype() dom.DocumentType {
	return this.doctype
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *document) CreateElement(name string) dom.Element {
	return NewNode(this, name, dom.ELEMENT_NODE).(dom.Element)
}

func (this *document) CreateComment(data string) dom.Comment {
	comment := NewNode(this, "#comment", dom.COMMENT_NODE).(*comment)
	comment.data = data
	return comment
}

func (this *document) CreateTextNode(data string) dom.Text {
	text := NewNode(this, "#text", dom.TEXT_NODE).(*text)
	text.data = data
	return text
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
