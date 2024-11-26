#!/usr/bin/env bash

set -e

ROOT_DIR=$(git rev-parse --show-toplevel)
STYLES_DIR=$ROOT_DIR/pkg/webapp
DIST_DIR=$STYLES_DIR/dist
OUT_DIR=$ROOT_DIR/static/css

cd $STYLES_DIR
bun install
bun run build
mkdir -p $OUT_DIR
mv $DIST_DIR/styles.css $OUT_DIR/styles.css
rm -rf $DIST_DIR
rm -rf $STYLES_DIR/node_modules

