
# Go parameters
GO=go
GOFLAGS = -ldflags "-s -w" 
BUILDDIR = build
WASM = $(wildcard cmd/wasm/*)
GOROOT = $(shell go env GOROOT)

# All targets
all: test httpserver $(WASM)
	@cp ${GOROOT}/misc/wasm/wasm_exec.js ${BUILDDIR}

# Rules for building
.PHONY: $(WASM)
$(WASM): mkdir
	@echo "Building $(BUILDDIR)/$(shell basename $@).html"
	@GOOS=js GOARCH=wasm $(GO) build -o ${BUILDDIR}/$@.wasm -tags js ${GOFLAGS} ./$@
	@sed 's|json.wasm|$@.wasm|' etc/wasm.html > ${BUILDDIR}/$(shell basename $@).html

.PHONY: httpserver
httpserver: mkdir
	$(GO) build -o $(BUILDDIR)/httpserver ${GOFLAGS} ./cmd/httpserver

.PHONY: test
test:
	$(GO) test -v ./pkg/...

.PHONY: jstest
jstest:
	$(GO) install github.com/agnivade/wasmbrowsertest
	@GOOS=js GOARCH=wasm $(GO) test -v -tags js -exec="${GOPATH}/bin/wasmbrowsertest" ./pkg/dom

.PHONY: mkdir
mkdir:
	@install -d $(BUILDDIR)

.PHONY: clean
clean: 
	@rm -fr $(BUILDDIR)
	$(GO) mod tidy
	$(GO) clean
