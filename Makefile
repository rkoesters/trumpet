GO         = go
BUILDFLAGS = -ldflags="-X main.version=$(VERSION)"
DEPS       = $(shell tools/list-deps.sh ./...)
VERSION    = $(shell git describe --always --dirty)

EXECNAME = trumpet

all: build

build:
	$(GO) build $(BUILDFLAGS) -o $(EXECNAME) ./cmd/trumpet

check:
	go fmt ./...
	golint -set_exit_status ./...

clean:
	$(GO) clean ./...
	rm -f $(EXECNAME)

config:
	tools/make-config.sh

deps:
	$(GO) get -u golang.org/x/lint/golint
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

install:
	$(GO) install $(BUILDFLAGS) ./cmd/trumpet

test:
	$(GO) test -cover -race ./...

.PHONY: all build check clean config deps install test
