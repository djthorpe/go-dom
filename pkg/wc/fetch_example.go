//go:build js

package wc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	dom "github.com/djthorpe/go-dom"
)

// ExampleFetcher demonstrates how to use the Fetcher component

// Example 1: Single fetch (manual mode)
func ExampleSingleFetch() {
	// Create a fetcher with no interval (manual mode only)
	fetcher := NewFetcher("https://api.example.com", 0)

	// Perform a single fetch
	fetcher.Fetch("/data", func(resp *FetchResponse) {
		if resp.IsSuccess() {
			fmt.Println("Data received:", resp.String())
		} else if resp.Error != nil {
			fmt.Println("Error:", resp.Error)
		} else {
			fmt.Printf("HTTP %d: %s\n", resp.Status, resp.StatusText)
		}
	})
}

// Example 2: Periodic fetch (automatic mode)
func ExamplePeriodicFetch() {
	// Create a fetcher that polls every 5 seconds
	fetcher := NewFetcher("https://api.example.com", 5*time.Second)

	// Start periodic fetching
	err := fetcher.Start("/status", func(resp *FetchResponse) {
		if resp.IsSuccess() {
			fmt.Println("Status update:", resp.String())
		} else if resp.Error != nil {
			fmt.Println("Fetch error:", resp.Error)
		}
	})

	if err != nil {
		fmt.Println("Failed to start fetcher:", err)
	}

	// Later... stop the fetcher
	// fetcher.Stop()
}

// Example 3: Fetch and parse JSON
func ExampleFetchJSON() {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	fetcher := NewFetcher("https://api.example.com", 0)

	fetcher.Fetch("/user/123", func(resp *FetchResponse) {
		var user User
		if err := resp.ParseJSON(&user); err != nil {
			fmt.Println("Failed to parse user:", err)
			return
		}
		fmt.Printf("User: %s (ID: %d)\n", user.Name, user.ID)
	})
}

// Example 4: Fetch and update DOM element
func ExampleFetchAndUpdateDOM(element dom.Element) {
	fetcher := NewFetcher("https://api.example.com", 10*time.Second)

	// Start periodic fetching and update the element
	fetcher.Start("/message", func(resp *FetchResponse) {
		if err := resp.UpdateElement(element); err != nil {
			fmt.Println("Failed to update element:", err)
		}
	})
}

// Example 5: Change interval on the fly
func ExampleDynamicInterval() {
	fetcher := NewFetcher("https://api.example.com", 5*time.Second)

	fetcher.Start("/data", func(resp *FetchResponse) {
		if resp.IsSuccess() {
			fmt.Println("Data:", resp.String())
		}
	})

	// Later, speed up the polling
	time.Sleep(30 * time.Second)
	fetcher.Stop()
	fetcher.SetInterval(2 * time.Second)
	fetcher.Start("/data", func(resp *FetchResponse) {
		if resp.IsSuccess() {
			fmt.Println("Faster data:", resp.String())
		}
	})
}

// Example 6: Using fetch options for CORS requests
func ExampleFetchWithCORS() {
	fetcher := NewFetcher("https://external-api.example.com", 0)

	// Configure CORS mode and credentials
	fetcher.SetFetchMode("cors").SetFetchCredentials("include")

	fetcher.Fetch("/api/data", func(resp *FetchResponse) {
		if resp.IsSuccess() {
			fmt.Println("CORS data:", resp.String())
		} else if resp.Error != nil {
			fmt.Println("Error:", resp.Error)
		}
	})
}

// Example 7: POST request with JSON body
func ExamplePostJSON() {
	type CreateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type CreateUserResponse struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	fetcher := NewFetcher("https://api.example.com", 0)

	// Prepare JSON payload
	payload := CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	jsonData, _ := json.Marshal(payload)

	// Make POST request
	fetcher.FetchWithOptions("/users", &FetchOptions{
		Method: "POST",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: bytes.NewReader(jsonData),
		Mode: "cors",
	}, func(resp *FetchResponse) {
		if !resp.IsSuccess() {
			fmt.Printf("Failed: HTTP %d\n", resp.Status)
			return
		}

		var result CreateUserResponse
		if err := resp.ParseJSON(&result); err != nil {
			fmt.Println("Parse error:", err)
			return
		}

		fmt.Printf("Created user #%d: %s\n", result.ID, result.Message)
	})
}

// Example 8: Custom headers and authentication
func ExampleWithAuthentication() {
	fetcher := NewFetcher("https://api.example.com", 10*time.Second)

	// Set authentication header for all requests
	fetcher.SetHeader("Authorization", "Bearer your-api-token-here")
	fetcher.SetHeader("X-API-Key", "your-api-key")

	// Start periodic authenticated fetching
	fetcher.Start("/protected/data", func(resp *FetchResponse) {
		if resp.Status == 401 {
			fmt.Println("Unauthorized - token may have expired")
			return
		}

		if resp.IsSuccess() {
			fmt.Println("Protected data:", resp.String())
		}
	})
}
