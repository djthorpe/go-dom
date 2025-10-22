//go:build !js

package dom

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type location struct {
	hash     string
	href     string
	hostname string
	pathname string
	port     string
	protocol string
	search   string
}

/////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewLocation() dom.Location {
	return &location{
		protocol: "file:",
		pathname: "/",
	}
}

///////////////////////////////////////////////////////////////////////////////
// LOCATION IMPLEMENTATION

func (l *location) Hash() string {
	return l.hash
}

func (l *location) SetHash(hash string) {
	l.hash = hash
}

func (l *location) Href() string {
	return l.href
}

func (l *location) SetHref(href string) {
	l.href = href
}

func (l *location) Hostname() string {
	return l.hostname
}

func (l *location) Pathname() string {
	return l.pathname
}

func (l *location) Port() string {
	return l.port
}

func (l *location) Protocol() string {
	return l.protocol
}

func (l *location) Search() string {
	return l.search
}

func (l *location) SetSearch(search string) {
	l.search = search
}

func (l *location) Reload() {
	// No-op for non-WASM builds
}

func (l *location) Replace(url string) {
	// No-op for non-WASM builds
	l.href = url
}
