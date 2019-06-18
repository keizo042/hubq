VERSION=0.0.1
PRJROOT=$(pwd)

.PHONY: build
build:
	@CGO_ENABLED=0 go build -o bin/hubq -ldflags "-X main.Version=$(VERSION)"  ./cmd/hubq


.PHONY: release
release:
	@CGO_ENABLED=0 go build -o bin/hubq -dflags "-X main.Version=$(VERSION)" $(PRJROOT)/cmd/hubq
