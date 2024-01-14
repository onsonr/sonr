# `x/identity`

The Identity module is responsible for managing native Sonr Accounts, their derived wallets, and associated user identification information.

## Concepts

### Account

### Decentralized Identifier (DID)

### Verifiable Credential (VC)

## State

Specify and describe structures expected to marshalled into the store, and their keys

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

### Tech Stack

sonrhq/identity is built on the following main stack:

- <img width='25' height='25' src='https://img.stackshare.io/service/1005/O6AczwfV_400x400.png' alt='Golang'/> [Golang](http://golang.org/) – Languages
- <img width='25' height='25' src='https://img.stackshare.io/service/2501/default_3cf1b307194b26782be5cb209d30360580ae5b3c.png' alt='Prometheus'/> [Prometheus](http://prometheus.io/) – Monitoring Tools
- <img width='25' height='25' src='https://img.stackshare.io/service/4393/ma2jqJKH_400x400.png' alt='Protobuf'/> [Protobuf](https://developers.google.com/protocol-buffers/) – Serialization Frameworks
- <img width='25' height='25' src='https://img.stackshare.io/service/4631/default_c2062d40130562bdc836c13dbca02d318205a962.png' alt='Shell'/> [Shell](https://en.wikipedia.org/wiki/Shell_script) – Shells
- <img width='25' height='25' src='https://img.stackshare.io/service/4670/default_d811b0ac72205af84aca21f967594338580be913.png' alt='gRPC'/> [gRPC](https://grpc.io/) – Remote Procedure Call (RPC)
- <img width='25' height='25' src='https://img.stackshare.io/service/8695/stretchr.png' alt='Testify'/> [Testify](https://github.com/stretchr/testify) – Go Testing
- <img width='25' height='25' src='https://img.stackshare.io/service/11563/actions.png' alt='GitHub Actions'/> [GitHub Actions](https://github.com/features/actions) – Continuous Integration
