---
title: HowIs
slug: 3M4V-howis
createdAt: 2022-04-26T14:53:09.000Z
updatedAt: 2022-04-28T19:49:44.000Z
---

# HowIs

## Overview
The `HowIs`Object function is included within `Channel` types. The `HowIs` Object illustrates how a channel is associated with a `registry` entry on the Sonr Network. A `registry` entry on the Sonr Network can either be for an `application` or a `name`

## Usage 
### Create HowIs

The following is an example of a request to make a new HowIs

```
MsgCreateHowIs {
  string creator;
  string did;
  ChannelDoc channel;
}
```

### Response from CreateHowIs

The following is an example of a response after creating a new HowIs

```
MsgCreateHowIsResponse {
    int32 code      // Code of the response
    string message  // Message of the response
    HowIs how_is    // HowIs of the Channel
}
```

### Update HowIs

The following is an example of a request to update a HowIs

```
MsgUpdateHowIs {
  string creator;
  string did;
  ChannelDoc channel;
}
```

### Response from UpdateHowIs

The following is an example of a response after Updating a HowIs

```
MsgUpdateHowIsResponse {
    int32 code     // Code of the response
    string message // Message of the response
    HowIs how_is   // HowIs of the Channel
}
```

### Deleting HowIs

The following is an example of a request to Delete a HowIs

```
MsgDeleteHowIs {
  string creator;
  string did;
}
```

### Response from DeleteHowIs

The following is an example of a response after Deleting a How Is

```
MsgDeleteHowIsResponse {
    int32 code     // Code of the response
    string message // Message of the response
}
```


































