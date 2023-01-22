#!/bin/bash
set -e

SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
PROTO_DIR=${PROJECT_DIR}/proto

# Generate the Go code
buf generate proto

cp -r github.com/sonr-hq/sonr/* .
rm -rf github.com
go mod tidy
