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

$(CMDS): $(SOURCES) Makefile
	$(GO) build -o $@ $(BUILDFLAGS) $(LDFLAGS) ./cmd/$@

deps:
	$(GO) get -u -t $(BUILDFLAGS) ./...

check:
	-$(GO) fmt ./...
	-$(GO) vet ./...
	-$(GOLINT) ./...

test:
	$(GO) test $(TESTFLAGS) ./...

clean:
	$(GO) clean ./...
	rm -f $(CMDS)

install: $(CMDS)
	mkdir -p $(DESTDIR)$(bindir)
	install $(CMDS) $(DESTDIR)$(bindir)

uninstall:
	rm $(addprefix $(DESTDIR)$(bindir)/,$(CMDS))

.PHONY: all check clean deps install test uninstall
