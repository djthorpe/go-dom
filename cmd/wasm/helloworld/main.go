package main

import (
	"fmt"

	// Modules
	. "github.com/djthorpe/go-dom/pkg/bootstrap"
)

func main() {
	doc := NewDocument()
	nav := doc.Append(doc.Nav())
	// ...add things into the nav
	fmt.Println("Document=", doc.Body().ParentElement().OuterHTML())
}
