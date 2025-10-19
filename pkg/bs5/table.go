package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Table struct {
	dom.Element
	thead *TableHead
	tbody *TableBody
	tfoot *TableFoot
}

type TableHead struct {
	dom.Element
	table *Table
}

type TableBody struct {
	dom.Element
	table *Table
}

type TableFoot struct {
	dom.Element
	table *Table
}

type TableRow struct {
	dom.Element
}

type TableHeaderCell struct {
	dom.Element
}

type TableCell struct {
	dom.Element
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Table creates a Bootstrap 5 table component
func (app *App) Table() *Table {
	// <table class="table">
	table := app.CreateElement("table")
	table.AddClass("table")

	return &Table{
		Element: table,
	}
}

////////////////////////////////////////////////////////////////////////
// TABLE METHODS

// Head creates or returns the table head
func (t *Table) Head() *TableHead {
	if t.thead == nil {
		// <thead>
		thead := t.Element.OwnerDocument().CreateElement("thead")

		// Insert at the beginning
		if t.Element.FirstChild() != nil {
			t.Element.InsertBefore(thead, t.Element.FirstChild())
		} else {
			t.Element.AppendChild(thead)
		}

		t.thead = &TableHead{
			Element: thead,
			table:   t,
		}
	}
	return t.thead
}

// Body creates or returns the table body
func (t *Table) Body() *TableBody {
	if t.tbody == nil {
		// <tbody>
		tbody := t.Element.OwnerDocument().CreateElement("tbody")
		t.Element.AppendChild(tbody)

		t.tbody = &TableBody{
			Element: tbody,
			table:   t,
		}
	}
	return t.tbody
}

// Foot creates or returns the table foot
func (t *Table) Foot() *TableFoot {
	if t.tfoot == nil {
		// <tfoot>
		tfoot := t.Element.OwnerDocument().CreateElement("tfoot")
		t.Element.AppendChild(tfoot)

		t.tfoot = &TableFoot{
			Element: tfoot,
			table:   t,
		}
	}
	return t.tfoot
}

// SetStriped adds zebra-striping to table rows
func (t *Table) SetStriped(striped bool) *Table {
	if striped {
		t.Element.AddClass("table-striped")
	} else {
		t.Element.RemoveClass("table-striped")
	}
	return t
}

// SetHoverable enables hover state on table rows
func (t *Table) SetHoverable(hoverable bool) *Table {
	if hoverable {
		t.Element.AddClass("table-hover")
	} else {
		t.Element.RemoveClass("table-hover")
	}
	return t
}

// SetBordered adds borders on all sides of the table and cells
func (t *Table) SetBordered(bordered bool) *Table {
	if bordered {
		t.Element.AddClass("table-bordered")
	} else {
		t.Element.RemoveClass("table-bordered")
	}
	return t
}

// SetBorderless removes all borders
func (t *Table) SetBorderless(borderless bool) *Table {
	if borderless {
		t.Element.AddClass("table-borderless")
	} else {
		t.Element.RemoveClass("table-borderless")
	}
	return t
}

// SetSmall makes the table more compact
func (t *Table) SetSmall(small bool) *Table {
	if small {
		t.Element.AddClass("table-sm")
	} else {
		t.Element.RemoveClass("table-sm")
	}
	return t
}

// SetResponsive makes the table responsive (horizontal scrolling)
func (t *Table) SetResponsive(responsive bool) *Table {
	// Note: For responsive, we typically wrap the table in a div
	// This method adds the class to the table itself for simplicity
	if responsive {
		// In a full implementation, we'd wrap in <div class="table-responsive">
		// For now, we'll just add a note that this should be handled by the container
		t.Element.AddClass("table-responsive")
	} else {
		t.Element.RemoveClass("table-responsive")
	}
	return t
}

// SetColor sets the table color variant
func (t *Table) SetColor(color ColorVariant) *Table {
	// Remove existing color classes
	colors := []ColorVariant{ColorPrimary, ColorSecondary, ColorSuccess, ColorDanger, ColorWarning, ColorInfo, ColorLight, ColorDark}
	for _, c := range colors {
		t.Element.RemoveClass("table-" + string(c))
	}
	// Add new color
	if color != "" {
		t.Element.AddClass("table-" + string(color))
	}
	return t
}

// AddClass adds a CSS class to the table
func (t *Table) AddClass(className string) *Table {
	t.Element.AddClass(className)
	return t
}

// RemoveClass removes a CSS class from the table
func (t *Table) RemoveClass(className string) *Table {
	t.Element.RemoveClass(className)
	return t
}

////////////////////////////////////////////////////////////////////////
// TABLE HEAD METHODS

// AddRow creates and adds a row to the table head
func (th *TableHead) AddRow() *TableRow {
	// <tr>
	tr := th.Element.OwnerDocument().CreateElement("tr")
	th.Element.AppendChild(tr)

	return &TableRow{
		Element: tr,
	}
}

// AddClass adds a CSS class to the table head
func (th *TableHead) AddClass(className string) *TableHead {
	th.Element.AddClass(className)
	return th
}

// RemoveClass removes a CSS class from the table head
func (th *TableHead) RemoveClass(className string) *TableHead {
	th.Element.RemoveClass(className)
	return th
}

// SetColor sets the table head color variant
func (th *TableHead) SetColor(color ColorVariant) *TableHead {
	// Remove existing color classes
	colors := []ColorVariant{ColorPrimary, ColorSecondary, ColorSuccess, ColorDanger, ColorWarning, ColorInfo, ColorLight, ColorDark}
	for _, c := range colors {
		th.Element.RemoveClass("table-" + string(c))
	}
	// Add new color
	if color != "" {
		th.Element.AddClass("table-" + string(color))
	}
	return th
}

////////////////////////////////////////////////////////////////////////
// TABLE BODY METHODS

// AddRow creates and adds a row to the table body
func (tb *TableBody) AddRow() *TableRow {
	// <tr>
	tr := tb.Element.OwnerDocument().CreateElement("tr")
	tb.Element.AppendChild(tr)

	return &TableRow{
		Element: tr,
	}
}

// AddClass adds a CSS class to the table body
func (tb *TableBody) AddClass(className string) *TableBody {
	tb.Element.AddClass(className)
	return tb
}

// RemoveClass removes a CSS class from the table body
func (tb *TableBody) RemoveClass(className string) *TableBody {
	tb.Element.RemoveClass(className)
	return tb
}

////////////////////////////////////////////////////////////////////////
// TABLE FOOT METHODS

// AddRow creates and adds a row to the table foot
func (tf *TableFoot) AddRow() *TableRow {
	// <tr>
	tr := tf.Element.OwnerDocument().CreateElement("tr")
	tf.Element.AppendChild(tr)

	return &TableRow{
		Element: tr,
	}
}

// AddClass adds a CSS class to the table foot
func (tf *TableFoot) AddClass(className string) *TableFoot {
	tf.Element.AddClass(className)
	return tf
}

// RemoveClass removes a CSS class from the table foot
func (tf *TableFoot) RemoveClass(className string) *TableFoot {
	tf.Element.RemoveClass(className)
	return tf
}

////////////////////////////////////////////////////////////////////////
// TABLE ROW METHODS

// AddHeaderCell adds a <th> header cell to the row
func (tr *TableRow) AddHeaderCell(text string) *TableHeaderCell {
	// <th>
	th := tr.Element.OwnerDocument().CreateElement("th")
	th.AppendChild(tr.Element.OwnerDocument().CreateTextNode(text))
	tr.Element.AppendChild(th)

	return &TableHeaderCell{
		Element: th,
	}
}

// AddCell adds a <td> data cell to the row
func (tr *TableRow) AddCell(text string) *TableCell {
	// <td>
	td := tr.Element.OwnerDocument().CreateElement("td")
	td.AppendChild(tr.Element.OwnerDocument().CreateTextNode(text))
	tr.Element.AppendChild(td)

	return &TableCell{
		Element: td,
	}
}

// AddCellWithNode adds a <td> data cell with a node as content
func (tr *TableRow) AddCellWithNode(node dom.Node) *TableCell {
	// <td>
	td := tr.Element.OwnerDocument().CreateElement("td")
	td.AppendChild(node)
	tr.Element.AppendChild(td)

	return &TableCell{
		Element: td,
	}
}

// SetColor sets the row color variant
func (tr *TableRow) SetColor(color ColorVariant) *TableRow {
	// Remove existing color classes
	colors := []ColorVariant{ColorPrimary, ColorSecondary, ColorSuccess, ColorDanger, ColorWarning, ColorInfo, ColorLight, ColorDark}
	for _, c := range colors {
		tr.Element.RemoveClass("table-" + string(c))
	}
	// Add new color
	if color != "" {
		tr.Element.AddClass("table-" + string(color))
	}
	return tr
}

// AddClass adds a CSS class to the table row
func (tr *TableRow) AddClass(className string) *TableRow {
	tr.Element.AddClass(className)
	return tr
}

// RemoveClass removes a CSS class from the table row
func (tr *TableRow) RemoveClass(className string) *TableRow {
	tr.Element.RemoveClass(className)
	return tr
}

////////////////////////////////////////////////////////////////////////
// TABLE HEADER CELL METHODS

// SetScope sets the scope attribute (col, row, colgroup, rowgroup)
func (th *TableHeaderCell) SetScope(scope string) *TableHeaderCell {
	th.Element.SetAttribute("scope", scope)
	return th
}

// SetColspan sets the colspan attribute
func (th *TableHeaderCell) SetColspan(colspan int) *TableHeaderCell {
	th.Element.SetAttribute("colspan", string(rune(colspan)))
	return th
}

// AddClass adds a CSS class to the header cell
func (th *TableHeaderCell) AddClass(className string) *TableHeaderCell {
	th.Element.AddClass(className)
	return th
}

// RemoveClass removes a CSS class from the header cell
func (th *TableHeaderCell) RemoveClass(className string) *TableHeaderCell {
	th.Element.RemoveClass(className)
	return th
}

////////////////////////////////////////////////////////////////////////
// TABLE CELL METHODS

// SetColspan sets the colspan attribute
func (tc *TableCell) SetColspan(colspan int) *TableCell {
	tc.Element.SetAttribute("colspan", string(rune(colspan)))
	return tc
}

// SetRowspan sets the rowspan attribute
func (tc *TableCell) SetRowspan(rowspan int) *TableCell {
	tc.Element.SetAttribute("rowspan", string(rune(rowspan)))
	return tc
}

// AddClass adds a CSS class to the table cell
func (tc *TableCell) AddClass(className string) *TableCell {
	tc.Element.AddClass(className)
	return tc
}

// RemoveClass removes a CSS class from the table cell
func (tc *TableCell) RemoveClass(className string) *TableCell {
	tc.Element.RemoveClass(className)
	return tc
}

// SetColor sets the cell color variant
func (tc *TableCell) SetColor(color ColorVariant) *TableCell {
	// Remove existing color classes
	colors := []ColorVariant{ColorPrimary, ColorSecondary, ColorSuccess, ColorDanger, ColorWarning, ColorInfo, ColorLight, ColorDark}
	for _, c := range colors {
		tc.Element.RemoveClass("table-" + string(c))
	}
	// Add new color
	if color != "" {
		tc.Element.AddClass("table-" + string(color))
	}
	return tc
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (t *Table) String() string {
	return "<bs5-table>"
}

func (th *TableHead) String() string {
	return "<bs5-table-head>"
}

func (tb *TableBody) String() string {
	return "<bs5-table-body>"
}

func (tf *TableFoot) String() string {
	return "<bs5-table-foot>"
}

func (tr *TableRow) String() string {
	return "<bs5-table-row>"
}

func (th *TableHeaderCell) String() string {
	return "<bs5-table-header-cell>"
}

func (tc *TableCell) String() string {
	return "<bs5-table-cell>"
}
