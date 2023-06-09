#!/bin/bash

echo "1. Create cert"
sudo apt install libnss3-tools
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
brew install mkcert
go install github.com/charmbracelet/gum@latest
ROOT_DOMAIN=$(gum input --placeholder "example.com")
ROOT_DOMAIN_CRT=/etc/ssl/$ROOT_DOMAIN/$ROOT_DOMAIN.crt
ROOT_DOMAIN_KEY=/etc/ssl/$ROOT_DOMAIN/$ROOT_DOMAIN.key
mkcert -cert-file $ROOT_DOMAIN_CRT -key-file $ROOT_DOMAIN_KEY $ROOT_DOMAIN api.$ROOT_DOMAIN rpc.$ROOT_DOMAIN grpc.$ROOT_DOMAIN p2p.$ROOT_DOMAIN localhost 127.0.0.1 ::1

echo "2. Setup NGINX"
sudo apt install nginx
sudo mkdir -p /etc/nginx/sites-available
sudo mkdir -p /etc/nginx/sites-enabled
sudo cat << EOF > /etc/nginx/sites-available/$ROOT_DOMAIN
server {
    listen 80;
    listen [::]:80;
    server_name $ROOT_DOMAIN;
    return 301 https://\$host\$request_uri;
}

server {
    listen 443 ssl;
    listen [::]:443 ssl;
    server_name $ROOT_DOMAIN;

    ssl_certificate $ROOT_DOMAIN_CRT;
    ssl_certificate_key $ROOT_DOMAIN_KEY;

    location / {
        proxy_pass http://localhost:26657;
        proxy_set_header Host \$host;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
    }
}

server {
    listen 443 ssl;
    listen [::]:443 ssl;
    server_name api.$ROOT_DOMAIN;

    ssl_certificate $ROOT_DOMAIN_CRT;
    ssl_certificate_key $ROOT_DOMAIN_KEY;

    location / {
        proxy_pass http://localhost:1317;
        proxy_set_header Host \$host;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
    }
}

server {
    listen 443 ssl;
    listen [::]:443 ssl;
    server_name rpc.$ROOT_DOMAIN;

    ssl_certificate $ROOT_DOMAIN_CRT;
    ssl_certificate_key $ROOT_DOMAIN_KEY;

    location / {
        proxy_pass http://localhost:26657;
        proxy_set_header Host \$host;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
    }
}

server {
    listen 443 ssl;
    listen [::]:443 ssl;
    server_name grpc.$ROOT_DOMAIN;

    ssl_certificate $ROOT_DOMAIN_CRT;
    ssl_certificate_key $ROOT_DOMAIN_KEY;

    location / {
        grpc_pass grpc://localhost:9090;
    }
}

server {
    listen 443 ssl;
    listen [::]:443 ssl;
    server_name p2p.$ROOT_DOMAIN;

    ssl_certificate $ROOT_DOMAIN_CRT;
    ssl_certificate_key $ROOT_DOMAIN_KEY;

    location / {
        proxy_pass http://localhost:26656;
        proxy_set_header Host \$host;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
    }
}
EOF
sudo ln -s /etc/nginx/sites-available/$ROOT_DOMAIN /etc/nginx/sites-enabled/$ROOT_DOMAIN

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


