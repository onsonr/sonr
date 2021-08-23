# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
SONR_ROOT_DIR=/Users/prad/Sonr
CORE_DIR=$(SONR_ROOT_DIR)/core
CORE_RPC_DIR=$(SONR_ROOT_DIR)/core/cmd/rpc
CORE_BIND_DIR=$(SONR_ROOT_DIR)/core/cmd/bind

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
PROTO_DIR_GO=$(SONR_ROOT_DIR)/core/pkg
PROTO_DIR_DART=$(SONR_ROOT_DIR)/plugin/lib/src/data/protobuf
PROTO_DIR_DOCS=$(SONR_ROOT_DIR)/docs
PROTO_DIR_RPC=$(SONR_ROOT_DIR)/electron/assets

# @ Proto Items Lists
PROTO_LIST_ALL=api.proto data.proto core.proto peer.proto error.proto user.proto
PROTO_LIST_CLIENT=api.proto data.proto peer.proto error.proto user.proto

# @ Proto Build Commands
PROTO_GEN_GO="--go_out=$(PROTO_DIR_GO)"
PROTO_GEN_RPC="--go-grpc_out=$(PROTO_DIR_GO)"
PROTO_GEN_DART="--dart_out=$(PROTO_DIR_DART)"
PROTO_GEN_DOCS="--doc_out=$(PROTO_DIR_DOCS)"

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
## [protobuf]     :   Compiles Protobuf models for Core Library and Plugin
protobuf:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 START PROTOBUFS COMPILE 🛸 -------------------"
	@echo "--------------------------------------------------------------"
	@cd $(PROTO_DEF_PATH) && protoc -I. --proto_path=$(PROTO_DEF_PATH) $(PROTO_GEN_DOCS) $(PROTO_LIST_ALL)
	@cd $(PROTO_DEF_PATH) && protoc -I. --proto_path=$(PROTO_DEF_PATH) $(PROTO_GEN_GO) $(PROTO_LIST_ALL)
	@cd $(PROTO_DEF_PATH) && protoc -I. --proto_path=$(PROTO_DEF_PATH) $(PROTO_GEN_RPC) $(PROTO_LIST_ALL)
	@cd $(PROTO_DEF_PATH) && protoc -I. --proto_path=$(PROTO_DEF_PATH) $(PROTO_GEN_DART) $(PROTO_LIST_CLIENT)
	@rm -rf $(PROTO_DIR_RPC)/proto
	@cp -R $(PROTO_DEF_PATH) $(PROTO_DIR_RPC)/proto
	@echo "✅ Finished Compiling ➡ " && date
	@echo ""

##
## [release]   :   Upload RPC Binary Artifact to S3
release: protobuf
	@echo "Building Artifacts..."
	@cd $(CORE_RPC_DIR) && goreleaser release --rm-dist
	@echo "Cleaning up build cache..."
	@cd $(CORE_DIR) && go mod tidy
	@rm -rf $(DIST_DIR_DARWIN_AMD)
	@rm -rf $(DIST_DIR_DARWIN_ARM)
	@rm -rf $(DIST_DIR_LINUX_AMD)
	@rm -rf $(DIST_DIR_LINUX_ARM)
	@rm -rf $(DIST_DIR_WIN)
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
