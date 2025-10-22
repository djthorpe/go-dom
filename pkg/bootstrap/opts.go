package bootstrap

import (

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type opts struct {
	name       name
	classList  TokenList
	attributes map[string]string
}

type Opt func(*opts) error

// Breakpoint defines the responsive breakpoint for containers
type Breakpoint string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	// BreakpointDefault is the default responsive container (max-width at each breakpoint)
	BreakpointDefault Breakpoint = ""
	// BreakpointSmall creates a container that is 100% wide until ≥576px
	BreakpointSmall Breakpoint = "sm"
	// BreakpointMedium creates a container that is 100% wide until ≥768px
	BreakpointMedium Breakpoint = "md"
	// BreakpointLarge creates a container that is 100% wide until ≥992px
	BreakpointLarge Breakpoint = "lg"
	// BreakpointXLarge creates a container that is 100% wide until ≥1200px
	BreakpointXLarge Breakpoint = "xl"
	// BreakpointXXLarge creates a container that is 100% wide until ≥1400px
	BreakpointXXLarge Breakpoint = "xxl"
	// BreakpointFluid creates a full-width container (100% at all breakpoints)
	BreakpointFluid Breakpoint = "fluid"
)

// Size defines button sizes
type Size string

const (
	// SizeDefault is the default button size
	SizeDefault Size = ""
	// SizeSmall creates a small button (btn-sm)
	SizeSmall Size = "sm"
	// SizeLarge creates a large button (btn-lg)
	SizeLarge Size = "lg"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewOpts(name name, opt ...Opt) (*opts, error) {
	o := &opts{
		name:       name,
		classList:  dom.NewTokenList(),
		attributes: make(map[string]string),
	}
	if err := o.apply(opt...); err != nil {
		return nil, err
	}
	return o, nil
}

func (o *opts) apply(opts ...Opt) error {
	for _, opt := range opts {
		if opt != nil {
			if err := opt(o); err != nil {
				return err
			}
		}
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC FUNCTIONS

func WithTheme(theme Color) Opt {
	return func(o *opts) error {
		switch theme {
		case LIGHT:
			o.attributes["data-bs-theme"] = "light"
		case DARK:
			o.attributes["data-bs-theme"] = "dark"
		default:
			return ErrBadParameter.Withf("Invalid theme color: %q", theme)
		}
		return nil
	}
}

func WithClass(class ...string) Opt {
	return func(o *opts) error {
		o.classList.Add(class...)
		return nil
	}
}

func WithoutClass(class ...string) Opt {
	return func(o *opts) error {
		o.classList.Remove(class...)
		return nil
	}
}

// WithBreakpoint sets the container breakpoint type.
// Use BreakpointFluid for full-width, BreakpointSmall/Medium/Large/XLarge/XXLarge for responsive containers,
// or BreakpointDefault for the default container behavior.
func WithBreakpoint(breakpoint Breakpoint) Opt {
	return func(o *opts) error {
		// Error if this isn't a container
		if o.name != ContainerComponent {
			return ErrBadParameter.Withf("Cannot use WithBreakpoint with component of type %q", o.name)
		}

		// Remove any existing container classes
		o.classList.Remove("container", "container-sm", "container-md", "container-lg", "container-xl", "container-xxl", "container-fluid")

		// Add the appropriate container class
		if breakpoint == BreakpointDefault || breakpoint == "" {
			o.classList.Add("container")
		} else {
			o.classList.Add("container-" + string(breakpoint))
		}
		return nil
	}
}

func WithBorder(position Position, color ...Color) Opt {
	return func(o *opts) error {
		// WithBorder works with heading, container, badge, and alert components
		if o.name != HeadingComponent && o.name != ContainerComponent && o.name != BadgeComponent && o.name != AlertComponent {
			return ErrBadParameter.Withf("Cannot use WithBorder with component of type %q", o.name)
		}

		// Remove all existing border classes
		o.classList.Remove(BorderAll.borderClassNames()...)

		// Add new border classes
		if pposition := position & BorderAll; pposition != 0 {
			o.classList.Add(pposition.borderClassNames()...)
		}

		// TODO: Remove any existing border color classes

		// Add border color
		if len(color) == 1 {
			o.classList.Add(color[0].className("border"))
		} else if len(color) > 1 {
			return ErrBadParameter.Withf("Only one color can be specified for borders")
		}

		// Return success
		return nil
	}
}

func WithBackground(color Color) Opt {
	return func(o *opts) error {
		// TODO: Remove any existing background color classes

		// Add background color
		o.classList.Add(color.className("bg"))

		// Return success
		return nil
	}
}

// WithColor sets the color for badge, alert, button, icon, link, and card components.
func WithColor(color Color) Opt {
	return func(o *opts) error {
		// TODO: Remove any existing color classes

		// Add appropriate color class based on component type
		switch o.name {
		case NavBarComponent:
			o.classList.Add(color.className("bg"))
		case BadgeComponent:
			o.classList.Add(color.className("text-bg"))
		case AlertComponent:
			o.classList.Add(color.className("alert"))
		case ButtonComponent:
			o.classList.Add(color.className("btn"))
		case IconComponent:
			o.classList.Add(color.className("text"))
		case LinkComponent:
			o.classList.Add(color.className("link"))
		case CardComponent:
			o.classList.Add(color.className("text-bg"))
		default:
			return ErrBadParameter.Withf("Cannot use WithColor with component of type %q", o.name)
		}

		// Return success
		return nil
	}
}

// WithMargin sets the margin for the specified position, where position can be one or more of
// TOP, BOTTOM, START, END. Margin is an integer from -5 to 5, where negative values indicate
// negative margins, 0 indicates no margin, and positive values indicate increasing margin sizes.
func WithMargin(position Position, margin int) Opt {
	return func(o *opts) error {
		// TODO: Remove any existing margin classes

		// Add margin classes
		o.classList.Add(position.marginClassNames(margin)...)

		// Return success
		return nil
	}
}

// WithPadding sets the padding for the specified position, where position can be one or more of
// TOP, BOTTOM, START, END. Padding is an integer from 1 to 5, where negative or zero values are not allowed.
func WithPadding(position Position, padding int) Opt {
	return func(o *opts) error {
		// Validate padding size (must be positive, no zero or negative)
		if padding <= 0 {
			return ErrBadParameter.Withf("Padding must be a positive value, got %d", padding)
		}

		// TODO: Remove any existing padding classes

		// Add padding classes
		o.classList.Add(position.paddingClassNames(padding)...)

		// Return success
		return nil
	}
}

// WithSize sets the size for button components (btn-sm, btn-lg)
// or button group components (btn-group-sm, btn-group-lg)
func WithSize(size Size) Opt {
	return func(o *opts) error {
		// WithSize only works with button or button-group components
		if o.name != ButtonComponent && o.name != ButtonGroupComponent {
			return ErrBadParameter.Withf("Cannot use WithSize with component of type %q", o.name)
		}

		// Skip if default size
		if size == SizeDefault {
			return nil
		}

		// TODO: Remove any existing size classes

		// Add size class based on component type
		if o.name == ButtonGroupComponent {
			o.classList.Add("btn-group-" + string(size))
		} else {
			o.classList.Add("btn-" + string(size))
		}

		// Return success
		return nil
	}
}

// WithAriaLabel sets the aria-label attribute for accessibility
func WithAriaLabel(label string) Opt {
	return func(o *opts) error {
		if label != "" {
			o.attributes["aria-label"] = label
		}
		return nil
	}
}

// WithAttribute sets a custom HTML attribute
func WithAttribute(key, value string) Opt {
	return func(o *opts) error {
		if key != "" {
			o.attributes[key] = value
		}
		return nil
	}
}

// WithFlex sets flexbox alignment classes based on position
// - CENTER or MIDDLE: adds "d-flex align-items-center"
// - START or TOP: adds "d-flex align-items-start"
// - END or BOTTOM: adds "d-flex align-items-end"
// - START|END: adds "d-flex flex-row" (horizontal flow)
// - TOP|BOTTOM: adds "d-flex flex-column" (vertical flow)
func WithFlex(position Position) Opt {
	return func(o *opts) error {
		// Always add d-flex
		o.classList.Add("d-flex")

		// Handle alignment
		switch {
		case position&CENTER != 0 || position&MIDDLE != 0:
			o.classList.Add("align-items-center")
		case position&START != 0 && position&END == 0:
			o.classList.Add("align-items-start")
		case position&TOP != 0 && position&BOTTOM == 0:
			o.classList.Add("align-items-start")
		case position&END != 0 && position&START == 0:
			o.classList.Add("align-items-end")
		case position&BOTTOM != 0 && position&TOP == 0:
			o.classList.Add("align-items-end")
		}

		// Handle flow direction
		if (position&START != 0) && (position&END != 0) {
			// Both START and END means horizontal flow
			o.classList.Add("flex-row")
		} else if (position&TOP != 0) && (position&BOTTOM != 0) {
			// Both TOP and BOTTOM means vertical flow
			o.classList.Add("flex-column")
		}

		return nil
	}
}
