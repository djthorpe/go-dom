package bootstrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	assert := assert.New(t)

	router := Router()
	assert.NotNil(router)
	assert.NotNil(router.Element())
	assert.NotNil(router.body)
	assert.True(router.body.ClassList().Contains("container-fluid"))
}

func TestRoute(t *testing.T) {
	assert := assert.New(t)

	route := Route("^home$", "^/$")
	assert.NotNil(route)
	assert.NotNil(route.Element())
}

func TestRouteMatches(t *testing.T) {
	assert := assert.New(t)

	route := Route("^home$", "^about$")

	assert.True(route.Matches("#home"))
	assert.True(route.Matches("home"))
	assert.True(route.Matches("#about"))
	assert.False(route.Matches("#other"))
	assert.False(route.Matches(""))
}

func TestRouteMatchesEmpty(t *testing.T) {
	assert := assert.New(t)

	// Route with no patterns matches empty hash
	route := Route()
	assert.True(route.Matches(""))
	assert.True(route.Matches("#"))
}

func TestRouterAddRoute(t *testing.T) {
	assert := assert.New(t)

	router := Router()
	route1 := Route("^home$")
	route2 := Route("^about$")

	router.AddRoute(route1)
	router.AddRoute(route2)

	assert.Len(router.routes, 2)
}

func TestRouterNavigate(t *testing.T) {
	assert := assert.New(t)

	router := Router()
	homeRoute := Route("^home$")
	aboutRoute := Route("^about$")

	router.AddRoute(homeRoute)
	router.AddRoute(aboutRoute)

	// Navigate to home
	router.Navigate("#home")
	assert.Equal("display:block", homeRoute.root.GetAttribute("style"))
	assert.Equal("display:none", aboutRoute.root.GetAttribute("style"))

	// Navigate to about
	router.Navigate("#about")
	assert.Equal("display:none", homeRoute.root.GetAttribute("style"))
	assert.Equal("display:block", aboutRoute.root.GetAttribute("style"))
}

func TestRouteAppend(t *testing.T) {
	assert := assert.New(t)

	route := Route("^test$")

	// Test appending string
	route.Append("Hello World")
	assert.Equal(1, len(route.root.ChildNodes()))

	// Test appending component
	para := Para()
	para.Append("Test paragraph")
	route.Append(para)
	assert.Equal(2, len(route.root.ChildNodes()))
}

func TestRouteShowHide(t *testing.T) {
	assert := assert.New(t)

	route := Route("^test$")

	// Initially hidden
	assert.Equal("display:none", route.root.GetAttribute("style"))

	// Show
	route.Show()
	assert.Equal("display:block", route.root.GetAttribute("style"))

	// Hide
	route.Hide()
	assert.Equal("display:none", route.root.GetAttribute("style"))
}
