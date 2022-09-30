APP = csvmodule
GOBASE = $(shell pwd)
GOBIN = $(GOBASE)/build/bin
LINT_PATH = $(GOBASE)/build/lint

.PHONY: lint install-golangci

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deps: ## Fetch required dependencies
	go mod tidy -compat=1.17
	go mod download

lint: ## Linter for developers
	$(LINT_PATH)/golangci-lint run --timeout=5m -c .golangci.yml

all: deps build lint test clean mac_build mac_docker

install-golangci: ## Install the correct version of lint
	GOBIN=$(LINT_PATH) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
