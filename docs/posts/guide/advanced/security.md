---
title: Security
id: security
displayed_sidebar: basicsSidebar
---
# Security
Both of Sonr's Highwawy and Motor Nodes require 3 separate pairs of keys to interact with the Network. An Account Key, Link Key, and Group Key.

### Master Key

This key is generated once a user claims a SName. This key provides a traditional seed phrase for recovery and creates a new DID Document for the User on the Sonr Blockchain. The `also_known_as` field will be pointing to the newly registered name.

### Device Key

The device key is an `ED25519` key stored locally on the device's Keychain that is set as the controller for the Users DIDDocument on the Sonr Blockchain. This key is used during all communication with the Sonr network and encrypts/decrypts data on muxed streams across the network.

### Authorize Key

The authorize key is a `multisignature` key that is compiled of all the Device Public Keys pointing to the user's ".snr" name. The public key of this pair will be the address pointing to the Account's mailbox. Sonr powered services leverage this key when verifying and authenticating users on their Applications.
