all:
	go build ./...
	(cd cmd/trumpet && go build)

.PHONY: all
