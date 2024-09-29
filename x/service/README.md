# `x/service`

The Service module is responsible for managing the registration and authorization of services within the Sonr ecosystem. It leverages
the native NFT module associated with DID Methods to provide a secure and verifiable mechanism for registering and authorizing services.

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
his is a module base generated with [`spawn`](https://github.com/rollchains/spawn).
