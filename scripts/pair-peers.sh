#!/bin/bash
nodes=( v1.beta v2.beta v3.beta v4.beta)

for i in "${nodes[@]}"; do scp script/sonrd.service root@$(dig "$i".sonr.ws +short):/etc/systemd/system/sonrd.service; done

NODE_ID=sonrd tendermint show-node-id
EXISTING_PEERS=toml-cli get /root/.sonr/config/config.toml p2p.persistent_peers

toml-cli set /root/.sonr/config/config.toml p2p.persistent_peers
