---
title: Overview
id: adr-overview
displayed_sidebar: resourcesSidebar
---

# ADR Guide
This guide serves as a walkthrough for ADRs 001-006. It provides a high level overview of each ADR with a special emphasis on interactions with the blockchain.

## Abstract
An Architecture Decision Record (ADR) is a document which describes in detail a significant addition to the Sonr platform. This includes a high level description all the way down to specific methods to be added. This guide will link each ADR for reference, but the blockchain interactions for each ADR will be synthesized and documented here.

### [ADR-001: Decentralized Identifiers](https://docs.sonr.io/posts/architecture/adr-001)

[Decentralized Identifiers (DIDs)](http://w3.org/TR/did-core) are the foundation of decentralization on Sonr. They provide a standardized way to support unfederated identities that are not owned by any individual organization—including Sonr.
	Any time a user or application is created or updated, a transaction on the blockchain is created. This may mean changing metadata for an account, such as a profile picture link; adding or removing a device (new phone, laptop, etc.); buying or transferring an alias, for example “alice.snr”; or authorizing another device as a controller, for example in the case you need to let an application act on your behalf. These transactions will always be done on behalf of the end user.

### [ADR-002: Object Schemas](https://docs.sonr.io/posts/architecture/adr-002)

Schemas and Objects together are the primary means for persistent data on Sonr. Schemas are data types, implemented as [IPLD](https://ipld.io) schemas that enable platform-agnostic models within Sonr and without. Objects on the other hand are simply data blobs on [IPFS](https://ipfs.io) (not to be confused with IPLD) that can be deserialized and represented by a schema.
	Schemas will be largely immutable, but creating them requires a transaction on the blockchain. Any updates to a schema are made in the form of a new schema and a deprecation of the old one. Objects by contrast are not stored on-chain. The only interaction between objects and the chain arises when buckets are introduced (more on that in the next section). These transactions will be done mostly by developers.

### [ADR-003: Buckets](https://docs.sonr.io/posts/architecture/adr-003)

Objects on Sonr will oftentimes be related to one another. For example, a user will have a vault containing all their encryption keys, a group of objects related to a particular application they use, or just personal information. These related objects will be tied together through buckets. A bucket is essentially a list of [CIDs](https://docs.ipfs.io/concepts/content-addressing/) or a set of small data blobs that is owned by an end user. Buckets also facilitate access control by acting as a central repository for shared encryption keys.
	Creating, updating, and deleting buckets all require transactions on the blockchain. This is likely to be done very frequently, with the cost on the end user.

### [ADR-004: Channels](https://docs.sonr.io/posts/architecture/adr-004)

The purpose of channels is to facilitate real-time communication between nodes on the network. In many cases, these channels need only exist for a short time to transfer data from one node to another or establish a connection—these are ephemeral channels. Ephemeral channels require no blockchain resources and no persisted data. Persistent channels on the other hand are created on-chain and never become unavailable unless deactivated. These channels are meant to facilitate things akin to Discord channels or notifications.

### [ADR-005: NFT Standard](https://docs.sonr.io/posts/architecture/adr-005)

This is just what it sounds like. Sonr will use the [CW721](https://docs.cosmwasm.com/cw-plus/0.9.0/cw721/spec/) Spec for implementing NFTs, enabling developers to build [CosmWasm](https://cosmwasm.com/) smart contracts for custom NFT logic. A transaction on-chain will be required for any mint, buy, sell, or transfer of NFTs.

### [ADR-006: Functions](https://docs.sonr.io/posts/architecture/adr-006)

Functions on Sonr allow developers to run serverless code, similar to AWS Lambda or GCP Cloud Functions. Creating, updating, and deleting these functions will involve a transaction to the chain, cost of execution will be determined from one of the following ways:
- total_cost = static_cost + runtime_fees (static_cost)
- total_cost=variable_cost+runtime_fees+ gas_fees (variable cost)
- total_cost=runtime_fees (runtime_cost) but execution is facilitated by an off-chain service (free). Developers will be the only users performing these transactions.
