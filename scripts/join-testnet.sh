#!/bin/bash

rm -rf $HOME/.sonr/config/genesis.json
doppler secrets get SONR_GENESIS_SIGNED --plain >> $HOME/.sonr/config/genesis.json
doppler secrets set NODE_INITIALIZED=true
