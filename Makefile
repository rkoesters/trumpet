GO     = go
GOLINT = golint

BUILDFLAGS =
LDFLAGS    = -ldflags="-X main.version=$(VERSION)"
TESTFLAGS  = -cover

bindir = $(shell $(GO) env GOPATH)/bin

CMDS    = $(shell ls cmd)
SOURCES = $(shell find . -type f -name '*.go')
VERSION = $(shell git describe --always --dirty)

all: $(CMDS)

$(CMDS): Makefile $(SOURCES)
	$(GO) build -o $@ $(BUILDFLAGS) $(LDFLAGS) ./cmd/$@

deps:
	$(GO) get -u -t $(BUILDFLAGS) ./...

check:
	-$(GO) fmt ./...
	-$(GO) vet ./...
	-$(GOLINT) ./...

test:
	$(GO) test $(TESTFLAGS) ./...

install: $(CMDS)
	mkdir -p $(DESTDIR)$(bindir)
	install $(CMDS) $(DESTDIR)$(bindir)

config:
	tools/make-config.sh

clean:
	$(GO) clean ./...
	rm -f $(CMDS)

.PHONY: all check clean config deps install test
