#!/bin/bash

set -e  # Exit on any error

# Function to compare version strings
version_gt() {
    test "$(printf '%s\n' "$@" | sort -V | head -n 1)" != "$1"
}

# Install commitizen if not present
if ! command -v cz &> /dev/null; then
    echo "Installing commitizen..."
    pip install --user commitizen
fi

# Get all tags and sort them by version
echo "Fetching all tags..."
git fetch --tags --force
TAGS=$(git tag -l "v*" | sort -V)
LATEST_TAG=$(echo "$TAGS" | tail -n1)

if [ -z "$LATEST_TAG" ]; then
    echo "No tags found"
    exit 1
fi

echo "Latest tag: $LATEST_TAG"

# Run commitizen to determine next version
echo "Running commitizen bump --dry-run..."
NEXT_VERSION=$(cz bump --dry-run --increment=patch 2>&1 | grep "tag to create: v" | cut -d "v" -f2)

if [ -z "$NEXT_VERSION" ]; then
    echo "Failed to determine next version"
    exit 1
fi

echo "Next version determined by commitizen: v$NEXT_VERSION"

# Check if the next version already exists
if echo "$TAGS" | grep -q "v$NEXT_VERSION"; then
    echo "ERROR: Version v$NEXT_VERSION already exists!"
    exit 1
fi

# Verify the next version is actually greater than the latest
if ! version_gt "$NEXT_VERSION" "${LATEST_TAG#v}"; then
    echo "ERROR: Next version v$NEXT_VERSION is not greater than current version $LATEST_TAG"
    exit 1
fi

echo "✅ Version v$NEXT_VERSION is valid and does not exist yet"
exit 0
