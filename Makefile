GO         = go
BUILDFLAGS = -v
DEPS       = $(shell tools/list-deps.sh ./...)

all: build

build:
	$(GO) build $(BUILDFLAGS) ./...

deps:
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

clean:
	$(GO) clean ./...

.PHONY: all build deps clean
