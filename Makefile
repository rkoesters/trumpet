GO         = go
BUILDFLAGS =
LDFLAGS    = -ldflags="-X main.version=$(VERSION)"
TESTFLAGS  = -cover

EXECNAME = trumpet
DEPS     = $(shell tools/list-deps.sh ./...)
VERSION  = $(shell git describe --always --dirty)

all: build

deps:
	$(GO) get -u $(BUILDFLAGS) golang.org/x/lint/golint
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

build:
	$(GO) build -o $(EXECNAME) $(BUILDFLAGS) $(LDFLAGS) ./cmd/trumpet

install:
	$(GO) install $(BUILDFLAGS) $(LDFLAGS) ./cmd/trumpet

check:
	$(GO) fmt ./...
	golint -set_exit_status ./...

test:
	$(GO) test $(TESTFLAGS) ./...

config:
	tools/make-config.sh

clean:
	$(GO) clean ./...
	rm -f $(EXECNAME)

.PHONY: all build check clean config deps install test
