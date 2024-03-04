#!/bin/bash

# remove existing daemon
rm -rf ~/.sonr

sonrd config keyring-backend $KEYRING
sonrd config chain-id $CHAIN_ID

# if $KEY exists it should be deleted
echo $MNEMONIC | sonrd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

sonrd init $MONIKER --chain-id $CHAIN_ID

# Update app.toml to reflect changes
if [ -n "$GRPC_ADDRESS" ]; then
  toml set $HOME/.sonr/config/app.toml grpc.address ${GRPC_ADDRESS} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
fi

if [ -n "$API_ENABLE" ]; then
  toml set $HOME/.sonr/config/app.toml api.enable ${API_ENABLE} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
fi

if [ -n "$API_SWAGGER" ]; then
  toml set $HOME/.sonr/config/app.toml api.swagger ${API_SWAGGER} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
fi

if [ -n "$API_ADDRESS" ]; then
  toml set $HOME/.sonr/config/app.toml api.address ${API_ADDRESS} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
fi

if [ -n "$MINIMUM_GAS_PRICES" ]; then
  toml set $HOME/.sonr/config/app.toml minimum-gas-prices ${MINIMUM_GAS_PRICES} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
fi

if [ -n "$SEEDS" ]; then
  toml set $HOME/.sonr/config/config.toml p2p.seeds ${SEEDS_VPC} > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
fi

if [ -n "$PERSISTENT_PEERS" ]; then
  toml set $HOME/.sonr/config/config.toml p2p.persistent_peers ${PERSISTENT_PEERS_VPC} > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
fi

if [ -n "$PRIVATE_PEER_IDS" ]; then
  toml set $HOME/.sonr/config/config.toml p2p.private_peer_ids ${PRIVATE_PEER_IDS} > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
fi

rm -rf $HOME/.sonr/config/genesis.json
doppler secrets get SONR_GENESIS --plain >> $HOME/.sonr/config/genesis.json
sonrd tendermint show-node-id >> node-id.txt
NODEID=$(tail -1 node-id.txt) && doppler secrets set NODE_ID=$NODEID
rm -rf node-id.txt
