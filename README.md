<div style="text-align: center;">

[![Sonr Logo Banner](.github/images/core-cover.png)](https://sonr.io)

</div>
<div style="text-align: left;">

[![Go Reference](https://pkg.go.dev/badge/github.com/sonrhq/sonr.svg)](https://pkg.go.dev/github.com/sonrhq/sonr)
![Discord](https://img.shields.io/discord/843061375160156170?logo=discord&label=Dev%20Chat)
[![GitHub all releases](https://img.shields.io/github/downloads/sonrhq/sonr/total)](https://github.com/sonrhq/sonr/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonrhq/sonr)](https://goreportcard.com/report/github.com/sonrhq/sonr)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=sonr-io_sonr&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
[![Test sonrd](https://github.com/sonrhq/sonr/actions/workflows/tests.yml/badge.svg)](https://github.com/sonrhq/sonr/actions/workflows/tests.yml)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=sonr-io_sonr&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/sonr_io?style=social&logo=twitter)
</div>

---

# `sonrhq/sonr`

Sonr is an ibc-enabled blockchain for decentralized identity.

- [x] Passkey based User Accounts. [__Docs__](https://sonr.io/docs/guides/database)
- [x] DKLS-MPC Powered Wallets _(No seed phrases)_. [__Docs__](https://sonr.io/docs/guides/auth)
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

## How it works

Sonr is a combination of decentralized primitives. Fundamentally, it is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS — providing users with a secure, portable decentralized identity.

<table style="table-layout:fixed; white-space: nowrap;">
  <tr>
    <th>Module</th>
    <th colspan=4>Description</th>
    <th>Status</th>
    <th>Deps</th>
  </tr>
  <tr>
    <th><code><a href="https://github.com/sonrhq/identity">x/identity</a></code></th>
    <th colspan=4>
    The Sonr Identity module is responsible for managing DID based <br />
    accounts using the MPC Protocol - <a href="https://sonr.io/whitepaper">Docs</a>
    </th>
    <th>
      <a href="https://github.com/sonrhq/identity/actions/workflows/ci.yml?query=branch%3Amaster++">
        <img src="https://github.com/sonrhq/identity/actions/workflows/ci.yml/badge.svg?branch=master" alt="CI Pipeline">
      </a>
    </th>
    <th>
      <code>x/auth</code>
      <br />
      <code>x/ibcaccounts</code>
      <br />
      <code>x/slashing</code>
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
    <th colspan=2>
      <code>x/bank</code>
      <br />
      <code>x/distribution</code>
      <br />
      <code>x/ibctransfer</code>
      <br />
      <code>x/staking</code>
    </th>
  </tr>
  <tr>
    <th><code><a href="https://github.com/sonrhq/service">x/service</a></code></th>
    <th colspan=4>
    The Service module is responsible for DAO Application Service <br />
    Configurations, and Passkey authentication - <a href="https://sonr.io/whitepaper">Docs</a>
    </th>
    <th>
      <a href="https://github.com/sonrhq/service/actions/workflows/ci.yml?query=branch%3Amaster++">
        <img src="https://github.com/sonrhq/service/actions/workflows/ci.yml/badge.svg?branch=master" alt="CI Pipeline">
      </a>
    </th>
    <th colspan=2>
      <code>x/group</code>
      <br />
      <code>x/identity</code>
      <br />
      <code>x/wasm</code>
    </th>
  </tr>
</table>

Sonr is built on top of the Cosmos SDK, which is a framework for building blockchain applications in Golang. We have built the above modules to provide a decentralized identity and asset management system.

## Documentation

For full documentation, visit [sonr.io/docs](https://sonr.io/docs). To see how to Contribute, visit [Getting Started](./docs/contribution/DEVELOPERS.md)

Its reccomended to install the following tools:

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

![Architecture](.github/images/architecture.svg)

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
