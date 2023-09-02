#!bin/bash


#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                    Utilities                                   ||
#// ! ||--------------------------------------------------------------------------------||

service_status() {
    if [ "$(uname)" == "Linux" ]; then
        sudo systemctl status $1
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

restart_service() {
    if [ "$(uname)" == "Linux" ]; then
        sudo systemctl daemon-reload
        if systemctl is-enabled --quiet $1; then
            sudo systemctl stop $1
            sudo systemctl restart $1
        fi
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

start_service() {
    if [ "$(uname)" == "Linux" ]; then
        if ! systemctl is-enabled --quiet $1; then
            sudo systemctl enable $1
            sudo systemctl start $1
        fi
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

stop_service() {
    if [ "$(uname)" == "Linux" ]; then
        if systemctl is-enabled --quiet $1; then
            sudo systemctl stop $1
        fi
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

download_release_file() {
    REPO=$1
    FILE=$2
    OUTPUT=$3
    wget https://github.com/$REPO/releases/latest/download/$FILE -O $OUTPUT
}

download_tarball_binary() {
    REPO=$1
    BINARY=$2
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m | tr '[:upper:]' '[:lower:]')
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

# Register icefiresql service
register_icefiresql_service() {
# Setup systemd for IceFireDB
sudo cat << EOF > /etc/systemd/system/icefiresql.service
[Unit]
Description=IceFireSQLite Service
After=network-online.target

[Service]
User=root
ExecStart=/usr/local/bin/icefiresql -c /var/lib/icefiresql/config.yml
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF

start_service icefiresql
}

# Setup systemd for sonrd
register_sonrd_service() {
sudo cat << EOF > /etc/systemd/system/sonrd.service
[Unit]
Description=Sonr Node Service
After=network-online.target

[Service]
User=root
ExecStart=/usr/local/bin/sonrd start
LimitNOFILE=4096
Environment=SONR_ENVIRONMENT=production
Environment=SONR_HIGHWAY_ICEFIREKV_HOST=localhost
Environment=SONR_HIGHWAY_ICEFIREKV_PORT=6001
Environment=SONR_HIGHWAY_ICEFIRESQL_HOST=localhost
Environment=SONR_HIGHWAY_ICEFIRESQL_PORT=23306
Environment=SONR_CHAIN_ID=sonr-testnet-1

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
    download_tarball_binary sonr-io/IceFireDB icefiresql
    mkdir -p /var/lib/icefiresql
    download_release_file sonr-io/IceFireDB config.sqlite.yml /var/lib/icefiresql/config.yml
}

register_services() {
    if [ "$(uname)" == "Linux" ]; then
        register_icefirekv_service
        register_icefiresql_service
        register_sonrd_service
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

upgrade() {
    download_tarball_binary sonr-io/sonr sonrd
    download_tarball_binary sonr-io/IceFireDB icefirekv
    download_tarball_binary sonr-io/IceFireDB icefiresql
    restart_service icefirekv
    restart_service icefiresql
    restart_service sonrd
}

#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                    Startup                                     ||
#// ! ||--------------------------------------------------------------------------------||

OPTIONS=("Install Sonr and Deps" "Register Systemd Service" "Upgrade Installation" "Check Service Status" "Exit")
PS3='Please enter your choice: '

select CHOICE in "${OPTIONS[@]}"
do
  echo ''
  case $CHOICE in
    "Install Sonr and Deps")
        install
        break
        ;;
    "Register Systemd Service")
        register_services
        break
        ;;
    "Upgrade Installation")
        upgrade
        break
        ;;
    "Check Service Status")
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
