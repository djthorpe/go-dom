// +build w

package dom

import (
	"syscall/js"

	"github.com/djthorpe/go-dom"
)

func Document() dom.Document {
	return js.Global().Get("document")
}

func Window() dom.Window {
	return js.Global()
}
