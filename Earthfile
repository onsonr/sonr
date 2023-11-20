VERSION 0.7
PROJECT sonrhq/sonr-testnet-0

faucet:
    FROM node:18.7-alpine
    ARG cosmosjsVersion=0.28.11
    RUN npm install @cosmjs/faucet@$cosmosjsVersion --global --production
    EXPOSE 4500
    ENTRYPOINT ["cosmos-faucet", "--chain-id", "sonr-testnet-0", "--node", "http://localhost:26657", "--denom", "usnr", "--keyring-backend", "test"]
    SAVE IMAGE sonrhq/faucet:latest
