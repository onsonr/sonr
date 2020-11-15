# Makefile that builds core and puts it into plugin repo
GOMOBILE=gomobile
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
	cd $(CORE_DIR) && git add .
	git-commitmsg | pbcopy
	git commit -m `pbpaste`

git:
	cd /Users/prad/Sonr/core && git add .
	git-commitmsg | pbcopy
	git commit -m = `pbpaste`

ios:
	$(info ************** BEGIN IOS BIND *****************)
	rm -rf $(IOS_BUILDDIR) 2>/dev/null
	mkdir -p $(IOS_BUILDDIR)
	eval $(BUILD_IOS)
	go mod tidy
	cd /System/Library/Sounds && afplay Glass.aiff

android:
	$(info ************ BEGIN ANDROID BIND *******************)
	rm -rf $(ANDROID_BUILDDIR) 2>/dev/null
	mkdir -p $(ANDROID_BUILDDIR)
	eval $(BUILD_ANDROID)
	go mod tidy
	cd /System/Library/Sounds && afplay Glass.aiff

clean:
	cd bind && $(GOCLEAN)
	rm -rf $(IOS_BUILDDIR)
	rm -rf $(ANDROID_BUILDDIR)