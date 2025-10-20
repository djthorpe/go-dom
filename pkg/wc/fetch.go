//go:build js

package wc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	dom "github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

// Fetcher provides periodic API fetching functionality
type Fetcher struct {
	baseURL  string
	interval time.Duration
	ticker   *time.Ticker
	stopChan chan bool
	running  bool
	client   *http.Client
	headers  http.Header
}

// FetchResponse represents a fetch response
type FetchResponse struct {
	Status     int
	StatusText string
	Data       []byte
	Error      error
}

// FetchCallback is called when a fetch completes
type FetchCallback func(*FetchResponse)

// FetchOptions configures fetch behavior
type FetchOptions struct {
	Method      string            // HTTP method (GET, POST, etc.)
	Headers     map[string]string // Custom headers
	Body        io.Reader         // Request body for POST/PUT
	Mode        string            // Fetch API mode: "cors", "no-cors", "same-origin"
	Credentials string            // Fetch API credentials: "omit", "same-origin", "include"
	Redirect    string            // Fetch API redirect: "follow", "error", "manual"
}

////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewFetcher creates a new API fetcher with a base URL and optional fetch interval.
// If interval is 0, fetching will only occur when manually triggered.
func NewFetcher(baseURL string, interval time.Duration) *Fetcher {
	return &Fetcher{
		baseURL:  baseURL,
		interval: interval,
		stopChan: make(chan bool),
		running:  false,
		client:   http.DefaultClient,
		headers:  make(http.Header),
	}
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// SetHeader sets a custom header for all requests
func (f *Fetcher) SetHeader(key, value string) *Fetcher {
	f.headers.Set(key, value)
	return f
}

// SetFetchMode sets the Fetch API mode (cors, no-cors, same-origin)
func (f *Fetcher) SetFetchMode(mode string) *Fetcher {
	f.headers.Set("js.fetch:mode", mode)
	return f
}

// SetFetchCredentials sets the Fetch API credentials (omit, same-origin, include)
func (f *Fetcher) SetFetchCredentials(credentials string) *Fetcher {
	f.headers.Set("js.fetch:credentials", credentials)
	return f
}

// SetFetchRedirect sets the Fetch API redirect behavior (follow, error, manual)
func (f *Fetcher) SetFetchRedirect(redirect string) *Fetcher {
	f.headers.Set("js.fetch:redirect", redirect)
	return f
}

// Fetch performs a single GET request to the specified path (relative to baseURL)
func (f *Fetcher) Fetch(path string, callback FetchCallback) {
	go f.doFetch(path, nil, callback)
}

// FetchWithOptions performs a single request with custom options
func (f *Fetcher) FetchWithOptions(path string, opts *FetchOptions, callback FetchCallback) {
	go f.doFetch(path, opts, callback)
}

// Start begins periodic fetching if an interval was configured
func (f *Fetcher) Start(path string, callback FetchCallback) error {
	if f.interval == 0 {
		return fmt.Errorf("cannot start periodic fetching: no interval configured")
	}

	if f.running {
		return fmt.Errorf("fetcher already running")
	}

	f.running = true
	f.ticker = time.NewTicker(f.interval)

	// Perform initial fetch immediately
	go f.doFetch(path, nil, callback)

	// Start periodic fetching
	go func() {
		for {
			select {
			case <-f.ticker.C:
				f.doFetch(path, nil, callback)
			case <-f.stopChan:
				return
			}
		}
	}()

	return nil
}

// Stop halts periodic fetching
func (f *Fetcher) Stop() {
	if !f.running {
		return
	}

	f.running = false
	if f.ticker != nil {
		f.ticker.Stop()
	}
	f.stopChan <- true
}

// IsRunning returns whether periodic fetching is active
func (f *Fetcher) IsRunning() bool {
	return f.running
}

// SetInterval updates the fetch interval (only takes effect after restart)
func (f *Fetcher) SetInterval(interval time.Duration) {
	f.interval = interval
}

// GetBaseURL returns the configured base URL
func (f *Fetcher) GetBaseURL() string {
	return f.baseURL
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// doFetch performs the actual fetch operation using net/http
func (f *Fetcher) doFetch(path string, opts *FetchOptions, callback FetchCallback) {
	url := f.baseURL + path

	// Determine HTTP method
	method := "GET"
	var body io.Reader
	if opts != nil {
		if opts.Method != "" {
			method = opts.Method
		}
		body = opts.Body
	}

	// Create the HTTP request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		callback(&FetchResponse{
			Error: fmt.Errorf("failed to create request: %w", err),
		})
		return
	}

	// Add default headers from fetcher
	for key, values := range f.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Add custom headers from options
	if opts != nil {
		// Add custom headers
		for key, value := range opts.Headers {
			req.Header.Set(key, value)
		}

		// Add fetch-specific options as special headers
		if opts.Mode != "" {
			req.Header.Set("js.fetch:mode", opts.Mode)
		}
		if opts.Credentials != "" {
			req.Header.Set("js.fetch:credentials", opts.Credentials)
		}
		if opts.Redirect != "" {
			req.Header.Set("js.fetch:redirect", opts.Redirect)
		}
	}

	// Perform the request in a goroutine to avoid blocking
	go func() {
		resp, err := f.client.Do(req)
		if err != nil {
			callback(&FetchResponse{
				Error: fmt.Errorf("fetch failed: %w", err),
			})
			return
		}
		defer resp.Body.Close()

		// Read the response body
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			callback(&FetchResponse{
				Status:     resp.StatusCode,
				StatusText: resp.Status,
				Error:      fmt.Errorf("failed to read response: %w", err),
			})
			return
		}

		// Call the callback with the response
		callback(&FetchResponse{
			Status:     resp.StatusCode,
			StatusText: resp.Status,
			Data:       data,
			Error:      nil,
		})
	}()
}

////////////////////////////////////////////////////////////////////////////////
// HELPER METHODS

// ParseJSON is a helper to parse JSON response data
func (r *FetchResponse) ParseJSON(v interface{}) error {
	if r.Error != nil {
		return r.Error
	}
	if r.Status < 200 || r.Status >= 300 {
		return fmt.Errorf("HTTP error: %d %s", r.Status, r.StatusText)
	}
	return json.Unmarshal(r.Data, v)
}

// String returns the response data as a string
func (r *FetchResponse) String() string {
	return string(r.Data)
}

// IsSuccess returns true if the HTTP status indicates success (2xx)
func (r *FetchResponse) IsSuccess() bool {
	return r.Error == nil && r.Status >= 200 && r.Status < 300
}

// UpdateElement is a helper to update a DOM element with the response
// This replaces all child nodes with a text node containing the response data
func (r *FetchResponse) UpdateElement(element dom.Element) error {
	if r.Error != nil {
		return r.Error
	}

	if !r.IsSuccess() {
		return fmt.Errorf("HTTP %d: %s", r.Status, r.StatusText)
	}

	// Remove all existing children
	for element.HasChildNodes() {
		element.RemoveChild(element.FirstChild())
	}

	// Create a text node with the response data
	doc := element.OwnerDocument()
	textNode := doc.CreateTextNode(r.String())
	element.AppendChild(textNode)

	return nil
}
