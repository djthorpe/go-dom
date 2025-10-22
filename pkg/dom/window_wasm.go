//go:build js

package dom

import (
	"fmt"
	"io"
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type window struct {
	*node
}

var _ dom.Window = (*window)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// GetWindow returns a global window object
func GetWindow() dom.Window {
	w := &window{}
	w.node = &node{Value: js.Global(), eventListeners: make(map[string][]js.Func)}
	return w
}

// GetWindowWithTitle returns a global window object with the specified title
func GetWindowWithTitle(title string) dom.Window {
	w := GetWindow()
	//	TOOD: w.Document().Set("title", title)
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
	return NewNode(this.Get("document")).(dom.Document)
}

func (this *window) Location() dom.Location {
	return NewLocation(this.Get("location"))
}

// Write the HTML inside a node
func (this *window) Write(w io.Writer, node dom.Node) (int, error) {
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
