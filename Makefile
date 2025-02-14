GO ?= go
GOBIN ?= $(shell echo `$(GO) env GOPATH`/bin)
GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)
MODULE_NAME ?= $(shell head -n1 go.mod | cut -f 2 -d ' ')

.PHONY: migrate
DIR ?= ./.data
CLEAR ?= false
migrate:
ifeq ($(CLEAR),true)
	$(GO) run test/data/migrate/main.go --dir $(DIR) --clear
else
	$(GO) run test/data/migrate/main.go --dir $(DIR)
endif

.PHONY: lint
lint:
	-$(GO) fmt ./...
	-$(GO) tool staticcheck ./...

.PHONY: test
PARALLELS ?= 10
test: lint
	-$(GO) test ./... -parallel $(PARALLELS) -coverprofile=cover.out
	-$(GO) tool cover -html=cover.out -o cover.html

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o backbar cmd/registry/main.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o hackbar cmd/tui/main.go
