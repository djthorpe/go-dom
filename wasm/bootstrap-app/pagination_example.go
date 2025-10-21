package main

import (
	"fmt"

	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddPaginationExamples adds pagination component examples to the app
func AddPaginationExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Example 1: Basic pagination with arrow icons
	card1 := app.Card()
	card1.AddClass("mb-4")
	card1.Header(app.H4(app.CreateTextNode("Basic Pagination with Icons")).Element)
	card1Body := card1.Body()

	pg1 := app.Pagination("Page navigation")

	// Previous with left arrow icon
	prevContent1 := app.CreateElement("span")
	prevIcon1 := app.Icon("chevron-left")
	prevIcon1.AddClass("me-1")
	prevContent1.AppendChild(prevIcon1.Element)
	prevContent1.AppendChild(app.CreateTextNode("Previous"))
	pg1.AddItemWithNode(prevContent1, "#", false, false).AddEventListener("click", func(target dom.Node) {
		fmt.Println("Previous clicked")
	})

	pg1.AddPage(1, "#page1", false)
	pg1.AddPage(2, "#page2", true) // Active page
	pg1.AddPage(3, "#page3", false)

	// Next with right arrow icon
	nextContent1 := app.CreateElement("span")
	nextContent1.AppendChild(app.CreateTextNode("Next"))
	nextIcon1 := app.Icon("chevron-right")
	nextIcon1.AddClass("ms-1")
	nextContent1.AppendChild(nextIcon1.Element)
	pg1.AddItemWithNode(nextContent1, "#", false, false).AddEventListener("click", func(target dom.Node) {
		fmt.Println("Next clicked")
	})

	card1Body.Element.AppendChild(pg1.Element)
	container.AppendChild(card1.Element)

	// Example 2: Pagination with disabled state and icons
	card2 := app.Card()
	card2.AddClass("mb-4")
	card2.Header(app.H4(app.CreateTextNode("Pagination with Disabled Previous")).Element)
	card2Body := card2.Body()

	pg2 := app.Pagination("Page navigation")

	// Disabled previous with icon
	prevContent2 := app.CreateElement("span")
	prevIcon2 := app.Icon("chevron-double-left")
	prevContent2.AppendChild(prevIcon2.Element)
	pg2.AddItemWithNode(prevContent2, "#", false, true) // Disabled

	pg2.AddPage(1, "#page1", true)
	pg2.AddPage(2, "#page2", false)
	pg2.AddPage(3, "#page3", false)

	// Next with icon
	nextContent2 := app.CreateElement("span")
	nextIcon2 := app.Icon("chevron-double-right")
	nextContent2.AppendChild(nextIcon2.Element)
	pg2.AddItemWithNode(nextContent2, "#", false, false)

	card2Body.Element.AppendChild(pg2.Element)
	container.AppendChild(card2.Element)

	// Example 3: Pagination with ellipsis and arrow icons
	card3 := app.Card()
	card3.AddClass("mb-4")
	card3.Header(app.H4(app.CreateTextNode("Pagination with Ellipsis and Icons")).Element)
	card3Body := card3.Body()

	pg3 := app.Pagination("Page navigation")

	// Previous with icon
	prevContent3 := app.CreateElement("span")
	prevIcon3 := app.Icon("arrow-left")
	prevContent3.AppendChild(prevIcon3.Element)
	pg3.AddItemWithNode(prevContent3, "#", false, false)

	pg3.AddPage(1, "#page1", false)
	pg3.AddEllipsis()
	pg3.AddPage(5, "#page5", false)
	pg3.AddPage(6, "#page6", true)
	pg3.AddPage(7, "#page7", false)
	pg3.AddEllipsis()
	pg3.AddPage(10, "#page10", false)

	// Next with icon
	nextContent3 := app.CreateElement("span")
	nextIcon3 := app.Icon("arrow-right")
	nextContent3.AppendChild(nextIcon3.Element)
	pg3.AddItemWithNode(nextContent3, "#", false, false)

	card3Body.Element.AppendChild(pg3.Element)
	container.AppendChild(card3.Element)

	// Example 4: Centered pagination with icons
	card4 := app.Card()
	card4.AddClass("mb-4")
	card4.Header(app.H4(app.CreateTextNode("Centered Pagination with Icons")).Element)
	card4Body := card4.Body()

	pg4 := app.Pagination("Page navigation")
	pg4.SetAlignment("center")

	// Previous with icon only
	prevContent4 := app.CreateElement("span")
	prevIcon4 := app.Icon("caret-left-fill")
	prevContent4.AppendChild(prevIcon4.Element)
	pg4.AddItemWithNode(prevContent4, "#", false, false)

	pg4.AddPage(1, "#page1", false)
	pg4.AddPage(2, "#page2", true)
	pg4.AddPage(3, "#page3", false)

	// Next with icon only
	nextContent4 := app.CreateElement("span")
	nextIcon4 := app.Icon("caret-right-fill")
	nextContent4.AppendChild(nextIcon4.Element)
	pg4.AddItemWithNode(nextContent4, "#", false, false)

	card4Body.Element.AppendChild(pg4.Element)
	container.AppendChild(card4.Element)

	// Example 5: Right-aligned pagination
	card5 := app.Card()
	card5.AddClass("mb-4")
	card5.Header(app.H4(app.CreateTextNode("Right-Aligned Pagination")).Element)
	card5Body := card5.Body()

	pg5 := app.Pagination("Page navigation")
	pg5.SetAlignment("end")
	pg5.AddPrevious("#", false)
	pg5.AddPage(1, "#page1", false)
	pg5.AddPage(2, "#page2", true)
	pg5.AddPage(3, "#page3", false)
	pg5.AddNext("#", false)

	card5Body.Element.AppendChild(pg5.Element)
	container.AppendChild(card5.Element)

	// Example 6: Pagination sizes with icons
	card6 := app.Card()
	card6.AddClass("mb-4")
	card6.Header(app.H4(app.CreateTextNode("Pagination Sizes with Icons")).Element)
	card6Body := card6.Body()

	sizesDiv := app.CreateElement("div")
	sizesDiv.AddClass("d-flex")
	sizesDiv.AddClass("flex-column")
	sizesDiv.AddClass("gap-3")

	// Large
	largeLabel := app.CreateElement("div")
	largeLabel.AppendChild(app.CreateTextNode("Large:"))
	sizesDiv.AppendChild(largeLabel)

	pgLarge := app.Pagination("Page navigation")
	pgLarge.SetSize(bs5.PaginationSizeLarge)

	prevLarge := app.CreateElement("span")
	prevLarge.AppendChild(app.Icon("chevron-bar-left").Element)
	pgLarge.AddItemWithNode(prevLarge, "#", false, false)

	pgLarge.AddPage(1, "#page1", false)
	pgLarge.AddPage(2, "#page2", true)
	pgLarge.AddPage(3, "#page3", false)

	nextLarge := app.CreateElement("span")
	nextLarge.AppendChild(app.Icon("chevron-bar-right").Element)
	pgLarge.AddItemWithNode(nextLarge, "#", false, false)

	sizesDiv.AppendChild(pgLarge.Element)

	// Default
	defaultLabel := app.CreateElement("div")
	defaultLabel.AppendChild(app.CreateTextNode("Default:"))
	sizesDiv.AppendChild(defaultLabel)

	pgDefault := app.Pagination("Page navigation")

	prevDefault := app.CreateElement("span")
	prevDefault.AppendChild(app.Icon("chevron-bar-left").Element)
	pgDefault.AddItemWithNode(prevDefault, "#", false, false)

	pgDefault.AddPage(1, "#page1", false)
	pgDefault.AddPage(2, "#page2", true)
	pgDefault.AddPage(3, "#page3", false)

	nextDefault := app.CreateElement("span")
	nextDefault.AppendChild(app.Icon("chevron-bar-right").Element)
	pgDefault.AddItemWithNode(nextDefault, "#", false, false)

	sizesDiv.AppendChild(pgDefault.Element)

	// Small
	smallLabel := app.CreateElement("div")
	smallLabel.AppendChild(app.CreateTextNode("Small:"))
	sizesDiv.AppendChild(smallLabel)

	pgSmall := app.Pagination("Page navigation")
	pgSmall.SetSize(bs5.PaginationSizeSmall)

	prevSmall := app.CreateElement("span")
	prevSmall.AppendChild(app.Icon("chevron-bar-left").Element)
	pgSmall.AddItemWithNode(prevSmall, "#", false, false)

	pgSmall.AddPage(1, "#page1", false)
	pgSmall.AddPage(2, "#page2", true)
	pgSmall.AddPage(3, "#page3", false)

	nextSmall := app.CreateElement("span")
	nextSmall.AppendChild(app.Icon("chevron-bar-right").Element)
	pgSmall.AddItemWithNode(nextSmall, "#", false, false)

	sizesDiv.AppendChild(pgSmall.Element)

	card6Body.Element.AppendChild(sizesDiv)
	container.AppendChild(card6.Element)

	return container.Element
}
