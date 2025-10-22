package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
)

func main() {
	// Create the bootstrap app
	bs.New().Append(
		bs.Container(
			bs.WithBreakpoint(bs.BreakpointLarge),
			bs.WithClass("mt-5", "mb-5"),
			bs.WithClass("border-top", "border-bottom", "border-left", "border-danger"),
		).Append(
			bs.Heading(1, bs.WithClass("text-primary")).Append("Hello, World"),
			"This is a simple Hello World application using the ",
			bs.Anchor("https://golang.org", bs.WithClass("text-primary")).Append(
				"Go programming language",
			),
		),
	)

	// Run the application
	select {}
}
