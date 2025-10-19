package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type Icon struct {
	dom.Element
}

////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Icon creates a Bootstrap Icon
// See https://icons.getbootstrap.com/ for available icon names
func (app *App) Icon(iconName string) *Icon {
	i := app.CreateElement("i")
	i.AddClass("bi")
	if iconName != "" {
		i.AddClass("bi-" + iconName)
	}

	return &Icon{
		Element: i,
	}
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// SetSize sets the icon size using font-size utility classes
func (i *Icon) SetSize(size string) *Icon {
	// Bootstrap font size utilities: fs-1 (largest) to fs-6 (smallest)
	i.AddClass(size)
	return i
}

// SetColor sets the icon color using text color utility classes
func (i *Icon) SetColor(color string) *Icon {
	// text-primary, text-secondary, text-success, text-danger, text-warning, text-info, etc.
	i.AddClass(color)
	return i
}

// AddClass adds a custom CSS class to the icon
func (i *Icon) AddClass(className string) *Icon {
	i.Element.AddClass(className)
	return i
}
