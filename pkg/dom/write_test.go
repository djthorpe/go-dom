package dom_test

import (
	"bytes"
	"strings"
	"testing"

	domPkg "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

func TestWrite_Document(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	// Add content to body
	body := doc.Body()
	if body != nil {
		p := doc.CreateElement("p")
		p.AppendChild(doc.CreateTextNode("Test content"))
		body.AppendChild(p)
	}

	buf := new(bytes.Buffer)
	n, err := win.Write(buf, doc)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output (%d bytes):\n%s", n, output)

	// Verify DOCTYPE is present
	if !strings.Contains(output, "<!DOCTYPE html>") {
		t.Error("Output missing DOCTYPE")
	}

	// Verify basic structure
	if !strings.Contains(output, "<html>") {
		t.Error("Output missing <html>")
	}

	// Verify our added content
	if body != nil && !strings.Contains(output, "<p>Test content</p>") {
		t.Error("Output missing test content")
	}
}

func TestWrite_Element(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	div := doc.CreateElement("div")
	div.SetAttribute("class", "container")
	div.SetAttribute("id", "main")
	p := doc.CreateElement("p")
	p.AppendChild(doc.CreateTextNode("Hello"))
	div.AppendChild(p)

	buf := new(bytes.Buffer)
	_, err := win.Write(buf, div)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output: %s", output)

	// Verify attributes
	if !strings.Contains(output, "class=\"container\"") {
		t.Error("Output missing class attribute")
	}
	if !strings.Contains(output, "id=\"main\"") {
		t.Error("Output missing id attribute")
	}
	if !strings.Contains(output, "<p>Hello</p>") {
		t.Error("Output missing paragraph content")
	}
}

func TestWrite_TextNode(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	text := doc.CreateTextNode("Hello, World!")
	buf := new(bytes.Buffer)
	_, err := win.Write(buf, text)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	if output != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got '%s'", output)
	}
}

func TestWrite_Comment(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	comment := doc.CreateComment("This is a comment")
	buf := new(bytes.Buffer)
	_, err := win.Write(buf, comment)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "<!--This is a comment-->") {
		t.Errorf("Expected comment, got '%s'", output)
	}
}

func TestWriteIndented_Document(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	// Add to body
	body := doc.Body()
	if body != nil {
		div := doc.CreateElement("div")
		div.SetAttribute("class", "container")
		p := doc.CreateElement("p")
		p.AppendChild(doc.CreateTextNode("Hello"))
		div.AppendChild(p)
		body.AppendChild(div)
	}

	buf := new(bytes.Buffer)
	n, err := win.WriteIndented(buf, doc, "  ")
	if err != nil {
		t.Fatalf("WriteIndented failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output (%d bytes):\n%s", n, output)

	// Verify indentation
	lines := strings.Split(output, "\n")

	// DOCTYPE should not be indented
	if !strings.HasPrefix(lines[0], "<!DOCTYPE") {
		t.Error("DOCTYPE should not be indented")
	}

	// Check that elements are indented
	foundIndentedHead := false
	foundIndentedBody := false
	for _, line := range lines {
		if strings.HasPrefix(line, "  <head>") {
			foundIndentedHead = true
		}
		if strings.HasPrefix(line, "  <body>") {
			foundIndentedBody = true
		}
	}

	if !foundIndentedHead {
		t.Error("Expected <head> to be indented with 2 spaces")
	}
	if !foundIndentedBody {
		t.Error("Expected <body> to be indented with 2 spaces")
	}
}

func TestWriteIndented_Tabs(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	html := doc.CreateElement("html")
	head := doc.CreateElement("head")
	html.AppendChild(head)
	doc.AppendChild(html)

	buf := new(bytes.Buffer)
	_, err := win.WriteIndented(buf, doc, "\t")
	if err != nil {
		t.Fatalf("WriteIndented failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output:\n%s", output)

	// Check for tab indentation
	if !strings.Contains(output, "\t<head>") {
		t.Error("Expected <head> to be indented with tab")
	}
}

func TestWriteIndented_InlineText(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	// Elements with single text child should be on one line
	p := doc.CreateElement("p")
	p.AppendChild(doc.CreateTextNode("Single line text"))

	buf := new(bytes.Buffer)
	_, err := win.WriteIndented(buf, p, "  ")
	if err != nil {
		t.Fatalf("WriteIndented failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output: %q", output)

	// Should be inline
	if strings.Contains(output, "\n  ") {
		t.Error("Text-only paragraph should be inline, not indented")
	}
	expected := "<p>Single line text</p>\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestWriteIndented_MixedContent(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	div := doc.CreateElement("div")
	p1 := doc.CreateElement("p")
	p1.AppendChild(doc.CreateTextNode("First"))
	p2 := doc.CreateElement("p")
	p2.AppendChild(doc.CreateTextNode("Second"))

	div.AppendChild(p1)
	div.AppendChild(p2)

	buf := new(bytes.Buffer)
	_, err := win.WriteIndented(buf, div, "  ")
	if err != nil {
		t.Fatalf("WriteIndented failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output:\n%s", output)

	// Should have multiple levels of indentation
	if !strings.Contains(output, "  <p>First</p>") {
		t.Error("Expected indented <p> elements")
	}
}

func TestWriteIndented_DefaultIndent(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	div := doc.CreateElement("div")
	doc.AppendChild(div)

	buf := new(bytes.Buffer)
	// Pass empty string - should default to 2 spaces
	_, err := win.WriteIndented(buf, div, "")
	if err != nil {
		t.Fatalf("WriteIndented failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output: %q", output)

	// Should use default indent (tested by checking output is not empty)
	if len(output) == 0 {
		t.Error("Expected non-empty output")
	}
}

func TestWriteIndented_Comment(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	div := doc.CreateElement("div")
	comment := doc.CreateComment("Test comment")
	div.AppendChild(comment)

	buf := new(bytes.Buffer)
	_, err := win.WriteIndented(buf, div, "  ")
	if err != nil {
		t.Fatalf("WriteIndented failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output:\n%s", output)

	// Comment should be indented
	if !strings.Contains(output, "  <!--Test comment-->") {
		t.Error("Expected comment to be indented")
	}
}

func TestWrite_ErrorCases(t *testing.T) {
	win := domPkg.GetWindow()

	buf := new(bytes.Buffer)
	_, err := win.Write(buf, nil)
	if err == nil {
		t.Error("Expected error when writing nil node")
	}

	_, err = win.WriteIndented(buf, nil, "  ")
	if err == nil {
		t.Error("Expected error when writing nil node with indent")
	}
}
