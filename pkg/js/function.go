//go:build js

package js

import (
	"syscall/js"
)

// CallbackFunc is the signature for JavaScript callback functions.
// The 'this' parameter is the JavaScript 'this' context.
// The 'args' parameter contains the arguments passed from JavaScript.
type CallbackFunc func(this Value, args []Value) any

// Function represents a JavaScript function.
type Function struct {
	fn js.Func
}

// Value returns the underlying js.Value representing the function.
func (f *Function) Value() Value {
	return f.fn.Value
}

// Call invokes the function with the given arguments.
// The 'this' context will be undefined.
func (f *Function) Call(args ...any) Value {
	return f.fn.Value.Invoke(args...)
}

// CallWithContext invokes the function with a specific 'this' context and arguments.
func (f *Function) CallWithContext(this Value, args ...any) Value {
	return f.fn.Value.Call("call", append([]any{this}, args...)...)
}

// NewFunction creates a new JavaScript function from a Go function.
// The callback receives 'this' context and arguments from JavaScript.
func NewFunction(fn CallbackFunc) *Function {
	return &Function{
		fn: js.FuncOf(fn),
	}
}

// Release releases the function references.
// This should be called when the function is no longer needed to prevent memory leaks.
func (f *Function) Release() {
	f.fn.Release()
}
