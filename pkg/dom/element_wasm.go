//go:build js

package dom

import (
	"fmt"
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-dom"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	*node
}

type style struct {
	js.Value
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

func (e *element) Style() dom.Style {
	return &style{e.Get("style")}
}

func (e *element) SetAttribute(name string, value string) dom.Attr {
	e.Call("setAttribute", name, value)
	return e.GetAttribute(name)
}

func (e *element) GetAttribute(name string) dom.Attr {
	// Use getAttributeNode instead of getAttribute
	// getAttribute returns a string, getAttributeNode returns an Attr object
	attrNode := e.Call("getAttributeNode", name)
	if attrNode.IsNull() {
		return nil
	}
	return NewNode(attrNode).(dom.Attr)
}

///////////////////////////////////////////////////////////////////////////////
// STYLE METHODS

func (s *style) Get(name string) string {
	return s.Value.Get(name).String()
}

func (s *style) Set(name string, value string) {
	s.Value.Set(name, value)
}
