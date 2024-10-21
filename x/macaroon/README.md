# `x/macaroon`

The Macaroon module is responsible for providing decentralized access control and service authorization for the Sonr ecosystem. It implements macaroon-based authentication and authorization mechanisms.

## Concepts

Macaroons are a type of bearer credential that allow for decentralized delegation, attenuation, and third-party caveats. This module implements the core functionality for creating, validating, and managing macaroons within the Sonr ecosystem.

## State

The module maintains the following state:

### Grant

Represents a permission grant with the following fields:
- `id`: Unique identifier (auto-incremented)
- `controller`: Address of the controller
- `subject`: Subject of the grant
- `origin`: Origin of the grant
- `expiry_height`: Block height at which the grant expires

### Macaroon

Represents a macaroon token with the following fields:
- `id`: Unique identifier (auto-incremented)
- `controller`: Address of the controller
- `subject`: Subject of the macaroon
- `origin`: Origin of the macaroon
- `expiry_height`: Block height at which the macaroon expires
- `macaroon`: The actual macaroon token

## State Transitions

State transitions occur through the following messages:
- `MsgUpdateParams`: Updates the module parameters
- `MsgIssueMacaroon`: Issues a new macaroon

## Messages

### MsgUpdateParams

Updates the module parameters. Can only be executed by the governance account.

Fields:
- `authority`: Address of the governance account
- `params`: New parameter values

### MsgIssueMacaroon

Issues a new macaroon for a given controller and origin.

Fields:
- `controller`: Address of the controller
- `origin`: Origin of the request in wildcard form
- `permissions`: Map of permissions
- `token`: Macaroon token to authenticate the operation

## Queries

The module provides the following queries:

- `Params`: Retrieves the current module parameters
- `RefreshToken`: Refreshes a macaroon token (post-authentication)
- `ValidateToken`: Validates a macaroon token (pre-authentication)

## Parameters

The module has the following parameters:

- `methods`: Defines the available DID methods
  - `default`: Default method
  - `supported`: List of supported methods
- `scopes`: Defines the set of scopes
  - `base`: Base scope
  - `supported`: List of supported scopes
- `caveats`: Defines the available caveats
  - `supported_first_party`: List of supported first-party caveats
  - `supported_third_party`: List of supported third-party caveats
- `transactions`: Defines the allowlist and denylist for transactions
  - `allowlist`: List of allowed transactions
  - `denylist`: List of denied transactions

## Events

The module may emit events related to macaroon issuance, validation, and refreshing. (Specific event details to be implemented)

## Client

The module provides gRPC endpoints for all queries and message types defined in the protobuf files.

## Future Improvements

- Implement more advanced caveat types
- Add support for third-party caveats
- Enhance macaroon revocation mechanisms
- Implement additional security features and checks

## Tests

(To be implemented: Acceptance tests for the module's functionality)

## Appendix

For more information on macaroons and their implementation, refer to the original macaroon paper: "Macaroons: Cookies with Contextual Caveats for Decentralized Authorization in the Cloud" by Arnar Birgisson, et al.
