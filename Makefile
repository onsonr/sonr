SHELL=/bin/bash

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
SONR_ROOT_DIR=/Users/prad/Developer
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CORE_DIR=$(SONR_ROOT_DIR)/core
DESKTOP_DIR=$(SONR_ROOT_DIR)/desktop

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
PROTO_DEF_PATH=/Users/prad/Developer/core/proto

# @ Proto Directories
PROTO_DIR_DART=$(SONR_ROOT_DIR)/mobile/lib/src
PROTO_DIR_TS=$(SONR_ROOT_DIR)/desktop/assets/proto
PROTO_LIST_ALL=${ROOT_DIR}/proto/**/*.proto
PROTO_LIST_API=${ROOT_DIR}/proto/api/*.proto
PROTO_LIST_COMMON=${ROOT_DIR}/proto/common/*.proto
PROTO_FILE_NODE_CLIENT=${ROOT_DIR}/proto/node/client.proto
PROTO_FILE_NODE_HIGHWAY=${ROOT_DIR}/proto/node/highway.proto
PROTO_LIST_PROTOCOLS=${ROOT_DIR}/proto/protocols/*.proto
PROTO_LIST_WALLET=${ROOT_DIR}/proto/wallet/*.proto
MODULE_NAME=github.com/sonr-io/core
GO_OPT_FLAG=--go_opt=module=${MODULE_NAME}
GRPC_OPT_FLAG=--go-grpc_opt=module=${MODULE_NAME}
PROTO_GEN_GO="--go_out=."
PROTO_GEN_RPC="--go-grpc_out=."
PROTO_GEN_DOCS="--doc_out=docs"
PROTO_GEN_DART="--dart_out=grpc:$(PROTO_DIR_DART)"
PROTO_GEN_TS="--ts_proto_out=$(PROTO_DIR_TS)"
TS_OPT_FLAG=--ts_proto_opt=outputClientImpl=grpc-node --ts_proto_opt=useOptionals=true --ts_proto_opt=outputServices=grpc-js --ts_proto_opt=env=node --ts_proto_opt=stringEnums=true --ts_proto_opt=esModuleInterop=true
PROTOC_TS_PLUGIN_PATH=/usr/local/bin/protoc-gen-ts

all:
	@echo "----"
	@echo "Sonr: Compiling Protobufs"
	@echo "----"
	@echo "Generating Protobuf Go code..."
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_GO) $(GO_OPT_FLAG)
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_RPC) $(GRPC_OPT_FLAG)

	@echo "Generating Protobuf Dart code..."
	@protoc $(PROTO_LIST_API) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)
	@protoc $(PROTO_LIST_COMMON) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)
	@protoc $(PROTO_FILE_NODE_CLIENT) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)

	@echo "Generating Protobuf Typescript code..."
	@rm -rf $(PROTO_DIR_TS)/data
	@mkdir -p ${PROTO_DIR_TS}
	@cp -R $(ROOT_DIR)/proto/api $(PROTO_DIR_TS)/api
	@cp -R $(ROOT_DIR)/proto/common $(PROTO_DIR_TS)/common
	@cp -R $(ROOT_DIR)/proto/node $(PROTO_DIR_TS)/node
	@cd $(DESKTOP_DIR) && yarn proto
	@rm -rf $(PROTO_DIR_TS)

	@echo "Generating Protobuf Docs..."
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DOCS)
	@echo "----"
	@echo "✅ Finished Compiling ➡ `date`"
