<h1 align="center">Sonr Blockchain</h1>

<div align="center">
  :trident: :dolphin: :godmode: :trident:
</div>
<div align="center">
  <strong>The Official Sonr Blockchain source code</strong>
</div>
<div align="center">
  A <code>easy-to-use</code> framework for building immersive decentralized applications.
</div>

<br />

<div align="center">
  <!-- Stability -->
    <img alt="CodeFactor Grade" src="https://img.shields.io/codefactor/grade/github/sonr-io/sonr/master?style=for-the-badge">
  <!-- NPM version -->
  <a href="https://godoc.org/github.com/sonr-io/blockchain">
  <img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge" />
  </a>
  <!-- Test Coverage -->
  <a href="https://codecov.io/github/choojs/choo">
<img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/sonr-io/blockchain?label=TLOC&style=for-the-badge">
  </a>
  <!-- Downloads -->
<img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/sonr_io?color=%2300ACEE&label=🐦 sonr_io&style=for-the-badge">
</div>

<div align="center">
  <h3>
    <a href="https://sonr.io">
      Home
    </a>
    <span> | </span>
    <a href="https://discord.gg/tjWMfvQZ7b">
      Discord
    </a>
    <span> | </span>
    <a href="https://github.com/sonr-io/blockchain/issues">
      Issues
    </a>
    <span> | </span>
      <!-- <span> | </span> -->
    <a href="https://docs.sonr.io">
      Docs
    </a>
     <span> | </span>
      <!-- <span> | </span> -->
    <a href="./CHANGELOG.md">
      Changelog
    </a>
  </h3>
</div>

<div align="center">
  <sub>The quickest way to production in Web3. Built with ❤︎ by the
  <a href="mailto:team@sonr.io">Sonr Team</a> and
  <a href="https://github.com/sonr-io/blockchain/graphs/contributors">
    contributors
  </a>
</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Getting Started](#getting-started)
  - [Overview](#overview)
  - [Configuration](#configuration)
  - [Resources](#resources)
    - [Architecture Decision Record's (ADR)](#architecture-decision-records-adr)
    - [Module Documentation](#module-documentation)
    - [Additional Specs](#additional-specs)
- [Install](#install)
  - [Requirements](#requirements)
    - [Development](#development)
  - [Release](#release)
- [Usage](#usage)
  - [Start the Blockchain](#start-the-blockchain)
  - [Run the Flutter Frontend](#run-the-flutter-frontend)
  - [Run the Vue.js Frontend](#run-the-vuejs-frontend)
  - [Starport CLI Reference](#starport-cli-reference)
- [Contributions](#contributions)
  - [Authors](#authors)
  - [Submitting a PR](#submitting-a-pr)
- [Acknowledgments](#acknowledgments)
- [License](#license)

## Getting Started

Sonr is building the most immersive DWeb experience for both Users and Developers alike. We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that’s composable and interoperable with all existing protocols.

For this we built our Networking layer in [Libp2p](“https://libp2p.io”) and our Layer 1 Blockchain with [Starport](“https://starport.com”). Our network comprises of two separate nodes: [Highway](“https://github.com/sonr-io/highway”) and [Motor](“https://github.com/sonr-io/motor”), which each have a specific use case on the network. In order to maximize the onboarding experience, we developed our own [Wallet](“https://github.com/sonr-io/wallet) which has value out of the gate!

### Overview

Peers in the network can dial other peers in the network to exchange messages using various transports, like QUIC, TCP, WebSocket, and Bluetooth. Modular design of the libp2p framework enables it to build drivers for other transports. Peers can run on any device, as a cloud service, mobile application or in the browser and talk to each other as long as they are connected through the same libp2p network.

### Configuration

This project is a pseudo-monorepo, meaning it has a single root directory and all of its packages are in subdirectories. The structure is as follows:

```text
/app             ->        Exported Starport app
/cmd             ->        Packaged libraries
  └─ sonrd       ->        +   Blockchain Binary
/docs            ->        Documentation.
/proto           ->        Cosmos SDK Protocol Definitions
/testutil        ->        Blockchain test utilities.
/vue             ->        Vue.js frontend for Cosmos SDK
/x               ->        Implementation of Cosmos-Sonr Schemas
  └─ bucket      ->        +   Collections of blobs and objects
  └─ channel     ->        +   Realtime Data Transmissions
  └─ object      ->        +   Verifiable Custom Objects
  └─ registry    ->        +   Name and Service Registration
```

### Resources

Docs and guides to help you understand the Sonr ecosystem.

#### Architecture Decision Record's (ADR)

- [ADR-001: Sonr DID Method Specification](./docs/adrs/ADR-001.md)

#### Module Documentation

- [Buckets](./x/bucket/README.md)
- [Channels](./x/channel/README.md)
- [Objects](./x/object/README.md)
- [Registry](./x/registry/README.md)

#### Additional Specs

- [Official Docs](https://docs.sonr.io)
- [DID Spec](https://docs.google.com/presentation/d/1AwSkO6s0UQ1YHCIOT5Ue5-gcIuvie5q4f3cAiAj6lr8/edit?usp=sharing)

## Install

To get a local copy up and running follow these simple steps.

### Requirements

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/docker-for-mac/install/mac/)
- [Taskfile](https://taskfile.dev)
  - `brew install go-task/tap/go-task`
- [Starport](https://docs.starport.network/guide/install.html#upgrading-your-starport-installation)

#### Development

1. Download the `starport` CLI tool.

```shell
// For Non M1 Systems
curl https://get.starport.network/starport! | bash

// For M1 Systems
curl https://get.starport.network/starport | bash # Install
sudo mv starport /usr/local/bin/ # Move to Directory
```

2. Serve the Blockchain

```sh
starport chain serve # Serve without resetting the chain
starport chain serve --reset-once # Reset the chain
```
#### Development without starport 
1. Initialize the chain: `./sonrd init my-node --chain-id sonr`
1. Add a key to your keyring (using test): `./sonrd keys add --keyring-backend test alice --home ~/.sonr`
1. Add the account as a genesis account: `./sonrd add-genesis-account $(./sonrd keys show alice -a) 1000000000000000stake,1000000000000snr`
1. Create a genesis transaction: `./sonrd gentx alice 1000000000000000stake --chain-id sonr`
1. Collect the genesis transactions: `./sonrd collect-gentxs`
1. Start the chain

#### <b>Or just run this</b>
```bash
./sonrd init my-node --chain-id sonr
./sonrd keys add --keyring-backend test alice --home ~/.sonr
./sonrd add-genesis-account $(./sonrd keys show alice -a) 1000000000000000stake,1000000000000snr
./sonrd gentx alice 1000000000000000stake --chain-id sonr
./sonrd collect-gentxs
./sonrd start
```
### Release

To install the latest version of the Sonr blockchain node's binary, execute the following command on your machine:

```shell
// For Non M1 Systems
curl https://sonr.network/sonr! | sudo bash

// For M1 Systems
curl https://sonr.network/sonr | bash # Install
sudo mv sonr /usr/local/bin/ # Move to Directory
```

## Usage

To launch the Sonr Blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Start the Blockchain

```shell
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Run the Flutter Frontend

Starport has scaffolded a Flutter-based mobile app in the `flutter` directory. Run the following commands to install dependencies and start the app:

```shell
cd flutter
flutter pub get
flutter run
```

### Run the Vue.js Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```text
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

### Starport CLI Reference

`starport chain serve`

- This is the command that starts the blockchain.
- By adding the flag `--reset-once` it will reset the blockchain on the first startup

`starport scaffold vue`

- This is the command that rescaffolds a vue frontend (web) interface for the Cosmos SDK.

`starport scaffold flutter`

- This is the command that rescaffolds a flutter frontend (mobile, desktop) interface for the Cosmos SDK.

`starport generate openapi`

- Generates an OpenAPI spec for your chain from your config.yml

`starport generate dart`

- Generate a Dart client

`starport generate vuex`

- Generate Vuex store for you chain's frontend from your config.yml

## Contributions

> Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated!**

### Authors

- [Prad Nukala](https://github.com/prnk28)

### Submitting a PR

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Acknowledgments

Tools, libraries, and frameworks that make the Sonr project possible:

- [Libp2p](https://libp2p.io/)
- [Cosmos](https://www.cosmos.network/)
- [Handshake](https://handshake.org/)

## License

This project facilitated under **Sonr Inc.** is distributed under the **GPLv3 License**. See `LICENSE.md` for more information.
