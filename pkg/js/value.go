//go:build js

package js

import (
	"syscall/js"
)

// Value is an alias for js.Value for convenience.
type Value = js.Value

// Global returns the JavaScript global object (window in browsers, global in Node.js).
func Global() js.Value {
	return js.Global()
}

// Undefined returns the JavaScript undefined value.
func Undefined() js.Value {
	return js.Undefined()
}

// Null returns the JavaScript null value.
func Null() js.Value {
	return js.Null()
}

// ValueOf returns a JavaScript value for the given Go value.
// Supported types: string, int, float64, bool, and nil.
func ValueOf(x interface{}) js.Value {
	return js.ValueOf(x)
}

// NewObject creates a new empty JavaScript object.
func NewObject() js.Value {
	return js.Global().Get("Object").New()
}

// NewArray creates a new JavaScript array with the given length.
// If no length is specified, creates an empty array.
func NewArray(length ...int) js.Value {
	if len(length) > 0 {
		return js.Global().Get("Array").New(length[0])
	}
	return js.Global().Get("Array").New()
}

// NewMap creates a new JavaScript Map.
func NewMap() js.Value {
	return js.Global().Get("Map").New()
}

// NewSet creates a new JavaScript Set.
func NewSet() js.Value {
	return js.Global().Get("Set").New()
}

// NewPromise creates a new JavaScript Promise with the given executor function.
// The executor function receives resolve and reject functions as js.Value arguments.
func NewPromise(executor func(resolve, reject js.Value)) js.Value {
	executorFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) >= 2 {
			executor(args[0], args[1])
		}
		return nil
	})
	defer executorFunc.Release()
	promise := js.Global().Get("Promise").New(executorFunc)
	executorFunc.Release()
	return promise
}

// NewError creates a new JavaScript Error with the given message.
func NewError(message string) js.Value {
	return js.Global().Get("Error").New(message)
}

// NewDate creates a new JavaScript Date object.
// If no arguments are provided, creates a Date for the current time.
func NewDate(args ...interface{}) js.Value {
	if len(args) > 0 {
		return js.Global().Get("Date").New(args...)
	}
	return js.Global().Get("Date").New()
}

// NewRegExp creates a new JavaScript RegExp with the given pattern and flags.
func NewRegExp(pattern string, flags ...string) js.Value {
	if len(flags) > 0 {
		return js.Global().Get("RegExp").New(pattern, flags[0])
	}
	return js.Global().Get("RegExp").New(pattern)
}

// Unwrap extracts the underlying js.Value from objects that have a JSValue() method.
// This is useful for working with DOM nodes and other wrapped JavaScript objects.
func Unwrap(v interface{}) js.Value {
	if unwrapper, ok := v.(interface{ JSValue() js.Value }); ok {
		return unwrapper.JSValue()
	}
	// If the value doesn't have JSValue(), try to convert it directly
	return ValueOf(v)
}
