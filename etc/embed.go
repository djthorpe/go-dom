package etc

import (
	_ "embed"
)

//go:embed wasm_exec.html
var WasmExecHTML []byte

//go:embed notify.html
var NotifyHTML []byte

//go:embed favicon.png
var FaviconPNG []byte
