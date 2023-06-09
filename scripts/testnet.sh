#!/bin/bash

go install github.com/charmbracelet/gum@latest
rm -rf $HOME/.sonr
NETWORK=$(gum choose "sonr-localnet-0" "sonr-devnet-1" "sonr-testnet-0")
DENOM=$(gum choose "usnr" "stake")
sonrd init $NETWORK --staking-bond-denom $DENOM
ACC_NAME=$(gum input --placeholder "Account Name")
sonrd keys add $ACC_NAME
ACC_KEY=$(sonrd keys show $ACC_NAME -a)
sonrd add-genesis-account $ACC_KEY 100000snr,1000000000000000000000000000$DENOM
sonrd gentx $ACC_NAME 1000000000000000$DENOM --chain-id $NETWORK
sonrd collect-gentxs
sudo sed -i '.bak' 's/stake/usnr/g' $HOME/.sonr/config/genesis.json
sonrd validate-genesis

