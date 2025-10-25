package bootstrap

///////////////////////////////////////////////////////////////////////////////
// TABLE-SPECIFIC OPTIONS

// WithStriped adds zebra-striping to table rows within the tbody
func WithStriped() Opt {
	return WithClass("table-striped")
}

// WithStripedColumns adds zebra-striping to table columns
func WithStripedColumns() Opt {
	return WithClass("table-striped-columns")
}

// WithHover enables a hover state on table rows within tbody
func WithHover() Opt {
	return WithClass("table-hover")
}

// WithActive highlights a table row or cell
func WithActive() Opt {
	return WithClass("table-active")
}

// WithBordered adds borders on all sides of the table and cells
func WithBordered() Opt {
	return WithClass("table-bordered")
}

// WithBorderless removes all borders from the table
func WithBorderless() Opt {
	return WithClass("table-borderless")
}

// WithSmall makes the table more compact by cutting cell padding in half
func WithSmall() Opt {
	return WithClass("table-sm")
}

// WithGroupDivider adds a thicker border between table groups (thead, tbody, tfoot)
// Adds class: table-group-divider
func WithGroupDivider() Opt {
	return WithClass("table-group-divider")
}

// WithAnimation enables animations for row operations (InsertBefore, Delete, Replace)
// Uses scaleY transforms for smooth height transitions
// This adds a special marker class that the Table constructor will detect
func WithAnimation() Opt {
	return WithClass("table-animated")
}
