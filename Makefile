.PHONY: clean all generate build deploy

MACOS_SERVICE=local.golangid
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
	rsync --progress ./$(PROGRAM_NAME) personal-www:~/bin/

build-deploy: generate
	unset CGO_ENABLED; \
	GOOS=linux GOARCH=amd64 go build ./cmd/$(PROGRAM_NAME)

install-local: deploy-local
	sudo cp ./cmd/$(PROGRAM_NAME)/$(PROGRAM_NAME).path    /etc/systemd/system/
	sudo cp ./cmd/$(PROGRAM_NAME)/$(PROGRAM_NAME).service /etc/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl enable $(PROGRAM_NAME)
	sudo systemctl start $(PROGRAM_NAME)

install-local-macos:
	cp cmd/$(PROGRAM_NAME)/$(MACOS_SERVICE).plist ~/Library/LaunchAgents/
	mkdir -p ~/bin
	CGO_ENABLED=0 go build ./cmd/www-golangid
	mv $(PROGRAM_NAME) ~/bin/
	launchctl load ~/Library/LaunchAgents/$(MACOS_SERVICE).plist
	launchctl start $(MACOS_SERVICE)

deploy-local: build
	sudo cp -f ./$(PROGRAM_NAME) /usr/local/bin/
