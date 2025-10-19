//go:build !js
// +build !js

package dom

import (
	"fmt"
	"io"
	"strconv"

	dom "github.com/djthorpe/go-dom"
	"golang.org/x/net/html"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type window struct {
	*document
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// GetWindow returns a global window object
func GetWindow() dom.Window {
	return &window{NewHTMLDocument("")}
}

// GetWindowWithTitle returns a global window object
func GetWindowWithTitle(title string) dom.Window {
	return &window{NewHTMLDocument(title)}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (w *window) String() string {
	str := "<DOMWindow"
	str += fmt.Sprint(" document=", w.document)
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Document returns window document
func (this *window) Document() dom.Document {
	return this.document
}

// Write the nodes inside a document, or the node itself otherwise
func (this *window) Write(w io.Writer, node dom.Node) (int, error) {
	if node == nil {
		return 0, dom.ErrBadParameter
	}
	switch node.NodeType() {
	case dom.DOCUMENT_TYPE_NODE:
		var s int
		for child := node.FirstChild(); child != nil; child = child.NextSibling() {
			if n, err := child.(nodevalue).write(w); err != nil {
				return 0, err
			} else {
				s += n
			}
		}
		return s, nil
	default:
		return node.(nodevalue).write(w)
	}
}

// Read in a document from a string, and set mimetype
func (this *window) Read(r io.Reader, mimetype string) (dom.Document, error) {
	var node dom.Element

	parser := html.NewTokenizer(r)
	doc := NewDocument()

	for {
		tt := parser.Next()
		switch tt {
		case html.ErrorToken:
			if err := parser.Err(); err == io.EOF {
				return doc, nil
			} else {
				return nil, err
			}
		case html.CommentToken:
			if node != nil {
				node.AppendChild(doc.CreateComment(string(parser.Text())))
			}
		case html.TextToken:
			if node != nil {
				node.AppendChild(doc.CreateTextNode(string(parser.Text())))
			}
		case html.DoctypeToken:
			fmt.Println("TODO doctype=", string(parser.Text()))
		case html.SelfClosingTagToken:
			elem := readCreateElement(doc, parser)
			if node == nil {
				doc.AppendChild(elem)
			} else {
				node.AppendChild(elem)
			}
		case html.StartTagToken:
			elem := readCreateElement(doc, parser)
			if node == nil {
				node = doc.AppendChild(elem).(dom.Element)
			} else {
				node = node.AppendChild(elem).(dom.Element)
			}
		case html.EndTagToken:
			tag, _ := parser.TagName()
			if node == nil || node.NodeName() != string(tag) {
				return nil, dom.ErrUnexpectedResponse.With("Unclosed tag: ", node.NodeName(), strconv.Quote(string(tag)))
			} else {
				node = node.ParentElement()
			}
		default:
			return nil, dom.ErrBadParameter.With(tt)
		}
	}
}

func readCreateElement(doc dom.Document, parser *html.Tokenizer) dom.Element {
	tag, hasattr := parser.TagName()
	elem := doc.CreateElement(string(tag))
	if hasattr {
		for {
			name, value, more := parser.TagAttr()
			if !more {
				break
			}
			elem.SetAttribute(string(name), string(value))
		}
	}
	return elem
}
