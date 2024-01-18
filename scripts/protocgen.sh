#!/usr/bin/env bash

set -e


cd proto
echo "Generating gogo proto code"
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep -q "option go_package" "$file" | grep -q ':0$'; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

buf generate --template buf.gen.pulsar.yaml

cd ..

cp -r github.com/sonrhq/sonr/x/* ./
rm -rf api && mkdir api
mv sonrhq/* ./api
rm -rf github.com sonrhq
