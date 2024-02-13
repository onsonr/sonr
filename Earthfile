VERSION 0.7
PROJECT sonrhq/testnet-1
FROM golang:1.21.5-alpine

RUN apk add --update --no-cache \
bash \
binutils \
ca-certificates \
coreutils \
curl \
git \
make

WORKDIR /app

# ---------------------------------------------------------------------
# -------------------
# [Network Services]
# -------------------

# build - builds the flavored ipfs gateway
build:
    FROM +base
    ARG goos=linux
    ARG goarch=amd64
    ENV GOOS=$goos
    ENV GOARCH=$goarch
    COPY . .
    RUN go build -o /app/sonrd ./cmd/sonrd
    SAVE ARTIFACT /app/sonrd AS LOCAL bin/sonrd

# docker - builds the docker image
docker:
    COPY +build/sonrd /usr/local/bin/sonrd
    COPY ./networks/local/entrypoint.sh ./entrypoint.sh
    RUN chmod +x /usr/local/bin/sonrd
    SAVE IMAGE sonrd:latest
