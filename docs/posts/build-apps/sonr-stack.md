---
title: The Sonr Stack
id: sonr-stack
displayed_sidebar: modulesSidebar
---
# The Sonr Stack
The Sonr Node has two different modes: Motor and Highway. The Motor node operates on every client implementation of Sonr.  The Highway node facilitates the flow of Motor Nodes interfacing on the Sonr network. The binary sizes by Operating System are below:

![td](https://archbee-image-uploads.s3.amazonaws.com/YigsjtwFFq_eX7dhChoeN/UplhsgArEk5gSYM7YpuQx_bdfc32b-7.png)

## 1. Motor

When the node operates in Light Mode it becomes a Peer in the Sonr Network. It has methods and receives callbacks for Discovery and Transmission updates.

Sonr enabled Apps or Wallets will have Discovery, Identity, and Transmission capabilities enabled by default - along with a full Motor Node. Motor Nodes automatically communicate with Highway Nodes when needed.

## 2. Highway

When the node operates in Highway Mode it works as a custodian in the Network to weed out bad actors and helps relay Peers amongst themselves. Along with custodial duty, Highway nodes manage the registration of HDNS Subdomains and interaction with the Sonr Blockchain.


<!--
[ts]("https://www.figma.com/file/gL4iAj7V42JSAsTxUE6J4G/Highway-SDK-Topology?node-id=0%3A1") -->



### Developer SDK

The Sonr Highway SDK is a cross-stack framework for developers to easily build decentralized applications (”Apps” for shorthand) on the Sonr platform. Modeled after Amazon Web Services and Google Firebase, the goal is to enable developers to build the future of applications on Web3 with ease and elegance.

The Highway SDK consists of five unique modules:

*   Channels

*   Objects

*   Channels

*   Storage
