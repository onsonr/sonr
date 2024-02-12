VERSION 0.7
PROJECT sonrhq/testnet-1
FROM golang:1.21.5-alpine

WORKDIR /app

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
    openssl \
    util-linux
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum


# -------------------
# [Network Services]
# -------------------

# build - builds the flavored ipfs gateway
build:
    FROM +deps
    ARG goos=linux
    ARG goarch=amd64
    ENV GOOS=$goos
    ENV GOARCH=$goarch
    COPY . .
    RUN go build -o /app/sonrd ./cmd/sonrd
    SAVE ARTIFACT /app/sonrd AS LOCAL bin/sonrd

# docker - builds the docker image
docker:
    FROM alpine:3.14
    COPY +build/sonrd /usr/local/bin/sonrd
    COPY ./networks/local/entrypoint.sh ./entrypoint.sh
    RUN chmod +x /usr/local/bin/sonrd
    SAVE IMAGE sonrd:latest
