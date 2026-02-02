# -*- Makefile -*-
# --------------------------------------------------------------------------
# Define variables

GIT_SUMMARY := $(shell git describe --tags --dirty --always 2>/dev/null || echo "v0.0.0-unknown")
BUILD_VERSION := $(shell git describe --tags --always 2>/dev/null || echo "v0.0.0-unknown")
BUILD_DATE := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

DIST_DIR ?= dist

DOCKER ?= podman

GO ?= go
GOFLAGS ?=
GO_LDFLAGS := -X 'main.gitSummary=$(GIT_SUMMARY)' -X 'main.buildDate=$(BUILD_DATE)'
GO_BUILD := $(GO) build $(GOFLAGS) -ldflags "$(GO_LDFLAGS)"
GO_TEST := $(GO) test $(GOFLAGS)

# --------------------------------------------------------------------------
# Targets

.PHONY: all
all: 
	@echo "NOOP"

.PHONY: preflight
preflight:
	# oapi-codegen v2 implements OpenAPI v3
	# go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	#
	# ^^ prefer go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest 
	# ^^ which will ensure the tool is managed by the go.mod file

.PHONY: generate
generate:
	${GO} generate ./...

.PHONY: build
build:
	$(GO_BUILD) -o $(DIST_DIR)/app .

.PHONY: clean
clean:
	${GO} clean ./...
	rm -rf $(DIST_DIR)
