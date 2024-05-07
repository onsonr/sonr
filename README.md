<div align="center">
<img src="https://pub-97e96d678cb448969765e4c1542e675a.r2.dev/github-sonr.png" width="256" height="256" />

---

<div style="text-align: center;">

[![Go Reference](https://pkg.go.dev/badge/github.com/didao-org/sonr.svg)](https://pkg.go.dev/github.com/didao-org/sonr)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/didao-org/sonr)
![GitHub Release Date - Published_At](https://img.shields.io/github/release-date/didao-org/sonr)
[![Static Badge](https://img.shields.io/badge/homepage-sonr.io-blue?style=flat-square)](https://sonr.io)
![Discord](https://img.shields.io/discord/843061375160156170?logo=discord&label=Discord%20Chat)

[![Go Report Card](https://goreportcard.com/badge/github.com/didao-org/sonr)](https://goreportcard.com/report/github.com/didao-org/sonr)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=sonrhq_sonr&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=sonrhq_sonr&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=sonr-io_sonr)
[![Mutable.ai Auto Wiki](https://img.shields.io/badge/Auto_Wiki-Mutable.ai-blue)](https://wiki.mutable.ai/di-dao/sonr)

</div>

## `sonr` - Docs, Proposals, & Networks

Sonr is a combination of decentralized primitives. Fundamentally, it is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS — providing users with a secure, portable decentralized identity.

<table style="table-layout:fixed; white-space: nowrap;">
  <tr>
    <th>Module</th>
    <th colspan=4>Description</th>
  </tr>
  <tr>
    <th><code><a href="https://github.com/didao-org/sonr/x/identity">x/identity</a></code></th>
    <th colspan=4>
    The Sonr Identity module is responsible for managing DID based <br />
    accounts using the MPC Protocol - <a href="https://sonr.io/whitepaper">Docs</a>
    </th>
  </tr>
  <tr>
    <th><code>x/oracle</code></th>
    <th colspan=4>
    The Oracle module is responsible for managing Staking delegations <br />
    rewards, and token transfers - <a href="https://sonr.io/whitepaper">Docs</a>
    </th>
  </tr>
  <tr>
    <th><code><a href="https://github.com/didao-org/sonr/x/service">x/service</a></code></th>
    <th colspan=4>
    The Service module is responsible for DAO Application Service <br />
    Configurations, and Passkey authentication - <a href="https://sonr.io/whitepaper">Docs</a>
    </th>
  </tr>
</table>

Sonr is built on top of the Cosmos SDK, which is a framework for building blockchain applications in Golang. We have built the above modules to provide a decentralized identity and asset management system.

## Usage

It's recommended to install the following tools:

-   [golang](https://golang.org/doc/install)
-   [grpcui](https://github.com/fullstorydev/grpcui)
-   [docker](https://docs.docker.com/get-docker/)
-   [spawn](https://github.com/rollchains/spawn)

## Status

-   [x] **Alpha**: Closed testing.
-   [x] **Private Devnet**: May have kinks. [See projects](https://sonr.io/dashboard).
-   [ ] **Public Testnet**: Stable for non-enterprise use. [Join it](#).
-   [ ] **Mainnet**: Coming soon. [Watch status](#).

We are currently in transitioning to Public Testnet. Watch [releases](https://github.com/didao-org/sonr/releases) of this repo to get notified of major updates.

## Architecture

Sonr is a [blockchain node](https://sonr.io/dashboard) which you can run locally, or use to join our testnet. You can sign up and start using Sonr without installing anything using our [dashboard](https://sonr.io/dashboard).

![Architecture](.github/assets/public/img/architecture.svg)

See [additional details](https://sonr.io/whitepaper) on these components in the whitepaper.

---

## Community & Support

-   [Forum](https://github.com/didao-org/sonr/discussions)
-   [Issues](https://github.com/didao-org/sonr/issues)
-   [Twitter](https://sonr.io/twitter)
-   [Dev Chat](https://sonr.io/discord)

## Acknowledgements

Sonr would not have been possible without the direct and indirect support of the following organizations and individuals:

-   **Protocol Labs**: For IPFS & Libp2p.
-   **Interchain Foundation**: For Cosmos & IBC.
-   **Tim Berners-Lee**: For the Internet.
-   **Satoshi Nakamoto**: For Bitcoin.
-   **Steve Jobs**: For Taste.

Thank you for your support and inspiration!
