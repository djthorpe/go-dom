//go:build wasm

package js

import (
	"syscall/js"

	// Namespace imports
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type fetchpromise struct {
	js.Value
}

type httpresponse struct {
	js.Value
}

///////////////////////////////////////////////////////////////////////////////
// FETCH API

// Fetch performs an HTTP request using the JavaScript Fetch API.
// Returns a FetchPromise that can be used to chain then/catch callbacks.
//
// Example:
//
//	Fetch("https://api.example.com/data").
//	    Then(func(response FetchResponse) {
//	        // Handle response
//	    }).
//	    Catch(func(err js.Value) {
//	        // Handle error
//	    })
func Fetch(url string, options ...js.Value) *fetchpromise {
	var opts js.Value
	if len(options) > 0 {
		opts = options[0]
	} else {
		opts = Undefined()
	}

	// Return the fetch promise
	return &fetchpromise{
		js.Global().Call("fetch", url, opts),
	}
}

///////////////////////////////////////////////////////////////////////////////
// FETCH PROMISE METHODS

// Then adds a callback for when the fetch succeeds.
// The callback receives a FetchResponse wrapping the JavaScript Response object.
// Returns the FetchPromise for chaining.
func (p *fetchpromise) Then(callback func(dom.HTTPResponse)) *fetchpromise {
	// Call the promise's then method
	p.Value = p.Value.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			response := &httpresponse{args[0]}
			callback(response)
		}
		return nil
	}))

	// Return the fetch promise for chaining
	return p
}

// Catch adds a callback for when the fetch fails.
// The callback receives the JavaScript error value.
// Returns the FetchPromise for chaining.
func (p *fetchpromise) Catch(callback func(js.Value)) *fetchpromise {
	// Call the promise's catch method
	p.Value = p.Value.Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			callback(args[0])
		}
		return nil
	}))

	// Return the fetch promise for chaining
	return p
}

// Finally adds a callback that runs regardless of success or failure.
// Returns the FetchPromise for chaining.
func (p *fetchpromise) Finally(callback func()) *fetchpromise {
	// Call the promise's finally method
	p.Value = p.Value.Call("finally", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback()
		return nil
	}))

	// Return the fetch promise for chaining
	return p
}

///////////////////////////////////////////////////////////////////////////////
// FETCH RESPONSE METHODS

	// Wait for either result or error
	select {
	case response := <-resultChan:
		return response, nil
	case err := <-errorChan:
		return FetchResponse{}, err
	}
}

///////////////////////////////////////////////////////////////////////////////
// FETCH RESPONSE METHODS

// Ok returns whether the response was successful (status 200-299).
func (r FetchResponse) Ok() bool {
	return r.value.Get("ok").Bool()
}

// Status returns the HTTP status code.
func (r FetchResponse) Status() int {
	return r.value.Get("status").Int()
}

// StatusText returns the HTTP status message.
func (r FetchResponse) StatusText() string {
	return r.value.Get("statusText").String()
}

// Headers returns the response headers as a JavaScript Headers object.
func (r FetchResponse) Headers() js.Value {
	return r.value.Get("headers")
}

// Text returns a promise that resolves to the response body as text.
func (r FetchResponse) Text(callback func(string)) {
	textPromise := r.value.Call("text")
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			callback(args[0].String())
		}
		return nil
	})
	textPromise.Call("then", jsFunc)
}

// JSON returns a promise that resolves to the response body parsed as JSON.
func (r FetchResponse) JSON(callback func(js.Value)) {
	jsonPromise := r.value.Call("json")
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			callback(args[0])
		}
		return nil
	})
	jsonPromise.Call("then", jsFunc)
}

// Blob returns a promise that resolves to the response body as a Blob.
func (r FetchResponse) Blob(callback func(js.Value)) {
	blobPromise := r.value.Call("blob")
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			callback(args[0])
		}
		return nil
	})
	blobPromise.Call("then", jsFunc)
}

// ArrayBuffer returns a promise that resolves to the response body as an ArrayBuffer.
func (r FetchResponse) ArrayBuffer(callback func(js.Value)) {
	arrayBufferPromise := r.value.Call("arrayBuffer")
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			callback(args[0])
		}
		return nil
	})
	arrayBufferPromise.Call("then", jsFunc)
}

// Value returns the underlying JavaScript Response value.
func (r FetchResponse) Value() js.Value {
	return r.value
}

///////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

// NewFetchOptions creates a JavaScript object for fetch options.
//
// Example:
//
//	opts := NewFetchOptions().
//	    Method("POST").
//	    Header("Content-Type", "application/json").
//	    Body(`{"name": "value"}`)
type FetchOptions struct {
	obj js.Value
}

// NewFetchOptions creates a new fetch options builder.
func NewFetchOptions() *FetchOptions {
	return &FetchOptions{
		obj: NewObject(),
	}
}

// Method sets the HTTP method (GET, POST, PUT, DELETE, etc.).
func (o *FetchOptions) Method(method string) *FetchOptions {
	o.obj.Set("method", method)
	return o
}

// Header sets a request header.
func (o *FetchOptions) Header(key, value string) *FetchOptions {
	if !o.obj.Get("headers").Truthy() {
		o.obj.Set("headers", NewObject())
	}
	o.obj.Get("headers").Set(key, value)
	return o
}

// Body sets the request body.
func (o *FetchOptions) Body(body string) *FetchOptions {
	o.obj.Set("body", body)
	return o
}

// Mode sets the request mode (cors, no-cors, same-origin, navigate).
func (o *FetchOptions) Mode(mode string) *FetchOptions {
	o.obj.Set("mode", mode)
	return o
}

// Credentials sets the credentials mode (omit, same-origin, include).
func (o *FetchOptions) Credentials(credentials string) *FetchOptions {
	o.obj.Set("credentials", credentials)
	return o
}

// Cache sets the cache mode (default, no-store, reload, no-cache, force-cache, only-if-cached).
func (o *FetchOptions) Cache(cache string) *FetchOptions {
	o.obj.Set("cache", cache)
	return o
}

// Redirect sets the redirect mode (follow, error, manual).
func (o *FetchOptions) Redirect(redirect string) *FetchOptions {
	o.obj.Set("redirect", redirect)
	return o
}

// Referrer sets the referrer.
func (o *FetchOptions) Referrer(referrer string) *FetchOptions {
	o.obj.Set("referrer", referrer)
	return o
}

// Value returns the JavaScript object representing the fetch options.
func (o *FetchOptions) Value() js.Value {
	return o.obj
}
