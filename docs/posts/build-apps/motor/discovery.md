---
title: Discovery
id: discovery
displayed_sidebar: modulesSidebar
---
# Discovery

The Sonr protocol uses various methods and fallbacks to ensure no friction discoverability when finding another user. Its node has an Exchange service that provides immediate discovery, validation, and verification access for every peer. The service operates in two separate modes: Local and Global, with the Local mode functioning similar to Airdrop and Global to Email.

# Local

Before the discovery process begins, nodes provide their Lat/Lng to join the local lobby - when location isn't provided, the node will operate on Global only mode. After receiving the values, an OLC (Open Location Code) is generated, which becomes the suffix for the new Local Topic. Peers that are validated then have access to Read/Write on the Local Room's DHT (Distributed Hash Table), where they would provide any Profile Info or Device Metadata.

Facilitating this is a combination of transports including mDNS, Bluetooth, Rendevouz Signalling, and the DHT.
After Peers are discovered, the cached copy of the DHT is streamed to any client-facing implementation bound to the node. The peers are then displayed on the UI.

# Global

The Global topic uses a lighter data structure for managing peers than rather the robust DHT used locally. To solve the issue of scale, we created Beam - a decentralized simple key/value store that has verifiable write access. By using Beam in combination with the node's HDNS Subdomain, we can successfully query the entire network in an instant, even when the node itself is offline.

# Offline

A common issue with P2P systems is handling the situation when the node is offline. With Sonr, we solve this issue by implementing the DIDCommMessaging Spec along with an IPFS mailbox where only users can access with their .snr/.
