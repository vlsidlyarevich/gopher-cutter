.DEFAULT_GOAL := build

build: $(GOFILES)
	$go build ./...
