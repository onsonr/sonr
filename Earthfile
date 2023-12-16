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

# clone - Clones the dependencies as git submodules
clone:
    FROM +repo
    WORKDIR /sonr
    GIT CLONE git@github.com:sonrhq/identity.git identity
    SAVE ARTIFACT identity AS LOCAL identity
    GIT CLONE git@github.com:sonrhq/service.git service
    SAVE ARTIFACT service AS LOCAL service
    GIT CLONE git@github.com:sonrhq/chain.git chain
    SAVE ARTIFACT chain AS LOCAL chain
    GIT CLONE git@github.com:sonrhq/rails.git rails
    SAVE ARTIFACT rails AS LOCAL rails

# generate - generates all code from proto files
generate:
    LOCALLY
    RUN make proto-gen
    FROM +deps
    COPY . .
    RUN sh ./scripts/protogen-orm.sh
    SAVE ARTIFACT sonrhq/identity AS LOCAL api
    SAVE ARTIFACT proto AS LOCAL proto
    RUN sh ./scripts/protocgen-docs.sh
    SAVE ARTIFACT docs AS LOCAL docs

# test - runs all tests
test:
    FROM +repo
    COPY . .
	RUN go test -v ./...
