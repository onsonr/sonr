#!/bin/bash

if [ -n "$SEEDS" ]; then
  toml set $HOME/.sonr/config/config.toml p2p.seeds ${SEEDS_VPC} > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
fi

if [ -n "$PERSISTENT_PEERS" ]; then
  toml set $HOME/.sonr/config/config.toml p2p.persistent_peers ${PERSISTENT_PEERS_VPC} > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
fi

if [ -n "$PRIVATE_PEER_IDS" ]; then
  toml set $HOME/.sonr/config/config.toml p2p.private_peer_ids ${PRIVATE_PEER_IDS} > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
fi
