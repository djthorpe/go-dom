//go:build js

package dom

import (
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
	return this.OuterHTML()
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

func (e *element) RemoveAttribute(name string) {
	e.Call("removeAttribute", name)
}

func (e *element) RemoveAttributeNode(attr dom.Attr) {
	if attr == nil {
		return
	}

	// Wrap in a deferred recover to handle JS errors gracefully
	defer func() {
		if r := recover(); r != nil {
			// Error occurred (attribute wasn't attached to this element)
			// Silently ignore
		}
	}()

	attrValue := toJSValue(attr)
	e.Call("removeAttributeNode", attrValue)
}

func (e *element) SetAttributeNode(attr dom.Attr) dom.Attr {
	if attr == nil {
		return nil
	}
	attrValue := toJSValue(attr)
	result := e.Call("setAttributeNode", attrValue)
	if result.IsNull() {
		return nil
	}
	return NewNode(result).(dom.Attr)
}

func (e *element) GetAttributeNames() []string {
	namesArray := e.Call("getAttributeNames")
	length := namesArray.Get("length").Int()
	names := make([]string, 0, length)
	for i := 0; i < length; i++ {
		names = append(names, namesArray.Index(i).String())
	}
	return names
}

func (e *element) GetElementsByClassName(className string) []dom.Element {
	nodeList := e.Call("getElementsByClassName", className)
	length := nodeList.Get("length").Int()
	result := make([]dom.Element, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, NewNode(nodeList.Index(i)).(dom.Element))
	}
	return result
}

func (e *element) GetElementsByTagName(tagName string) []dom.Element {
	nodeList := e.Call("getElementsByTagName", tagName)
	length := nodeList.Get("length").Int()
	result := make([]dom.Element, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, NewNode(nodeList.Index(i)).(dom.Element))
	}
	return result
}

func (e *element) Remove() {
	e.Call("remove")
}

func (e *element) ReplaceWith(nodes ...dom.Node) {
	if len(nodes) == 0 {
		return
	}

	// Convert nodes to JS values
	args := make([]interface{}, len(nodes))
	for i, node := range nodes {
		args[i] = toJSValue(node)
	}

	e.Call("replaceWith", args...)
}

func (e *element) InsertAdjacentElement(position string, element dom.Element) dom.Element {
	if element == nil {
		return nil
	}

	elemValue := toJSValue(element)
	result := e.Call("insertAdjacentElement", position, elemValue)
	if result.IsNull() {
		return nil
	}
	return NewNode(result).(dom.Element)
}

func (e *element) ID() string {
	return e.Get("id").String()
}

func (e *element) SetID(id string) {
	e.Set("id", id)
}

func (e *element) ClassName() string {
	return e.Get("className").String()
}

func (e *element) SetClassName(className string) {
	e.Set("className", className)
}

func (e *element) Children() []dom.Element {
	children := e.Get("children")
	length := children.Get("length").Int()
	result := make([]dom.Element, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, NewNode(children.Index(i)).(dom.Element))
	}
	return result
}

func (e *element) ChildElementCount() int {
	return e.Get("childElementCount").Int()
}

func (e *element) FirstElementChild() dom.Element {
	child := e.Get("firstElementChild")
	if child.IsNull() {
		return nil
	}
	return NewNode(child).(dom.Element)
}

func (e *element) LastElementChild() dom.Element {
	child := e.Get("lastElementChild")
	if child.IsNull() {
		return nil
	}
	return NewNode(child).(dom.Element)
}

func (e *element) NextElementSibling() dom.Element {
	sibling := e.Get("nextElementSibling")
	if sibling.IsNull() {
		return nil
	}
	return NewNode(sibling).(dom.Element)
}

func (e *element) PreviousElementSibling() dom.Element {
	sibling := e.Get("previousElementSibling")
	if sibling.IsNull() {
		return nil
	}
	return NewNode(sibling).(dom.Element)
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
				// Wrap as Element first if possible, then fall back to Node
				// This ensures Component() can find components on the clicked element
				if targetNode := NewNode(target); targetNode != nil {
					callback(targetNode)
				}
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

// RemoveEventListener removes all event listeners of the specified type and releases their resources
func (e *element) RemoveEventListener(eventType string) {
	if e.eventListeners == nil {
		return
	}

	// Get all listeners for this event type
	listeners := e.eventListeners[eventType]
	if listeners == nil {
		return
	}

	// Remove each listener from the DOM and release the js.Func
	for _, jsCallback := range listeners {
		e.Call("removeEventListener", eventType, jsCallback)
		jsCallback.Release()
	}

	// Remove from the map
	delete(e.eventListeners, eventType)
}

// ReleaseEventListeners removes all event listeners and releases their resources
// Call this when discarding the element to prevent memory leaks
func (e *element) ReleaseEventListeners() {
	if e.eventListeners == nil {
		return
	}

	// Remove and release all listeners
	for eventType, listeners := range e.eventListeners {
		for _, jsCallback := range listeners {
			e.Call("removeEventListener", eventType, jsCallback)
			jsCallback.Release()
		}
	}

	// Clear the map
	e.eventListeners = nil
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
