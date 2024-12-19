#!/bin/bash

set -e

# Ensure we're in the right directory
ROOT_DIR=$(git rev-parse --show-toplevel)
cd $ROOT_DIR

DOPPLER_TOKEN=$(skate get DOPPLER_NETWORK)

ACC0=$(doppler secrets get KEY0_NAME --plain)
ACC1=$(doppler secrets get KEY1_NAME --plain)
MNEM0=$(doppler secrets get KEY0_MNEMONIC --plain)
MNEM1=$(doppler secrets get KEY1_MNEMONIC --plain)
CHAIN_ID=$(doppler secrets get CHAIN_ID --plain)

# Run the node setup with all variables properly exported
KEY0_NAME=$ACC0 KEY0_MNEMONIC=$MNEM0 KEY1_NAME=$ACC1 KEY1_MNEMONIC=$MNEM1 CHAIN_ID=$CHAIN_ID CLEAN=true sh ./scripts/test_node.sh

