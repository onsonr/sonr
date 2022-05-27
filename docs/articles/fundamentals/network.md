---
id: network
title: Network
displayed_sidebar: basicsSidebar
---

# The Sonr Network
---
Sonr is fundamentally a decentralized peer-to-peer protocol built on a multi-asset, 
Proof-of-Stake (PoS) blockchain. The Sonr network is essentially a mesh of nodes that 
communicate over the Sonr protocol.

Built with [Cosmos SDK](https://docs.cosmos.network/), the Sonr blockchain can natively
interoperate with other blockchains through the 
[Inter-Blockchain Communication Protocol (IBC)](https://github.com/cosmos/ibc).

## Sonr Protocol

Built on a modular design, the Sonr protocol is in fact a combination of multiple
open sub-protocols that stack on top of each other, composing features that support the 
entire Sonr ecosystem. Each sub-protocol takes care of an individual logical part of 
that ecosystem.

### ADRs

Architectural Decision Records (ADRs) are documents used by the Sonr team as the actual
specs of the protocol, meaning they should always drive the implementation. Currently, 
the format of ADRs allows for collaboration through a revision system.

## Node Types
---

### Motor

Motor is Sonr's reference light client implementation, and will have all the properties of a
typical [Tendermint Light Client](https://docs.tendermint.com/master/tendermint-core/light-client.html),
and is the best option for quickly interacting with the Sonr network. Some of its capabilities are:

- Discovery
- 
- Transmission
- Registration
- Validation
- Highway API

For details, refer to [The Sonr Stack](/articles/sonr-stack#1-motor)

### Highway 



For details, refer to [The Sonr Stack](/articles/sonr-stack#2-highway)

