SHELL := /bin/zsh # Set Shell

# Go Commands
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind

# Plugin Directories
IOS_BUILDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
IOS_ARTIFACT= $(IOS_BUILDDIR)/Core.framework
ANDROID_BUILDDIR=/Users/prad/Sonr/plugin/android/libs
ANDROID_ARTIFACT= $(ANDROID_BUILDDIR)/io.sonr.core.aar

# Proto Directories
PB_PATH="/Users/prad/Sonr/core/internal/models"
CORE_PB_DIR="/Users/prad/Sonr/core/internal/models"
PLUGIN_PB_DIR="/Users/prad/Sonr/plugin/lib/models"

# Proto Build Commands
PB_BUILD_CORE="--go_out=$(CORE_PB_DIR)"
PB_BUILD_PLUGIN="--dart_out=$(PLUGIN_PB_DIR)"

all: protoc ios android
	@go mod tidy
	@cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅   FINISHED ALL TASKS  ✅ ✅ ✅  --------------"
	@echo "--------------------------------------------------------------"


## android :    Builds Android Bind at Plugin Path
android:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "--------------- 🤖 BEGIN ANDROID BIND 🤖 ----------------------"
	@echo "--------------------------------------------------------------"
	cd bind && GODEBUG=asyncpreemptoff=1 gomobile bind -ldflags='-s -w' -target=android/arm64 -v -o $(ANDROID_ARTIFACT)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🤖  COMPLETE ANDROID BIND 🤖  -------------------"
	@echo "--------------------------------------------------------------"
	@echo ""


## ios     :    Builds iOS Bind at Plugin Path
ios:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 BEGIN IOS BIND 📱 ---------------------------"
	@echo "--------------------------------------------------------------"
	cd bind && GODEBUG=asyncpreemptoff=1 gomobile bind -ldflags='-s -w' -target=ios/arm64 -v -o $(IOS_ARTIFACT)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 COMPLETE IOS BIND 📱 ------------------------"
	@echo "--------------------------------------------------------------"
	@echo ""


## protoc  :    Compiles Protobuf models for Core Library and Plugin
protoc:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 START PROTOBUFS COMPILE 🛸 -------------------"
	@echo "--------------------------------------------------------------"
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_CORE) api.proto data.proto core.proto
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_PLUGIN) api.proto data.proto
	@echo "Finished Compiling ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 COMPILED ALL PROTOBUFS 🛸 --------------------"
	@echo "--------------------------------------------------------------"
	@echo ""


## reset   :    Cleans Gomobile, Removes Framworks from Plugin, and Inits Gomobile
reset:
	cd bind && $(GOCLEAN)
	go mod tidy
	rm -rf $(IOS_BUILDDIR)
	rm -rf $(ANDROID_BUILDDIR)
	mkdir -p $(IOS_BUILDDIR)
	mkdir -p $(ANDROID_BUILDDIR)
	cd bind && gomobile init


help : Makefile
	@sed -n 's/^##//p' $<
