.DEFAULT_GOAL := build

fmt:
        go fmt ./...
.PHONY:fmt

lint: fmt
        golangci-lint ./...
.PHONY:lint

build: vet
        go build ?
.PHONY:build
