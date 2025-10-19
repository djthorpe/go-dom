package bs5

import (
	"fmt"

	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Progress struct {
	dom.Element
	bar       dom.Element
	isStacked bool // true if this is a stacked segment
}

type ProgressStacked struct {
	dom.Element
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Progress creates a Bootstrap 5 progress bar component
func (app *App) Progress(value int, label string) *Progress {
	// Create wrapper div
	wrapper := app.CreateElement("div")
	wrapper.AddClass("progress")
	wrapper.SetAttribute("role", "progressbar")
	wrapper.SetAttribute("aria-label", label)
	wrapper.SetAttribute("aria-valuenow", fmt.Sprintf("%d", value))
	wrapper.SetAttribute("aria-valuemin", "0")
	wrapper.SetAttribute("aria-valuemax", "100")

	// Create progress bar
	bar := app.CreateElement("div")
	bar.AddClass("progress-bar")
	bar.SetAttribute("style", fmt.Sprintf("width: %d%%", value))

	wrapper.AppendChild(bar)

	return &Progress{
		Element:   wrapper,
		bar:       bar,
		isStacked: false,
	}
}

// ProgressStacked creates a stacked progress bar container
func (app *App) ProgressStacked(items ...dom.Node) *ProgressStacked {
	// Create wrapper div
	wrapper := app.CreateElement("div")
	wrapper.AddClass("progress-stacked")

	// Add progress items
	for _, item := range items {
		wrapper.AppendChild(item)
	}

	return &ProgressStacked{
		Element: wrapper,
	}
}

// ProgressStackedSegment creates a progress bar segment for use in stacked progress bars
// For stacked bars, the width is set on the wrapper, not the bar
func (app *App) ProgressStackedSegment(value int, label string) *Progress {
	// Create wrapper div
	wrapper := app.CreateElement("div")
	wrapper.AddClass("progress")
	wrapper.SetAttribute("role", "progressbar")
	wrapper.SetAttribute("aria-label", label)
	wrapper.SetAttribute("aria-valuenow", fmt.Sprintf("%d", value))
	wrapper.SetAttribute("aria-valuemin", "0")
	wrapper.SetAttribute("aria-valuemax", "100")
	wrapper.SetAttribute("style", fmt.Sprintf("width: %d%%", value))

	// Create progress bar (no width set here for stacked segments)
	bar := app.CreateElement("div")
	bar.AddClass("progress-bar")

	wrapper.AppendChild(bar)

	return &Progress{
		Element:   wrapper,
		bar:       bar,
		isStacked: true,
	}
}

////////////////////////////////////////////////////////////////////////
// METHODS

// SetValue updates the progress value
func (p *Progress) SetValue(value int) *Progress {
	p.Element.SetAttribute("aria-valuenow", fmt.Sprintf("%d", value))

	if p.isStacked {
		// For stacked segments, set width on the wrapper element
		p.Element.SetAttribute("style", fmt.Sprintf("width: %d%%", value))
	} else {
		// For regular progress bars, set width on the bar element
		p.bar.SetAttribute("style", fmt.Sprintf("width: %d%%", value))
	}
	return p
}

// SetLabel sets the accessible label
func (p *Progress) SetLabel(label string) *Progress {
	p.Element.SetAttribute("aria-label", label)
	return p
}

// SetHeight sets the height of the progress bar
func (p *Progress) SetHeight(height string) *Progress {
	currentStyle := p.Element.GetAttribute("style")
	var style string
	if currentStyle != nil {
		style = currentStyle.Value() + "; "
	}
	style += fmt.Sprintf("height: %s", height)
	p.Element.SetAttribute("style", style)
	return p
}

// SetColor sets the background color of the progress bar
func (p *Progress) SetColor(color ColorVariant) *Progress {
	// Remove existing bg- classes
	p.bar.RemoveClass("bg-primary")
	p.bar.RemoveClass("bg-secondary")
	p.bar.RemoveClass("bg-success")
	p.bar.RemoveClass("bg-danger")
	p.bar.RemoveClass("bg-warning")
	p.bar.RemoveClass("bg-info")
	p.bar.RemoveClass("bg-light")
	p.bar.RemoveClass("bg-dark")

	// Add new color class
	p.bar.AddClass("bg-" + string(color))
	return p
}

// SetStriped sets the striped style
func (p *Progress) SetStriped(striped bool) *Progress {
	if striped {
		p.bar.AddClass("progress-bar-striped")
	} else {
		p.bar.RemoveClass("progress-bar-striped")
	}
	return p
}

// SetAnimated sets the animated style (requires striped to be true)
func (p *Progress) SetAnimated(animated bool) *Progress {
	if animated {
		p.bar.AddClass("progress-bar-striped")
		p.bar.AddClass("progress-bar-animated")
	} else {
		p.bar.RemoveClass("progress-bar-animated")
	}
	return p
}

// ShowLabel displays the value as text inside the progress bar
func (p *Progress) ShowLabel(show bool, app *App) *Progress {
	// Clear existing content
	for child := p.bar.FirstChild(); child != nil; child = p.bar.FirstChild() {
		p.bar.RemoveChild(child)
	}

	if show {
		// Get current value
		valueAttr := p.Element.GetAttribute("aria-valuenow")
		if valueAttr != nil {
			textNode := app.CreateTextNode(valueAttr.Value() + "%")
			p.bar.AppendChild(textNode)
		}
	}
	return p
}

// SetTextColor sets the text color using Bootstrap text-bg utility
func (p *Progress) SetTextColor(color ColorVariant) *Progress {
	// Remove existing text-bg classes
	p.bar.RemoveClass("text-bg-primary")
	p.bar.RemoveClass("text-bg-secondary")
	p.bar.RemoveClass("text-bg-success")
	p.bar.RemoveClass("text-bg-danger")
	p.bar.RemoveClass("text-bg-warning")
	p.bar.RemoveClass("text-bg-info")
	p.bar.RemoveClass("text-bg-light")
	p.bar.RemoveClass("text-bg-dark")

	// Add new text-bg class
	p.bar.AddClass("text-bg-" + string(color))
	return p
}

// AddClass adds a CSS class to the progress wrapper
func (p *Progress) AddClass(className string) *Progress {
	p.Element.AddClass(className)
	return p
}

// RemoveClass removes a CSS class from the progress wrapper
func (p *Progress) RemoveClass(className string) *Progress {
	p.Element.RemoveClass(className)
	return p
}

// AddClass adds a CSS class to the stacked progress wrapper
func (s *ProgressStacked) AddClass(className string) *ProgressStacked {
	s.Element.AddClass(className)
	return s
}

// RemoveClass removes a CSS class from the stacked progress wrapper
func (s *ProgressStacked) RemoveClass(className string) *ProgressStacked {
	s.Element.RemoveClass(className)
	return s
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (p *Progress) String() string {
	return "<bs5-progress>"
}

func (s *ProgressStacked) String() string {
	return "<bs5-progressstacked>"
}
