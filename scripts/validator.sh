#!/bin/bash

echo "1. Reset genesis config"
rm -rf .sonr
sonrd init sonr-devnet-0 --staking-bond-denom usnr
sonrd keys add validator
sonrd add-genesis-account $(sonrd keys show validator -a) 100000snr,10000000000000000000000000usnr
sonrd gentx validator 1000000000usnr
sonrd collect-gentxs
sonrd validate-genesis

echo "2. Updating genesis for staking denom"
sudo sed -i 's/stake/usnr/g' .sonr/config/genesis.json

echo "3. Setting up system service"
sudo cat << EOF > /etc/systemd/system/sonrd.service
[Unit]
Description=Sonr Chain Daemon
After=network-online.target

[Service]
User=root
ExecStart=/usr/local/bin/sonrd start --rpc.laddr tcp://0.0.0.0:26657
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF

echo "4. Enable Service"
sudo systemctl daemon-reload
sudo systemctl enable sonrd
