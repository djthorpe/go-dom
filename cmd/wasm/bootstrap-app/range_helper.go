//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/djthorpe/go-dom"
)

// setupRangeValueDisplay sets up a range input to update a display element
func setupRangeValueDisplay(rangeInput dom.Element, displayId string) {
	// Use syscall/js to set up the event listener directly
	rangeInput.AddEventListener("input", func(node dom.Node) {
		// Access the DOM element directly via JavaScript
		doc := js.Global().Get("document")
		rangeEl := doc.Call("getElementById", rangeInput.GetAttribute("id").Value())
		displayEl := doc.Call("getElementById", displayId)

		if !rangeEl.IsUndefined() && !rangeEl.IsNull() && !displayEl.IsUndefined() && !displayEl.IsNull() {
			// Get the current value from the range input property (not attribute)
			value := rangeEl.Get("value").String()
			// Update the display element's text content
			displayEl.Set("textContent", value)
		}
	})
}
