GO ?= go
GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)
MODULE_NAME ?= $(shell head -n1 go.mod | cut -f 2 -d ' ')
PARALLELS ?= 10

.PHONY: test
test:
	-$(GO) test ./... -parallel $(PARALLELS) -coverprofile=cover.out
	$(GO) tool cover -html=cover.out -o cover.html
	rm cover.out

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o backbar cmd/registry/main.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o hackbar cmd/tui/main.go
