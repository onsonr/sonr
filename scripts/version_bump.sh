#!/bin/bash
set -e

# Get the current version from most recent git tag
export VERSION=$(git describe --tags --abbrev=0)

# Check if the version is a valid semantic version
if ! [[ $VERSION =~ ^v?[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Invalid version format: $VERSION"
    exit 1
fi

# Check if the version has already been bumped

