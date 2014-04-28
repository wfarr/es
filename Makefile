GO ?= go
GOPATH := $(GOPATH)

#all: build
default: test

build: clean
	$(GO) build

test: build
	$(GO) test

fmt:
	$(GO) fmt

clean:
	$(GO) clean
