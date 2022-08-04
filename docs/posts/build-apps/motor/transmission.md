---
title: Transmission
id: transmission
displayed_sidebar: modulesSidebar
---
# Transmission

We believe that the most neglected aspect of the current internet is File Transfer. Simply put, there's been little change in how file transfer has worked since the invention of Mediafire. Right now sharing files is equivalent to mailing a letter to your next-door neighbor instead of just handing it to them yourself.

On Sonr your data is transmitted with encrypted p2p channels without the use of a middle-man server. Assuming that you are sending a file, it would be directly sent to the peer itself as opposed to being hosted on a monolithic server. In the event where a peer is Offline, IPFS Storage is used and the bucket itself is signed with the peer's public key.

## Efficiency

On all write streams, Sonr uses the FastCDC chunking algorithm which makes direct transmission 30% faster.

## Identifiers

Favorited/Pinned data will be assigned a Decentralized Identifier which would then be resolvable by every Peer on the Network.
