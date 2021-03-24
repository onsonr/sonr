SHELL=/bin/zsh # Set Shell
SONR_ROOT_DIR=/Users/prad/Sonr # Set this to Folder of Sonr
ANDROID_BINDDIR=/Users/prad/Sonr/plugin/android/libs
IOS_BINDDIR=/Users/prad/Sonr/plugin/ios/Frameworks

# @ Packaging Vars/Commands
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind -ldflags='-s -w' -v

# @ Bind Directories
BIND_DIR=/Users/prad/Sonr/core/bind
IOS_ARTIFACT= $(IOS_BINDDIR)/Core.framework
ANDROID_ARTIFACT= $(ANDROID_BINDDIR)/io.sonr.core.aar

# @ Proto Directories
PB_PATH=/Users/prad/Sonr/core/internal/models
CONTACT_PB_DIR=/Users/prad/Sonr/contact/lib/src/data/models
CORE_PB_DIR=/Users/prad/Sonr/core/internal/models
PLUGIN_PB_DIR=/Users/prad/Sonr/plugin/lib/src/core/models
PROTO_DOC_DIR=/Users/prad/Sonr/docs/proto

# @ Proto Build Commands
PB_BUILD_CONTACT="--dart_out=$(CONTACT_PB_DIR)"
PB_BUILD_CORE="--go_out=$(CORE_PB_DIR)"
PB_BUILD_PLUGIN="--dart_out=$(PLUGIN_PB_DIR)"

all: Makefile
	@figlet -f slant Sonr Core Actions
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
	cd $(BIND_DIR) && $(GOBIND) -target=android -o $(ANDROID_ARTIFACT)
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
	cd $(BIND_DIR) && $(GOBIND) -target=ios -o $(IOS_ARTIFACT)
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
	@cd internal/models && protoc --doc_out=$(PROTO_DOC_DIR) --doc_opt=html,index.html api.proto data.proto core.proto user.proto
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_CORE) api.proto data.proto core.proto user.proto
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_CONTACT) user.proto
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_PLUGIN) api.proto data.proto user.proto
	@echo "Finished Compiling ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 COMPILED ALL PROTOBUFS 🛸 --------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

## [clean]     :   Cleans Gomobile, Removes Framworks from Plugin, and Initializes Gomobile
clean:
	cd $(BIND_DIR) && $(GOCLEAN)
	go mod tidy
	go clean -cache -x
	go clean -modcache -x
	rm -rf $(IOS_BINDDIR)
	rm -rf $(ANDROID_BINDDIR)
	mkdir -p $(IOS_BINDDIR)
	mkdir -p $(ANDROID_BINDDIR)
	cd $(BIND_DIR) && gomobile init
