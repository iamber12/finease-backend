# These shell flags are REQUIRED for an early exit in case any program called by make errors!
.SHELLFLAGS=-euo pipefail -c
SHELL := /bin/bash

.PHONY: all fmt clean check build tidy goimports golangci-lint

# Set the GOBIN environment variable so that dependencies will be installed
# always in the same place, regardless of the value of GOPATH
CACHE := $(PWD)/.cache
export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)

all: build

build: tidy ## Build binaries
	@CGO_ENABLED=0 go build -a -o $(GOBIN)/finease-backend ./cmd/main.go

run: build
	@bin/stryds-server

tidy:
	@go mod tidy

verify:
	@go mod verify

fmt:
	@go fmt ./...
