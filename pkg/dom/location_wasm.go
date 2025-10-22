//go:build js && wasm
// +build js,wasm

package dom

import (
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type location struct {
	js.Value
}

var _ dom.Location = (*location)(nil)

/////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewLocation(jsLocation js.Value) dom.Location {
	return &location{jsLocation}
}

///////////////////////////////////////////////////////////////////////////////
// LOCATION IMPLEMENTATION

func (l *location) Hash() string {
	return l.Get("hash").String()
}

func (l *location) SetHash(hash string) {
	l.Set("hash", hash)
}

func (l *location) Href() string {
	return l.Get("href").String()
}

func (l *location) SetHref(href string) {
	l.Set("href", href)
}

func (l *location) Hostname() string {
	return l.Get("hostname").String()
}

func (l *location) Pathname() string {
	return l.Get("pathname").String()
}

func (l *location) Port() string {
	return l.Get("port").String()
}

func (l *location) Protocol() string {
	return l.Get("protocol").String()
}

func (l *location) Search() string {
	return l.Get("search").String()
}

func (l *location) SetSearch(search string) {
	l.Set("search", search)
}

func (l *location) Reload() {
	l.Call("reload")
}

func (l *location) Replace(url string) {
	l.Call("replace", url)
}
