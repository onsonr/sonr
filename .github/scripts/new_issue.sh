#!/bin/bash

set -e

ROOT_DIR=$(git rev-parse --show-toplevel)

# Extract scope name and path using jq, and pass it to fzf for selection
SCOPE=$(cat "$ROOT_DIR/.github/scopes.json" | jq -r '.[] | "\(.name)"' | fzf --prompt "Select scope:")
DOCS=$(cat "$ROOT_DIR/.github/scopes.json" | jq -r ".[] | select(.name == \"$SCOPE\") | .docs[].url")

# Write Title
TITLE=$(gum input --placeholder "Issue Title...")

# Write Goal
GOAL=$(mods --role "determine-issue-goal" "$SCOPE $TITLE")

# Input Requirements
REQUIREMENTS=()
while true; do
    if [ ${#REQUIREMENTS[@]} -ge 2 ]; then
        if ! gum confirm "Do you want to add another requirement?"; then
            break
        fi
    fi
    REQUIREMENT=$(gum input --placeholder "Add a requirement...")
    if [ -n "$REQUIREMENT" ]; then
        REQUIREMENTS+=("$REQUIREMENT")
    else
        echo "Requirement cannot be empty. Please enter a valid requirement."
    fi
done

create_body() {
    echo "### Goal(s):"
    echo "$GOAL"
    echo "### Requirements:"
    for i in "${!REQUIREMENTS[@]}"; do
        echo "$(($i + 1)). ${REQUIREMENTS[$i]}"
    done
    echo "### Resources:"
    while IFS= read -r doc; do
        echo "- $doc"
    done <<< "$DOCS"
}

ISSUE_BODY=$(create_body)

# Function to collect output
preview_output() {
    echo "# ($SCOPE) $TITLE"
    echo "$ISSUE_BODY"
}

# Display the formatted output
preview_output | gum format

# Confirm to create a GitHub issue
if gum confirm "Do you want to create a new GitHub issue with this information?"; then
    # Create a new GitHub issue using the gh CLI
    gh issue create --repo onsonr/sonr --title "($SCOPE) $TITLE" --body "$ISSUE_BODY"
else
  exit 1
fi
