---
id: how-it-works
title: How It Works
displayed_sidebar: basicsSidebar
---
# How it works
Sonr decentralizes application databases, making data universally composable and reusable across applications. The network consists of three core parts: &#x20;

*   Highly-scalable, decentralized infrastructure for data availability and consensus,&#x20;

*   Marketplace of community-created data models

*   Suite of standard APIs for storing, updating, and retrieving data from those models.

# Fundamental Building Blocks

1.  Verifiable User Identity

2.  Realtime Peer-to-Peer Network

3.  Permission-less Application Hub

4.  IBC Enabled Blockchain

5.  Object Storage (Data)

6.  Blob Storage (Files)

7.  (Location based discovery) (Device Discovery)

# Verifiable User Identity

Our base DID (or decentralized identifier) follows a syntactic structure of the **root** (your DID), followed by a **method** (or in our case, in every case with this SDK, Sonr), then followed by a **public key**

![DID and channel schema](https://archbee-image-uploads.s3.amazonaws.com/YigsjtwFFq_eX7dhChoeN/ze9buUbapxPP7S5ROVXn__6e60b2d-screenshot2022-03-10at25108pm.png)


# Realtime Peer to Peer Network

Our network is build upon `LibP2P` a peer to peer communication protocol. This allows Highway Nodes and Motor nodes to transmit data, share files, and locate one another. Each motor node in our network is assigned a `DHT` address which allows for unique addressing of each motor in our application network.


# IBC (Inter Blockhain Communication)

IBC can be implemented by any consensus algorithm that supports cheaply verifiable finality with any state machine that supports vector commitments. IBC defines a set of low-level primitives for authentication, transport, and ordering, and a set of application-level standards for asset & data semantics. Ledgers which support compatible standards can be connected together without any special permissions.
