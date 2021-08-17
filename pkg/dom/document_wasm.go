//go:build js

package dom

import (
	"fmt"

	dom "github.com/djthorpe/go-dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type document struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *document) Body() dom.Element {
	return NewNode(this.Get("body")).(dom.Element)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *document) CreateElement(name string) dom.Element {
	return NewNode(this.Call("createElement", name)).(dom.Element)
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *document) String() string {
	str := "<DOMDocument"
	str += fmt.Sprint(" ", this.node)
	return str + ">"
}
