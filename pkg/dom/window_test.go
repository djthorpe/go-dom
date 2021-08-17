package dom_test

import (
	"testing"

	// Modules
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func Test_Window_001(t *testing.T) {
	window := NewWindow()
	if window == nil {
		t.Fatal("Window() returned nil")
	} else {
		t.Log("window=", window)
	}

	document := window.Document()
	if document == nil {
		t.Fatal("Document() returned nil")
	} else {
		t.Log("document=", document)
	}
}
