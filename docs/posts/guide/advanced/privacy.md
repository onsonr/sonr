---
title: Privacy
slug: zlY7-privacy
createdAt: 2022-04-15T20:04:42.000Z
updatedAt: 2022-04-23T19:51:52.000Z
---
# Privacy
## Overview

Our protocols are built on top of block chain technology, letting us take advantage of it's distributive anonymous nature. All data stored is denormalized. Writes on our chain are only related by did, and publicKey. With each message sent on our network handshaked, with TCP/TLS.&#x20;

### Encrypted Messaging

All data stored, communicated on our network is end to end encrypted. Using `EDC25519` elliptic curve standard, each operation requires a key exchange inorder to operate on user data. This makes it so your data can only opened by you or users you trust.

### Decentralized Data Storage

We use `IPFS` to store `Blobs`, and `Objects`. Each item stored is sharded, meaning your data is stored in fragments across our network. Each item is encrypted using the `EDC25519` standard when stored on the nodes file system. Each `Highway` node is associated with an IPFS instance meaning your data can be sharded across our network with high Encryption standards.

#
