# Sonr Highway Node

`highway-go` is a the golang implementation for the custodial node on the **Sonr Network**. It is responsible for managing blockchain transactions, and complex interactions between Motor nodes and services.

## About The Project

Sonr is building the most immersive DWeb experience for both Users and Developers alike. We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that’s composable and interoperable with all existing protocols.

For this we built our Networking layer in [Libp2p](“https://libp2p.io”) and our Layer 1 Blockchain with [Starport](“https://starport.com”). Our network comprises of two separate nodes: [Highway](“https://github.com/sonr-io/highway”) and [Motor](“https://github.com/sonr-io/motor”), which each have a specific use case on the network. In order to maximize the onboarding experience, we developed our own [Wallet](“https://github.com/sonr-io/wallet) which has value out of the gate!

### Documentation

All documentation inluding API Reference, Guides, and Recipes are available on the [Sonr Docs](“https://docs.sonr.io”) website.

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Requirements

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/docker-for-mac/install/mac/)
- [Taskfile](https://taskfile.dev)
  - `brew install go-task/tap/go-task`

### Installation

1. Download the `sonr-io/sonr` blockchain node.

  ```shell
  // For Non M1 Systems
  curl https://sonr.network/sonr! | sudo bash

  // For M1 Systems
  curl https://sonr.network/sonr | bash # Install
  sudo mv sonr /usr/local/bin/ # Move to Directory
  ```

2. Run the Sonr Blockchain Node

  ```sh
  sonrd start
  ```

3. Run the `sonr-io/highway-go` server with `task run`.


### Structure

This project is a pseudo-monorepo, meaning it has a single root directory and all of its packages are in subdirectories. The structure is as follows:

```text
/cmd             ->        CLI commands for the project
/grpc            ->        Highway Service gRPC implementation
/pkg             ->        Protocol Services for Sonr Core
  └─ acccount    ->        +   Service and Account Management
  └─ client      ->        +   Blockchain Client
/proto           ->        Highway API Schema and Protobuf Definitions
/remix           ->        Remix frontend
```

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project facilitated under **Sonr Inc.** is distributed under the **GPLv3 License**. See `LICENSE.md` for more information.

## Acknowledgements

- [Libp2p](https://libp2p.io/)
- [Textile](https://www.textile.io/)
- [Handshake](https://handshake.org/)
