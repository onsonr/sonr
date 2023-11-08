# Sonr Testnet Deployment

This directory contains the necessary files to stand up a Sonr validator pool
consisting of 4 validator nodes, 4 sentry nodes, and a faucet using Docker Compose.

## Prerequisites

- Docker and Docker Compose installed on your machine
- A valid Doppler token to access the necessary secrets

## Usage

1. Clone this repository to your local machine:

```shell
   git clone https://github.com/sonrlabs/sonr-validator-pool.git
```

1. Navigate to the sonr-validator-pool directory:

```shell
   cd sonr-validator-pool
```

1. Set the Doppler token as an environment variable:

```shell
 export DOPPLER_TOKEN=<your-doppler-token>
```

1. Start the validator pool:

```shell
   docker-compose up -d
```

> This will start the validator nodes, sentry nodes, and faucet in detached mode.

1. Check the logs to make sure everything is running correctly:

```shell
   docker-compose logs -f
```

> This will show the logs of all the containers. Press Ctrl+C to exit.

1. Use the faucet to fund your validator nodes:

- The faucet is available at <http://localhost:8000>.
- To fund a validator node, copy its address from the logs and
  paste it into the faucet's input field.
- Click the "Fund" button to send tokens to the validator node.

1. Monitor the validator nodes using the Sonr Explorers

- The Sonr Explorer is available at <http://localhost:8080>.
- Use the validator node addresses to search for them in the explorer.
- You can also use the explorer to monitor the network status and transactions.

### Configuration

The validator pool can be configured using the following environment variables:

- DOPPLER_TOKEN: The Doppler token to access the necessary secrets.
- MONIKER: The moniker of the validator nodes (default: sonr-validator).
- CHAIN_ID: The chain ID of the Sonr network (default: sonr-testnet).
- KEYRING: The keyring backend to use for the validator nodes (default: os).
- KEY: The name of the key to use for the validator nodes (default: validator).
- STAKE: The amount of tokens to stake for each validator node (default: 1000000000000000000).
- GRPC_ADDRESS: The gRPC address of the validator nodes (default: 0.0.0.0:9090).
- API_ENABLE: Whether to enable the API for the validator nodes (default: true).
- API_SWAGGER: Whether to enable the Swagger UI for the validator nodes
  (default: true).
- API_ADDRESS: The API address of the validator nodes (default: 0.0.0.0:1317).
- MINIMUM_GAS_PRICES: The minimum gas prices for the validator nodes (default: 0.025usnr).
- SEEDS: The seeds for the sentry nodes (default: "").
- PERSISTENT_PEERS: The persistent peers for the validator nodes (default: "").
- PRIVATE_PEER_IDS: The private peer IDs for the validator nodes (default: "").

You can set these environment variables in the .env file in the
sonr-validator-pool directory, or use **Doppler+Taskfile** to set them.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
