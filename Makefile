TEST ?= $$(go list ./... | grep -v /vendor/)
BUILD_TAG ?= oslinfra

all: bin

bin: generate

quickdev: generate
	@CGO_ENABLED=0 go build -i -tags='$(BUILD_TAGS)' -o bin/oslinfra

test:
	CGO_ENABLED=0 go test -tags='$(BUILD_TAGS)' $(TEST) $(TESTARGS) -timeout=20m -parallel=4

generate:
	go generate $(go list ./... | grep -v /vendor/)

.PHONY: bin default generate test
