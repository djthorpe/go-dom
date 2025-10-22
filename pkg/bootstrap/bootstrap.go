package bootstrap

import (
	// Packages
	"fmt"

	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type app struct {
	Document
	component
}

// Ensure that app implements Application interface
var _ Application = (*app)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	AppComponent name = "wasmbuild-bootstrap-app"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new bootstrap application
func New() *app {
	// Get a document object
	doc := dom.GetWindow().Document()

	// Append a div to the body which will contain the application
	root := doc.CreateElement("DIV")
	root.SetAttribute("id", string(AppComponent))
	doc.Body().AppendChild(root)

	// Listen for hashchange events
	dom.GetWindow().AddEventListener("hashchange", func(event Event) {
		// Handle the hash change event
		fmt.Println("Hash changed", dom.GetWindow().Location().Hash())
	})

	// Return the document
	return &app{
		Document: doc,
		component: component{
			name: AppComponent,
			root: root,
		},
	}
}
