SHELL := /bin/zsh # Set Shell

# GoMobile Commands
#GODEBUG=asyncpreemptoff=1
GCMSG=git-commitmsg
GOBIND=gomobile bind

# GoMobile Directories
IOS_BUILDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
ANDROID_BUILDDIR=/Users/prad/Sonr/plugin/android/libs

# Platform Specific Parameters
IOS_ARTIFACT=$(IOS_BUILDDIR)/Core.framework
IOS_TARGET=ios
ANDROID_ARTIFACT=$(ANDROID_BUILDDIR)/io.sonr.core.aar
ANDROID_TARGET=android/arm64
LD_FLAGS='-s -w'
 
# Proto Directories
PB_PATH="/Users/prad/Sonr/core/internal/models"
CORE_PB_DIR="/Users/prad/Sonr/core/internal/models"
PLUGIN_PB_DIR="/Users/prad/Sonr/plugin/lib/models"

# Proto Build Commands
PB_CLEAN_CORE="cd $(PB_PATH) && find ./ -name "*.go" -not -name "*.proto" -exec rm {} \;"
PB_CLEAN_PLUGIN="cd $(PLUGIN_PB_DIR) && find ./ -name "*.dart" -not -name "models.dart" -exec rm {} \;"
PB_BUILD_CORE="--go_out=$(CORE_PB_DIR)"
PB_BUILD_PLUGIN="--dart_out=$(PLUGIN_PB_DIR)"

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

clean:
	cd bind && $(GOCLEAN)
	go mod tidy
	rm -rf $(IOS_BUILDDIR)
	rm -rf $(ANDROID_BUILDDIR)
	eval $(PB_CLEAN_CORE) 2>/dev/null
	eval $(PB_CLEAN_PLUGIN) 