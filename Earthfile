VERSION 0.7
PROJECT sonrhq/testnet-1

FROM golang:1.21.5-alpine
WORKDIR /app

# install dependencies
deps:
    FROM +base
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

# docker - builds the docker image
docker:
    FROM +build
    COPY +build/sonrd /usr/local/bin/sonrd
    COPY ./networks/local/entrypoint.sh ./entrypoint.sh
    RUN chmod +x /usr/local/bin/sonrd
    SAVE IMAGE sonrd:latest

# matrix-config - builds the matrix configuration
matrix-config:
    LOCALLY
    RUN mkdir -p ./tmp/config

    FROM matrixdotorg/dendrite-monolith:latest
    ARG serverName=localhost
    COPY ./networks/local/matrix-config.yaml ./matrix-config.yaml
    SAVE ARTIFACT matrix-config.yaml AS LOCAL matrix-config.yaml

# matrix-local - setups the local matrix server
matrix-local:
    LOCALLY
    RUN git clone https://github.com/matrix-org/dendrite ./tmp/dendrite
    RUN cd ./tmp/dendrite && go build -o bin/ ./cmd/...
    RUN ./tmp/dendrite/bin/generate-keys --private-key ./tmp/dendrite/matrix_key.pem
    RUN ./tmp/dendrite/bin/generate-keys --tls-cert ./tmp/dendrite/server.crt --tls-key ./tmp/dendrite/server.key
    RUN cp ./tmp/dendrite/dendrite-sample.yaml ./tmp/dendrite/dendrite.yaml
    RUN ./tmp/dendrite/bin/dendrite --tls-cert ./tmp/dendrite/server.crt --tls-key ./tmp/dendrite/server.key --config ./tmp/dendrite/dendrite.yaml

