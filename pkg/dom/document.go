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

func NewDocument() *document {
	return NewNode(nil, "#document", dom.DOCUMENT_NODE, "").(*document)
}

func NewHTMLDocument(title string) *document {
	doc := NewDocument()

	// Set doctype, root, head, body, charset and title to document
	doc.doctype = NewNode(doc, "html", dom.DOCUMENT_TYPE_NODE, "").(dom.DocumentType)
	doc.AppendChild(doc.CreateElement("html"))
	doc.head = doc.FirstChild().AppendChild(doc.CreateElement("head")).(dom.Element)
	doc.charset = doc.head.AppendChild(doc.CreateElement("meta")).(dom.Element)
	if title != "" {
		titlenode := doc.head.AppendChild(doc.CreateElement("title")).(dom.Element)
		titlenode.AppendChild(doc.CreateTextNode(title))
	}
	doc.body = doc.FirstChild().AppendChild(doc.CreateElement("body")).(dom.Element)

	// Return the document
	return doc
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *document) NextSibling() dom.Node {
	// NO-OP
	return nil
}

func (this *document) PreviousSibling() dom.Node {
	// NO-OP
	return nil
}

func (this *document) Body() dom.Element {
	return this.body
}

func (this *document) Doctype() dom.DocumentType {
	return this.doctype
}

func (this *document) DocumentElement() dom.Element {
	return this.FirstChild().(dom.Element)
}

func (doc *document) Title() string {
	if doc.head == nil {
		return ""
	}
	return fmt.Sprint(doc.node.v())
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *document) AppendChild(child dom.Node) dom.Node {
	if child.NodeType() != dom.ELEMENT_NODE {
		return nil
	}
	if this.HasChildNodes() {
		return nil
	}
	return this.AppendChild(child)
}

func (this *document) CloneNode(deep bool) dom.Node {
	clone := NewNode(this.document, this.name, this.nodetype, this.cdata).(*document)
	if this.doctype != nil {
		clone.doctype = this.doctype.CloneNode(deep).(dom.DocumentType)
	}
	if root := this.FirstChild(); root != nil {
		clone.AppendChild(root.CloneNode(deep).(dom.Element))
	}
	// TODO: set head, body, charset from this.root
	return clone
}

func (this *document) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	// NO-OP
	return nil
}

func (this *document) RemoveChild(child dom.Node) {
	// NO-OP
}

func (this *document) ReplaceChild(dom.Node, dom.Node) {
	// NO-OP
}

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

func (this *document) ActiveElement() dom.Element {
	// Not supported in non-WASM builds
	return nil
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
