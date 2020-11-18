SHELL := /bin/zsh
# Makefile that builds core and puts it into plugin repo
GOMOBILE=gomobile
GCMSG=git-commitmsg
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind
CORE_DIR=
IOS_BUILDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
IOS_ARTIFACT=$(IOS_BUILDDIR)/Core.framework
ANDROID_BUILDDIR=/Users/prad/Sonr/plugin/android/libs
ANDROID_ARTIFACT=$(ANDROID_BUILDDIR)/io.sonr.core.aar
IOS_TARGET=ios/arm64
ANDROID_TARGET=android

BUILD_IOS="cd bind && $(GOCLEAN) &&  $(GOBIND) -target=$(IOS_TARGET) -v -o $(IOS_ARTIFACT)"
BUILD_ANDROID="cd bind && $(GOCLEAN) && $(GOBIND) -target=$(ANDROID_TARGET) -v -o $(ANDROID_ARTIFACT)"

all: ios android 
	cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo "**************************************************************"
	@echo "************** FINISHED IOS/ANDROID BINDINGS *****************"
	@echo "**************************************************************"


proto:
	cd pkg/proto && protoc -I=. --go_out=. ./models.proto

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

clean:
	cd bind && $(GOCLEAN)
	go mod tidy
	rm -rf $(IOS_BUILDDIR)
	rm -rf $(ANDROID_BUILDDIR)