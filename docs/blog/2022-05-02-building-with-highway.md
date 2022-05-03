---
title: Building with Highway
slug: building-with-highway
date: 2022-04-15T21:43:34.000Z
authors: [mcjcloud]
tags: [hello, docusaurus]
---



---

# Overview

The highway is a single binary which allows for interfacing with the Sonr Blockchain (see **'Using the CLI'** for information on commands). The highway is also equipped with a REST server.  The following is a diagram outlining the topology of highway and available features.

We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that’s composable and interoperable with all existing protocols. For this, we built our Networking layer in Libp2p and our Layer 1 Blockchain with Starport. Our network comprises of two separate nodes: Highway and Motor, which each have a specific use case on the network. In order to maximize the onboarding experience, we developed our own Wallet which has value right out of the gate!



[t]("https://www.figma.com/file/kZVXK3yJOxmukNdckjh2RT/Highway-SDK?node-id=2%3A12")



# Using the CLI

our Highway-sdk comes with a set of CLI commands&#x20;



```none
serve - Serves our GRPC and HTTP servers on the specified ports in our enviorment files
```



# Using REST

The highway is capable of running an http server (REST) with 'serve' ports can be specified

*   Register

*   Authentication

*   Objects&#x20;

    *   Create

    *   Update

    *   Deactivate

*   Buckets

    *   Create

    *   Update

    *   Deactivate

*   Channels

    *   Create

    *   Hide

*   registry

    *   query

    *   exists


