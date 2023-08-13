#!/bin/bash

# remove existing daemon
rm -rf ~/.sonr*

sonrd config keyring-backend $KEYRING
sonrd config chain-id $CHAIN_ID

# if $KEY exists it should be deleted
echo $MNEMONIC | sonrd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

sonrd init $MONIKER --chain-id $CHAIN_ID

# Update config.toml to reflect changes
toml set $HOME/.sonr/config/config.toml rpc.laddr ${RPC_LADDR} > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
toml set $HOME/.sonr/config/app.toml grpc.address ${GRPC_ADDRESS} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
toml set $HOME/.sonr/config/app.toml api.enable ${API_ENABLE} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
toml set $HOME/.sonr/config/app.toml api.swagger ${API_SWAGGER} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
toml set $HOME/.sonr/config/app.toml api.address ${API_ADDRESS} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
toml set $HOME/.sonr/config/app.toml minimum-gas-prices ${MINIMUM_GAS_PRICES} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml



# Collect genesis tx
sonrd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
sonrd validate-genesis

