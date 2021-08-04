SERVICE		?= $(shell basename `go list`)
VERSION		?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(PWD)/.version 2> /dev/null || echo v0)
PACKAGE		?= $(shell go list)
PACKAGES	?= $(shell go list ./...)
FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: help clean fmt lint vet test test-cover build all

default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

all:    ## clean, format, build and unit test
	make clean-all
	make gofmt
	make build
	make test

install:    ## build and install go application executable
	go install -v ./...

env:    ## Print useful environment variables to stdout
	@echo $(CURDIR)
	@echo $(SERVICE)
	@echo $(PACKAGE)
	@echo $(VERSION)

clean:  ## go clean
	go clean

fmt:    ## format the go source files
	go fmt ./...
	goimports -w $(FILES)

lint:   ## run go lint on the source files
	golint $(PACKAGES)

vet:    ## run go vet on the source files
	go vet ./...

doc:    ## generate godocs and start a local documentation webserver on port 8085
	godoc -http=:8085 -index

build:
	go build example/main.go

test:
	go test -v ./... -short

test-it:
	go test -v ./...

# Generate test coverage
test-cover:     ## Run test coverage and generate html report
	rm -fr coverage
	mkdir coverage
	go list -f '{{if gt (len .TestGoFiles) 0}}"go test -covermode count -coverprofile {{.Name}}.coverprofile -coverpkg ./... {{.ImportPath}}"{{end}}' ./... | xargs -I {} bash -c {}
	echo "mode: count" > coverage/cover.out
	grep -h -v "^mode:" *.coverprofile >> "coverage/cover.out"
	rm *.coverprofile
	go tool cover -html=coverage/cover.out -o=coverage/cover.html

test-all: test test-cover

binary:   ## Build Golang application binary with settings to enable it to run in a Docker scratch container.
	CGO_ENABLED=0 GOOS=linux go build  -ldflags '-s' -installsuffix cgo example/main.go
