#!/bin/bash

SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
MOTOR_LIB_DIR=${PROJECT_DIR}/bind/motor-mobile
MOTOR_WASM_DIR=${PROJECT_DIR}/bind/motor-wasm

while getopts "iaw" opt; do
  echo "🔷 Setting up build Environment..."
  VERSION=$(git describe --tags --abbrev=0)
  BUILDDIR=${PROJECT_DIR}/build
  mkdir -p ${BUILDDIR}

  case $opt in
    a)
      ANDROID_ARTIFACT=${BUILDDIR}/io.sonr.motor.aar
      ANDROID_TAR_BALL=${BUILDDIR}/motor-${VERSION}-android.tar.gz

      echo "🔷 Binding Android Artifact Version ${VERSION}..."
      cd ${MOTOR_LIB_DIR}
      gomobile bind -ldflags='-s -w' -target=android/arm64 -o ${ANDROID_ARTIFACT} -v
      rm -rf ${BUILDDIR}/io.sonr.motor-sources.jar
      tar -czvf ${ANDROID_TAR_BALL} ${ANDROID_ARTIFACT}
      rm -rf ${ANDROID_ARTIFACT}
      echo "✅ Android Tarball written to: ${ANDROID_TAR_BALL}"
      ;;
    i)
      IOS_ARTIFACT=${BUILDDIR}/Motor.xcframework
      IOS_TAR_BALL=${BUILDDIR}/motor-${VERSION}-ios.tar.gz

      echo "🔷 Binding iOS Artifact Version ${VERSION}..."
      cd ${MOTOR_LIB_DIR}
      gomobile bind -ldflags='-s -w' -target=ios -prefix=SNR  -o ${IOS_ARTIFACT} -v
      tar -czvf ${IOS_TAR_BALL} ${IOS_ARTIFACT}
      rm -rf ${IOS_ARTIFACT}
      echo "✅ iOS Tarball written to: ${IOS_TAR_BALL}"
      ;;
    w)
      WASM_ARTIFACT=${BUILDDIR}/sonr-motor.wasm
      WASM_TAR_BALL=${BUILDDIR}/motor-${VERSION}-wasm.tar.gz

      echo "🔷 Binding WebAssembly Artifact Version ${VERSION}..."
      cd ${MOTOR_WASM_DIR}
      GOOS=js GOARCH=wasm go build -tags wasm -o ${WASM_ARTIFACT} -v
      tar -czvf ${BUILDDIR}/motor-${VERSION}-wasm.tar.gz ${WASM_ARTIFACT}
      rm -rf ${WASM_ARTIFACT}
      echo "✅ WebAssembly Tarball written to: ${WASM_TAR_BALL}"
      ;;
    ?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done
