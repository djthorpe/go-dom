package main

import (
	"syscall/js"

	// Modules
	dom "github.com/djthorpe/go-dom/pkg/dom"
)

func main() {
	// Define a custom element class
	defineCustomElement()

	// Now we can use the Go DOM wrapper to create custom elements!
	document := dom.GetWindow().Document()
	body := document.Body()

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

func defineCustomElement() {
	// Create a minimal ES6 class with just the constructor
	// Then add methods programmatically to the prototype
	classDefinition := `
		class MyGreeting extends HTMLElement {
			constructor() {
				super();
			}
		}
		MyGreeting
	`

	// Execute the class definition and get a reference to the class
	myGreetingClass := js.Global().Call("eval", classDefinition)

	// Get the prototype
	prototype := myGreetingClass.Get("prototype")

	// Add connectedCallback method
	prototype.Set("connectedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

		// Create wrapper div
		wrapper := js.Global().Get("document").Call("createElement", "div")
		wrapper.Set("innerHTML", "Hello, <strong>"+nameStr+"</strong>!")

		// Create style element (scoped to shadow DOM)
		styleElem := js.Global().Get("document").Call("createElement", "style")
		styleElem.Set("textContent", `
			div {
				border: 2px solid blue;
				padding: 10px;
				margin: 10px;
				border-radius: 5px;
				background-color: #f0f8ff;
			}
		`)

		// Append style and content to shadow DOM
		shadow.Call("appendChild", styleElem)
		shadow.Call("appendChild", wrapper)

		return nil
	}))

	// Add disconnectedCallback method
	prototype.Set("disconnectedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("log", "Custom element removed from DOM")
		return nil
	}))

	// Add attributeChangedCallback method
	prototype.Set("attributeChangedCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) >= 3 {
			attrName := args[0].String()
			oldValue := args[1]
			newValue := args[2].String()

			js.Global().Get("console").Call("log", "Attribute changed:", attrName, "from", oldValue, "to", newValue)

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
	}))

	// Define observedAttributes as a static getter
	observedAttrs := js.Global().Get("Array").New()
	observedAttrs.Call("push", "name")

	// Use Object.defineProperty to create a static getter
	js.Global().Get("Object").Call("defineProperty", myGreetingClass, "observedAttributes", map[string]interface{}{
		"get": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return observedAttrs
		}),
	})

	// Register the custom element
	js.Global().Get("customElements").Call("define", "my-greeting", myGreetingClass)
}
