---
description: >-
  Sonr is a peer-to-peer identity and asset management system that leverages DID
  documents, Webauthn, and IPFS.
---

# ADR-001: Sonr Base Foundations

## Abstract

**Sonr is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS — providing users with a secure, user-friendly way to manage their digital identity and assets.**

Sonr is a Cosmos powered blockchain which is powered by a TenderMint validation mechanism. The default consensus for TenderMint is DPoS and works with our current ABCI implementation for Transaction Verification.

---

## Concepts

- `Delegated Proof of Stake (DPoS)`

  Delegated Proof of Stake (DPoS) is a consensus mechanism used in blockchain networks where token holders can vote for delegates to validate transactions and secure the network. It allows for faster transaction confirmations and energy efficiency compared to traditional Proof of Work (PoW) systems. This is the underlying mechanism Sonr uses in order to solve the BFT problem.

- `IPFS`
- `Matrix Protocol`

  The Matrix Protocol is a decentralized communication protocol that enables secure and interoperable messaging across different platforms and networks. It provides a framework for real-time communication and collaboration, allowing users to send messages, share files, and make voice and video calls.

  The protocol focuses on providing privacy, security, and end-to-end encryption for communication between users. It is designed to be open and extensible, allowing developers to build on top of it and create innovative applications. The Matrix Protocol aims to create a decentralized and federated communication network that is not controlled by any single entity.

- `Libp2p`
- `Motor Node`
- `Highway Protocol`

---

## Architecture

![Sonr Holistic Authentication Feature Set](https://prod-files-secure.s3.us-west-2.amazonaws.com/b4e83706-0f19-4f5e-9020-40fcf1b9dda3/60bb656f-a282-4fa1-a054-f1360a5f673a/diagram-export_10_6_2023.svg)

Sonr Holistic Authentication Feature Set

### Native Protocol Assets

On Sonr there are two tokens which operate within protocols and execution.

| SNR                                                          | USNR                                                                                               |
| ------------------------------------------------------------ | -------------------------------------------------------------------------------------------------- |
| Used for paying for operations provided on the Sonr network. | Used for staking claim in the network for elevated permissions and calling permissioned operations |
| $R(A)=1\_{SNR}$                                              | $1\_{SNR} = 1,000,000\_{uSNR}$                                                                     |

_Coin age is irrelevant. All coins that are mature will add the same staking weight — resulting in stable, consistent interest only for active wallets and only with small inputs._

### Stakeholder Roles & Responsibilities

**General Public**

- Vote on validator selection
- Vote on software upgrades
- Report Validators/Organizations

---

**Sonr Treasury**

- Propose/vote on new Grant themes
- Propose/Vote Software Upgrades
- Jail Validators/Organizations

**Organizations**

- Propose new permission types
- Propose/vote on validators
- Vote on Software upgrades

---

**Node Operators**

- Propose/vote on validators
- Propose/Vote New Permissions
- Propose/Vote Software Upgrades

### Core Blockchain Modules

| Module        | Summary                                                                | Author   |
| ------------- | ---------------------------------------------------------------------- | -------- |
| x/identity    | Manages identity primitives native to the Sonr blockchain.             | Sonr     |
| x/service     | Manages service record operations, validations, reporting, and gateway | Sonr     |
| x/oracle      | Manages the deposit of cross-chain tokens, and price feeds             | Sonr     |
| x/permissions | Manages the available permissions services can request from accounts   | Sonr     |
| x/bank        | Used to mint new token or to deposit cross-chain assets                | Cosmos   |
| x/capability  | Used to distinguish between different account/addresses                | Cosmos   |
| x/gov         | Used to establish Governance guidelines for operator nodes             | Cosmos   |
| x/group       | On-chain multi-signature accounts for Organizations                    | Cosmos   |
| x/slashing    | Used by Sonr foundation to facilitate bad actors on the network        | Cosmos   |
| x/upgrade     | Used for software binary upgrades across nodes                         | Cosmos   |
| x/ibcaccounts | Used to leverage IBC Host/Controller Accounts across chains            | IBC      |
| x/ibctransfer | Used to provide bridge-less transfers natively for accounts            | IBC      |
| x/wasm        | Used for deploying Rust smart contracts to the Sonr network.           | CosmWasm |

### Value-Demand Relationship

The **Demand** is rooted from participation in the governance process, determining upgrades to the protocol or allocation of resources, and **supply** is the result of the tokens eligible to participate in the governance process.

\$$ \text{Token Price \textbf{(SNR)\}} = \text{\underline{Fundamental Value\}} + \text{\textbf{Token Specific} Non-Fundamentals} + \text{\textbf{Market/Industry} Non-Fundamentals} \$$

Establishing fundamental value via governance incorporates seven elements of economic design:

1. **Scope of Decisions**
2. **Stakeholders**
3. **Policy Research & Development**
4. **Proposal Process**
5. **Information Distribution Systems**
6. **Decision Making Procedures**
7. **Implementation & Property Rights**

The **price** of a token is the exchange rate of the token for fiat. The **value** of the token can be modeled in terms of fundamental demand drivers and effective supply. The **demand** of the token is driven by the **property rights granted by the token**, and the effective **supply** is driven by the **number of tokens to which a specific set of rights are granted**.

---

## Deployment Timeline

### 1. Establish Community

Build a loyal community of token holders that represent our platforms stakeholder groups.

- All of our existing token holders (this includes our founder and early investors) will be staking their Sonr Token in order to be incentivized to hold tokens and participate in voting
- We will be actively networking with DeFi projects, blockchain service providers, and etc. in order to participate

### 2. Polling System

Our initial polling system will be token weighted and not enforceable, in order to build engagement.

- Users will submit qualitative updates to the system and have them up/downvoted
- The core team will then assess the submissions and select viable community suggestions into upgrades
  - Along side this, our team will locate and commission the development of a proposal for any technical upgrades without enforceable code proposals
  - We will be able to provide clarity, external security, and in applicable cases the economic evaluation from the impacts of proposals to the community.

### 3. On-Chain Execution

We will be introducing an on-chain, executable code-based governance system in stages with majority rule voting.

**Stage 1:** Grants Program

- Incorporate community voting for grant recipient selection
- Allow users to submit proposals for grant allocation

**Stage 2:** Subsidies and Rewards

- Proposals to update subsidies and rewards can be submitted by Core Sonr Team or third parties
- Core Team will be responsible for considering funding security, and economic audits for proposals

**Stage 3:** Technical Upgrades

- These include pricing, security, and bug fixes
- Proposals will be spearheaded by the Sonr Team and/or its delegates

### 4. Maintenance

When on-chain governance has come to fruition, the core Sonr team will be also maintaining the polling-based system with subsidized development to minimize barrier to entry. This will be incorporating governance within the Motor Nebula Widget itself.

---

## References

- [Decentralized Identifiers by W3C](https://www.w3.org/TR/did-core/)
- [Delegated Proof of Stake by BitShares](https://how.bitshares.works/en/master/technology/dpos.html)
