VERSION 0.7
PROJECT sonrhq/sonr-testnet-0

ARG branch=master

IMPORT github.com/sonrhq/chain:$branch AS sonrd


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

# testnet - installs the latest version of infisical and configures it for the testnet
testnet:
    FROM +base
    RUN curl -1sLf 'https://dl.cloudsmith.io/public/infisical/infisical-cli/setup.alpine.sh' | bash
    RUN apk add infisical

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
