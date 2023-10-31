#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
PATH := $(PATH):/usr/local/go/bin
LEDGER_ENABLED ?= false
DOCKER := $(shell which docker)
E2E_UPGRADE_VERSION := "v17"
#SHELL := /bin/bash

GH_IMAGE=ghcr.io/sonr-io/sonrd
DO_IMAGE=registry.digitalocean.com/sonrhq/sonrd

GO_VERSION := $(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f 2)
GO_MODULE := $(shell cat go.mod | grep "module " | cut -d ' ' -f 2)
GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
protoVer=0.11.2
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

ifeq (,$(VERSION))
  VERSION := $(shell git describe --tags)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

LEDGER_ENABLED ?= false
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf:1.0.0-rc8
BUILDDIR ?= $(CURDIR)/build
HTTPS_GIT := https://github.com/sonrhq/core.git

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc cleveldb
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=sonr \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=sonrd \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif


all: install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/sonrd

build:
	go build $(BUILD_FLAGS) -o bin/sonrd ./cmd/sonrd

docker-build:
	docker build -t $(GH_IMAGE) -t $(DO_IMAGE) -f Dockerfile .
	docker push $(GH_IMAGE) --all-tags

lint:
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' -not -name '*.gw.go' | xargs go run mvdan.cc/gofumpt -w .
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' -not -name '*.gw.go' | xargs go run github.com/client9/misspell/cmd/misspell -w
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' -not -name '*.gw.go' | xargs go run golang.org/x/tools/cmd/goimports -w -local github.com/notional-labs/centauri
.PHONY: lint

###############################################################################
###                                  Proto                                  ###
###############################################################################

protoVer=0.11.6
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
containerProtoGen=proto-gen-$(protoVer)
containerProtoFmt=proto-fmt-$(protoVer)

proto-all: proto-format proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protocgen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=main

.PHONY: proto-all proto-gen proto-format proto-lint proto-check-breaking

###############################################################################
###                             Interchain test                             ###
###############################################################################

# Executes start chain tests via interchaintest
ictest-start-cosmos:
	cd tests/interchaintest && go test -race -v -run TestStartComposable .

ictest-validator:
	cd tests/interchaintest && go test -race -v -run TestValidator .

# Executes start chain tests via interchaintest
ictest-start-polkadot:
	cd tests/interchaintest && go test -timeout=25m -race -v -run TestPolkadotComposableChainStart .

# Executes IBC tests via interchaintest
ictest-ibc:
	cd tests/interchaintest && go test -timeout=25m -race -v -run TestComposablePicassoIBCTransfer .

# Executes Basic Upgrade Chain tests via interchaintest
ictest-upgrade:
	cd tests/interchaintest && go test -timeout=25m -race -v -run TestComposableUpgrade .

# Executes all tests via interchaintest after compling a local image as juno:local
ictest-all: ictest-start-cosmos ictest-start-polkadot ictest-ibc

# Executes push wasm client tests via interchaintest
ictest-push-wasm:
	cd tests/interchaintest && go test -race -v -run TestPushWasmClientCode .

.PHONY: ictest-start-cosmos ictest-start-polkadot ictest-ibc ictest-push-wasm ictest-all

###############################################################################
###                                Release                                  ###
###############################################################################
GORELEASER_IMAGE := ghcr.io/goreleaser/goreleaser-cross:v$(GO_VERSION)
COSMWASM_VERSION := $(shell go list -m github.com/CosmWasm/wasmvm | sed 's/.* //')

release:
	goreleaser release --clean

release-dry-run:
	docker run \
		--rm \
		-e COSMWASM_VERSION=$(COSMWASM_VERSION) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/sonrd \
		-w /go/src/sonrd \
		$(GORELEASER_IMAGE) \
		release \
		--clean \
		--skip-publish

release-snapshot:
	docker run \
		--rm \
		-e COSMWASM_VERSION=$(COSMWASM_VERSION) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/sonrd \
		-w /go/src/sonrd \
		$(GORELEASER_IMAGE) \
		release \
		--clean \
		--snapshot \
		--skip-validate \
		--skip-publish


###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	@echo "--> Running linter"
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout=10m
	@docker run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md"

format:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./... --fix
	@go run mvdan.cc/gofumpt -l -w x/ app/ ante/ tests/
	@docker run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md" --fix

mdlint:
	@echo "--> Running markdown linter"
	@docker run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md"

markdown:
	@docker run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md" --fix



###############################################################################
###                                Protobuf                                 ###
###############################################################################

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main
