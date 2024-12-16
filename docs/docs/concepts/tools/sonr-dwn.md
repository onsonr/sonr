# `x/dwn`

The DWN module is responsible for the management of IPFS deployed Decentralized Web Nodes (DWNs) and their associated data.

## Concepts

The DWN module introduces several key concepts:

1. Decentralized Web Node (DWN): A distributed network for storing and sharing data.
2. Schema: A structure defining the format of various data types in the dwn.
3. IPFS Integration: The module can interact with IPFS for decentralized data storage.

## State

The DWN module maintains the following state:

### DWN State

The DWN state is stored using the following structure:

```protobuf
message DWN {
  uint64 id = 1;
  string alias = 2;
  string cid = 3;
  string resolver = 4;
}
```

This state is indexed by ID, alias, and CID for efficient querying.

### Params State

The module parameters are stored in the following structure:

```protobuf
message Params {
  bool ipfs_active = 1;
  bool local_registration_enabled = 2;
  Schema schema = 4;
}
```

### Schema State

The Schema state defines the structure for various data types:

```protobuf
message Schema {
  int32 version = 1;
  string account = 2;
  string asset = 3;
  string chain = 4;
  string credential = 5;
  string did = 6;
  string jwk = 7;
  string grant = 8;
  string keyshare = 9;
  string profile = 10;
}
```

## State Transitions

State transitions in the DWN module are primarily triggered by:

1. Updating module parameters
2. Allocating new dwns
3. Syncing DID documents

## Messages

The DWN module defines the following message:

1. `MsgUpdateParams`: Used to update the module parameters.

```protobuf
message MsgUpdateParams {
  string authority = 1;
  Params params = 2;
}
```

## Begin Block

No specific begin-block operations are defined for this module.

## End Block

No specific end-block operations are defined for this module.

## Hooks

The DWN module does not define any hooks.

## Events

The DWN module does not explicitly define any events. However, standard Cosmos SDK events may be emitted during state transitions.

## Client

The DWN module provides the following gRPC query endpoints:

1. `Params`: Queries all parameters of the module.
2. `Schema`: Queries the DID document schema.
3. `Allocate`: Initializes a Target DWN available for claims.
4. `Sync`: Queries the DID document by its ID and returns required information.

## Params

The module parameters include:

- `ipfs_active` (bool): Indicates if IPFS integration is active.
- `local_registration_enabled` (bool): Indicates if local registration is enabled.
- `schema` (Schema): Defines the structure for various data types in the dwn.

## Future Improvements

Potential future improvements could include:

1. Enhanced IPFS integration features.
2. Additional authentication mechanisms beyond WebAuthn.
3. Improved DID document management and querying capabilities.

## Tests

Acceptance tests should cover:

1. Parameter updates
2. DWN state management
3. Schema queries
4. DWN allocation process
5. DID document syncing

## Appendix

| Concept                                     | Description                                                                                                                                                                           |
| ------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Decentralized Web Node (DWN)                | A decentralized, distributed, and secure network of nodes that store and share data. It is a decentralized alternative to traditional web hosting services.                           |
| Decentralized Identifier (DID)              | A unique identifier that is created, owned, and controlled by the user. It is used to establish a secure and verifiable digital identity.                                             |
| HTMX (Hypertext Markup Language eXtensions) | A set of extensions to HTML that allow for the creation of interactive web pages. It is used to enhance the user experience and provide additional functionality to web applications. |
| IPFS (InterPlanetary File System)           | A decentralized, peer-to-peer network for storing and sharing data. It is a distributed file system that allows for the creation and sharing of content across a network of nodes.    |
| WebAuthn (Web Authentication)               | A set of APIs that allow websites to request user authentication using biometric or non-biometric factors.                                                                            |
| WebAssembly (Web Assembly)                  | A binary instruction format for a stack-based virtual machine.                                                                                                                        |
| Verifiable Credential (VC)                  | A digital statement that can be cryptographically verified.                                                                                                                           |
