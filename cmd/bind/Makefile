GO_CMD ?= go
GOMOBILE_CMD ?= gomobile

PROJECT_ROOT = github.com/sonr-io/core
MOBILE_PROJECT_ROOT = $(PROJECT_ROOT)/cmd/bind

MOBILE_BUILD_PACKAGES := $(MOBILE_PROJECT_ROOT)/...

MAKEFILE_DIR := ${CURDIR}
BUILD_DIR = $(MAKEFILE_DIR)/build
LDFLAGS ?= '-s -w'

ANDROID_BUILD_DIR = $(BUILD_DIR)/android
ANDROID_OUTPUT_FILE_NAME = io.sonr.core.aar
ifeq ($(ANDROID_TARGET),)
	ANDROID_TARGET = android/arm
endif
ANDROID_JAVA_PKG = io.sonr.core

IOS_BUILD_DIR = $(BUILD_DIR)/ios
IOS_OUTPUT_FILE_NAME = Core.xcframework
ifeq ($(IOS_TARGET),)
    IOS_TARGET = ios
endif

IS_MACOS = false
ifneq ($(OS),Windows_NT)
	OS := $(shell uname)
	ifeq ($(OS),Darwin)
		IS_MACOS = true
	endif
endif

help:
	@echo 'Commands:'
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null \
		| awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' \
		| sort \
		| egrep -v -e '^[^[:alnum:]]' -e '^$@$$' \
		| grep -v / \
		| sed 's/^/	$(HELP_MSG_PREFIX)make /'

.PHONY: all
all: clean dependencies unit-test bindings

.PHONY: clean
clean: clean-build

.PHONY: clean-build
clean-build: clean-android clean-ios
	@rm -Rf $(BUILD_DIR)

.PHONY: clean-android
clean-android:
	@rm -Rf $(ANDROID_BUILD_DIR)

.PHONY: clean-ios
clean-ios:
	@rm -Rf $(IOS_BUILD_DIR)

.PHONY: dependencies
dependencies:
	@echo '----------------------------------------'
	@echo '   Installing dependencies              '
	@echo '----------------------------------------'
	@GO111MODULE=off go get -u golang.org/x/mobile/cmd/...
	$(GOMOBILE_CMD) init -v

.PHONY: bindings
bindings: clean-build bindings-android bindings-ios

.PHONY: bindings-android
bindings-android:
	@echo '----------------------------------------'
	@echo '   gomobile: Creating Android bindings  '
	@echo '----------------------------------------'
	@mkdir -p $(ANDROID_BUILD_DIR)
	$(GOMOBILE_CMD) bind -v -ldflags $(LDFLAGS) -target=$(ANDROID_TARGET) -javapkg=$(ANDROID_JAVA_PKG) -o=$(ANDROID_BUILD_DIR)/$(ANDROID_OUTPUT_FILE_NAME) $(MOBILE_BUILD_PACKAGES)

.PHONY: bindings-ios
bindings-ios:
	@echo '----------------------------------------'
	@echo '   gomobile: Creating iOS bindings      '
	@echo '----------------------------------------'
ifeq ($(IS_MACOS),true)
	@mkdir -p $(IOS_BUILD_DIR)
	@CGO_LDFLAGS_ALLOW='-fembed-bitcode' $(GOMOBILE_CMD) bind -v -ldflags $(LDFLAGS) -target=$(IOS_TARGET) -o=$(IOS_BUILD_DIR)/$(IOS_OUTPUT_FILE_NAME) $(MOBILE_BUILD_PACKAGES)
else
	@echo 'Skipping iOS bindings. Only available for Darwin operating systems.'
endif

.PHONY: unit-test
unit-test:
	@echo '----------------------------------------'
	@echo '   Running unit tests                   '
	@echo '----------------------------------------'
	@go test $(MOBILE_BUILD_PACKAGES) -count=1 -race -timeout=10m -cover
