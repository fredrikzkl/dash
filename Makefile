.DEFAULT_GOAL := dash

.PHONY: all deps format build test lint shadow checkmake help

dash: deps tidy format fix-imports lint vet test build

## Download makefile dependencies
deps:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest

## Tidy and format project code
format:
	gofmt -w .

tidy:
	go mod tidy

fix-imports:
	goimports -w .

build:
	go build -o dash .

install: build
	sudo mv dash /usr/local/bin/dash

run:
	go run .

## Run project tests
test:
	go test ./...

## Run linters
lint:
	golangci-lint run --config .golangci.yml --timeout 3m --verbose

## Run go vet on the project
vet:
	go vet -vettool=$(shell which shadow) ./...

modernize:
	modernize -fix ./...

tail:
	tail -f debug.log

clear-debug:
	> debug.log
