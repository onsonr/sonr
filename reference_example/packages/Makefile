MAKEFILE_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
MANIFEST_FILE = $(MAKEFILE_DIR)/Manifest.yml

UTILS_DIR = $(MAKEFILE_DIR)/utils
UTIL_MANIFEST_GET_BIN = $(UTILS_DIR)/manifest_get/manifest_get.sh
UTIL_MANIFEST_GET = $(UTILS_DIR)/manifest_get
UTIL_MAVEN_FORMAT = $(UTILS_DIR)/maven_format
UTIL_MAVEN_FORMAT_REQ = $(UTIL_MAVEN_FORMAT)/requirements.txt
UTIL_MAVEN_FORMAT_CORE_BIN = $(UTIL_MAVEN_FORMAT)/maven_format_core.py
UTIL_MAVEN_PUBLISH = $(UTILS_DIR)/maven_publish
UTIL_MAVEN_PUBLISH_REQ = $(UTIL_MAVEN_PUBLISH)/requirements.txt
UTIL_MAVEN_PUBLISH_CORE_BIN = $(UTIL_MAVEN_PUBLISH)/maven_publish_core.py
UTIL_COCOAPOD_FORMAT = $(UTILS_DIR)/cocoapod_format
UTIL_COCOAPOD_FORMAT_REQ = $(UTIL_COCOAPOD_FORMAT)/requirements.txt
UTIL_COCOAPOD_FORMAT_BRIDGE_BIN = $(UTIL_COCOAPOD_FORMAT)/cocoapod_format_bridge.py
UTIL_COCOAPOD_FORMAT_CORE_BIN = $(UTIL_COCOAPOD_FORMAT)/cocoapod_format_core.py
UTIL_COCOAPOD_PUBLISH = $(UTILS_DIR)/cocoapod_publish
UTIL_COCOAPOD_PUBLISH_REQ = $(UTIL_COCOAPOD_PUBLISH)/requirements.txt
UTIL_COCOAPOD_PUBLISH_BRIDGE_BIN = $(UTIL_COCOAPOD_PUBLISH)/cocoapod_publish_bridge.py
UTIL_COCOAPOD_PUBLISH_CORE_BIN = $(UTIL_COCOAPOD_PUBLISH)/cocoapod_publish_core.py
UTIL_BINTRAY_FORMAT = $(UTILS_DIR)/bintray_format
UTIL_BINTRAY_FORMAT_REQ = $(UTIL_BINTRAY_FORMAT)/requirements.txt
UTIL_BINTRAY_FORMAT_IOS_BIN = $(UTIL_BINTRAY_FORMAT)/bintray_format_ios.py
UTIL_BINTRAY_PUBLISH = $(UTILS_DIR)/bintray_publish
UTIL_BINTRAY_PUBLISH_REQ = $(UTIL_BINTRAY_PUBLISH)/requirements.txt
UTIL_BINTRAY_PUBLISH_ANDROID_BIN = $(UTIL_BINTRAY_PUBLISH)/bintray_publish_android.py
UTIL_BINTRAY_PUBLISH_IOS_BIN = $(UTIL_BINTRAY_PUBLISH)/bintray_publish_ios.py
BUILD_DIR = $(MAKEFILE_DIR)/build
PIP = pip3

MANIFEST_GET_FUNC=$(or $(shell $(UTIL_MANIFEST_GET_BIN) $(1)),$(error "Can't get <$(1)> from Manifest.yml"))
DEV_VERSION := 0.0.42-dev
VERSION := $(or $(GOMOBILE_IPFS_VERSION),$(DEV_VERSION))
ANDROID_GROUP_ID := $(shell echo $(call MANIFEST_GET_FUNC,global.group_id) | tr . /)
ANDROID_CORE_ARTIFACT_ID := $(call MANIFEST_GET_FUNC,go_core.android.artifact_id)
ANDROID_APP_FILENAME := $(call MANIFEST_GET_FUNC,android_demo_app.filename)
IOS_CORE_PACKAGE := $(call MANIFEST_GET_FUNC,go_core.ios.package)
IOS_APP_FILENAME := $(call MANIFEST_GET_FUNC,ios_demo_app.filename)

GO_DIR = $(MAKEFILE_DIR)/../go
GO_SRC = $(shell find $(GO_DIR) -name \*.go)
GO_MOD_FILES = $(GO_DIR)/go.mod $(GO_DIR)/go.sum

CORE_PACKAGE = github.com/ipfs-shipyard/gomobile-ipfs/go/bind/core
EXT_PACKAGE ?=
GOMOBILE = $(GOPATH)/bin/gomobile
GOMOBILE_OPT ?=
GOMOBILE_TARGET ?=

ANDROID_DIR = $(MAKEFILE_DIR)/../android
ANDROID_SRC = $(shell git ls-files $(ANDROID_DIR) | grep -v '.gitignore')
ANDROID_BUILD_DIR = $(BUILD_DIR)/android
ANDROID_BUILD_DIR_INT = $(ANDROID_BUILD_DIR)/intermediates
ANDROID_BUILD_DIR_INT_CORE = $(ANDROID_BUILD_DIR_INT)/core
ANDROID_CORE = $(ANDROID_BUILD_DIR_INT_CORE)/core.aar
ANDROID_BUILD_DIR_MAV = $(ANDROID_BUILD_DIR)/maven
ANDROID_BUILD_DIR_MAV_CORE = $(ANDROID_BUILD_DIR_MAV)/$(ANDROID_GROUP_ID)/$(ANDROID_CORE_ARTIFACT_ID)/$(VERSION)
ANDROID_OUTPUT_APK = $(ANDROID_DIR)/app/build/outputs/apk/release/app-release.apk
ANDROID_BUILD_DIR_APP = $(ANDROID_BUILD_DIR)/app/$(VERSION)
ANDROID_BUILD_DIR_APP_APK = $(ANDROID_BUILD_DIR_APP)/$(ANDROID_APP_FILENAME)-$(VERSION).apk

IOS_DIR = $(MAKEFILE_DIR)/../ios
IOS_SRC = $(shell git ls-files $(IOS_DIR) | grep -v '.gitignore')
IOS_BUILD_DIR = $(BUILD_DIR)/ios
IOS_BUILD_DIR_INT = $(IOS_BUILD_DIR)/intermediates
IOS_BUILD_DIR_INT_CORE = $(IOS_BUILD_DIR_INT)/core
IOS_CORE = $(IOS_BUILD_DIR_INT_CORE)/Core.framework
IOS_BUILD_DIR_CCP = $(IOS_BUILD_DIR)/cocoapods
IOS_BUILD_DIR_CCP_CORE = $(IOS_BUILD_DIR_CCP)/$(IOS_CORE_PACKAGE)/$(VERSION)
IOS_WORKSPACE = $(IOS_DIR)/Example.xcworkspace
IOS_APP_PLIST = $(IOS_WORKSPACE)/release_export.plist
IOS_BUILD_DIR_APP = $(IOS_BUILD_DIR)/app/$(VERSION)
IOS_BUILD_DIR_APP_IPA = $(IOS_BUILD_DIR_APP)/$(IOS_APP_FILENAME)-$(VERSION).ipa
IOS_BUILD_DIR_INT_APP = $(IOS_BUILD_DIR_INT)/app
IOS_BUILD_DIR_INT_APP_IPA = $(IOS_BUILD_DIR_INT_APP)/ipa
IOS_BUILD_DIR_INT_APP_IPA_OUTPUT = $(IOS_BUILD_DIR_INT_APP_IPA)/Example.ipa
IOS_BUILD_DIR_INT_APP_ARCHIVE = $(IOS_BUILD_DIR_INT_APP)/archive
IOS_BUILD_DIR_INT_APP_ARCHIVE_OUTPUT = $(IOS_BUILD_DIR_INT_APP_ARCHIVE)/app-release.xcarchive

DOC_DIR = $(MAKEFILE_DIR)/../docs
ANDROID_DOC_DIR = $(DOC_DIR)/android
IOS_DOC_DIR = $(DOC_DIR)/ios

.PHONY: help build_core build_core.android build_core.ios build_demo build_demo.android build_demo.ios clean clean.android clean.ios docgen docgen.android docgen.ios fail_on_dev publish publish_bridge publish_bridge.android publish_bridge.ios publish_core publish_core.android publish_core.ios publish_demo publish_demo.android publish_demo.ios re re.android re.ios test test_bridge test_bridge.android test_bridge.ios test_core

help:
	@echo 'Commands:'
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null \
		| awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' \
		| sort \
		| egrep -v -e '^[^[:alnum:]]' -e '^$@$$' \
		| grep -v / \
		| sed 's/^/	$(HELP_MSG_PREFIX)make /'

# Build rules
build_core: build_core.android build_core.ios

build_core.android: $(ANDROID_BUILD_DIR_MAV_CORE)

$(ANDROID_BUILD_DIR_MAV_CORE): $(ANDROID_CORE) $(MANIFEST_FILE)
	@echo '------------------------------------'
	@echo '   Android Core: Maven formatting   '
	@echo '------------------------------------'
	@if [ "$$($(PIP) freeze | grep -f $(UTIL_MAVEN_FORMAT_REQ) | wc -l)" != "$$(wc -l < $(UTIL_MAVEN_FORMAT_REQ))" ]; then \
		echo 'Installing pip dependencies:'; $(PIP) install -r $(UTIL_MAVEN_FORMAT_REQ); echo; \
	fi
	@$(UTIL_MAVEN_FORMAT_CORE_BIN) && touch $(ANDROID_BUILD_DIR_MAV_CORE)
	@echo 'Done!'

$(ANDROID_CORE): $(ANDROID_BUILD_DIR_INT_CORE) $(GO_SRC) $(GO_MOD_FILES) | $(GOMOBILE)
	@echo '------------------------------------'
	@echo '   Android Core: Gomobile binding   '
	@echo '------------------------------------'
	@$(GOMOBILE) bind -v $(GOMOBILE_OPT) -target=android$(GOMOBILE_TARGET) -o $(ANDROID_CORE) $(CORE_PACKAGE) $(EXT_PACKAGE)
	@echo 'Done!'

$(ANDROID_BUILD_DIR_INT_CORE):
	@mkdir -p $(ANDROID_BUILD_DIR_INT_CORE)

build_core.ios: $(IOS_BUILD_DIR_CCP_CORE)

$(IOS_BUILD_DIR_CCP_CORE): $(IOS_CORE) $(MANIFEST_FILE)
	@echo '------------------------------------'
	@echo '   iOS Core: CocoaPod formatting   '
	@echo '------------------------------------'
	@if [ "$$($(PIP) freeze | grep -f $(UTIL_COCOAPOD_FORMAT_REQ) | wc -l)" != "$$(wc -l < $(UTIL_COCOAPOD_FORMAT_REQ))" ]; then \
		echo 'Installing pip dependencies:'; $(PIP) install -r $(UTIL_COCOAPOD_FORMAT_REQ); echo; \
	fi
	@$(UTIL_COCOAPOD_FORMAT_CORE_BIN) && touch $(IOS_BUILD_DIR_CCP_CORE)
	@echo 'Done!'

$(IOS_CORE): $(IOS_BUILD_DIR_INT_CORE) $(GO_SRC) $(GO_MOD_FILES) | $(GOMOBILE)
	@echo '------------------------------------'
	@echo '     iOS Core: Gomobile binding     '
	@echo '------------------------------------'
	@$(GOMOBILE) bind $(GOMOBILE_OPT) -target=ios$(GOMOBILE_TARGET) -o $(IOS_CORE) $(CORE_PACKAGE) $(EXT_PACKAGE)
	@echo 'Done!'

$(IOS_BUILD_DIR_INT_CORE):
	@mkdir -p $(IOS_BUILD_DIR_INT_CORE)

build_demo: build_demo.android build_demo.ios

build_demo.android: $(ANDROID_BUILD_DIR_APP_APK)

$(ANDROID_BUILD_DIR_APP_APK): $(ANDROID_BUILD_DIR_APP) $(ANDROID_OUTPUT_APK) $(MANIFEST)
	@echo '------------------------------------'
	@echo '  Android Demo: apk path creation   '
	@echo '------------------------------------'
	@cp $(ANDROID_OUTPUT_APK) $(ANDROID_BUILD_DIR_APP_APK)
	@echo 'Built .apk available in: $(ANDROID_BUILD_DIR_APP)'
	@echo 'Done!'

$(ANDROID_OUTPUT_APK): $(ANDROID_SRC) $(MANIFEST) $(ANDROID_BUILD_DIR_MAV_CORE)
	@echo '------------------------------------'
	@echo '   Android Demo: Gradle building    '
	@echo '------------------------------------'
	@cd $(ANDROID_DIR) && ./gradlew app:build
	@touch $(ANDROID_OUTPUT_APK)
	@echo 'Done!'

$(ANDROID_BUILD_DIR_APP):
	@mkdir -p $(ANDROID_BUILD_DIR_APP)

build_demo.ios: $(IOS_BUILD_DIR_APP_IPA)

$(IOS_BUILD_DIR_APP_IPA): $(IOS_BUILD_DIR_INT_APP_IPA_OUTPUT) $(MANIFEST)
	@echo '------------------------------------'
	@echo '    iOS Demo: Bintray formatting    '
	@echo '------------------------------------'
	@if [ "$$($(PIP) freeze | grep -f $(UTIL_BINTRAY_FORMAT_REQ) | wc -l)" != "$$(wc -l < $(UTIL_BINTRAY_FORMAT_REQ))" ]; then \
		echo 'Installing pip dependencies:'; $(PIP) install -r $(UTIL_BINTRAY_FORMAT_REQ); echo; \
	fi
	@$(UTIL_BINTRAY_FORMAT_IOS_BIN)
	@echo 'Done!'

$(IOS_BUILD_DIR_INT_APP_IPA_OUTPUT): $(IOS_BUILD_DIR_INT_APP_IPA) $(IOS_BUILD_DIR_INT_APP_ARCHIVE_OUTPUT)
	@echo '------------------------------------'
	@echo '   iOS Demo: XCode building ipa     '
	@echo '------------------------------------'
	@xcodebuild -exportArchive \
		-archivePath $(IOS_BUILD_DIR_INT_APP_ARCHIVE_OUTPUT) \
		-exportOptionsPlist $(IOS_APP_PLIST) \
		-exportPath $(IOS_BUILD_DIR_INT_APP_IPA)
	@touch $(IOS_BUILD_DIR_INT_APP_IPA_OUTPUT)
	@echo 'Done!'

$(IOS_BUILD_DIR_INT_APP_IPA):
	@mkdir -p $(IOS_BUILD_DIR_INT_APP_IPA)

$(IOS_BUILD_DIR_INT_APP_ARCHIVE_OUTPUT): $(IOS_BUILD_DIR_INT_APP_ARCHIVE) $(IOS_CORE) $(IOS_SRC)
	@echo '------------------------------------'
	@echo '  iOS Demo: XCode building archive  '
	@echo '------------------------------------'
	@xcodebuild archive \
		-workspace $(IOS_WORKSPACE) \
		-scheme Example \
		-configuration Release \
		-sdk iphoneos \
		-archivePath $(IOS_BUILD_DIR_INT_APP_ARCHIVE_OUTPUT)
	@touch $(IOS_BUILD_DIR_INT_APP_ARCHIVE_OUTPUT)
	@echo 'Done!'

$(IOS_BUILD_DIR_INT_APP_ARCHIVE):
	@mkdir -p $(IOS_BUILD_DIR_INT_APP_ARCHIVE)

# Publish rules
publish: publish_core publish_bridge publish_demo

publish_core: publish_core.android publish_core.ios

publish_core.android: fail_on_dev build_core.android
	@echo '------------------------------------'
	@echo '   Android Core: Maven publishing   '
	@echo '------------------------------------'
	@if [ "$$($(PIP) freeze | grep -f $(UTIL_MAVEN_PUBLISH_REQ) | wc -l)" != "$$(wc -l < $(UTIL_MAVEN_PUBLISH_REQ))" ]; then \
		echo 'Installing pip dependencies:'; $(PIP) install -r $(UTIL_MAVEN_PUBLISH_REQ); echo; \
	fi
	@$(UTIL_MAVEN_PUBLISH_CORE_BIN)
	@echo 'Done!'

publish_core.ios: fail_on_dev build_core.ios
	@echo '------------------------------------'
	@echo '   iOS Core: CocoaPod publishing   '
	@echo '------------------------------------'
	@if [ "$$($(PIP) freeze | grep -f $(UTIL_COCOAPOD_PUBLISH_REQ) | wc -l)" != "$$(wc -l < $(UTIL_COCOAPOD_PUBLISH_REQ))" ]; then \
		echo 'Installing pip dependencies:'; $(PIP) install -r $(UTIL_COCOAPOD_PUBLISH_REQ); echo; \
	fi
	@$(UTIL_COCOAPOD_PUBLISH_CORE_BIN)
	@echo 'Done!'

publish_bridge: publish_bridge.android publish_bridge.ios

publish_bridge.android: fail_on_dev build_core.android
	@echo '------------------------------------'
	@echo '  Android Bridge: Maven publishing  '
	@echo '------------------------------------'
	@cd $(ANDROID_DIR) && ./gradlew bridge:bintrayUpload
	@echo 'Done!'

publish_bridge.ios: fail_on_dev build_core.ios
	@echo '------------------------------------'
	@echo '  iOS Bridge: CocoaPod publishing   '
	@echo '------------------------------------'
	@if [ "$$($(PIP) freeze | grep -f $(UTIL_COCOAPOD_FORMAT_REQ) -f $(UTIL_COCOAPOD_PUBLISH_REQ) | wc -l)" != "$$(cat $(UTIL_COCOAPOD_FORMAT_REQ) $(UTIL_COCOAPOD_PUBLISH_REQ) | sort | uniq | wc -l )" ]; then \
		echo 'Installing pip dependencies:'; $(PIP) install -r $(UTIL_COCOAPOD_FORMAT_REQ) -r $(UTIL_COCOAPOD_PUBLISH_REQ); echo; \
	fi
	@$(UTIL_COCOAPOD_FORMAT_BRIDGE_BIN) && $(UTIL_COCOAPOD_PUBLISH_BRIDGE_BIN)
	@echo 'Done!'

publish_demo: publish_demo.android publish_demo.ios

publish_demo.android: fail_on_dev build_demo.android
	@echo '------------------------------------'
	@echo '  Android Demo: Bintray publishing  '
	@echo '------------------------------------'
	@if [ "$$($(PIP) freeze | grep -f $(UTIL_BINTRAY_PUBLISH_REQ) | wc -l)" != "$$(wc -l < $(UTIL_BINTRAY_PUBLISH_REQ))" ]; then \
		echo 'Installing pip dependencies:'; $(PIP) install -r $(UTIL_BINTRAY_PUBLISH_REQ); echo; \
	fi
	$(UTIL_BINTRAY_PUBLISH_ANDROID_BIN)
	@echo 'Done!'

publish_demo.ios: fail_on_dev build_demo.ios
	@echo '------------------------------------'
	@echo '    iOS Demo: Bintray publishing    '
	@echo '------------------------------------'
	@if [ "$$($(PIP) freeze | grep -f $(UTIL_BINTRAY_PUBLISH_REQ) | wc -l)" != "$$(wc -l < $(UTIL_BINTRAY_PUBLISH_REQ))" ]; then \
		echo 'Installing pip dependencies:'; $(PIP) install -r $(UTIL_BINTRAY_PUBLISH_REQ); echo; \
	fi
	$(UTIL_BINTRAY_PUBLISH_IOS_BIN)
	@echo 'Done!'

# Doc generation rules
docgen: docgen.android docgen.ios

docgen.android: $(ANDROID_DOC_DIR) build_core.android
	@echo '------------------------------------'
	@echo '   Android Bridge: Doc generation   '
	@echo '------------------------------------'
	@cd $(ANDROID_DIR) && ./gradlew bridge:javadoc
	@cp -rf $(ANDROID_DIR)/bridge/javadoc/* $(ANDROID_DOC_DIR)
	@echo 'Done!'

$(ANDROID_DOC_DIR):
	@mkdir -p $(ANDROID_DOC_DIR)

docgen.ios: $(IOS_DOC_DIR) build_core.ios
	@echo '------------------------------------'
	@echo '     iOS Bridge: Doc generation     '
	@echo '------------------------------------'
	@cd $(IOS_DIR)/Bridge && \
		jazzy -o $(IOS_DOC_DIR) \
		--readme $(IOS_DIR)/../README.md \
		--module 'GomobileIPFS' \
		--title 'Gomobile-IPFS - iOS Bridge' \
		--github_url 'https://github.com/ipfs-shipyard/gomobile-ipfs' \
		--github-file-prefix 'https://github.com/ipfs-shipyard/gomobile-ipfs/tree/master/ios/Bridge'
	@echo 'Done!'

$(IOS_DOC_DIR):
	@mkdir -p $(IOS_DOC_DIR)

# Test rules
test: test_core test_bridge

test_bridge: test_bridge.android test_bridge.ios

test_bridge.android: build_core.android
	@echo '------------------------------------'
	@echo '   Android Bridge: running test     '
	@echo '------------------------------------'
	@cd $(ANDROID_DIR) && ./gradlew bridge:test && \
	EMULATOR=$$(emulator -avd -list-avds | tail -n1); \
	if [ -z "$$EMULATOR" ]; then \
		>&2 echo "No emulator found to run the test";	\
		exit 1;	\
	fi;	\
	emulator -avd $$EMULATOR -no-boot-anim -no-window -no-snapshot-save -gpu swiftshader_indirect -noaudio & EMULATOR_PID=$$!; \
	adb wait-for-device shell 'while [[ -z $$(getprop sys.boot_completed) ]]; do sleep 1; done;'; \
	(cd $(ANDROID_DIR) && ./gradlew bridge:connectedAndroidTest) || \
	(kill $$EMULATOR_PID; exit 1) && \
	(kill $$EMULATOR_PID; echo 'Done!')

test_bridge.ios: build_core.ios
	@echo '------------------------------------'
	@echo '     iOS Bridge: running test       '
	@echo '------------------------------------'
	@DESTINATION=$$(xcodebuild -showdestinations -project $(IOS_DIR)/Bridge/GomobileIPFS.xcodeproj -scheme GomobileIPFS | awk '/Ineligible destinations for/ {exit} {print}' | grep 'platform:iOS Simulator' | awk -F 'id:' '{print $$2}' | cut -d',' -f1 | tail -n1); \
	if [ -z "$$DESTINATION" ]; then \
		>&2 echo "No compatible simulator found to run the test";	\
		exit 1;	\
	fi;	\
	xcodebuild test -project $(IOS_DIR)/Bridge/GomobileIPFS.xcodeproj -scheme GomobileIPFS -sdk iphonesimulator -destination "platform=iOS Simulator,id=$$DESTINATION"
	@echo 'Done!'

test_core:
	@echo '------------------------------------'
	@echo '       Go Core: running test        '
	@echo '------------------------------------'
	@cd $(GO_DIR) && go test -v ./...
	@echo 'Done!'

# Misc rules
$(GOMOBILE):
	@echo '------------------------------------'
	@echo '      Go: installing modules        '
	@echo '------------------------------------'
	go mod download
	go install golang.org/x/mobile/cmd/gomobile
	go install golang.org/x/mobile/cmd/gobind
	@$(GOMOBILE) init -v
	@echo 'Done!'

fail_on_dev:
	@if [ "$(VERSION)" == "$(DEV_VERSION)" ]; then \
		>&2 echo "Can't publish a dev version: GOMOBILE_IPFS_VERSION env variable not set";	\
		exit 1; \
	fi

# Clean rules
clean: clean.android clean.ios

clean.android:
	@echo '------------------------------------'
	@echo '  Android Core: removing build dir  '
	@echo '------------------------------------'
	@rm -rf $(ANDROID_BUILD_DIR)
	@echo 'Done!'

clean.ios:
	@echo '------------------------------------'
	@echo '    iOS Core: removing build dir    '
	@echo '------------------------------------'
	@rm -rf $(IOS_BUILD_DIR)
	@echo 'Done!'
