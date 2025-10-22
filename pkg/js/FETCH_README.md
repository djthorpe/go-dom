# JavaScript Fetch API for Go WebAssembly

This package provides a Go-friendly wrapper around the JavaScript Fetch API for WebAssembly applications.

## Features

- **Promise-based API**: Chain `Then()` and `Catch()` callbacks
- **FetchOptions builder**: Fluent API for configuring requests
- **Response helpers**: Easy access to JSON, text, blob, and array buffer responses
- **Type-safe**: Wraps JavaScript values in Go types

## Basic Usage

### Simple GET Request

```go
import "github.com/djthorpe/go-wasmbuild/pkg/js"

js.Fetch("https://api.example.com/data").
    Then(func(response js.FetchResponse) {
        if response.Ok() {
            response.JSON(func(data js.Value) {
                title := data.Get("title").String()
                println("Title:", title)
            })
        }
    }).
    Catch(func(err js.Value) {
        println("Error:", err.Get("message").String())
    })
```

### POST Request with JSON

```go
opts := js.NewFetchOptions().
    Method("POST").
    Header("Content-Type", "application/json").
    Header("Authorization", "Bearer token123").
    Body(`{"name": "John", "age": 30}`).
    Mode("cors").
    Credentials("include")

js.Fetch("https://api.example.com/users", opts.Value()).
    Then(func(response js.FetchResponse) {
        if response.Ok() {
            println("User created! Status:", response.Status())
        }
    }).
    Catch(func(err js.Value) {
        println("Failed:", err.String())
    })
```

### GET Request with Text Response

```go
js.Fetch("https://example.com/document.txt").
    Then(func(response js.FetchResponse) {
        response.Text(func(text string) {
            println("Document content:", text)
        })
    }).
    Catch(func(err js.Value) {
        println("Error:", err.String())
    })
```

### Using Finally

```go
js.Fetch("https://api.example.com/data").
    Then(func(response js.FetchResponse) {
        // Handle response
    }).
    Catch(func(err js.Value) {
        // Handle error
    }).
    Finally(func() {
        println("Request completed (success or failure)")
    })
```

## FetchOptions Builder

The `FetchOptions` type provides a fluent API for configuring fetch requests:

```go
opts := js.NewFetchOptions().
    Method("PUT").                           // HTTP method
    Header("Content-Type", "application/json").
    Header("X-Custom-Header", "value").
    Body(`{"update": true}`).                // Request body
    Mode("cors").                            // cors, no-cors, same-origin
    Credentials("include").                  // omit, same-origin, include
    Cache("no-cache").                       // default, no-store, reload, etc.
    Redirect("follow").                      // follow, error, manual
    Referrer("https://example.com")

js.Fetch("/api/resource", opts.Value())
```

## FetchResponse Methods

The `FetchResponse` type wraps the JavaScript Response object:

- `Ok() bool` - Returns true for 2xx status codes
- `Status() int` - Returns the HTTP status code
- `StatusText() string` - Returns the status message
- `Headers() js.Value` - Returns the Headers object
- `Text(callback func(string))` - Reads response as text
- `JSON(callback func(js.Value))` - Parses response as JSON
- `Blob(callback func(js.Value))` - Reads response as Blob
- `ArrayBuffer(callback func(js.Value))` - Reads response as ArrayBuffer
- `Value() js.Value` - Returns the underlying JavaScript Response

## FetchPromise Methods

The `FetchPromise` type allows chaining:

- `Then(callback func(FetchResponse)) *FetchPromise` - Handle success
- `Catch(callback func(js.Value)) *FetchPromise` - Handle errors
- `Finally(callback func()) *FetchPromise` - Always runs after completion
- `Await() (FetchResponse, error)` - Synchronously wait for completion (blocks goroutine)

## Advanced Examples

### Error Handling

```go
js.Fetch("https://api.example.com/data").
    Then(func(response js.FetchResponse) {
        if !response.Ok() {
            println("HTTP Error:", response.Status(), response.StatusText())
            return
        }
        
        response.JSON(func(data js.Value) {
            // Process data
        })
    }).
    Catch(func(err js.Value) {
        println("Network error:", err.Get("message").String())
    })
```

### Multiple Headers

```go
opts := js.NewFetchOptions().
    Method("GET").
    Header("Accept", "application/json").
    Header("X-API-Key", "your-key").
    Header("X-Request-ID", "req-123")
```

### CORS Requests

```go
opts := js.NewFetchOptions().
    Mode("cors").
    Credentials("include").
    Header("Content-Type", "application/json")

js.Fetch("https://external-api.com/data", opts.Value())
```

### DELETE Request

```go
opts := js.NewFetchOptions().
    Method("DELETE").
    Header("Authorization", "Bearer token")

js.Fetch("https://api.example.com/resource/123", opts.Value()).
    Then(func(response js.FetchResponse) {
        if response.Ok() {
            println("Resource deleted")
        }
    })
```

## Notes

- All callbacks run asynchronously
- Function references are not automatically released (see comments in source)
- For production use, consider implementing proper lifecycle management for `js.FuncOf` callbacks
- The `Await()` method blocks the current goroutine - use with caution

## See Also

- [MDN Fetch API Documentation](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API)
- [JavaScript Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)
