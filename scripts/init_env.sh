#!/usr/bin/env bash

set -e

ROOT_DIR=$(git rev-parse --show-toplevel)

# Check if pkl-gen-go is installed to $GOPATH/bin
if [ ! -f "$GOPATH/bin/pkl-gen-go" ]; then
  echo "pkl-gen-go not found. Installing..."
  go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest
fi

# Check if sqlc is installed to $GOPATH/bin
if [ ! -f "$GOPATH/bin/sqlc" ]; then
  echo "sqlc not found. Installing..."
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
fi

# Check if goreleaser is installed to $GOPATH/bin
if [ ! -f "$GOPATH/bin/goreleaser" ]; then
  echo "goreleaser not found. Installing..."
  go install github.com/goreleaser/goreleaser/v2@latest
fi

# Check if templ is installed to $GOPATH/bin
if [ ! -f "$GOPATH/bin/templ" ]; then
  echo "templ not found. Installing..."
  go install github.com/a-h/templ/cmd/templ@latest
fi

