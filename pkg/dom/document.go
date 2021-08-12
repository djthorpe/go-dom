// +build !w

package dom

import "github.com/djthorpe/go-dom"

type document struct {
	node
}

func Document() dom.Document {
	return &document{}
}
