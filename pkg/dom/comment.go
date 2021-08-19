//go:build !js

package dom

import (
	"fmt"
	"html"
	"io"

	dom "github.com/djthorpe/go-dom"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type comment struct {
	*node
}

/////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	startcomment = []byte("<!--")
	endcomment   = []byte("-->")
)

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *comment) String() string {
	str := "<DOMComment"
	str += fmt.Sprintf(" data=%q length=%v", this.Data(), this.Length())
	return str + ">"
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *comment) NextSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return this.parent.(nodevalue).nextChild(this)
	}
}

func (this *comment) PreviousSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return this.parent.(nodevalue).previousChild(this)
	}
}

func (this *comment) Data() string {
	return this.cdata
}

func (this *comment) Length() int {
	return len(this.cdata)
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *comment) AppendChild(child dom.Node) dom.Node {
	// NO-OP
	return nil
}

func (this *comment) CloneNode(bool) dom.Node {
	return NewNode(this.document, this.name, this.nodetype, this.cdata)
}

func (this *comment) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	// NO-OP
	return nil
}

func (this *comment) RemoveChild(child dom.Node) {
	// NO-OP
}

func (this *comment) ReplaceChild(dom.Node, dom.Node) {
	// NO-OP
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *comment) v() *node {
	return this.node
}

func (this *comment) write(w io.Writer) (int, error) {
	s := 0
	if n, err := w.Write(startcomment); err != nil {
		return 0, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte(html.EscapeString(this.cdata))); err != nil {
		return 0, err
	} else {
		s += n
	}
	if n, err := w.Write(endcomment); err != nil {
		return 0, err
	} else {
		s += n
	}
	return s, nil
}
