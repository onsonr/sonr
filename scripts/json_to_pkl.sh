#!/usr/bin/env bash

set -e

SOURCE=$1
OUTPUT=$2

ROOT_DIR=$(git rev-parse --show-toplevel)
cd $ROOT_DIR

mkdir -p $OUTPUT
pkl eval package://pkg.pkl-lang.org/pkl-pantry/org.json_schema.contrib@1.0.0#/generate.pkl -m . -p source="$SOURCE" -p output="$OUTPUT"
