VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
MAKEFLAGS += --silent
PID := /tmp/.$(PROJECTNAME).pid
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
	@GOBIN=$(GOBIN) go generate ./...
	echo "done"

.PHONY: compile
compile: install build

.PHONY: start-server
start-server: stop-server
	echo "starting server... \c"
	@$(GOBIN)/$(PROJECTNAME) 2>&1 & echo $$! > $(PID)
	echo "done"
	cat $(PID) | sed "/^/s/^/  \>  PID: /"

.PHONY: stop-server
stop-server:
	echo "stopping server... \c"
	@touch $(PID)
	@kill `cat $(PID)` 2> /dev/null || true
	@rm $(PID)
	echo "done"

.PHONY: start
start: compile start-server

.PHONY: stop
stop: stop-server

.PHONY: test
test:
	@go test -short ${ALLFILES}

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

.PHONY: update
update:
	@clear
	make stop
	git pull origin dev
	make start

.PHONY: help
help:
	echo "Choose a command run in $(PROJECTNAME):"
	echo " install - installs all dependencies for the project"
	echo " build - builds a binary"
	echo " compile - installs all dependencies and builds a binary"
	echo " clean - cleans the cache and cleans up the build files"
