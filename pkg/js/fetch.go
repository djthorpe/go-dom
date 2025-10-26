//go:build js && wasm

package js

import (
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Request wraps a JavaScript Request object
type Request struct {
	value js.Value
}

// Response wraps a JavaScript Response object
type Response struct {
	value js.Value
}

// Headers wraps a JavaScript Headers object
type Headers struct {
	value js.Value
}

// RequestInit represents options for creating a fetch request
type RequestInit struct {
	Method      string            // HTTP method (GET, POST, PUT, DELETE, etc.)
	Headers     map[string]string // Request headers
	Body        interface{}       // Request body (string, FormData, or js.Value)
	Mode        string            // Request mode: "cors", "no-cors", "same-origin", "navigate"
	Credentials string            // Credentials: "omit", "same-origin", "include"
	Cache       string            // Cache mode: "default", "no-store", "reload", "no-cache", "force-cache", "only-if-cached"
	Redirect    string            // Redirect mode: "follow", "error", "manual"
	Referrer    string            // Referrer URL or "no-referrer", "client"
	Integrity   string            // Subresource integrity value
	KeepAlive   bool              // Keep connection alive
	Signal      js.Value          // AbortSignal for cancellation
}

///////////////////////////////////////////////////////////////////////////////
// FETCH

// Fetch performs a fetch request to the specified URL with optional RequestInit options.
// Returns a Promise that resolves to a Response.
func Fetch(url string, init ...*RequestInit) *Promise {
	var promiseValue js.Value
	if len(init) > 0 && init[0] != nil {
		promiseValue = js.Global().Call("fetch", url, buildRequestInit(init[0]))
	} else {
		promiseValue = js.Global().Call("fetch", url)
	}
	return NewPromiseFromValue(promiseValue)
}

// FetchWithRequest performs a fetch request using a Request object.
// Returns a Promise that resolves to a Response.
func FetchWithRequest(request *Request, init ...*RequestInit) *Promise {
	var promiseValue js.Value
	if len(init) > 0 && init[0] != nil {
		promiseValue = js.Global().Call("fetch", request.value, buildRequestInit(init[0]))
	} else {
		promiseValue = js.Global().Call("fetch", request.value)
	}
	return NewPromiseFromValue(promiseValue)
}

///////////////////////////////////////////////////////////////////////////////
// REQUEST

// NewRequest creates a new Request object with the specified URL and options
func NewRequest(url string, init ...*RequestInit) *Request {
	var req js.Value
	if len(init) > 0 && init[0] != nil {
		req = js.Global().Get("Request").New(url, buildRequestInit(init[0]))
	} else {
		req = js.Global().Get("Request").New(url)
	}
	return &Request{value: req}
}

// NewRequestFromValue creates a new Request wrapper from a js.Value
func NewRequestFromValue(value js.Value) *Request {
	return &Request{value: value}
}

// Clone creates a copy of the Request
func (r *Request) Clone() *Request {
	return &Request{value: r.value.Call("clone")}
}

// Method returns the HTTP method of the request
func (r *Request) Method() string {
	return r.value.Get("method").String()
}

// URL returns the URL of the request
func (r *Request) URL() string {
	return r.value.Get("url").String()
}

// Headers returns the Headers object for the request
func (r *Request) Headers() *Headers {
	return &Headers{value: r.value.Get("headers")}
}

// Mode returns the mode of the request
func (r *Request) Mode() string {
	return r.value.Get("mode").String()
}

// Credentials returns the credentials mode of the request
func (r *Request) Credentials() string {
	return r.value.Get("credentials").String()
}

// Cache returns the cache mode of the request
func (r *Request) Cache() string {
	return r.value.Get("cache").String()
}

// Redirect returns the redirect mode of the request
func (r *Request) Redirect() string {
	return r.value.Get("redirect").String()
}

// Referrer returns the referrer of the request
func (r *Request) Referrer() string {
	return r.value.Get("referrer").String()
}

// Integrity returns the subresource integrity value
func (r *Request) Integrity() string {
	return r.value.Get("integrity").String()
}

// BodyUsed returns whether the body has been read
func (r *Request) BodyUsed() bool {
	return r.value.Get("bodyUsed").Bool()
}

// ArrayBuffer returns a Promise that resolves to an ArrayBuffer
func (r *Request) ArrayBuffer() *Promise {
	return NewPromiseFromValue(r.value.Call("arrayBuffer"))
}

// Blob returns a Promise that resolves to a Blob
func (r *Request) Blob() *Promise {
	return NewPromiseFromValue(r.value.Call("blob"))
}

// FormData returns a Promise that resolves to FormData
func (r *Request) FormData() *Promise {
	return NewPromiseFromValue(r.value.Call("formData"))
}

// JSON returns a Promise that resolves to the parsed JSON
func (r *Request) JSON() *Promise {
	return NewPromiseFromValue(r.value.Call("json"))
}

// Text returns a Promise that resolves to the text content
func (r *Request) Text() *Promise {
	return NewPromiseFromValue(r.value.Call("text"))
}

// JSValue returns the underlying js.Value
func (r *Request) JSValue() js.Value {
	return r.value
}

///////////////////////////////////////////////////////////////////////////////
// RESPONSE

// NewResponse creates a new Response object with optional body and init options
func NewResponse(body interface{}, init map[string]interface{}) *Response {
	var resp js.Value
	if body == nil {
		resp = js.Global().Get("Response").New()
	} else if init != nil {
		initObj := buildResponseInit(init)
		resp = js.Global().Get("Response").New(Unwrap(body), initObj)
	} else {
		resp = js.Global().Get("Response").New(Unwrap(body))
	}
	return &Response{value: resp}
}

// NewResponseFromValue creates a new Response wrapper from a js.Value
func NewResponseFromValue(value js.Value) *Response {
	return &Response{value: value}
}

// Clone creates a copy of the Response
func (r *Response) Clone() *Response {
	return &Response{value: r.value.Call("clone")}
}

// Type returns the type of the response (e.g., "basic", "cors")
func (r *Response) Type() string {
	return r.value.Get("type").String()
}

// URL returns the URL of the response
func (r *Response) URL() string {
	return r.value.Get("url").String()
}

// Redirected returns whether the response is the result of a redirect
func (r *Response) Redirected() bool {
	return r.value.Get("redirected").Bool()
}

// Status returns the HTTP status code
func (r *Response) Status() int {
	return r.value.Get("status").Int()
}

// StatusText returns the HTTP status text
func (r *Response) StatusText() string {
	return r.value.Get("statusText").String()
}

// Ok returns true if the status is in the range 200-299
func (r *Response) Ok() bool {
	return r.value.Get("ok").Bool()
}

// Headers returns the Headers object for the response
func (r *Response) Headers() *Headers {
	return &Headers{value: r.value.Get("headers")}
}

// BodyUsed returns whether the body has been read
func (r *Response) BodyUsed() bool {
	return r.value.Get("bodyUsed").Bool()
}

// ArrayBuffer returns a Promise that resolves to an ArrayBuffer
func (r *Response) ArrayBuffer() *Promise {
	return NewPromiseFromValue(r.value.Call("arrayBuffer"))
}

// Blob returns a Promise that resolves to a Blob
func (r *Response) Blob() *Promise {
	return NewPromiseFromValue(r.value.Call("blob"))
}

// FormData returns a Promise that resolves to FormData
func (r *Response) FormData() *Promise {
	return NewPromiseFromValue(r.value.Call("formData"))
}

// JSON returns a Promise that resolves to the parsed JSON
func (r *Response) JSON() *Promise {
	return NewPromiseFromValue(r.value.Call("json"))
}

// Text returns a Promise that resolves to the text content
func (r *Response) Text() *Promise {
	return NewPromiseFromValue(r.value.Call("text"))
}

// JSValue returns the underlying js.Value
func (r *Response) JSValue() js.Value {
	return r.value
}

// Static Response methods

// ResponseError creates a network error Response
func ResponseError() *Response {
	return &Response{value: js.Global().Get("Response").Call("error")}
}

// ResponseRedirect creates a redirect Response with the given URL and status
func ResponseRedirect(url string, status int) *Response {
	return &Response{value: js.Global().Get("Response").Call("redirect", url, status)}
}

///////////////////////////////////////////////////////////////////////////////
// HEADERS

// NewHeaders creates a new Headers object, optionally initialized with values
func NewHeaders(init ...map[string]string) *Headers {
	var headers js.Value
	if len(init) > 0 && init[0] != nil {
		obj := NewObject()
		for key, value := range init[0] {
			obj.Set(key, value)
		}
		headers = js.Global().Get("Headers").New(obj)
	} else {
		headers = js.Global().Get("Headers").New()
	}
	return &Headers{value: headers}
}

// NewHeadersFromValue creates a new Headers wrapper from a js.Value
func NewHeadersFromValue(value js.Value) *Headers {
	return &Headers{value: value}
}

// Append appends a value to an existing header or creates a new header
func (h *Headers) Append(name, value string) {
	h.value.Call("append", name, value)
}

// Delete removes a header
func (h *Headers) Delete(name string) {
	h.value.Call("delete", name)
}

// Get returns the value of a header
func (h *Headers) Get(name string) string {
	val := h.value.Call("get", name)
	if val.IsNull() {
		return ""
	}
	return val.String()
}

// Has checks if a header exists
func (h *Headers) Has(name string) bool {
	return h.value.Call("has", name).Bool()
}

// Set sets a header value, replacing any existing value
func (h *Headers) Set(name, value string) {
	h.value.Call("set", name, value)
}

// Entries returns all header entries as a map
func (h *Headers) Entries() map[string]string {
	result := make(map[string]string)
	entries := h.value.Call("entries")

	for {
		entry := entries.Call("next")
		if entry.Get("done").Bool() {
			break
		}
		value := entry.Get("value")
		key := value.Index(0).String()
		val := value.Index(1).String()
		result[key] = val
	}

	return result
}

// Keys returns all header names
func (h *Headers) Keys() []string {
	var keys []string
	keysIterator := h.value.Call("keys")

	for {
		entry := keysIterator.Call("next")
		if entry.Get("done").Bool() {
			break
		}
		keys = append(keys, entry.Get("value").String())
	}

	return keys
}

// Values returns all header values
func (h *Headers) Values() []string {
	var values []string
	valuesIterator := h.value.Call("values")

	for {
		entry := valuesIterator.Call("next")
		if entry.Get("done").Bool() {
			break
		}
		values = append(values, entry.Get("value").String())
	}

	return values
}

// JSValue returns the underlying js.Value
func (h *Headers) JSValue() js.Value {
	return h.value
}

///////////////////////////////////////////////////////////////////////////////
// ABORT CONTROLLER

// AbortController wraps a JavaScript AbortController object
type AbortController struct {
	value js.Value
}

// NewAbortController creates a new AbortController
func NewAbortController() *AbortController {
	return &AbortController{
		value: js.Global().Get("AbortController").New(),
	}
}

// Signal returns the AbortSignal associated with this controller
func (ac *AbortController) Signal() js.Value {
	return ac.value.Get("signal")
}

// Abort aborts the request associated with this controller
func (ac *AbortController) Abort() {
	ac.value.Call("abort")
}

// JSValue returns the underlying js.Value
func (ac *AbortController) JSValue() js.Value {
	return ac.value
}

///////////////////////////////////////////////////////////////////////////////
// HELPERS

// buildRequestInit converts RequestInit to a JavaScript object
func buildRequestInit(init *RequestInit) js.Value {
	obj := NewObject()

	if init.Method != "" {
		obj.Set("method", init.Method)
	}

	if len(init.Headers) > 0 {
		headers := NewHeaders(init.Headers)
		obj.Set("headers", headers.JSValue())
	}

	if init.Body != nil {
		obj.Set("body", Unwrap(init.Body))
	}

	if init.Mode != "" {
		obj.Set("mode", init.Mode)
	}

	if init.Credentials != "" {
		obj.Set("credentials", init.Credentials)
	}

	if init.Cache != "" {
		obj.Set("cache", init.Cache)
	}

	if init.Redirect != "" {
		obj.Set("redirect", init.Redirect)
	}

	if init.Referrer != "" {
		obj.Set("referrer", init.Referrer)
	}

	if init.Integrity != "" {
		obj.Set("integrity", init.Integrity)
	}

	if init.KeepAlive {
		obj.Set("keepalive", true)
	}

	if !init.Signal.IsUndefined() && !init.Signal.IsNull() {
		obj.Set("signal", init.Signal)
	}

	return obj
}

// buildResponseInit converts a map to a JavaScript object for Response
func buildResponseInit(init map[string]interface{}) js.Value {
	obj := NewObject()

	for key, value := range init {
		switch v := value.(type) {
		case *Headers:
			obj.Set(key, v.JSValue())
		default:
			obj.Set(key, ValueOf(v))
		}
	}

	return obj
}
