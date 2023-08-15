VERSION 0.7
FROM golang:1.19-alpine3.17
WORKDIR /go-workdir

build:
    RUN apk update && apk add ca-certificates gcc g++ build-base linux-headers

    COPY go.mod go.sum ./
    RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

    RUN set -eux; \
    export ARCH=$(uname -m); \
    WASM_VERSION=$(go list -m all | grep github.com/CosmWasm/wasmvm | awk '{print $2}'); \
    if [ ! -z "${WASM_VERSION}" ]; then \
    wget -O /lib/libwasmvm_muslc.a https://github.com/CosmWasm/wasmvm/releases/download/${WASM_VERSION}/libwasmvm_muslc.${ARCH}.a; \
    fi; \
    go mod download;

    COPY . .
    RUN go build -mod=readonly \
        -tags "netgo,ledger,muslc" \
        -ldflags \
        "-X github.com/cosmos/cosmos-sdk/version.Name="sonr" \
        -X github.com/cosmos/cosmos-sdk/version.AppName="sonrd" \
        -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
        -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
        -X github.com/cosmos/cosmos-sdk/version.BuildTags=netgo,ledger,muslc \
        -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
        -trimpath -o build/sonrd ./cmd/sonrd/main.go
    SAVE ARTIFACT build/sonrd

docker:
    COPY +build/sonrd .
    ENTRYPOINT ["/go-workdir/sonrd"]
    SAVE IMAGE sonrd:latest
