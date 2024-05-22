# `x/did`

The Decentralized Identity module is responsible for managing native Sonr Accounts, their derived wallets, and associated user identification information.

## Concepts

### Account

An Account represents a user's identity within the Sonr ecosystem. It includes information such as the user's public key, associated wallets, and other identification details.

### Decentralized Identifier (DID)

A Decentralized Identifier (DID) is a unique identifier that is created, owned, and controlled by the user. It is used to establish a secure and verifiable digital identity.

### Verifiable Credential (VC)

A Verifiable Credential (VC) is a digital statement that can be cryptographically verified. It contains claims about a subject (e.g., a user) and is issued by a trusted authority.

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
