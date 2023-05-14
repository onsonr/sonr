sonrd init sonr-devnet-0 --staking-bond-denom usnr
sonrd keys add validator
sonrd add-genesis-account $(sonrd keys show validator -a) 100000snr,10000000000000000000000000usnr
sonrd gentx validator 1000000000usnr
sonrd collect-gentxs
sonrd validate-genesis
