#!/bin/bash

SCRIPTDIR=$(dirname "$0")
cd ${SCRIPTDIR}/../../
PROJECT_DIR=$(pwd);
CORE_BIND_DIR=${PROJECT_DIR}/cmd/lib
cd ${PROJECT_DIR}/../
ROOT_DIR=$(pwd);
PLUGIN_ANDROID=${ROOT_DIR}/plugin/android/libs
PLUGIN_IOS=${ROOT_DIR}/plugin/ios/Frameworks
ANDROID_ARTIFACT=io.sonr.core.aar
IOS_ARTIFACT=Core.xcframework
cd ${PROJECT_DIR}

echo ""
echo "⚡️ Sonr Bind Script for Mobile"
echo ""
echo "1. Choose Platform (ios, android)"
read -p "Platform: " platform
echo ""
echo "2. Choose Output Path (plugin, build)"
read -p "Output Path: " output

echo "- (1/3) Installing Dependencies [ ~2min ]"
go get -u golang.org/x/mobile/bind

if [ $platform == android ]; then
    echo "- (2/3) Binding Java AAR 🤖 [ >8min ]"
    cd ${CORE_BIND_DIR}
    gomobile init
elif [ $platform == ios ]; then
    echo "- (2/3) Binding Objective-C Framework 📱 [ >8min ]"
    cd ${CORE_BIND_DIR}
    gomobile init
else
    echo "Invalid platform"
    exit 1
fi

echo ""
echo "----------------- (Build Output) -------------------"
if [ $platform == android ]; then
    if [ $output == plugin ]; then
        gtime gomobile bind -ldflags='-s -w' -v -target=android -o ${PLUGIN_ANDROID}/${ANDROID_ARTIFACT}
        echo ""
    elif [ $output == build ]; then
        mkdir -p ${PROJECT_DIR}/build
        gtime gomobile bind -ldflags='-s -w' -v -target=android -o ${PROJECT_DIR}/build/${ANDROID_ARTIFACT}
        echo ""
    else
        echo "Invalid output"
        exit 1
    fi
elif [ $platform == ios ]; then
    if [ $output == plugin ]; then
        gtime gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${PLUGIN_IOS}/${IOS_ARTIFACT}
        echo ""
    elif [ $output == build ]; then
        mkdir -p ${PROJECT_DIR}/build
        gtime gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${PROJECT_DIR}/build/${IOS_ARTIFACT}
        echo ""
    else
        echo "Invalid output"
        exit 1
    fi
else
    exit 1
fi

go mod tidy

if [ $platform == android ]; then
    echo "✅  Finished Binding for Android ➡ `date`"
elif [ $platform == ios ]; then
    echo "✅  Finished Binding for iOS ➡ `date`"
else
    exit 1
fi

