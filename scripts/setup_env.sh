#!/bin/bash

ROOT_DIR=$(git rev-parse --show-toplevel)

# Function to install a Go binary if it doesn't exist
function go_install() {
	if ! command -v "$1" &>/dev/null; then
		echo "Installing $1..."
		go install "github.com/$1"
	fi
}

# Function to install a gh extension if it doesn't exist. Check gh <extension> for checking if installed
function gh_ext_install() {
	gh extension install "$1"
}

function main() {
	go_install go-task/task/v3/cmd/task@latest
	go_install a-h/templ/cmd/templ@latest
	go_install goreleaser/goreleaser/v2@latest

	gh_ext_install johnmanjiro13/gh-bump
}

main
