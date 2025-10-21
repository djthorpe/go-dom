//go:build wasm

package dom

import (
	"fmt"
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	*node
	eventListeners map[string][]js.Func // Store event listeners to prevent GC
}

type style struct {
	js.Value
}

var _ dom.Element = (*element)(nil)

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
	return e.GetAttributeNode(name)
}

func (e *element) GetAttributeNode(name string) dom.Attr {
	// Use getAttributeNode to get the Attr object
	attrNode := e.Call("getAttributeNode", name)
	if attrNode.IsNull() {
		return nil
	}
	return NewNode(attrNode).(dom.Attr)
}

func (e *element) HasAttribute(name string) bool {
	return e.Call("hasAttribute", name).Bool()
}

func (e *element) ClassList() dom.TokenList {
	// Return a tokenlist that wraps the real DOM element's classList
	classList := e.Get("classList")
	return &tokenlist{
		classList: classList,
	}
}

func (e *element) GetAttribute(name string) string {
	// Use getAttribute which directly returns a string
	result := e.Call("getAttribute", name)
	if result.IsNull() {
		return ""
	}
	return result.String()
}

func (e *element) AddEventListener(eventType string, callback func(dom.Node)) dom.Element {
	// Initialize event listeners map if needed
	if e.eventListeners == nil {
		e.eventListeners = make(map[string][]js.Func)
	}

	// Create a JS function wrapper
	jsCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			// Create a Node from the event target
			target := args[0].Get("target")
			if !target.IsUndefined() && !target.IsNull() {
				callback(NewNode(target))
			}
		}
		return nil
	})

	// Store the callback to prevent garbage collection
	e.eventListeners[eventType] = append(e.eventListeners[eventType], jsCallback)

	// Add event listener
	e.Call("addEventListener", eventType, jsCallback)

	return e
}

func (e *element) Blur() {
	e.Call("blur")
}

func (e *element) Focus() {
	e.Call("focus")
}

///////////////////////////////////////////////////////////////////////////////
// STYLE METHODS

func (s *style) Get(name string) string {
	return s.Value.Get(name).String()
}

func (s *style) Set(name string, value string) {
	s.Value.Set(name, value)
}
