package etc

import (
	_ "embed"
)

//go:embed wasm_exec.html
var WasmExecHTML []byte

//go:embed notify.js
var NotifyJS []byte

//go:embed notify.css
var NotifyCSS []byte

//go:embed favicon.png
var FaviconPNG []byte
