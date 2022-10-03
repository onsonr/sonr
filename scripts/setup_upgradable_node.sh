#!/bin/bash
export DAEMON_NAME=sonrd
export SONR_V1_HOME=/home/ut/sonr_cosmvisor/temp/
export SONR_V2_HOME=/home/ut/sonr_cosmvisor/temp/
export DAEMON_HOME=/path/to/sonrd-home
# Tmp cosmovisor directory. In this directory, we are going to
# setup Cosmovisor directory structures and copy this directory
# into $DAEMON_HOME
export TMP_COSMOVISOR_DIR=/path/to/tmp-cosmovisor
# update if needed
export DAEMON_ALLOW_DOWNLOAD_BINARIES=false
export DAEMON_RESTART_AFTER_UPGRADE=true
# setup cored binaries
cd $SONR_V1_HOME
make build
cp ./build/sonrd $TMP_COSMOVISOR_DIR/genesis/bin/
cd $SONR_V2_HOME
make build
cp ./build/sonrd $TMP_COSMOVISOR_DIR/upgrades/v2.0.0/bin/
# setup cored data
cd $SONR_V1_HOME
ignite chain init --home $DAEMON_HOME
cp -R $TMP_COSMOVISOR_DIR $DAEMON_HOME/
cosmovisor run start --home $DAEMON_HOME
