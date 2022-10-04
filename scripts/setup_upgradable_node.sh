#!/bin/bash
export DAEMON_NAME=sonrd
export SONR_V1_HOME=/home/ut/sonr_cosmvisor/comovisor_temp/bin_1
export SONR_V2_HOME=/home/ut/sonr_cosmvisor/comovisor_temp/bin_2
export DAEMON_HOME=/home/ut/sonr_cosmvisor/comovisor_temp/daemon
# Tmp cosmovisor directory. In this directory, we are going to
# setup Cosmovisor directory structures and copy this directory
# into $DAEMON_HOME
export TMP_COSMOVISOR_DIR=/home/ut/sonr_cosmvisor/temp/
# update if needed
export DAEMON_ALLOW_DOWNLOAD_BINARIES=false
export DAEMON_RESTART_AFTER_UPGRADE=true
# setup cored binaries
ignite chain build -t linux:amd64 --output .
cp sonrd $SONR_V1_HOME
cp sonrd $SONR_V2_HOME
# setup cored data
ignite chain init --home $DAEMON_HOME
#cp -R $TMP_COSMOVISOR_DIR $DAEMON_HOME/
cosmovisor run start --home $DAEMON_HOME
