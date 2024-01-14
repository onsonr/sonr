#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./docs/static
cd proto
buf generate --template buf.gen.swagger.yaml
