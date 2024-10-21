# `x/service`

The Service module is responsible for managing the registration and authorization of services within the Sonr ecosystem. It provides a secure and verifiable mechanism for registering and authorizing services using Decentralized Identifiers (DIDs).

## Concepts

- **Service**: A decentralized service on the Sonr Blockchain with properties such as ID, authority, origin, name, description, category, tags, and expiry height.
- **Profile**: Represents a DID alias with properties like ID, subject, origin, and controller.
- **Metadata**: Contains information about a service, including name, description, category, icon, and tags.

## State

The module uses the following state structures:

### Metadata

Stores information about services:
- Primary key: `id` (auto-increment)
- Unique index: `origin`
- Fields: id, origin, name, description, category, icon (URI), tags

### Profile

Stores DID alias information:
- Primary key: `id`
- Unique index: `subject,origin`
- Fields: id, subject, origin, controller

## Messages

### MsgUpdateParams

Updates the module parameters. Can only be executed by the governance account.

### MsgRegisterService

Registers a new service on the blockchain. Requires a valid TXT record in DNS for the origin.

## Params

The module has the following parameters:
- `categories`: List of allowed service categories
- `types`: List of allowed service types

## Query

The module provides the following query:

### Params

Retrieves all parameters of the module.

## Client

### gRPC

The module provides a gRPC Query service with the following RPC:
- `Params`: Get all parameters of the module

### CLI

(TODO: Add CLI commands for interacting with the module)

## Events

(TODO: List and describe event tags used by the module)

## Future Improvements

- Implement service discovery mechanisms
- Add support for service reputation and rating systems
- Enhance service metadata with more detailed information
- Implement service update and deactivation functionality

## Tests

(TODO: Add acceptance tests for the module)

## Appendix

This module is part of the Sonr blockchain project and interacts with other modules such as DID and NFT modules to provide a comprehensive decentralized service ecosystem.
