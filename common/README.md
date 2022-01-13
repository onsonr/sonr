
# Sonr DID

**Sonr DID** is a module for building, verifying, specifying, and parsing DIDs for the [Sonr](https://sonr.io) blockchain.

## Description

Sonr is building the most immersive DWeb experience for both Users and Developers alike. We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that’s composable and interoperable with all existing protocols.

For this we built our Networking layer in [Libp2p](“https://libp2p.io”) and our Layer 1 Blockchain with [Starport](“https://starport.com”). Our network comprises of two separate nodes: [Highway](“https://github.com/sonr-io/highway”) and [Motor](“https://github.com/sonr-io/motor”), which each have a specific use case on the network. In order to maximize the onboarding experience, we developed our own [Wallet](“https://github.com/sonr-io/wallet) which has value out of the gate!

## Getting Started

### Dependencies

- [Golang](https://go.dev)

### Installing

To install the latest version of the Sonr blockchain node's binary, execute the following command on your machine:

``` shell
go get -u https://github.com/sonr-io/did
```

### Configuration

This project is a pseudo-monorepo, meaning it has a single root directory and all of its packages are in subdirectories. The structure is as follows:

``` text
/beam            ->        Real-time Key/Value Store
/common          ->        Core data types and functions.
/device          ->        Node Device management
/docs            ->        Documentation.
/exchange        ->        Data Transfer related Models.
/host            ->        Libp2p Host Configuration
/identity        ->        Identity management models and interfaces
/node            ->        Highway and Motor node builder configuration
/proto           ->        Protobuf Definition Files.
/transmit        ->        Protocol for byte transmission between nodes
/types           ->        Protobuf Compiled Types
  └─ cpp         ->        +   C++ Definition Files
  └─ go          ->        +   Golang Definition Files
  └─ java        ->        +   Java Definition Files
/wallet          ->        Interfaces for managing Universal Wallet
```

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Authors

Contributors names and contact info

- [Prad Nukala](“https://github.com/prnk28”)

## License

This project facilitated under **Sonr Inc.** is distributed under the **GPLv3 License**. See `LICENSE.md` for more information.

## Acknowledgments

Inspiration, code snippets, etc.

- [W3 Specification](https://w3c-ccg.github.io/did-spec)
- [Handshake](https://handshake.org/)
