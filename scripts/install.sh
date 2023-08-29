#!/bin/bash

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | tr '[:upper:]' '[:lower:]')
DOWNLOAD_URL="https://github.com/sonr-io/sonr/releases/latest/download/sonrd-${OS}-${ARCH}"

if wget -q --show-error "${DOWNLOAD_URL}" -O /usr/local/bin/sonrd; then
    echo "File downloaded successfully."
else
    echo "Failed to download the file."
fi
