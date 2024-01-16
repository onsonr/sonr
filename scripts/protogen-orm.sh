#!/usr/bin/env bash

set -e


cd proto
find ./sonrhq/identity/module/v1 -name 'state_query.proto' -delete
echo "Generating orm code"
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name 'state.proto'); do
      buf generate --template buf.gen.proto.yaml
      buf generate --template buf.gen.yaml
  done
done

cd ..
