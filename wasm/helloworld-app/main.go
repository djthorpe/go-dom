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
			bs.WithBorder(bs.BorderAll, bs.PRIMARY),
		).Append(
			bs.Heading(1, bs.WithClass("text-primary")).Append("Hello, World"),
		),
	)

	// Run the application
	select {}
}
