package dom

import (
	dom "github.com/djthorpe/go-wasmbuild"
)

// componentFactory is the registered factory function for creating components from elements
var componentFactory func(dom.Element) dom.Component

// RegisterComponentFactory registers a factory function that can create components from elements
// This should be called by component libraries (like bootstrap) during initialization
func RegisterComponentFactory(factory func(dom.Element) dom.Component) {
	componentFactory = factory
}

// componentFromElement uses the registered factory to create a component from an element
func componentFromElement(elem dom.Element) dom.Component {
	if componentFactory == nil {
		return nil
	}
	return componentFactory(elem)
}
