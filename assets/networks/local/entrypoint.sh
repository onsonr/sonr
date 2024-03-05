#!/bin/bash

KEY="alice"
CHAINID="sonr-localnet-1"
MONIKER="florence"
KEYALGO="secp256k1"
KEYRING="test"
HOME="./data"
SONR_HOME="$HOME/.sonr"

# remove existing daemon
rm -rf $HOME/.sonr
mkdir -p $HOME/.sonr

cp ./assets/public $SONR_HOME/public -r

# if $KEY exists it should be deleted
echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | sonrd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

sonrd init $MONIKER --chain-id $CHAINID --home $HOME/.sonr --default-denom usnr

# Allocate genesis accounts (cosmos formatted addresses)
sonrd genesis add-genesis-account $KEY 100000000000000000000000000usnr --keyring-backend $KEYRING

# Sign genesis transaction
sonrd genesis gentx $KEY 1000000000000000000000usnr --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
sonrd genesis collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
sonrd genesis validate-genesis
