# `x/vault`

The Vault module is responsible for the management of IPFS deployed Decentralized Web Nodes (DWNs) and their associated data.

## Concepts

## State

Specify and describe structures expected to marshalled into the store, and their keys

### Account State

The Account state includes the user's public key, associated wallets, and other identification details. It is stored using the user's DID as the key.

### Credential State

The Credential state includes the claims about a subject and is stored using the credential ID as the key.

## State Transitions

Standard state transition operations triggered by hooks, messages, etc.

## Messages

Specify message structure(s) and expected state machine behaviour(s).

## Begin Block

Specify any begin-block operations.

## End Block

Specify any end-block operations.

## Hooks

Describe available hooks to be called by/from this module.

## Events

List and describe event tags used.

## Client

List and describe CLI commands and gRPC and REST endpoints.

## Params

List all module parameters, their types (in JSON) and identitys.

## Future Improvements

Describe future improvements of this module.

## Tests

Acceptance tests.

## Appendix

Supplementary details referenced elsewhere within the spec.

| Concept                                     | Description                                                                                                                                                                           |
| ------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Decentralized Web Node (DWN)                | A decentralized, distributed, and secure network of nodes that store and share data. It is a decentralized alternative to traditional web hosting services.                           |
| Decentralized Identifier (DID)              | A unique identifier that is created, owned, and controlled by the user. It is used to establish a secure and verifiable digital identity.                                             |
| HTMX (Hypertext Markup Language eXtensions) | A set of extensions to HTML that allow for the creation of interactive web pages. It is used to enhance the user experience and provide additional functionality to web applications. |
| IPFS (InterPlanetary File System)           | A decentralized, peer-to-peer network for storing and sharing data. It is a distributed file system that allows for the creation and sharing of content across a network of nodes.    |
| WebAuthn (Web Authentication)               | A set of APIs that allow websites to request user authentication using biometric or non-biometric factors.                                                                            |
| WebAssembly (Web Assembly)                  | A binary instruction format for a stack-based virtual machine.                                                                                                                        |
| Verifiable Credential (VC)                  | A digital statement that can be cryptographically verified.                                                                                                                           |
