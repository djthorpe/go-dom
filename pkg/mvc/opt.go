package mvc

import (
	"slices"
	"strings"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////////////////
// TYPES

// Opt is a function which can apply options to a view
type Opt func(*opt) error

// opt is a private struct which holds options
type opt struct {
	name  string
	id    string
	class []string
	attr  map[string]string
}

/////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func applyOpts(element Element, opts ...Opt) error {
	var o opt

	// Existing name from element
	if element == nil {
		return ErrBadParameter.Withf("Missing Element")
	} else if name := element.GetAttribute(DataComponentAttrKey); name == "" {
		return ErrBadParameter.Withf("Element does not have a valid component name")
	} else {
		o.id = element.ID()
		o.name = element.GetAttribute(DataComponentAttrKey)
		o.class = element.ClassList().Values()
		o.attr = make(map[string]string)

		for _, attr := range element.Attributes() {
			attrName := attr.Name()

			// Skip id, class, and data-component as they're handled separately
			if attrName != "id" && attrName != "class" && attrName != DataComponentAttrKey {
				o.attr[attrName] = attr.Value()
			}
		}
	}

	// Apply each option to the opt struct
	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return err
		}
	}

	// Apply ID if set
	if o.id != "" {
		element.SetID(o.id)
	}

	// Apply classes if set
	if len(o.class) > 0 {
		element.SetAttribute("class", strings.Join(o.classes(), " "))
	}

	// Apply attributes if set
	for key, value := range o.attr {
		element.SetAttribute(key, value)
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func WithClass(classes ...string) Opt {
	return func(o *opt) error {
		o.class = append(o.class, classes...)
		return nil
	}
}

func WithoutClass(classes ...string) Opt {
	return func(o *opt) error {
		for _, class := range classes {
			o.class = slices.DeleteFunc(o.class, func(c string) bool {
				return c == class
			})
		}
		return nil
	}
}

func WithAttr(key, value string) Opt {
	return func(o *opt) error {
		if o.attr == nil {
			o.attr = make(map[string]string)
		}
		o.attr[key] = value
		return nil
	}
}

func WithoutAttr(keys ...string) Opt {
	return func(o *opt) error {
		if o.attr == nil {
			return nil
		}
		for _, key := range keys {
			delete(o.attr, key)
		}
		return nil
	}
}

func WithID(id string) Opt {
	return func(o *opt) error {
		o.id = id
		return nil
	}
}

/////////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Return space-separated class string without duplicates
func (o *opt) classes() []string {
	return slices.Compact(o.class)
}
