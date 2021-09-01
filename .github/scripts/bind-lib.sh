#!/bin/bash
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind -ldflags='-s -w' -v
GOBIND_ANDROID=$(GOBIND) -target=android
GOBIND_IOS=$(GOBIND) -target=ios -bundleid=io.sonr.core

echo "🔷 Binding Library..."
BASEDIR=$(dirname "$0")
cd ${BASEDIR}/../../


