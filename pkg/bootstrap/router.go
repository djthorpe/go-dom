package bootstrap

import (
	"regexp"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// router is a container that manages routes and displays content based on URL hash
type router struct {
	component
	routes []*route
}

// route represents a single route with one or more regex patterns
type route struct {
	component
	patterns []*regexp.Regexp
}

// Ensure router implements Component interface
var _ Component = (*router)(nil)
var _ Component = (*route)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Router creates a new router with a container-fluid body.
// The router manages multiple routes and displays content based on the URL hash.
// Use AddRoute() to add routes, and Navigate() to switch between them.
//
// Example:
//
//	router := Router()
//	homeRoute := Route("^$", "^home$").Append(Heading(1).Append("Home"))
//	router.AddRoute(homeRoute)
//	router.Navigate(window.Location().Hash())
func Router(opt ...Opt) *router {
	// Create wrapper div
	root := dom.GetWindow().Document().CreateElement("DIV")

	// Create container-fluid for the body where routes are appended
	body := dom.GetWindow().Document().CreateElement("DIV")
	body.SetAttribute("class", "container-fluid")
	root.AppendChild(body)

	// Apply options
	if opts, err := NewOpts(ContainerComponent); err != nil {
		panic(err)
	} else if err := opts.apply(opt...); err != nil {
		panic(err)
	} else {
		// Set class list on root
		classes := opts.classList.Values()
		if len(classes) > 0 {
			root.SetAttribute("class", strings.Join(classes, " "))
		}

		// Set attributes
		for key, value := range opts.attributes {
			root.SetAttribute(key, value)
		}
	}

	return &router{
		component: component{
			name: ContainerComponent,
			root: root,
			body: body,
		},
		routes: make([]*route, 0),
	}
}

// Route creates a new route with the given regex patterns.
// Each pattern is a regular expression that matches against the URL hash.
// Content added via Append() is displayed when the route matches.
//
// Example:
//
//	Route("^home$", "^$").Append(Heading(1).Append("Home Page"))
//	Route("^user/\\d+$").Append(Para().Append("User profile"))
func Route(patterns ...string) *route {
	// Compile regex patterns
	compiledPatterns := make([]*regexp.Regexp, 0, len(patterns))
	for _, pattern := range patterns {
		if re, err := regexp.Compile(pattern); err == nil {
			compiledPatterns = append(compiledPatterns, re)
		}
	}

	// Create content container
	root := dom.GetWindow().Document().CreateElement("DIV")
	root.SetAttribute("style", "display:none")

	return &route{
		component: component{
			name: ContainerComponent,
			root: root,
		},
		patterns: compiledPatterns,
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// AddRoute adds a route to the router
func (r *router) AddRoute(route *route) *router {
	r.routes = append(r.routes, route)
	r.body.AppendChild(route.Element())
	return r
}

// Navigate updates the router to show the route matching the given hash
func (r *router) Navigate(hash string) *router {
	// Hide all routes
	for _, route := range r.routes {
		route.Hide()
	}

	// Show matching route(s)
	for _, route := range r.routes {
		if route.Matches(hash) {
			route.Show()
		}
	}

	return r
}

///////////////////////////////////////////////////////////////////////////////
// ROUTE METHODS

// Matches returns true if the hash matches any of the route's patterns
func (r *route) Matches(hash string) bool {
	// Remove leading # if present
	if len(hash) > 0 && hash[0] == '#' {
		hash = hash[1:]
	}

	// Empty hash matches empty pattern
	if hash == "" && len(r.patterns) == 0 {
		return true
	}

	// Check all patterns
	for _, pattern := range r.patterns {
		if pattern.MatchString(hash) {
			return true
		}
	}

	return false
}

// Show makes the route visible
func (r *route) Show() *route {
	r.root.SetAttribute("style", "display:block")
	return r
}

// Hide makes the route invisible
func (r *route) Hide() *route {
	r.root.SetAttribute("style", "display:none")
	return r
}
