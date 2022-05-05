#!/bin/bash

SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
MOTOR_DIR=${PROJECT_DIR}/cmd/motor-lib

while getopts "iav:" opt; do
  case $opt in
    v)
      echo "🔷 Binding with Version: $OPTARG" >&2
      # Set Environment Variables
      VERSION=$OPTARG
      ;;
    a)
      echo "🔷 Setting up build Environment..."
      ANDROID_BUILD_PATH="motorlib_android_${VERSION}_arm64"
      ANDROID_OUT=${PROJECT_DIR}/build/${ANDROID_BUILD_PATH}
      ANDROID_ARTIFACT=${ANDROID_OUT}/io.sonr.motor.aar
      mkdir -p ${ANDROID_OUT}

      echo "🔷 Binding Android..."
      gomobile bind -ldflags='-s -w' -target=android/arm64 -o ${ANDROID_ARTIFACT} -v
      ;;
    i)
      echo "🔷 Setting up build Environment..."
      IOS_BUILD_PATH="motorlib_darwin_${VERSION}_arm64"
      IOS_OUT=${PROJECT_DIR}/build/${IOS_BUILD_PATH}
      IOS_ARTIFACT=${IOS_OUT}/SonrMotor.xcframework
      mkdir -p ${IOS_OUT}

      echo "🔷 Binding iOS..."
      gomobile bind -ldflags='-s -w' -target=ios/arm64 -o ${IOS_ARTIFACT} -v
      ;;
    ?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done
