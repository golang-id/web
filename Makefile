.PHONY: clean all generate build deploy

PROGRAM_NAME=www-golangid

all: install

clean:
	go clean -i ./cmd/$(PROGRAM_NAME)
	rm -f ./$(PROGRAM_NAME)
	find ./content -name "*.html" -delete

generate:
	go generate

build: generate
	go build ./cmd/$(PROGRAM_NAME)

install: generate
	go install ./cmd/$(PROGRAM_NAME)

serve: generate
	DEBUG=1 go run ./cmd/$(PROGRAM_NAME) -port=5080

deploy: build-deploy
	rsync --progress ./$(PROGRAM_NAME) golangid-webserver:~/bin/

build-deploy: generate
	unset CGO_ENABLED; \
	GOOS=linux GOARCH=amd64 go build ./cmd/$(PROGRAM_NAME)

install-local: deploy-local
	sudo cp ./cmd/$(PROGRAM_NAME)/$(PROGRAM_NAME).path    /etc/systemd/system/
	sudo cp ./cmd/$(PROGRAM_NAME)/$(PROGRAM_NAME).service /etc/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl enable $(PROGRAM_NAME)
	sudo systemctl start $(PROGRAM_NAME)

deploy-local: build
	sudo cp -f ./$(PROGRAM_NAME) /usr/local/bin/
