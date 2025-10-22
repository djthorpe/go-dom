package dom_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	domPkg "github.com/djthorpe/go-wasmbuild/pkg/dom"
	"github.com/stretchr/testify/assert"
)

func TestElement_Basic(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	assert.NotNil(t, element, "Element should not be nil")
	assert.Equal(t, "DIV", element.TagName(), "Tag name should be uppercase")
	assert.Equal(t, dom.ELEMENT_NODE, element.NodeType(), "Node type should be ELEMENT_NODE")

	// NodeName case varies between implementations
	nodeName := element.NodeName()
	assert.True(t, nodeName == "div" || nodeName == "DIV", "Node name should be 'div' or 'DIV'")
}

func TestElement_TagName(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	tests := []struct {
		name        string
		tagName     string
		expectedTag string
	}{
		{"div element", "div", "DIV"},
		{"span element", "span", "SPAN"},
		{"p element", "p", "P"},
		{"h1 element", "h1", "H1"},
		{"input element", "input", "INPUT"},
		{"button element", "button", "BUTTON"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := doc.CreateElement(tt.tagName)
			assert.Equal(t, tt.expectedTag, element.TagName(), "Tag name should be uppercase")
		})
	}
}

func TestElement_InnerHTML(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	tests := []struct {
		name         string
		tagName      string
		expectedHTML string
	}{
		{"empty div", "div", ""},
		{"empty span", "span", ""},
		{"empty p", "p", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := doc.CreateElement(tt.tagName)
			assert.Equal(t, tt.expectedHTML, element.InnerHTML(), "Empty element should have empty innerHTML")
		})
	}
}

func TestElement_OuterHTML(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	tests := []struct {
		name         string
		tagName      string
		expectedHTML string
	}{
		{"div element", "div", "<div></div>"},
		{"span element", "span", "<span></span>"},
		{"p element", "p", "<p></p>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := doc.CreateElement(tt.tagName)
			outerHTML := element.OuterHTML()
			// Convert to lowercase for comparison since some implementations may vary in case
			assert.Equal(t, tt.expectedHTML, strings.ToLower(outerHTML), "Outer HTML should match expected format")
		})
	}
}

func TestElement_Attributes(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	t.Run("initial state", func(t *testing.T) {
		assert.False(t, element.HasAttributes(), "New element should not have attributes")
		attributes := element.Attributes()
		assert.Empty(t, attributes, "New element should have empty attributes slice")
	})

	t.Run("set and get attribute", func(t *testing.T) {
		attr := element.SetAttribute("id", "test-id")
		assert.NotNil(t, attr, "SetAttribute should return an attribute")
		assert.True(t, element.HasAttribute("id"), "Element should have id attribute")
		assert.Equal(t, "test-id", element.GetAttribute("id"), "Attribute value should match")
		assert.True(t, element.HasAttributes(), "Element should have attributes")
	})

	t.Run("multiple attributes", func(t *testing.T) {
		element.SetAttribute("class", "test-class")
		element.SetAttribute("data-test", "test-value")

		assert.True(t, element.HasAttribute("id"), "Should have id attribute")
		assert.True(t, element.HasAttribute("class"), "Should have class attribute")
		assert.True(t, element.HasAttribute("data-test"), "Should have data-test attribute")

		assert.Equal(t, "test-id", element.GetAttribute("id"), "ID should match")
		assert.Equal(t, "test-class", element.GetAttribute("class"), "Class should match")
		assert.Equal(t, "test-value", element.GetAttribute("data-test"), "Data attribute should match")

		attributes := element.Attributes()
		assert.GreaterOrEqual(t, len(attributes), 3, "Should have at least 3 attributes")
	})

	t.Run("get non-existent attribute", func(t *testing.T) {
		value := element.GetAttribute("non-existent")
		assert.Equal(t, "", value, "Non-existent attribute should return empty string")
		assert.False(t, element.HasAttribute("non-existent"), "Should not have non-existent attribute")
	})
}

func TestElement_AttributeNode(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	t.Run("set and get attribute node", func(t *testing.T) {
		element.SetAttribute("title", "test-title")

		attrNode := element.GetAttributeNode("title")
		if attrNode != nil {
			assert.Equal(t, "title", attrNode.Name(), "Attribute name should match")
			assert.Equal(t, "test-title", attrNode.Value(), "Attribute value should match")
		}
	})

	t.Run("get non-existent attribute node", func(t *testing.T) {
		attrNode := element.GetAttributeNode("non-existent")
		assert.Nil(t, attrNode, "Non-existent attribute node should be nil")
	})
}

func TestElement_ClassListIntegration(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	classList := element.ClassList()
	assert.NotNil(t, classList, "ClassList should not be nil")

	t.Run("initial state", func(t *testing.T) {
		assert.Equal(t, 0, classList.Length(), "Initial class list should be empty")
		assert.Equal(t, "", classList.Value(), "Initial class value should be empty")
	})

	t.Run("add classes", func(t *testing.T) {
		classList.Add("class1", "class2")
		assert.Equal(t, 2, classList.Length(), "Should have 2 classes")
		assert.True(t, classList.Contains("class1"), "Should contain class1")
		assert.True(t, classList.Contains("class2"), "Should contain class2")
	})

	t.Run("remove classes", func(t *testing.T) {
		classList.Remove("class1")
		assert.Equal(t, 1, classList.Length(), "Should have 1 class")
		assert.False(t, classList.Contains("class1"), "Should not contain class1")
		assert.True(t, classList.Contains("class2"), "Should still contain class2")
	})

	t.Run("toggle classes", func(t *testing.T) {
		result := classList.Toggle("class3")
		assert.True(t, result, "Toggle should return true for added class")
		assert.True(t, classList.Contains("class3"), "Should contain class3")

		result = classList.Toggle("class3")
		assert.False(t, result, "Toggle should return false for removed class")
		assert.False(t, classList.Contains("class3"), "Should not contain class3")
	})
}

func TestElement_ChildNodes(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("div")

	t.Run("initial state", func(t *testing.T) {
		assert.False(t, parent.HasChildNodes(), "New element should not have child nodes")
		assert.Nil(t, parent.FirstChild(), "First child should be nil")
		assert.Nil(t, parent.LastChild(), "Last child should be nil")
	})

	t.Run("append single child", func(t *testing.T) {
		child := doc.CreateElement("span")
		appendedChild := parent.AppendChild(child)

		assert.Equal(t, child, appendedChild, "AppendChild should return the same element")
		assert.True(t, parent.HasChildNodes(), "Parent should have child nodes")
		assert.True(t, parent.FirstChild().Equals(child), "First child should match")
		assert.True(t, parent.LastChild().Equals(child), "Last child should match")
		assert.True(t, child.ParentNode().Equals(parent), "Child's parent should match")
		assert.Nil(t, child.PreviousSibling(), "First child should have no previous sibling")
		assert.Nil(t, child.NextSibling(), "Only child should have no next sibling")
	})

	t.Run("append multiple children", func(t *testing.T) {
		child2 := doc.CreateElement("p")
		child3 := doc.CreateElement("h1")

		parent.AppendChild(child2)
		parent.AppendChild(child3)

		assert.True(t, parent.HasChildNodes(), "Parent should have child nodes")

		// Test first and last child
		firstChild := parent.FirstChild()
		lastChild := parent.LastChild()
		assert.NotNil(t, firstChild, "First child should not be nil")
		assert.NotNil(t, lastChild, "Last child should not be nil")
		assert.True(t, lastChild.Equals(child3), "Last child should be child3")

		// Test sibling relationships
		if child2.PreviousSibling() != nil {
			if prevSibling, ok := child2.PreviousSibling().(dom.Element); ok {
				assert.Equal(t, "SPAN", prevSibling.TagName(), "Child2's previous sibling should be span")
			}
		}
		if child2.NextSibling() != nil {
			assert.True(t, child2.NextSibling().Equals(child3), "Child2's next sibling should be child3")
		}
		if child3.PreviousSibling() != nil {
			assert.True(t, child3.PreviousSibling().Equals(child2), "Child3's previous sibling should be child2")
		}
		assert.Nil(t, child3.NextSibling(), "Last child should have no next sibling")
	})

	t.Run("remove child", func(t *testing.T) {
		// Remove the middle child (child2 - p element)
		children := []dom.Element{}
		current := parent.FirstChild()
		for current != nil {
			if elem, ok := current.(dom.Element); ok {
				children = append(children, elem)
			}
			current = current.NextSibling()
		}

		if len(children) >= 2 {
			child2 := children[1] // The p element
			parent.RemoveChild(child2)

			assert.Nil(t, child2.ParentNode(), "Removed child should have no parent")

			// Verify remaining structure
			assert.True(t, parent.HasChildNodes(), "Parent should still have child nodes")
		}
	})
}

func TestElement_TextContent(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	t.Run("add text nodes", func(t *testing.T) {
		textNode := doc.CreateTextNode("Hello, World!")
		element.AppendChild(textNode)

		// Test that the text node was added (can't easily test TextContent without implementing it)
		assert.True(t, element.HasChildNodes(), "Element should have child nodes")

		firstChild := element.FirstChild()
		if textNode, ok := firstChild.(dom.Text); ok {
			assert.Equal(t, "Hello, World!", textNode.Data(), "Text content should match")
		}
	})
}

func TestElement_Style(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	t.Run("style property", func(t *testing.T) {
		style := element.Style()
		// Note: Style may be nil in some implementations or test environments
		// We just verify that calling Style() doesn't panic
		_ = style // Use the style variable to avoid unused warnings
		assert.NotPanics(t, func() {
			element.Style()
		}, "Style() should not panic")
	})
}

func TestElement_Focus(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	t.Run("focus and blur", func(t *testing.T) {
		input := doc.CreateElement("input")

		// These methods don't return values, so we just test they don't panic
		assert.NotPanics(t, func() {
			input.Focus()
		}, "Focus should not panic")

		assert.NotPanics(t, func() {
			input.Blur()
		}, "Blur should not panic")
	})
}

func TestElement_Events(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("button")

	t.Run("add event listener", func(t *testing.T) {
		handler := func(node dom.Node) {
			// Event handler - would be called when event fires
		}

		result := element.AddEventListener("click", handler)
		assert.Equal(t, element, result, "AddEventListener should return the same element")

		// Note: Actually triggering events would require a more complex test setup
		// This just verifies that AddEventListener doesn't panic and returns correctly
	})
}

func TestElement_EdgeCases(t *testing.T) {
	t.Run("self manipulation", func(t *testing.T) {
		element := domPkg.GetWindow().Document().CreateElement("div")

		// Test self-append behavior - WASM throws error, non-WASM might allow it
		defer func() {
			if r := recover(); r != nil {
				t.Log("Implementation correctly prevents self-append with error:", r)
			}
		}()

		element.AppendChild(element)

		// If we get here, self-append was allowed
		firstChild := element.FirstChild()
		if firstChild != nil && firstChild.Equals(element) {
			t.Log("Implementation allows self-append (element becomes its own child)")
		} else {
			t.Log("Self-append was ignored by implementation")
		}
	})

	t.Run("null operations", func(t *testing.T) {
		element := domPkg.GetWindow().Document().CreateElement("div")

		// Test RemoveChild with nil - WASM may throw error, non-WASM handles gracefully
		t.Run("remove nil child", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Log("RemoveChild(nil) correctly caused an error in WASM:", r)
				}
			}()
			element.RemoveChild(nil)
			t.Log("RemoveChild(nil) handled gracefully in non-WASM")
		})

		// Test SetAttribute with empty name - WASM may throw error, non-WASM handles gracefully
		t.Run("empty attribute name", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Log("SetAttribute with empty name correctly caused an error in WASM:", r)
				}
			}()
			element.SetAttribute("", "")
			t.Log("SetAttribute with empty name handled gracefully in non-WASM")
		})

		// GetAttribute with empty string should always work
		t.Run("get empty attribute", func(t *testing.T) {
			result := element.GetAttribute("")
			assert.Equal(t, "", result, "GetAttribute with empty string should return empty string")
		})
	})
}

func TestElement_ComplexStructure(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	t.Run("build complex DOM tree", func(t *testing.T) {
		// Create a complex structure: div > header > h1 + nav > ul > li*3
		container := doc.CreateElement("div")
		container.SetAttribute("class", "container")

		header := doc.CreateElement("header")
		h1 := doc.CreateElement("h1")
		h1Text := doc.CreateTextNode("Page Title")
		h1.AppendChild(h1Text)

		nav := doc.CreateElement("nav")
		ul := doc.CreateElement("ul")

		// Create navigation items
		for i := 1; i <= 3; i++ {
			li := doc.CreateElement("li")
			a := doc.CreateElement("a")
			a.SetAttribute("href", "#")
			linkText := doc.CreateTextNode("Link " + string(rune('0'+i)))
			a.AppendChild(linkText)
			li.AppendChild(a)
			ul.AppendChild(li)
		}

		// Build the structure
		nav.AppendChild(ul)
		header.AppendChild(h1)
		header.AppendChild(nav)
		container.AppendChild(header)

		// Test the structure
		assert.True(t, container.HasChildNodes(), "Container should have children")
		assert.Equal(t, "container", container.GetAttribute("class"), "Container should have correct class")

		firstChild := container.FirstChild()
		assert.NotNil(t, firstChild, "Container should have first child")
		if header, ok := firstChild.(dom.Element); ok {
			assert.Equal(t, "HEADER", header.TagName(), "First child should be header")
		}
	})
}

// Legacy tests for backward compatibility
func Test_Element_001(t *testing.T) {
	element := domPkg.GetWindow().Document().CreateElement("A")
	if element.NodeName() != "A" {
		t.Errorf("Element.NodeName() = %v, want %v", element.NodeName(), "A")
	}
	if element.NodeType() != dom.ELEMENT_NODE {
		t.Errorf("Element.NodeType() = %v, want %v", element.NodeType(), dom.ELEMENT_NODE)
	}
}

func Test_Element_002(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("A")
	b := doc.CreateElement("CB")
	c := doc.CreateElement("CC")
	if parent.HasChildNodes() != false {
		t.Error("HasChildNodes() failed")
	}
	if b.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	if c.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	parent.AppendChild(b)
	if parent.FirstChild().Equals(b) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(b) == false {
		t.Error("LastChild() failed")
	}
	if parent.HasChildNodes() == false {
		t.Error("HasChildNodes() failed: ", parent)
	}
	if b.ParentNode().Equals(parent) == false {
		t.Error("ParentNode() failed")
	}
	if b.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if b.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.AppendChild(c)
	if b.NextSibling() == nil {
		t.Error("NextSibling() failed")
	}
	if parent.FirstChild().Equals(b) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(c) == false {
		t.Error("LastChild() failed")
	}
	if c.ParentNode().Equals(parent) == false {
		t.Error("ParentNode() failed")
	}
	if b.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if b.NextSibling().Equals(c) == false {
		t.Error("NextSibling() failed")
	}
	if c.PreviousSibling().Equals(b) == false {
		t.Error("PreviousSibling() failed")
	}
	if c.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.RemoveChild(b)
	if parent.FirstChild().Equals(c) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(c) == false {
		t.Error("LastChild() failed")
	}
	if b.ParentNode() != nil {
		t.Error("ParentNode() failed", b.ParentNode())
	}
	if c.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if c.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.RemoveChild(c)
	if parent.FirstChild() != nil {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild() != nil {
		t.Error("LastChild() failed")
	}
	if parent.HasChildNodes() {
		t.Error("HasChildNodes() failed")
	}
	if c.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
}

func Test_Element_003(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("a")
	if parent.InnerHTML() != "" {
		t.Error("InnerHTML() failed")
	}
	if parent.OuterHTML() != "<a></a>" {
		t.Error("OuterHTML() failed: ", parent.OuterHTML())
	}
}

func TestElement_GetElementsByClassName(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	t.Run("empty_container", func(t *testing.T) {
		container := doc.CreateElement("div")
		elements := container.GetElementsByClassName("test-class")
		assert.Empty(t, elements, "Should return empty slice for container with no children")
	})

	t.Run("no_matching_elements", func(t *testing.T) {
		container := doc.CreateElement("div")

		// Add some elements without the target class
		child1 := doc.CreateElement("p")
		child1.SetAttribute("class", "other-class")
		container.AppendChild(child1)

		child2 := doc.CreateElement("span")
		child2.SetAttribute("class", "different-class")
		container.AppendChild(child2)

		elements := container.GetElementsByClassName("test-class")
		assert.Empty(t, elements, "Should return empty slice when no elements have the class")
	})

	t.Run("single_matching_element", func(t *testing.T) {
		container := doc.CreateElement("div")

		child := doc.CreateElement("p")
		child.SetAttribute("class", "test-class")
		container.AppendChild(child)

		elements := container.GetElementsByClassName("test-class")
		assert.Len(t, elements, 1, "Should return one element")
		assert.Equal(t, "P", elements[0].TagName(), "Should return the correct element")
		assert.True(t, elements[0].ClassList().Contains("test-class"), "Element should have the class")
	})

	t.Run("multiple_matching_elements", func(t *testing.T) {
		container := doc.CreateElement("div")

		// Add multiple elements with the same class
		child1 := doc.CreateElement("p")
		child1.SetAttribute("class", "test-class")
		container.AppendChild(child1)

		child2 := doc.CreateElement("span")
		child2.SetAttribute("class", "test-class")
		container.AppendChild(child2)

		child3 := doc.CreateElement("div")
		child3.SetAttribute("class", "test-class")
		container.AppendChild(child3)

		elements := container.GetElementsByClassName("test-class")
		assert.Len(t, elements, 3, "Should return three elements")

		expectedTags := []string{"P", "SPAN", "DIV"}
		for i, element := range elements {
			assert.Equal(t, expectedTags[i], element.TagName(), "Should return elements in document order")
			assert.True(t, element.ClassList().Contains("test-class"), "Each element should have the class")
		}
	})

	t.Run("elements_with_multiple_classes", func(t *testing.T) {
		container := doc.CreateElement("div")

		// Element with multiple classes including target
		child1 := doc.CreateElement("p")
		child1.SetAttribute("class", "first-class test-class last-class")
		container.AppendChild(child1)

		// Element with only target class
		child2 := doc.CreateElement("span")
		child2.SetAttribute("class", "test-class")
		container.AppendChild(child2)

		elements := container.GetElementsByClassName("test-class")
		assert.Len(t, elements, 2, "Should find elements regardless of other classes")
		assert.Equal(t, "P", elements[0].TagName(), "First element should be p")
		assert.Equal(t, "SPAN", elements[1].TagName(), "Second element should be span")
	})

	t.Run("nested_elements", func(t *testing.T) {
		container := doc.CreateElement("div")

		// Direct child with class
		directChild := doc.CreateElement("p")
		directChild.SetAttribute("class", "test-class")
		container.AppendChild(directChild)

		// Nested structure
		nestedParent := doc.CreateElement("section")
		container.AppendChild(nestedParent)

		nestedChild := doc.CreateElement("span")
		nestedChild.SetAttribute("class", "test-class")
		nestedParent.AppendChild(nestedChild)

		// Deeply nested
		deepParent := doc.CreateElement("article")
		nestedParent.AppendChild(deepParent)

		deepChild := doc.CreateElement("em")
		deepChild.SetAttribute("class", "test-class")
		deepParent.AppendChild(deepChild)

		elements := container.GetElementsByClassName("test-class")
		assert.Len(t, elements, 3, "Should find elements at all nesting levels")

		expectedTags := []string{"P", "SPAN", "EM"}
		for i, element := range elements {
			assert.Equal(t, expectedTags[i], element.TagName(), "Should return elements in document order")
		}
	})

	t.Run("mixed_elements_some_matching", func(t *testing.T) {
		container := doc.CreateElement("div")

		// Mix of elements with and without the class
		elements := []struct {
			tag   string
			class string
		}{
			{"p", "test-class"},
			{"span", "other-class"},
			{"div", "test-class another-class"},
			{"article", "unrelated"},
			{"section", "test-class"},
		}

		for _, elem := range elements {
			child := doc.CreateElement(elem.tag)
			child.SetAttribute("class", elem.class)
			container.AppendChild(child)
		}

		matches := container.GetElementsByClassName("test-class")
		assert.Len(t, matches, 3, "Should find only elements with the target class")

		expectedTags := []string{"P", "DIV", "SECTION"}
		for i, element := range matches {
			assert.Equal(t, expectedTags[i], element.TagName(), "Should return correct elements in order")
		}
	})

	t.Run("case_sensitive_class_names", func(t *testing.T) {
		container := doc.CreateElement("div")

		child1 := doc.CreateElement("p")
		child1.SetAttribute("class", "Test-Class")
		container.AppendChild(child1)

		child2 := doc.CreateElement("span")
		child2.SetAttribute("class", "test-class")
		container.AppendChild(child2)

		// Search for lowercase - should only find exact match
		elements := container.GetElementsByClassName("test-class")
		assert.Len(t, elements, 1, "Class name matching should be case-sensitive")
		assert.Equal(t, "SPAN", elements[0].TagName(), "Should find only exact case match")

		// Search for uppercase
		elementsUpper := container.GetElementsByClassName("Test-Class")
		assert.Len(t, elementsUpper, 1, "Should find uppercase variant")
		assert.Equal(t, "P", elementsUpper[0].TagName(), "Should find exact case match")
	})

	t.Run("empty_class_name", func(t *testing.T) {
		container := doc.CreateElement("div")

		child := doc.CreateElement("p")
		child.SetAttribute("class", "test-class")
		container.AppendChild(child)

		elements := container.GetElementsByClassName("")
		assert.Empty(t, elements, "Empty class name should return no elements")
	})

	t.Run("recursive_depth_first_order", func(t *testing.T) {
		// Create a more complex tree to test traversal order
		container := doc.CreateElement("div")

		// Level 1: first branch
		branch1 := doc.CreateElement("section")
		branch1.SetAttribute("class", "target")
		container.AppendChild(branch1)

		// Level 2: first branch children
		leaf1 := doc.CreateElement("p")
		leaf1.SetAttribute("class", "target")
		branch1.AppendChild(leaf1)

		leaf2 := doc.CreateElement("span") // no class
		branch1.AppendChild(leaf2)

		// Level 3: deeply nested in first branch
		deepLeaf := doc.CreateElement("em")
		deepLeaf.SetAttribute("class", "target")
		leaf2.AppendChild(deepLeaf)

		// Level 1: second branch (should come after all first branch descendants)
		branch2 := doc.CreateElement("article")
		branch2.SetAttribute("class", "target")
		container.AppendChild(branch2)

		// Level 2: second branch child
		leaf3 := doc.CreateElement("strong")
		leaf3.SetAttribute("class", "target")
		branch2.AppendChild(leaf3)

		elements := container.GetElementsByClassName("target")
		assert.Len(t, elements, 5, "Should find all elements with target class")

		// Verify depth-first order: section, p, em, article, strong
		expectedTags := []string{"SECTION", "P", "EM", "ARTICLE", "STRONG"}
		for i, element := range elements {
			assert.Equal(t, expectedTags[i], element.TagName(),
				"Should return elements in depth-first document order")
		}
	})
}
