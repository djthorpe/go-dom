package mvc

import (
	// Namespace imports

	. "github.com/djthorpe/go-wasmbuild"
	"github.com/djthorpe/go-wasmbuild/pkg/dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// View represents a UI component in the interface
type View interface {
	// Return the view name
	Name() string

	// Return the view ID, if set
	ID() string

	// Return the view's root element
	Root() Element

	// Set the view's body element
	Body(Element)

	// Empty the view's body element, and return the view
	Empty() View

	// Insert text, Element or View children at the top of the view body
	Insert(children ...any) View

	// Append text, Element or View children at the bottom of the view body
	Append(children ...any) View

	// Add an event listener to the view's root element
	AddEventListener(event string, handler func(Node)) View

	// Apply options to the view
	Apply(opts ...Opt) View
}

// Implementation of View interface
type view struct {
	name string
	root Element
	body Element
}

// Ensure that view implements View interface
var _ View = (*view)(nil)

// Constructor function for views
type ViewConstructorFunc func(Element) View

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// The attribute key which identifies a wasmbuild component
	DataComponentAttrKey = "data-wasmbuild-component"
)

var (
	// All the registered views
	views = make(map[string]ViewConstructorFunc, 50)
)

// RegisterView registers a view constructor function for a given name,
// so that the view can be created on-demand
func RegisterView(name string, constructor ViewConstructorFunc) {
	if _, exists := views[name]; exists {
		panic("View already registered: " + name)
	}
	views[name] = constructor
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new empty view, applying any options to it
func NewView(name string, tagName string, opts ...Opt) View {
	if _, exists := views[name]; !exists {
		panic("View not registered: " + name)
	}

	// Create the view
	v := &view{
		name: name,
		root: elementFactory(tagName),
	}

	// Set the data-wasmbuild-component attribute
	v.root.SetAttribute(DataComponentAttrKey, name)

	// Apply options to the view
	if len(opts) > 0 {
		if err := applyOpts(v.root, opts...); err != nil {
			panic(err)
		}
	}

	// Return the view
	return v
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *view) String() string {
	return v.Root().OuterHTML()
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (v *view) Name() string {
	return v.name
}

func (v *view) ID() string {
	return v.root.ID()
}

func (v *view) Root() Element {
	return v.root
}

func (v *view) Body(element Element) {
	v.body = element
}

func (v *view) Empty() View {
	element := v.body
	if element == nil {
		element = v.root
	}
	element.SetInnerHTML("")
	return v
}

func (v *view) Insert(children ...any) View {
	target := v.body
	if target == nil {
		target = v.root
	}
	firstChild := target.FirstChild()
	for _, child := range children {
		if firstChild == nil {
			target.AppendChild(NodeFromAny(child))
		} else {
			target.InsertBefore(NodeFromAny(child), firstChild)
		}
	}
	return v
}

func (v *view) Append(children ...any) View {
	target := v.body
	if target == nil {
		target = v.root
	}
	for _, child := range children {
		target.AppendChild(NodeFromAny(child))
	}
	return v
}

func (v *view) AddEventListener(event string, handler func(Node)) View {
	v.root.AddEventListener(event, handler)
	return v
}

func (v *view) Apply(opts ...Opt) View {
	if err := applyOpts(v.root, opts...); err != nil {
		panic(err)
	}
	return v
}

///////////////////////////////////////////////////////////////////////////////
// UTILITY METHODS

// NodeFromAny returns a Node from a string, Element, or View
// or returns nil if the type is unsupported
func NodeFromAny(child any) Node {
	switch c := child.(type) {
	case string:
		return textFactory(c)
	case Element:
		return c
	case View:
		return c.Root()
	default:
		panic(ErrInternalAppError.Withf("NodeFromAny: unsupported: %T", child))
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Create a new DOM element to be attached to a view
func elementFactory(tagName string) Element {
	return dom.GetWindow().Document().CreateElement(tagName)
}

// Create a new DOM text node to be attached to a view
func textFactory(text string) Node {
	return dom.GetWindow().Document().CreateTextNode(text)
}
