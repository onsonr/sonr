#!/bin/bash
echo "🔷 Building Library..."
SCRIPTDIR=$(dirname "$0")
cd ${SCRIPTDIR}/../../

echo "Setting up Project"
PROJECT_DIR=$(pwd)
CORE_BIND_DIR=${PROJECT_DIR}/cmd/bind
mkdir -p ${PROJECT_DIR}/build
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init

echo "Building for iOS"
cd ${CORE_BIND_DIR} && gomobile bind -ldflags='-s -w' -target=ios -bundleid=io.sonr.core -o ${PROJECT_DIR}/build/Core.framework
echo "✅  Finished Binding for iOS ➡ `date`"
