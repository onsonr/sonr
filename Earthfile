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
    RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest
    RUN go install github.com/a-h/templ/cmd/templ@latest
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum


# test - runs tests on x/identity and x/service
test:
    FROM +deps
    COPY . .
    RUN go test -v ./...

# -----------------
# [BUILD Commands]
# -----------------

# build-daemon - builds binary
daemon:
    FROM +deps
    WORKDIR /app
    COPY . .
    ARG version=$EARTHLY_GIT_REFS
    ARG commit=$EARTHLY_BUILD_SHA

    COPY . .
    RUN  go build -ldflags "-X main.Version=$version -X main.Commit=$commit" -o sonrd ./cmd/sonrd/main.go
    SAVE ARTIFACT sonrd AS LOCAL bin/sonrd

# build-daemon - builds binary
studio:
    LOCALLY
    RUN  cd ./cmd/studio && wails build

# docker-daemon - builds the binary cmd docker image
docker:
    FROM +daemon
    ARG tag=latest
    ARG commit=$EARTHLY_BUILD_SHA
    ARG version=$EARTHLY_GIT_REFS
    COPY +daemon/sonrd sonrd
    ENTRYPOINT [ "/app/sonrd" ]
    EXPOSE 26657
    EXPOSE 1317
    EXPOSE 26656
    EXPOSE 9090
    SAVE IMAGE sonrhq/sonrd:$tag sonrd:$tag

# -------------------
# [PROTOBUF Commands]
# -------------------

# generate - generates all code from proto files
generate:
    FROM +deps
    COPY . .
    RUN sh ./scripts/protogen-orm.sh
    SAVE ARTIFACT sonr AS LOCAL api
    SAVE ARTIFACT proto AS LOCAL proto
    LOCALLY
    RUN templ generate


# breaking - runs tests on x/identity and x/service with breaking changes
breaking:
    FROM +deps
    COPY . .
    RUN cd proto && buf breaking --against buf.build/sonrhq/sonr
