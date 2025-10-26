package main

import (
	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func main() {
	// Make a new application
	app := mvc.New("Bart App")

	// Create a heading
	app.Append(
		Alerts(),
		Containers(),
		Badges(),
		Rules(),
		Icons(),
	)

	// Run the application
	select {}
}
