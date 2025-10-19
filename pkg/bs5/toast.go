package bs5

import (
	"strconv"

	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type Toast struct {
	dom.Element
	header dom.Element
	body   dom.Element
}

type ToastContainer struct {
	dom.Element
}

////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Toast creates a new toast notification
func (app *App) Toast() *Toast {
	toast := app.Document.CreateElement("div")
	toast.AddClass("toast")
	toast.SetAttribute("role", "alert")
	toast.SetAttribute("aria-live", "assertive")
	toast.SetAttribute("aria-atomic", "true")

	return &Toast{
		Element: toast,
	}
}

// ToastContainer creates a container for stacking toasts
func (app *App) ToastContainer() *ToastContainer {
	container := app.Document.CreateElement("div")
	container.AddClass("toast-container")
	container.AddClass("position-fixed")

	return &ToastContainer{
		Element: container,
	}
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - TOAST

// AddHeader adds a header to the toast with title and optional timestamp
func (t *Toast) AddHeader(title string, timestamp string, app *App) *Toast {
	if t.header == nil {
		t.header = app.Document.CreateElement("div")
		t.header.AddClass("toast-header")
		t.Element.AppendChild(t.header)
	}

	// Add title
	strong := app.Document.CreateElement("strong")
	strong.AddClass("me-auto")
	strong.AppendChild(app.CreateTextNode(title))
	t.header.AppendChild(strong)

	// Add timestamp if provided
	if timestamp != "" {
		small := app.Document.CreateElement("small")
		small.AppendChild(app.CreateTextNode(timestamp))
		t.header.AppendChild(small)
	}

	return t
}

// AddCloseButton adds a close button to the header
func (t *Toast) AddCloseButton(app *App) *Toast {
	if t.header == nil {
		t.AddHeader("", "", app)
	}

	btn := app.Document.CreateElement("button")
	btn.SetAttribute("type", "button")
	btn.AddClass("btn-close")
	btn.SetAttribute("data-bs-dismiss", "toast")
	btn.SetAttribute("aria-label", "Close")
	t.header.AppendChild(btn)

	return t
}

// SetBody sets the body content of the toast
func (t *Toast) SetBody(content dom.Node, app *App) *Toast {
	if t.body == nil {
		t.body = app.Document.CreateElement("div")
		t.body.AddClass("toast-body")
		t.Element.AppendChild(t.body)
	} else {
		// Clear existing content
		for t.body.FirstChild() != nil {
			t.body.RemoveChild(t.body.FirstChild())
		}
	}

	t.body.AppendChild(content)
	return t
}

// Body returns the toast body element
func (t *Toast) Body() dom.Element {
	return t.body
}

// Header returns the toast header element
func (t *Toast) Header() dom.Element {
	return t.header
}

// SetAutoHide sets whether the toast should automatically hide
func (t *Toast) SetAutoHide(autohide bool) *Toast {
	if autohide {
		t.Element.SetAttribute("data-bs-autohide", "true")
	} else {
		t.Element.SetAttribute("data-bs-autohide", "false")
	}
	return t
}

// SetDelay sets the delay in milliseconds before hiding the toast
func (t *Toast) SetDelay(delay int) *Toast {
	t.Element.SetAttribute("data-bs-delay", strconv.Itoa(delay))
	return t
}

// SetAnimation sets whether to apply CSS fade transition
func (t *Toast) SetAnimation(animation bool) *Toast {
	if animation {
		t.Element.SetAttribute("data-bs-animation", "true")
	} else {
		t.Element.SetAttribute("data-bs-animation", "false")
	}
	return t
}

// SetColor applies a background color scheme to the toast
func (t *Toast) SetColor(color ColorVariant) *Toast {
	t.AddClass("text-bg-" + string(color))
	t.AddClass("border-0")
	return t
}

// SetPlacement sets the position of the toast using a toast container
func (t *Toast) SetPlacement(placement string) *Toast {
	t.Element.AddClass(placement)
	return t
}

// AddClass adds a CSS class to the toast
func (t *Toast) AddClass(class string) *Toast {
	t.Element.AddClass(class)
	return t
}

// RemoveClass removes a CSS class from the toast
func (t *Toast) RemoveClass(class string) *Toast {
	t.Element.RemoveClass(class)
	return t
}

// AddEventListener adds an event listener to the toast
func (t *Toast) AddEventListener(event string, handler func(dom.Node)) *Toast {
	t.Element.AddEventListener(event, handler)
	return t
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - TOAST CONTAINER

// SetPosition sets the position of the toast container
// Common positions: "top-0 start-0", "top-0 end-0", "bottom-0 start-0", "bottom-0 end-0"
func (tc *ToastContainer) SetPosition(position string) *ToastContainer {
	// Parse space-separated classes
	classes := []string{}
	current := ""
	for _, char := range position {
		if char == ' ' {
			if current != "" {
				classes = append(classes, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	if current != "" {
		classes = append(classes, current)
	}

	// Add all classes
	for _, class := range classes {
		tc.Element.AddClass(class)
	}
	return tc
}

// SetPadding adds padding to the container
func (tc *ToastContainer) SetPadding(padding string) *ToastContainer {
	tc.Element.AddClass(padding)
	return tc
}

// AddToast adds a toast to the container
func (tc *ToastContainer) AddToast(toast *Toast) *ToastContainer {
	tc.Element.AppendChild(toast.Element)
	return tc
}

// AddClass adds a CSS class to the container
func (tc *ToastContainer) AddClass(class string) *ToastContainer {
	tc.Element.AddClass(class)
	return tc
}

// RemoveClass removes a CSS class from the container
func (tc *ToastContainer) RemoveClass(class string) *ToastContainer {
	tc.Element.RemoveClass(class)
	return tc
}
