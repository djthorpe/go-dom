package bs

import (
	"fmt"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type icon struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewIcon = "mvc-bs-icon"
)

func init() {
	RegisterView(ViewIcon, newIconFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Icon(name string) *icon {
	return &icon{NewView(ViewIcon, "I", WithClass("bi-"+name))}
}

func newIconFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "I" {
		panic(fmt.Sprintf("newIconFromElement: invalid tag name %q", tagName))
	}
	return &icon{NewViewWithElement(element)}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (icon *icon) Append(children ...any) View {
	panic("Append: not supported for icon")
}

func (icon *icon) Insert(children ...any) View {
	panic("Insert: not supported for icon")
}
