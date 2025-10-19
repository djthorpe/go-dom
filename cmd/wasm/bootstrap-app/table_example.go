package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddTableExamples adds table component examples to the app
func AddTableExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Example 1: Basic table
	card1 := app.Card()
	card1.AddClass("mb-4")
	card1.Header(app.H4(app.CreateTextNode("Basic Table")).Element)
	card1Body := card1.Body()

	table1 := app.Table()

	// Head
	headRow1 := table1.Head().AddRow()
	headRow1.AddHeaderCell("#").SetScope("col")
	headRow1.AddHeaderCell("First").SetScope("col")
	headRow1.AddHeaderCell("Last").SetScope("col")
	headRow1.AddHeaderCell("Handle").SetScope("col")

	// Body
	row1_1 := table1.Body().AddRow()
	row1_1.AddHeaderCell("1").SetScope("row")
	row1_1.AddCell("Mark")
	row1_1.AddCell("Otto")
	row1_1.AddCell("@mdo")

	row1_2 := table1.Body().AddRow()
	row1_2.AddHeaderCell("2").SetScope("row")
	row1_2.AddCell("Jacob")
	row1_2.AddCell("Thornton")
	row1_2.AddCell("@fat")

	row1_3 := table1.Body().AddRow()
	row1_3.AddHeaderCell("3").SetScope("row")
	row1_3.AddCell("Larry")
	row1_3.AddCell("Bird")
	row1_3.AddCell("@twitter")

	card1Body.Element.AppendChild(table1.Element)
	container.AppendChild(card1.Element)

	// Example 2: Striped table
	card2 := app.Card()
	card2.AddClass("mb-4")
	card2.Header(app.H4(app.CreateTextNode("Striped Rows")).Element)
	card2Body := card2.Body()

	table2 := app.Table()
	table2.SetStriped(true)

	// Head
	headRow2 := table2.Head().AddRow()
	headRow2.AddHeaderCell("#").SetScope("col")
	headRow2.AddHeaderCell("First").SetScope("col")
	headRow2.AddHeaderCell("Last").SetScope("col")
	headRow2.AddHeaderCell("Handle").SetScope("col")

	// Body
	for i := 1; i <= 5; i++ {
		row := table2.Body().AddRow()
		row.AddHeaderCell(string(rune('0' + i))).SetScope("row")
		row.AddCell("Data " + string(rune('0'+i)))
		row.AddCell("Info " + string(rune('0'+i)))
		row.AddCell("@handle" + string(rune('0'+i)))
	}

	card2Body.Element.AppendChild(table2.Element)
	container.AppendChild(card2.Element)

	// Example 3: Hoverable table
	card3 := app.Card()
	card3.AddClass("mb-4")
	card3.Header(app.H4(app.CreateTextNode("Hoverable Rows")).Element)
	card3Body := card3.Body()

	table3 := app.Table()
	table3.SetHoverable(true)

	// Head
	headRow3 := table3.Head().AddRow()
	headRow3.AddHeaderCell("#").SetScope("col")
	headRow3.AddHeaderCell("Name").SetScope("col")
	headRow3.AddHeaderCell("Status").SetScope("col")

	// Body
	row3_1 := table3.Body().AddRow()
	row3_1.AddHeaderCell("1").SetScope("row")
	row3_1.AddCell("Active User")
	row3_1.AddCell("Online")

	row3_2 := table3.Body().AddRow()
	row3_2.AddHeaderCell("2").SetScope("row")
	row3_2.AddCell("Pending User")
	row3_2.AddCell("Away")

	card3Body.Element.AppendChild(table3.Element)
	container.AppendChild(card3.Element)

	// Example 4: Small table
	card4 := app.Card()
	card4.AddClass("mb-4")
	card4.Header(app.H4(app.CreateTextNode("Small Table")).Element)
	card4Body := card4.Body()

	table4 := app.Table()
	table4.SetSmall(true)

	// Head
	headRow4 := table4.Head().AddRow()
	headRow4.AddHeaderCell("#").SetScope("col")
	headRow4.AddHeaderCell("First").SetScope("col")
	headRow4.AddHeaderCell("Last").SetScope("col")

	// Body
	row4_1 := table4.Body().AddRow()
	row4_1.AddHeaderCell("1").SetScope("row")
	row4_1.AddCell("Mark")
	row4_1.AddCell("Otto")

	row4_2 := table4.Body().AddRow()
	row4_2.AddHeaderCell("2").SetScope("row")
	row4_2.AddCell("Jacob")
	row4_2.AddCell("Thornton")

	card4Body.Element.AppendChild(table4.Element)
	container.AppendChild(card4.Element)

	// Example 5: Bordered table
	card5 := app.Card()
	card5.AddClass("mb-4")
	card5.Header(app.H4(app.CreateTextNode("Bordered Table")).Element)
	card5Body := card5.Body()

	table5 := app.Table()
	table5.SetBordered(true)

	// Head
	headRow5 := table5.Head().AddRow()
	headRow5.AddHeaderCell("#").SetScope("col")
	headRow5.AddHeaderCell("First").SetScope("col")
	headRow5.AddHeaderCell("Last").SetScope("col")

	// Body
	row5_1 := table5.Body().AddRow()
	row5_1.AddHeaderCell("1").SetScope("row")
	row5_1.AddCell("Mark")
	row5_1.AddCell("Otto")

	row5_2 := table5.Body().AddRow()
	row5_2.AddHeaderCell("2").SetScope("row")
	row5_2.AddCell("Jacob")
	row5_2.AddCell("Thornton")

	card5Body.Element.AppendChild(table5.Element)
	container.AppendChild(card5.Element)

	// Example 6: Table with colored rows
	card6 := app.Card()
	card6.AddClass("mb-4")
	card6.Header(app.H4(app.CreateTextNode("Table with Colored Rows")).Element)
	card6Body := card6.Body()

	table6 := app.Table()

	// Head
	headRow6 := table6.Head().AddRow()
	headRow6.AddHeaderCell("Class").SetScope("col")
	headRow6.AddHeaderCell("Heading").SetScope("col")
	headRow6.AddHeaderCell("Status").SetScope("col")

	// Body with colored rows
	row6_1 := table6.Body().AddRow()
	row6_1.SetColor(bs5.ColorPrimary)
	row6_1.AddCell("Primary")
	row6_1.AddCell("Cell")
	row6_1.AddCell("Active")

	row6_2 := table6.Body().AddRow()
	row6_2.SetColor(bs5.ColorSuccess)
	row6_2.AddCell("Success")
	row6_2.AddCell("Cell")
	row6_2.AddCell("Completed")

	row6_3 := table6.Body().AddRow()
	row6_3.SetColor(bs5.ColorDanger)
	row6_3.AddCell("Danger")
	row6_3.AddCell("Cell")
	row6_3.AddCell("Error")

	row6_4 := table6.Body().AddRow()
	row6_4.SetColor(bs5.ColorWarning)
	row6_4.AddCell("Warning")
	row6_4.AddCell("Cell")
	row6_4.AddCell("Pending")

	card6Body.Element.AppendChild(table6.Element)
	container.AppendChild(card6.Element)

	// Example 7: Table with head and foot
	card7 := app.Card()
	card7.AddClass("mb-4")
	card7.Header(app.H4(app.CreateTextNode("Table with Header and Footer")).Element)
	card7Body := card7.Body()

	table7 := app.Table()
	table7.SetStriped(true)

	// Head with dark background
	table7.Head().SetColor(bs5.ColorDark)
	headRow7 := table7.Head().AddRow()
	headRow7.AddHeaderCell("Product").SetScope("col")
	headRow7.AddHeaderCell("Price").SetScope("col")
	headRow7.AddHeaderCell("Quantity").SetScope("col")

	// Body
	row7_1 := table7.Body().AddRow()
	row7_1.AddCell("Widget A")
	row7_1.AddCell("$10.00")
	row7_1.AddCell("5")

	row7_2 := table7.Body().AddRow()
	row7_2.AddCell("Widget B")
	row7_2.AddCell("$15.00")
	row7_2.AddCell("3")

	row7_3 := table7.Body().AddRow()
	row7_3.AddCell("Widget C")
	row7_3.AddCell("$20.00")
	row7_3.AddCell("2")

	// Foot
	footRow7 := table7.Foot().AddRow()
	footRow7.AddHeaderCell("Total").SetScope("row")
	footRow7.AddCell("$95.00")
	footRow7.AddCell("10")

	card7Body.Element.AppendChild(table7.Element)
	container.AppendChild(card7.Element)

	return container.Element
}
