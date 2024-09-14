<div align="center" style="text-align: center;">
<img src="https://pub-97e96d678cb448969765e4c1542e675a.r2.dev/github-sonr.png" width="256" height="256" />

# `sonr` - Sonr Chain

[![Go Reference](https://pkg.go.dev/badge/github.com/onsonr/sonr.svg)](https://pkg.go.dev/github.com/onsonr/sonr)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/onsonr/sonr)
![GitHub Release Date - Published_At](https://img.shields.io/github/release-date/onsonr/sonr)
[![Static Badge](https://img.shields.io/badge/homepage-sonr.io-blue?style=flat-square)](https://sonr.io)
![Discord](https://img.shields.io/discord/843061375160156170?logo=discord&label=Discord%20Chat)

[![Go Report Card](https://goreportcard.com/badge/github.com/onsonr/sonr)](https://goreportcard.com/report/github.com/onsonr/sonr)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=sonrhq_sonr&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=sonrhq_sonr&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
[![Mutable.ai Auto Wiki](https://img.shields.io/badge/Auto_Wiki-Mutable.ai-blue)](https://wiki.mutable.ai/di-dao/sonr)

</div>
<br />

Sonr is a combination of decentralized primitives. Fundamentally, it is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS—providing users with a secure, portable decentralized identity.

<br />

## Components

### `sonrd`

The main blockchain node that runs the `sonr` chain. It is responsible for maintaining the state of the chain, including IPFS based vaults, and did documents.

### `vault`

The `vault` is a wasm module that is compiled and deployed to IPFS on behalf of the user. It is responsible for storing and retrieving encrypted data.

- SQLite Database backend
- Encryption via admonition
- Authentication via webauthn
- Authorization via Macroons
- HTTP API

## Acknowledgements

Sonr would not have been possible without the direct and indirect support of the following organizations and individuals:

- **Protocol Labs**: For IPFS & Libp2p.
- **Interchain Foundation**: For Cosmos & IBC.
- **Tim Berners-Lee**: For the Internet.
- **Satoshi Nakamoto**: For Bitcoin.
- **Steve Jobs**: For Taste.

<br />

## Community & Support

- [Forum](https://github.com/onsonr/sonr/discussions)
- [Issues](https://github.com/onsonr/sonr/issues)
- [Twitter](https://sonr.io/twitter)
- [Dev Chat](https://sonr.io/discord)
