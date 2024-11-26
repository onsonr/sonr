#!/usr/bin/env bash

set -e

ROOT_DIR=$(git rev-parse --show-toplevel)

bunx pkl project package $ROOT_DIR/pkl/*/

for dir in .out/*/; do
    folder=$(basename "$dir")
    rclone copy "$dir" "r2:pkljar/$folder"
done

rm -rf .out
