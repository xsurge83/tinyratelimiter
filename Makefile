
export GOBIN ?= $(shell pwd)/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Lint parameters
LINTCMD=golangci-lint
LINTCMD_ARGS=run --enable-all --deadline=5m

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

all: clean cover lint  

.PHONY: test
test:
	$(GOTEST)  -race -v

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

build:
	$(GOBUILD) -v ./...


.PHONY: lint
lint:
	$(LINTCMD) $(LINTCMD_ARGS) ./...

.PHONY: clean 
clean:
	rm -f cover.*