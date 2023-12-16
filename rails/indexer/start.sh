#!bin/bash

cosmos-indexer update-denoms \
--update-all \
--log.pretty = true \
--log.level = debug \
--base.index-chain = true \
--base.start-block 1 \
--base.end-block -1 \
--base.throttling 2.005 \
--base.rpc-workers 1 \
--base.reindex true \
--base.prevent-reattempts true \
--base.api http://sonrd:1317 \
--probe.rpc http://sonrd:26657 \
--probe.account-prefix idx \
--probe.chain-id sonr-testnet-1 \
--probe.chain-name sonr \
--database.host postgres \
--database.database postgres \
--database.user taxuser \
--database.password password

cosmos-indexer index \
--log.pretty = true \
--log.level = debug \
--base.index-chain = true \
--base.start-block 1 \
--base.end-block -1 \
--base.throttling 2.005 \
--base.rpc-workers 1 \
--base.reindex true \
--base.prevent-reattempts true \
--base.api http://sonrd:1317 \
--probe.rpc http://sonrd:26657 \
--probe.account-prefix idx \
--probe.chain-id sonr-testnet-1 \
--probe.chain-name sonr \
--database.host postgres \
--database.database postgres \
--database.user taxuser \
--database.password password
