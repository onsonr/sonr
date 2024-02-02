#

[![Sonr Logo Banner](.github/assets/core-cover.png)](https://sonr.io)

<div style="text-align: center;">

[![Go Reference](https://pkg.go.dev/badge/github.com/sonrhq/sonr.svg)](https://pkg.go.dev/github.com/sonrhq/sonr)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/sonrhq/sonr)
![GitHub Release Date - Published_At](https://img.shields.io/github/release-date/sonrhq/sonr)
[![Static Badge](https://img.shields.io/badge/homepage-sonr.io-blue?style=flat-square)](https://sonr.io)
![Discord](https://img.shields.io/discord/843061375160156170?logo=discord&label=Discord%20Chat)

[![Go Report Card](https://goreportcard.com/badge/github.com/sonrhq/sonr)](https://goreportcard.com/report/github.com/sonrhq/sonr)
[![Run Tests](https://github.com/sonrhq/sonr/actions/workflows/run-tests.yaml/badge.svg)](https://github.com/sonrhq/sonr/actions/workflows/run-tests.yaml)
[![Merge Queue](https://github.com/sonrhq/sonr/actions/workflows/merge-queue.yaml/badge.svg)](https://github.com/sonrhq/sonr/actions/workflows/merge-queue.yaml)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=sonrhq_sonr&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=sonrhq_sonr)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=sonrhq_sonr&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=sonrhq_sonr&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)

</div>

---

__*Sonr is an ibc-enabled blockchain for decentralized identity*__.

- [x] Passkey based User Accounts. [__Docs__](https://sonr.io/docs/guides/database)
- [x] DKLS-MPC Powered Wallets *(No seed phrases)*. [__Docs__](https://sonr.io/docs/guides/auth)
- [x] IPFS Database and Storage. [__Docs__](https://sonr.io/docs/guides/storage)
  - [x] Redis. [__Docs__](https://sonr.io/docs/guides/api#rest-api-overview)
  - [x] MySQL. [__Docs__](https://sonr.io/docs/guides/api#graphql-api-overview)
  - [ ] Matrix chat. [Status](https://github.com/sonrhq/sonr/issues/783)
- [x] IBC Integrations.
  - [ ] Nomic. [Status](https://github.com/sonrhq/sonr/issues/784)
  - [ ] Evmos. [Status](https://github.com/sonrhq/sonr/issues/785)
  - [ ] Osmosis. [Status](https://github.com/sonrhq/sonr/issues/786)
- [x] Smart Contracts. [__Docs__](https://sonr.io/docs/guides/storage)
- [ ] Typescript Client SDKs. [Status](https://github.com/sonr-io/front/milestone/2)
- [ ] Dashboard. [Status](https://github.com/sonr-io/front/milestone/1)

## Features

Sonr is a combination of decentralized primitives. Fundamentally, it is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS — providing users with a secure, portable decentralized identity.

<table style="table-layout:fixed; white-space: nowrap;">
  <tr>
    <th>Module</th>
    <th colspan=4>Description</th>
    <th>Status</th>
  </tr>
  <tr>
    <th><code><a href="https://github.com/sonrhq/sonr/x/identity">x/identity</a></code></th>
    <th colspan=4>
    The Sonr Identity module is responsible for managing DID based <br />
    accounts using the MPC Protocol - <a href="https://sonr.io/whitepaper">Docs</a>
    </th>
    <th>
      <a href="https://github.com/sonrhq/sonr/actions/workflows/run-tests.yaml">
        <img src="https://github.com/sonrhq/sonr/actions/workflows/run-tests.yaml/badge.svg" alt="CI Pipeline">
      </a>
    </th>
  </tr>
  <tr>
    <th><code>x/oracle</code></th>
    <th colspan=4>
    The Oracle module is responsible for managing Staking delegations <br />
    rewards, and token transfers - <a href="https://sonr.io/whitepaper">Docs</a>
    </th>
    <th>
      <center>
      🚧
      </center>
    </th>
  </tr>
  <tr>
    <th><code><a href="https://github.com/sonrhq/sonr/x/service">x/service</a></code></th>
    <th colspan=4>
    The Service module is responsible for DAO Application Service <br />
    Configurations, and Passkey authentication - <a href="https://sonr.io/whitepaper">Docs</a>
    </th>
    <th>
      <a href="https://github.com/sonrhq/sonr/x/service/actions/workflows/ci.yml?query=branch%3Amaster++">
        <img src="https://github.com/sonrhq/sonr/x/service/actions/workflows/ci.yml/badge.svg?branch=master" alt="CI Pipeline">
      </a>
    </th>
  </tr>
</table>

Sonr is built on top of the Cosmos SDK, which is a framework for building blockchain applications in Golang. We have built the above modules to provide a decentralized identity and asset management system.

## Documentation

For full documentation, visit [sonr.io/docs](https://sonr.io/docs). To see how to Contribute, visit [Getting Started](./docs/contribution/DEVELOPERS.md)

It's recommended to install the following tools:

- [golang](https://golang.org/doc/install)
- [grpcui](https://github.com/fullstorydev/grpcui)
- [docker](https://docs.docker.com/get-docker/)
- [earthly](https://earthly.dev/get-earthly)
- [buf](https://docs.buf.build/installation)

## Status

- [X] __Alpha__: Closed testing.
- [X] __Private Devnet__: May have kinks. [See projects](https://sonr.io/dashboard).
- [ ] __Public Testnet__: Stable for non-enterprise use. [Join it](https://sonr.io/dashboard).
- [ ] __Mainnet__: Coming soon. [Watch status](https://sonr.io/docs/guides/getting-started/features#feature-status).

We are currently in transitioning to Public Testnet. Watch [releases](https://github.com/sonrhq/sonr/releases) of this repo to get notified of major updates.

## Architecture

Sonr is a [blockchain node](https://sonr.io/dashboard) which you can run locally, or use to join our testnet. You can sign up and start using Sonr without installing anything using our [dashboard](https://sonr.io/dashboard).

![Architecture](.github/assets/architecture.svg)

See [additional details](https://sonr.io/whitepaper) on these components in the whitepaper.

---

## Community & Support

- [Forum](https://github.com/sonrhq/sonr/discussions)
- [Issues](https://github.com/sonrhq/sonr/issues)
- [Twitter](https://sonr.io/twitter)
- [Dev Chat](https://sonr.io/discord)

## Acknowledgements

Sonr would not have been possible without the direct and indirect support of the following organizations and individuals:

- __Protocol Labs__: For IPFS & Libp2p.
- __Interchain Foundation__: For Cosmos & IBC.
- __Tim Berners-Lee__: For the Internet.
- __Satoshi Nakamoto__: For Bitcoin.
- __Steve Jobs__: For Taste.

Thank you for your support and inspiration!
