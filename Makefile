GO         = go
BUILDFLAGS = -v
DEPS       = $(shell tools/list-deps.sh ./...)

all:
	$(GO) build $(BUILDFLAGS) ./...
	(cd cmd/trumpet && $(GO) build $(BUILDFLAGS))

deps:
	$(GO) get -u $(BUILDFLAGS) $(DEPS)

clean:
	$(GO) clean ./...

.PHONY: all deps clean
