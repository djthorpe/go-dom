# go-dom

> *Experimental!* Please note this repository contains code which depends
on experimental features of the Go language.

Implements the document object model (DOM) for Go - which is
runnable both in a native go environment or via a web browser
when compiling for WASM target.

Presently go version 1.17 has been tested. [Tinygo](https://tinygo.org/) should
eventually be supported in order to facilitate smaller binary sizes.

## Running Tests

You should run tests both in the native go environment and in the WASM
environment. For the latter, Google Chrome needs to be installed (other browsers
may work but have not been tested). The commands for testing are:

```bash
git clone git@github.com:djthorpe/go-dom.git
cd go-dom
make test && make jstest
```

## Example: Hello, World 

There is a simple "Hello, World" example which runs both on the command line
and in the web browser. In order to run it in the web browser, you need to
serve the binary file using a HTTP server, included:

```bash
git clone git@github.com:djthorpe/go-dom.git
cd go-dom
make httpserver cmd/wasm/helloworld
cd build && ./httpserver -port :9090 
open http://localhost:9090/helloworld.html
```

