package main

import (
	// Modules
	dom "github.com/djthorpe/go-dom/pkg/dom"
)

func main() {
	document := dom.GetWindow().Document()
	body := document.Body()
	h1 := document.CreateElement("h1")
	h1.AppendChild(document.CreateTextNode("Hello, World!"))
	body.AppendChild(h1)
}
