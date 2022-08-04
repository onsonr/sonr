[![Release Workflow](https://github.com/sonr-io/sonr/actions/workflows/release.yml/badge.svg?branch=dev)](https://github.com/sonr-io/sonr/actions/workflows/release.yml)
<h1 align="center">Sonr Core</h1>

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
  <a href="https://godoc.org/github.com/sonr-io/sonr">
  <img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge" />
  </a>
  <!-- Test Coverage -->
  <a href="https://codecov.io/github/choojs/choo">
<img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/sonr-io/sonr?label=TLOC&style=for-the-badge">
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
    <a href="https://discord.gg/6Z3RmWs257">
      Discord
    </a>
    <span> | </span>
    <a href="https://github.com/sonr-io/sonr/issues">
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
  <sub>The quickest way to production in Web5. Built with ❤︎ by the
  <a href="mailto:team@sonr.io">Sonr Team</a> and
  <a href="https://github.com/sonr-io/sonr/graphs/contributors">
    contributors
  </a>
</div>

---

Sonr is a platform for developers to build decentralized applications which put user privacy first and foremost. It weds decentralized storage technologies such as [IPFS](https://ipfs.io) and [libp2p](https://libp2p.io) with an intuitive, firebase-like developer experience.

Sonr's platform aims to be the most immersive and powerful DWeb experience for both Users and Developers alike. We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that’s composable and interoperable with all existing protocols.

For this we built our Networking layer in [Libp2p](“https://libp2p.io”) and our Layer 1 Blockchain with [Starport](“https://starport.com”). Our network comprises of two separate nodes: [Highway](“https://github.com/sonr-io/sonr/tree/dev/pkg”) and [Motor](“https://github.com/sonr-io/sonr/tree/dev/motor), which each have a specific use case on the network.

For a more in-depth look at building on the Sonr network, check out our [docs](https://docs.sonr.io).

### Project Structure

The `sonr` repo follows the Go project structure outlined in https://github.com/golang-standards/project-layout.

The core packages (`/pkg`) is structured as follows:

```text
/app                ->          Main blockchain executable
/bind               ->          Binded motor executable for ios, android, and wasm
/cmd                ->          Executable binaries
/docker             ->          Docker container files
/docs               ->          Documentation with Docusaurus
/internal           ->          Identity management models and interfaces
/pkg                ->          Core packages for all executables
  └─ client         ->          +   Blockchain Client utilities
  └─ config         ->          +   Configuration settings for Motor and Highway nodes
  └─ crypto         ->          +   Cryptographic primitives and Wallet implementation
  └─ did            ->          +   DID management utilities
  └─ fs             ->          +   File System utilities for Motor
  └─ host           ->          +   Libp2p host for Motor & Highway nodes
/proto              ->          Protobuf Definition Files
/scripts            ->          Project specific scripts
/testutil           ->          Testing utilities for simulations
/vue                ->          Vue based wallet UI
/x                  ->          Cosmos Blockchain Implementation
  └─ bucket         ->          +   See /docs/articles/reference/ADR-003.md
  └─ channel        ->          +   See /docs/articles/reference/ADR-004.md
  └─ registry       ->          +   See /docs/articles/reference/ADR-001.md
  └─ schema         ->          +   See /docs/articles/reference/ADR-002.md
```

### Contributing
**State of the library** - This library is under *extremely* active development. The API can change without notice. Checkout the issues and PRs to be informed about any development.
Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Core Team

Contributors names and contact info

- [Prad Nukala](https://github.com/prnk28)
- [Nick Tindle](https://github.com/ntindle)
- [Josh Long](https://github.com/joshLong145)
- [Brayden Cloud](https://github.com/mcjcloud)
- [Ian Perez](https://github.com/brokecollegekidwithaclothingobsession)

### Acknowledgments

Inspiration, dependencies, compliance, or just pure appreciation!

- [Libp2p](https://libp2p.io/)
- [Handshake](https://handshake.org/)
- [Ignite](https://ignite.com/)
- [IPFS](https://ipfs.io/)
- [W3C](https://www.w3.org/)
- [Fido Alliance](https://fidoalliance.org/)


### License

This project facilitated under **Sonr Inc.** is distributed under the **GPLv3 License**. See `LICENSE.md` for more information.
