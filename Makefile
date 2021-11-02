SHELL=/bin/bash

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
SONR_ROOT_DIR=/Users/prad/Developer
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CORE_DIR=$(SONR_ROOT_DIR)/core
DESKTOP_DIR=$(SONR_ROOT_DIR)/desktop
MOBILE_DIR=$(SONR_ROOT_DIR)/mobile
CORE_FULL_DIR=$(SONR_ROOT_DIR)/core/cmd/sonrd
CORE_BIND_DIR=$(SONR_ROOT_DIR)/core/cmd/lib
ELECTRON_BIN_DIR=$(SONR_ROOT_DIR)/electron/assets/bin/darwin
PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
LD_LIBRARY_PATH=/opt/homebrew/bin/ffmpeg/4.4_2/lib

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
PROTO_DEF_PATH=/Users/prad/Developer/core/proto
APP_ROOT_DIR =/Users/prad/Developer/mobile/

# @ Packaging Vars/Commands
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind -ldflags='-s -w' -v
GOBIND_ANDROID=$(GOBIND) -target=android/arm64 -androidapi=24
GOBIND_IOS=$(GOBIND) -target=ios/arm64 -bundleid=io.sonr.core

# @ Bind Directories
BIND_DIR_ANDROID=$(SONR_ROOT_DIR)/mobile/android/libs
BIND_DIR_IOS=$(SONR_ROOT_DIR)/mobile/ios/Frameworks
BIND_IOS_ARTIFACT= $(BIND_DIR_IOS)/Core.xcframework
BIND_ANDROID_ARTIFACT= $(BIND_DIR_ANDROID)/io.sonr.core.aar

# @ Proto Directories
PROTO_DIR_DART=$(SONR_ROOT_DIR)/mobile/lib/src
PROTO_DIR_TS=$(SONR_ROOT_DIR)/desktop/assets/proto
PROTO_LIST_ALL=${ROOT_DIR}/proto/**/*.proto
PROTO_LIST_API=${ROOT_DIR}/proto/api/*.proto
PROTO_LIST_COMMON=${ROOT_DIR}/proto/common/*.proto
PROTO_FILE_NODE_CLIENT=${ROOT_DIR}/proto/node/client.proto
PROTO_FILE_NODE_HIGHWAY=${ROOT_DIR}/proto/node/highway.proto
PROTO_LIST_PROTOCOLS=${ROOT_DIR}/proto/protocols/*.proto
PROTO_LIST_WALLET=${ROOT_DIR}/proto/wallet/*.proto
MODULE_NAME=github.com/sonr-io/core
GO_OPT_FLAG=--go_opt=module=${MODULE_NAME}
GRPC_OPT_FLAG=--go-grpc_opt=module=${MODULE_NAME}
PROTO_GEN_GO="--go_out=."
PROTO_GEN_RPC="--go-grpc_out=."
PROTO_GEN_DOCS="--doc_out=docs"
PROTO_GEN_DART="--dart_out=grpc:$(PROTO_DIR_DART)"
PROTO_GEN_TS="--ts_proto_out=$(PROTO_DIR_TS)"
TS_OPT_FLAG=--ts_proto_opt=outputClientImpl=grpc-node --ts_proto_opt=useOptionals=true --ts_proto_opt=outputServices=grpc-js --ts_proto_opt=env=node --ts_proto_opt=stringEnums=true --ts_proto_opt=esModuleInterop=true
PROTOC_TS_PLUGIN_PATH=/usr/local/bin/protoc-gen-ts

# @ Distribution Release Variables
DIST_DIR=$(SONR_ROOT_DIR)/core/cmd/rpc/dist
DIST_DIR_DARWIN_AMD=$(DIST_DIR)/sonr-rpc_darwin_amd64
DIST_DIR_DARWIN_ARM=$(DIST_DIR)/sonr-rpc_darwin_arm64
DIST_DIR_LINUX_AMD=$(DIST_DIR)/sonr-rpc_linux_amd64
DIST_DIR_LINUX_ARM=$(DIST_DIR)/sonr-rpc_linux_arm64
DIST_DIR_WIN=$(DIST_DIR)/sonr-rpc_windows_amd64
DIST_ZIP_WIN=$(DIST_DIR)/*.zip

all: Makefile
	@figlet -f larry3d Sonr Core
	@echo ''
	@sed -n 's/^##//p ' $<

## bind        :   Binds Android and iOS for Plugin Path
bind: protobuf bind.ios bind.android
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo ""
	@echo ""
	@echo "----------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅  SUCCESFUL MOBILE BIND  ✅ ✅ ✅  --------------"
	@echo "----------------------------------------------------------------"


## └─ android       - Android AAR
bind.android:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "--------------- 🤖 START ANDROID BIND 🤖 ----------------------"
	@echo "--------------------------------------------------------------"
	@go get golang.org/x/mobile/bind
	@gomobile init
	cd $(CORE_BIND_DIR) && $(GOBIND_ANDROID) -o $(BIND_ANDROID_ARTIFACT)
	@echo "✅ Finished Binding ➡ `date`"
	@echo ""


## └─ ios           - iOS Framework
bind.ios:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 START IOS BIND 📱 ---------------------------"
	@echo "--------------------------------------------------------------"
	@go get golang.org/x/mobile/bind
	cd $(CORE_BIND_DIR) && $(GOBIND_IOS) -o $(BIND_IOS_ARTIFACT)
	@echo "✅ Finished Binding ➡ `date`"
	@echo ""

##
## [proto]     :   Compiles Protobuf models for Core Library and Plugin
protobuf:
	@echo "----"
	@echo "Sonr: Compiling Protobufs"
	@echo "----"
	@echo "Generating Protobuf Go code..."
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_GO) $(GO_OPT_FLAG)
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_RPC) $(GRPC_OPT_FLAG)

	@echo "Generating Protobuf Dart code..."
	@protoc $(PROTO_LIST_API) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)
	@protoc $(PROTO_LIST_COMMON) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)
	@protoc $(PROTO_FILE_NODE_CLIENT) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)

	@echo "Generating Protobuf Typescript code..."
	@rm -rf $(PROTO_DIR_TS)/data
	@mkdir -p ${PROTO_DIR_TS}
	@cp -R $(ROOT_DIR)/proto/api $(PROTO_DIR_TS)/api
	@cp -R $(ROOT_DIR)/proto/common $(PROTO_DIR_TS)/common
	@cp -R $(ROOT_DIR)/proto/node $(PROTO_DIR_TS)/node
	@cd $(DESKTOP_DIR) && yarn proto
	@rm -rf $(PROTO_DIR_TS)

	@echo "Generating Protobuf Docs..."
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DOCS)
	@echo "----"
	@echo "✅ Finished Compiling ➡ `date`"

##
## [release]   :   Upload RPC Binary Artifact to S3
release: protobuf
	@echo "Building Artifacts..."
	@cd $(CORE_FULL_DIR) && goreleaser release --rm-dist
	@echo "Cleaning up build cache..."
	@cd $(CORE_DIR) && go mod tidy
	@rm -rf $(ELECTRON_BIN_DIR)
	@mkdir -p $(ELECTRON_BIN_DIR)
	@mv $(DIST_DIR_DARWIN_ARM) $(ELECTRON_BIN_DIR)
	@rm -rf $(DIST_DIR)
	@echo "✅ Finished Releasing RPC Binary ➡ `date`"
	@cd /System/Library/Sounds && afplay Glass.aiff

## [clean]     :   Reinitializes Gomobile and Removes Framworks from Plugin
clean:
	cd $(CORE_BIND_DIR) && $(GOCLEAN)
	go mod tidy
	go clean -cache -x
	rm -rf $(BIND_DIR_IOS)
	rm -rf $(BIND_DIR_ANDROID)
	mkdir -p $(BIND_DIR_IOS)
	mkdir -p $(BIND_DIR_ANDROID)
	cd $(CORE_BIND_DIR) && gomobile init
