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

type text struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *text) String() string {
	str := "<DOMText"
	str += fmt.Sprintf(" data=%q length=%v", this.Data(), this.Length())
	return str + ">"
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *text) NextSibling() dom.Node {
	return nextSibling(this.parent, this)
}

func (this *text) PreviousSibling() dom.Node {
	return previousSibling(this.parent, this)
}

func (this *text) Data() string {
	return this.cdata
}

func (this *text) Length() int {
	return len(this.cdata)
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *text) CloneNode(bool) dom.Node {
	return NewNode(this.document, this.name, this.nodetype, this.cdata)
}

func (this *text) AppendChild(child dom.Node) dom.Node {
	// NO-OP
	return nil
}

func (this *text) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	// NO-OP
	return nil
}

func (this *text) RemoveChild(child dom.Node) {
	// NO-OP
}

func (this *text) ReplaceChild(dom.Node, dom.Node) {
	// NO-OP
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *text) v() *node {
	return this.node
}

func (this *text) write(w io.Writer) (int, error) {
	return w.Write([]byte(html.EscapeString(this.cdata)))
}
