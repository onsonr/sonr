SHELL=/bin/zsh # Set Shell

# @ Go Commands
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind

# @ Bind Directories
IOS_BINDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
IOS_ARTIFACT= $(IOS_BINDDIR)/Core.framework
ANDROID_BINDDIR=/Users/prad/Sonr/plugin/android/libs
ANDROID_ARTIFACT= $(ANDROID_BINDDIR)/io.sonr.core.aar

# @ Build Directories
MAC_BUILDDIR=/Users/prad/Sonr/core/build/darwin
MAC_ARTIFACT=$(MAC_BUILDDIR)/Sonr.app/Contents/MacOS/sonr_core
WIN_BUILDDIR=/Users/prad/Sonr/core/build/win
WIN_ARTIFACT=$(WIN_BUILDDIR)/sonr-core.exe

# @ Proto Directories
PB_PATH="/Users/prad/Sonr/core/internal/models"
CONTACT_PB_DIR="/Users/prad/Sonr/contact/lib/src/data/models"
CORE_PB_DIR="/Users/prad/Sonr/core/internal/models"
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
	cd bind && gomobile bind -ldflags='-s -w' -target=android -v -o $(ANDROID_ARTIFACT)
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
	cd bind && gomobile bind -ldflags='-s -w' -target=ios -v -o $(IOS_ARTIFACT)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 COMPLETE IOS BIND 📱 ------------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

## build       :   Builds Darwin and Windows Builds at Build Path
build: proto build.darwin build.win
	@go mod tidy
	@cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo ""
	@echo "------------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅   FINISHED DESKTOP BUILD  ✅ ✅ ✅  --------------"
	@echo "------------------------------------------------------------------"

## └─ darwin        - MacOS executable
build.darwin:
	@echo ""
	@echo ""
	@echo "-----------------------------------------------------------"
	@echo "------------- 🖥  START DARWIN BUILD  🖥  -------------------"
	@echo "-----------------------------------------------------------"
	@go clean -cache
	@go mod tidy
	cd pkg && packr build -o $(MAC_ARTIFACT)
	@packr clean
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Building ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🖥  COMPLETED DAWIN BULD  🖥  -------------------"
	@echo "--------------------------------------------------------------"

## └─ win           - Windows executable
build.win:
	@echo ""
	@echo ""
	@echo "-----------------------------------------------------------"
	@echo "------------- 🪟 START WINDOWS BUILD 🪟 --------------------"
	@echo "-----------------------------------------------------------"
	@go clean -cache
	cd pkg && GOOS=windows GOARCH=amd64 packr build -ldflags -H=windowsgui -o $(WIN_ARTIFACT)
	@packr clean
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Building ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🪟 COMPLETED WINDOWS BULD 🪟 --------------------"
	@echo "--------------------------------------------------------------"
	@echo ""


## deploy      :   Package into Desktop Installers
deploy: proto deploy.mac
	@go mod tidy
	@cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo ""
	@echo "-------------------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅   FINISHED PACKAGING INSTALLERS  ✅ ✅ ✅  --------------"
	@echo "-------------------------------------------------------------------------"

## └─ mac           - MacOS DMG
# https://github.com/create-dmg/create-dmg
deploy.mac: build.darwin
	rm $(MAC_BUILDDIR)/Sonr-Installer.dmg
	create-dmg \
  --volname "Sonr Installer" \
  --volicon $(MAC_BUILDDIR)/meta/"volume.icns" \
  --background $(MAC_BUILDDIR)/meta/"volume-bg-alt.png" \
  --window-pos 200 120 \
  --window-size 800 400 \
  --icon-size 125 \
  --icon "Sonr.app" 182 172 \
  --hide-extension "Sonr.app" \
  --app-drop-link 618 167 \
  $(MAC_BUILDDIR)"/Sonr-Installer.dmg" \
  $(MAC_BUILDDIR)"/Sonr.app"
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "--------------------------------------------------------------"
	@echo "------------- 🖥  Packaged MacOS  🖥  -------------------------"
	@echo "--------------------------------------------------------------"

##
## [proto]     :   Compiles Protobuf models for Core Library and Plugin
proto:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 START PROTOBUFS COMPILE 🛸 -------------------"
	@echo "--------------------------------------------------------------"
	@cd internal/models && protoc --doc_out=$(PROTO_DOC_OUT) --doc_opt=html,index.html api.proto data.proto core.proto user.proto
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_CORE) api.proto data.proto core.proto user.proto
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_CONTACT) api.proto data.proto user.proto
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_PLUGIN) user.proto
	@echo "Finished Compiling ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 COMPILED ALL PROTOBUFS 🛸 --------------------"
	@echo "--------------------------------------------------------------"
	@echo ""


## [run]       :   Builds and Runs for Darwin
run:
	@echo ""
	@echo ""
	@echo "-----------------------------------------------------------"
	@echo "------------- 🖥  START DARWIN BUILD  🖥  -------------------"
	@echo "-----------------------------------------------------------"
	@go clean -cache
	@go mod tidy
	cd pkg && packr build -o $(MAC_ARTIFACT)
	@packr clean
	@echo "Finished Building ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🖥  RUN DAWIN BULD  🖥  -------------------------"
	@echo "--------------------------------------------------------------"
	@echo ""
	@cd $(MAC_BUILDDIR) && ./sonr_core

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
