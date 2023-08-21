#!/usr/bin/env bash

# Generate the Go code
buf generate proto

# Vault/Auth
cp -r internal/gen/* ./types
rm -rf internal/gen
rm -rf types/core
rm -rf types/sonr

# Blockchain
cp -r github.com/sonrhq/core/* .
rm -rf github.com
go mod tidy
