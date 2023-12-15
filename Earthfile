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

WORKDIR /sonr

# setup - clones all repositories
setup:
	LOCALLY
	GIT CLONE git@github.com:sonrhq/chain.git chain
	GIT CLONE git@github.com:sonrhq/identity.git identity
	GIT CLONE git@github.com:sonrhq/service.git service
	GIT CLONE git@github.com:sonrhq/rails.git rails

# gomod - downloads and caches all dependencies for earthly. go.mod and go.sum will be updated locally.
gomod:
    FROM +base
    COPY ./go.mod ./go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# test - runs all tests
test:
    FROM +gomod
    COPY . .
	RUN go test -v ./...
