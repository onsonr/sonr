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

platform=""
output=""

# Functions for printing help
usage() {                                 # Function: Print a help message.
  echo ""
  echo "🚨  ERROR: Invalid Parameters"
  echo "Usage: sonr-io/core/bind.sh [ -p PLATFORM (ios, android) ] [ -o OUTPUT (plugin, build) ]" 1>&2
  echo ""
}
exit_abnormal() {                         # Function: Exit with error.
  usage
  exit 1
}
cd ${PROJECT_DIR}

# Parse arguments
while getopts ":p:o:" options; do         # Loop: Get the next option;
  case "${options}" in                    #
    p)                                    # If the option is n,
      platform=${OPTARG}                      # set $NAME to specified value.
      ;;
    o)                                    # If the option is t,
      output=${OPTARG}                     # Set $TIMES to specified value.
      ;;
    :)                                    # If expected argument omitted:
      echo "Error: -${OPTARG} requires an argument."
      exit_abnormal                       # Exit abnormally.
      ;;
    *)                                    # If unknown (any other) option:
      exit_abnormal                       # Exit abnormally.
      ;;
  esac
done

echo ""
echo "⚡️ Sonr Bind Script for Mobile"
echo ""
if [ "${platform}" == "" ]; then
    echo "1. Choose Platform (ios, android)"
    read -p "Platform: " platform
else
    echo " └─ Building for [ ${platform} ] 🔥 "
fi
if [ "${output}" == "" ]; then
    echo "2. Choose Output Path (plugin, build)"
    read -p "Output Path: " output
else
    echo " └─ Output to [ ${output} ] 📩 "
fi
echo ""

echo "- (1/4) Installing Dependencies [ ~2min ]"
go get -u golang.org/x/mobile/bind
if [ $platform == android ]; then
    echo "- (2/4) Binding Java AAR 🤖 [ >8min ]"
    cd ${CORE_BIND_DIR}
    gomobile init
elif [ $platform == ios ]; then
    echo "- (2/4) Binding Objective-C Framework 📱 [ >8min ]"
    cd ${CORE_BIND_DIR}
    gomobile init
elif [ $platform == all ]; then
    echo "- (2/4) Binding All Frameworks 🌎  [ >16min ]"
    cd ${CORE_BIND_DIR}
    gomobile init
else
    exit_abnormal
fi

echo ""
echo "----------------- (Build Output) -------------------"
if [ $platform == android ]; then
    if [ $output == plugin ]; then
        gtime -q gomobile bind -ldflags='-s -w' -v -target=android -o ${PLUGIN_ANDROID}/${ANDROID_ARTIFACT}
        echo ""
    elif [ $output == build ]; then
        mkdir -p ${PROJECT_DIR}/build
        gtime -q gomobile bind -ldflags='-s -w' -v -target=android -o ${PROJECT_DIR}/build/${ANDROID_ARTIFACT}
        echo ""
    else
        exit_abnormal
    fi
elif [ $platform == ios ]; then
    if [ $output == plugin ]; then
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${PLUGIN_IOS}/${IOS_ARTIFACT}
        echo ""
    elif [ $output == build ]; then
        mkdir -p ${PROJECT_DIR}/build
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${PROJECT_DIR}/build/${IOS_ARTIFACT}
        echo ""
    else
        exit_abnormal
    fi
elif [ $platform == all ]; then
    if [ $output == plugin ]; then
        echo "Building iOS (3/4)"
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${PLUGIN_IOS}/${IOS_ARTIFACT}
        echo ""
        echo "Building Android (4/4)"
        gtime -q gomobile bind -ldflags='-s -w' -v -target=android -o ${PLUGIN_ANDROID}/${ANDROID_ARTIFACT}
        echo ""
    elif [ $output == build ]; then
        mkdir -p ${PROJECT_DIR}/build
        echo "Building iOS (3/4)"
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${PROJECT_DIR}/build/${IOS_ARTIFACT}
        echo ""
        echo "Building Android (4/4)"
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${PROJECT_DIR}/build/${IOS_ARTIFACT}
        echo ""
    else
        exit_abnormal
    fi
else
    exit_abnormal
fi
echo ""
go mod tidy

if [ $platform == android ]; then
    echo "✅  Finished Binding for Android"
elif [ $platform == ios ]; then
    echo "✅  Finished Binding for iOS"
else
    echo "✅  Finished Binding for All Platforms"
fi
exit 0
