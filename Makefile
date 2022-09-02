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
bind: bind.ios bind.android bind.web

## └─ android       - Android AAR
bind.android:
	sh $(SCRIPTS_DIR)/bind.sh -a

## └─ ios           - iOS Framework
bind.ios:
	sh $(SCRIPTS_DIR)/bind.sh -i

## └─ web           - iOS Framework
bind.web:
	sh $(SCRIPTS_DIR)/bind.sh -w

## proto       :   Compiles Go Proto Files and pushes to Buf.Build
proto: proto.go proto.buf

## └─ go            - Generate to x/*/types and thirdparty/types/*
proto.go:
	ignite generate proto-go --yes
	go mod tidy
	@echo "✅ Generated Go Proto Files"

## └─ buf           - Build and push to buf.build/sonr-io/blockchain
proto.buf:
	cd $(ROOT_DIR)/proto && buf mod update && buf build && buf push
	@echo "✅ Pushed Protos to Buf.Build"

## clean       :   Clean all artifacts and tidy
clean:
	rm -rf ./build
	rm -rf ./tmp
	rm -rf ./dist
	rm -rf ./io.sonr.motor.aar
	rm -rf ./sonr-motor.wasm
	rm -rf ./SonrMotor.xcframework
	rm -rf ./docs/.docusaurus/
	rm -rf ./docs/build
	go mod tidy
