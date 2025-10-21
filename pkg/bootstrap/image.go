package bootstrap

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type image struct {
	component
}

// Ensure that image implements Component interface
var _ Component = (*image)(nil)

////////////////////////////////////////////////////////////////////////////////
// CONSTRUCTOR

// Image creates a Bootstrap image element with src and optional alt text
// The src parameter sets the image source URL, and options can be used to style the image
//
// Example:
//
//	Image("photo.jpg", WithAriaLabel("Profile photo"))
//	Image("logo.png", WithClass("img-fluid"))
//	Image("thumbnail.jpg", WithClass("img-thumbnail"))
func Image(src string, opt ...Opt) *image {
	// Create an img element
	root := dom.GetWindow().Document().CreateElement("IMG")

	// Set src attribute
	if src != "" {
		root.SetAttribute("src", src)
	}

	// Apply options
	if opts, err := NewOpts(ImageComponent); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			root.SetAttribute(key, value)
		}

		// Return the image component
		return &image{
			component: component{
				name: ImageComponent,
				root: root,
			},
		}
	}
}
