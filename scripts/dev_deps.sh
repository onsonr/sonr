#!/bin/bash

# Determine the operating system
os=$(uname -s)
arch=$(uname -m)


install_pkl() {
    # Check if the pkl executable is already installed
    if command -v pkl &> /dev/null; then
        echo "Pkl is already installed. Version:"
        pkl --version
        return
    fi

    # Set the download URL based on the OS and architecture
    case "${os}-${arch}" in
    "Darwin-arm64")
        url="https://github.com/apple/pkl/releases/download/0.25.3/pkl-macos-aarch64"
        ;;
    "Darwin-x86_64")
        url="https://github.com/apple/pkl/releases/download/0.25.3/pkl-macos-amd64"
        ;;
    "Linux-aarch64")
        url="https://github.com/apple/pkl/releases/download/0.25.3/pkl-linux-aarch64"
        ;;
    "Linux-x86_64")
        url="https://github.com/apple/pkl/releases/download/0.25.3/pkl-linux-amd64"
        ;;
    *)
        echo "Unsupported operating system or architecture: ${os}-${arch}"
        exit 1
        ;;
    esac

    # Download the pkl executable
    curl -L -o pkl "${url}"

    # Make the pkl executable
    chmod +x pkl

    # Move the pkl executable to /usr/local/bin
    sudo mv pkl /usr/local/bin/

    # Verify the installation
    echo "Pkl installed successfully. Version:"
    pkl --version
}

install_pkl
go mod download
