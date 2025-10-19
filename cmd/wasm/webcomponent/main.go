package main

import (
	// Modules
	"fmt"

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
	defineCustomElement(document, "Greeting", "my-greeting")

	// Create custom elements using the Go DOM wrapper
	customElem := document.CreateElement("my-greeting")
	customElem.SetAttribute("name", "World")
	customElem.Style().Set("margin", "10px")
	body.AppendChild(customElem)

	customElem2 := document.CreateElement("my-greeting")
	customElem2.SetAttribute("name", "Go WASM")
	customElem2.Style().Set("margin", "10px")
	body.AppendChild(customElem2)
}

// Create a custom element class that extends HTMLElement
func defineCustomElement(doc Document, className, elementName string) {
	myGreetingClass := js.NewClass(className, js.GetClass("HTMLElement"))

	// Add connectedCallback method
	myGreetingClass.NewFunction("connectedCallback", func(this js.Value, args []js.Value) any {
		// Attach shadow DOM
		shadow := this.Call("attachShadow", map[string]any{
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
	myGreetingClass.NewFunction("disconnectedCallback", func(this js.Value, args []js.Value) any {
		fmt.Println("Custom element removed from DOM")
		return nil
	})

	// Add attributeChangedCallback method
	myGreetingClass.NewFunction("attributeChangedCallback", func(this js.Value, args []js.Value) any {
		if len(args) >= 3 {
			attrName := args[0].String()
			oldValue := args[1]
			newValue := args[2].String()
			fmt.Println("Attribute changed:", attrName, "from", oldValue, "to", newValue)

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
	js.Global().Get("customElements").Call("define", elementName, myGreetingClass.Value())
}
