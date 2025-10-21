package bootstrap_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	domPkg "github.com/djthorpe/go-wasmbuild/pkg/dom"
	"github.com/stretchr/testify/assert"
)

func TestContainer_Basic(t *testing.T) {
	container := bs.Container()

	assert.NotNil(t, container, "Container should not be nil")
	assert.NotNil(t, container.Element(), "Container element should not be nil")
	assert.Equal(t, dom.ELEMENT_NODE, container.Element().NodeType(), "Container should be an element node")
	assert.Equal(t, "DIV", container.Element().TagName(), "Container should be a div element")
}

func TestContainer_DefaultClass(t *testing.T) {
	container := bs.Container()
	element := container.Element()

	assert.True(t, element.HasAttribute("class"), "Container should have class attribute")
	assert.Equal(t, "container", element.GetAttribute("class"), "Container should have 'container' class")

	classList := element.ClassList()
	assert.NotNil(t, classList, "Container should have class list")
	assert.Equal(t, 1, classList.Length(), "Container should have exactly one class")
	assert.True(t, classList.Contains("container"), "Container should contain 'container' class")
}

func TestContainer_FluidOption(t *testing.T) {
	container := bs.Container(bs.WithContainerFluid())
	element := container.Element()

	assert.True(t, element.HasAttribute("class"), "Fluid container should have class attribute")
	assert.Equal(t, "container-fluid", element.GetAttribute("class"), "Fluid container should have 'container-fluid' class")

	classList := element.ClassList()
	assert.Equal(t, 1, classList.Length(), "Fluid container should have exactly one class")
	assert.True(t, classList.Contains("container-fluid"), "Fluid container should contain 'container-fluid' class")
	assert.False(t, classList.Contains("container"), "Fluid container should not contain 'container' class")
}

func TestContainer_OuterHTML(t *testing.T) {
	tests := []struct {
		name         string
		constructor  func() dom.Component
		expectedHTML string
	}{
		{
			name:         "default container",
			constructor:  func() dom.Component { return bs.Container() },
			expectedHTML: `<div class="container"></div>`,
		},
		{
			name:         "fluid container",
			constructor:  func() dom.Component { return bs.Container(bs.WithContainerFluid()) },
			expectedHTML: `<div class="container-fluid"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := tt.constructor()
			outerHTML := container.Element().OuterHTML()
			// Normalize to lowercase for comparison since different implementations may vary
			assert.Equal(t, tt.expectedHTML, strings.ToLower(outerHTML), "Outer HTML should match expected format")
		})
	}
}

func TestContainer_WithAdditionalClasses(t *testing.T) {
	tests := []struct {
		name            string
		options         []bs.Opt
		expectedClasses []string
	}{
		{
			name:            "container with single additional class",
			options:         []bs.Opt{bs.WithClass("my-custom-class")},
			expectedClasses: []string{"container", "my-custom-class"},
		},
		{
			name:            "container with multiple additional classes",
			options:         []bs.Opt{bs.WithClass("class1", "class2", "class3")},
			expectedClasses: []string{"container", "class1", "class2", "class3"},
		},
		{
			name:            "fluid container with additional classes",
			options:         []bs.Opt{bs.WithContainerFluid(), bs.WithClass("custom-fluid")},
			expectedClasses: []string{"container-fluid", "custom-fluid"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(tt.options...)
			classList := container.Element().ClassList()

			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Should have expected number of classes")

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Should contain class: %s", expectedClass)
			}
		})
	}
}

func TestContainer_ComponentInterface(t *testing.T) {
	container := bs.Container()

	// Test that container implements Component interface
	var component dom.Component = container
	assert.NotNil(t, component, "Container should implement Component interface")
	assert.Equal(t, container.Element(), component.Element(), "Element() method should return same element")
}

func TestContainer_ElementProperties(t *testing.T) {
	container := bs.Container()
	element := container.Element()

	t.Run("basic element properties", func(t *testing.T) {
		assert.Equal(t, "DIV", element.TagName(), "Element should be a div")
		assert.Equal(t, dom.ELEMENT_NODE, element.NodeType(), "Should be element node type")
		assert.Equal(t, "", element.InnerHTML(), "Should have empty inner HTML initially")
	})

	t.Run("attributes", func(t *testing.T) {
		assert.True(t, element.HasAttributes(), "Should have attributes")
		assert.True(t, element.HasAttribute("class"), "Should have class attribute")
		assert.False(t, element.HasAttribute("id"), "Should not have id attribute initially")
	})

	t.Run("child nodes", func(t *testing.T) {
		assert.False(t, element.HasChildNodes(), "Should have no child nodes initially")
		assert.Nil(t, element.FirstChild(), "First child should be nil")
		assert.Nil(t, element.LastChild(), "Last child should be nil")
	})
}

func TestContainer_ContentManipulation(t *testing.T) {
	container := bs.Container()
	element := container.Element()

	// We need to get a document to create child elements
	doc := element.OwnerDocument()
	if doc == nil {
		// Fallback if OwnerDocument is not implemented
		doc = domPkg.GetWindow().Document()
	}

	t.Run("add child elements", func(t *testing.T) {
		// Create and add child elements
		child1 := doc.CreateElement("p")
		child1.AppendChild(doc.CreateTextNode("Paragraph 1"))

		child2 := doc.CreateElement("span")
		child2.AppendChild(doc.CreateTextNode("Span content"))

		element.AppendChild(child1)
		element.AppendChild(child2)

		assert.True(t, element.HasChildNodes(), "Should have child nodes after adding")
		assert.NotNil(t, element.FirstChild(), "First child should not be nil")
		assert.NotNil(t, element.LastChild(), "Last child should not be nil")
		assert.NotEqual(t, "", element.InnerHTML(), "Inner HTML should not be empty after adding children")
	})

	t.Run("set attributes", func(t *testing.T) {
		element.SetAttribute("id", "my-container")
		element.SetAttribute("data-test", "container-test")

		assert.True(t, element.HasAttribute("id"), "Should have id attribute")
		assert.True(t, element.HasAttribute("data-test"), "Should have data-test attribute")
		assert.Equal(t, "my-container", element.GetAttribute("id"), "ID should match")
		assert.Equal(t, "container-test", element.GetAttribute("data-test"), "Data attribute should match")
	})
}

func TestContainer_ClassListOperations(t *testing.T) {
	container := bs.Container()
	classList := container.Element().ClassList()

	t.Run("initial state", func(t *testing.T) {
		assert.Equal(t, 1, classList.Length(), "Should start with one class")
		assert.True(t, classList.Contains("container"), "Should contain container class")
	})

	t.Run("add classes", func(t *testing.T) {
		classList.Add("custom-class", "another-class")
		assert.Equal(t, 3, classList.Length(), "Should have 3 classes")
		assert.True(t, classList.Contains("custom-class"), "Should contain custom-class")
		assert.True(t, classList.Contains("another-class"), "Should contain another-class")
		assert.True(t, classList.Contains("container"), "Should still contain container class")
	})

	t.Run("remove classes", func(t *testing.T) {
		classList.Remove("custom-class")
		assert.Equal(t, 2, classList.Length(), "Should have 2 classes")
		assert.False(t, classList.Contains("custom-class"), "Should not contain custom-class")
		assert.True(t, classList.Contains("another-class"), "Should still contain another-class")
		assert.True(t, classList.Contains("container"), "Should still contain container class")
	})

	t.Run("toggle classes", func(t *testing.T) {
		result := classList.Toggle("toggle-class")
		assert.True(t, result, "Toggle should return true for added class")
		assert.True(t, classList.Contains("toggle-class"), "Should contain toggle-class")

		result = classList.Toggle("toggle-class")
		assert.False(t, result, "Toggle should return false for removed class")
		assert.False(t, classList.Contains("toggle-class"), "Should not contain toggle-class")
	})
}

func TestContainer_MultipleContainers(t *testing.T) {
	// Test creating multiple containers
	container1 := bs.Container()
	container2 := bs.Container(bs.WithContainerFluid())
	container3 := bs.Container(bs.WithClass("custom"))

	assert.NotEqual(t, container1.Element(), container2.Element(), "Different containers should have different elements")
	assert.NotEqual(t, container1.Element(), container3.Element(), "Different containers should have different elements")
	assert.NotEqual(t, container2.Element(), container3.Element(), "Different containers should have different elements")

	// Verify each has correct classes
	assert.True(t, container1.Element().ClassList().Contains("container"), "Container1 should have container class")
	assert.True(t, container2.Element().ClassList().Contains("container-fluid"), "Container2 should have container-fluid class")
	assert.True(t, container3.Element().ClassList().Contains("container"), "Container3 should have container class")
	assert.True(t, container3.Element().ClassList().Contains("custom"), "Container3 should have custom class")
}

func TestContainer_OptionsOrder(t *testing.T) {
	tests := []struct {
		name            string
		options         []bs.Opt
		expectedClasses []string
		description     string
	}{
		{
			name:            "fluid then custom class",
			options:         []bs.Opt{bs.WithContainerFluid(), bs.WithClass("custom")},
			expectedClasses: []string{"container-fluid", "custom"},
			description:     "WithContainerFluid should replace container with container-fluid, then add custom",
		},
		{
			name:            "custom class then fluid",
			options:         []bs.Opt{bs.WithClass("custom"), bs.WithContainerFluid()},
			expectedClasses: []string{"container-fluid", "custom"},
			description:     "Order should not matter for final result",
		},
		{
			name:            "multiple custom classes with fluid",
			options:         []bs.Opt{bs.WithClass("class1"), bs.WithContainerFluid(), bs.WithClass("class2")},
			expectedClasses: []string{"container-fluid", "class1", "class2"},
			description:     "All custom classes should be preserved with fluid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(tt.options...)
			classList := container.Element().ClassList()

			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Should have expected number of classes: %s", tt.description)

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Should contain class %s: %s", expectedClass, tt.description)
			}
		})
	}
}

func TestContainer_EdgeCases(t *testing.T) {
	t.Run("empty options", func(t *testing.T) {
		container := bs.Container()
		assert.NotNil(t, container, "Container with no options should not be nil")
		assert.Equal(t, "container", container.Element().GetAttribute("class"),
			"Default container should have container class")
	})

	t.Run("nil options handling", func(t *testing.T) {
		// This should not panic
		assert.NotPanics(t, func() {
			container := bs.Container(nil)
			assert.NotNil(t, container, "Container should handle nil option gracefully")
		}, "Container creation should handle nil options")
	})

	t.Run("duplicate classes", func(t *testing.T) {
		container := bs.Container(bs.WithClass("container", "duplicate", "duplicate"))
		classList := container.Element().ClassList()

		// Count how many times "duplicate" appears
		duplicateCount := 0
		for _, class := range classList.Values() {
			if class == "duplicate" {
				duplicateCount++
			}
		}
		assert.Equal(t, 1, duplicateCount, "Duplicate classes should be handled properly")
	})

	t.Run("empty class names", func(t *testing.T) {
		// This tests edge case handling for empty strings
		assert.NotPanics(t, func() {
			container := bs.Container(bs.WithClass("", "valid-class", ""))
			classList := container.Element().ClassList()
			assert.True(t, classList.Contains("container"), "Should still have container class")
			assert.True(t, classList.Contains("valid-class"), "Should have valid class")
		}, "Should handle empty class names gracefully")
	})
}

// Legacy test for backward compatibility
func Test_Container_001(t *testing.T) {
	assert := assert.New(t)

	// Create a container
	container := bs.Container()
	assert.Equal(`<div class="container"></div>`, strings.ToLower(container.Element().OuterHTML()))

	// Create a fluid container
	fluid := bs.Container(bs.WithContainerFluid())
	assert.Equal(`<div class="container-fluid"></div>`, strings.ToLower(fluid.Element().OuterHTML()))
}
