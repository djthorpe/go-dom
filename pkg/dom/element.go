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
	attrs []dom.Attr
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
	return nextSibling(this.parent, this)
}

func (this *element) PreviousSibling() dom.Node {
	return previousSibling(this.parent, this)
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
	result := make([]dom.Attr, len(this.attrs))
	for i, attr := range this.attrs {
		result[i] = attr
	}
	return result
}

func (this *element) HasAttributes() bool {
	return len(this.attrs) > 0
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *element) v() *node {
	return this.node
}
