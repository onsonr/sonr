#!/bin/bash

# Allocate genesis accounts (cosmos formatted addresses)
sonrd add-genesis-account $KEY ${BALANCE} --keyring-backend $KEYRING

# Sign genesis transaction
sonrd gentx $KEY $STAKE --keyring-backend $KEYRING --chain-id $CHAIN_ID
