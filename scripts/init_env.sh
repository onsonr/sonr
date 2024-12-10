#!/usr/bin/env bash

set -e

ROOT_DIR=$(git rev-parse --show-toplevel)
export TURNSTILE_SITE_KEY=$(skate get TURNSTILE_SITE_KEY)
#
# export KEY="user1"
# export KEY2="user2"
#
# export CHAIN_ID=${CHAIN_ID:-"sonr-testnet-1"}
# export MONIKER="florence"
# export KEYALGO="secp256k1"
# export KEYRING=${KEYRING:-"test"}
# export HOME_DIR=$(eval echo "${HOME_DIR:-"~/.sonr"}")
# export BINARY=${BINARY:-sonrd}
# export DENOM=${DENOM:-usnr}
#
# export CLEAN=${CLEAN:-"true"}
# export RPC=${RPC:-"26657"}
# export REST=${REST:-"1317"}
# export PROFF=${PROFF:-"6969"}
# export P2P=${P2P:-"26656"}
# export GRPC=${GRPC:-"9090"}
# export GRPC_WEB=${GRPC_WEB:-"9091"}
# export ROSETTA=${ROSETTA:-"8420"}
# export BLOCK_TIME=${BLOCK_TIME:-"5s"}

# Check if process-compose is installed to $ROOT_DIR/build
if [ ! -f "$ROOT_DIR/bin/process-compose" ]; then
  echo "process-compose not found. Installing..."
  sh -c "$(curl --location https://raw.githubusercontent.com/F1bonacc1/process-compose/main/scripts/get-pc.sh)" -- -d -b $ROOT_DIR/bin
fi

