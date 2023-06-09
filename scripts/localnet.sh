#!/bin/bash

go install github.com/charmbracelet/gum@latest
rm -rf $HOME/.sonr
SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
BIN_DIR=${PROJECT_DIR}/dist/core_darwin_arm64
NETWORK=sonr-localnet-0
DENOM=usnr
cd $BIN_DIR
chmod +x ./sonrd
./sonrd init $NETWORK --staking-bond-denom $DENOM
ACC_NAME=validator
./sonrd keys add $ACC_NAME
ACC_KEY=$(./sonrd keys show $ACC_NAME -a)
./sonrd add-genesis-account $ACC_KEY 100000snr,1000000000000000000000000000$DENOM
./sonrd gentx $ACC_NAME 1000000000000000$DENOM --chain-id $NETWORK
./sonrd collect-gentxs
sudo sed -i '.bak' 's/stake/usnr/g' $HOME/.sonr/config/genesis.json
./sonrd validate-genesis
