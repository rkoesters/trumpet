GO         = go
BUILDFLAGS = -ldflags="-X main.version=$(VERSION)"
DEPS       = $(shell tools/list-deps.sh ./...)
VERSION    = $(shell git describe --always --dirty)

EXECNAME = trumpet

all: build

build:
	$(GO) build $(BUILDFLAGS) -o $(EXECNAME) ./cmd/trumpet

install:
	$(GO) install $(BUILDFLAGS) ./cmd/trumpet

clean:
	$(GO) clean ./...
	rm -f $(EXECNAME)

config:
	tools/make-config.sh

deps:
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

test:
	$(GO) test -cover -race ./...

.PHONY: all build install clean config deps test
