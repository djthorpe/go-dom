package main

import (
	"fmt"

	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func NavBar() mvc.View {
	navbar := bs.NavBar(bs.WithColor(bs.DANGER_SUBTLE)).Content(
		bs.NavItem(
			bs.Icon("info-circle-fill"),
			" About",
		),
		bs.NavItem(
			bs.Icon("phone-fill"),
			" Contact",
		),
	)

	fmt.Println(navbar)

	return navbar
}
