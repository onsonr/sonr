# Sonr Network Rails

Dependencies for the Sonr blockchain network. These containers are designed to be preconfigured in order to operate out-of-the-box with a [sonr validator node](https://github.com/sonrhq/sonr).

## Installation

This project uses Earthfile as a basis to build, test, install, and deploy the project. Earthfile is a simple, yet powerful, way to define your project's build process and environment.

### Getting Started

Installation for __Linux__:

```bash
sudo /bin/sh -c 'wget https://github.com/earthly/earthly/releases/latest/download/earthly-linux-amd64 -O /usr/local/bin/earthly && chmod +x /usr/local/bin/earthly && /usr/local/bin/earthly bootstrap --with-autocomplete'
```

Installation for __MacOS__:

```bash
brew install earthly && earthly bootstrap
```

## Usage

### 1. Build

Building an individual package:

```bash
cd <PACKAGE> && earthly +build
```

Build all the packages:

```bash
earthly +build-all
```

## Dependencies

Sonr has chosen to follow Decentralized principles and leverage a tech stack which promotes Self-Sovereign Identity. We have chosen to use the following technologies.

### `Dendrite`

Provides a Matrix homeserver for validators to run in tandem with Sonr. Enabling the SonrAppService provides decentralized communication amongst Sonr identities. Preconfigured to work with the [sonr validator node](https://github.com/sonrhq/sonr).

### `Cosmos-Faucet`

Provides a faucet for the Cosmos SDK based blockchain. Preconfigured to work with the [sonr validator node](https://github.com/sonrhq/sonr).

### `IPFS`

Configured to allow for IPNS Entries and IPFS Pinning.

### `Nginx`

We use the nginx auto-proxy container in order to streamline docker stack based deployment.

#### Usage in Docker Compose

Add the proxy as a container and map the port.

```yaml
nginx-proxy:
    image: sonr-nginx:latest # WARNING: This image is only available after an Earthly Build
    ports:
      - "80:80"
    volumes:
      - /var/run/docker.sock:/tmp/docker.sock:ro
```

Add the following labels to any container you want to proxy.

```yaml
  whoami:
    image: jwilder/whoami
    expose:
      - "8000" # This is the port that the container exposes
    environment:
      - VIRTUAL_HOST=whoami.example # This is the domain name that the container will proxy to
      - VIRTUAL_PORT=8000 # This is the port that the container exposes
```

### `Postgres`

We leverage Postgres for Transaction Indexing for `sonrd` and matrix message storage for `dendrite`. In production we opt for the DigitalOcean Hosted Database.
