GOPATH := ${PWD}:${PWD}/_vendor
export GOPATH

default: install

env: env
	go env

install: fmt
	go install ./...

build: fmt vet
	go build ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

vendor_clean:
	rm -rf ./_vendor/src
