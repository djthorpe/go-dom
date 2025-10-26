package dom

// Component interface
type Component interface {
	// Return the components type name
	Name() string

	// Return the component's ID
	ID() string

	// Return the component's root element
	Element() Element

	// Empty the component's root element
	Empty() Component

	// Insert text, Element or Component children at the top of this component
	Insert(children ...any) Component

	// Append text, Element or Component children to this component
	Append(children ...any) Component

	// Add an event listener to the component's root element
	AddEventListener(event string, handler func(Node)) Component

	// Apply options to the component
	Apply(opts ...any) Component
}

// Application interface
type Application interface {
	Component
}
