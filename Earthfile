VERSION 0.7
PROJECT sonrhq/sonr-testnet-0

IMPORT ./chain AS chain
IMPORT ./identity AS identity
IMPORT ./service AS service
IMPORT ./studio AS studio

build:
    BUILD chain+build
    BUILD studio+build
    BUILD +faucet

generate:
    BUILD identity+generate
    BUILD service+generate

test:
    BUILD identity+test
    BUILD service+test

faucet:
    FROM node:18.7-alpine
    ARG cosmosjsVersion=0.28.11
    RUN npm install @cosmjs/faucet@$cosmosjsVersion --global --production
    COPY ./scripts/start-faucet.sh /usr/local/bin/faucet
    RUN chmod +x /usr/local/bin/faucet
    EXPOSE 4500
    ENTRYPOINT ["faucet"]
    SAVE IMAGE sonrhq/faucet:latest

build-vm:
    FROM nixos/nix
    RUN nix-channel --add https://nixos.org/channels/nixpkgs-unstable
    RUN nix-channel --update
    RUN nix-env -iA nixpkgs.bash
    RUN nix-env -iA nixpkgs.curl
    RUN nix-env -iA nixpkgs.git
    RUN nix-env -iA nixpkgs.go
    RUN nix-env -iA nixpkgs.go-task
    RUN nix-env -iA nixpkgs.docker
    RUN nix-env -iA nixpkgs.earthly
    RUN nix-env -iA nixpkgs.nodejs
    RUN nix-env -iA nixpkgs.gum
    RUN nix-env -iA nixpkgs.bob
    SAVE IMAGE sonrhq/vmbase:latest
