[![CodeFactor](https://www.codefactor.io/repository/github/sonr-io/core/badge/release?s=ee02a1b599502678b3d583aa5b6d1f55d2137ded)](https://www.codefactor.io/repository/github/sonr-io/core/overview/release)
[![CI Workflow 🐿](https://github.com/sonr-io/core/actions/workflows/ci.yml/badge.svg)](https://github.com/sonr-io/core/actions/workflows/ci.yml)

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/sonr-io/core">
    <img src="https://uploads-ssl.webflow.com/60e4b57e5960f8d0456720e7/60fbc0e3fcdf204c7ed9946b_Github%20-%20Core.png" alt="Logo" height="275">
  </a>

  <p align="center">
  Core Framework that manages the Sonr Libp2p node in Go, Handles File Management, Connection to Peer, and Pub-Sub for Lobby.
    <a href="https://github.com/sonr-io/core"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/sonr-io/core">View Demo</a>
    ·
    <a href="https://github.com/sonr-io/core/issues">Report Bug</a>
    ·
    <a href="https://github.com/sonr-io/core/issues">Request Feature</a>
  </p>
</p>
<br />

_By [Sonr](https://www.sonr.io), creators of [Sonr Protocol](https://www.twitter.com/SonrProtocol)_

---

<!-- ABOUT THE PROJECT -->

## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

### Built With

- [Golang]()
- [Node.js]()
- [Flutter]()

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
- **golang**

  ```sh
  go get github.com/golang/dep/cmd/dep
  ```

### Installation

1. Clone the repo

   ```sh
   git clone https://github.com/sonr-io/sonr.git
   ```

2. Use Makefile Commands

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

3. Pushing the image to ghcr.io *optional*

  ```sh
  docker push ghcr.io/sonr-io/snrd:latest
  ```

### Structure

This project is a pseudo-monorepo, meaning it has a single root directory and all of its packages are in subdirectories. The structure is as follows:

```text
/app            ->        Exposed common application code from modules
/cmd            ->        Packaged Binaries.
  └─ bind       ->        +   Desktop Binary
  └─ daemon     ->        +   Binded Mobile Framework (aar, xcframework)
  └─ highway    ->        +   Sonr Custodian Node (desktop, server)
/common         ->        Shared Types
/docs           ->        Documentation and Specifications
/device         ->        Current Node Device management
/internal       ->        Internal Code. (Networking, Emitter, FileSystem, etc.)
  └─ beam       ->        +   Shared Protobuf Models, Generic Types, and Enums.
  └─ did        ->        +   Core data types and functions.
  └─ host       ->        +   Libp2p Configuration
/node           ->        Central Node for Sonr Network
/pkg            ->        Protocol Services for Sonr Core
  └─ discover   ->        +   Shared Protobuf Models, Generic Types, and Enums.
  └─ exchange   ->        +   Data Transfer related Models.
  └─ identity   ->        +   Node Peer related Models.
  └─ registry   ->        +   Creates and Registers Libp2p RPC Service Handlers.
  └─ transmit   ->        +   Creates an Interface which manages libp2p pubsub topics.
/proto          ->        Protobuf Definition Files.
/wallet         ->        Universal Wallet Interface for Sonr Core.
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
- [Handshake](https://handshake.org/)
- [Flutter](https://flutter.dev/)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/sonr-io/core.svg?style=for-the-badge
[contributors-url]: https://github.com/sonr-io/core/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/sonr-io/core.svg?style=for-the-badge
[forks-url]: https://github.com/sonr-io/core/network/members
[stars-shield]: https://img.shields.io/github/stars/sonr-io/core.svg?style=for-the-badge
[stars-url]: https://github.com/sonr-io/core/stargazers
[issues-shield]: https://img.shields.io/github/issues/sonr-io/core.svg?style=for-the-badge
[issues-url]: https://github.com/sonr-io/core/issues
[license-shield]: https://img.shields.io/github/license/sonr-io/core.svg?style=for-the-badge
[license-url]: https://github.com/sonr-io/core/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/sonr-io
