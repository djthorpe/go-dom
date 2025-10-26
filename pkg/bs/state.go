package bs

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithDisabled adds a disabled attribute to a view
func WithDisabled(disabled bool) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewButton {
			return fmt.Errorf("WithDisabled: invalid view type %q", o.Name())
		}
		if disabled {
			return mvc.WithAttr("disabled", "disabled")(o)
		} else {
			return mvc.WithoutAttr("disabled")(o)
		}
	}
}

// WithActive adds an active attribute to a view
func WithActive(active bool) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewButton {
			return fmt.Errorf("WithActive: invalid view type %q", o.Name())
		}
		if active {
			return mvc.WithClass("active")(o)
		} else {
			return mvc.WithoutClass("active")(o)
		}
	}
}
