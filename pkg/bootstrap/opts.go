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
	classList TokenList
}

type Opt func(*opts) error

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewOpts(opt ...Opt) (*opts, error) {
	o := &opts{
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

func WithContainerFluid() Opt {
	return func(o *opts) error {
		o.classList.Remove("container")
		o.classList.Add("container-fluid")
		return nil
	}
}
