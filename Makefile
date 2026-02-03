MACOS_SERVICE=local.golangid
PROGRAM_NAME=www-golangid

.PHONY: all
all: install

.PHONY: clean
clean:
	go clean -i ./cmd/$(PROGRAM_NAME)
	rm -f ./$(PROGRAM_NAME)
	find ./content -name "*.html" -delete

.PHONY: embed
embed:
	go run ./cmd/www-golangid embed

.PHONY: build
build: embed
	go build ./cmd/$(PROGRAM_NAME)

.PHONY: install
install: embed
	go install ./cmd/$(PROGRAM_NAME)

.PHONY: dev
dev:
	go run ./cmd/$(PROGRAM_NAME) -dev


##---- Deployment.

.PHONY: deploy-build
deploy-build: embed
	unset CGO_ENABLED; \
	GOOS=linux GOARCH=amd64 go build ./cmd/$(PROGRAM_NAME)

.PHONY: deploy-remote
deploy-remote: deploy-build
	rsync --progress ./$(PROGRAM_NAME) golang-id.org:/data/app/bin/

.PHONY: on-webhook
on-webhook: CGO_ENABLED=0
on-webhook: GOOS=linux
on-webhook: GOARCH=amd64
on-webhook: build
	sudo rsync --progress ./$(PROGRAM_NAME) /data/app/bin/$(PROGRAM_NAME)


##---- Local installation.

.PHONY: install-local
install-local: deploy-local
	sudo cp ./cmd/$(PROGRAM_NAME)/$(PROGRAM_NAME).path    /etc/systemd/system/
	sudo cp ./cmd/$(PROGRAM_NAME)/$(PROGRAM_NAME).service /etc/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl enable $(PROGRAM_NAME)
	sudo systemctl start $(PROGRAM_NAME)


##---- Local installation on macos.

.PHONY: macos-install-local
macos-install-local:
	cp cmd/$(PROGRAM_NAME)/$(MACOS_SERVICE).plist ~/Library/LaunchAgents/
	mkdir -p ~/bin
	CGO_ENABLED=0 go build ./cmd/www-golangid
	mv $(PROGRAM_NAME) ~/bin/
	launchctl load ~/Library/LaunchAgents/$(MACOS_SERVICE).plist
	launchctl start $(MACOS_SERVICE)

.PHONY: macos-deploy-local
macos-deploy-local: build
	sudo cp -f ./$(PROGRAM_NAME) /usr/local/bin/
