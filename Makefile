GO         = go
BUILDFLAGS =
LDFLAGS    = -ldflags="-X main.version=$(VERSION)"
TESTFLAGS  = -cover

CMDS    = $(shell ls cmd)
SOURCES = $(shell find . -type f -name '*.go')
DEPS    = $(shell tools/list-deps.sh ./...)
VERSION = $(shell git describe --always --dirty)

all: $(CMDS)

$(CMDS): Makefile $(SOURCES)
	$(GO) build -o $@ $(BUILDFLAGS) $(LDFLAGS) ./cmd/$@

deps:
	$(GO) get -u $(BUILDFLAGS) golang.org/x/lint/golint
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

check:
	$(GO) fmt ./...
	golint -set_exit_status ./...

test:
	$(GO) test $(TESTFLAGS) ./...

install:
	$(GO) install $(BUILDFLAGS) $(LDFLAGS) ./cmd/trumpet

config:
	tools/make-config.sh

clean:
	$(GO) clean ./...
	rm -f $(CMDS)

.PHONY: all check clean config deps install test
