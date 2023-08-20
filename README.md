<div style="text-align: center;">

[![Sonr Logo Banner](.github/images/core-cover.png)](https://sonr.io)
[![Go Reference](https://pkg.go.dev/badge/github.com/sonrhq/core.svg)](https://pkg.go.dev/github.com/sonrhq/core)
[![Test sonrd](https://github.com/sonrhq/core/actions/workflows/tests.yml/badge.svg)](https://github.com/sonrhq/core/actions/workflows/tests.yml)
[![Release sonrd](https://github.com/sonrhq/core/actions/workflows/release.yml/badge.svg)](https://github.com/sonrhq/core/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonrhq/core)](https://goreportcard.com/report/github.com/sonrhq/core)

</div>

---

# Sonr

Sonr is a <strong>peer-to-peer identity</strong> and <strong>asset management system</strong> that leverages <italic>DID Documents, WebAuthn, and IPFS</italic> - to provide users with a <strong>secure, user-friendly</strong> way to manage their <strong>digital identity and assets.</strong>

- [x] Hosted Postgres Database. [Docs](https://sonr.io/docs/guides/database)
- [x] Authentication and Authorization. [Docs](https://sonr.io/docs/guides/auth)
- [x] Auto-generated APIs.
  - [x] REST. [Docs](https://sonr.io/docs/guides/api#rest-api-overview)
  - [x] GraphQL. [Docs](https://sonr.io/docs/guides/api#graphql-api-overview)
  - [x] Realtime subscriptions. [Docs](https://sonr.io/docs/guides/api#realtime-api-overview)
- [x] Functions.
  - [x] Database Functions. [Docs](https://sonr.io/docs/guides/database/functions)
  - [x] Edge Functions [Docs](https://sonr.io/docs/guides/functions)
- [x] File Storage. [Docs](https://sonr.io/docs/guides/storage)
- [x] AI + Vector/Embeddings Toolkit. [Docs](https://sonr.io/docs/guides/ai)
- [x] Dashboard

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

Sonr is a combination of open source tools. We’re building the features of Firebase using enterprise-grade, open source products. If the tools and communities exist, with an MIT, Apache 2, or equivalent open license, we will use and support that tool. If the tool doesn't exist, we build and open source it ourselves. Sonr is not a 1-to-1 mapping of Firebase. Our aim is to give developers a Firebase-like developer experience using open source tools.

### Architecture

Sonr is a [decentralized platform](https://sonr.io/dashboard). You can sign up and start using Sonr without installing anything.
You can also [self-host](https://sonr.io/docs/guides/hosting/overview) and [develop locally](https://sonr.io/docs/guides/local-development).

![Architecture](.github/images/architecture.svg)

- [L1 Blockchain](https://www.postgresql.org/) is an object-relational database system with over 30 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance.
- [API Gateway](https://github.com/sonr/realtime) is an Elixir server that allows you to listen to PostgreSQL inserts, updates, and deletes using websockets. Realtime polls Postgres' built-in replication functionality for database changes, converts changes to JSON, then broadcasts the JSON over websockets to authorized clients.
- [CosmWasm Contracts](http://postgrest.org/) is a web server that turns your PostgreSQL database directly into a RESTful API
- [IcefireDB Redis/SQL](http://github.com/sonr/pg_graphql/) a PostgreSQL extension that exposes a GraphQL API
- [IBC Channel](https://github.com/sonr/postgres-meta) is a RESTful API for managing your Postgres, allowing you to fetch tables, add roles, and run queries, etc.
- [Matrix](https://github.com/sonr/gotrue) is an JWT based API for managing users and issuing JWT tokens.
- [Libp2p](https://github.com/Kong/kong) is a cloud-native API gateway.

### Client libraries

Our approach for client libraries is modular. Each sub-library is a standalone implementation for a single external system. This is one of the ways we support existing tools.

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
    <th><a href="https://github.com/sonr/gotrue" target="_blank" rel="noopener noreferrer">GoTrue</a></th>
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
    <td><a href="https://github.com/sonr-community/gotrue-lang" target="_blank" rel="noopener noreferrer">gotrue-lang</a></td>
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
    <td><a href="https://github.com/sonr/gotrue-js" target="_blank" rel="noopener noreferrer">gotrue-js</a></td>
    <td><a href="https://github.com/sonr/realtime-js" target="_blank" rel="noopener noreferrer">realtime-js</a></td>
    <td><a href="https://github.com/sonr/storage-js" target="_blank" rel="noopener noreferrer">storage-js</a></td>
    <td><a href="https://github.com/sonr/functions-js" target="_blank" rel="noopener noreferrer">functions-js</a></td>
  </tr>
    <tr>
    <td>Flutter</td>
    <td><a href="https://github.com/sonr-io/sonr-flutter" target="_blank" rel="noopener noreferrer">sonr-flutter</a></td>
    <td><a href="https://github.com/sonr/postgrest-dart" target="_blank" rel="noopener noreferrer">postgrest-dart</a></td>
    <td><a href="https://github.com/sonr/gotrue-dart" target="_blank" rel="noopener noreferrer">gotrue-dart</a></td>
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
    <td><a href="https://github.com/sonr-community/gotrue-csharp" target="_blank" rel="noopener noreferrer">gotrue-csharp</a></td>
    <td><a href="https://github.com/sonr-community/realtime-csharp" target="_blank" rel="noopener noreferrer">realtime-csharp</a></td>
    <td><a href="https://github.com/sonr-community/storage-csharp" target="_blank" rel="noopener noreferrer">storage-csharp</a></td>
    <td><a href="https://github.com/sonr-community/functions-csharp" target="_blank" rel="noopener noreferrer">functions-csharp</a></td>
  </tr>
  <tr>
    <td>Go</td>
    <td>-</td>
    <td><a href="https://github.com/sonr-community/postgrest-go" target="_blank" rel="noopener noreferrer">postgrest-go</a></td>
    <td><a href="https://github.com/sonr-community/gotrue-go" target="_blank" rel="noopener noreferrer">gotrue-go</a></td>
    <td>-</td>
    <td><a href="https://github.com/sonr-community/storage-go" target="_blank" rel="noopener noreferrer">storage-go</a></td>
    <td><a href="https://github.com/sonr-community/functions-go" target="_blank" rel="noopener noreferrer">functions-go</a></td>
  </tr>
  <tr>
    <td>Java</td>
    <td>-</td>
    <td>-</td>
    <td><a href="https://github.com/sonr-community/gotrue-java" target="_blank" rel="noopener noreferrer">gotrue-java</a></td>
    <td>-</td>
    <td><a href="https://github.com/sonr-community/storage-java" target="_blank" rel="noopener noreferrer">storage-java</a></td>
    <td>-</td>
  </tr>
  <tr>
    <td>Kotlin</td>
    <td><a href="https://github.com/sonr-community/sonr-kt" target="_blank" rel="noopener noreferrer">sonr-kt</a></td>
    <td><a href="https://github.com/sonr-community/sonr-kt/tree/master/Postgrest" target="_blank" rel="noopener noreferrer">postgrest-kt</a></td>
    <td><a href="https://github.com/sonr-community/sonr-kt/tree/master/GoTrue" target="_blank" rel="noopener noreferrer">gotrue-kt</a></td>
    <td><a href="https://github.com/sonr-community/sonr-kt/tree/master/Realtime" target="_blank" rel="noopener noreferrer">realtime-kt</a></td>
    <td><a href="https://github.com/sonr-community/sonr-kt/tree/master/Storage" target="_blank" rel="noopener noreferrer">storage-kt</a></td>
    <td><a href="https://github.com/sonr-community/sonr-kt/tree/master/Functions" target="_blank" rel="noopener noreferrer">functions-kt</a></td>
  </tr>
  <tr>
    <td>Python</td>
    <td><a href="https://github.com/sonr-community/sonr-py" target="_blank" rel="noopener noreferrer">sonr-py</a></td>
    <td><a href="https://github.com/sonr-community/postgrest-py" target="_blank" rel="noopener noreferrer">postgrest-py</a></td>
    <td><a href="https://github.com/sonr-community/gotrue-py" target="_blank" rel="noopener noreferrer">gotrue-py</a></td>
    <td><a href="https://github.com/sonr-community/realtime-py" target="_blank" rel="noopener noreferrer">realtime-py</a></td>
    <td><a href="https://github.com/sonr-community/storage-py" target="_blank" rel="noopener noreferrer">storage-py</a></td>
    <td><a href="https://github.com/sonr-community/functions-py" target="_blank" rel="noopener noreferrer">functions-py</a></td>
  </tr>
  <tr>
    <td>Ruby</td>
    <td><a href="https://github.com/sonr-community/sonr-rb" target="_blank" rel="noopener noreferrer">sonr-rb</a></td>
    <td><a href="https://github.com/sonr-community/postgrest-rb" target="_blank" rel="noopener noreferrer">postgrest-rb</a></td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
  </tr>
  <tr>
    <td>Rust</td>
    <td>-</td>
    <td><a href="https://github.com/sonr-community/postgrest-rs" target="_blank" rel="noopener noreferrer">postgrest-rs</a></td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
    <td>-</td>
  </tr>
  <tr>
    <td>Swift</td>
    <td><a href="https://github.com/sonr-community/sonr-swift" target="_blank" rel="noopener noreferrer">sonr-swift</a></td>
    <td><a href="https://github.com/sonr-community/postgrest-swift" target="_blank" rel="noopener noreferrer">postgrest-swift</a></td>
    <td><a href="https://github.com/sonr-community/gotrue-swift" target="_blank" rel="noopener noreferrer">gotrue-swift</a></td>
    <td><a href="https://github.com/sonr-community/realtime-swift" target="_blank" rel="noopener noreferrer">realtime-swift</a></td>
    <td><a href="https://github.com/sonr-community/storage-swift" target="_blank" rel="noopener noreferrer">storage-swift</a></td>
    <td><a href="https://github.com/sonr-community/functions-swift" target="_blank" rel="noopener noreferrer">functions-swift</a></td>
  </tr>
  <tr>
    <td>Godot Engine (GDScript)</td>
    <td><a href="https://github.com/sonr-community/godot-engine.sonr" target="_blank" rel="noopener noreferrer">sonr-gdscript</a></td>
    <td><a href="https://github.com/sonr-community/postgrest-gdscript" target="_blank" rel="noopener noreferrer">postgrest-gdscript</a></td>
    <td><a href="https://github.com/sonr-community/gotrue-gdscript" target="_blank" rel="noopener noreferrer">gotrue-gdscript</a></td>
    <td><a href="https://github.com/sonr-community/realtime-gdscript" target="_blank" rel="noopener noreferrer">realtime-gdscript</a></td>
    <td><a href="https://github.com/sonr-community/storage-gdscript" target="_blank" rel="noopener noreferrer">storage-gdscript</a></td>
    <td><a href="https://github.com/sonr-community/functions-gdscript" target="_blank" rel="noopener noreferrer">functions-gdscript</a></td>
  </tr>
  <!-- /notranslate -->
</table>


## Acknowledgements

The Sonr project would not have been possible without the direct and indirect support of the following organizations and individuals:

#### Protocol Labs
> Sonr is built on top of the [IPFS](https://ipfs.io/) protocol. IPFS is a peer-to-peer hypermedia protocol designed to make the web faster, safer, and more open.

#### Interchain Foundation
> Sonr is built on top of the [Cosmos](https://cosmos.network/) protocol. Cosmos is a decentralized network of independent parallel blockchains, each powered by BFT consensus algorithms like Tendermint consensus.

#### Tim Berners-Lee
> Sonr is built on top of the [Solid](https://solidproject.org/) protocol. Solid is a specification that lets people store their data securely in decentralized data stores called Pods (personal online data stores) and lets users control who can access what parts of their data using Solid's authentication and authorization systems.

#### Satoshi Nakamoto
> Sonr is built on top of the [Bitcoin](https://bitcoin.org/) protocol. Bitcoin is a decentralized peer-to-peer electronic cash system that does not rely on any central authority like a government or financial institution.

#### "Apples Spirit"
> Sonr is built on top of the [Ethereum](https://ethereum.org/) protocol. Ethereum is a decentralized platform that runs smart contracts: applications that run exactly as programmed without any possibility of downtime, censorship, fraud or third-party interference.

---

We appreciate you for your inspiration and support.
