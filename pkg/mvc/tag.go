package mvc

import (
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// tag wraps a generic HTML tag element
type tag struct {
	Element
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Tag(tagName string, opts ...Opt) *tag {
	e := elementFactory(tagName)
	if len(opts) > 0 {
		if err := applyOpts(e, opts...); err != nil {
			panic(err)
		}
	}
	return &tag{Element: e}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (t *tag) Append(children ...any) *tag {
	for _, child := range children {
		t.AppendChild(NodeFromAny(child))
	}
	return t
}
