SHELL := /bin/zsh
# GoMobile Commands
GOMOBILE=gomobile
GCMSG=git-commitmsg
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind

# GoMobile Directories
IOS_BUILDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
ANDROID_BUILDDIR=/Users/prad/Sonr/plugin/android/libs

# Proto Directories
CORE_PROTO_DIR="/Users/prad/Sonr/core/pkg/models"
PLUGIN_PROTO_DIR="/Users/prad/Sonr/plugin/lib/models"

# Platform Specific Parameters
IOS_ARTIFACT=$(IOS_BUILDDIR)/Core.framework
ANDROID_ARTIFACT=$(ANDROID_BUILDDIR)/io.sonr.core.aar
IOS_TARGET=ios/arm64
ANDROID_TARGET=android

# Gomobile Bind Commands
BUILD_IOS="cd bind && $(GOCLEAN) &&  $(GOBIND) -target=$(IOS_TARGET) -v -o $(IOS_ARTIFACT)"
BUILD_ANDROID="cd bind && $(GOCLEAN) && $(GOBIND) -target=$(ANDROID_TARGET) -v -o $(ANDROID_ARTIFACT)"

all: proto ios android 
	cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo "**************************************************************"
	@echo "************** FINISHED IOS/ANDROID BINDINGS *****************"
	@echo "**************************************************************"

ios:
	@echo ""
	@echo "***********************************************"
	@echo "************** BEGIN IOS BIND *****************"
	@echo "***********************************************"
	rm -rf $(IOS_BUILDDIR) 2>/dev/null
	mkdir -p $(IOS_BUILDDIR)
	eval $(BUILD_IOS)
	go mod tidy
	cd /System/Library/Sounds && afplay Glass.aiff
	@echo ""

android:
	@echo ""
	@echo "***************************************************"
	@echo "************** BEGIN ANDROID BIND *****************"
	@echo "***************************************************"
	rm -rf $(ANDROID_BUILDDIR) 2>/dev/null
	mkdir -p $(ANDROID_BUILDDIR)
	eval $(BUILD_ANDROID)
	go mod tidy
	cd /System/Library/Sounds && afplay Glass.aiff
	@echo ""

proto:
	cd protobuf && protoc -I=. --go_out=$(CORE_PROTO_DIR) ./models.proto
	cd protobuf && protoc -I=. --dart_out=$(PLUGIN_PROTO_DIR) ./models.proto

clean:
	cd bind && $(GOCLEAN)
	go mod tidy
	rm -rf $(IOS_BUILDDIR)
	rm -rf $(ANDROID_BUILDDIR)