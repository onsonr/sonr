SHELL=/bin/bash

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
SONR_ROOT_DIR=/Users/prad/Developer
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CORE_DIR=$(SONR_ROOT_DIR)/core
DESKTOP_DIR=$(SONR_ROOT_DIR)/desktop

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
PROTO_DEF_PATH=/Users/prad/Developer/core/proto

# @ Proto Directories
PROTO_LIST_ALL=${ROOT_DIR}/proto/**/*.proto
MODULE_NAME=github.com/sonr-io/core
GO_OPT_FLAG=--go_opt=module=${MODULE_NAME}
GRPC_OPT_FLAG=--go-grpc_opt=module=${MODULE_NAME}
PROTO_GEN_GO="--go_out=."
PROTO_GEN_RPC="--go-grpc_out=."
PROTO_GEN_DOCS="--doc_out=docs"

all:
	@echo "----"
	@echo "Sonr: Compiling Protobufs"
	@echo "----"
	@echo "Generating Protobuf Go code..."
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_GO) $(GO_OPT_FLAG)
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_RPC) $(GRPC_OPT_FLAG)

	@echo "Generating Protobuf Docs..."
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DOCS)
	@echo "----"
	@echo "✅ Finished Compiling ➡ `date`"
