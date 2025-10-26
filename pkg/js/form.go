//go:build js && wasm

package js

import (
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Form wraps a JavaScript form element
type Form struct {
	value js.Value
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewForm creates a new Form wrapper from a js.Value
func NewForm(value js.Value) *Form {
	return &Form{value: value}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// FormData creates and returns a FormData object from this form
func (f *Form) FormData() *FormData {
	return NewFormDataFromElement(f.value)
}

// Submit submits the form
func (f *Form) Submit() {
	f.value.Call("submit")
}

// Reset resets the form
func (f *Form) Reset() {
	f.value.Call("reset")
}

// JSValue returns the underlying js.Value
func (f *Form) JSValue() js.Value {
	return f.value
}
