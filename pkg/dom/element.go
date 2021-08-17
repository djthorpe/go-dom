//go:build !js
// +build !js

package dom

import (
	"fmt"

	dom "github.com/djthorpe/go-dom"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *element) String() string {
	str := "<DOMElement"
	str += fmt.Sprint(" ", this.node)
	return str + ">"
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *element) InnerHTML() string {
	return "TODO"
}

func (this *element) OuterHTML() string {
	return "TODO"
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *element) AppendChild(child dom.Node) dom.Node {
	this.appendchild(child, this)
	return child
}

func (this *element) NextSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return nextSibling(this.parent, this)
	}
}

func (this *element) PreviousSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return previousSibling(this.parent, this)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *element) v() *node {
	return this.node
}
