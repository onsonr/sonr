#!/bin/bash
echo "🛠  Building iOS Framework..."
SCRIPTDIR=$(dirname "$0")
cd ${SCRIPTDIR}/../../

echo "- (1/4) Setting up Project [ <1min ]"
PROJECT_DIR=$(pwd);
cd ${PROJECT_DIR}/../
ROOT_DIR=$(pwd);
cd ${PROJECT_DIR}
CORE_BIND_DIR=${PROJECT_DIR}/cmd/lib
BIND_DIR_IOS=${ROOT_DIR}/plugin/ios/Frameworks
BIND_IOS_ARTIFACT=${BIND_DIR_IOS}/Core.xcframework

echo "- (2/4) Installing Dependencies [ ~2min ]"
go install github.com/joho/godotenv/cmd/godotenv@latest
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init

echo "- (3/4) Binding Objective-C Framework [ >8min ]"
cd ${CORE_BIND_DIR}
echo ""
echo "----------------- (Build Output) -------------------"
gtime gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${BIND_IOS_ARTIFACT}
echo "✅  Finished Binding for iOS ➡ `date`"
