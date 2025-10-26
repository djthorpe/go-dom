package main

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/dom"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func main() {
	// Get the document
	document := dom.GetWindow().Document()

	// Create a DIV
	div := mvc.Div().Append("Hello, World!").AddEventListener("click", func(target Node) {
		// Obtain the view from the node
		if view := mvc.ViewFromNode(target); view != nil {
			fmt.Println("click view:", view.Name())
		} else {
			fmt.Println("click no view")
		}
	})

	// Insert the DIV into the body
	document.Body().InsertBefore(
		div.Root(), document.Body().FirstChild(),
	)

	fmt.Println(document.Body())

	// Run the application
	select {}
}
