VERSION 0.7
FROM golang:1.21.5-alpine
WORKDIR /app

# install dependencies
deps:
    RUN apk add --no-cache git
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# build - builds the flavored ipfs gateway
build:
    BUILD +hway
    BUILD +sonrd

# hway - builds the flavored ipfs gateway
hway:
    FROM +deps
    COPY . .
    ARG goos=linux
    ARG goarch=amd64
    ENV GOOS=$goos
    ENV GOARCH=$goarch
    RUN go build -o /app/hway ./cmd/hway
    SAVE ARTIFACT /app/hway AS LOCAL bin/hway

# sonrd - builds the flavored ipfs gateway
sonrd:
    FROM +deps
    COPY . .
    ARG goos=linux
    ARG goarch=amd64
    ENV GOOS=$goos
    ENV GOARCH=$goarch
    RUN go build -o /app/sonrd ./cmd/sonrd
    SAVE ARTIFACT /app/sonrd AS LOCAL bin/sonrd
