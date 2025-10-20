# go-wasmbuild

**wasmbuild** is a command-line tool for compiling applications in the Go programming language ("golang") into 
[WASM](https://webassembly.org/) so they can be run in a web browser. 

It provides a development environment so you can test your code as you
develop it. There are also various packages which implement bridges to the web browser runtime, and several
popular JavaScript libtaties, so you can - for example - create reactive web components in your golang
application.

> *Experimental!* Please note this repository contains code which depends on experimental features of the Go language,
> and is also a very rough and ready implementation itself.

## Overview

This repository provides three main components:

1. **WASM Build Server** (`cmd/wasmbuild`) - A feature-rich development server for building and serving golang WASM
   applications with live reload, error notifications, and automatic dependency watching
2. **DOM Package** (`pkg/dom`) - A Go implementation of the Document Object Model (DOM) API that works both natively and in WASM environments
3. **Bootstrap 5 Package** (`pkg/bs5`) - Experimental Go wrappers for Bootstrap 5 components (very rough and ready)

There are also some examples of developing front-end applications in the `cmd/wasm` folder.

## WASM Build Server

The `wasmbuild` tool provides a modern development experience for WASM applications written in Go:

### Features

- **Compile golang to WASM** - Seamlessly compiles Go applications to WebAssembly
- **Serve WASM Applications** - Hosts your WASM apps with a built-in HTTP server, including serving the bootstrap HTML and JS needed to run Go WASM applications
- **Bootstrap 5 Support** - Optional `--bs-5` flag includes Bootstrap 5 CSS, JavaScript, and Icons in the HTML template
- **Automatic Dependency Discovery** - Watches all local package dependencies (no manual configuration needed)
- **Server-Sent Events (SSE)** - Efficient communication between server and browser
- **Live Reload** - Automatically recompiles and reloads the browser when source files change
- **Real-time Error Display** - Compilation errors appear directly in the browser with full error messages

### Quick Start

Build the wasmbuild to the `build/` directory:

```bash
make wasmbuild
```

Serve your WASM application with live reload:

```bash
./build/wasmbuild serve --watch ./cmd/wasm/helloworld
```

Or compile without serving, which also copies the required bootstrapping assets into the build folder:

```bash
./build/wasmbuild compile ./cmd/wasm/helloworld
```

### Usage

```bash
# Serve with live reload (discovers and watches local dependencies automatically)
./build/wasmbuild serve --watch <path-to-wasm-app>

# Serve with Bootstrap 5 support (includes Bootstrap CSS/JS in the HTML)
./build/wasmbuild serve --watch --bs-5 <path-to-wasm-app>

# Serve with verbose logging
./build/wasmbuild serve --verbose --watch <path-to-wasm-app>

# Compile to a specific output directory
./build/wasmbuild compile --output ./build <path-to-wasm-app>

# Custom listen address
./build/wasmbuild serve --listen 0.0.0.0:8080 <path-to-wasm-app>
```

The server automatically:

- Discovers all local package dependencies using `go list`
- Watches the main application and all dependent packages for changes
- Recompiles when any `.go` file changes
- Sends compilation errors to the browser in real-time
- Triggers browser reload on successful compilation

## Bootstrap 5 Package (Experimental)

> **Note:** The Bootstrap 5 wrapper (`pkg/bs5`) is very rough and ready. It provides
> Go bindings for Bootstrap 5 components but is not production-ready and should be
> considered experimental.

The `pkg/bs5` package provides Go wrappers for Bootstrap 5 components, allowing you to build
Bootstrap-based web applications in Go/WASM. A comprehensive demo application showcasing all
components is available in `cmd/wasm/bootstrap-app`.

### Bootstrap Components

- Card-based layouts with headers and footers
- Navigation bars with dropdowns and color schemes
- Alerts, badges, buttons with icon support
- Modals with forms and input groups
- Accordions, pagination, progress bars
- Toast notifications, offcanvas panels
- Tables, tabs, breadcrumbs
- Responsive grid system (Row/Col)
- Form controls with validation
- Bootstrap Icons integration

### Bootstrap Quick Start

Build and run the Bootstrap demo, which demonstrates all the components which can be created using the `pkg/bs5` package:

```bash
./build/wasmbuild serve --watch ./cmd/wasm/bootstrap-app --bs-5
```

See [`cmd/wasm/bootstrap-app/README.md`](cmd/wasm/bootstrap-app/README.md) for detailed documentation.

## DOM Package

Implements the document object model (DOM) for Go - which is
runnable both in a native go environment or via a web browser
when compiling for WASM target.

Presently go version 1.24 has been tested. [Tinygo](https://tinygo.org/) should
eventually be supported in order to facilitate smaller binary sizes.

### Hello, World

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

See the [WebAssembly documentation](https://go.dev/wiki/WebAssembly) for running
WASM within a web browser. There is some helper code which allows you to run it
from the command line:

```bash
git clone git@github.com:djthorpe/go-dom.git
cd go-dom
make wasmbuild
./build/wasmbuild serve cmd/wasm/helloworld
```

Then you can simply view the page at <http://localhost:9090/> and check the console log to see the output.

## Running Tests

(TODO)

You should run tests both in the native go environment and in the WASM
environment. For the latter, Google Chrome needs to be installed (other browsers
may work but have not been tested). The commands for testing are:

```bash
git clone git@github.com:djthorpe/go-dom.git
cd go-dom
make test && make jstest
```

Testing in a WASM environment uses [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest). Please see the documentation for that package for more information on testing in the WASM environment and browser support.
