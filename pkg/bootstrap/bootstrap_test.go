package bootstrap_test

import (
	"testing"

	// Modules
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	assert "github.com/stretchr/testify/assert"
)

func Test_App_001(t *testing.T) {
	assert := assert.New(t)

	// Create the application
	app := bs.New()
	assert.NotNil(app)

	// The DocType should be 'html'
	doctype := app.Doctype()
	assert.NotNil(doctype)
	assert.Equal("html", doctype.NodeName())

	// The root element should be a div with the correct id
	root := app.Root()
	assert.NotNil(root)
	assert.Equal("DIV", root.NodeName())
	assert.Equal(bs.AppIdentifier, root.GetAttribute("id"))
}

func Test_App_002(t *testing.T) {
	assert := assert.New(t)

	// Create the application
	app := bs.New()

	// Append a text node to the root
	app.Append("Hello, World!")
	assert.Len(app.Root().ChildNodes(), 1)
	assert.Equal("Hello, World!", app.Root().TextContent())
}
