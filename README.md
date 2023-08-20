<div style="text-align: center;">

[![Sonr Logo Banner](.github/images/core-cover.png)](https://sonr.io)
[![Go Reference](https://pkg.go.dev/badge/github.com/sonrhq/core.svg)](https://pkg.go.dev/github.com/sonrhq/core)
[![Test sonrd](https://github.com/sonrhq/core/actions/workflows/tests.yml/badge.svg)](https://github.com/sonrhq/core/actions/workflows/tests.yml)
[![Release sonrd](https://github.com/sonrhq/core/actions/workflows/release.yml/badge.svg)](https://github.com/sonrhq/core/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonrhq/core)](https://goreportcard.com/report/github.com/sonrhq/core)

</div>

---

# Sonr

Sonr is an ibc-enabled blockchain for decentralized identity.

- [x] Passkey based User Accounts. [Docs](https://sonr.io/docs/guides/database)
- [x] DKLS MPC Powered Wallets _(No seed phrases)_. [Docs](https://sonr.io/docs/guides/auth)
- [x] IPFS Database and Storage. [Docs](https://sonr.io/docs/guides/storage)
  - [x] Redis. [Docs](https://sonr.io/docs/guides/api#rest-api-overview)
  - [x] MySQL. [Docs](https://sonr.io/docs/guides/api#graphql-api-overview)
  - [ ] User mailboxes. [Docs](https://sonr.io/docs/guides/api#realtime-api-overview)
  - [ ] Realtime subscriptions. [Docs](https://sonr.io/docs/guides/api#realtime-api-overview)
  - [ ] Matrix chat. [Docs](https://sonr.io/docs/guides/api#realtime-api-overview)
- [x] IBC Integrations.
  - [ ] Nomic. [Docs](https://sonr.io/docs/guides/database/functions)
  - [ ] Evmos. [Docs](https://sonr.io/docs/guides/functions)
  - [ ] Osmosis. [Docs](https://sonr.io/docs/guides/functions)
- [x] Smart Contracts  _(CosmWasm)_. [Docs](https://sonr.io/docs/guides/storage)
- [ ] Typescript Client SDKs. [Docs](https://sonr.io/docs/guides/ai)
- [ ] Dashboard. [Repo](https://github.com/sonr-io/front)

## Documentation

For full documentation, visit [sonr.io/docs](https://sonr.io/docs)

To see how to Contribute, visit [Getting Started](./docs/wiki/DEVELOPERS.md)

## Community & Support

- [Community Forum](https://github.com/sonr-io/sonr/discussions). Best for: help with building, discussion about database best practices.
- [GitHub Issues](https://github.com/sonr-io/sonr/issues). Best for: bugs and errors you encounter using Sonr.
- [Email Support](https://sonr.io/docs/support#business-support). Best for: problems with your database or infrastructure.
- [Discord](https://discord.sonr.com). Best for: sharing your applications and hanging out with the community.

## Status

- [x] **Alpha**: We are testing Sonr with a closed set of customers
- [x] **Private Devnet**: Try it over at [sonr.com/dashboard](https://sonr.io/dashboard).
- [ ] **Public Testnet**: Stable enough for most non-enterprise use-cases, But go easy on us - there are a few kinks. Try it over at [sonr.io/dashboard](https://sonr.io/dashboard).
- [ ] **Mainnet**: General Availability with a DEX enlisted token, [watch status](https://sonr.io/docs/guides/getting-started/features#feature-status).

We are currently in Private Devnet. Watch "releases" of this repo to get notified of major updates.

## How it works

Sonr is a combination of decentralized primitives. Fundamentally, it is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS — providing users with a secure, portable decentralized identity.

Sonr is built on top of the Cosmos SDK, which is a framework for building blockchain applications in Golang. We use these modules:
- `x/auth`
- `x/bank`
- `x/staking`
- `x/slashing`
- `x/gov`
- `x/params`
- `x/distribution`
- `x/upgrade`
- `x/ibc`
- `x/ibc/applications/transfer`
- `x/wasm`


### Architecture

Sonr is a [blockchain node](https://sonr.io/dashboard) which you can run locally, or use to join our testnet. You can sign up and start using Sonr without installing anything using our [dashboard](https://sonr.io/dashboard).

![Architecture](.github/images/architecture.svg)

See [additional details](https://sonr.io/docs) on these components.

### Client libraries

Our approach for client libraries is uniform. Abstract away any blockchain specific details, and provide a simple interface for developers to use. We have a few client libraries that we maintain, and provide [guidelines](#) for community maintained libraries.

<table style="table-layout:fixed; white-space: nowrap;">
  <tr>
    <th>Language</th>
    <th>Client</th>
    <th colspan="5">Feature-Clients (bundled in Sonr client)</th>
  </tr>
  <!-- notranslate -->
  <tr>
    <th></th>
    <th>Sonr</th>
    <th><a href="https://github.com/postgrest/postgrest" target="_blank" rel="noopener noreferrer">PostgREST</a></th>
    <th><a href="https://github.com/sonr/realtime" target="_blank" rel="noopener noreferrer">Realtime</a></th>
    <th><a href="https://github.com/sonr/storage-api" target="_blank" rel="noopener noreferrer">Storage</a></th>
    <th>Functions</th>
  </tr>
  <!-- TEMPLATE FOR NEW ROW -->
  <!-- START ROW
  <tr>
    <td>lang</td>
    <td><a href="https://github.com/sonr-community/sonr-lang" target="_blank" rel="noopener noreferrer">sonr-lang</a></td>
    <td><a href="https://github.com/sonr-community/postgrest-lang" target="_blank" rel="noopener noreferrer">postgrest-lang</a></td>
    <td><a href="https://github.com/sonr-community/realtime-lang" target="_blank" rel="noopener noreferrer">realtime-lang</a></td>
    <td><a href="https://github.com/sonr-community/storage-lang" target="_blank" rel="noopener noreferrer">storage-lang</a></td>
  </tr>
  END ROW -->
  <!-- /notranslate -->
  <th colspan="7">⚡️ Official ⚡️</th>
  <!-- notranslate -->
  <tr>
    <td>JavaScript (TypeScript)</td>
    <td><a href="https://github.com/sonr-io/sonr-js" target="_blank" rel="noopener noreferrer">sonr-js</a></td>
    <td><a href="https://github.com/sonr/postgrest-js" target="_blank" rel="noopener noreferrer">postgrest-js</a></td>
    <td><a href="https://github.com/sonr/realtime-js" target="_blank" rel="noopener noreferrer">realtime-js</a></td>
    <td><a href="https://github.com/sonr/storage-js" target="_blank" rel="noopener noreferrer">storage-js</a></td>
    <td><a href="https://github.com/sonr/functions-js" target="_blank" rel="noopener noreferrer">functions-js</a></td>
  </tr>
    <tr>
    <td>Flutter</td>
    <td><a href="https://github.com/sonr-io/sonr-flutter" target="_blank" rel="noopener noreferrer">sonr-flutter</a></td>
    <td><a href="https://github.com/sonr/postgrest-dart" target="_blank" rel="noopener noreferrer">postgrest-dart</a></td>
    <td><a href="https://github.com/sonr/realtime-dart" target="_blank" rel="noopener noreferrer">realtime-dart</a></td>
    <td><a href="https://github.com/sonr/storage-dart" target="_blank" rel="noopener noreferrer">storage-dart</a></td>
    <td><a href="https://github.com/sonr/functions-dart" target="_blank" rel="noopener noreferrer">functions-dart</a></td>
  </tr>
  <!-- /notranslate -->
  <th colspan="7">💚 Community 💚</th>
  <!-- notranslate -->
  <tr>
    <td>C#</td>
    <td><a href="https://github.com/sonr-community/sonr-csharp" target="_blank" rel="noopener noreferrer">sonr-csharp</a></td>
    <td><a href="https://github.com/sonr-community/postgrest-csharp" target="_blank" rel="noopener noreferrer">postgrest-csharp</a></td>
    <td><a href="https://github.com/sonr-community/realtime-csharp" target="_blank" rel="noopener noreferrer">realtime-csharp</a></td>
    <td><a href="https://github.com/sonr-community/storage-csharp" target="_blank" rel="noopener noreferrer">storage-csharp</a></td>
    <td><a href="https://github.com/sonr-community/functions-csharp" target="_blank" rel="noopener noreferrer">functions-csharp</a></td>
  </tr>
  <tr>
    <td>Go</td>
    <td>-</td>
    <td><a href="https://github.com/sonr-community/postgrest-go" target="_blank" rel="noopener noreferrer">postgrest-go</a></td>
    <td>-</td>
    <td><a href="https://github.com/sonr-community/storage-go" target="_blank" rel="noopener noreferrer">storage-go</a></td>
    <td><a href="https://github.com/sonr-community/functions-go" target="_blank" rel="noopener noreferrer">functions-go</a></td>
  </tr>
  <tr>
    <td>Java</td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
    <td><a href="https://github.com/sonr-community/storage-java" target="_blank" rel="noopener noreferrer">storage-java</a></td>
    <td>-</td>
  </tr>
  <tr>
    <td>Swift</td>
    <td><a href="https://github.com/sonr-community/sonr-swift" target="_blank" rel="noopener noreferrer">sonr-swift</a></td>
    <td><a href="https://github.com/sonr-community/postgrest-swift" target="_blank" rel="noopener noreferrer">postgrest-swift</a></td>
    <td><a href="https://github.com/sonr-community/realtime-swift" target="_blank" rel="noopener noreferrer">realtime-swift</a></td>
    <td><a href="https://github.com/sonr-community/storage-swift" target="_blank" rel="noopener noreferrer">storage-swift</a></td>
    <td><a href="https://github.com/sonr-community/functions-swift" target="_blank" rel="noopener noreferrer">functions-swift</a></td>
  </tr>
  <!-- /notranslate -->
</table>

---

## Acknowledgements

Sonr would not have been possible without the direct and indirect support of the following organizations and individuals:

- **Protocol Labs**: For IPFS & Libp2p.
- **Interchain Foundation**: For Cosmos & IBC.
- **Tim Berners-Lee**: For the Internet.
- **Satoshi Nakamoto**: For Bitcoin.
- **Steve Jobs**: For Taste.

Thank you for your support and inspiration!
