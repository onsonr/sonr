#!/bin/bash

rm -rf $HOME/gentx-$MONIKER.json

# Sign genesis transaction
if [ -n "$ACCOUNT_NUMBER" ]; then
    sonrd gentx $KEY $STAKE --keyring-backend $KEYRING --chain-id $CHAIN_ID --account-number $ACCOUNT_NUMBER --output-document $HOME/gentx-$MONIKER.json
    GENTX=$(cat $HOME/gentx-$MONIKER.json) && doppler secrets set SONR_GENTX=$GENTX
fi
