package bs5

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// CONSTANTS

// Bootstrap 5 color variants
type ColorVariant string

const (
	ColorPrimary   ColorVariant = "primary"
	ColorSecondary ColorVariant = "secondary"
	ColorSuccess   ColorVariant = "success"
	ColorDanger    ColorVariant = "danger"
	ColorWarning   ColorVariant = "warning"
	ColorInfo      ColorVariant = "info"
	ColorLight     ColorVariant = "light"
	ColorDark      ColorVariant = "dark"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Badge struct {
	dom.Element
	color ColorVariant
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Badge creates a badge span element with the specified color variant
func (app *App) Badge(color ColorVariant, children ...dom.Node) *Badge {
	span := app.CreateElement("span")
	span.AddClass("badge")
	span.AddClass(fmt.Sprintf("text-bg-%s", color))

	for _, child := range children {
		span.AppendChild(child)
	}

	return &Badge{
		Element: span,
		color:   color,
	}
}

////////////////////////////////////////////////////////////////////////
// METHODS

// Color returns the badge's color variant
func (b *Badge) Color() ColorVariant {
	return b.color
}

// AddClass adds a CSS class to the badge
func (b *Badge) AddClass(className string) *Badge {
	b.Element.AddClass(className)
	return b
}

// RemoveClass removes a CSS class from the badge
func (b *Badge) RemoveClass(className string) *Badge {
	b.Element.RemoveClass(className)
	return b
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (b *Badge) String() string {
	return fmt.Sprintf("<bs5-badge color=%s>", b.color)
}
