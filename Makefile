GO         = go
BUILDFLAGS = -v
DEPS       = $(shell tools/list-deps.sh ./...)

all: deps build

build:
	$(GO) build $(BUILDFLAGS) ./...
	(cd cmd/trumpet && $(GO) build $(BUILDFLAGS))

deps:
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

config:
	tools/make-config.sh

clean:
	$(GO) clean ./...

.PHONY: all build clean config deps
