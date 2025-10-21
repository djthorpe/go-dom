//go:build js && wasm

package main

import "syscall/js"

// showToast shows a Bootstrap toast using the JavaScript API
func showToast(toastId string) {
	// Get the toast element from DOM
	toastEl := js.Global().Get("document").Call("getElementById", toastId)
	if !toastEl.IsUndefined() && !toastEl.IsNull() {
		// Get Bootstrap Toast class and create instance
		bootstrap := js.Global().Get("bootstrap")
		if !bootstrap.IsUndefined() {
			toastClass := bootstrap.Get("Toast")
			toastInstance := toastClass.Call("getOrCreateInstance", toastEl)
			toastInstance.Call("show")
		}
	}
}
