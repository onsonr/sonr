#!/bin/bash

KEY="mykey"
CHAINID="test-1"
MONIKER="localtestnet"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"

# remove existing daemon
rm -rf ~/.sonr*

sonrd config keyring-backend $KEYRING
sonrd config chain-id $CHAINID

# if $KEY exists it should be deleted
echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | sonrd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

sonrd init $MONIKER --chain-id $CHAINID

# Allocate genesis accounts (cosmos formatted addresses)
sonrd add-genesis-account $KEY 100000000000000000000000000stake --keyring-backend $KEYRING

# Sign genesis transaction
sonrd gentx $KEY 1000000000000000000000stake --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
sonrd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
sonrd validate-genesis

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
sonrd start --pruning=nothing  --minimum-gas-prices=0.0001stake --rpc.laddr tcp://0.0.0.0:26657
