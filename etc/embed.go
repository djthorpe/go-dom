package etc

import (
	_ "embed"
)

//go:embed bootstrap.html
var BootstrapHTML []byte

//go:embed bs5.html
var Bootstrap5 []byte

//go:embed notify.js
var NotifyJS []byte

//go:embed notify.css
var NotifyCSS []byte

//go:embed favicon.png
var FaviconPNG []byte
