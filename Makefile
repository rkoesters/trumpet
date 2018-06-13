GO         = go
BUILDFLAGS =
DEPS       = $(shell tools/list-deps.sh ./...)
VERSION    = $(shell git describe --always --dirty)

EXECNAME = trumpet

all: build

build:
	$(GO) build $(BUILDFLAGS) -ldflags="-X main.version=$(VERSION)" -o $(EXECNAME) ./cmd/trumpet

clean:
	$(GO) clean ./...
	rm -f $(EXECNAME)

config:
	tools/make-config.sh

deps:
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

test:
	$(GO) test -cover -race ./...

.PHONY: all build clean config deps test
