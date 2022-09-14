#!/bin/bash
set -e

SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
MOTOR_LIB_DIR=${PROJECT_DIR}/bind/motor-mobile
MOTOR_WASM_DIR=${PROJECT_DIR}/bind/motor-wasm

while getopts "iawm" opt; do
  echo "🔷 Setting up build Environment..."
  VERSION=$(git describe --tags --abbrev=0)
  BUILDDIR=${PROJECT_DIR}/build
  mkdir -p ${BUILDDIR}

  case $opt in
    a)
      ANDROID_ARTIFACT=${BUILDDIR}/io.sonr.motor.aar
      echo "🔷 Binding Android Artifact Version ${VERSION}..."
      cd ${MOTOR_LIB_DIR}
      gomobile bind -ldflags='-s -w' -target=android/arm64 -o ${ANDROID_ARTIFACT} -v
      rm -rf ${BUILDDIR}/io.sonr.motor-sources.jar
      cd ${BUILDDIR}

      if [ "$TAR_COMPRESS" = true ] ; then
        echo "🔷 Compressing Android Artifact..."
        tar -czf ${BUILDDIR}/motor-${VERSION}-android.tar.gz io.sonr.motor.aar
        rm -rf ${ANDROID_ARTIFACT}
        echo "✅ Android Tarball written to: ${ANDROID_TAR_BALL}"
      fi
      ;;
    i)
      IOS_ARTIFACT=${BUILDDIR}/Motor.xcframework
      echo "🔷 Binding iOS Artifact Version ${VERSION}..."
      cd ${MOTOR_LIB_DIR}
      gomobile bind -ldflags='-s -w' -target=ios -prefix=SNR  -o ${IOS_ARTIFACT} -v
      cd ${BUILDDIR}

      if [ "$TAR_COMPRESS" = true ] ; then
        echo "🔷 Compressing iOS Artifact..."
        tar -czf ${BUILDDIR}/motor-${VERSION}-ios.tar.gz Motor.xcframework
        rm -rf ${IOS_ARTIFACT}
        echo "✅ iOS Tarball written to: ${IOS_TAR_BALL}"
      fi
      ;;
    w)
      WASM_ARTIFACT=${BUILDDIR}/sonr-motor.wasm
      WASM_TAR_BALL=${BUILDDIR}/motor-${VERSION}-wasm.tar.gz

      echo "🔷 Binding WebAssembly Artifact Version ${VERSION}..."
      cd ${MOTOR_WASM_DIR}
      GOOS=js GOARCH=wasm go build -tags wasm -o ${WASM_ARTIFACT} -v
      cd ${BUILDDIR}

      if [ "$TAR_COMPRESS" = true ] ; then
        echo "🔷 Compressing WebAssembly Artifact..."
        tar -czf ${WASM_TAR_BALL} sonr-motor.wasm
        rm -rf ${WASM_ARTIFACT}
      fi
      echo "✅ WebAssembly Tarball written to: ${WASM_TAR_BALL}"
      ;;
    ?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done
