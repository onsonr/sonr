#!/bin/bash
set -e

SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
LICENSE=${PROJECT_DIR}/LICENSE.md
MOTOR_LIB_DIR=${PROJECT_DIR}/bind/motor-mobile
MOTOR_WASM_DIR=${PROJECT_DIR}/bind/motor-wasm




while getopts "iaw" opt; do
  echo "🔷 Setting up build Environment..."
  VERSION=$(git describe --tags --abbrev=0)
  BUILDDIR=${PROJECT_DIR}/build
  mkdir -p ${BUILDDIR}

  if ! command -v gomobile &> /dev/null
  then
    echo "gomobile could not be found. Installing it..."
    go install golang.org/x/mobile/cmd/gomobile@latest
    gomobile init
  fi

  case $opt in
    a)
      ANDROID_ARTIFACT=${BUILDDIR}/io.sonr.motor.aar
      echo "🔷 Binding Android Artifact - Version:${VERSION}..."
      gomobile bind -ldflags='-s -w' -target=android/arm64 -o ${ANDROID_ARTIFACT} -androidapi 19 -v ${MOTOR_LIB_DIR}
      rm -rf ${BUILDDIR}/io.sonr.motor-sources.jar
      cd ${BUILDDIR}

      if [ "$TAR_COMPRESS" = true ] ; then
        echo "🔷 Compressing Android Artifact..."
        zip -r ${BUILDDIR}/motor_android.zip io.sonr.motor.aar
        rm -rf ${ANDROID_ARTIFACT}
        echo "✅ Android Tarball written to: ${ANDROID_TAR_BALL}"
      fi
      ;;
    i)
      IOS_ARTIFACT=${BUILDDIR}/Motor.xcframework
      echo "🔷 Binding Universal iOS/macOS Artifact - Version: ${VERSION}..."
      gomobile bind -ldflags='-s -w' -target=ios,iossimulator,macos -prefix=SNR  -o ${IOS_ARTIFACT} -v ${MOTOR_LIB_DIR}
      cd ${BUILDDIR}
      cp ${LICENSE} ${IOS_ARTIFACT}/LICENSE.md

      if [ "$TAR_COMPRESS" = true ] ; then
        echo "🔷 Compressing iOS Artifact..."
        zip -r ${BUILDDIR}/motor_universal.zip Motor.xcframework
        rm -rf ${IOS_ARTIFACT}
        echo "✅ iOS Tarball written to: ${IOS_TAR_BALL}"
      fi
      ;;
    w)
      WASM_ARTIFACT=${BUILDDIR}/sonr-motor.wasm
      echo "🔷 Binding WebAssembly Artifact Version ${VERSION}..."
      GOOS=js GOARCH=wasm go build -tags wasm -o ${WASM_ARTIFACT} -v ${MOTOR_WASM_DIR}
      cd ${BUILDDIR}

      if [ "$TAR_COMPRESS" = true ] ; then
        echo "🔷 Compressing WebAssembly Artifact..."
        zip -r ${BUILDDIR}/motor_wasm.zip sonr-motor.wasm
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
