VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
MAKEFLAGS += --silent
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(GOPATH)/src/github.com/n7down/iota/cmd/iota/*.go
ALLFILES=$(shell find . -name '*.go')

.PHONY: install
install:
	echo "installing... \c"
	@go get ./...
	echo "done"

.PHONY: build
build: clean
	echo "building... \c"
	@GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)
	echo "done"

.PHONY: generate
generate:
	echo "generating dependency files... \c"
	@protoc --go_out=plugins=grpc:internal/pb/settings internal/pb/settings/settings.proto
	@GOBIN=$(GOBIN) go generate ./...
	echo "done"

.PHONY: compile
compile: install build

.PHONY: test
test:
	go test -v ./...

.PHONY: vet
vet:
	@go vet ${ALLFILES}

.PHONY: lint
lint:
	@for file in ${ALLFILES); do \
		golint $$file ; \
		done

.PHONY: clean
clean:
	echo "cleaning build cache... \c"
	@go clean
	@rm -rf bin/
	echo "done"

.PHONY: docker-build
docker-build:
	echo "building with docker"
	docker-compose up -d

.PHONY: docker-logs-follow
docker-logs:
	docker-compose logs -f

.PHONY: help
help:
	echo "Choose a command run in $(PROJECTNAME):"
	echo " install - installs all dependencies for the project"
	echo " build - builds a binary"
	echo " compile - installs all dependencies and builds a binary"
	echo " clean - cleans the cache and cleans up the build files"
