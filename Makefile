# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

TARGETOS ?= darwin
TARGETARCH ?= arm64
BUILD_VERSION ?= $(shell git describe --always --dirty)
BINARY_SUFFIX ?= ""

.PHONY: all
all: build

##@ Build
.PHONY: clean
clean: ## Delete all built binaries.
	rm -rf ./bin ./vendor ./out

.PHONY: build
build: ## Build the command binaries.
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build \
        -o bin/hathora-${TARGETOS}-${TARGETARCH}${BINARY_SUFFIX} \
        -ldflags "-X 'github.com/hathora/ci/internal/commands.BuildVersion=${BUILD_VERSION}'" \
        hathora/main.go

.PHONY: lint
lint: ## Lints the project, logging any warnings or errors without modifying any files.
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63
	golangci-lint run ./...

.PHONY: fmt
fmt: ## Reformat all code with the go fmt command.
	go fmt ./...

.PHONY: vet
vet: ## Run vet on all code with the go vet command.
	go vet ./...

##@ Tests
.PHONY: test
test: ## Unit test all modules.
	go test -v -parallel=1 -race ./...

.PHONY: test-short
test-short: ## Unit test all modules in short mode.
	go test -v -parallel=1 -race -short ./...

##@ Misc.
# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php
.PHONY: help
help: ## Display usage help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
