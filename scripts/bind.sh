#!/bin/bash

SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
MOTOR_LIB_DIR=${PROJECT_DIR}/cmd/motor-lib

while getopts "iaw" opt; do
  case $opt in
    a)
      echo "🔷 Setting up build Environment..."
      ANDROID_OUT=${PROJECT_DIR}/build/android
      ANDROID_ARTIFACT=${ANDROID_OUT}/io.sonr.motor.aar
      mkdir -p ${ANDROID_OUT}

      echo "🔷 Binding Android..."
      cd ${MOTOR_LIB_DIR}
      gomobile bind -ldflags='-s -w' -target=android/arm64 -o ${ANDROID_ARTIFACT} -v
      ;;
    i)
      echo "🔷 Setting up build Environment..."
      IOS_OUT=${PROJECT_DIR}/build/ios
      IOS_ARTIFACT=${IOS_OUT}/SonrMotor.xcframework
      mkdir -p ${IOS_OUT}

      echo "🔷 Binding iOS..."
      cd ${MOTOR_LIB_DIR}
      gomobile bind -ldflags='-s -w' -target=ios/arm64 -o ${IOS_ARTIFACT} -v
      ;;
    w)
      echo "🔷 Setting up build Environment..."
      WASM_OUT=${PROJECT_DIR}/build/js
      WASM_ARTIFACT=${WASM_OUT}/sonr-motor.wasm
      mkdir -p ${WASM_OUT}

      echo "🔷 Binding Web..."
      cd ${MOTOR_LIB_DIR}
      GOOS=js GOARCH=wasm go build -tags wasm -o ${WASM_ARTIFACT} -v
      ;;
    ?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done
