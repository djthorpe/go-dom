package main

import (
	"fmt"

	// Modules
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func main() {
	window := GetWindow()
	window.Document().SetTitle("Hello, World!")
	fmt.Println(window.Document().DocumentElement().OuterHTML())
}
