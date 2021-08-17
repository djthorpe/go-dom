//go:build js

package dom

import (
	"fmt"
	"syscall/js"
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
	return this.Get("innerHTML").String()
}

func (this *element) OuterHTML() string {
	return this.Get("outerHTML").String()
}

/////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *element) v() js.Value {
	return this.node.v()
}
