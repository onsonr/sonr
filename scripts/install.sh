#!bin/bash


#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                    Utilities                                   ||
#// ! ||--------------------------------------------------------------------------------||


start_service() {
    sudo systemctl daemon-reload
    if ! systemctl is-enabled --quiet $1; then
        sudo systemctl enable $1
        sudo systemctl start $1
    fi
}

stop_service() {
    sudo systemctl daemon-reload
    if systemctl is-enabled --quiet $1; then
        sudo systemctl stop $1
        sudo systemctl disable $1
    fi
}

download_tarball_binary() {
    REPO=$1
    BINARY=$2
    OS=$(uname -s)
    ARCH=$(uname -m)
    wget https://github.com/$REPO/releases/latest/download/$BINARY-$OS-$ARCH.tar.gz
    sudo tar -xvf $BINARY-$OS-$ARCH.tar.gz -C /usr/local/bin
    rm $BINARY-$OS-$ARCH.tar.gz
}

#// ! ||--------------------------------------------------------------------------------||
#// ! ||                              Install Dependencies                              ||
#// ! ||--------------------------------------------------------------------------------||

install() {
    download_tarball_binary sonr-io/sonr sonrd
    download_tarball_binary sonr-io/IceFireDB icefirekv
}

#// ! ||------------------------------------------------------------------------------||
#// ! ||                                    Actions                                   ||
#// ! ||------------------------------------------------------------------------------||

# Register icefirekv service
register_icefirekv_service() {
# Setup systemd for IceFireDB
sudo cat << EOF > /etc/systemd/system/icefirekv.service
[Unit]
Description=IceFireKV Service
After=network-online.target

[Service]
User=root
ExecStart=/usr/local/bin/icefirekv start
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF

start_service icefirekv
}

# Setup systemd for sonrd
register_sonrd_service() {
sudo cat << EOF > /etc/systemd/system/sonrd.service
[Unit]
Description=Sonr Blockchain Node
After=network-online.target

[Service]
User=root
ExecStart=/usr/local/bin/sonrd start --home /root/.sonr --rpc.laddr tcp://0.0.0.0:26657
LimitNOFILE=4096
Environment=SONR_ENVIRONMENT=production
Environment=SONR_VALIDATOR_ADDRESS=$VALIDATOR_ADDRESS
Environment=SONR_PUBLIC_DOMAIN=$ROOT_DOMAIN
Environment=SONR_CHAIN_ID=$SONR_CHAIN_ID
Environment=SONR_ACCOUNT_ICEFIRE_ENABLED=true

[Install]
WantedBy=multi-user.target
EOF
start_service sonrd
}

#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                    Startup                                     ||
#// ! ||--------------------------------------------------------------------------------||

INIT="Initialize Sonr Validator"
UPGRADE="Upgrade to latest Sonr binary"
STATUS="Check status of System Services"
RESET="Reset Sonr Validator configuration"
EXIT="Exit"

CHOICE=$(gum choose --header "Select action..." "$INIT" "$STATUS" "$UPGRADE" "$RESET" "$EXIT")

if [ "$CHOICE" == "$INIT" ]; then
    reset_all
    install_fireice
    install_latest
    init_sonr
    gum confirm "Would you like to register nginx service?" && register_nginx
    gum confirm "Would you like to register sonrd service?" && register_sonrd
    exit 0
elif [ "$CHOICE" == "$UPGRADE" ]; then
    mkdir -p $HOME/.sonr-backup
    cp -r $HOME/.sonr/config/config.toml $HOME/.sonr-backup/config.toml
    cp -r $HOME/.sonr/config/app.toml $HOME/.sonr-backup/app.toml
    reset_sonr
    install_latest
    init_sonr
    mv $HOME/.sonr-backup/config.toml $HOME/.sonr/config/config.toml
    mv $HOME/.sonr-backup/app.toml $HOME/.sonr/config/app.toml
    rm -rf $HOME/.sonr-backup
    sudo systemctl daemon-reload
    register_sonrd
    exit 0
elif [ "$CHOICE" == "$RESET" ]; then
    reset_all
    exit 0
elif [ "$CHOICE" == "$STATUS" ]; then
    sudo systemctl status sonrd
    exit 0
elif [ "$CHOICE" == "$EXIT" ]; then
    echo "Exiting..."
    exit 1
fi
