#!/bin/bash 
export SONRD_V1_HOME=/path/to/sonrd-v1
export SONRD_V1_BIN=$SONRD_V1_HOME/cosmovisor/genesis/bin/cored
export DAEMON_HOME=/path/to/sonrd-home
# Send SoftwareUpgrade proposal - Upgrade Name: v2.0.0
$SONRD_V1_BIN tx gov submit-proposal software-upgrade v2.0.0 --title v2.0.0 --description v2.0.0 --upgrade-height 40 --from validator1 --yes --home $DAEMON_HOME --chain-id app_9000-1
# Deposit for the proposal - Proposal ID: 1
$SONRD_V1_BIN tx gov deposit 1 10000000atoken --from validator1 --yes --home $DAEMON_HOME --chain-id app_9000-1
# Vote for the proposal