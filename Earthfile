VERSION 0.7
PROJECT sonrhq/testnet-1

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
    jq \
    less \
    make \
    nodejs \
    npm \
    openssl \
    util-linux

# repo - Creates repository container environment
repo:
	FROM +base
    ARG EARTHLY_GIT_BRANCH

    GIT CLONE --branch $EARTHLY_GIT_BRANCH git@github.com:sonrhq/sonr.git sonr
    CACHE --sharing shared sonr
    WORKDIR /sonr

    COPY ./go.mod ./go.sum ./
    RUN go mod download
    CACHE --sharing shared /go/pkg/mod

    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# test - runs all tests
test:
    FROM +repo
	RUN go test -v ./...
