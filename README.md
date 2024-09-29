<div align="center" style="text-align: center;">

# `sonr` - Sonr Chain

[![Go Reference](https://pkg.go.dev/badge/github.com/onsonr/sonr.svg)](https://pkg.go.dev/github.com/onsonr/sonr)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/onsonr/sonr)
![GitHub Release Date - Published_At](https://img.shields.io/github/release-date/onsonr/sonr)
[![Static Badge](https://img.shields.io/badge/homepage-sonr.io-blue?style=flat-square)](https://sonr.io)

[![Go Report Card](https://goreportcard.com/badge/github.com/onsonr/sonr)](https://goreportcard.com/report/github.com/onsonr/sonr)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=sonrhq_sonr&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)

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

Sonr would not have been possible without the direct and indirect support of the following individuals:

- **Juan Benet**: For the IPFS Ecosystem.
- **Satoshi Nakamoto**: For Bitcoin.
- **Steve Jobs**: For User first UX.
- **Tim Berners-Lee**: For the Internet.

<br />

## Community & Support

- [Forum](https://github.com/onsonr/sonr/discussions)
- [Issues](https://github.com/onsonr/sonr/issues)
- [Twitter](https://sonr.io/twitter)
- [Dev Chat](https://sonr.io/discord)
