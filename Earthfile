VERSION 0.7
PROJECT sonrhq/sonr-testnet-0

IMPORT github.com/sonrhq/chain:master as chain
IMPORT github.com/sonrhq/studio:master as studio

build:
    BUILD chain+build
    BUILD studio+build

dev-down:
    LOCALLY
    RUN task down
