# Set V to 1 for verbose output from the Makefile
Q=$(if $V,,@)
SRC=$(shell find . -type f -name '*.go')

# protoc-gen-go constraints
GEN_GO_BIN ?= protoc-gen-go
GEN_GO_MIN_VERSION ?= 1.36.1
GEN_GO_VERSION ?= $(shell $(GEN_GO_BIN) --version | awk -F ' v' '{print $$NF}')

# protoc-gen-go-grpc constraints
GEN_GRPC_BIN ?= protoc-gen-go-grpc
GEN_GRPC_MIN_VERSION ?= 1.5.1
GEN_GRPC_VERSION ?= $(shell $(GEN_GRPC_BIN) --version | awk -F ' ' '{print $$NF}')

all: lint generate test

ci: test

.PHONY: all ci

#########################################
# Build
#########################################

build: ;

#########################################
# Bootstrapping
#########################################

bootstra%:
	$Q curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.49.0
	$Q go install golang.org/x/vuln/cmd/govulncheck@latest
	$Q go install gotest.tools/gotestsum@v1.8.1
	$Q go install -mod=readonly google.golang.org/protobuf/cmd/protoc-gen-go
	$Q go install -mod=readonly google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: bootstrap

#########################################
# Test
#########################################

test:
	$Q $(GOFLAGS) gotestsum -- -coverpkg=./... -coverprofile=coverage.out -covermode=atomic ./...

race:
	$Q $(GOFLAGS) gotestsum -- -race ./...

.PHONY: test race

#########################################
# Linting
#########################################

fmt:
	$Q goimports -local github.com/golangci/golangci-lint -l -w $(SRC)

lint: SHELL:=/bin/bash
lint:
	$Q LOG_LEVEL=error golangci-lint run --config <(curl -s https://raw.githubusercontent.com/smallstep/workflows/master/.golangci.yml) --timeout=30m
	$Q govulncheck ./...

.PHONY: fmt lint

#########################################
# Generate
#########################################

generate: check-gen-go-version check-gen-grpc-version
	@# remove any previously generated protobufs & gRPC files
	@find . \
		-type f \
		-name "*.pb.go" \
		-delete

	@# generate the corresponding protobufs & gRPC code files
	$Q protoc \
		--proto_path=spec \
		--go_opt=module=github.com/smallstep/linkedca \
		--go_out=. \
		--go-grpc_opt=module=github.com/smallstep/linkedca \
		--go-grpc_out=. \
		$(shell find spec -type f -name "*.proto")	

.PHONY: generate

#########################################
# Tool constraints
#########################################

check-gen-go-version:
	@if ! printf "%s\n%s" "$(GEN_GO_MIN_VERSION)" "$(GEN_GO_VERSION)" | sort -V -C; then \
		echo "Your $(GEN_GO_BIN) version (v$(GEN_GO_VERSION)) is older than the minimum required (v$(GEN_GO_MIN_VERSION))."; \
		exit 1; \
	fi

.PHONY: check-gen-go-version

check-gen-grpc-version:
	@if ! printf "%s\n%s" "$(GEN_GRPC_MIN_VERSION)" "$(GEN_GRPC_VERSION)" | sort -V -C; then \
		echo "Your $(GEN_GRPC_BIN) version (v$(GEN_GRPC_VERSION)) is older than the minimum required (v$(GEN_GRPC_MIN_VERSION))."; \
		exit 1; \
	fi

.PHONY: check-gen-grpc-version
