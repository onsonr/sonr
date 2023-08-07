# ! ||--------------------------------------------------------------------------------||
# ! ||                                  Cosmjs Faucet                                 ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux node:18.7-alpine AS cosmos-faucet

ENV COSMJS_VERSION=0.28.11

RUN npm install @cosmjs/faucet@${COSMJS_VERSION} --global --production

ENV FAUCET_CONCURRENCY=2
ENV FAUCET_PORT=4500
ENV FAUCET_GAS_PRICE=0.001stake
# Prepared keys for determinism
ENV FAUCET_MNEMONIC="decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry"
ENV FAUCET_ADDRESS_PREFIX=idx
ENV FAUCET_TOKENS="usnr, snr"
ENV FAUCET_CREDIT_AMOUNT_STAKE=100
ENV FAUCET_CREDIT_AMOUNT_TOKEN=100
ENV FAUCET_COOLDOWN_TIME=0

EXPOSE 4500

ENTRYPOINT [ "cosmos-faucet" ]


# ! ||--------------------------------------------------------------------------------||
# ! ||                                  Sonrd Builder                                 ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux golang:1.19-alpine AS builder

ENV SONR_VERSION=master
RUN apk add --update gcc g++ make git curl

WORKDIR /root
RUN git clone --depth 1 --branch ${SONR_VERSION} https://github.com/sonrhq/core.git sonr

WORKDIR /root/sonr
RUN go build -o ./build/sonrd ./cmd/sonrd/main.go


# ! ||--------------------------------------------------------------------------------||
# ! ||                               Sonr in Production                               ||
# ! ||--------------------------------------------------------------------------------||

FROM --platform=linux alpine

WORKDIR /root

# Copy the sonrd binary from the builder stage and local config
COPY --from=builder /root/sonr/build/sonrd /usr/local/bin/sonrd
COPY sonr.yml .

# Setup environment variables
ENV KEY="alice"
ENV CHAIN_ID=sonr-1
ENV MONIKER=florence
ENV KEYALGO=secp256k1
ENV KEYRING=test
ENV MNEMONIC="decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry"

# Initialize the node
RUN echo $MNEMONIC | sonrd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover
RUN sonrd init ${MONIKER} --chain-id ${CHAIN_ID} --default-denom usnr --home /root/.sonr
RUN sonrd add-genesis-account $KEY 100000000000000000000000000usnr,1000000000000000snr --keyring-backend $KEYRING
RUN sonrd gentx $KEY 1000000000000000000000usnr --keyring-backend $KEYRING --chain-id $CHAIN_ID
RUN sonrd collect-gentxs

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
EXPOSE 9090

CMD [ "sonrd", "start" ]
