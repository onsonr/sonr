#!/bin/bash

sudo mkdir -p /var/log/sonrd 
sudo touch /var/log/sonrd/digitaloceand.log 
sudo touch /var/log/sonrd/sonrd_error.log 


sonrd init my-node --chain-id sonr
# ## Setup an alice account
export ALICE_TEXT=$(sonrd keys add --keyring-backend test alice --home /root/.sonr)
export ALICE_ACCOUNT=$(echo $ALICE_TEXT| grep address | awk '{ print $2 }')
echo "ALICE_TEXT:"; echo $ALICE_TEXT \n
echo $ALICE_TEXT >> alice.txt

# Setup a bob account
export BOB_TEXT=$(sonrd keys add --keyring-backend test bob --home /root/.sonr| grep address | awk '{ print $2 }')
echo "BOB_TEXT:"; echo $BOB_TEXT \n
echo $BOB_TEXT >> bob.txt
# Setup the genesis accounts with $$
sonrd add-genesis-account $(sonrd keys show -a alice --keyring-backend test) 1000000000000000stake,1000000000000snr
sonrd add-genesis-account $(sonrd keys show -a bob --keyring-backend test) 1000000000000000stake,1000000000000snr
# Setup alice as the validator
sonrd gentx alice 1000000000000000stake --keyring-backend test --chain-id sonr
# collect the genesis transaction
sonrd collect-gentxs
# Echo the account addresses
echo "ALICE_ACCOUNT:"; echo $(sonrd keys show -a alice --keyring-backend test)
echo "BOB_ACCOUNT:"; echo $(sonrd keys show -a bob --keyring-backend test)

# Enable the ssh API in /root/.sonr/config/app.toml using https://github.com/MinseokOh/toml-cli
toml-cli set /root/.sonr/config/app.toml api.enable true
toml-cli set /root/.sonr/config/app.toml api.swagger true
toml-cli set /root/.sonr/config/app.toml grpc-web true

toml-cli set /root/.sonr/config/client.toml chain-id sonr

toml-cli set /root/.sonr/config/config.toml rpc.laddr tcp://0.0.0.0:26657
