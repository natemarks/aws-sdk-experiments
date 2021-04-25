.DEFAULT_GOAL := help

# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))
OUT := aws-sdk-experiments
PKG := github.com/natemarks/aws-sdk-experiments
VERSION := 0.0.0
COMMIT := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
CMDS := describe_clusters listS3Contents

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build-clean: ## delete and re-create teh build directory
	@rm -rf build
	mkdir build

$(CMDS):
	go build  -v -o build/$@ \
	-ldflags="-X 'github.com/natemarks/aws-sdk-experiments/app/build.Version=${COMMIT}'" ${PKG}/cmd/$@

# this would let you run the build on any docker host without install ing go
#$(CMDS):  ## Creates a target for each executable package (example: describe_clusters => cmd/describe_clusters)
#	mkdir -p build
#	docker run --rm \
#	--env GOOS=darwin \
#	--env GOARCH=amd64 \
#	-v $(shell pwd):/usr/src/myapp \
#	-w /usr/src/myapp golang:1.15.8 go build -v \
#	-o /usr/src/myapp/build/$@ \
#	-ldflags="-X 'github.com/natemarks/aws-sdk-experiments/app/build.Version=${COMMIT}'" ${PKG}/cmd/$@

build_commands: build-clean $(CMDS) ## BUIld all of the executables


test:
	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint build_commands

.PHONY: build_commands