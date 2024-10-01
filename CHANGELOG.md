## v0.10.0 (2024-10-01)

### Feat

- allow manual release triggers

## v0.9.0 (2024-10-01)

### Feat

- add Input and RegistrationForm models
- add new utility classes
- add login and registration pages
- add tailwindcss utilities

### Fix

- **cta**: Fix typo in CTA title
- change bento section title to reflect security focus

### Refactor

- move home page sections to home package

## v0.8.0 (2024-10-01)

### Feat

- add support for ARM64 architecture

## v0.7.0 (2024-10-01)

### Feat

- add DWN resolver field
- add stats section to homepage
- implement hero section using Pkl
- add PKL schema for message formats

### Fix

- adjust hero image dimensions
- **Input**: Change  type from  to
- update hero image height in config.pkl

### Refactor

- rename motrd to motr
- update hero image dimensions
- move nebula configuration to static file
- rename buf-publish.yml to publish-assets.yml
- remove unused  field from

## v0.6.0 (2024-09-30)

### Feat

- add Homebrew tap for sonr

## v0.5.0 (2024-09-30)

### Feat

- update release workflow to use latest tag

## v0.4.0 (2024-09-30)

### Feat

- **dwn**: add wasm build for dwn
- add macaroon and oracle genesis states
- add scheduled binary release workflow
- introduce process-compose for process management
- add counter animation to hero section
- add registration page

### Fix

- Enable scheduled release workflow

### Refactor

- remove old changelog entries
- remove unnecessary checkout in scheduled-release workflow
- rename build ID to sonr
- remove unnecessary release existence check
- move dwn wasm build to pkg directory

## v0.3.1 (2024-09-29)

### Refactor

- move nebula/pages to pkg/nebula/pages

## v0.3.0 (2024-09-29)

### Feat

- add buf.lock for proto definitions

### Fix

- remove unused linting rules
- update proto breaking check target to master branch

### Refactor

- remove unused lock files and configurations

## v0.2.0 (2024-09-29)

### Feat

- disable goreleaser workflow
- update workflows to include master branch
- remove global style declaration
- **oracle**: add oracle module
- optimize IPFS configuration for better performance
- add local IPFS bootstrap script and refactor devbox config
- add AllocateVault HTTP endpoint
- add WebAuthn credential management functionality
- remove unused coins interface
- remove global integrity proof from genesis state
- add vault module
- enable buf.build publishing on master and develop branches
- add Gitflow workflow for syncing branches
- add automated production release workflow
- **ui**: implement profile page
- add automated production release workflow
- **did**: remove unused proto files
- add enums.pulsar.go file for PermissionScope enum (#4)
- add initial DID implementation
- remove builder interface
- add basic UI for block explorer
- add Usage: pkl [OPTIONS] COMMAND [ARGS]...
- use SQLite embedded driver
- add DID method for each coin
- Expand KeyType enum and update KeyInfo message in genesis.proto
- Add whitelisted key types to genesis params
- Add DID grants protobuf definition
- Add fields to KeyInfo struct to distinguish CBOR and standard blockchain key types
- Add new message types for AssetInfo, ChainInfo, Endpoint, ExplorerInfo, FeeInfo, and KeyInfo
- run sonr-node container in testnet network and make network external
- Add docker-compose.yaml file to start a Sonr testnet node
- configure Sonr testnet environment
- Update Dockerfile to start and run a testnet
- add Equal methods for AssetInfo and ChainInfo types
- Add ProveWitness and SyncVault RPCs
- Add MsgRegisterService to handle service registration
- Add MsgRegisterService to handle service registration
- add enums.pulsar.go file for PermissionScope enum

### Fix

- ensure go version is up-to-date
- use GITHUB_TOKEN for version bump workflow
- update account table interface to use address, chain and network
- **ci**: update docker vm release workflow with new token
- use mnemonic phrases for test account keys
- reduce motr proxy shutdown timeout
- **nebula**: use bunx for tailwindcss build
- **proto**: update protobuf message index numbers
- **ante**: reduce POA rate floor and ceiling
- Update proc_list_width in mprocs.yaml
- Add service to database when registering
- pin added did documents to local ipfs node
- remove extra spaces in typeUrl
- **release**: remove unnecessary quotes in tag pattern
- remove unused imports and simplify KeyInfo message
- bind node ports to localhost
- Update docker-compose network name to dokploy-network
- Update network name to dokploy
- remove unused port mapping
- Update docker-compose.yaml to use correct volume path
- update docker-compose volume name
- Update docker-compose.yaml to use shell directly for sonrd command
- replace "sh" with "/bin/sh" in docker-compose.yaml command
- Update runner image dependencies for debian-11
- **deps**: update golang image to 1.21
- **chains**: update nomic chain build target
- Remove unused `Meta` message from `genesis.proto`
- Add ProveWitness and SyncVault RPCs

### Refactor

- adjust source directory for config files (#1102)
- Use actions/checkout@v4
- remove unused master branch from CI workflow
- rename github token secret
- remove unnecessary x-cloak styles
- optimize oracle genesis proto
- remove unused code related to whitelisted assets
- update buf publish source directory
- adjust devbox configuration to reflect nebula changes
- rename msg_server.go to rpc.go
- remove devbox integration
- move dwn package to app/config
- move configuration files to app directory
- extract root command creation to separate file
- move ipfs setup to function
- remove unnecessary proxy config
- rename script to
- move DWN proxy server logic to separate file
- use htmx instead of dwn for vault client
- remove unused environment variables
- simplify verification method structure
- use staking keeper in DID keeper
- remove unused dependencies
- remove unused image building workflow
- add field to
- Update KeyKind Enum to have proper naming conventions
- Update `DIDNamespace` to have proper naming convention
- expose ports directly in docker-compose
- remove unused port mappings
- streamline script execution
- use CMD instead of ENTRYPOINT in Dockerfile
- **deps**: Upgrade Debian base image to 11
- Simplify the types and properties to keep a consistent structure for the blockchain
- remove PERMISSION_SCOPE_IDENTIFIERS_ENS enum value
