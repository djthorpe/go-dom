package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Badges() mvc.View {
	return bs.Container().Append(
		"Badges ",
		bs.Badge().Append("Hello, World!"),
		bs.Badge(bs.WithColor(bs.INFO), bs.WithMargin(bs.X, 1)).Append("Info"),
		bs.Badge(bs.WithColor(bs.LIGHT), bs.WithBorder(bs.All), bs.WithMargin(bs.X, 1)).Append("Light"),
		bs.Badge(bs.WithColor(bs.DANGER), bs.WithMargin(bs.X, 1)).Append("Danger"),
		bs.Badge(bs.WithColor(bs.DANGER), bs.WithMargin(bs.X, 1)).Append("Danger"),
	)
}
