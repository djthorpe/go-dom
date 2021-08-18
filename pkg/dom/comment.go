//go:build !js

package dom

import (
	"fmt"

	dom "github.com/djthorpe/go-dom"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type comment struct {
	*node
	data string
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *comment) String() string {
	str := "<DOMComment"
	str += fmt.Sprintf(" data=%q length=%v", this.Data(), this.Length())
	return str + ">"
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *comment) Data() string {
	return this.data
}

func (this *comment) Length() int {
	return len(this.data)
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *comment) AppendChild(child dom.Node) dom.Node {
	return nil
}

func (this *comment) CloneNode(bool) dom.Node {
	clone := NewNode(this.document, this.name, this.nodetype).(*text)
	clone.data = this.data
	return clone
}

func (this *comment) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
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
