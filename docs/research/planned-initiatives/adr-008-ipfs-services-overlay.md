# ADR-008: IPFS Services Overlay

## Context

Add relevant information and links in under 4 sentences to give clarity to the challenge this proposal is fixing.

***

## O**bjective**

* Insert goal #1
* Insert goal #2
* Insert goal #3

***

## Solution

1 Sentence elevator pitch for this proposal explaining the improved changes.

#### Solution Detail Title

More on detail number one.

#### Solution Detail Title

More on detail number two.

***

## Definitions

*   `UCAN`

    Description of the new term
*   `CID`

    Description of the new term
*   IceFireDB

    Description of the new term
*

***

## Sequence Methods

### Sequence Method #1

More information on sequence method one

***

## Economic Impact

### Network Fees

|                                                                | Sender        | Receiver  | Fees       |
| -------------------------------------------------------------- | ------------- | --------- | ---------- |
| Authenticating with a Service                                  | Service Owner | Validator | SNR 1      |
| Open an encrypted channel                                      | Service Owner | Validator | SNR 0.25/s |
| Register a service record                                      | Service Owner | Treasury  | SNR 100    |
| Verify a Signature                                             | Service Owner | Validator | SNR 0.1    |
| Send a Message to a Users Inbox                                | Service Owner | Validator | SNR 2      |
| Issue a Credential for the User on a Service                   | Service Owner | Validator | SNR 3      |
| Validate a Credential for the User on a Service                | Service Owner | Validator | SNR 1      |
| Lookup User Identifier existence in Zero-knowledge Accumulator | Service Owner | Validator | SNR 0.5    |

### Staking Incentives

| Action                   | Minimum Delegation | Unlock Period |
| ------------------------ | ------------------ | ------------- |
| Persisting a Username    | USNR 200,000,000   | 30 Days       |
| Elevate Developer Access | USNR 500,000,000   | 12 Months     |

***

## Implementation

### API Methods

| Summary                     | Method | Endpoint        |
| --------------------------- | ------ | --------------- |
| Query for an Identity       | GET    | /core/identity/ |
| Send a transaction on-chain | POST   | /core/identity/ |

### Status

This proposal is **under development** by the core Sonr Team.

| Development Phase | Devnet  |
| ----------------- | ------- |
| Target Completion | Q4 2023 |

***

## References

* [Sonr ADR-001: Decentralized Identity System](https://www.notion.so/ADR-002-Decentralized-Identity-Specification-01102d0fa712448b8893fe1bdc689d1e?pvs=21)
* [Sonr ADR-003: Decentralized Network-Relaying Services](https://www.notion.so/ADR-003-Authoritative-Application-Records-9b579f508d14454bbe995c9dc430c345?pvs=21)
