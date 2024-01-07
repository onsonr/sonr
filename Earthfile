# sonrhq/chain: Earthfile
# ---------------------------------------------------------------------
VERSION 0.7
PROJECT sonrhq/testnet-1

FROM golang:1.21.5-alpine
IMPORT ../identity AS identity
IMPORT ../service AS service
WORKDIR /chain
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
    RUN go install github.com/a-h/templ/cmd/templ@latest
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

# runner - builds the docker image
runner:
    FROM gcr.io/distroless/static-debian11
    ARG tag=latest
    COPY +build/sonrd /usr/local/bin/sonrd
    EXPOSE 26657
    EXPOSE 1317
    EXPOSE 26656
    EXPOSE 9090
    SAVE IMAGE --push sonrhq/sonrd:$tag ghcr.io/sonrhq/sonrd:$tag sonrd:$tag

# test - runs tests on x/identity and x/service
test:
    FROM +deps
    BUILD identity+test
    BUILD service+test

# breaking - runs tests on x/identity and x/service with breaking changes
breaking:
    BUILD identity+breaking
    BUILD service+breaking

# templates - runs protogen, and templ generate on all modules and root
templates:
    LOCALLY
    RUN templ generate
    BUILD identity+templates
    BUILD service+templates
