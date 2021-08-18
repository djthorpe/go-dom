package main

import (
	"fmt"

	// Modules
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func main() {
	// Create HTML document and set title and body for document
	doc := NewWindow().Document()
	doc.Title().SetInnerHTML("Hello, world!")
	doc.Body().AppendChild(doc.CreateTextNode("Hello, world!"))
	fmt.Println(doc)
}
