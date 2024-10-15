## v0.6.0 (2024-10-15)

### Feat

- add User-Agent and Platform to session
- introduce AuthState enum for authentication state

## v0.5.12 (2024-10-11)

## v0.5.11 (2024-10-10)

### Feat

- nebula assets served from CDN
- use CDN for nebula frontend assets
- add static hero section content to homepage
- add wrangler scripts for development, build, and deployment
- remove build configuration
- move gateway web code to dedicated directory
- add PubKey fast reflection
- **macaroon**: add transaction allowlist/denylist caveats
- add PR labeler
- **devbox**: remove hway start command
- add GitHub Actions workflow for running tests
- add workflow for deploying Hway to Cloudflare Workers
- Publish configs to R2
- integrate nebula UI with worker-assets-gen
- extract reusable layout components
- Implement service worker for IPFS vault
- implement CDN support for assets
- add payment method support
- add support for public key management
- add ModalForm component
- add LoginStart and RegisterStart routes
- implement authentication views
- add json tags to config structs
- implement templ forms for consent privacy, credential assert, credential register, and profile details
- **vault**: introduce assembly of the initial vault
- add client logos to homepage
- add tailwind utility classes
- implement new profile card component

### Fix

- Correct source directory for asset publishing
- install dependencies before nebula build
- update Schema service to use new API endpoint
- fix broken logo image path

### Refactor

- remove unnecessary branch configuration from scheduled release workflow
- update dwn configuration generation import path
- use nebula/routes instead of nebula/global
- move index template to routes package
- remove cdn package and move assets to global styles
- move nebula assets to hway build directory
- remove docker build and deployment
- rename internal/session package to internal/ctx
- remove unused fields from
- rename PR_TEMPLATE to PULL_REQUEST_TEMPLATE
- remove devbox.json init hook
- rename sonrd dockerfile to Dockerfile
- remove unused dependency
- rename 'global/cdn' to 'assets'
- move CDN assets to separate folder
- move Pkl module definitions to dedicated package
- move CDN assets to js/ folder
- remove unused component templates
- move ui components to global
- move view handlers to router package

## v0.5.10 (2024-10-07)

### Feat

- **blocks**: remove button component

## v0.5.9 (2024-10-06)

### Feat

- add Motr support
- update UIUX PKL to utilize optional fields

### Fix

- Update source directory for asset publishing

## v0.5.8 (2024-10-04)

### Refactor

- Remove unused logs configuration

## v0.5.7 (2024-10-04)

### Feat

- **devbox**: use process-compose for testnet services
- remove motr.mjs dependency
- add markdown rendering to issue templates
- update issue templates for better clarity
- add issue templates for tracking and task issues
- add issue templates for bug report and tracking
- introduce docker-compose based setup

### Refactor

- update issue template headings
- rename bug-report issue template to bug

## v0.5.6 (2024-10-03)

### Feat

- add hway and sonr processes to dev environment

## v0.5.5 (2024-10-03)

### Feat

- add rudimentary DidController table
- update home section with new features
- introduce Home model and refactor views
- **nebula**: create Home model for home page

### Refactor

- reorganize pkl files for better separation of concerns
- rename msg_server_test.go to rpc_test.go

## v0.5.4 (2024-10-02)

## v0.5.3 (2024-10-02)

### Fix

- remove unnecessary telegram message template

## v0.5.2 (2024-10-02)

### Feat

- **service**: integrate group module (#1104)

### Refactor

- revert version bump to 0.5.1

## v0.5.1 (2024-10-02)

### Refactor

- move Motr API to state package

## v0.5.0 (2024-10-02)

### Feat

- allow multiple macaroons with the same id

## v0.4.5 (2024-10-02)

### Fix

- use correct secret for docker login

## v0.4.4 (2024-10-02)

## v0.4.3 (2024-10-02)

### Feat

- **release**: add docker images for sonrd and motr
- update homepage with new visual design
- add DID to vault genesis schema
- add video component
- add video component
- add hx-get attribute to primary button in hero section

### Fix

- **layout**: add missing favicon
- **hero**: Use hx-swap for primary button to prevent flicker

### Refactor

- use single GITHUB_TOKEN for release workflow
- update workflow variables

## v0.4.2 (2024-10-01)

### Refactor

- use single GITHUB_TOKEN for release workflow

## v0.4.1 (2024-10-01)

### Feat

- Implement session management
- allow manual release triggers
- add Input and RegistrationForm models
- add new utility classes
- add login and registration pages
- add tailwindcss utilities
- add support for ARM64 architecture
- add DWN resolver field
- add stats section to homepage
- implement hero section using Pkl
- add PKL schema for message formats
- add Homebrew tap for sonr
- update release workflow to use latest tag

### Fix

- **version**: update version number to 0.4.0
- update release workflow to use latest tag
- **versioning**: revert version to 0.9.0
- **cta**: Fix typo in CTA title
- change bento section title to reflect security focus
- adjust hero image dimensions
- **Input**: Change type from to
- update hero image height in config.pkl

### Refactor

- move home page sections to home package
- rename motrd to motr
- update hero image dimensions
- move nebula configuration to static file
- rename buf-publish.yml to publish-assets.yml
- remove unused field from

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
