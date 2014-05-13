GO ?= go
GOLINT ?= golint
GOPATH := $(GOPATH)

#all: build
default: test

build: clean
	$(GO) build

dist: clean
	mkdir -p pkg pkg/linux-amd64 pkg/darwin-amd64
	GOOS=darwin GOARCH=amd64 $(GO) build -o pkg/darwin-amd64/es
	GOOS=linux GOARCH=amd64 $(GO) build -o pkg/linux-amd64/es

test: build fmt lint
	$(GO) test -v

lint:
	$(GOLINT) .

fmt:
	$(GO) fmt

clean:
	rm -rf pkg
	$(GO) clean

release:
	GOOS=linux GOARCH=amd64
