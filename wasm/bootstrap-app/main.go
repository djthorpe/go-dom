package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
)

func main() {
	// Create the bootstrap app
	bs.New().Append(
		NavBar(),
		NavExamples(),
		NavBarExamples(),
		BadgeExamples(),
		ButtonExamples(),
		AlertExamples(),
		IconExamples(),
		LinkExamples(),
		CardExamples(),
		ImageExamples(),
	)

	// Run the application
	select {}
}
