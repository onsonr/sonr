#!/bin/bash

sonrd config set client chain-id $SONR_CHAIN_ID
sonrd config set client keyring-backend $SONR_KEYRING_BACKEND
sonrd config set app api.enable true
sonrd config set app api.swagger true
sonrd config set app api.address $SONR_API_ADDRESS
sonrd keys add alice --recover $VALIDATOR_MNEMONIC
sonrd keys add bob --recover $FAUCET_MNEMONIC

