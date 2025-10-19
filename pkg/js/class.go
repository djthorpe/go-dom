//go:build js

package js

import (
	"regexp"
	"syscall/js"
)

var (
	reClassName = regexp.MustCompile(`^[A-Z][A-Za-z0-9_]*$`)
)

// Class represents a JavaScript ES6 class constructor.
type Class struct {
	value js.Value
}

// NewClass creates a new empty ES6 class with the given name.
// The class extends the specified parent class (or pass nil for no parent).
// Returns a pointer to a Class wrapper around the class constructor.
func NewClass(name string, parent *Class) *Class {
	var classValue js.Value

	// Check class name is valid
	if !reClassName.MatchString(name) {
		return nil
	}

	if parent == nil {
		// Create a class with no parent
		classDefinition := `
			class ` + name + ` {
				constructor() {
				}
			}
			` + name + `
		`
		classValue = js.Global().Call("eval", classDefinition)
	} else {
		// Create a class that extends the parent
		// We need to make the parent available in the eval scope
		classDefinition := `
			(function(ParentClass) {
				class ` + name + ` extends ParentClass {
					constructor() {
						super();
					}
				}
				return ` + name + `;
			})
		`

		// Execute and pass the parent class
		factory := js.Global().Call("eval", classDefinition)
		classValue = factory.Invoke(parent.value)
	}

	return &Class{value: classValue}
}

// GetClass retrieves an existing class from the global scope by name.
// Returns a pointer to a Class wrapper, or nil if not found.
func GetClass(name string) *Class {
	class := js.Global().Get(name)
	if class.IsUndefined() || class.IsNull() {
		return nil
	}

	// Verify it's actually a function/class (constructors are functions in JS)
	if class.Type() != js.TypeFunction {
		return nil
	}

	return &Class{value: class}
}

// Prototype returns the prototype object of the class.
func (c *Class) Prototype() js.Value {
	return c.value.Get("prototype")
}

// New creates a new instance of the class by calling its constructor.
// The args parameter contains the arguments to pass to the constructor.
func (c *Class) New(args ...any) js.Value {
	return js.Global().Get("Reflect").Call("construct", c.value, args)
}

// Value returns the underlying js.Value representing the class constructor.
func (c *Class) Value() js.Value {
	return c.value
}

// NewFunction creates a new JavaScript function and sets it on the class prototype.
// The name parameter specifies the method name on the prototype.
// Returns the created Function for further manipulation if needed.
func (c *Class) NewFunction(name string, fn CallbackFunc) *Function {
	function := NewFunction(fn)
	c.Prototype().Set(name, function.Value())
	return function
}
