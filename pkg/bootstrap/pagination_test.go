package bootstrap_test

import (
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	"github.com/stretchr/testify/assert"
)

func TestPagination_Basic(t *testing.T) {
	pagination := bs.Pagination("Page navigation")
	assert.NotNil(t, pagination)
	assert.NotNil(t, pagination.Element())
}

func TestPagination_TagName(t *testing.T) {
	pagination := bs.Pagination("Page navigation")
	element := pagination.Element()
	assert.Equal(t, "NAV", element.TagName())
}

func TestPagination_DefaultAriaLabel(t *testing.T) {
	pagination := bs.Pagination("")
	element := pagination.Element()
	assert.Equal(t, "Page navigation", element.GetAttribute("aria-label"))
}

func TestPagination_CustomAriaLabel(t *testing.T) {
	pagination := bs.Pagination("Search results pages")
	element := pagination.Element()
	assert.Equal(t, "Search results pages", element.GetAttribute("aria-label"))
}

func TestPagination_HasULChild(t *testing.T) {
	pagination := bs.Pagination("Page navigation")
	element := pagination.Element()
	children := element.Children()
	assert.Equal(t, 1, len(children))
	assert.Equal(t, "UL", children[0].TagName())
}

func TestPagination_ULHasPaginationClass(t *testing.T) {
	pagination := bs.Pagination("Page navigation")
	element := pagination.Element()
	children := element.Children()
	assert.True(t, children[0].ClassList().Contains("pagination"))
}

func TestPagination_AppendText(t *testing.T) {
	pagination := bs.Pagination("Page navigation").Append("1")
	element := pagination.Element()
	ul := element.Children()[0]
	liElements := ul.Children()

	assert.Equal(t, 1, len(liElements))
	assert.Equal(t, "LI", liElements[0].TagName())
	assert.True(t, liElements[0].ClassList().Contains("page-item"))
	assert.Equal(t, "1", liElements[0].TextContent())
}

func TestPagination_AppendMultiple(t *testing.T) {
	pagination := bs.Pagination("Page navigation").
		Append("Previous").
		Append("1").
		Append("2").
		Append("Next")

	element := pagination.Element()
	ul := element.Children()[0]
	liElements := ul.Children()

	assert.Equal(t, 4, len(liElements))
	assert.Equal(t, "Previous", liElements[0].TextContent())
	assert.Equal(t, "1", liElements[1].TextContent())
	assert.Equal(t, "2", liElements[2].TextContent())
	assert.Equal(t, "Next", liElements[3].TextContent())
}

func TestPagination_AppendComponent(t *testing.T) {
	link := bs.Link("#").Append("Page 1")
	pagination := bs.Pagination("Page navigation").Append(link)

	element := pagination.Element()
	ul := element.Children()[0]
	liElements := ul.Children()

	assert.Equal(t, 1, len(liElements))
	assert.Equal(t, "LI", liElements[0].TagName())
	assert.True(t, liElements[0].ClassList().Contains("page-item"))

	// Check that the link is inside the li
	liChildren := liElements[0].Children()
	assert.Equal(t, 1, len(liChildren))
	assert.Equal(t, "A", liChildren[0].TagName())
}

func TestPagination_Insert(t *testing.T) {
	pagination := bs.Pagination("Page navigation").
		Append("2")

	// Insert uses the inherited Component method which inserts into root (nav)
	// It doesn't wrap in <li class="page-item"> like Append does
	pagination.Insert("Text node")

	element := pagination.Element()
	children := element.ChildNodes()

	// Should have 2 children: the inserted text node and the ul
	assert.True(t, len(children) >= 2)
	// First child should be the text node
	assert.Equal(t, "Text node", children[0].TextContent())
}

func TestPagination_IsComponent(t *testing.T) {
	pagination := bs.Pagination("Page navigation")
	var _ dom.Component = pagination
	assert.NotNil(t, pagination)
}

func TestPagination_ComponentMethods(t *testing.T) {
	pagination := bs.Pagination("Page navigation")

	// Test Element() method
	element := pagination.Element()
	assert.NotNil(t, element)
	assert.Equal(t, "NAV", element.TagName())

	// Test Append() returns Component
	result := pagination.Append("Test")
	assert.NotNil(t, result)
	var _ dom.Component = result

	// Test Insert() returns Component
	result = pagination.Insert("First")
	assert.NotNil(t, result)
	var _ dom.Component = result
}

func TestPagination_AllPageItemsHaveClass(t *testing.T) {
	pagination := bs.Pagination("Page navigation").
		Append("1").
		Append("2").
		Append("3")

	element := pagination.Element()
	ul := element.Children()[0]
	liElements := ul.Children()

	for i, li := range liElements {
		assert.True(t, li.ClassList().Contains("page-item"),
			"Element %d should have page-item class", i)
	}
}

func TestPagination_OuterHTML(t *testing.T) {
	pagination := bs.Pagination("Test navigation").
		Append("1").
		Append("2")

	element := pagination.Element()
	html := element.OuterHTML()

	// HTML output uses lowercase tags (matching browser behavior)
	assert.Contains(t, html, "<nav")
	assert.Contains(t, html, "aria-label=\"Test navigation\"")
	assert.Contains(t, html, "<ul")
	assert.Contains(t, html, "class=\"pagination\"")
	assert.Contains(t, html, "<li")
	assert.Contains(t, html, "class=\"page-item\"")
}

func TestPagination_WithOptions(t *testing.T) {
	pagination := bs.Pagination("Page navigation",
		bs.WithClass("custom-class"),
		bs.WithAttribute("data-test", "value"))

	element := pagination.Element()
	assert.True(t, element.ClassList().Contains("custom-class"))
	assert.Equal(t, "value", element.GetAttribute("data-test"))
}

func TestPagination_EmptyAriaLabel(t *testing.T) {
	pagination := bs.Pagination("")
	element := pagination.Element()
	ariaLabel := element.GetAttribute("aria-label")

	// When empty string is passed, it should default to "Page navigation"
	assert.Equal(t, "Page navigation", ariaLabel)
}

func TestPagination_ChainedAppend(t *testing.T) {
	pagination := bs.Pagination("Pages")

	result := pagination.
		Append("1").
		Append("2").
		Append("3")

	// Verify chaining returns a Component
	var _ dom.Component = result

	// Verify all items were added
	element := pagination.Element()
	ul := element.Children()[0]
	assert.Equal(t, 3, len(ul.Children()))
}
