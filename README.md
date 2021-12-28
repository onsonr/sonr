# Sonr Core

## About The Project

Sonr is building the most simple and intuitive Decentralized Web experience for users and developers alike with our revolutionary blockchain and universal digital wallet.

### Built With

- [Golang](https://go.dev)
- [Libp2p](https://libp2p.io)

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Installation

1. Clone the repo

   ```sh
   git clone https://github.com/sonr-io/sonr.git
   ```

2. Install NPM packages

  ```bash
  # Binds Android and iOS for Plugin Path
  make bind

  # Binds iOS Framework ONLY
  make bind.ios

  # Binds AAR for Android ONLY
  make bind.android

  # Compiles Protobuf models for Core Library and Plugin
  make proto

  # Binds Binary, Creates Protobufs, and Updates App
  make upgrade

  # Reinitializes Gomobile and Removes Framworks from Plugin
  make clean
  ```

Docker Instructions

1. Build the Docker image

   ```sh
   docker build -t ghcr.io/sonr-io/snrd .
   ```

2. Run the Docker image

  ```sh
  docker run -it -p 443:26225 ghcr.io/sonr-io/snrd
  ```

### Structure

This project is a pseudo-monorepo, meaning it has a single root directory and all of its packages are in subdirectories. The structure is as follows:

``` text
/cmd            ->        Packaged Binaries.
  └─ bin        ->        +   Daemon RPC for Desktop Builds.
  └─ highway    ->        +   Sonr Custodian Node (desktop, server)
  └─ lib        ->        +   Binded Mobile Framework (aar, framework)
/docs           ->        Documentation.
/extensions     ->        Sonr Extension's for platform integrations (Figma, Chrome, Native, etc.)
/internal       ->        Internal Code. (Networking, Emitter, FileSystem, etc.)
  └─ api        ->        +   Shared Protobuf Models, Generic Types, and Enums.
  └─ common     ->        +   Core data types and functions.
  └─ device     ->        +   Current Node Device management
  └─ host       ->        +   Libp2p Configuration
  └─ keychain   ->        +   Keychain for Private/Public Keys
  └─ node       ->        +   Central Node for Sonr Network
/pkg            ->        Protocol Services for Sonr Core
  └─ domain     ->        +   Shared Protobuf Models, Generic Types, and Enums.
  └─ exchange   ->        +   Data Transfer related Models.
  └─ lobby      ->        +   Node Peer related Models.
  └─ mailbox    ->        +   Creates and Registers Libp2p RPC Service Handlers.
  └─ transfer   ->        +   Creates an Interface which manages libp2p pubsub topics.
/proto          ->        Protobuf Definition Files.
/tools          ->        API Services utilized in the project.
  └─ config     ->        +   File System structure management
  └─ internet   ->        +   Namebase, REST, and DNS Resolver
  └─ state      ->        +   State Machine Management
```

<!-- ROADMAP -->

## Roadmap

See the [open issues](https://github.com/sonr-io/core/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

Prad Nukala - [TheSonrApp](https://twitter.com/TheSonrApp) - pradn@sonr.io

Project Link: [Github](https://github.com/sonr-io/core) - [Discord](https://sonr.io) - [Website](https://sonr.io)

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

- [Libp2p](https://libp2p.io/)
- [Textile](https://www.textile.io/)
- [Handshake](https://handshake.org/)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/sonr-io/core.svg?style=for-the-badge
[contributors-url]: https://github.com/sonr-io/core/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/sonr-io/core.svg?style=for-the-badge
[forks-url]: https://github.com/sonr-io/core/network/members
