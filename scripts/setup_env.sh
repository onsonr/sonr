#!/bin/bash

set -e

ROOT_DIR=$(git rev-parse --show-toplevel)

# Function to install a Go binary if it doesn't exist
function go_install() {
	if ! command -v "$1" &>/dev/null; then
		echo "Installing $1..."
		go install "github.com/$1"
	fi
}

# Function to install a Cargo binary if it doesn't exist
function cargo_install() {
	if ! command -v "$1" &>/dev/null; then
		echo "Installing $1..."
		cargo install "$1"
	fi
}

# Function to install a uv tool if it doesn't exist
function uv_install() {
	if ! command -v "$1" &>/dev/null; then
		echo "Installing $1..."
		uv tool install "$1" --force
	fi
}

# Function to initialize git credentials
function set_git() {
	git config --global user.name "Darp Alakun"
	git config --global user.email "i@prad.nu"

	# Check if the GITHUB_TOKEN is set then authenticate with it if not ignore
	if [[ -z ${GITHUB_TOKEN} ]]; then
		echo "GITHUB_TOKEN is not set. Please set it before running this script."
		exit 1
	else
		gh auth login --with-token <<<"${GITHUB_TOKEN}"
	fi
}

function get_deps() {
	go_install go-task/task/v3/cmd/task@latest
	go_install x-motemen/ghq@latest
	go_install a-h/templ/cmd/templ@latest

	cargo_install ripgrep
	cargo_install fd-find
	cargo_install eza

	uv_install aider-chat
}

function clone_repos() {
	ghq get github.com/onsonr/sonr
	ghq get github.com/onsonr/nebula
	ghq get github.com/onsonr/hway
}

function main() {
	get_deps
	set_git
}

main
