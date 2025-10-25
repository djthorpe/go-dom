# go-wasmbuild

**wasmbuild** is a command-line tool for compiling applications from the Go programming language ("golang") into
[WASM](https://webassembly.org/) so they can be run as web applications in a browser.

It provides a development environment so you can test your code as you
develop it. There are packages for several
popular JavaScript libraries, so you can - for example - create reactive web components in golang.

> *Experimental!* Please note this repository contains code which depends on experimental features of the Go language.

This repository provides three main components:

1. **WASM Build Server** (`cmd/wasmbuild`) - A feature-rich development server for building and serving golang WASM
   applications with live reload, error notifications, and automatic dependency watching
2. **DOM Package** (`pkg/dom`) - A Go implementation of the Document Object Model (DOM) API that works both natively and in WASM environments
3. **Bootstrap 5 Package** (`pkg/bootstrap`) - Go wrappers for Bootstrap 5 components

There are also some examples of developing front-end applications in the `wasm` folder.

## WASM Build Server

The `wasmbuild` tool provides a modern development experience for WASM applications written in Go. Please see [the documentation for `wasmbuild`](cmd/wasmbuild/README.md) for full details.

Example usage:

```bash
# Install wasmbuild
go install github.com/djthorpe/go-wasmbuild/cmd/wasmbuild

# Start the wasmbuild server, watch for changes
wasmbuild serve -w path/to/your/go/project

# Navigate to http://localhost:8080 in your browser for development

# Build a production version of your app
wasmbuild build -o ./dist path/to/your/go/project
```

### Features

- **Compile golang to WASM** - Seamlessly compiles Go applications to WebAssembly
- **Serve WASM Applications** - Hosts your WASM apps with a built-in HTTP server, including serving the bootstrap HTML and JS needed to run Go WASM applications
- **Automatic Dependency Discovery** - Watches all local package dependencies (no manual configuration needed)
- **Live Reload** - Automatically recompiles and reloads the browser when source files change
- **Real-time Error Display** - Compilation errors appear directly in the browser with full error messages

## Bootstrap

The `pkg/bootstrap` package provides Go wrappers for Bootstrap 5 components, allowing you to build Bootstrap-based web applications in Go/WASM. A comprehensive demo application showcasing all components is available in `wasm/bootstrap-app`.

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

## Running Tests

You can run tests both in the native go environment and in the WASM
environment.  The commands for testing are:

```bash
git clone git@github.com:djthorpe/go-wasmbuild.git
cd go-wasmbuild
# Test in native go environment
make test
# Test in browser environment
make jstest
```

Testing in a WASM environment uses [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest). Please see the documentation for that package for more information on testing in the WASM environment and browser support.

## License

This project is licensed under the Apache License, Version 2.0. The Apache License is a permissive free software license that:

- **Allows** commercial use, modification, distribution, and private use
- **Requires** preservation of copyright and license notices, and inclusion of a NOTICE file if one exists
- **Provides** an express grant of patent rights from contributors to users
- **Does not** provide trademark rights or warranties
- **Requires** stating significant changes made to the software

For the complete license terms, see the [LICENSE](LICENSE) file or visit <http://www.apache.org/licenses/LICENSE-2.0>.
