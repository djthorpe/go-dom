package bs5

import (
	// Packages
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Container struct {
	dom.Element
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func (app *App) Container(children ...dom.Node) *Container {
	div := app.CreateElement("div")
	div.AddClass("container")
	for _, child := range children {
		div.AppendChild(child)
	}
	return &Container{div}
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (c *Container) String() string {
	str := "<bs5-container"
	return str + ">"
}
