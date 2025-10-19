//go:build !js

package dom

import (
	"bytes"
	"fmt"
	"strings"

	dom "github.com/djthorpe/go-dom"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	*node
	attrs map[string]dom.Attr
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *element) String() string {
	str := "<DOMElement"
	str += fmt.Sprint(" ", this.node)
	return str + ">"
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *element) NextSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return this.parent.(nodevalue).nextChild(this)
	}
}

func (this *element) PreviousSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return this.parent.(nodevalue).previousChild(this)
	}
}

func (this *element) InnerHTML() string {
	buf := new(bytes.Buffer)
	for child := this.FirstChild(); child != nil; child = child.NextSibling() {
		child.(nodevalue).write(buf)
	}
	return buf.String()
}

func (this *element) OuterHTML() string {
	buf := new(bytes.Buffer)
	this.write(buf)
	return buf.String()
}

func (this *element) TagName() string {
	if name := this.NodeName(); strings.HasPrefix(name, "#") {
		return name
	} else {
		return strings.ToUpper(name)
	}
}

func (this *element) Attributes() []dom.Attr {
	result := make([]dom.Attr, 0, len(this.attrs))
	for _, attr := range this.attrs {
		result = append(result, attr)
	}
	return result
}

func (this *element) HasAttributes() bool {
	return len(this.attrs) > 0
}

func (this *element) Style() dom.Style {
	// Not implemented for non-WASM builds
	return nil
}

func (this *element) SetAttribute(name, value string) dom.Attr {
	attr := this.document.CreateAttribute(name)
	attr.SetValue(value)
	attr.(nodevalue).v().parent = this
	this.attrs[name] = attr
	return attr
}

func (this *element) GetAttribute(name string) dom.Attr {
	return this.attrs[name]
}

func (this *element) AddClass(className string) {
	// Get current class attribute
	classAttr := this.GetAttribute("class")
	if classAttr == nil {
		// No class attribute exists, create one
		this.SetAttribute("class", className)
		return
	}

	// Parse existing classes
	classes := strings.Fields(classAttr.Value())

	// Check if class already exists
	for _, c := range classes {
		if c == className {
			return // Class already present
		}
	}

	// Add the new class
	classes = append(classes, className)
	this.SetAttribute("class", strings.Join(classes, " "))
}

func (this *element) RemoveClass(className string) {
	// Get current class attribute
	classAttr := this.GetAttribute("class")
	if classAttr == nil {
		return // No classes to remove
	}

	// Parse existing classes
	classes := strings.Fields(classAttr.Value())

	// Filter out the class to remove
	filtered := make([]string, 0, len(classes))
	for _, c := range classes {
		if c != className {
			filtered = append(filtered, c)
		}
	}

	// Update or remove the class attribute
	if len(filtered) == 0 {
		delete(this.attrs, "class")
	} else {
		this.SetAttribute("class", strings.Join(filtered, " "))
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *element) v() *node {
	return this.node
}
