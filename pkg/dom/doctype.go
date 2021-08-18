//go:build !js

package dom

import (
	"fmt"

	dom "github.com/djthorpe/go-dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type doctype struct {
	*node
	publicid string
	systemid string
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *doctype) String() string {
	str := "<DOMDocumentType"
	if name := this.Name(); name != "" {
		str += fmt.Sprintf(" name=%q", name)
	}
	if publicid := this.PublicId(); publicid != "" {
		str += fmt.Sprintf(" publicId=%q", publicid)
	}
	if systemid := this.SystemId(); systemid != "" {
		str += fmt.Sprintf(" systemId=%q", systemid)
	}
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *doctype) Name() string {
	return this.name
}
func (this *doctype) PublicId() string {
	return this.publicid
}

func (this *doctype) SystemId() string {
	return this.systemid
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *doctype) AppendChild(child dom.Node) dom.Node {
	return nil
}

func (this *doctype) CloneNode(bool) dom.Node {
	clone := NewNode(this.document, this.name, this.nodetype).(*doctype)
	clone.publicid = this.publicid
	clone.systemid = this.systemid
	return clone
}

func (this *doctype) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	return nil
}

func (this *doctype) RemoveChild(child dom.Node) {
	// NO-OP
}

func (this *doctype) ReplaceChild(dom.Node, dom.Node) {
	// NO-OP
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *doctype) v() *node {
	return this.node
}
