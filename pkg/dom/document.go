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
	charset dom.Element
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewHTMLDocument(title string) *document {
	doc := NewNode(nil, "#document", dom.DOCUMENT_NODE, "").(*document)

	// Append doctype, head, body and title to document
	doc.doctype = doc.AppendChild(NewNode(doc, "html", dom.DOCUMENT_TYPE_NODE, "")).(dom.DocumentType)
	html := doc.AppendChild(doc.CreateElement("html"))
	doc.head = html.AppendChild(doc.CreateElement("head")).(dom.Element)
	doc.charset = doc.head.AppendChild(doc.CreateElement("meta")).(dom.Element)
	doc.body = html.AppendChild(doc.CreateElement("body")).(dom.Element)
	if title != "" {
		titlenode := doc.head.AppendChild(doc.CreateElement("title")).(dom.Element)
		titlenode.AppendChild(doc.CreateTextNode(title))
	}

	// Return the document
	return doc
}

func NewXMLDocument(root string) *document {
	doc := NewNode(nil, "#document", dom.DOCUMENT_NODE, "").(*document)
	doc.body = doc.AppendChild(doc.CreateElement(root)).(dom.Element)

	// Return the document
	return doc
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *document) NextSibling() dom.Node {
	return nextSibling(this.parent, this)
}

func (this *document) PreviousSibling() dom.Node {
	return previousSibling(this.parent, this)
}

func (this *document) Body() dom.Element {
	return this.body
}

func (this *document) Doctype() dom.DocumentType {
	return this.doctype
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *document) CreateElement(name string) dom.Element {
	return NewNode(this, name, dom.ELEMENT_NODE, "").(dom.Element)
}

func (this *document) CreateComment(cdata string) dom.Comment {
	return NewNode(this, "#comment", dom.COMMENT_NODE, cdata).(dom.Comment)
}

func (this *document) CreateTextNode(cdata string) dom.Text {
	return NewNode(this, "#text", dom.TEXT_NODE, cdata).(dom.Text)
}

func (this *document) CreateAttribute(name string) dom.Attr {
	return NewNode(this, name, dom.ATTRIBUTE_NODE, "").(dom.Attr)
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
