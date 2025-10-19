package main

import (
	// Modules
	dom "github.com/djthorpe/go-dom/pkg/dom"
	js "github.com/djthorpe/go-dom/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-dom"
)

func main() {

	// Now we can use the Go DOM wrapper to create custom elements!
	document := dom.GetWindow().Document()
	body := document.Body()

	// Define a custom element class
	defineCustomElement(document)

	// Create custom elements using the Go DOM wrapper
	customElem := document.CreateElement("my-greeting")
	customElem.SetAttribute("name", "World")
	customElem.Style().Set("margin", "10px")
	body.AppendChild(customElem)

	customElem2 := document.CreateElement("my-greeting")
	customElem2.SetAttribute("name", "Go WASM")
	customElem2.Style().Set("margin", "10px")
	body.AppendChild(customElem2)

	// Also add a regular h1 for comparison
	h1 := document.CreateElement("h1")
	h1.AppendChild(document.CreateTextNode("Hello from Go DOM!"))
	body.AppendChild(h1)
}

func defineCustomElement(doc Document) {
	// Create a custom element class that extends HTMLElement using the jsutil package
	htmlElement := js.GetClass("HTMLElement")
	myGreetingClass := js.NewClass("MyGreeting", htmlElement)

	// Add connectedCallback method
	myGreetingClass.NewFunction("connectedCallback", func(this js.Value, args []js.Value) interface{} {
		// Attach shadow DOM
		shadow := this.Call("attachShadow", map[string]interface{}{
			"mode": "open",
		})

		// Get the 'name' attribute or default to 'Guest'
		name := this.Call("getAttribute", "name")
		nameStr := "Guest"
		if !name.IsNull() {
			nameStr = name.String()
		}

		// Create wrapper div using Go DOM
		wrapper := doc.CreateElement("div")
		// Get the underlying js.Value for shadow DOM manipulation
		wrapperJS := js.Unwrap(wrapper)
		wrapperJS.Set("innerHTML", "Hello, <strong>"+nameStr+"</strong>!")

		// Create style element (scoped to shadow DOM) using Go DOM
		styleElem := doc.CreateElement("style")
		styleJS := js.Unwrap(styleElem)
		styleJS.Set("textContent", `
			div {
				border: 2px solid blue;
				padding: 10px;
				margin: 10px;
				border-radius: 5px;
				background-color: #f0f8ff;
			}
		`)

		// Append style and content to shadow DOM
		shadow.Call("appendChild", styleJS)
		shadow.Call("appendChild", wrapperJS)

		return nil
	})

	// Add disconnectedCallback method
	myGreetingClass.NewFunction("disconnectedCallback", func(this js.Value, args []js.Value) interface{} {
		console := js.GetConsole()
		console.Log("Custom element removed from DOM")
		return nil
	})

	// Add attributeChangedCallback method
	myGreetingClass.NewFunction("attributeChangedCallback", func(this js.Value, args []js.Value) interface{} {
		if len(args) >= 3 {
			attrName := args[0].String()
			oldValue := args[1]
			newValue := args[2].String()

			console := js.GetConsole()
			console.Log("Attribute changed:", attrName, "from", oldValue, "to", newValue)

			// Re-render when 'name' attribute changes
			if attrName == "name" {
				// Update the shadow DOM content
				shadow := this.Get("shadowRoot")
				if !shadow.IsNull() {
					wrapper := shadow.Call("querySelector", "div")
					if !wrapper.IsNull() {
						wrapper.Set("innerHTML", "Hello, <strong>"+newValue+"</strong>!")
					}
				}
			}
		}
		return nil
	})

	// Define observedAttributes as a static getter
	observedAttrs := js.NewArray()
	observedAttrs.Call("push", "name")

	// Use Object.defineProperty to create a static getter
	observedAttributesGetter := js.NewFunction(func(this js.Value, args []js.Value) interface{} {
		return observedAttrs
	})
	js.Global().Get("Object").Call("defineProperty", myGreetingClass.Value(), "observedAttributes", map[string]interface{}{
		"get": observedAttributesGetter.Value(),
	})

	// Register the custom element
	js.Global().Get("customElements").Call("define", "my-greeting", myGreetingClass.Value())
}
