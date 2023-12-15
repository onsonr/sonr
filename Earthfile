VERSION 0.7
PROJECT sonrhq/testnet-1

LOCALLY
GIT CLONE git@github.com:sonrhq/chain.git chain
GIT CLONE git@github.com:sonrhq/identity.git identity
GIT CLONE git@github.com:sonrhq/rails.git rails
GIT CLONE git@github.com:sonrhq/service.git service


IMPORT ./chain AS chain
IMPORT ./identity AS identity
IMPORT ./rails AS rails
IMPORT ./service AS service

# build - Initializes the base repository and clones dependencies
build:
	FROM sonrhq/builder:latest
	BUILD identity+test
	BUILD service+test
	BUILD chain+build
	BUILD rails+build
