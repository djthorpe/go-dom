package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Alerts() mvc.View {
	return bs.Container().Append(
		"Alerts ",
		bs.Alert().Append("Hello, World!"),
		bs.Alert(bs.WithColor(bs.INFO), bs.WithMargin(bs.X, 1)).Append("Info"),
		bs.Alert(bs.WithColor(bs.LIGHT), bs.WithBorder(bs.All), bs.WithMargin(bs.X, 1)).Append("Light"),
		bs.Alert(bs.WithColor(bs.DANGER), bs.WithMargin(bs.X, 1)).Append("Danger"),
		bs.Alert(bs.WithColor(bs.DANGER), bs.WithMargin(bs.X, 1)).Append("Danger"),
	)
}
