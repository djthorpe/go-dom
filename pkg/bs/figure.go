package bs

import (

	// Namespace imports
	"fmt"

	. "github.com/djthorpe/go-wasmbuild"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// figure are elements that represent figure views
type figure struct {
	View
}

// figurecaption are elements that represent figure caption, within the figure
type figurecaption struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewFigure        = "mvc-bs-figure"
	ViewFigureCaption = "mvc-bs-figurecaption"
)

func init() {
	RegisterView(ViewFigure, newFigureFromElement)
	RegisterView(ViewFigureCaption, newFigureCaptionFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Figure(opt ...Opt) *figure {
	return &figure{NewView(ViewFigure, "FIGURE", append(opt, WithClass("figure"))...)}
}

func newFigureFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "FIGURE" {
		panic(fmt.Sprintf("newFigureFromElement: invalid tag name %q", tagName))
	}
	return &figure{NewViewWithElement(element)}
}

func newFigureCaptionFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "FIGCAPTION" {
		panic(fmt.Sprintf("newFigureCaptionFromElement: invalid tag name %q", tagName))
	}
	return &figurecaption{NewViewWithElement(element)}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (figure *figure) Append(children ...any) View {
	return figure.View.Append(figure.prepare(children)...)
}

func (figure *figure) Insert(children ...any) View {
	// For Insert, we need special handling because View.Insert adds at the beginning
	// but we want the caption to remain at the end
	prepared := figure.prepare(children)

	// If prepare returned both image and caption, we need to insert only the image
	// and then append the caption
	if len(prepared) == 2 {
		// Insert the image (first element)
		figure.View.Insert(prepared[0])
		// Append the caption (second element) to keep it at the end
		figure.View.Append(prepared[1])
		return figure
	}

	// Otherwise just insert normally
	return figure.View.Insert(prepared...)
}

func (figure *figure) Caption(children ...any) *figure {
	// Remove any existing caption first
	caption := figure.caption()
	if caption != nil {
		caption.Root().Remove()
	}

	// We append the caption to the figure at the bottom of the view
	figure.View.Append(NewView(ViewFigureCaption, "FIGCAPTION", WithClass("figure-caption")).Append(children...))

	// Return the figure
	return figure
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Prepare the children for adding to the body
func (figure *figure) prepare(children []any) []any {
	if len(children) == 0 {
		return children
	}

	// Only accept a single image
	if len(children) != 1 {
		panic(fmt.Sprintf("Append/Insert: figure accepts only one child at a time, got %d", len(children)))
	}

	// Verify it's an image
	img, ok := children[0].(*image)
	if !ok {
		panic(fmt.Sprintf("Append/Insert: figure only accepts image components, got %T", children[0]))
	}

	// Add "figure-img" class to the image
	img.Opts(mvc.WithClass("figure-img"))

	// Detach the caption if it exists so we can re-append it after the image
	caption := figure.caption()
	var captionElement Element
	if caption != nil {
		captionElement = caption.Root()
		captionElement.Remove()
	}

	// Return the image, and the caption will be re-added after
	result := []any{img}
	if captionElement != nil {
		result = append(result, captionElement)
	}

	return result
}

// Return the figurecaption element
func (figure *figure) caption() *figurecaption {
	element := figure.Root().LastElementChild()
	if element == nil || element.TagName() != "FIGCAPTION" {
		return nil
	}
	return newFigureCaptionFromElement(element).(*figurecaption)
}
