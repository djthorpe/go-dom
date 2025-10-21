//go:build !wasm

package dom

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	*node
	classlist *tokenlist
	attrs     map[string]dom.Attr
}

var _ dom.Element = (*element)(nil)

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

func (this *element) GetAttribute(name string) string {
	attr := this.attrs[name]
	if attr == nil {
		return ""
	}
	return attr.Value()
}

func (this *element) GetAttributeNode(name string) dom.Attr {
	return this.attrs[name]
}

func (this *element) HasAttribute(name string) bool {
	_, exists := this.attrs[name]
	return exists
}

func (this *element) ClassList() dom.TokenList {
	// Sync classlist with class attribute if needed
	classAttr := this.GetAttribute("class")
	if classAttr != "" {
		// Parse class attribute and update classlist
		classes := strings.Fields(classAttr)
		this.classlist = NewTokenList(classes...)
	} else if this.classlist == nil {
		this.classlist = NewTokenList()
	}
	return this.classlist
}

func (this *element) AddEventListener(eventType string, callback func(dom.Node)) dom.Element {
	// Event listeners are not supported in non-WASM builds
	// This is a no-op since there's no event loop outside the browser
	return this
}

func (this *element) Blur() {
	// Not supported in non-WASM builds
}

func (this *element) Focus() {
	// Not supported in non-WASM builds
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *element) v() *node {
	return this.node
}

func (this *element) write(w io.Writer) (int, error) {
	s := 0

	// Write opening tag with attributes
	tag := "<" + this.node.name

	// Add attributes
	if len(this.attrs) > 0 {
		for _, attr := range this.attrs {
			tag += fmt.Sprintf(" %s=%q", attr.Name(), attr.Value())
		}
	}

	tag += ">"
	if n, err := w.Write([]byte(tag)); err != nil {
		return 0, err
	} else {
		s += n
	}

	// Write children
	for _, child := range this.node.children {
		if n, err := child.(nodevalue).write(w); err != nil {
			return 0, err
		} else {
			s += n
		}
	}

	// Write closing tag
	if n, err := w.Write([]byte("</" + this.node.name + ">")); err != nil {
		return 0, err
	} else {
		s += n
	}

	return s, nil
}
