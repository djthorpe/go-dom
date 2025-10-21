
# Go parameters
GO=go
GOFLAGS = -ldflags "-s -w" 
BUILDDIR = build
WASM = $(wildcard wasm/*)
GOROOT = $(shell go env GOROOT)

# All targets
all: wasmbuild $(WASM)

# Rules for building
.PHONY: $(WASM)
$(WASM): mkdir
	@echo -n 'Building '
	@$(BUILDDIR)/wasmbuild compile -o ${BUILDDIR}/$(shell basename $@).wasm ./$@

.PHONY: wasmbuild
wasmbuild: mkdir
	@echo 'Building wasmbuild'
	@$(GO) build -o $(BUILDDIR)/wasmbuild ${GOFLAGS} ./cmd/wasmbuild

.PHONY: test
test:
	$(GO) test -v ./pkg/...

.PHONY: jstest
jstest: clean
	@$(GO) install github.com/agnivade/wasmbrowsertest@latest
	@GOOS=js GOARCH=wasm $(GO) test -v -exec="wasmbrowsertest" ./pkg/dom
	@GOOS=js GOARCH=wasm $(GO) test -v -exec="wasmbrowsertest" ./pkg/bootstrap

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
