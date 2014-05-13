GO ?= go
GOLINT ?= golint
GOPATH := $(GOPATH)

#all: build
default: test

build: clean
	$(GO) build

test: build fmt lint
	$(GO) test -v

lint:
	$(GOLINT) .

fmt:
	$(GO) fmt

clean:
	$(GO) clean
