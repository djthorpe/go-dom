# BART Station Viewer

A WebAssembly application that demonstrates the Fetcher component by displaying real-time BART (Bay Area Rapid Transit) station data.

## Features

- **Real BART API Integration**: Fetches live station data from the official BART API
- **Periodic Auto-Refresh**: Automatically updates station data every 10 seconds
- **Beautiful UI**: Bootstrap 5 cards with station details including:
  - Station name and abbreviation
  - Address and location information
  - County information
- **Interactive Controls**:
  - Manual fetch button
  - Start/Stop auto-refresh
  - Real-time status indicators with timestamps
- **Error Handling**: Displays loading states and error messages

## Technical Implementation

### Fetcher Component Usage

This app demonstrates the `pkg/wc.Fetcher` component which provides:

```go
// Create fetcher with 10-second interval
fetcher := wc.NewFetcher("https://api.bart.gov", 10*time.Second)

// Perform a single fetch
fetcher.Fetch("/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y", callback)

// Start periodic fetching
fetcher.Start("/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y", callback)

// Stop periodic fetching
fetcher.Stop()
```

### Response Handling

The Fetcher uses `net/http` which is automatically converted to browser Fetch API calls:

```go
displayStations := func(resp *wc.FetchResponse) {
    // Check for errors
    if resp.Error != nil {
        // Handle error
        return
    }
    
    // Check HTTP status
    if !resp.IsSuccess() {
        // Handle HTTP error
        return
    }
    
    // Parse JSON response
    var data StationsResponse
    if err := resp.ParseJSON(&data); err != nil {
        // Handle parse error
        return
    }
    
    // Use the data
    stations := data.Root.Stations.Station
    // ... display stations
}
```

## Building

```bash
# Build the WASM binary
GOOS=js GOARCH=wasm go build -o build/bart-app.wasm ./cmd/wasm/bart-app

# Copy wasm_exec.js if needed
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" build/
```

## Running

Use the wasmserver to serve the application:

```bash
# From the project root
./build/wasmserver serve --watch ./cmd/wasm/bart-app

# Or serve directly
./build/wasmserver serve --html etc/bart.html ./cmd/wasm/bart-app
```

Then open your browser to `http://localhost:9090`

## API Information

This app uses the official BART API:

- **Endpoint**: `https://api.bart.gov/api/stn.aspx`
- **Command**: `stns` (station list)
- **Format**: JSON
- **Documentation**: <https://api.bart.gov/docs/overview/examples.aspx>

## Key Concepts Demonstrated

1. **HTTP Fetching in WASM**: Using Go's `net/http` package in WebAssembly
2. **JSON Parsing**: Unmarshaling API responses into Go structs
3. **DOM Manipulation**: Creating and updating HTML elements from Go
4. **Periodic Tasks**: Using time.Ticker for auto-refresh
5. **Event Handling**: Button clicks and user interactions
6. **Error Handling**: Graceful error display and recovery
7. **Bootstrap Integration**: Creating responsive UI with Bootstrap 5
8. **State Management**: Enabling/disabling buttons based on fetcher state

## File Structure

- `main.go` - Main application code with Fetcher usage
- `../../etc/bart.html` - HTML template with Bootstrap 5
- `../../pkg/wc/fetch.go` - Reusable Fetcher component
- `../../pkg/wc/fetch_example.go` - Additional Fetcher examples

## Future Enhancements

Potential improvements:

- Add real-time train departures for each station
- Filter stations by county or line
- Show station map with coordinates
- Add favorite stations feature
- Display station alerts and advisories
