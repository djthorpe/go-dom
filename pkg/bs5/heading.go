package bs5

import (
	"fmt"

	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Heading struct {
	dom.Element
	level int
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// H1 creates an H1 heading element
func (app *App) H1(children ...dom.Node) *Heading {
	return app.heading(1, children...)
}

// H2 creates an H2 heading element
func (app *App) H2(children ...dom.Node) *Heading {
	return app.heading(2, children...)
}

// H3 creates an H3 heading element
func (app *App) H3(children ...dom.Node) *Heading {
	return app.heading(3, children...)
}

// H4 creates an H4 heading element
func (app *App) H4(children ...dom.Node) *Heading {
	return app.heading(4, children...)
}

// H5 creates an H5 heading element
func (app *App) H5(children ...dom.Node) *Heading {
	return app.heading(5, children...)
}

// H6 creates an H6 heading element
func (app *App) H6(children ...dom.Node) *Heading {
	return app.heading(6, children...)
}

// heading is the internal helper to create heading elements
func (app *App) heading(level int, children ...dom.Node) *Heading {
	if level < 1 || level > 6 {
		panic("heading level must be between 1 and 6")
	}

	tag := fmt.Sprintf("h%d", level)
	h := app.CreateElement(tag)

	for _, child := range children {
		h.AppendChild(child)
	}

	return &Heading{
		Element: h,
		level:   level,
	}
}

////////////////////////////////////////////////////////////////////////
// METHODS

// Level returns the heading level (1-6)
func (h *Heading) Level() int {
	return h.level
}

// AddClass adds a CSS class to the heading
func (h *Heading) AddClass(className string) *Heading {
	h.Element.AddClass(className)
	return h
}

// RemoveClass removes a CSS class from the heading
func (h *Heading) RemoveClass(className string) *Heading {
	h.Element.RemoveClass(className)
	return h
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (h *Heading) String() string {
	return fmt.Sprintf("<bs5-h%d>", h.level)
}
