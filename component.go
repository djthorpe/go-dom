package dom

// Application interface
type Application interface {
	// Return the root element of the application
	Root() Element
}

// Component interface
type Component interface {
	// Return the component's root element
	Element() Element
}
