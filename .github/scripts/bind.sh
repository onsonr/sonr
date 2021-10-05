#!/bin/bash

# ----------------------------------------------------------------------
# ----- <Build Variables> ----------------------------------------------
SCRIPTDIR=$(dirname "$0")
cd ${SCRIPTDIR}/../../
PROJECT_DIR=$(pwd);
BIND_DIR=${PROJECT_DIR}/cmd/lib
BUILD_DIR=${PROJECT_DIR}/build
cd ${PROJECT_DIR}/../
ROOT_DIR=$(pwd);
PLUGIN_ANDROID=${ROOT_DIR}/mobile/android/libs
PLUGIN_IOS=${ROOT_DIR}/mobile/ios/Frameworks
ANDROID_ARTIFACT=io.sonr.core.aar
IOS_ARTIFACT=Core.xcframework
cd ${PROJECT_DIR}
platform=""
output=""
# ----- <Build Variables/> -----------------------------------------------
# ------------------------------------------------------------------------



# ------------------------------------------------------------------------
# ----- <Process Functions> ----------------------------------------------
# Functions for printing help
usage() {                                 # Function: Print a help message.
  echo "Usage: sonr-io/core/bind.sh [ -p PLATFORM (all, ios, android) ] [ -o OUTPUT (plugin, build) ]" 1>&2
  echo ""
}

# Function: Echo Android and init gomobile
init_android(){
  echo "- (2/3) Setup gomobile for Java AAR 🤖 [ >8min ]"
  go get -u golang.org/x/mobile/bind
  cd ${BIND_DIR}
  gomobile init
}

# Function: Echo iOS and init gomobile
init_ios() {
  echo "- (2/3) Setup gomobile for Objective-C Framework 📱 [ >8min ]"
  go get -u golang.org/x/mobile/bind
  cd ${BIND_DIR}
  gomobile init
}

# Function: Echo All and init gomobile
init_all() {
  echo "- (2/4) Setup gomobile for All Frameworks 🌎  [ >16min ]"
  go get -u golang.org/x/mobile/bind
  cd ${BIND_DIR}
  gomobile init
}

# Function: Exit with error.
exit_abnormal() {
  echo ""
  echo "🚨  ERROR: Invalid Parameters"
  usage
  go mod tidy
  exit 1
}

# Function: Exit after Success.
exit_success() {
    go mod tidy
    echo ""
    echo "-----------------------------"
    echo "🎉  SUCCESS: Bind complete"
    echo "-----------------------------"
    exit 0
}
# ----- <Process Functions/> ----------------------------------------------
# -------------------------------------------------------------------------



# ------------------------------------------------------------------------
# ----- <Input Flags> ----------------------------------------------------
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
# ----- <Input Flags/> ---------------------------------------------------
# ------------------------------------------------------------------------



# ------------------------------------------------------------------------
# ----- <Input Reader> ---------------------------------------------------
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
# ----- <Input Reader/> --------------------------------------------------
# ------------------------------------------------------------------------



# ------------------------------------------------------------------------
# ----- <Dir Setup> ------------------------------------------------------
if [ $output == plugin ]; then
    echo "- (1/3) Initialize Plugin Directories 📁 [ <1min ]"
    mkdir -p ${PLUGIN_ANDROID}
    mkdir -p ${PLUGIN_IOS}
elif [ $output == build ]; then
    echo "- (1/3) Initialize Build Directory 📁 [ <1min ]"
    mkdir -p ${BUILD_DIR}
else
    exit_abnormal
fi
# ----- <Dir Setup/> -----------------------------------------------------
# ------------------------------------------------------------------------



# ------------------------------------------------------------------------
# ----- <Build> ----------------------------------------------------------
echo ""
if [ $platform == android ]; then
    init_android
    echo ""
    echo "----------------- (AAR: Build Output) -------------------"
    if [ $output == plugin ]; then
        gtime -q -format=%E gomobile bind -ldflags='-s -w' -v -target=android -o ${PLUGIN_ANDROID}/${ANDROID_ARTIFACT}
        exit_success
    elif [ $output == build ]; then
        gtime -q -format=%E gomobile bind -ldflags='-s -w' -v -target=android -o ${BUILD_DIR}/${ANDROID_ARTIFACT}
        exit_success
    else
        exit_abnormal
    fi
elif [ $platform == ios ]; then
    init_ios
    echo ""
    echo "----------------- (xcframework: Build Output) -------------------"
    if [ $output == plugin ]; then
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios -bundleid=io.sonr.core -o ${PLUGIN_IOS}/${IOS_ARTIFACT}
        exit_success
    elif [ $output == build ]; then
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios -bundleid=io.sonr.core -o ${BUILD_DIR}/${IOS_ARTIFACT}
        exit_success
    else
        exit_abnormal
    fi
elif [ $platform == all ]; then
    init_all
    if [ $output == plugin ]; then
        echo "- (3/4) Binding iOS Framework"
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${PLUGIN_IOS}/${IOS_ARTIFACT}
        echo "- (4/4) Binding Android Framework"
        gtime -q -format=%E gomobile bind -ldflags='-s -w' -v -target=android -o ${PLUGIN_ANDROID}/${ANDROID_ARTIFACT}
        exit_success
    elif [ $output == build ]; then
        echo "- (3/4) Binding iOS Framework"
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${BUILD_DIR}/${IOS_ARTIFACT}
        echo "- (4/4) Binding Android Framework"
        gtime -q --format=%E gomobile bind -ldflags='-s -w' -v -target=ios/arm64 -bundleid=io.sonr.core -o ${BUILD_DIR}/${IOS_ARTIFACT}
        exit_success
    else
        exit_abnormal
    fi
else
    exit_abnormal
fi
# ----- <Build/> ---------------------------------------------------------
# ------------------------------------------------------------------------
