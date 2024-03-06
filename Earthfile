# -----------------------------------------------------
# Purpose - Setting up local development environment
# -----------------------------------------------------

VERSION 0.7
PROJECT sonrhq/testnet-1

# install dependencies
deps:
    FROM golang:1.21.5-alpine
    WORKDIR /app
    RUN apk add --no-cache git
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

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

# current-node - copies local files to the current build
current-node:
    FROM alpine:3.14
    COPY ./bin/sonrd /usr/local/bin/sonrd
    RUN chmod +x /usr/local/bin/sonrd
    SAVE IMAGE sonrd:latest

# current-hway - copies local files to the current build
current-hway:
    FROM alpine:3.14
    COPY ./bin/hway /usr/local/bin/hway
    RUN chmod +x /usr/local/bin/hway
    SAVE IMAGE hway:latest
