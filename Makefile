SHELL=/bin/zsh # Set Shell
# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
SONR_ROOT_DIR=/Users/prad/Sonr

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
PROTO_DEF_PATH=/Users/prad/Sonr/core/api
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
PROTO_DIR_JS=$(SONR_ROOT_DIR)/electron/src/data/proto

# @ Proto Items Lists
PROTO_LIST_ALL=api.proto data.proto core.proto peer.proto error.proto rpc.proto user.proto
PROTO_LIST_CLIENT=api.proto data.proto peer.proto error.proto user.proto

# @ Proto Build Commands
PROTO_GEN_GO="--go_out=$(PROTO_DIR_GO)"
PROTO_GEN_RPC="--go-grpc_out=$(PROTO_DIR_GO)"
PROTO_GEN_DART="--dart_out=$(PROTO_DIR_DART)"
PROTO_GEN_DOCS="--doc_out=$(PROTO_DIR_DOCS)"
PROTO_CP_JS=$(PROTO_DEF_PATH)/{api.proto,data.proto,peer.proto,error.proto,rpc.proto,user.proto}

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
	@cd api && protoc -I. --proto_path=$(PROTO_DEF_PATH) $(PROTO_GEN_DOCS) $(PROTO_LIST_ALL)
	@cd api && protoc -I. --proto_path=$(PROTO_DEF_PATH) $(PROTO_GEN_GO) $(PROTO_LIST_ALL)
	@cd api && protoc -I. --proto_path=$(PROTO_DEF_PATH) $(PROTO_GEN_RPC) $(PROTO_LIST_ALL)
	@cd api && protoc -I. --proto_path=$(PROTO_DEF_PATH) $(PROTO_GEN_DART) $(PROTO_LIST_CLIENT)
	@cp $(PROTO_CP_JS) $(PROTO_DIR_JS)
	@echo "✅ Finished Compiling ➡ " && date
	@echo ""

##
## [run]       :   Run GRPC Server for Desktop Development
run:
	@cd cmd && go run server.go

## [upgrade]   :   Binds Binary, Creates Protobufs, and Updates App
upgrade: proto bind.ios bind.android
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
##               (u) => upgrade
##               (c) => clean
b:bind
bi:bind.ios
ba:bind.android
p:proto
u:upgrade
c:clean
