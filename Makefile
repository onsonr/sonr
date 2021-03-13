SHELL=/bin/zsh # Set Shell

# @ Packaging Vars/Commands
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind -ldflags='-s -w' -target=android -v -o
BIND_DIR="/Users/prad/Sonr/core/cmd/bind"

# @ Bind Directories
IOS_BINDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
IOS_ARTIFACT= $(IOS_BINDDIR)/Core.framework
ANDROID_BINDDIR=/Users/prad/Sonr/plugin/android/libs
ANDROID_ARTIFACT= $(ANDROID_BINDDIR)/io.sonr.core.aar

# @ Proto Directories
PB_PATH="/Users/prad/Sonr/core/pkg/models"
CONTACT_PB_DIR="/Users/prad/Sonr/contact/lib/src/data/models"
CORE_PB_DIR="/Users/prad/Sonr/core/pkg/models"
PLUGIN_PB_DIR="/Users/prad/Sonr/plugin/lib/src/core/models"
PROTO_DOC_OUT="/Users/prad/Sonr/docs/proto"

# @ Proto Build Commands
PB_BUILD_CONTACT="--dart_out=$(CONTACT_PB_DIR)"
PB_BUILD_CORE="--go_out=$(CORE_PB_DIR)"
PB_BUILD_PLUGIN="--dart_out=$(PLUGIN_PB_DIR)"

all: Makefile
	@echo '--- Sonr Core Module Actions ---'
	@echo ''
	@sed -n 's/^##//p ' $<

## bind        :   Binds Android and iOS for Plugin Path
bind: proto bind.ios bind.android
	@go mod tidy
	@cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo ""
	@echo "----------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅   FINISHED MOBILE BIND  ✅ ✅ ✅  --------------"
	@echo "----------------------------------------------------------------"


## └─ android       - Android AAR
bind.android:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "--------------- 🤖 BEGIN ANDROID BIND 🤖 ----------------------"
	@echo "--------------------------------------------------------------"
	@go get golang.org/x/mobile/bind
	@gomobile init
	cd bind && $(GOBIND) $(ANDROID_ARTIFACT)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🤖  COMPLETE ANDROID BIND 🤖  -------------------"
	@echo "--------------------------------------------------------------"
	@echo ""


## └─ ios           - iOS Framework
bind.ios:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 BEGIN IOS BIND 📱 ---------------------------"
	@echo "--------------------------------------------------------------"
	@go get golang.org/x/mobile/bind
	cd bind && $(GOBIND) $(IOS_ARTIFACT)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 COMPLETE IOS BIND 📱 ------------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

##
## [proto]     :   Compiles Protobuf models for Core Library and Plugin
proto:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 START PROTOBUFS COMPILE 🛸 -------------------"
	@echo "--------------------------------------------------------------"
	@cd pkg/models && protoc --doc_out=$(PROTO_DOC_OUT) --doc_opt=html,index.html api.proto data.proto core.proto user.proto
	@cd pkg/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_CORE) api.proto data.proto core.proto user.proto
	@cd pkg/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_CONTACT) api.proto data.proto user.proto
	@cd pkg/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_PLUGIN) user.proto
	@echo "Finished Compiling ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 COMPILED ALL PROTOBUFS 🛸 --------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

## [reset]     :   Cleans Gomobile, Removes Framworks from Plugin, and Initializes Gomobile
reset:
	cd bind && $(GOCLEAN)
	go mod tidy
	go clean -cache -x
	go clean -modcache -x
	rm -rf $(IOS_BINDDIR)
	rm -rf $(ANDROID_BINDDIR)
	mkdir -p $(IOS_BINDDIR)
	mkdir -p $(ANDROID_BINDDIR)
	cd bind && gomobile init
