## 0.1.0 (2024-09-05)

### Feat

- add SQLite database support
- Add targets for templ and vault in Makefile and use only make in devbox.json
- Add models.go file with database table structs
- Convert constant SQL queries to functions in queries.go and update db.go to use prepared statements
- Simplify db.go implementation
- Update the db implementation to use the provided go library
- Add DBConfig and DBOption types
- Add DIDNamespace and PermissionScope enums
- Add database enum types
- Update `createPermissionsTable` to match Permissions struct
- Add createKeysharesTable to internal/db/db.go
- Add constant SQL queries to queries.go and use prepared statements in db.go
- Update createProfilesTable and add createPropertiesTable
- Update the `createCredentialsTable` method to match the proper Credential struct
- Add keyshares table
- Implement database layer for Vault node
- introduce database layer
- Add method to initialize SQLite database
- add persistent SQLite database support in WASM
- **orm**: remove unused ORM models
- implement API endpoints for profile management
- Merge zkgate.go and zkprop.go logic
- Uncomment and modify zkgate code to work with Property struct
- Add zkgate.go file
- add WASM build step to devbox.json
- add KeyCurve and KeyType to KeyInfo in genesis
- Update the `CreateWitness` and `CreateAccumulator` and `VerifyWitness` and `UpdateAccumulator` to Use the new `Accumulator` and `Witness` types. Then Clean up the code in the file and refactor the marshalling methods
- add basic vault command operations
- add initial wasm entrypoint
- Implement IPFS file, location, and filesystem abstractions
- add IPFS file system abstraction
- Add AddFile and AddFolder methods
- Update GetCID and GetIPNS functions to read data from IPFS node
- Add local filesystem check for IPFS and IPNS
- Improve IPFS client initialization and mount checking
- Update `EncodePublicKey` to be the inverse of `DecodePublicKey`
- add DID model definitions
- add DID method for each coin
- Expand KeyType enum and update KeyInfo message in genesis.proto
- Add whitelisted key types to genesis params
- Add DID grants protobuf definition
- Add fields to KeyInfo struct to distinguish CBOR and standard blockchain key types
- Add new message types for AssetInfo, ChainInfo, Endpoint, ExplorerInfo, FeeInfo, and KeyInfo
- run sonr-node container in testnet network and make network external
- Add docker-compose.yaml file to start a Sonr testnet node

### Fix

- update Makefile to use sonrd instead of wasmd
- Remove unused statement map and prepare statements
- Refactor database connection and statement handling
- update db implementation to use go-sqlite3 v0.18.2
- Reorder the SQL statements in the tables.go file
- update go-sqlite3 dependency to version 1.14.23
- Update module names in protobuf files
- Ensure IPFS client is initialized before pinning CID
- Use Unixfs().Get() instead of Cat() for IPFS and IPNS content retrieval
- Initialize IPFS client and check for mounted directories
- update default assets with correct asset types
- Fix EncodePublicKey method in KeyInfo struct
- remove unused imports and simplify KeyInfo message
- bind node ports to localhost
- Update docker-compose network name to dokploy-network
- Update network name to dokploy
- remove unused port mapping
- Update docker-compose.yaml to use correct volume path
- update docker-compose volume name
- Update docker-compose.yaml to use shell directly for sonrd command
- replace "sh" with "/bin/sh" in docker-compose.yaml command

### Refactor

- remove unused template file
- Replace PrimaryKey with Property struct in zkprop.go
- remove unused FileSystem interface
- remove unused functions and types
- update AssetInfo protobuf definition
- add field to
- Update KeyKind Enum to have proper naming conventions
- Update `DIDNamespace` to have proper naming convention
- expose ports directly in docker-compose
- remove unused port mappings
- streamline script execution
- use CMD instead of ENTRYPOINT in Dockerfile

## v0.0.1 (2024-08-28)

### Feat

- configure Sonr testnet environment
- Update Dockerfile to start and run a testnet
- add Equal methods for AssetInfo and ChainInfo types
- Add ProveWitness and SyncVault RPCs
- Add MsgRegisterService to handle service registration
- Add MsgRegisterService to handle service registration
- add enums.pulsar.go file for PermissionScope enum

### Fix

- Update runner image dependencies for debian-11
- **deps**: update golang image to 1.21
- **chains**: update nomic chain build target
- Remove unused `Meta` message from `genesis.proto`
- Add ProveWitness and SyncVault RPCs

### Refactor

- **deps**: Upgrade Debian base image to 11
- Simplify the types and properties to keep a consistent structure for the blockchain
- remove PERMISSION_SCOPE_IDENTIFIERS_ENS enum value
