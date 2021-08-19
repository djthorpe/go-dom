# go-dom

> *Experimental!* Please note this repository contains code which depends
on experimental features of the Go language.

Implements the document object model (DOM) for Go - which is
runnable both in a native go environment or via a web browser
when compiling for WASM target.

Presently go version 1.17 has been tested. [Tinygo](https://tinygo.org/) should
eventually be supported in order to facilitate smaller binary sizes.

## Hello, World

In order to access the current HTML document, access the `window` object. You
can dump the current document contents from the root element, or set the
title of the document:

```go
package main

import (
  . "github.com/djthorpe/go-dom/pkg/dom"
)

func main() {
  window := GetWindow()
  window.Document().SetTitle("Hello, World!")
  fmt.Println(window.Document().DocumentElement().OuterHTML())
}
```

To compile and run this example natively, run:

```bash
git clone git@github.com:djthorpe/go-dom.git
cd go-dom
go run ./cmd/wasm/helloworld
```

To run this within a web browser, you'll need to compile it to WASM instead:

```bash
git clone git@github.com:djthorpe/go-dom.git
cd go-dom
GOOS=js GOARCH=wasm go build -o build/helloworld.wasm ./cmd/wasm/helloworld
```

See the instructions [here](https://github.com/golang/go/wiki/WebAssembly) for running
WASM within a web browser. There is some helper code which allows you to run it
from the command line:

```bash
git clone git@github.com:djthorpe/go-dom.git
cd go-dom
make httpserver cmd/wasm/helloworld
cd build && ./httpserver -port :9090 
```

Then you can simply view the page at [`http://localhost:9090/helloworld.html`](http://localhost:9090/helloworld.html) and check the console log to see the output.

## Running Tests

You should run tests both in the native go environment and in the WASM
environment. For the latter, Google Chrome needs to be installed (other browsers
may work but have not been tested). The commands for testing are:

```bash
git clone git@github.com:djthorpe/go-dom.git
cd go-dom
make test && make jstest
```

Testing in a WASM environment uses [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest). Please see the documentation for that package for more information
on testing in the WASM environment and browser support.
