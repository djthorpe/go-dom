package main

import (
	"fmt"

	// Modules
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func main() {
	doc := Document()
	fmt.Println("Helloworld! Document=", doc)
}
