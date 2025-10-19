//go:build js

package dom

import (
	"fmt"

	// Packages
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
	return this.Get("innerHTML").String()
}

func (this *element) OuterHTML() string {
	return this.Get("outerHTML").String()
}

func (e *element) TagName() string {
	return e.Get("tagName").String()
}

func (e *element) Attributes() []dom.Attr {
	attrs := e.Get("attributes")
	length := attrs.Get("length").Int()
	result := make([]dom.Attr, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, NewNode(attrs.Call("item", i)).(dom.Attr))
	}
	return result
}

func (e *element) HasAttributes() bool {
	return e.Call("hasAttributes").Bool()
}

func (e *element) SetAttribute(name string, value string) dom.Attr {
	e.Call("setAttribute", name, value)
	return e.GetAttribute(name)
}

func (e *element) GetAttribute(name string) dom.Attr {
	return NewNode(e.Call("getAttribute", name)).(dom.Attr)
}
