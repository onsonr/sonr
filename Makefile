SHELL := /bin/zsh # Set Shell

# GoMobile Commands
GOMOBILE=gomobile
GCMSG=git-commitmsg
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind

# GoMobile Directories
IOS_BUILDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
ANDROID_BUILDDIR=/Users/prad/Sonr/plugin/android/libs

# Platform Specific Parameters
IOS_ARTIFACT=$(IOS_BUILDDIR)/Core.framework
IOS_TARGET=ios/arm64
ANDROID_ARTIFACT=$(ANDROID_BUILDDIR)/io.sonr.core.aar
ANDROID_TARGET=android

# Gomobile Build Commands
BUILD_IOS="cd bind && $(GOCLEAN) &&  $(GOBIND) -target=$(IOS_TARGET) -v -o $(IOS_ARTIFACT)"
BUILD_ANDROID="cd bind && $(GOCLEAN) && $(GOBIND) -target=$(ANDROID_TARGET) -v -o $(ANDROID_ARTIFACT)"

# Proto Directories
PB_PATH="/Users/prad/Sonr/core/proto"
CORE_PB_DIR="/Users/prad/Sonr/core/pkg/models"
PLUGIN_PB_DIR="/Users/prad/Sonr/plugin/lib/models"

# Proto Build Commands
PB_FOR_GO="--go_out=$(CORE_PB_DIR)"
PB_FOR_DART="--dart_out=$(PLUGIN_PB_DIR)"

all: protoc ios android 
	@cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅   FINISHED ALL TASKS  ✅ ✅ ✅  --------------"
	@echo "--------------------------------------------------------------"

ios:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 BEGIN IOS BIND 📱 ---------------------------"
	@echo "--------------------------------------------------------------"
	@rm -rf $(IOS_BUILDDIR) 2>/dev/null
	@mkdir -p $(IOS_BUILDDIR)
	eval $(BUILD_IOS)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 COMPLETE IOS BIND 📱 ------------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

android:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "--------------- 🤖 BEGIN ANDROID BIND 🤖 ----------------------"
	@echo "--------------------------------------------------------------"
	@rm -rf $(ANDROID_BUILDDIR) 2>/dev/null
	@mkdir -p $(ANDROID_BUILDDIR)
	eval $(BUILD_ANDROID)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🤖  COMPLETE ANDROID BIND 🤖  -------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

protoc:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 START PROTOBUFS COMPILE 🛸 -------------------"
	@echo "--------------------------------------------------------------"
	@cd proto && protoc -I. --proto_path=$(PB_PATH) $(PB_FOR_GO) data.proto event.proto message.proto user.proto error.proto
	@cd proto && protoc -I. --proto_path=$(PB_PATH) $(PB_FOR_DART) data.proto event.proto message.proto user.proto error.proto
	@echo "Finished Compiling ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 COMPILED ALL PROTOBUFS 🛸 --------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

clean:
	cd bind && $(GOCLEAN)
	go mod tidy
	rm -rf $(IOS_BUILDDIR)
	rm -rf $(ANDROID_BUILDDIR)