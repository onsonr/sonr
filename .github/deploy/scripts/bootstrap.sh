#!/usr/bin/env bash

set -e

# Ensure we're in the right directory
ROOT_DIR=$(git rev-parse --show-toplevel)
cd $ROOT_DIR

DOPPLER_TOKEN=$(skate get DOPPLER_NETWORK)

ACC0=$(doppler secrets get KEY0_NAME --plain --project sonr --config test)
ACC1=$(doppler secrets get KEY1_NAME --plain --project sonr --config test)
MNEM0=$(doppler secrets get KEY0_MNEMONIC --plain --project sonr --config test)
MNEM1=$(doppler secrets get KEY1_MNEMONIC --plain --project sonr --config test)
CHAIN_ID=$(doppler secrets get CHAIN_ID --plain --project sonr --config test)
TX_INDEX_INDEXER=$(doppler secrets get TX_INDEXER --plain --project sonr --config test)
TX_INDEX_PSQL_CONN=$(doppler secrets get TX_PSQL_CONN --plain --project sonr --config test)

# Run the node setup with all variables properly exported
CLEAN=true KEY0_NAME=$ACC0 KEY0_MNEMONIC=$MNEM0 KEY1_NAME=$ACC1 KEY1_MNEMONIC=$MNEM1 CHAIN_ID=$CHAIN_ID TX_INDEX_INDEXER=$TX_INDEX_INDEXER TX_INDEX_PSQL_CONN=$TX_INDEX_PSQL_CONN sh ./start.sh

