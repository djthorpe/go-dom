//go:build js

package dom

import (
	"fmt"
	"io"
	"syscall/js"

	"github.com/djthorpe/go-dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type window struct {
	js.Value
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	w = &window{js.Global()}
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewWindow() dom.Window {
	return w
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *window) String() string {
	str := "<DOMWindow"
	if doc := this.Document(); doc != nil {
		str += fmt.Sprint(" document=", doc)
	}
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *window) Document() dom.Document {
	return NewNode(w.Get("document")).(dom.Document)
}

// Write the HTML inside a node
func (this *window) Write(node dom.Node, w io.Writer) (int, error) {
	var s int
	if node == nil {
		return 0, dom.ErrBadParameter
	}
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		html := child.(nodevalue).v().Get("outerHTML").String()
		if n, err := w.Write([]byte(html)); err != nil {
			return 0, err
		} else {
			s += n
		}
	}
	return s, nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS
