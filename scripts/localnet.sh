#!/bin/bash

SONRD_BIN=$1

KEY="alice"
CHAINID="sonr-localnet-1"
MONIKER="florence"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"

# remove existing daemon
rm -rf ~/.sonr*

$SONRD_BIN config keyring-backend $KEYRING
$SONRD_BIN config chain-id $CHAINID

# if $KEY exists it should be deleted
echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | $SONRD_BIN keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

$SONRD_BIN init $MONIKER --chain-id $CHAINID --home $HOME/.sonr

# Allocate genesis accounts (cosmos formatted addresses)
$SONRD_BIN add-genesis-account $KEY 100000000000000000000000000usnr --keyring-backend $KEYRING

# Sign genesis transaction
$SONRD_BIN gentx $KEY 1000000000000000000000usnr --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
$SONRD_BIN collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
$SONRD_BIN validate-genesis

$SONRD_BIN start --log_level $LOGLEVEL --pruning=nothing --rpc.unsafe --minimum-gas-prices=0.000006usnr --trace
