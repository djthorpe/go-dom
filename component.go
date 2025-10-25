package dom

// Component interface
type Component interface {
	// Return the component's root element
	Element() Element

	// Insert text, Element or Component children at the top of this component
	Insert(children ...any) Component

	// Append text, Element or Component children to this component
	Append(children ...any) Component
}

// Application interface
type Application interface {
	Component
}
