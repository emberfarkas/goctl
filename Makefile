VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)

.PHONY: init
init:
	go install github.com/UnnoTed/fileb0x@latest

.PHONY: build
build:
	go build -ldflags "-X github.com/emberfarkas/goctl/internal/version.Version=$(VERSION) -X github.com/emberfarkas/goctl/internal/version.Branch=$(BRANCH) -X github.com/emberfarkas/goctl/internal/version.Revision=$(REVISION) -X github.com/emberfarkas/goctl/internal/version.BuildDate=$(BUILD_DATE)"