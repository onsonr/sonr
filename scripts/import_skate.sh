#!/usr/bin/env bash

set -e

export KEY="user1"
export KEY2="user2"

export CHAIN_ID=${CHAIN_ID:-"sonr-testnet-1"}
export MONIKER="florence"
export KEYALGO="secp256k1"
export KEYRING=${KEYRING:-"test"}
export HOME_DIR=$(eval echo "${HOME_DIR:-"~/.sonr"}")
export BINARY=${BINARY:-sonrd}
export DENOM=${DENOM:-usnr}

export CLEAN=${CLEAN:-"true"}
export RPC=${RPC:-"26657"}
export REST=${REST:-"1317"}
export PROFF=${PROFF:-"6969"}
export P2P=${P2P:-"26656"}
export GRPC=${GRPC:-"9090"}
export GRPC_WEB=${GRPC_WEB:-"9091"}
export ROSETTA=${ROSETTA:-"8420"}
export BLOCK_TIME=${BLOCK_TIME:-"5s"}
