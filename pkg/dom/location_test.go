//go:build !js

package dom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	assert := assert.New(t)

	window := GetWindow()
	location := window.Location()

	assert.NotNil(location)
	assert.Equal("file:", location.Protocol())
	assert.Equal("/", location.Pathname())
	assert.Equal("", location.Hash())
}

func TestLocationHash(t *testing.T) {
	assert := assert.New(t)

	window := GetWindow()
	location := window.Location()

	// Test setting and getting hash
	location.SetHash("#test")
	assert.Equal("#test", location.Hash())

	location.SetHash("#another-section")
	assert.Equal("#another-section", location.Hash())
}

func TestLocationSearch(t *testing.T) {
	assert := assert.New(t)

	window := GetWindow()
	location := window.Location()

	// Test setting and getting search
	location.SetSearch("?foo=bar")
	assert.Equal("?foo=bar", location.Search())

	location.SetSearch("?key=value&other=thing")
	assert.Equal("?key=value&other=thing", location.Search())
}

func TestLocationHref(t *testing.T) {
	assert := assert.New(t)

	window := GetWindow()
	location := window.Location()

	// Test setting and getting href
	location.SetHref("https://example.com/path?query=1#section")
	assert.Equal("https://example.com/path?query=1#section", location.Href())
}
