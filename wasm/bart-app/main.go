package main

import (
	"fmt"
	"time"

	// Packages
	dom "github.com/djthorpe/go-dom"
	bs5 "github.com/djthorpe/go-wasmbuild/pkg/bs5"
)

func main() {
	// Create the Bootstrap app
	app := bs5.New("BART Station Viewer")

	// Create the main container
	container := app.Container()
	container.Element.AddClass("mt-5")

	// Add BART logo
	logoDiv := app.CreateElement("div")
	logoDiv.AddClass("text-center")
	logoDiv.AddClass("mb-4")

	logo := app.CreateElement("img")
	logo.SetAttribute("src", "https://upload.wikimedia.org/wikipedia/commons/2/26/Bart-logo.svg")
	logo.SetAttribute("alt", "BART Logo")
	logo.AddClass("img-fluid")
	logo.SetAttribute("style", "max-width: 200px;")
	logoDiv.AppendChild(logo)
	container.Element.AppendChild(logoDiv)

	// Add title
	title := app.H1(app.CreateTextNode("BART Station Viewer"))
	title.AddClass("mb-4")
	title.AddClass("text-center")
	container.Element.AppendChild(title.Element)

	// Add description
	desc := app.CreateElement("p")
	desc.AddClass("lead")
	desc.AddClass("mb-4")
	desc.AddClass("text-center")
	desc.AppendChild(app.CreateTextNode("Fetching BART station data using the Fetcher component"))
	container.Element.AppendChild(desc)

	// Create control buttons
	btnRow := app.CreateElement("div")
	btnRow.AddClass("mb-4")

	fetchBtn := app.Button(bs5.ColorPrimary, app.CreateTextNode("Fetch Stations"))
	fetchBtn.Element.AddClass("me-2")

	startBtn := app.Button(bs5.ColorSuccess, app.CreateTextNode("Start Auto-Refresh (10s)"))
	startBtn.Element.AddClass("me-2")

	stopBtn := app.Button(bs5.ColorDanger, app.CreateTextNode("Stop Auto-Refresh"))
	stopBtn.Element.AddClass("me-2")
	stopBtn.Element.SetAttribute("disabled", "true")

	statusSpan := app.CreateElement("span")
	statusSpan.AddClass("ms-3")
	statusSpan.AddClass("text-muted")
	statusSpan.AppendChild(app.CreateTextNode("Ready"))

	btnRow.AppendChild(fetchBtn.Element)
	btnRow.AppendChild(startBtn.Element)
	btnRow.AppendChild(stopBtn.Element)
	btnRow.AppendChild(statusSpan)
	container.Element.AppendChild(btnRow)

	// Create loading alert
	loadingSpinner := app.CreateElement("span")
	loadingSpinner.AddClass("spinner-border")
	loadingSpinner.AddClass("spinner-border-sm")
	loadingSpinner.AddClass("me-2")
	loadingAlert := app.Alert(bs5.ColorInfo, loadingSpinner, app.CreateTextNode("Fetching BART stations..."))
	loadingAlert.AddClass("d-none")
	container.Element.AppendChild(loadingAlert.Element)

	// Create error alert
	errorAlert := app.Alert(bs5.ColorDanger)
	errorAlert.AddClass("d-none")
	container.Element.AppendChild(errorAlert.Element)

	// Create station list container
	stationGrid := app.CreateElement("div")
	stationGrid.AddClass("row")
	stationGrid.AddClass("g-3")
	container.Element.AppendChild(stationGrid)

	// Add container to body
	app.Body().AppendChild(container.Element)

	// BART API response structures
	type Station struct {
		Name    string `json:"name"`
		Abbr    string `json:"abbr"`
		GTFSLat string `json:"gtfs_latitude"`
		GTFSLon string `json:"gtfs_longitude"`
		Address string `json:"address"`
		City    string `json:"city"`
		County  string `json:"county"`
		State   string `json:"state"`
		Zipcode string `json:"zipcode"`
	}

	type StationsResponse struct {
		Root struct {
			Stations struct {
				Station []Station `json:"station"`
			} `json:"stations"`
		} `json:"root"`
	}

	// Create fetcher
	fetcher := wc.NewFetcher("https://api.bart.gov", 10*time.Second)

	// Helper function to update status
	updateStatus := func(message, className string) {
		for statusSpan.HasChildNodes() {
			statusSpan.RemoveChild(statusSpan.FirstChild())
		}
		statusSpan.RemoveClass("text-muted")
		statusSpan.RemoveClass("text-success")
		statusSpan.RemoveClass("text-danger")
		statusSpan.AddClass(className)
		statusSpan.AppendChild(app.CreateTextNode(message))
	}

	// Fetch callback to display stations
	displayStations := func(resp *wc.FetchResponse) {
		// Hide loading indicator
		loadingAlert.RemoveClass("d-none")
		loadingAlert.AddClass("d-none")

		if resp.Error != nil {
			// Show error
			errorAlert.RemoveClass("d-none")
			for errorAlert.Element.HasChildNodes() {
				errorAlert.Element.RemoveChild(errorAlert.Element.FirstChild())
			}
			errorAlert.Element.AppendChild(app.CreateTextNode(fmt.Sprintf("Error: %v", resp.Error)))
			updateStatus("Error occurred", "text-danger")
			return
		}

		if !resp.IsSuccess() {
			// Show HTTP error
			errorAlert.RemoveClass("d-none")
			for errorAlert.Element.HasChildNodes() {
				errorAlert.Element.RemoveChild(errorAlert.Element.FirstChild())
			}
			errorAlert.Element.AppendChild(app.CreateTextNode(fmt.Sprintf("HTTP Error: %d %s", resp.Status, resp.StatusText)))
			updateStatus("HTTP error occurred", "text-danger")
			return
		}

		// Hide error
		errorAlert.AddClass("d-none")

		// Parse JSON response
		var data StationsResponse
		if err := resp.ParseJSON(&data); err != nil {
			errorAlert.RemoveClass("d-none")
			for errorAlert.Element.HasChildNodes() {
				errorAlert.Element.RemoveChild(errorAlert.Element.FirstChild())
			}
			errorAlert.Element.AppendChild(app.CreateTextNode(fmt.Sprintf("Parse error: %v", err)))
			updateStatus("Parse error occurred", "text-danger")
			return
		}

		// Clear existing stations
		for stationGrid.HasChildNodes() {
			stationGrid.RemoveChild(stationGrid.FirstChild())
		}

		// Display stations
		stations := data.Root.Stations.Station
		updateStatus(fmt.Sprintf("Loaded %d stations at %s", len(stations), time.Now().Format("15:04:05")), "text-success")

		for _, station := range stations {
			// Create column div
			colDiv := app.CreateElement("div")
			colDiv.AddClass("col-md-6")
			colDiv.AddClass("col-lg-4")

			// Create card using bs5 Card component
			card := app.Card()
			card.Element.AddClass("h-100")

			// Get or create card body
			cardBody := card.Body()

			// Station name
			stationTitle := app.H5(app.CreateTextNode(station.Name))
			stationTitle.AddClass("card-title")
			cardBody.Element.AppendChild(stationTitle.Element)

			// Station abbreviation badge
			abbrBadge := app.Badge(bs5.ColorPrimary, app.CreateTextNode(station.Abbr))
			abbrBadge.Element.AddClass("mb-2")
			cardBody.Element.AppendChild(abbrBadge.Element)

			// Address
			if station.Address != "" {
				addrP := app.CreateElement("p")
				addrP.AddClass("card-text")
				addrP.AddClass("small")
				addrP.AddClass("mb-1")

				addrIcon := app.Icon("geo-alt-fill")
				addrIcon.Element.AddClass("me-1")
				addrP.AppendChild(addrIcon.Element)
				addrP.AppendChild(app.CreateTextNode(station.Address))
				cardBody.Element.AppendChild(addrP)
			}

			// City
			if station.City != "" {
				cityP := app.CreateElement("p")
				cityP.AddClass("card-text")
				cityP.AddClass("small")
				cityP.AddClass("mb-1")
				cityP.AppendChild(app.CreateTextNode(fmt.Sprintf("%s, %s %s", station.City, station.State, station.Zipcode)))
				cardBody.Element.AppendChild(cityP)
			}

			// County
			if station.County != "" {
				countyP := app.CreateElement("p")
				countyP.AddClass("card-text")
				countyP.AddClass("small")
				countyP.AddClass("text-muted")
				countyP.AppendChild(app.CreateTextNode(fmt.Sprintf("County: %s", station.County)))
				cardBody.Element.AppendChild(countyP)
			}

			colDiv.AppendChild(card.Element)
			stationGrid.AppendChild(colDiv)
		}
	}

	// Fetch button handler
	fetchBtn.Element.AddEventListener("click", func(node dom.Node) {
		loadingAlert.RemoveClass("d-none")
		updateStatus("Fetching...", "text-primary")
		fetcher.Fetch("/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y", displayStations)
	})

	// Start button handler
	startBtn.Element.AddEventListener("click", func(node dom.Node) {
		loadingAlert.RemoveClass("d-none")
		updateStatus("Starting auto-refresh...", "text-primary")

		err := fetcher.Start("/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y", displayStations)
		if err != nil {
			updateStatus(fmt.Sprintf("Error: %v", err), "text-danger")
			loadingAlert.AddClass("d-none")
			return
		}

		// Disable start and fetch buttons
		startBtn.Element.SetAttribute("disabled", "true")
		fetchBtn.Element.SetAttribute("disabled", "true")

		// Enable stop button by accessing underlying JS value
		if jsVal, ok := stopBtn.Element.(interface {
			JSValue() interface {
				Call(string, ...interface{}) interface{}
			}
		}); ok {
			jsVal.JSValue().Call("removeAttribute", "disabled")
		}
	})

	// Stop button handler
	stopBtn.Element.AddEventListener("click", func(node dom.Node) {
		fetcher.Stop()
		updateStatus("Auto-refresh stopped", "text-muted")

		// Disable stop button
		stopBtn.Element.SetAttribute("disabled", "true")

		// Enable start and fetch buttons
		if jsVal, ok := startBtn.Element.(interface {
			JSValue() interface {
				Call(string, ...interface{}) interface{}
			}
		}); ok {
			jsVal.JSValue().Call("removeAttribute", "disabled")
		}
		if jsVal, ok := fetchBtn.Element.(interface {
			JSValue() interface {
				Call(string, ...interface{}) interface{}
			}
		}); ok {
			jsVal.JSValue().Call("removeAttribute", "disabled")
		}
	})

	// Perform initial fetch
	loadingAlert.RemoveClass("d-none")
	updateStatus("Loading initial data...", "text-primary")
	fetcher.Fetch("/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y", displayStations)

	// Keep the program running
	select {}
}
