---
title: Channels Overview
id: overview
displayed_sidebar: modulesSidebar
---

# `x/channels`

The Sonr channel module is used to store the records of the active pubsub topics associated with Applications powered by the Sonr Network. Each record contains a `ChannelDoc` which describes the Topic configuration and status of the channel. Each channel is required to have a set RegisteredType to pass as a payload with ChannelMessages.

## Overview

The record type utilized in the **Channel module** is the `HowIs` type. This type provides both an interface to utilize VerifiableCredentials and modify the ChannelDoc type definition

### Objects Relation

While channels determine and facilitate the actions passed through an application through realtime data streams, validators - or in our case, **Objects**, are essential to validating that data stream. Objects also make decisions as to which buckets user created data will be stored in.

### Examples

*   Realtime document editing -> imagine a decentralized version of Notion!

*   Shared device positioning and tracking -> think location services/device mapping.

*   Group Chat messaging -> p2p messaging, airdrops, all decentralized.

*   Secure direct transactions of data -> decentralized Dropbox or WeTransfer.

## Usage

> Blockchain Methods supplied by Channel Module. Full implementation is still a work in progress.

### `CreateChannel()` - Records a new Channel configuration for a specified application on Sonr.

```Text
- (`string`) Creator                : The Account Address signing this message
- (`Session`) Session               : The Session for the authenticated user
- (`string`) Label                  : Name of the channel defined by developer
- (`string`) Description            : Description of the channel defined by developer
- (`ObjectDoc`) ObjectToRegister    : The registered verified type to be sent in channel messages
```

## Status Codes

```
200 - SUCCESS
300 - MULTIPLE CHOICE
304 - NOT MODIFIED
400 - BAD REQUEST
401 - NOT AUTHORIZED

```
