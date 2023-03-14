all: lint test
.PHONY: all

lint:
	golangci-lint run --timeout 600s
test:
	go test -v -failfast -timeout 600s ./...

.DEFAULT_GOAL := lint

