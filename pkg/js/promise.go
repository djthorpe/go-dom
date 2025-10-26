//go:build js && wasm

package js

import (
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Promise wraps a JavaScript Promise object
type Promise struct {
	value Value
}

// PromiseCallback is a function that handles Promise resolution or rejection
type PromiseCallback func(value Value) interface{}

///////////////////////////////////////////////////////////////////////////////
// CONSTRUCTORS

// NewPromiseFromValue creates a Promise wrapper from a Value
func NewPromiseFromValue(value Value) *Promise {
	return &Promise{value: value}
}

// PromiseResolve creates a Promise that is resolved with the given value
func PromiseResolve(value interface{}) *Promise {
	jsValue := Global().Get("Promise").Call("resolve", Unwrap(value))
	return &Promise{value: jsValue}
}

// PromiseReject creates a Promise that is rejected with the given reason
func PromiseReject(reason interface{}) *Promise {
	jsValue := Global().Get("Promise").Call("reject", Unwrap(reason))
	return &Promise{value: jsValue}
}

// PromiseAll returns a Promise that resolves when all promises in the array resolve,
// or rejects with the reason of the first promise that rejects
func PromiseAll(promises ...*Promise) *Promise {
	array := NewArray(len(promises))
	for i, p := range promises {
		array.SetIndex(i, p.value)
	}
	jsValue := Global().Get("Promise").Call("all", array)
	return &Promise{value: jsValue}
}

// PromiseAllSettled returns a Promise that resolves after all promises have settled
// (each may resolve or reject)
func PromiseAllSettled(promises ...*Promise) *Promise {
	array := NewArray(len(promises))
	for i, p := range promises {
		array.SetIndex(i, p.value)
	}
	jsValue := Global().Get("Promise").Call("allSettled", array)
	return &Promise{value: jsValue}
}

// PromiseAny returns a Promise that resolves as soon as any promise resolves,
// or rejects if all promises reject
func PromiseAny(promises ...*Promise) *Promise {
	array := NewArray(len(promises))
	for i, p := range promises {
		array.SetIndex(i, p.value)
	}
	jsValue := Global().Get("Promise").Call("any", array)
	return &Promise{value: jsValue}
}

// PromiseRace returns a Promise that resolves or rejects as soon as one of the promises
// resolves or rejects, with the value or reason from that promise
func PromiseRace(promises ...*Promise) *Promise {
	array := NewArray(len(promises))
	for i, p := range promises {
		array.SetIndex(i, p.value)
	}
	jsValue := Global().Get("Promise").Call("race", array)
	return &Promise{value: jsValue}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Then attaches fulfillment and rejection handlers to the promise and returns a new promise
// resolving to the return value of the called handler
func (p *Promise) Then(onFulfilled PromiseCallback, onRejected ...PromiseCallback) *Promise {
	fulfilledFunc := js.FuncOf(func(this Value, args []Value) interface{} {
		if len(args) > 0 {
			return onFulfilled(args[0])
		}
		return nil
	})

	var result Value
	if len(onRejected) > 0 && onRejected[0] != nil {
		rejectedFunc := js.FuncOf(func(this Value, args []Value) interface{} {
			if len(args) > 0 {
				return onRejected[0](args[0])
			}
			return nil
		})
		result = p.value.Call("then", fulfilledFunc, rejectedFunc)
		// Note: Don't release these functions as they may be called multiple times
		// The JavaScript runtime will handle garbage collection
	} else {
		result = p.value.Call("then", fulfilledFunc)
	}

	return &Promise{value: result}
}

// Catch attaches a rejection handler callback to the promise and returns a new promise
// resolving to the return value of the callback if it is called, or to its original
// fulfillment value if the promise is instead fulfilled
func (p *Promise) Catch(onRejected PromiseCallback) *Promise {
	rejectedFunc := js.FuncOf(func(this Value, args []Value) interface{} {
		if len(args) > 0 {
			return onRejected(args[0])
		}
		return nil
	})

	result := p.value.Call("catch", rejectedFunc)
	return &Promise{value: result}
}

// Finally attaches a handler that is called when the promise is settled (fulfilled or rejected)
// The handler is called without any arguments
func (p *Promise) Finally(onFinally func()) *Promise {
	finallyFunc := js.FuncOf(func(this Value, args []Value) interface{} {
		onFinally()
		return nil
	})

	result := p.value.Call("finally", finallyFunc)
	return &Promise{value: result}
}

// JSValue returns the underlying Value
func (p *Promise) JSValue() Value {
	return p.value
}

///////////////////////////////////////////////////////////////////////////////
// AWAIT HELPERS

// Await is a helper function that waits for a promise to resolve and returns the result
// This creates a Go channel and blocks until the promise settles
// Returns (value, error) where error is set if the promise was rejected
func (p *Promise) Await() (Value, error) {
	resultChan := make(chan Value, 1)
	errorChan := make(chan error, 1)

	// Set up the then/catch handlers
	p.Then(func(value Value) interface{} {
		resultChan <- value
		return nil
	}).Catch(func(reason Value) interface{} {
		// Convert JS error to Go error
		errorChan <- js.Error{Value: reason}
		return nil
	})

	// Wait for either result or error
	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errorChan:
		return Undefined(), err
	}
}

///////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

// WrapPromise wraps a Value that represents a Promise into a Promise struct
// This is useful when you get a promise from calling JavaScript APIs
func WrapPromise(value Value) *Promise {
	return &Promise{value: value}
}

// IsPromise checks if a Value is a Promise
func IsPromise(value Value) bool {
	if value.Type() != js.TypeObject {
		return false
	}
	promiseConstructor := Global().Get("Promise")
	return value.InstanceOf(promiseConstructor)
}
