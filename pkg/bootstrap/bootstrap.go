package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type app struct {
	Document

	// The root element in which the application is contained
	root Element
}

// Ensure that app implements Application interface
var _ Application = (*app)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	AppIdentifier = "wasmbuild-bootstrap-app"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new bootstrap application
func New() *app {
	// Get a document object
	doc := dom.GetWindow().Document()

	// Append a div to the body which will contain the application
	root := doc.CreateElement("DIV")
	root.SetAttribute("id", AppIdentifier)
	doc.Body().AppendChild(root)

	// Return the document
	return &app{
		Document: doc,
		root:     root,
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (app *app) Root() Element {
	return app.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (app *app) Append(children ...any) Application {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = app.Document.CreateTextNode(str)
		}

		// Append to root
		app.root.AppendChild(child.(Node))
	}

	// Return the application for chaining
	return app
}
