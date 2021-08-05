SHELL=/bin/zsh # Set Shell
# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
SONR_ROOT_DIR=/Users/prad/Sonr
CORE_DIR=$(SONR_ROOT_DIR)/core
CORE_CMD_DIR=$(SONR_ROOT_DIR)/core/cmd

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
PROTO_DEF_PATH=/Users/prad/Sonr/core/pkg/proto
APP_ROOT_DIR =/Users/prad/Sonr/app

# @ Packaging Vars/Commands
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind -ldflags='-s -w' -v
GOBIND_ANDROID=$(GOBIND) -target=android
GOBIND_IOS=$(GOBIND) -target=ios -bundleid=io.sonr.core

# @ Bind Directories
BIND_DIR_CORE=$(SONR_ROOT_DIR)/core/bind
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
DIST_DIR=$(SONR_ROOT_DIR)/core/cmd/dist
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
bind: proto bind.ios bind.android
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
	cd $(BIND_DIR_CORE) && $(GOBIND_ANDROID) -o $(BIND_ANDROID_ARTIFACT)
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
	cd $(BIND_DIR_CORE) && $(GOBIND_IOS) -o $(BIND_IOS_ARTIFACT)
	@echo "✅ Finished Binding ➡ " && date
	@echo ""

##
## [proto]     :   Compiles Protobuf models for Core Library and Plugin
proto:
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
release:
	@echo "Bumping Release Version.."
	@cd $(CORE_DIR) && gitmoji -c
	@cd $(CORE_DIR) && bump patch
	@echo "Bumping Release Version... DONE"
	@echo "Building Artifacts..."
	@cd $(CORE_CMD_DIR) && goreleaser release --rm-dist
	@cd $(CORE_DIR) && git push origin --tags
	@cd $(CORE_DIR) && git push
	@echo "Cleaning up build cache..."
	@cd $(CORE_DIR) && go mod tidy
	@rm -rf $(DIST_DIR_DARWIN_AMD)
	@rm -rf $(DIST_DIR_DARWIN_ARM)
	@rm -rf $(DIST_DIR_LINUX_AMD)
	@rm -rf $(DIST_DIR_LINUX_ARM)
	@rm -rf $(DIST_DIR_WIN)
	@echo "✅ Finished Releasing RPC Binary ➡ " && date
	@cd /System/Library/Sounds && afplay Glass.aiff

## [upgrade]   :   Binds Binary, Creates Protobufs, and Updates App
upgrade: proto bind.ios bind.android
	@go mod tidy
	@echo "-----------------------------------------------------------"
	@echo "------------- 🔄  START PLUGIN UPDATE 🔄 -------------------"
	@echo "------------------------------------------------------------"
	cd $(APP_ROOT_DIR) && make update
	@echo ""

## [clean]     :   Reinitializes Gomobile and Removes Framworks from Plugin
clean:
	cd $(BIND_DIR) && $(GOCLEAN)
	go mod tidy
	go clean -cache -x
	rm -rf $(BIND_DIR_IOS)
	rm -rf $(BIND_DIR_ANDROID)
	mkdir -p $(BIND_DIR_IOS)
	mkdir -p $(BIND_DIR_ANDROID)
	cd $(BIND_DIR_CORE) && gomobile init

##
##
## Shortcuts   : (b) => bind
##               └─ (bi) => bind.ios
##               └─ (ba) => bind.android
##               (p) => proto
##               (r) => release
##               (u) => upgrade
##               (c) => clean
b:bind
bi:bind.ios
ba:bind.android
p:proto
r:release
u:upgrade
c:clean
