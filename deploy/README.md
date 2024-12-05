# Deployment

This directory contains the configuration files for deploying the Sonr blockchain.

## Contents

- `devnet`: Configuration for deploying the Sonr blockchain on the devnet (local development).
- `testnet`: Configuration for deploying the Sonr blockchain on the testnet (current prod setup)

## Usage

Configuration is automatically loaded from the `PKL` files in the root of the repository. These templates are generated during deployment initialization.

To deploy the total network, run the following command:

```bash
devbox run <network>
```

Replace `<network>` with either `devnet` or `testnet` from the root of the repository.

## Components

### Sonr

The Sonr blockchain is deployed using the `sonrd` binary. This binary is built using the `Makefile` in the root of the repository.

### IPFS

IPFS is deployed using the `ipfs` binary. This binary is built using the `Makefile` in the root of the repository.

### Hway

Hway is deployed using the `hway` binary. This binary is built using the `Makefile` in the root of the repository.

### Synapse

Synapse is deployed using the `matrix-synapse` binary. This binary is built using the `Makefile` in the root of the repository.

### Tigerbeetle

Tigerbeetle is deployed using the `tigerbeetle` binary. This binary is built using the `Makefile` in the root of the repository.
