#!/bin/bash
set -e

SCRIPTS_DIR=$(dirname "$0")
BUF_VERSION="1.12.0"
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
TMP_DIR=${PROJECT_DIR}/tmp
npm i -g mintlify
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ${TMP_DIR}

# Substitute BIN for your bin directory.
# Substitute VERSION for the current released version.
  curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${TMP_DIR}/buf" && \

chmod +x "${TMP_DIR}/buf"

sudo curl https://get.ignite.com/cli! | sudo bash
sudo mv ${TMP_DIR}/task /usr/local/bin
sudo mv ${TMP_DIR}/buf /usr/local/bin
rm -rf ${TMP_DIR}
