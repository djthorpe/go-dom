//go:build js && wasm
// +build js,wasm

package dom

import (
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type event struct {
	js.Value
}

var _ dom.Event = (*event)(nil)

/////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewEvent(jsEvent js.Value) dom.Event {
	return &event{jsEvent}
}

///////////////////////////////////////////////////////////////////////////////
// EVENT IMPLEMENTATION

func (e *event) Type() string {
	return e.Get("type").String()
}

func (e *event) Target() dom.Node {
	return NewNode(e.Get("target"))
}

func (e *event) CurrentTarget() dom.Node {
	return NewNode(e.Get("currentTarget"))
}

func (e *event) Bubbles() bool {
	return e.Get("bubbles").Bool()
}

func (e *event) Cancelable() bool {
	return e.Get("cancelable").Bool()
}

func (e *event) DefaultPrevented() bool {
	return e.Get("defaultPrevented").Bool()
}

func (e *event) PreventDefault() {
	e.Call("preventDefault")
}

func (e *event) StopPropagation() {
	e.Call("stopPropagation")
}

func (e *event) StopImmediatePropagation() {
	e.Call("stopImmediatePropagation")
}
