# `x/did`

The Decentralized Identity module is responsible for managing native Sonr Accounts, their derived wallets, and associated user identification information.

## State

The DID module maintains several key state structures:

### Controller State

The Controller state represents a Sonr DWN Vault. It includes:
- Unique identifier (number)
- DID
- Sonr address
- Ethereum address
- Bitcoin address
- Public key
- Keyshares pointer
- Claimed block
- Creation block

### Assertion State

The Assertion state includes:
- DID
- Controller
- Subject
- Public key
- Assertion type
- Accumulator (metadata)
- Creation block

### Authentication State

The Authentication state includes:
- DID
- Controller
- Subject
- Public key
- Credential ID
- Metadata
- Creation block

### Verification State

The Verification state includes:
- DID
- Controller
- DID method
- Issuer
- Subject
- Public key
- Verification type
- Metadata
- Creation block

## State Transitions

State transitions are triggered by the following messages:
- LinkAssertion
- LinkAuthentication
- UnlinkAssertion
- UnlinkAuthentication
- ExecuteTx
- UpdateParams

## Messages

The DID module defines the following messages:

1. MsgLinkAuthentication
2. MsgLinkAssertion
3. MsgExecuteTx
4. MsgUnlinkAssertion
5. MsgUnlinkAuthentication
6. MsgUpdateParams

Each message triggers specific state machine behaviors related to managing DIDs, authentications, assertions, and module parameters.

## Query

The DID module provides the following query endpoints:

1. Params: Query all parameters of the module
2. Resolve: Query the DID document by its ID
3. Sign: Sign a message with the DID document
4. Verify: Verify a message with the DID document

## Params

The module parameters include:
- Allowed public keys (map of KeyInfo)
- Conveyance preference
- Attestation formats

## Client

The module provides gRPC and REST endpoints for all defined messages and queries.

## Future Improvements

Potential future improvements could include:
1. Enhanced privacy features for DID operations
2. Integration with more blockchain networks
3. Support for additional key types and cryptographic algorithms
4. Improved revocation mechanisms for credentials and assertions

## Tests

Acceptance tests should cover all major functionality, including:
- Creating and managing DIDs
- Linking and unlinking assertions and authentications
- Executing transactions with DIDs
- Querying and resolving DIDs
- Parameter updates

## Appendix

### Account

An Account represents a user's identity within the Sonr ecosystem. It includes information such as the user's public key, associated wallets, and other identification details.

### Decentralized Identifier (DID)

A Decentralized Identifier (DID) is a unique identifier that is created, owned, and controlled by the user. It is used to establish a secure and verifiable digital identity.

### Verifiable Credential (VC)

A Verifiable Credential (VC) is a digital statement that can be cryptographically verified. It contains claims about a subject (e.g., a user) and is issued by a trusted authority.

### Key Types

The module supports various key types, including:
- Role
- Algorithm (e.g., ES256, EdDSA, ES256K)
- Encoding (e.g., hex, base64, multibase)
- Curve (e.g., P256, P384, P521, X25519, X448, Ed25519, Ed448, secp256k1)

### JSON Web Key (JWK)

The module supports JSON Web Keys (JWK) for representing cryptographic keys, including properties such as key type (kty), curve (crv), and coordinates (x, y) for EC and OKP keys, as well as modulus (n) and exponent (e) for RSA keys.
