package dom

///////////////////////////////////////////////////////////////////////////////
// INTERFACES

// Fetch implements https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API
type Fetch interface {
	// Methods
	Then(func(HTTPResponse) Fetch)
}

type HTTPResponse interface {
	// Methods
}
