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
			bs.WithMargin(bs.TOP|bs.BOTTOM, 5),
		).Append(
			bs.Heading(1, bs.WithTextAlign(bs.CENTER)).Append("Hello, World!"),
		),
	)

	// Run the application
	select {}
}
