#!/bin/bash

# remove existing daemon
if [ -n "$NODE_INITIALIZED" ]; then
  exit 0
fi

rm -rf ~/.sonr
sonrd config keyring-backend $KEYRING
sonrd config chain-id $CHAIN_ID

# if $KEY exists it should be deleted
echo $MNEMONIC | sonrd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

sonrd init $MONIKER --chain-id $CHAIN_ID
# Update app.toml to reflect changes
if [ -n "$GRPC_ADDRESS" ]; then
  sed -i "s/grpc.address = .*/grpc.address = \"$GRPC_ADDRESS\"/" $HOME/.sonr/config/app.toml
fi

if [ -n "$API_ENABLE" ]; then
  sed -i "s/api.enable = .*/api.enable = $API_ENABLE/" $HOME/.sonr/config/app.toml
fi

if [ -n "$API_SWAGGER" ]; then
  sed -i "s/api.swagger = .*/api.swagger = $API_SWAGGER/" $HOME/.sonr/config/app.toml
fi

if [ -n "$API_ADDRESS" ]; then
  sed -i "s/api.address = .*/api.address = \"$API_ADDRESS\"/" $HOME/.sonr/config/app.toml
fi

if [ -n "$MINIMUM_GAS_PRICES" ]; then
  sed -i "s/minimum-gas-prices = .*/minimum-gas-prices = \"$MINIMUM_GAS_PRICES\"/" $HOME/.sonr/config/app.toml
fi

if [ -n "$SEEDS" ]; then
  sed -i "s/p2p.seeds = .*/p2p.seeds = \"$SEEDS\"/" $HOME/.sonr/config/config.toml
fi

if [ -n "$PERSISTENT_PEERS" ]; then
  sed -i "s/p2p.persistent_peers = .*/p2p.persistent_peers = \"$PERSISTENT_PEERS\"/" $HOME/.sonr/config/config.toml
fi

if [ -n "$PRIVATE_PEER_IDS" ]; then
  sed -i "s/p2p.private_peer_ids = .*/p2p.private_peer_ids = \"$PRIVATE_PEER_IDS\"/" $HOME/.sonr/config/config.toml
fi

rm -rf $HOME/.sonr/config/genesis.json
doppler secrets get SONR_GENESIS --plain >> $HOME/.sonr/config/genesis.json
sonrd tendermint show-node-id >> node-id.txt
NODEID=$(tail -1 node-id.txt) && doppler secrets set NODE_ID=$NODEID
rm -rf node-id.txt

rm -rf $HOME/gentx-$MONIKER.json

# Sign genesis transaction
if [ -n "$ACCOUNT_NUMBER" ]; then
    sonrd gentx $KEY $STAKE --keyring-backend $KEYRING --chain-id $CHAIN_ID --account-number $ACCOUNT_NUMBER --output-document $HOME/gentx-$MONIKER.json
    GENTX=$(cat $HOME/gentx-$MONIKER.json) && doppler secrets set SONR_GENTX=$GENTX
fi
