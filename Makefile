# Makefile that builds core and puts it into plugin repo
<<<<<<< HEAD
GOMOBILE=/Users/prad/go/bin/gomobile
GOCLEAN=$(GOMOBILE) clean
=======
GOMOBILE=gomobile
>>>>>>> parent of 88a087f... Merge pull request #20 from sonr-io/feature/events
GOBIND=$(GOMOBILE) bind
IOS_BUILDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
IOS_ARTIFACT=$(IOS_BUILDDIR)/Core.framework
ANDROID_BUILDDIR=/Users/prad/Sonr/plugin/android/libs
ANDROID_ARTIFACT=$(ANDROID_BUILDDIR)/io.sonr.core.aar
IOS_TARGET=ios
ANDROID_TARGET=android
LDFLAGS='-s -w'
IMPORT_PATH=github.com/sonr-io/core/bind

BUILD_IOS="cd bind && $(GOBIND) -target=$(IOS_TARGET) -v -o $(IOS_ARTIFACT)"
BUILD_ANDROID="cd bind && $(GOBIND) -target=$(ANDROID_TARGET) -v -o $(ANDROID_ARTIFACT)"

all: ios android 
	cd /System/Library/Sounds && afplay Hero.aiff

ios:
	$(info ************** BEGIN IOS BIND *****************)
	rm -rf $(IOS_BUILDDIR) 2>/dev/null
	mkdir -p $(IOS_BUILDDIR)
	eval $(BUILD_IOS)
<<<<<<< HEAD
	cd $(EXAMPLE_DIR) && $(FLUTCLEAN)
	cd /System/Library/Sounds && afplay Glass.aiff
=======
>>>>>>> parent of 88a087f... Merge pull request #20 from sonr-io/feature/events

android:
	$(info ************ BEGIN ANDROID BIND *******************)
	rm -rf $(ANDROID_BUILDDIR) 2>/dev/null
	mkdir -p $(ANDROID_BUILDDIR)
	eval $(BUILD_ANDROID)
	cd /System/Library/Sounds && afplay Glass.aiff

clean:
	rm -rf $(IOS_BUILDDIR)
	rm -rf $(ANDROID_BUILDDIR) 