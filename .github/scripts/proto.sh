#!/bin/bash
SCRIPTDIR=$(dirname "$0")
cd ${SCRIPTDIR}/../../
PROJECT_DIR=$(pwd);
cd ${PROJECT_DIR}/../
ROOT_DIR=$(pwd);
PROTO_DIR_DART=${ROOT_DIR}/plugin/lib/src
PROTO_LIST_ALL=${PROJECT_DIR}/proto/**/*.proto
PROTO_LIST_CLIENT=${PROJECT_DIR}/proto/node/*.proto
PROTO_LIST_COMMON=${PROJECT_DIR}/proto/common/*.proto
MODULE_NAME=github.com/sonr-io/core
GO_OPT_FLAG=--go_opt=module=${MODULE_NAME}
GRPC_OPT_FLAG=--go-grpc_opt=module=${MODULE_NAME}
PROTO_GEN_GO="--go_out=."
PROTO_GEN_RPC="--go-grpc_out=."
PROTO_GEN_DOCS="--doc_out=docs"
PROTO_GEN_DART="--dart_out=grpc:$(PROTO_DIR_DART)"
outputs=""

# Functions for printing help
usage() {                                 # Function: Print a help message.
  echo "Usage: sonr-io/core/proto.sh [ -o OUTPUT (all, dart, go, ts) ]" 1>&2
  echo ""
}

compile_docs() {
    echo "Compiling docs"
    protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DOCS)
}

compile_go() {
  echo "Generating Protobuf Go code..."
  protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_GO) $(GO_OPT_FLAG)
  protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_RPC) $(GRPC_OPT_FLAG)
}

compile_dart() {
    echo "Generating Protobuf Dart code..."
    protoc $(PROTO_LIST_CLIENT) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)
    protoc $(PROTO_LIST_COMMON) --proto_path=$(ROOT_DIR) $(PROTO_GEN_DART)
}

# ------------------------------------------------------------------------
# ----- <Input Flags> ----------------------------------------------------
while getopts ":o:" options; do         # Loop: Get the next option;
  case "${options}" in                    #
    o)                                    # If the option is t,
      outputs=${OPTARG}                     # Set $TIMES to specified value.
      ;;
    :)                                    # If expected argument omitted:
      echo "Error: -${OPTARG} requires an argument."
      exit_abnormal                       # Exit abnormally.
      ;;
    *)                                    # If unknown (any other) option:
      exit_abnormal                       # Exit abnormally.
      ;;
  esac
done
# ----- <Input Flags/> ---------------------------------------------------
# -------------------------------------------------

# ------------------------------------------------------------------------
# ----- <Input Reader> ---------------------------------------------------
echo ""
echo "🌈  Sonr Proto Compile"
echo ""
if [ "${outputs}" == "" ]; then
    echo "1. Choose Compile Outputs (all, dart, go, ts)"
    read -p "Output Path: " output
else
    echo " └─ Compilng for [ ${outputs} ] 📩 "
fi
echo ""
# ----- <Input Reader/> --------------------------------------------------
# ------------------------------------------------------------------------
