GO         = go
BUILDFLAGS =
DEPS       = $(shell tools/list-deps.sh ./...)

all: build

build:
	cd cmd/trumpet && $(GO) build $(BUILDFLAGS)

clean:
	$(GO) clean ./...

config:
	tools/make-config.sh

deps:
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

test:
	$(GO) test ./...

.PHONY: all build clean config deps test
