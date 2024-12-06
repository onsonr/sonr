#!/bin/bash

set -e

ROOT_DIR=$(git rev-parse --show-toplevel)

# Extract scope name and path using jq, and pass it to fzf for selection
SCOPE=$(cat $ROOT_DIR/.github/scopes.json | jq -r '.[] | "\(.name)"' | fzf --prompt "Select scope:")
DOCS=$(cat $ROOT_DIR/.github/scopes.json | jq -r ".[] | select(.name == \"$SCOPE\") | .docs[]")
# Split the selected scope into name and path
echo $SCOPE
echo $DOCS

# Write Title
TITLE=$(gum write --placeholder "Title...")

# Write Goal
GOAL=$(mods --role "determine-issue-goal" "$SCOPE $TITLE")

REQUIREMENTS=$(gum write --placeholder "Requirements...")

echo "Title: $TITLE"
echo "Goal: $GOAL"
echo "Requirements: $REQUIREMENTS"
