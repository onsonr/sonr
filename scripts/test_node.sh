#!/bin/bash
# Run this script to quickly install, setup, and run the current version of the network without docker.
#
# Examples:
# CHAIN_ID="local-1" HOME_DIR="~/.core" BLOCK_TIME="1000ms" CLEAN=true sh scripts/test_node.sh
# CHAIN_ID="local-2" HOME_DIR="~/.core" CLEAN=true RPC=36657 REST=2317 PROFF=6061 P2P=36656 GRPC=8090 GRPC_WEB=8091 ROSETTA=8081 BLOCK_TIME="500ms" sh scripts/test_node.sh

export KEY0_NAME=${KEY0_NAME:-"user0"}
export KEY0_MNEMONIC=${KEY0_MNEMONIC:-"decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry"}
export KEY1_NAME="user2"
export KEY1_MNEMONIC=${KEY1_MNEMONIC:-"wealth flavor believe regret funny network recall kiss grape useless pepper cram hint member few certain unveil rather brick bargain curious require crowd raise"}

export TX_INDEX_INDEXER=${TX_INDEX_INDEXER:-"kv"}
export TX_INDEX_PSQL_CONN=${TX_INDEX_PSQL_CONN:-""}
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
export PROFF=${PROFF:-"6060"}
export P2P=${P2P:-"26656"}
export GRPC=${GRPC:-"9090"}
export GRPC_WEB=${GRPC_WEB:-"9091"}
export ROSETTA=${ROSETTA:-"8080"}
export BLOCK_TIME=${BLOCK_TIME:-"5s"}

ROOT_DIR=$(git rev-parse --show-toplevel)

# if which binary does not exist, exit
if [ -z `which $BINARY` ]; then
  echo "Ensure $BINARY is installed and in your PATH"
  exit 1
fi

alias BINARY="$BINARY --home=$HOME_DIR"

command -v $BINARY > /dev/null 2>&1 || { echo >&2 "$BINARY command not found. Ensure this is setup / properly installed in your GOPATH (make install)."; exit 1; }
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

set_config() {
  $BINARY config set client chain-id $CHAIN_ID
  $BINARY config set client keyring-backend $KEYRING
}
set_config

from_scratch () {
  # Fresh install on current branch
  make install

  # remove existing daemon files.
  if [ ${#HOME_DIR} -le 2 ]; then
      echo "HOME_DIR must be more than 2 characters long"
      return
  fi
  rm -rf $HOME_DIR && echo "Removed $HOME_DIR"

  # reset values if not set already after whipe
  set_config

  add_key() {
    echo "Adding key: $1"
    key=$1
    mnemonic=$2
  }

  # idx1efd63aw40lxf3n4mhf7dzhjkr453axur9vjt6y
  echo "$KEY0_MNEMONIC" | BINARY keys add $KEY0_NAME --keyring-backend $KEYRING --algo $KEYALGO --recover
  echo "$KEY1_MNEMONIC" | BINARY keys add $KEY1_NAME --keyring-backend $KEYRING --algo $KEYALGO --recover

  # chain initial setup
  BINARY init $MONIKER --chain-id $CHAIN_ID --default-denom $DENOM

  update_test_genesis () {
    cat $HOME_DIR/config/genesis.json | jq "$1" > $HOME_DIR/config/tmp_genesis.json && mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json
  }

  # === CORE MODULES ===

  # Block
  update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'

  # Gov
  update_test_genesis `printf '.app_state["gov"]["params"]["min_deposit"]=[{"denom":"%s","amount":"1000000"}]' $DENOM`
  update_test_genesis '.app_state["gov"]["params"]["voting_period"]="30s"'
  update_test_genesis '.app_state["gov"]["params"]["expedited_voting_period"]="15s"'

  # staking
  update_test_genesis `printf '.app_state["staking"]["params"]["bond_denom"]="%s"' $DENOM`
  update_test_genesis '.app_state["staking"]["params"]["min_commission_rate"]="0.050000000000000000"'

  # mint
  update_test_genesis `printf '.app_state["mint"]["params"]["mint_denom"]="%s"' $DENOM`

  # crisis
  update_test_genesis `printf '.app_state["crisis"]["constant_fee"]={"denom":"%s","amount":"1000"}' $DENOM`

  # === CUSTOM MODULES ===
  # globalfee
  update_test_genesis `printf '.app_state["globalfee"]["params"]["minimum_gas_prices"]=[{"amount":"0.000000000000000000","denom":"%s"}]' $DENOM`
  # tokenfactory
  update_test_genesis '.app_state["tokenfactory"]["params"]["denom_creation_fee"]=[]'
  update_test_genesis '.app_state["tokenfactory"]["params"]["denom_creation_gas_consume"]=100000'
  # poa
  update_test_genesis '.app_state["poa"]["params"]["admins"]=["idx10d07y265gmmuvt4z0w9aw880jnsr700j9kqcfa"]'

  # Allocate genesis accounts
  BINARY genesis add-genesis-account $KEY0_NAME 10000000$DENOM,900snr --keyring-backend $KEYRING
  BINARY genesis add-genesis-account $KEY1_NAME 10000000$DENOM,800snr --keyring-backend $KEYRING

  # Sign genesis transaction
  BINARY genesis gentx $KEY0_NAME 1000000$DENOM --keyring-backend $KEYRING --chain-id $CHAIN_ID

  BINARY genesis collect-gentxs

  BINARY genesis validate-genesis
  err=$?
  if [ $err -ne 0 ]; then
    echo "Failed to validate genesis"
    return
  fi
}

# check if CLEAN is not set to false
if [ "$CLEAN" != "false" ]; then
  echo "Starting from a clean state"
  from_scratch
fi

echo "Starting node..."

# Tx Index
if [ "$TX_INDEX_PSQL_CONN" != "" ]; then
  awk -v conn="$TX_INDEX_PSQL_CONN" '/^psql-conn = / {$0 = "psql-conn = \"" conn "\""} 1' $HOME_DIR/config/config.toml > temp && mv temp $HOME_DIR/config/config.toml
fi


# Opens the RPC endpoint to outside connections
sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/c\laddr = "tcp:\/\/0.0.0.0:'$RPC'"/g' $HOME_DIR/config/config.toml
sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' $HOME_DIR/config/config.toml

# REST endpoint
sed -i 's/address = "tcp:\/\/localhost:1317"/address = "tcp:\/\/0.0.0.0:'$REST'"/g' $HOME_DIR/config/app.toml
sed -i 's/enable = false/enable = true/g' $HOME_DIR/config/app.toml

# peer exchange
sed -i 's/pprof_laddr = "localhost:6060"/pprof_laddr = "localhost:'$PROFF_LADDER'"/g' $HOME_DIR/config/config.toml
sed -i 's/laddr = "tcp:\/\/0.0.0.0:26656"/laddr = "tcp:\/\/0.0.0.0:'$P2P'"/g' $HOME_DIR/config/config.toml

# GRPC
sed -i 's/address = "localhost:9090"/address = "0.0.0.0:'$GRPC'"/g' $HOME_DIR/config/app.toml
sed -i 's/address = "localhost:9091"/address = "0.0.0.0:'$GRPC_WEB'"/g' $HOME_DIR/config/app.toml

# Rosetta Api
sed -i 's/address = ":8080"/address = "0.0.0.0:'$ROSETTA'"/g' $HOME_DIR/config/app.toml
sed -i 's/indexer = "kv"/indexer = "'$TX_INDEX_INDEXER'"/g' $HOME_DIR/config/config.toml

# Faster blocks
sed -i 's/timeout_commit = "5s"/timeout_commit = "'$BLOCK_TIME'"/g' $HOME_DIR/config/config.toml

# Start the node with 0 gas fees
BINARY start --pruning=nothing  --minimum-gas-prices=0$DENOM --rpc.laddr="tcp://0.0.0.0:$RPC" --grpc.address="0.0.0.0:$GRPC" --grpc-web.enable=true 
