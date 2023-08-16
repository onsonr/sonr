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


# ! ||----------------------------------------------------------------------------------||
# ! ||                               Sonr Standalone Node                               ||
# ! ||----------------------------------------------------------------------------------||
FROM --platform=linux alpine AS sonr-node

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


# ! ||-----------------------------------------------------------------------------||
# ! ||                               Sonr Base Image                               ||
# ! ||-----------------------------------------------------------------------------||
FROM --platform=linux alpine AS sonr-base

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
