# These shell flags are REQUIRED for an early exit in case any program called by make errors!
.SHELLFLAGS=-euo pipefail -c
SHELL := /bin/bash

.PHONY: all fmt clean check build tidy goimports golangci-lint

# Set the GOBIN environment variable so that dependencies will be installed
# always in the same place, regardless of the value of GOPATH
CACHE := $(PWD)/.cache
export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)

DOCKER ?= docker

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

tests:
	@go test ./test/...

db-setup: db-teardown
	@$(DOCKER) volume create finease-data || true
	$(DOCKER) run -d \
	--name finease-db \
	--publish 5432:5432 \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=postgres \
	-e POSTGRES_DB=postgres \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v finease-data:/var/lib/postgresql/data \
	--restart on-failure \
	postgres:13

db-bootstrap: db-setup build
	$(GOBIN)/finease-backend migrate --db-name=postgres --db-user=postgres --db-password=postgres --db-host=localhost --db-port=5432

db-teardown:
	@$(DOCKER) container stop finease-db || true
	@$(DOCKER) container rm finease-db || true
