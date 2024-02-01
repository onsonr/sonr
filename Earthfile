# sonrhq/chain: Earthfile
# ---------------------------------------------------------------------
VERSION 0.7
PROJECT sonrhq/testnet-1
FROM golang:1.21.5-alpine

# ---------------------------------------------------------------------

# deps - downloads dependencies
deps:
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
    nodejs \
    npm \
    openssl \
    util-linux
    COPY go.mod go.sum ./
    RUN go mod download
    RUN npm install -g swagger-combine
    RUN npm install @bufbuild/buf
    FROM ghcr.io/cosmos/proto-builder:0.14.0
    RUN go install github.com/kollalabs/protoc-gen-openapi@latest
    RUN go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm@latest
	RUN go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm-proto@latest
    RUN go install github.com/a-h/templ/cmd/templ@latest
    RUN go install github.com/go-task/task/v3/cmd/task@latest
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# -------------------
# [PROTOBUF Commands]
# -------------------

# generate - generates all code from proto files
generate:
    FROM +deps
    COPY . .
    RUN task gen:proto-orm
    SAVE ARTIFACT sonrhq AS LOCAL api
    SAVE ARTIFACT proto AS LOCAL proto


# breaking - runs tests on x/identity and x/service with breaking changes
breaking:
    FROM +deps
    COPY . .
    RUN cd proto && buf breaking --against buf.build/sonrhq/sonr

# build - builds and configures monolithic dendrite matrix homeserver
build:
    FROM matrixdotorg/dendrite-monolith:latest
    ARG serverName=matrix.sonr.run
    SAVE IMAGE --push sonrhq/dendrite:latest

# init - generates private key for matrix homeserver
init:
    FROM matrixdotorg/dendrite-monolith:latest
    ARG serverName=matrix.sonr.run
    ARG psqlConnURI=postgres://postgres:pwd@postgres:5432/postgres?sslmode=disable
    RUN /usr/bin/generate-keys -private-key matrix_key.pem
    SAVE ARTIFACT matrix_key.pem AS LOCAL matrix_key.pem
    RUN /usr/bin/generate-config -server $serverName -db $psqlConnURI > dendrite.yaml
    SAVE ARTIFACT dendrite.yaml AS LOCAL dendrite.yaml
