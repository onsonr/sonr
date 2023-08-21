<div style="text-align: center;">

[![Sonr Logo Banner](.github/images/core-cover.png)](https://sonr.io)

</div>
<div style="text-align: left;">

[![Go Reference](https://pkg.go.dev/badge/github.com/sonrhq/core.svg)](https://pkg.go.dev/github.com/sonrhq/core)
[![Test sonrd](https://github.com/sonrhq/core/actions/workflows/tests.yml/badge.svg)](https://github.com/sonrhq/core/actions/workflows/tests.yml)
[![Release sonrd](https://github.com/sonrhq/core/actions/workflows/release.yml/badge.svg)](https://github.com/sonrhq/core/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonrhq/core)](https://goreportcard.com/report/github.com/sonrhq/core)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=sonr-io_sonr&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=sonr-io_sonr&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
</div>

---

# `sonr-io/sonr`

Sonr is an ibc-enabled blockchain for decentralized identity.

- [x] Passkey based User Accounts. [__Docs__](https://sonr.io/docs/guides/database)
- [x] DKLS-MPC Powered Wallets _(No seed phrases)_. [__Docs__](https://sonr.io/docs/guides/auth)
- [x] IPFS Database and Storage. [__Docs__](https://sonr.io/docs/guides/storage)
  - [x] Redis. [__Docs__](https://sonr.io/docs/guides/api#rest-api-overview)
  - [x] MySQL. [__Docs__](https://sonr.io/docs/guides/api#graphql-api-overview)
  - [ ] User mailboxes. [Status](https://github.com/sonrhq/core/issues/781)
  - [ ] Realtime subscriptions. [Status](https://github.com/sonrhq/core/issues/782)
  - [ ] Matrix chat. [Status](https://github.com/sonrhq/core/issues/783)
- [x] IBC Integrations.
  - [ ] Nomic. [Status](https://github.com/sonrhq/core/issues/784)
  - [ ] Evmos. [Status](https://github.com/sonrhq/core/issues/785)
  - [ ] Osmosis. [Status](https://github.com/sonrhq/core/issues/786)
- [x] Smart Contracts. [__Docs__](https://sonr.io/docs/guides/storage)
- [ ] Typescript Client SDKs. [Status](https://github.com/sonr-io/front/milestone/2)
- [ ] Dashboard. [Status](https://github.com/sonr-io/front/milestone/1)

## Documentation

For full documentation, visit [sonr.io/docs](https://sonr.io/docs)

To see how to Contribute, visit [Getting Started](./docs/wiki/DEVELOPERS.md)

## Community & Support

- [Forum](https://github.com/sonrhq/core/discussions)
- [Issues](https://github.com/sonrhq/core/issues)
- [Twitter](https://sonr.io/twitter)
- [Dev Chat](https://sonr.io/discord)

## Status

- [X] __Alpha__: Closed testing.
- [X] __Private Devnet__: May have kinks. [See projects](https://sonr.io/dashboard).
- [ ] __Public Testnet__: Stable for non-enterprise use. [Join it](https://sonr.io/dashboard).
- [ ] __Mainnet__: Coming soon. [Watch status](https://sonr.io/docs/guides/getting-started/features#feature-status).

We are currently in transitioning to Public Testnet. Watch [releases](https://github.com/sonrhq/core/releases) of this repo to get notified of major updates.

## How it works

Sonr is a combination of decentralized primitives. Fundamentally, it is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS — providing users with a secure, portable decentralized identity.

Sonr is built on top of the Cosmos SDK, which is a framework for building blockchain applications in Golang. We use these modules:

- `x/auth`
- `x/bank`
- `x/distribution`
- `x/ibc`
- `x/ibc/applications/transfer`
- `x/gov`
- `x/params`
- `x/slashing`
- `x/staking`
- `x/upgrade`
- `x/wasm`

### Architecture

Sonr is a [blockchain node](https://sonr.io/dashboard) which you can run locally, or use to join our testnet. You can sign up and start using Sonr without installing anything using our [dashboard](https://sonr.io/dashboard).

![Architecture](.github/images/architecture.svg)

See [additional details](https://sonr.io/docs) on these components.

### Client libraries

Our approach for client libraries is uniform. Abstract away any blockchain specific details, and provide a simple interface for developers to use. We have a few client libraries that we maintain, and provide [guidelines](https://sonr.io) for community maintained libraries.

- TODO

---

## Acknowledgements

Sonr would not have been possible without the direct and indirect support of the following organizations and individuals:

- __Protocol Labs__: For IPFS & Libp2p.
- __Interchain Foundation__: For Cosmos & IBC.
- __Tim Berners-Lee__: For the Internet.
- __Satoshi Nakamoto__: For Bitcoin.
- __Steve Jobs__: For Taste.

Thank you for your support and inspiration!
