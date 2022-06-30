---
title: Overview
id: overview
displayed_sidebar: modulesSidebar
---

# `x/registry`
The Sonr registry module is used to store the records of user accounts and applications. Each record contains a DIDDocument and additional WebAuthn credential information.

## Overview

The record type utilized in the **Registry module** is the `WhoIs` type. This type provides both an interface to utilize WebAuthn credentials and a means to access the record's DIDDocument.

## Usage

> Blockchain Methods supplied by Registry Module. Full implementation is still a work in progress.

### `RegisterName()` - Register's a new '.snr' domain name for an account

```Text
- (`string`) NameToRegister     : The name to register
- (`string`) Creator            : The Account Address signing this message
- (`Credential`) Credential     : Webauthn credential to use for registration
- (`map`) Metadata              : Metadata to attach to the `WhoIs` record
```

### `RegisterApplication()` - Register's a new Application for the Sonr Network

```Text
- (`string`) Creator                : The Account Address signing this message
- (`Credential`) Credential         : Webauthn credential to use for registration
- (`string`) ApplicationName        : The Name of the Application being registered
- (`string`) ApplicationDescription : Short about description of the App
- (`string`) ApplicationURL         : Website/Homepage of the App
- (`string`) ApplicationCategory    : Category of the Application Type
```

### `AccessName()` - Accesses a particular name essentially a "Login" function

```Text
- (`string`) Creator            : The Account Address signing this message
- (`Credential`) Credential     : Webauthn credential to use for registration
- (`string`) Name               : The name to authenticate and retreive data
```

### `AccessApplication()` - Accesses a particular application essentially a "Register" function

```Text
- (`string`) AppName                : The Name of the Application being accessed
- (`string`) Creator                : The Account Address signing this message
- (`Credential`) Credential         : Webauthn Credential of the authenticated user
```

## Record Type: `WhoIs`



```
- ('string') Name           : Name is the registered name of the User or Application
- ('string') DID            : DID is the DID of the account
- ('bytes') Document       : Document is the DID Document of the registered name and account encoded as JSON
- ('string') Creator        : Creator is the Account Address of the creator of the DID Document
- ('Credential') Credentials    : Credentials are the biometric info of the registered name and account encoded with public key
- ('Type') Type           : Type is the type of the registered name
- ('string') MetaData       : Additional Metadata for associated WhoIs
- ('string') Timestamp      : Timestamp is the time of the last update of the DID Document
- ('string') IsActive       : IsActive is the status of the DID Document
- ('Enum') Type             : Type is the type of the registered name

```



## Status Codes



```
200 - SUCCESS
300 - MULTIPLE CHOICE
304 - NOT MODIFIED
400 - BAD REQUEST
401 - NOT AUTHORIZED
```


