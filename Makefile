# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
SONR_ROOT_DIR=/Users/prad/Sonr
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CORE_DIR=$(SONR_ROOT_DIR)/core
CORE_RPC_DIR=$(SONR_ROOT_DIR)/core/cmd/rpc
CORE_BIND_DIR=$(SONR_ROOT_DIR)/core/cmd/bind
ELECTRON_BIN_DIR=$(SONR_ROOT_DIR)/electron/assets/bin/darwin

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
PROTO_DEF_PATH=/Users/prad/Sonr/core/proto
APP_ROOT_DIR =/Users/prad/Sonr/app

# @ Packaging Vars/Commands
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind -ldflags='-s -w' -v
GOBIND_ANDROID=$(GOBIND) -target=android
GOBIND_IOS=$(GOBIND) -target=ios -bundleid=io.sonr.core

# @ Bind Directories
BIND_DIR_ANDROID=$(SONR_ROOT_DIR)/plugin/android/libs
BIND_DIR_IOS=$(SONR_ROOT_DIR)/plugin/ios/Frameworks
BIND_IOS_ARTIFACT= $(BIND_DIR_IOS)/Core.framework
BIND_ANDROID_ARTIFACT= $(BIND_DIR_ANDROID)/io.sonr.core.aar

# @ Proto Directories
PROTO_DIR_DART=$(SONR_ROOT_DIR)/plugin/lib/src
falafel=$(which falafel)

# Name of the package for the generated APIs.
pkg="bind"

# The package where the protobuf definitions originally are found.
target_pkg="github.com/sonr-io/core/proto"

# A mapping from grpc service to name of the custom listeners. The grpc server
# must be configured to listen on these.
listeners="location=locationLis profileunlocker=profileUnlockerLis"

# Set to 1 to create boiler plate grpc client code and listeners. If more than
# one proto file is being parsed, it should only be done once.
mem_rpc=1
opts="package_name=$(pkg),target_package=$(target_pkg),mem_rpc=$(mem_rpc)"
PROTO_LIST_ALL=${ROOT_DIR}/proto/**/*.proto
PROTO_LIST_CLIENT=${ROOT_DIR}/proto/client/*.proto
PROTO_LIST_COMMON=${ROOT_DIR}/proto/common/*.proto
MODULE_NAME=github.com/sonr-io/core
GO_OPT_FLAG=--go_opt=module=${MODULE_NAME}
GRPC_OPT_FLAG=--go-grpc_opt=module=${MODULE_NAME}
PROTO_GEN_GO="--go_out=."
PROTO_GEN_DOCS="--doc_out=./docs"
PROTO_GEN_DART="--dart_out=grpc:$(PROTO_DIR_DART)"

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
	@echo "✅ Finished Binding ➡ " && date
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
	@echo "✅ Finished Binding ➡ " && date
	@echo ""

##
## [proto]     :   Compiles Protobuf models for Core Library and Plugin
protobuf:
	@echo "----"
	@echo "Sonr: Compiling Protobufs"
	@echo "----"
	@echo "Generating Protobuf Go code..."
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_GO) $(GO_OPT_FLAG)
	@echo "Generating Protobuf Go RPC code..."
# @protoc $(PROTO_LIST_CLIENT) --plugin=protoc-gen-custom=$(falafel) --custom_opt=$(opts) --proto_path=$(ROOT_DIR) $(PROTO_GEN_RPC)
	@echo "Generating Protobuf Dart code..."
	@protoc $(PROTO_LIST_CLIENT) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)
	@protoc $(PROTO_LIST_COMMON) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)
	@echo "----"
	@echo "✅ Finished Compiling ➡ `date`"

##
## [release]   :   Upload RPC Binary Artifact to S3
release: protobuf
	@echo "Building Artifacts..."
	@cd $(CORE_RPC_DIR) && goreleaser release --rm-dist
	@echo "Cleaning up build cache..."
	@cd $(CORE_DIR) && go mod tidy
	@rm -rf $(ELECTRON_BIN_DIR)
	@mkdir -p $(ELECTRON_BIN_DIR)
	@mv $(DIST_DIR_DARWIN_ARM) $(ELECTRON_BIN_DIR)
	@rm -rf $(DIST_DIR)
	@echo "✅ Finished Releasing RPC Binary ➡ " && date
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
