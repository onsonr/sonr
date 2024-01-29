---
description: >-
  A specification for managing foreign DNS Domains as relaying parties and
  services in the Sonr network.
---

# ADR-003: Service Records

## O**bjective**

* Alleviate decentralized organization management with on-chain implementation
* Enforce Compliance with WebAuthn Relaying Party Specification for consistency
* Enable Domain Origin support with DNS TXT Record Verification

***

## Solution

#### Service Record Primitive

#### On-Chain Multisig Accounts for Organizations

#### Traditional & Handshake (HNS) DNS Verification

***

## Definitions

*   `DIDCommMessaging`

    Description of the new term
*   `Relaying Party`

    Description of the new term
*   `Permissions`

    Description of the new term

***

## Sequence Methods

#### 1. Creating a Sonr Developer Account

The Blockchain creates an on-chain multisig using the `x/group` module and sets the delegator as the owner.

* During Testnet; this is alleviated by requests to the Faucet airdrop
* Stake **$SNR 500** to unlock elevated developer permissions (i.e. Creating scoped Personal Access Tokens, Registering Service Records, etc.)

#### 2. Configuring a Service

Upon successful authentication a New record of the event is stored on chain with an encrypted fingerprint and the account which was authenticated.

* Permission scope for app determined by utilized Network features
* Validator initializes record with parameters for functioning as a Relaying party and the minimum user identifier requirements

***

## Economic Impact

### Network Fees

|                                                                | Sender        | Receiver  |
| -------------------------------------------------------------- | ------------- | --------- |
| Authenticating with a Service                                  | Service Owner | Validator |
| Open an encrypted channel                                      | Service Owner | Validator |
| Register a service record                                      | Service Owner | Treasury  |
| Verify a Signature                                             | Service Owner | Validator |
| Send a Message to a Users Inbox                                | Service Owner | Validator |
| Issue a Credential for the User on a Service                   | Service Owner | Validator |
| Validate a Credential for the User on a Service                | Service Owner | Validator |
| Lookup User Identifier existence in Zero-knowledge Accumulator | Service Owner | Validator |

### Staking Incentives

| Action                   | Minimum Delegation | Unlock Period |
| ------------------------ | ------------------ | ------------- |
| Persisting a Username    | USNR 200,000,000   | 30 Days       |
| Elevate Developer Access | USNR 500,000,000   | 12 Months     |

***

## Implementation

***

## Status

This proposal is **under development** by the core Sonr Team.

| Development Phase | Devnet  |
| ----------------- | ------- |
| Target Completion | Q4 2023 |

***

## References

* [Sonr ADR-001: Decentralized Identity System](https://www.notion.so/ADR-002-Decentralized-Identity-Specification-01102d0fa712448b8893fe1bdc689d1e?pvs=21)
* [Sonr ADR-003: Decentralized Network-Relaying Services](https://www.notion.so/ADR-003-Authoritative-Application-Records-9b579f508d14454bbe995c9dc430c345?pvs=21)
