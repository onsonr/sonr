# sonrhq/chain: Earthfile
# ---------------------------------------------------------------------
VERSION 0.7
PROJECT sonrhq/testnet-1

FROM golang:1.21-alpine3.18
IMPORT github.com/sonrhq/identity:main AS identity
IMPORT github.com/sonrhq/service:main AS service
IMPORT ./cmd AS cmd
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
    RUN  go build -ldflags "-X main.Version=$version -X main.Commit=$commit" -o /usr/bin/sonrd ./cmd/sonrd/main.go
    SAVE ARTIFACT /usr/bin/sonrd AS LOCAL bin/sonrd

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
    BUILD deploy+build-ipfs
    BUILD deploy+build-faucet
    BUILD deploy+build-dendrite
    BUILD deploy+build-standalone
