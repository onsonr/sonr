#!/bin/bash

SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
MOTOR_LIB_DIR=${PROJECT_DIR}/cmd/motor-lib
MOTOR_WASM_DIR=${PROJECT_DIR}/cmd/motor-wasm

while getopts "iawv:" opt; do
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
      cd ${MOTOR_LIB_DIR}
      gomobile bind -ldflags='-s -w' -target=android/arm64 -o ${ANDROID_ARTIFACT} -v
      ;;
    i)
      echo "🔷 Setting up build Environment..."
      IOS_BUILD_PATH="motorlib_darwin_${VERSION}_arm64"
      IOS_OUT=${PROJECT_DIR}/build/${IOS_BUILD_PATH}
      IOS_ARTIFACT=${IOS_OUT}/SonrMotor.xcframework
      mkdir -p ${IOS_OUT}

      echo "🔷 Binding iOS..."
      cd ${MOTOR_LIB_DIR}
      gomobile bind -ldflags='-s -w' -target=ios/arm64 -o ${IOS_ARTIFACT} -v
      ;;
    w)
      echo "🔷 Setting up build Environment..."
      WASM_BUILD_PATH="motorlib_js_${VERSION}_wasm"
      WASM_OUT=${PROJECT_DIR}/build/${WASM_BUILD_PATH}
      WASM_ARTIFACT=${WASM_OUT}/motorlib.wasm
      mkdir -p ${WASM_OUT}

      echo "🔷 Binding Web..."
      cd ${MOTOR_WASM_DIR}
      GOOS=js GOARCH=wasm go build -o ${WASM_ARTIFACT} -v
      ;;
    ?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done
