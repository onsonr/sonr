#!bin/bash

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | tr '[:upper:]' '[:lower:]')

wget https://github.com/sonr-io/sonr/releases/latest/download/sonrd-${OS}-${ARCH} -O /usr/local/bin/sonrd
