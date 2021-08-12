
# Go parameters
GO=go
GOFLAGS = -ldflags "-s -w $(GOLDFLAGS)" 
BUILDDIR = build
TAGS = w
WASM = $(wildcard cmd/wasm/*)
GOROOT = $(shell go env GOROOT)

# All targets
all: test mkdir $(WASM)
	@cp ${GOROOT}/misc/wasm/wasm_exec.js ${BUILDDIR}

# Rules for building
.PHONY: $(WASM)
$(WASM): 
	@echo "Building $(BUILDDIR)/$(shell basename $@).html"
	@GOOS=js GOARCH=wasm $(GO) build -o ${BUILDDIR}/$@.wasm -tags "$(TAGS)" ${GOFLAGS} ./$@
	@sed 's|json.wasm|$@.wasm|' etc/wasm.html > ${BUILDDIR}/$(shell basename $@).html

.PHONY: test
test:
	$(GO) test -tags "$(TAGS)" ./pkg/...

.PHONY: mkdir
mkdir:
	@install -d $(BUILDDIR)

.PHONY: clean
clean: 
	@rm -fr $(BUILDDIR)
	$(GO) mod tidy
	$(GO) clean
