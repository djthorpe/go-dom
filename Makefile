
# Go parameters
GO=go
GOFLAGS = -ldflags "-s -w" 
BUILDDIR = build
WASM = $(wildcard cmd/wasm/*)
GOROOT = $(shell go env GOROOT)

# All targets
all: wasmbuild $(WASM)
	@cp ${GOROOT}/misc/wasm/wasm_exec.js ${BUILDDIR}

# Rules for building
.PHONY: $(WASM)
$(WASM): mkdir
	@echo "Building $(BUILDDIR)/$(shell basename $@).html"
	@GOOS=js GOARCH=wasm $(GO) build -o ${BUILDDIR}/$@.wasm -tags js ${GOFLAGS} ./$@

.PHONY: wasmbuild
wasmbuild: mkdir
	$(GO) build -o $(BUILDDIR)/wasmbuild ${GOFLAGS} ./cmd/wasmbuild

.PHONY: test
test:
	$(GO) test -v ./pkg/...

.PHONY: jstest
jstest: clean
	$(GO) install github.com/agnivade/wasmbrowsertest@latest
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
