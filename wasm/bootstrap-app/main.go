package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
)

func main() {
	// Create the bootstrap app
	bs.New().Append(
		BadgeExamples(),
		ButtonExamples(),
		AlertExamples(),
	)

	// Run the application
	select {}
}
