ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"

# ! ||--------------------------------------------------------------------------------||
# ! ||                                  Cosmjs Faucet                                 ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux node:18.7-alpine AS sonr-faucet

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

ENV COSMJS_VERSION=0.28.11

RUN npm install @cosmjs/faucet@${COSMJS_VERSION} --global --production

ENV FAUCET_CONCURRENCY=4
ENV FAUCET_PORT=4500
ENV FAUCET_GAS_PRICE=0.0000usnr
# Prepared keys for determinism
ENV FAUCET_MNEMONIC="decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry"
ENV FAUCET_ADDRESS_PREFIX=idx
ENV FAUCET_TOKENS="usnr, snr"
ENV FAUCET_CREDIT_AMOUNT_STAKE=1000
ENV FAUCET_CREDIT_AMOUNT_TOKEN=100
ENV FAUCET_COOLDOWN_TIME=0

EXPOSE 4500

ENTRYPOINT [ "cosmos-faucet" ]

# ! ||-----------------------------------------------------------------------||
# ! ||                                  TKMS                                 ||
# ! ||-----------------------------------------------------------------------||
FROM --platform=linux rust:1.64.0-alpine AS tkms-builder

RUN apk update
RUN apk add libusb-dev=1.0.26-r0 musl-dev git

ENV LOCAL=/usr/local
ENV RUSTFLAGS=-Ctarget-feature=+aes,+ssse3
ENV TMKMS_VERSION=v0.12.2

WORKDIR /root
RUN git clone --branch ${TMKMS_VERSION} https://github.com/iqlusioninc/tmkms.git
WORKDIR /root/tmkms
RUN cargo build --release --features=softsign

# The production image starts here
FROM --platform=linux alpine AS tkms

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

COPY --from=tkms-builder /root/tmkms/target/release/tmkms ${LOCAL}/bin

ENTRYPOINT [ "tmkms" ]


# ! ||--------------------------------------------------------------------------------||
# ! ||                                  Sonrd Builder                                 ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux golang:1.19-alpine AS sonr-builder

ARG arch=x86_64

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers

# Download go dependencies
WORKDIR /root
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download


# Cosmwasm - Download correct libwasmvm version
RUN set -eux; \
    export ARCH=$(uname -m); \
    WASM_VERSION=$(go list -m all | grep github.com/CosmWasm/wasmvm | awk '{print $2}'); \
    if [ ! -z "${WASM_VERSION}" ]; then \
    wget -O /lib/libwasmvm_muslc.a https://github.com/CosmWasm/wasmvm/releases/download/${WASM_VERSION}/libwasmvm_muslc.${ARCH}.a; \
    fi; \
    go mod download;

# Copy the remaining files
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    GOWORK=off go build \
    -mod=readonly \
    -tags "netgo,ledger,muslc" \
    -ldflags \
    "-X github.com/cosmos/cosmos-sdk/version.Name="sonr" \
    -X github.com/cosmos/cosmos-sdk/version.AppName="sonrd" \
    -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=netgo,ledger,muslc \
    -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o /root/sonr/build/sonrd ./cmd/sonrd/main.go



# ! ||-----------------------------------------------------------------------------||
# ! ||                               Sonr Base Image                               ||
# ! ||-----------------------------------------------------------------------------||
FROM alpine AS sonr-base

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

# Copy sonrd binary and config
COPY --from=sonr-builder /root/sonr/build/sonrd /usr/local/bin/sonrd
COPY sonr.yml sonr.yml
COPY scripts scripts
ENV SONR_LAUNCH_CONFIG=/sonr.yml

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080


# ! ||----------------------------------------------------------------------------------||
# ! ||                               Sonr Standalone Node                               ||
# ! ||----------------------------------------------------------------------------------||

FROM alpine AS sonr-node

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

# Copy sonrd binary and config
COPY --from=sonr-builder /root/sonr/build/sonrd /usr/local/bin/sonrd
COPY sonr.yml sonr.yml
COPY scripts scripts
ENV SONR_LAUNCH_CONFIG=/sonr.yml

# Download, extract, and install the toml-cli binary
RUN apk add --update curl
RUN curl -LO https://github.com/gnprice/toml-cli/releases/latest/download/toml-0.2.3-x86_64-linux.tar.gz && \
    tar -xvf toml-0.2.3-x86_64-linux.tar.gz && \
    mv toml-0.2.3-x86_64-linux/toml /usr/local/bin && \
    rm toml-0.2.3-x86_64-linux.tar.gz && \
    rm -rf toml-0.2.3-x86_64-linux

# Setup localnet environment
RUN sh scripts/localnet.sh

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
EXPOSE 9090

CMD [ "sonrd", "start" ]


# ! ||---------------------------------------------------------------------------------||
# ! ||                               Sonr Validator Node                               ||
# ! ||---------------------------------------------------------------------------------||

FROM sonr-base AS sonr-validator

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

# Download, extract, and install the toml-cli binary
RUN apk add --update curl wget
RUN curl -LO https://github.com/gnprice/toml-cli/releases/latest/download/toml-0.2.3-x86_64-linux.tar.gz && \
    tar -xvf toml-0.2.3-x86_64-linux.tar.gz && \
    mv toml-0.2.3-x86_64-linux/toml /usr/local/bin && \
    rm toml-0.2.3-x86_64-linux.tar.gz && \
    rm -rf toml-0.2.3-x86_64-linux

# Install Doppler CLI
RUN wget -q -t3 'https://packages.doppler.com/public/cli/rsa.8004D9FF50437357.key' -O /etc/apk/keys/cli@doppler-8004D9FF50437357.rsa.pub && \
    echo 'https://packages.doppler.com/public/cli/alpine/any-version/main' | tee -a /etc/apk/repositories && \
    apk add doppler

# Setup build args
ARG DOPPLER_TOKEN

# Create encrypted snapshot for high availability
RUN doppler secrets download doppler.encrypted.json

# Initialize the node
RUN doppler run --command='echo $MNEMONIC | sonrd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover'
RUN doppler run --command='sonrd init ${MONIKER} --chain-id ${CHAIN_ID} --home /root/.sonr'
RUN doppler run --command='sonrd add-genesis-account $KEY ${BALANCE} --keyring-backend $KEYRING'
RUN doppler run --command='sonrd gentx $KEY ${STAKE} --keyring-backend $KEYRING --chain-id $CHAIN_ID --moniker ${MONIKER} --website ${WEBSITE} --note ${NOTE} --security-contact ${SECURITY_CONTACT}'
RUN doppler run --command='sonrd collect-gentxs'

# Update config.toml
RUN doppler run --command='toml set $HOME/.sonr/config/config.toml rpc.laddr ${RPC_LADDR} > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml'
RUN doppler run --command='toml set $HOME/.sonr/config/app.toml grpc.address ${GRPC_ADDRESS} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml'
RUN doppler run --command='toml set $HOME/.sonr/config/app.toml api.enable ${API_ENABLE} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml'
RUN doppler run --command='toml set $HOME/.sonr/config/app.toml api.swagger ${API_SWAGGER} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml'
RUN doppler run --command='toml set $HOME/.sonr/config/app.toml api.address ${API_ADDRESS} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml'
RUN doppler run --command='toml set $HOME/.sonr/config/app.toml minimum-gas-prices ${MINIMUM_GAS_PRICES} > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml'

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
EXPOSE 9090

CMD [ "doppler", "run", "--", "sonrd", "start" ]
