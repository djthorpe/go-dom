//go:build !js

package dom

import (
	"fmt"

	dom "github.com/djthorpe/go-dom"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type text struct {
	*node
	data string
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

func (this *text) Data() string {
	return this.data
}

func (this *text) Length() int {
	return len(this.data)
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *text) AppendChild(child dom.Node) dom.Node {
	return nil
}

func (this *text) CloneNode(bool) dom.Node {
	clone := NewNode(this.document, this.name, this.nodetype).(*text)
	clone.data = this.data
	return clone
}

func (this *text) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
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
