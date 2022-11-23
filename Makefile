SHELL=/bin/bash

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
SCRIPTS_DIR=$(ROOT_DIR)/scripts
PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
export VERSION := $(shell echo $(shell git describe --always --match "v*") | sed 's/^v//')
export TMVERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
export COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= false
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
MOCKS_DIR = $(CURDIR)/tests/mocks
HTTPS_GIT := https://github.com/sonr-io/sonr.git
DOCKER := $(shell which docker)
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)
# RocksDB is a native dependency, so we don't assume the library is installed.
# Instead, it must be explicitly enabled and we warn when it is not.
ENABLE_ROCKSDB ?= true

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

ifeq (secp,$(findstring secp,$(COSMOS_BUILD_OPTIONS)))
  build_tags += libsecp256k1_sdk
endif

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=sonr \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=sonrd \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
			-X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TMVERSION)

ifeq ($(ENABLE_ROCKSDB),true)
  BUILD_TAGS += rocksdb_build
  test_tags += rocksdb_build
endif

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  ifneq ($(ENABLE_ROCKSDB),true)
    $(error Cannot use RocksDB backend unless ENABLE_ROCKSDB=true)
  endif
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
endif

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# Check for debug option
ifeq (debug,$(findstring debug,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -gcflags "all=-N -l"
endif

all: Makefile
	@echo ''
	@sed -n 's/^##//p ' $<

## Makefile
## > The following Makefile is used for various actions for the Sonr project.
##
## bind        :   Binds Android, iOS and Web for Plugin Path
bind: bind.ios bind.mac bind.android bind.web

## └─ android       - Android AAR
bind.android:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -a

## └─ ios           - iOS Framework
bind.ios:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -i

## └─ mac           - Mac Framework
bind.mac:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -m

## └─ web           - WASM Framework
bind.web:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -w

## └─ tar           - Build All & Tar Compress
bind.tar:
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -a
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -i
	TAR_COMPRESS=true && sh $(SCRIPTS_DIR)/bind.sh -w

build-all:
	GOOS=linux GOARCH=amd64 go build -o ./build/sonr-linux-amd64 ./cmd/sonrd/main.go
	GOOS=linux GOARCH=arm64 go build -o ./build/sonr-linux-arm64 ./cmd/sonrd/main.go
	GOOS=darwin GOARCH=amd64 go build -o ./build/sonr-darwin-amd64 ./cmd/sonrd/main.go

do-checksum:
	cd build && sha256sum sonr-linux-amd64 sonr-linux-arm64 sonr-darwin-amd64 > myproject_checksum

build-with-checksum: build-all do-checksum


## proto       :   Compiles Go Proto Files and pushes to Buf.Build
proto: proto.go proto.buf

## └─ go            - Generate to x/*/types and thirdparty/types/*
proto.go:
	ignite generate proto-go --yes
	go mod tidy
	@echo "✅ Generated Go Proto Files"

## └─ buf           - Build and push to buf.build/sonr-io/blockchain
proto.buf:
	cd $(ROOT_DIR)/proto && buf mod update && buf build
	@echo "✅ Pushed Protos to Buf.Build"

## └─ publish       - Compiles protos, buf.build publish, Zips protos in build
proto.publish:
	cd $(ROOT_DIR)/proto && buf mod update && buf build
	@echo "✅ Pushed Protos to Buf.Build"
	cp -r proto/ build/proto/
