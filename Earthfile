# sonrhq/chain: Earthfile
# ---------------------------------------------------------------------
VERSION 0.7
PROJECT sonrhq/testnet-1

FROM golang:1.21-alpine3.18
IMPORT github.com/sonrhq/identity AS identity
IMPORT github.com/sonrhq/service AS service
IMPORT ./cmd AS cmd
IMPORT ./rails AS rails
IMPORT ./deploy AS deploy
WORKDIR /chain
# ---------------------------------------------------------------------

# Initial Setup
RUN apk add --update --no-cache \
    bash \
    binutils \
    ca-certificates \
    coreutils \
    curl \
    findutils \
    g++ \
    git \
    grep \
    make \
    openssl \
    util-linux


# ---------------------------------------------------------------------

# dev - Starts a development chain
dev:
    BUILD +docker
    LOCALLY
    RUN docker compose -f ./deploy/docker-compose.dev.yml up --quiet-pull --build --remove-orphans

# deps - downloads dependencies
deps:
    COPY go.mod go.sum ./
    RUN go mod download
    # Output these back in case go mod download changes them.
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# build - builds binary
build:
    FROM +deps
    ARG version=$EARTHLY_GIT_REFS
    ARG commit=$EARTHLY_BUILD_SHA

    COPY . .
    RUN  go build -ldflags "-X main.Version=$version -X main.Commit=$commit" -o bin/sonrd ./cmd/sonrd/main.go
    SAVE ARTIFACT bin/sonrd AS LOCAL /usr/bin/sonrd

# docker - builds the docker image
docker:
    ARG tag=latest
    COPY +build/sonrd /usr/local/bin/sonrd
    EXPOSE 26657
    EXPOSE 1317
    EXPOSE 26656
    EXPOSE 9090
    SAVE IMAGE sonrhq/sonrd:$tag ghcr.io/sonrhq/sonrd:$tag sonrd-runner:$tag

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

# scripts - Stores the repository scripts in the local cache
scripts:
    COPY ./scripts ./src
    SAVE ARTIFACT src AS LOCAL scripts

# test - runs tests on x/identity and x/service
test:
    FROM +deps
    BUILD identity+test
    BUILD service+test

# breaking - runs tests on x/identity and x/service with breaking changes
breaking:
    FROM +deps
    BUILD identity+breaking
    BUILD service+breaking

# rails - builds the required rails to run sonr
rails:
    ARG --required tunnelToken
    ARG --required infisicalToken
    BUILD rails+ipfs
    BUILD rails+faucet  --infisicalToken $infisicalToken
    BUILD rails+dendrite  --infisicalToken $infisicalToken
    BUILD rails+cloudflared --tunnelToken $tunnelToken
