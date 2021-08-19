package dom_test

import (
	"bytes"
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

func Test_Window_002(t *testing.T) {
	window := NewWindow()
	tests := []struct {
		mime, data string
	}{
		{"text/xml", "<warning>Beware of the tiger</warning>"},
		{"image/svg+xml", "<circle cx=\"50\" cy=\"50\" r=\"50\"/>"},
		{"text/html", "<strong>Beware of the leopard</strong>"},
		{"text/html", "<!DOCTYPE html><html><head><title>test</title><meta charset=\"utf-8\"></head><body></body></html>"},
	}
	for _, test := range tests {
		buf := bytes.NewBufferString(test.data)
		if doc, err := window.Read(buf, test.mime); err != nil {
			t.Error(err)
		} else {
			t.Log("doc=", doc)
		}
	}
}
