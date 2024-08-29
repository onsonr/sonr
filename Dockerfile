FROM golang:1.22-alpine AS go-builder

SHELL ["/bin/sh", "-ecuxo", "pipefail"]

RUN apk add --no-cache ca-certificates build-base git

WORKDIR /code

COPY go.mod go.sum ./
RUN set -eux; \
  export ARCH=$(uname -m); \
  WASM_VERSION=$(go list -m all | grep github.com/CosmWasm/wasmvm || true); \
  if [ ! -z "${WASM_VERSION}" ]; then \
  WASMVM_REPO=$(echo $WASM_VERSION | awk '{print $1}');\
  WASMVM_VERS=$(echo $WASM_VERSION | awk '{print $2}');\
  wget -O /lib/libwasmvm_muslc.a https://${WASMVM_REPO}/releases/download/${WASMVM_VERS}/libwasmvm_muslc.$(uname -m).a;\
  fi; \
  go mod download;

# Copy over code
COPY . /code

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
# then log output of file /code/bin/sonrd
# then ensure static linking
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build \
  && file /code/build/sonrd \
  && echo "Ensuring binary is statically linked ..." \
  && (file /code/build/sonrd | grep "statically linked")

# --------------------------------------------------------
FROM debian:11-slim

COPY --from=go-builder /code/build/sonrd /usr/bin/sonrd


# Install dependencies for Debian 11
RUN apt-get update && apt-get install -y \
  curl \
  make \
  bash \
  jq \
  sed \
  && rm -rf /var/lib/apt/lists/*

COPY scripts/test_node.sh /usr/bin/test_node.sh 

WORKDIR /opt

# rest server, tendermint p2p, tendermint rpc
EXPOSE 1317 26656 26657

ENTRYPOINT ["/usr/bin/sonrd"]
