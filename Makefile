.PHONY: all generate build deploy

all: build

generate:
	go generate

build: generate
	go build ./cmd/golangid

deploy: build
	rsync ./golangid gcp-webserver:~/bin/
