# Set V to 1 for verbose output from the Makefile
Q=$(if $V,,@)
SRC=$(shell find . -type f -name '*.go')

# Go tools
GOIMPORTS=golang.org/x/tools/cmd/goimports
GOLANGCI_LINT=github.com/golangci/golangci-lint/v2/cmd/golangci-lint
GOLANGCI_LINT_CONFIG_URL=https://raw.githubusercontent.com/smallstep/workflows/master/.golangci.yml
GOTESTSUM=gotest.tools/gotestsum
GOVULNCHECK=golang.org/x/vuln/cmd/govulncheck
BUF=github.com/bufbuild/buf/cmd/buf

all: generate lint test

ci: test

.PHONY: all ci

#########################################
# Build
#########################################

build: ;

#########################################
# Test
#########################################

test:
	$Q $(GOFLAGS) go tool $(GOTESTSUM) -- -coverpkg=./... -coverprofile=coverage.out -covermode=atomic ./...

race:
	$Q $(GOFLAGS) go tool $(GOTESTSUM) -- -race ./...

.PHONY: test race

#########################################
# Linting
#########################################

fmt:
	$Q go tool $(GOIMPORTS) -local github.com/smallstep/linkedca -l -w $(SRC)

lint: SHELL:=/bin/bash
lint:
	$Q LOG_LEVEL=error go tool $(GOLANGCI_LINT) run --config <(curl -s $(GOLANGCI_LINT_CONFIG_URL)) --timeout=30m ./...
	$Q go tool $(GOVULNCHECK) ./...

.PHONY: fmt lint

#########################################
# Generate
#########################################

generate:
	@# remove any previously generated protobufs & gRPC files
	@find . \
		-type f \
		-name "*.pb.go" \
		-delete

	@# generate the corresponding protobufs & gRPC code files
	$Q go tool $(BUF) generate -o .

.PHONY: generate

