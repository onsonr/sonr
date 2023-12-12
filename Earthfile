VERSION 0.7
PROJECT sonrhq/sonr-testnet-0

FROM golang:1.21-alpine3.17
RUN apk add --update --no-cache \
    bash \
    bash-completion \
    binutils \
    ca-certificates \
    clang-extra-tools \
    coreutils \
    curl \
    findutils \
    g++ \
    git \
    grep \
    gum \
    jq \
    less \
    make \
    nodejs \
    npm \
    openssl \
    util-linux

WORKDIR /sonr
COPY . .

# gomod - downloads and caches all dependencies for earthly. go.mod and go.sum will be updated locally.
gomod:
    FROM +base
    RUN go work sync
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum
    SAVE ARTIFACT go.work AS LOCAL go.work
    SAVE ARTIFACT go.work.sum AS LOCAL go.work.sum


faucet:
    FROM node:18.7-alpine
    ARG cosmosjsVersion=0.28.11
    RUN npm install @cosmjs/faucet@$cosmosjsVersion --global --production
    COPY ./scripts/start-faucet.sh /usr/local/bin/faucet
    RUN chmod +x /usr/local/bin/faucet
    EXPOSE 4500
    ENTRYPOINT ["faucet"]
    SAVE IMAGE sonrhq/faucet:latest


# deps - downloads and caches all dependencies for earthly. go.mod and go.sum will be updated locally.
deps:
    FROM +base
    RUN npm install -g swagger-combine
    RUN npm install @bufbuild/buf
    FROM ghcr.io/cosmos/proto-builder:0.14.0
    RUN go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm@latest
	RUN go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm-proto@latest

# proto - generates all code from proto files
proto:
    FROM +deps
    COPY . .
    RUN sh ./scripts/protocgen.sh
    SAVE ARTIFACT common/crypto/*.go AS LOCAL common/crypto/*.go

# test - runs all tests
test:
    FROM +base
    COPY . .
    RUN go test -v ./...
