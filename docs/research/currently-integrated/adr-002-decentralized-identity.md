---
description: >-
  Sonr aims to set a new standard in digital identity by aligning with W3C
  specifications for Decentralized Identifiers (DIDs).
---

# ADR-002: Decentralized Identity

## Context

Sonr aims to set a new standard in digital identity by aligning with W3C specifications for Decentralized Identifiers (DIDs). The current system's reliance on third-party identification and seed-phrases for wallet onboarding presents privacy risks and usability challenges. This proposal aims to eliminate these issues by enhancing our **`x/identity`** module.

## O**bjective**

* Enforce Compliance by implementing Decentralized Identifiers from W3C
* Prevent Applications and Services from associating 3rd-party identification with wallets
* Eliminate cumbersome onboarding by removing seed-phrases

***

## Solution

The `x/identity` module is responsible for the management of DIDDocuments on-chain and resolution of private user identifiers off-chain. In addition to standard type definitions, this module provides a gateway for Sonr clients to be able to authorize wallet accounts over a REST API.

#### Multi-party Computation

The multi-party computation feature aims to enhance the security and privacy of decentralized identity systems. It allows for secure computation on private user identifiers without exposing them to any single party.

#### Interblockchain Communication

The interblockchain communication or IBC, enables the exchange of information and data between different blockchain networks. When using the Cosmos SDK, this functionality is offered out of the box. Resulting in interoperability and seamless integration between different IBC Protocol enabled blockchains.

***

## Definitions

*   `**Decentralized Identifier (DID)**`

    A globally unique identifier that enables verifiable, self-sovereign digital identities. DIDs are fully under the control of the DID subject, independent from any centralized registry, identity provider, or certificate authority.
*   `**DID Method**`

    Specifies the syntax and procedures for specific DID schemes. It defines how to read, write and resolve DIDs for a particular blockchain or storage network.
*   `**DID Document**`

    A JSON-LD document that describes the DID, including its public keys, authentication protocols, and service endpoints for interaction.
*   `**DID Controller**`

    An entity that has the capability to make changes to a DID Document, essentially having control over the DID.
*   `**DID Resolver**`

    A system function that takes a DID as input and produces a DID Document as output, essentially translating the identifier into a form that the system can process.

***

## Sequence Methods

#### 1. Self Signed Registration

More information on sequence method one

* More information on sequence method one
* More information on sequence method one

#### 2. Service-Proxied Registration

![Service-Proxied Registration Diagram.svg](https://prod-files-secure.s3.us-west-2.amazonaws.com/b4e83706-0f19-4f5e-9020-40fcf1b9dda3/6b2a62a6-dbf2-42fd-a54f-57917c052986/Service-Proxied\_Registration\_Diagram.svg)

More information on sequence method one.

* More information on sequence method one
* More information on sequence method one

#### 3. Resolving Decentralized Identifiers

![Resolution Process Diagram.svg](https://prod-files-secure.s3.us-west-2.amazonaws.com/b4e83706-0f19-4f5e-9020-40fcf1b9dda3/c50528ef-1374-408e-8bb0-8a41c1b4bc2e/Resolution\_Process\_Diagram.svg)

More information on sequence method one

* More information on sequence method one
* More information on sequence method one

***

### Keyshare Mapping

[https://app.eraser.io/workspace/XRrdDcwLV7mOFOhhGWcH?origin=share](https://app.eraser.io/workspace/XRrdDcwLV7mOFOhhGWcH?origin=share)

| Term           | Definition                                                                                  |
| -------------- | ------------------------------------------------------------------------------------------- |
| Alice          | Represents encrypted data stored on-chain and backed up in IPFS.                            |
| Bob            | Encrypted by a Zero Knowledge (zk) accumulator using traditional user-provided identifiers. |
| Pub            | Public information or public key.                                                           |
| SecData        | Secret data or private key information.                                                     |
| Curve          | Details about the cryptographic curve used.                                                 |
| Keyshare       | A fragment of a cryptographic key.                                                          |
| DID Controller | The controlling entity that interacts with the keyshares from both Alice and Bob.           |

## Economic Impact

### Value Generation

### Staking Incentives

| Action                   | Minimum Delegation | Unlock Period |
| ------------------------ | ------------------ | ------------- |
| Persisting a Username    | USNR 200,000,000   | 30 Days       |
| Elevate Developer Access | USNR 500,000,000   | 12 Months     |

***

## Implementation

### Formatting Identifiers

### Supported Methods

`did:idx`

### Document Model

### Multi-Party Computation

***

## Status

This proposal is **under development** by the core Sonr Team.

| Development Phase | Testnet |
| ----------------- | ------- |
| Target Completion | Q4 2023 |

***

## References

* [Decentralized Identifiers by W3C](https://www.w3.org/TR/did-core/)
