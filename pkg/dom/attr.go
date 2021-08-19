//go:build !js

package dom

import (
	"fmt"
	"html"
	"io"
	"strconv"

	dom "github.com/djthorpe/go-dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type attr struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *attr) String() string {
	str := "<DOMAttribute"
	if name := this.Name(); name != "" {
		str += fmt.Sprintf(" %v=%q", name, this.Value())
	}
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *attr) NextSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return this.parent.(nodevalue).nextChild(this)
	}
}

func (this *attr) PreviousSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return this.parent.(nodevalue).previousChild(this)
	}
}

func (this *attr) Name() string {
	return this.name
}

func (this *attr) Value() string {
	return this.cdata
}

func (this *attr) SetValue(cdata string) {
	this.cdata = cdata
}

func (this *attr) OwnerElement() dom.Element {
	return this.ParentElement()
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *attr) AppendChild(child dom.Node) dom.Node {
	return nil
}

func (this *attr) CloneNode(bool) dom.Node {
	return NewNode(this.document, this.name, this.nodetype, this.cdata)
}

func (this *attr) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	// NO-OP
	return nil
}

func (this *attr) RemoveChild(child dom.Node) {
	// NO-OP
}

func (this *attr) ReplaceChild(dom.Node, dom.Node) {
	// NO-OP
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *attr) v() *node {
	return this.node
}

func (this *attr) write(w io.Writer) (int, error) {
	return w.Write([]byte(this.name + "=" + strconv.Quote(html.EscapeString(this.cdata))))
}
