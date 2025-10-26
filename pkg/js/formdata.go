//go:build js && wasm

package js

import (
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// FormData wraps a JavaScript FormData object
type FormData struct {
	value js.Value
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewFormData creates a new FormData wrapper from a js.Value
func NewFormData(value js.Value) *FormData {
	return &FormData{value: value}
}

// NewFormDataFromElement creates a new FormData object from a form element
func NewFormDataFromElement(formElement js.Value) *FormData {
	formData := js.Global().Get("FormData").New(formElement)
	return &FormData{value: formData}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Keys returns a slice of all keys in the FormData
func (fd *FormData) Keys() []string {
	var keys []string

	// Get the keys iterator
	keysIterator := fd.value.Call("keys")

	// Iterate through all keys
	for {
		entry := keysIterator.Call("next")
		if entry.Get("done").Bool() {
			break
		}
		keys = append(keys, entry.Get("value").String())
	}

	return keys
}

// Values returns a slice of all values in the FormData
func (fd *FormData) Values() []string {
	var values []string

	// Get the values iterator
	valuesIterator := fd.value.Call("values")

	// Iterate through all values
	for {
		entry := valuesIterator.Call("next")
		if entry.Get("done").Bool() {
			break
		}
		values = append(values, entry.Get("value").String())
	}

	return values
}

// Entries returns a map of all key-value pairs in the FormData
func (fd *FormData) Entries() map[string]string {
	result := make(map[string]string)

	// Get the entries iterator
	entries := fd.value.Call("entries")

	// Iterate through all entries
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

// Get returns the value associated with the given key
func (fd *FormData) Get(key string) string {
	return fd.value.Call("get", key).String()
}

// GetAll returns all values associated with the given key
func (fd *FormData) GetAll(key string) []string {
	jsArray := fd.value.Call("getAll", key)
	length := jsArray.Length()

	values := make([]string, length)
	for i := 0; i < length; i++ {
		values[i] = jsArray.Index(i).String()
	}

	return values
}

// Has checks if the FormData contains the given key
func (fd *FormData) Has(key string) bool {
	return fd.value.Call("has", key).Bool()
}

// Set sets the value for the given key
func (fd *FormData) Set(key, value string) {
	fd.value.Call("set", key, value)
}

// Append appends a value to the given key
func (fd *FormData) Append(key, value string) {
	fd.value.Call("append", key, value)
}

// Delete removes the given key and its values
func (fd *FormData) Delete(key string) {
	fd.value.Call("delete", key)
}

// JSValue returns the underlying js.Value
func (fd *FormData) JSValue() js.Value {
	return fd.value
}
