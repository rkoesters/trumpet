GO         = go
BUILDFLAGS =
LDFLAGS    = -ldflags="-X main.version=$(VERSION)"
TESTFLAGS  = -cover

EXECNAME = trumpet
SOURCES  = $(shell find . -type f -name '*.go')
DEPS     = $(shell tools/list-deps.sh ./...)
VERSION  = $(shell git describe --always --dirty)

all: $(EXECNAME)

deps:
	$(GO) get -u $(BUILDFLAGS) golang.org/x/lint/golint
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

$(EXECNAME): Makefile $(SOURCES)
	$(GO) build -o $@ $(BUILDFLAGS) $(LDFLAGS) ./cmd/trumpet

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

.PHONY: all check clean config deps install test
