//go:build !js

package dom

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type event struct {
	eventType string
	target    dom.Node
}

/////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewEvent(eventType string, target dom.Node) dom.Event {
	return &event{eventType, target}
}

///////////////////////////////////////////////////////////////////////////////
// EVENT IMPLEMENTATION

func (e *event) Type() string {
	return e.eventType
}

func (e *event) Target() dom.Node {
	return e.target
}

func (e *event) CurrentTarget() dom.Node {
	return e.target
}

func (e *event) Bubbles() bool {
	return false
}

func (e *event) Cancelable() bool {
	return false
}

func (e *event) DefaultPrevented() bool {
	return false
}

func (e *event) PreventDefault() {
	// No-op for non-WASM builds
}

func (e *event) StopPropagation() {
	// No-op for non-WASM builds
}

func (e *event) StopImmediatePropagation() {
	// No-op for non-WASM builds
}
