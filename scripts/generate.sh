#!/usr/bin/env bash

set -eo pipefail

# get protoc executions
go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos 2>/dev/null
SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
PROTO_DIR=${PROJECT_DIR}/proto

echo "Generating gogo proto code"
cd ${PROTO_DIR}
proto_dirs=$(find ./core -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package $file &>/dev/null; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

cd ..

# move proto files to the right places
#
# Note: Proto files are suffixed with the current binary version.
rm -rf github.com

go mod tidy -compat=1.19


# # Generate the Go code
# buf generate proto

# # Vault/Auth
# cp -r internal/gen/* ./types
# rm -rf internal/gen
# rm -rf types/core
# rm -rf types/sonr

# # Blockchain
# cp -r github.com/sonrhq/core/* .
# rm -rf github.com
# go mod tidy
