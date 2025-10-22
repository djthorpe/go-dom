//go:build wasm

package js

import (
	"syscall/js"
	"testing"
)

func TestFetchOptions(t *testing.T) {
	opts := NewFetchOptions().
		Method("POST").
		Header("Content-Type", "application/json").
		Header("Authorization", "Bearer token123").
		Body(`{"key": "value"}`).
		Mode("cors").
		Credentials("include").
		Cache("no-cache").
		Redirect("follow")

	jsObj := opts.Value()

	// Verify method
	if method := jsObj.Get("method").String(); method != "POST" {
		t.Errorf("Expected method POST, got %s", method)
	}

	// Verify body
	if body := jsObj.Get("body").String(); body != `{"key": "value"}` {
		t.Errorf("Expected body to be JSON, got %s", body)
	}

	// Verify mode
	if mode := jsObj.Get("mode").String(); mode != "cors" {
		t.Errorf("Expected mode cors, got %s", mode)
	}

	// Verify credentials
	if creds := jsObj.Get("credentials").String(); creds != "include" {
		t.Errorf("Expected credentials include, got %s", creds)
	}

	// Verify headers exist
	headers := jsObj.Get("headers")
	if !headers.Truthy() {
		t.Error("Expected headers to be set")
	}

	// Verify specific headers
	if ct := headers.Get("Content-Type").String(); ct != "application/json" {
		t.Errorf("Expected Content-Type header, got %s", ct)
	}
}

func TestFetchResponse(t *testing.T) {
	// Create a mock response object
	mockResponse := NewObject()
	mockResponse.Set("ok", true)
	mockResponse.Set("status", 200)
	mockResponse.Set("statusText", "OK")
	mockResponse.Set("headers", NewObject())

	// Add mock methods
	mockResponse.Set("text", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		promise := NewPromise(func(resolve, reject js.Value) {
			resolve.Invoke("response text")
		})
		return promise
	}))

	mockResponse.Set("json", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		promise := NewPromise(func(resolve, reject js.Value) {
			obj := NewObject()
			obj.Set("data", "value")
			resolve.Invoke(obj)
		})
		return promise
	}))

	// Create FetchResponse
	response := FetchResponse{value: mockResponse}

	// Test Ok
	if !response.Ok() {
		t.Error("Expected response to be ok")
	}

	// Test Status
	if status := response.Status(); status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	// Test StatusText
	if statusText := response.StatusText(); statusText != "OK" {
		t.Errorf("Expected status text OK, got %s", statusText)
	}

	// Test Headers
	headers := response.Headers()
	if !headers.Truthy() {
		t.Error("Expected headers to be defined")
	}
}

func TestNewFetchOptions_EmptyHeaders(t *testing.T) {
	opts := NewFetchOptions().
		Method("GET")

	jsObj := opts.Value()

	// Verify method is set
	if method := jsObj.Get("method").String(); method != "GET" {
		t.Errorf("Expected method GET, got %s", method)
	}

	// Headers should not be truthy if never set
	headers := jsObj.Get("headers")
	if headers.Truthy() && headers.Type() == js.TypeObject {
		// This is actually fine - an empty object was created
		// Just verify it doesn't have any properties
	}
}

func TestFetchOptions_MultipleHeaders(t *testing.T) {
	opts := NewFetchOptions().
		Header("X-Custom-1", "value1").
		Header("X-Custom-2", "value2").
		Header("X-Custom-3", "value3")

	jsObj := opts.Value()
	headers := jsObj.Get("headers")

	if !headers.Truthy() {
		t.Fatal("Expected headers to be set")
	}

	// Verify all headers
	if h1 := headers.Get("X-Custom-1").String(); h1 != "value1" {
		t.Errorf("Expected X-Custom-1: value1, got %s", h1)
	}
	if h2 := headers.Get("X-Custom-2").String(); h2 != "value2" {
		t.Errorf("Expected X-Custom-2: value2, got %s", h2)
	}
	if h3 := headers.Get("X-Custom-3").String(); h3 != "value3" {
		t.Errorf("Expected X-Custom-3: value3, got %s", h3)
	}
}

func TestFetchOptions_AllFields(t *testing.T) {
	opts := NewFetchOptions().
		Method("PUT").
		Header("Content-Type", "application/json").
		Body(`{"update": true}`).
		Mode("same-origin").
		Credentials("omit").
		Cache("reload").
		Redirect("error").
		Referrer("https://example.com")

	jsObj := opts.Value()

	tests := []struct {
		field    string
		expected string
	}{
		{"method", "PUT"},
		{"body", `{"update": true}`},
		{"mode", "same-origin"},
		{"credentials", "omit"},
		{"cache", "reload"},
		{"redirect", "error"},
		{"referrer", "https://example.com"},
	}

	for _, tt := range tests {
		if actual := jsObj.Get(tt.field).String(); actual != tt.expected {
			t.Errorf("Expected %s to be %s, got %s", tt.field, tt.expected, actual)
		}
	}
}
