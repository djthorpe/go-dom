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
	name      name
	classList TokenList
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

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewOpts(name name, opt ...Opt) (*opts, error) {
	o := &opts{
		name:      name,
		classList: dom.NewTokenList(),
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
