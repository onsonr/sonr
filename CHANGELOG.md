## v0.2.0 (2024-09-21)

### Feat

- add automated production release workflow
- **did**: remove unused proto files
- add enums.pulsar.go file for PermissionScope enum (#4)
- add initial DID implementation
- remove builder interface
- add basic UI for block explorer
- add Usage: pkl [OPTIONS] COMMAND [ARGS]...
- use SQLite embedded driver

### Fix

- Update proc_list_width in mprocs.yaml
- Add service to database when registering
- pin added did documents to local ipfs node
- remove extra spaces in typeUrl
- **release**: remove unnecessary quotes in tag pattern

### Refactor

- simplify verification method structure
- use staking keeper in DID keeper
- remove unused dependencies
- remove unused image building workflow

## v0.1.0 (2024-09-05)

### Feat

- add DID method for each coin
- Expand KeyType enum and update KeyInfo message in genesis.proto
- Add whitelisted key types to genesis params
- Add DID grants protobuf definition
- Add fields to KeyInfo struct to distinguish CBOR and standard blockchain key types
- Add new message types for AssetInfo, ChainInfo, Endpoint, ExplorerInfo, FeeInfo, and KeyInfo
- run sonr-node container in testnet network and make network external
- Add docker-compose.yaml file to start a Sonr testnet node

### Fix

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

- add  field to
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
