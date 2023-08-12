#!/bin/bash

KEY="alice"
CHAINID="sonr-localnet-1"
MONIKER="florence"
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
sonrd add-genesis-account $KEY 100000000000000000000000000usnr --keyring-backend $KEYRING

# Sign genesis transaction
sonrd gentx $KEY 1000000000000000000000usnr --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
sonrd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
sonrd validate-genesis

# Update config.toml to reflect changes
toml set $HOME/.sonr/config/config.toml rpc.laddr tcp://0.0.0.0:26657 > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
toml set $HOME/.sonr/config/app.toml grpc.address 0.0.0.0:9000 > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
toml set $HOME/.sonr/config/app.toml api.enable true > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
toml set $HOME/.sonr/config/app.toml api.swagger true > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
toml set $HOME/.sonr/config/app.toml api.address tcp://0.0.0.0:1317 > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
toml set $HOME/.sonr/config/app.toml minimum-gas-prices 0.0000snr > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
