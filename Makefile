SHELL=/bin/bash

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
SCRIPTS_DIR=$(ROOT_DIR)/scripts

all: Makefile
	@echo ''
	@sed -n 's/^##//p ' $<

## Makefile
## > The following Makefile is used for various actions for the Sonr project.
##
## bind        :   Binds Android, iOS and Web for Plugin Path
bind: bind.ios bind.mac bind.android bind.web

## └─ android       - Android AAR
bind.android:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -a

## └─ ios           - iOS Framework
bind.ios:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -i

## └─ mac           - Mac Framework
bind.mac:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -m

## └─ web           - WASM Framework
bind.web:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -w

## └─ tar           - Build All & Tar Compress
bind.tar:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -a
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -i
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -w

## build       :   Builds macos and ubuntu releases for Sonr
build:
	env GOOS=linux GOARCH=amd64 go build -o ./build/sonr-linux-amd64 ./cmd/sonrd/main.go
	env GOOS=linux GOARCH=arm64 go build -o ./build/sonr-linux-arm64 ./cmd/sonrd/main.go
	env GOOS=darwin GOARCH=amd64 go build -o ./build/sonr-darwin-amd64 ./cmd/sonrd/main.go

## proto       :   Compiles Go Proto Files and pushes to Buf.Build
proto: proto.go proto.buf

## └─ go            - Generate to x/*/types and thirdparty/types/*
proto.go:
	ignite generate proto-go --yes
	go mod tidy
	@echo "✅ Generated Go Proto Files"

## └─ buf           - Build and push to buf.build/sonr-io/blockchain
proto.buf:
	cd $(ROOT_DIR)/proto && buf mod update && buf build
	@echo "✅ Pushed Protos to Buf.Build"

## └─ publish       - Compiles protos, buf.build publish, Zips protos in build
proto.publish:
	cd $(ROOT_DIR)/proto && buf mod update && buf build
	@echo "✅ Pushed Protos to Buf.Build"
	cp -r proto/ build/proto/

## clean       :   Clean all artifacts and tidy
clean:
	rm -rf ./build
	rm -rf ./tmp
	rm -rf ./dist
	go mod tidy
