VERSION 0.7
PROJECT sonrhq/sonrd

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

    GIT CLONE --branch $EARTHLY_GIT_BRANCH git@github.com:sonrhq/service.git service
    CACHE --sharing shared service
    WORKDIR /service

    COPY ./go.mod ./go.sum ./
    RUN go mod download
    CACHE --sharing shared /go/pkg/mod
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# deps - downloads and caches all dependencies for earthly. go.mod and go.sum will be updated locally.
deps:
    FROM +base
    RUN npm install -g swagger-combine
    RUN npm install @bufbuild/buf
    FROM ghcr.io/cosmos/proto-builder:0.14.0
    RUN go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm@latest
	RUN go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm-proto@latest
    SAVE IMAGE deps

# generate - generates all code from proto files
generate:
    LOCALLY
    RUN make proto-all
    FROM +deps
    COPY . .
    RUN sh ./scripts/protogen-orm.sh
    SAVE ARTIFACT sonrhq/service AS LOCAL api
    SAVE ARTIFACT proto AS LOCAL proto
    RUN sh ./scripts/protocgen-docs.sh
    SAVE ARTIFACT docs AS LOCAL docs

# lint - lints the protobuf files
lint:
    LOCALLY
    RUN make proto-format
    RUN make proto-lint

# test - runs all tests
test:
    FROM +repo
    COPY . .
	RUN go test -v ./...

# breaking - runs buf change detection
breaking:
    FROM +deps
    COPY . .
    RUN cd proto && buf breaking --against buf.build/sonrhq/service
