# -*- Makefile -*-
# -------------------------------------

DIST_DIR = dist

DOCKER ?= docker
GO = go

.PHONY: all
all: build

.PHONY: build
build: ${DIST_DIR}
	$(GO) build -o $(DIST_DIR)/kubever .

${DIST_DIR}:
	mkdir -p $(DIST_DIR)

.PHONY: package
package:
	$(DOCKER) build -t kubever:latest .

.PHONY: clean
clean:
	rm -rf $(DIST_DIR)
