sonrd init testnet
sonrd keys add validator
sonrd add-genesis-account $(sonrd keys show validator -a) 1000000000000000000000000000stake,1000000000000000000000000000uatom
sonrd gentx validator 1000000000000000000000000000stake --chain-id testnet
sonrd collect-gentxs
sonrd validate-genesis
