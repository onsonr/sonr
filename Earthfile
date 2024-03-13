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

# gen: go generate prisma
gen:
    FROM +deps
    COPY . .
    RUN cd internal/prisma/hway && go run github.com/steebchen/prisma-client-go generate
    SAVE ARTIFACT internal/prisma/hway/db AS LOCAL internal/prisma/hway/db
    RUN cd internal/prisma/indexer && go run github.com/steebchen/prisma-client-go generate
    SAVE ARTIFACT internal/prisma/indexer/db AS LOCAL internal/prisma/indexer/db
    RUN cd internal/prisma/matrix && go run github.com/steebchen/prisma-client-go generate
    SAVE ARTIFACT internal/prisma/matrix/db AS LOCAL internal/prisma/matrix/db


# build - builds the flavored ipfs gateway
build:
    BUILD +hway
    BUILD +sonrd

# hway - builds the flavored ipfs gateway
hway:
    FROM +gen
    ARG goos=linux
    ARG goarch=amd64
    ENV GOOS=$goos
    ENV GOARCH=$goarch
    RUN go build -o /app/hway ./cmd/hway
    SAVE ARTIFACT /app/hway AS LOCAL bin/hway

# sonrd - builds the flavored ipfs gateway
sonrd:
    FROM +gen
    ARG goos=linux
    ARG goarch=amd64
    ENV GOOS=$goos
    ENV GOARCH=$goarch
    RUN go build -o /app/sonrd ./cmd/sonrd
    SAVE ARTIFACT /app/sonrd AS LOCAL bin/sonrd
