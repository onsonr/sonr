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


#// ! ||------------------------------------------------------------------------------||
#// ! ||                                   Services                                   ||
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
#// ! ||                              Install Dependencies                              ||
#// ! ||--------------------------------------------------------------------------------||

install() {
    download_tarball_binary sonr-io/sonr sonrd
    download_tarball_binary sonr-io/IceFireDB icefirekv
}

register_services() {
    register_icefirekv_service
    register_sonrd_service
}

upgrade() {
    stop_service sonrd
    stop_service icefirekv
    download_tarball_binary sonr-io/sonr sonrd
    download_tarball_binary sonr-io/IceFireDB icefirekv
    start_service icefirekv
    start_service sonrd
}

#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                    Startup                                     ||
#// ! ||--------------------------------------------------------------------------------||

OPTIONS=("Initialize Sonr Validator" "Register System Services on Linux" "Upgrade to latest Sonr binary" "Check status of System Services" "Exit")
PS3='Please enter your choice: '

select CHOICE in "${OPTIONS[@]}"
do
  case $CHOICE in
    "Initialize Sonr Validator")
        install
        break
        ;;
    "Register System Services on Linux")
        register_services
        break
        ;;
    "Upgrade to latest Sonr binary")
        upgrade
        break
        ;;
    "Check status of System Services")
        sudo systemctl status sonrd
        break
        ;;
    "Exit")
        echo "Exiting..."
        exit 1
        ;;
    *) echo "invalid option $REPLY";;
  esac
done
