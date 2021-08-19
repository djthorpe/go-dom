//go:build !js

package dom

import (
	"fmt"
	"io"
	"strconv"

	dom "github.com/djthorpe/go-dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type doctype struct {
	*node
	publicid string
	systemid string
}

/////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	startdoctype = []byte("<!DOCTYPE ")
	enddoctype   = []byte(">\n")
)

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

func (this *doctype) NextSibling() dom.Node {
	return nextSibling(this.parent, this)
}

func (this *doctype) PreviousSibling() dom.Node {
	return previousSibling(this.parent, this)
}

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
	clone := NewNode(this.document, this.name, this.nodetype, this.cdata).(*doctype)
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

func (this *doctype) write(w io.Writer) (int, error) {
	s := 0
	if n, err := w.Write(startdoctype); err != nil {
		return 0, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte(this.name)); err != nil {
		return 0, err
	} else {
		s += n
	}
	if this.publicid != "" {
		if n, err := w.Write([]byte(" PUBLIC " + strconv.Quote(this.publicid))); err != nil {
			return 0, err
		} else {
			s += n
		}
		if this.systemid != "" {
			if n, err := w.Write([]byte(" " + strconv.Quote(this.systemid))); err != nil {
				return 0, err
			} else {
				s += n
			}
		}
	} else if this.systemid != "" {
		if n, err := w.Write([]byte(" SYSTEM " + strconv.Quote(this.systemid))); err != nil {
			return 0, err
		} else {
			s += n
		}
	}
	if n, err := w.Write(enddoctype); err != nil {
		return 0, err
	} else {
		s += n
	}
	return s, nil
}
