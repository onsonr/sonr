#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Function to check if a command exists.
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

echo "Starting the build process for Caddy with Cloudflare DNS module..."

# Check if Go is installed
if ! command_exists go; then
    echo "Go is not installed. Please install Go before running this script."
    exit 1
fi

# Set Go environment variables
export GOPATH=$(go env GOPATH)
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

# Install xcaddy if not present
if ! command_exists xcaddy; then
    echo "xcaddy not found. Installing xcaddy..."
    curl -sSfL https://raw.githubusercontent.com/caddyserver/xcaddy/master/install.sh | bash -s -- -b $GOBIN
fi

# Build Caddy with the Cloudflare DNS module
echo "Building Caddy with the Cloudflare DNS module..."
xcaddy build --with github.com/caddy-dns/cloudflare
mv caddy ../build/caddy
echo "Caddy has been built successfully with the Cloudflare DNS module."

# Optional: Move the caddy binary to /usr/local/bin (requires sudo)
# echo "Moving caddy to /usr/local/bin (requires sudo)..."
# sudo mv caddy /usr/local/bin/

# echo "Caddy has been installed to /usr/local/bin."
