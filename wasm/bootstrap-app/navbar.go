package main

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func NavBar() mvc.View {
	navbar := bs.NavBar(
		bs.WithColor(bs.LIGHT),
		bs.WithSize(bs.SizeMedium),
		bs.WithBorder(bs.Bottom, bs.PRIMARY),
		bs.WithMargin(bs.Bottom, 3),
	).Content(
		bs.NavItem(
			bs.Icon("info-circle-fill"),
			" About",
		),
		bs.NavItem(
			bs.Icon("phone-fill"),
			" Contact",
		),
	).(mvc.ViewWithCaption).Caption("My Application")

	navbar.AddEventListener("click", func(node Node) {
		view := mvc.ViewFromNode(node)
		if view != nil && view.Name() == bs.ViewNavItem {
			fmt.Println(" ===> ", view.Name())
		}
	})

	return navbar
}
