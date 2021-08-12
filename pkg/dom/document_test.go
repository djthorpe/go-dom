// +build !w

package dom_test

import (
	"testing"

	// Modules
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func Test_Document_001(t *testing.T) {
	doc := Document()
	t.Log(doc)
}
