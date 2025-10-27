package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// dropdown are elements to create a dropdown menu for a nav, narbar or button group
type dropdown struct {
	View
}

// dropdownitem are elements to insert into a dropdown menu
type dropdownitem struct {
	View
}

// dropdownheader are header elements to insert into a dropdown menu
type dropdownheader struct {
	View
}

// dropdowndivider are elements to insert into a dropdown menu
type dropdowndivider struct {
	View
}

var _ View = (*dropdown)(nil)
var _ View = (*dropdownitem)(nil)
var _ View = (*dropdownheader)(nil)
var _ View = (*dropdowndivider)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewDropdown        = "mvc-bs-dropdown"
	ViewDropdownItem    = "mvc-bs-dropdownitem"
	ViewDropdownHeader  = "mvc-bs-dropdownheader"
	ViewDropdownDivider = "mvc-bs-dropdowndivider"
)

func init() {
	RegisterView(ViewDropdown, nil)
	RegisterView(ViewDropdownItem, nil)
	RegisterView(ViewDropdownHeader, nil)
	RegisterView(ViewDropdownDivider, nil)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS
