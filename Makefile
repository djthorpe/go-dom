
# Go parameters
GO=go
GOFLAGS = -ldflags "-s -w" 
BUILDDIR = build
WASM = $(wildcard cmd/wasm/*)
GOROOT = $(shell go env GOROOT)

# All targets
all: wasmserver $(WASM)
	@cp ${GOROOT}/misc/wasm/wasm_exec.js ${BUILDDIR}

# Rules for building
.PHONY: $(WASM)
$(WASM): mkdir
	@echo "Building $(BUILDDIR)/$(shell basename $@).html"
	@GOOS=js GOARCH=wasm $(GO) build -o ${BUILDDIR}/$@.wasm -tags js ${GOFLAGS} ./$@

.PHONY: wasmserver
wasmserver: mkdir
	$(GO) build -o $(BUILDDIR)/wasmserver ${GOFLAGS} ./cmd/wasmserver

.PHONY: test
test:
	$(GO) test -v ./pkg/...

.PHONY: jstest
jstest: clean
	@if [ ! -x "${GOPATH}/bin/wasmbrowsertest" ]; then \
		echo "wasmbrowsertest not found, installing..."; \
		$(GO) install github.com/agnivade/wasmbrowsertest@latest; \
	fi
	@GOOS=js GOARCH=wasm $(GO) test -v -tags js -exec="wasmbrowsertest" ./pkg/dom

.PHONY: mkdir
mkdir:
	@install -d $(BUILDDIR)

.PHONY: tidy
tidy: 
	$(GO) mod tidy

.PHONY: clean
clean: tidy
	@rm -fr $(BUILDDIR)
	$(GO) clean
