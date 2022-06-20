---
title: Interface
id: interface
displayed_sidebar: modulesSidebar
---

# Motor Interface
## Overview

### Interaction Methods
These are methods that are handled only between the user and their data, or directly with another peer.
```Text

 ClearMessage()
 DownloadVault()
 OpenMailbox()
 RequestPeer()
 RespondPeer()
 SendMessage()
 UploadVault()
```


### Highway Methods
These methods interact with a Sonr Highway node which communicates with the underlying blockchain
```Text
 AccessApp()
 InteractObject()
 ListenChannel()
 LinkDevice()
 LinkPeer()
 ReadBucket()
 ReadObject()
 RequestBucket()
 RespondBucket()
```


### Node Methods
These are developer facing methods that manage the state of the underlying motor node.
```Text

 Start()
 Stop()
 Pause()
 Resume()
 Status()
```

### AccessApp
User authenticates a Registered Application on Sonr with their DID Based Multisignature key for all their devices. Creates a new Bucket inside the User bucket for the newly provisioned Application.

### InteractObject
Users map the new data for a specific type definition presented in the UI, and push the updated data to the corresponding application in their Bucket. This utilizes the JWE process in order to encrypt data from the User end.


### ListenChannel
User specifies which application data stream to begin reading for data. The returned channel is a listenable stream or callback depending on Device architecture.


### LinkDevice
This method allows motor based applications to link an addition WebAuthN credential to their top-level


### ReadBucket
This method be utilized with pulling a Application Specific buckets type definitions, functions, etc. In order to render the payloads onto the frontend UI.


### ReadObject
This method begins the request process for reading the individual object values of another User's app bucket. Provisioned users will automatically get access (if they called RequestBucket already) and unprovisioned users will send the RequestBucket to the corresponding peer.

### RequestBucket
This method is utilized for accessing another users application specific data holistically. This would be utilized for full access to a peers App Data config


### RespondBucket
This method is utilized for responding to a request from another user from the mailbox folder in their User specific bucket.


### ClearMessage()
Removes a message from the mailbox with provided ID.


### DownloadVault
Retreives a file or buffer from the user data vault from in their Bucket - with either a specified CID or REGEX Match string for querying files.


### OpenMailbox
Retreives the list of messages stored in the user Mailbox, with the id and their corresponding buffer body.


### RequestPeer
Adds a connected peer into the root level DID as a InvocationMethod in order to have ongoing connection.


### RespondPeer
Approves/Declines a request from another peer to connect as an InvocationMethod on the top level DID Document for a User.


### SendMessage
Send a message to another peer's mailbox utilizing any arbitrary buffer of data, signed with the users DIDDocument. The other user then fetches the public key from the sending user's DIDDocument on the sonr blockchain in order to verify the authenticity of the sender.

### UploadVault
Uploads data to the User specific bucket in their Vault directory.
