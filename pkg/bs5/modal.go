package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Modal struct {
	dom.Element
	dialog  dom.Element
	content dom.Element
	header  *ModalHeader
	body    *ModalBody
	footer  *ModalFooter
	id      string
}

type ModalHeader struct {
	dom.Element
	modal *Modal
}

type ModalBody struct {
	dom.Element
	modal *Modal
}

type ModalFooter struct {
	dom.Element
	modal *Modal
}

type ModalSize string

////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	ModalSizeSmall  ModalSize = "sm"
	ModalSizeMedium ModalSize = ""
	ModalSizeLarge  ModalSize = "lg"
	ModalSizeXLarge ModalSize = "xl"
)

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Modal creates a Bootstrap 5 modal component
func (app *App) Modal(id string) *Modal {
	// Create modal structure
	// <div class="modal fade" tabindex="-1">
	modal := app.CreateElement("div")
	modal.AddClass("modal")
	modal.AddClass("fade")
	modal.SetAttribute("id", id)
	modal.SetAttribute("tabindex", "-1")
	modal.SetAttribute("aria-hidden", "true")

	// <div class="modal-dialog">
	dialog := app.CreateElement("div")
	dialog.AddClass("modal-dialog")

	// <div class="modal-content">
	content := app.CreateElement("div")
	content.AddClass("modal-content")

	// Build structure
	dialog.AppendChild(content)
	modal.AppendChild(dialog)

	return &Modal{
		Element: modal,
		dialog:  dialog,
		content: content,
		id:      id,
	}
}

////////////////////////////////////////////////////////////////////////
// METHODS

// SetSize sets the modal size
func (m *Modal) SetSize(size ModalSize) *Modal {
	// Remove existing size classes
	m.dialog.RemoveClass("modal-sm")
	m.dialog.RemoveClass("modal-lg")
	m.dialog.RemoveClass("modal-xl")

	// Add new size class if not medium (default)
	if size != ModalSizeMedium {
		m.dialog.AddClass("modal-" + string(size))
	}
	return m
}

// SetCentered centers the modal vertically
func (m *Modal) SetCentered(centered bool) *Modal {
	if centered {
		m.dialog.AddClass("modal-dialog-centered")
	} else {
		m.dialog.RemoveClass("modal-dialog-centered")
	}
	return m
}

// SetScrollable makes the modal body scrollable
func (m *Modal) SetScrollable(scrollable bool) *Modal {
	if scrollable {
		m.dialog.AddClass("modal-dialog-scrollable")
	} else {
		m.dialog.RemoveClass("modal-dialog-scrollable")
	}
	return m
}

// SetFullscreen makes the modal fullscreen
func (m *Modal) SetFullscreen(fullscreen bool) *Modal {
	if fullscreen {
		m.dialog.AddClass("modal-fullscreen")
	} else {
		m.dialog.RemoveClass("modal-fullscreen")
	}
	return m
}

// Header creates or returns the modal header
func (m *Modal) Header(children ...dom.Node) *ModalHeader {
	if m.header == nil {
		// <div class="modal-header">
		header := m.Element.OwnerDocument().CreateElement("div")
		header.AddClass("modal-header")

		// Add children
		for _, child := range children {
			header.AppendChild(child)
		}

		m.content.AppendChild(header)
		m.header = &ModalHeader{
			Element: header,
			modal:   m,
		}
	}
	return m.header
}

// Body creates or returns the modal body
func (m *Modal) Body(children ...dom.Node) *ModalBody {
	if m.body == nil {
		// <div class="modal-body">
		body := m.Element.OwnerDocument().CreateElement("div")
		body.AddClass("modal-body")

		// Add children
		for _, child := range children {
			body.AppendChild(child)
		}

		m.content.AppendChild(body)
		m.body = &ModalBody{
			Element: body,
			modal:   m,
		}
	}
	return m.body
}

// Footer creates or returns the modal footer
func (m *Modal) Footer(children ...dom.Node) *ModalFooter {
	if m.footer == nil {
		// <div class="modal-footer">
		footer := m.Element.OwnerDocument().CreateElement("div")
		footer.AddClass("modal-footer")

		// Add children
		for _, child := range children {
			footer.AppendChild(child)
		}

		m.content.AppendChild(footer)
		m.footer = &ModalFooter{
			Element: footer,
			modal:   m,
		}
	}
	return m.footer
}

// AddClass adds a CSS class to the modal
func (m *Modal) AddClass(className string) *Modal {
	m.Element.AddClass(className)
	return m
}

// RemoveClass removes a CSS class from the modal
func (m *Modal) RemoveClass(className string) *Modal {
	m.Element.RemoveClass(className)
	return m
}

// AddEventListener adds an event listener to the modal
// Common Bootstrap modal events: show.bs.modal, shown.bs.modal, hide.bs.modal, hidden.bs.modal
func (m *Modal) AddEventListener(eventType string, callback func(dom.Node)) *Modal {
	m.Element.AddEventListener(eventType, callback)
	return m
}

// ID returns the modal's ID
func (m *Modal) ID() string {
	return m.id
}

////////////////////////////////////////////////////////////////////////
// MODAL HEADER METHODS

// AddCloseButton adds a close button to the modal header
func (mh *ModalHeader) AddCloseButton() *ModalHeader {
	// <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	closeBtn := mh.Element.OwnerDocument().CreateElement("button")
	closeBtn.SetAttribute("type", "button")
	closeBtn.AddClass("btn-close")
	closeBtn.SetAttribute("data-bs-dismiss", "modal")
	closeBtn.SetAttribute("aria-label", "Close")

	mh.Element.AppendChild(closeBtn)
	return mh
}

// AddTitle adds a title to the modal header
func (mh *ModalHeader) AddTitle(title string) *ModalHeader {
	// <h5 class="modal-title">
	titleEl := mh.Element.OwnerDocument().CreateElement("h5")
	titleEl.AddClass("modal-title")
	titleEl.AppendChild(mh.Element.OwnerDocument().CreateTextNode(title))

	// Insert title at the beginning
	if mh.Element.FirstChild() != nil {
		mh.Element.InsertBefore(titleEl, mh.Element.FirstChild())
	} else {
		mh.Element.AppendChild(titleEl)
	}
	return mh
}

// AddClass adds a CSS class to the modal header
func (mh *ModalHeader) AddClass(className string) *ModalHeader {
	mh.Element.AddClass(className)
	return mh
}

// RemoveClass removes a CSS class from the modal header
func (mh *ModalHeader) RemoveClass(className string) *ModalHeader {
	mh.Element.RemoveClass(className)
	return mh
}

////////////////////////////////////////////////////////////////////////
// MODAL BODY METHODS

// AddClass adds a CSS class to the modal body
func (mb *ModalBody) AddClass(className string) *ModalBody {
	mb.Element.AddClass(className)
	return mb
}

// RemoveClass removes a CSS class from the modal body
func (mb *ModalBody) RemoveClass(className string) *ModalBody {
	mb.Element.RemoveClass(className)
	return mb
}

////////////////////////////////////////////////////////////////////////
// MODAL FOOTER METHODS

// AddClass adds a CSS class to the modal footer
func (mf *ModalFooter) AddClass(className string) *ModalFooter {
	mf.Element.AddClass(className)
	return mf
}

// RemoveClass removes a CSS class from the modal footer
func (mf *ModalFooter) RemoveClass(className string) *ModalFooter {
	mf.Element.RemoveClass(className)
	return mf
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (m *Modal) String() string {
	return "<bs5-modal>"
}

func (mh *ModalHeader) String() string {
	return "<bs5-modal-header>"
}

func (mb *ModalBody) String() string {
	return "<bs5-modal-body>"
}

func (mf *ModalFooter) String() string {
	return "<bs5-modal-footer>"
}
