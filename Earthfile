# sonrhq/chain: Earthfile
# ---------------------------------------------------------------------
VERSION 0.7
PROJECT sonrhq/testnet-1

FROM golang:1.21-alpine3.18
IMPORT github.com/sonrhq/identity AS identity
IMPORT github.com/sonrhq/service AS service
IMPORT ./rails AS rails
IMPORT ./deploy AS deploy
WORKDIR /chain
# ---------------------------------------------------------------------

# Initial Setup
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


# ---------------------------------------------------------------------

# dev - Starts a development chain
dev:
    BUILD +docker
    LOCALLY
    RUN docker compose -f ./deploy/docker-compose.dev.yml up --quiet-pull --build --remove-orphans

# deps - downloads dependencies
deps:
    COPY go.mod go.sum ./
    RUN go mod download
    # Output these back in case go mod download changes them.
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# build - builds binary
build:
    FROM +deps
    ARG version=$EARTHLY_GIT_REFS
    ARG commit=$EARTHLY_BUILD_SHA

    COPY . .
    RUN  go build -ldflags "-X main.Version=$version -X main.Commit=$commit" -o bin/sonrd ./cmd/sonrd/main.go
    SAVE ARTIFACT bin/sonrd AS LOCAL bin/sonrd

# docker - builds the docker image
docker:
    ARG tag=latest
    COPY +build/sonrd /usr/local/bin/sonrd
    EXPOSE 26657
    EXPOSE 1317
    EXPOSE 26656
    EXPOSE 9090
    ENTRYPOINT ["sonrd"]
    SAVE IMAGE sonrhq/sonrd:$tag ghcr.io/sonrhq/sonrd:$tag

# runner - Creates a containerized node with preconfigured keys
runner:
    ARG tag=latest
    ARG --secret --required infisicalToken
    ARG --required mountVolume

    ENV INFISICAL_TOKEN=$infisicalToken
    VOLUME $mountVolume:/root/.sonr

    RUN apt-get update && apt-get install -y bash curl && curl -1sLf \
    'https://dl.cloudsmith.io/public/infisical/infisical-cli/setup.deb.sh' | bash \
    && apt-get update && apt-get install -y infisical

    ARG chainId=sonr-testnet-1
    ARG enableSwagger=true
    ARG enableAPI=true
    ARG faucetBalance=1000000snr
    ARG faucetKey=bob
    ARG genesisBalance=10000000snr
    ARG keyringBackend=test
    ARG moniker=austin
    ARG validatorKey=alice
    ARG vestingAmount=1000000snr

    COPY +build/sonrd .

    RUN ./sonrd config set client chain-id $chainId
    RUN ./sonrd config set client keyring-backend $keyringBackend
    RUN ./sonrd config set app api.enable $enableAPI
    RUN ./sonrd config set app api.swagger $enableSwagger
    RUN ./sonrd keys add $validatorKey --recover $validatorMnemonic
    RUN ./sonrd keys add $faucetKey --recover $faucetMnemonic

    RUN ./sonrd init $moniker --chain-id $chainId --default-denom snr --home $home
    RUN ./sonrd genesis add-genesis-account $validatorKey $genesisBalance --chain-id $chainId --home $home
    RUN ./sonrd genesis add-genesis-account $faucetKey $faucetBalance --chain-id $chainId --home $home
    RUN ./sonrd genesis gentx $validatorKey $vestingAmount --chain-id $chainId --home $home
    RUN ./sonrd genesis collect-gentxs --home $home

    EXPOSE 26657
    EXPOSE 1317
    EXPOSE 26656
    EXPOSE 9090

    CMD ["/chain/sonrd start"]
    SAVE IMAGE --push sonrhq/sonrd:$tag-standalone ghcr.io/sonrhq/sonrd:$tag-standalone

# generate - generates all code from proto files
generate:
    LOCALLY
    RUN make proto-gen
    FROM +deps
    COPY . .
    RUN sh ./scripts/protogen-orm.sh
    SAVE ARTIFACT sonrhq/identity AS LOCAL api
    SAVE ARTIFACT proto AS LOCAL proto
    RUN sh ./scripts/protocgen-docs.sh
    SAVE ARTIFACT docs AS LOCAL docs

# test - runs tests on x/identity and x/service
test:
    FROM +deps
    BUILD identity+test
    BUILD service+test

# breaking - runs tests on x/identity and x/service with breaking changes
breaking:
    FROM +deps
    BUILD identity+breaking
    BUILD service+breaking
