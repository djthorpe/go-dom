//go:build !js
// +build !js

package dom

import (
	"fmt"
	"io"

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

func NewWindow() dom.Window {
	return &window{NewHTMLDocument("")}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *window) String() string {
	str := "<DOMWindow"
	str += fmt.Sprint(" document=", this.document)
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
	parser := html.NewTokenizer(r)
	depth := 0
	for {
		tt := parser.Next()
		switch tt {
		case html.ErrorToken:
			if err := parser.Err(); err == io.EOF {
				return nil, nil
			} else {
				return nil, err
			}
		case html.CommentToken:
			if depth > 0 {
				fmt.Println("comment=", string(parser.Text()))
			}
		case html.TextToken:
			if depth > 0 {
				fmt.Println("text=", string(parser.Text()))
			}
		case html.DoctypeToken:
			fmt.Println("doctype=", string(parser.Text()))
		case html.SelfClosingTagToken:
			tn, _ := parser.TagName()
			fmt.Println("self closing tag=", string(tn))
		case html.StartTagToken, html.EndTagToken:
			tn, _ := parser.TagName()
			if tt == html.StartTagToken {
				depth++
			} else {
				depth--
			}
			fmt.Println(depth, "tt=", tt, " name=", string(tn))
		}
	}
}
