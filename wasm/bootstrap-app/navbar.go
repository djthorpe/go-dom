package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func NavBar() mvc.View {
	return bs.NavBar(bs.WithColor(bs.DANGER_SUBTLE)).Header(
		bs.Icon("house-door-fill"),
		" Home",
	).Content(
		bs.NavItem(
			bs.Icon("info-circle-fill"),
			" About",
		),
		bs.NavItem(
			bs.Icon("phone-fill"),
			" Contact",
		),
	)
}
