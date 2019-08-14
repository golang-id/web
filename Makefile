.PHONY: clean all generate build deploy

all: build

clean:
	go clean -i ./cmd/golangid
	rm -f ./golangid
	find ./content -name "*.html" -delete

generate:
	go generate

build: generate
	go build ./cmd/golangid

serve: generate
	DEBUG=1 go run ./cmd/golangid

deploy: build
	rsync --progress ./golangid gcp-webserver:~/bin/
