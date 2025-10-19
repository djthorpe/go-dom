package bs5

import (
	"fmt"

	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Row struct {
	dom.Element
}

type Col struct {
	dom.Element
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Row creates a Bootstrap 5 row (grid row)
func (app *App) Row() *Row {
	div := app.CreateElement("div")
	div.AddClass("row")
	return &Row{Element: div}
}

// Col creates a Bootstrap 5 column (grid column)
// Usage: Col() for auto-width, or use SetSize/SetBreakpoint methods
func (app *App) Col() *Col {
	div := app.CreateElement("div")
	div.AddClass("col")
	return &Col{Element: div}
}

////////////////////////////////////////////////////////////////////////
// ROW METHODS

// SetGutters sets the gutter size between columns
// size: 0-5 (0 = no gutters, 5 = largest gutters)
func (r *Row) SetGutters(size int) *Row {
	if size >= 0 && size <= 5 {
		r.Element.AddClass(fmt.Sprintf("g-%d", size))
	}
	return r
}

// SetGuttersX sets horizontal gutters only
func (r *Row) SetGuttersX(size int) *Row {
	if size >= 0 && size <= 5 {
		r.Element.AddClass(fmt.Sprintf("gx-%d", size))
	}
	return r
}

// SetGuttersY sets vertical gutters only
func (r *Row) SetGuttersY(size int) *Row {
	if size >= 0 && size <= 5 {
		r.Element.AddClass(fmt.Sprintf("gy-%d", size))
	}
	return r
}

// AddClass adds a CSS class to the row
func (r *Row) AddClass(className string) *Row {
	r.Element.AddClass(className)
	return r
}

// RemoveClass removes a CSS class from the row
func (r *Row) RemoveClass(className string) *Row {
	r.Element.RemoveClass(className)
	return r
}

// AppendChild adds a child node to the row
func (r *Row) AppendChild(child dom.Node) *Row {
	r.Element.AppendChild(child)
	return r
}

////////////////////////////////////////////////////////////////////////
// COLUMN METHODS

// SetSize sets the column width (1-12)
// Use 0 or "auto" for auto-sizing
func (c *Col) SetSize(size interface{}) *Col {
	// Remove existing col classes
	c.Element.RemoveClass("col")
	for i := 1; i <= 12; i++ {
		c.Element.RemoveClass(fmt.Sprintf("col-%d", i))
	}
	c.Element.RemoveClass("col-auto")

	switch v := size.(type) {
	case int:
		if v == 0 {
			c.Element.AddClass("col-auto")
		} else if v >= 1 && v <= 12 {
			c.Element.AddClass(fmt.Sprintf("col-%d", v))
		}
	case string:
		if v == "auto" {
			c.Element.AddClass("col-auto")
		}
	}
	return c
}

// SetBreakpoint sets responsive column size at a breakpoint
// breakpoint: "sm", "md", "lg", "xl", "xxl"
// size: 1-12, 0 for auto, or "auto"
func (c *Col) SetBreakpoint(breakpoint string, size interface{}) *Col {
	// Remove existing breakpoint classes
	for i := 1; i <= 12; i++ {
		c.Element.RemoveClass(fmt.Sprintf("col-%s-%d", breakpoint, i))
	}
	c.Element.RemoveClass(fmt.Sprintf("col-%s-auto", breakpoint))
	c.Element.RemoveClass(fmt.Sprintf("col-%s", breakpoint))

	switch v := size.(type) {
	case int:
		if v == 0 {
			c.Element.AddClass(fmt.Sprintf("col-%s-auto", breakpoint))
		} else if v >= 1 && v <= 12 {
			c.Element.AddClass(fmt.Sprintf("col-%s-%d", breakpoint, v))
		}
	case string:
		if v == "auto" {
			c.Element.AddClass(fmt.Sprintf("col-%s-auto", breakpoint))
		} else if v == "" {
			c.Element.AddClass(fmt.Sprintf("col-%s", breakpoint))
		}
	}
	return c
}

// SetOffset sets the column offset (1-11)
func (c *Col) SetOffset(offset int) *Col {
	if offset >= 1 && offset <= 11 {
		c.Element.AddClass(fmt.Sprintf("offset-%d", offset))
	}
	return c
}

// SetOffsetBreakpoint sets responsive column offset at a breakpoint
func (c *Col) SetOffsetBreakpoint(breakpoint string, offset int) *Col {
	if offset >= 1 && offset <= 11 {
		c.Element.AddClass(fmt.Sprintf("offset-%s-%d", breakpoint, offset))
	}
	return c
}

// AddClass adds a CSS class to the column
func (c *Col) AddClass(className string) *Col {
	c.Element.AddClass(className)
	return c
}

// RemoveClass removes a CSS class from the column
func (c *Col) RemoveClass(className string) *Col {
	c.Element.RemoveClass(className)
	return c
}

// AppendChild adds a child node to the column
func (c *Col) AppendChild(child dom.Node) *Col {
	c.Element.AppendChild(child)
	return c
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (r *Row) String() string {
	return "<bs5-row>"
}

func (c *Col) String() string {
	return "<bs5-col>"
}
