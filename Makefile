# Change shell to bash
SHELL := /bin/bash

# Define Paths
export REPO_ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

include $(REPO_ROOT)/makefiles/help.mk

.PHONY: test
test: ## Run go test on all packages
	go test $(REPO_ROOT)/...

.PHONY: lint
lint: ## Run golangci-lit
	docker run \
		--rm \
		--userns=host \
		--volume $(REPO_ROOT):/app \
		--workdir /app \
		--interactive \
		--tty golangci/golangci-lint:latest \
			golangci-lint run \
			--verbose

.PHONY: test-verbose
test-verbose: ## Run go test on all packages with -v
	go test -v -count=1 $(REPO_ROOT)/...
