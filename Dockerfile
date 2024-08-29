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
FROM alpine:3.16

COPY --from=go-builder /code/build/sonrd /usr/bin/sonrd
COPY scripts/test_node.sh /usr/local/bin/test_node.sh

# Install dependencies
RUN apk add --no-cache curl make bash jq sed

WORKDIR /opt

# Set environment variables
ENV CHAIN_ID=local-1 \
    HOME_DIR=/root/.core \
    BINARY=sonrd \
    DENOM=usnr \
    KEYRING=test \
    KEY=user1 \
    KEY2=user2 \
    CLEAN=true \
    RPC=26657 \
    REST=1317 \
    PROFF=6060 \
    P2P=26656 \
    GRPC=9090 \
    GRPC_WEB=9091 \
    ROSETTA=8080 \
    BLOCK_TIME=5s

# Expose ports
EXPOSE 1317 26656 26657 9090 9091 8080

# Create entrypoint script
RUN echo '#!/bin/sh' > /usr/local/bin/entrypoint.sh && \
    echo 'set -e' >> /usr/local/bin/entrypoint.sh && \
    echo 'bash /usr/local/bin/test_node.sh' >> /usr/local/bin/entrypoint.sh && \
    chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
