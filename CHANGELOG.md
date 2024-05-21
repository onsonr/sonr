## 0.5.0 (2024-05-13)

### Feat

- **app**: rename ChainApp to SonrApp and initialize proxy in app.go refactor(app): change ChainApp to SonrApp in msg_filter_test.go refactor(app): adjust function parameter in 'app/encoding.go' from 'ChainApp' to 'SonrApp' refactor(app): rename ChainApp to SonrApp in export.go refactor(app): change ChainApp to SonrApp in sim_test.go setupSimulationApp function refactor(app): rename ChainApp instances to SonrApp in test_helpers.go refactor(app): rename ChainApp to SonrApp in test_support.go refactor(app): rename ChainApp to SonrApp in upgrades.go fix(cmd/sonrd): replace ChainApp with SonrApp in commands.go fix(devbox.json): remove 'up' and 'kill' scripts from 'scripts' section feat(go.mod): add bool64/cache dependency to version v0.4.7 feat: add bool64/cache v0.4.7 to go.sum dependencies feat(proxy): add cache management functionalities for session and challenge in pkg/proxy/cache.go feat(proxy): add new file challenge.go with Challenge struct feat(proxy): add session struct to session.go file feat(proxy): add new 'token.go' file in 'pkg/proxy' directory feat(pkg/vault): remove config.go file feat(pkg/vault): remove ipfs.go file feat(vault): remove vault.go and associated structures
- **.gitignore**: add terraform related files to ignore list
- Delete 'interchaintest-downloader' subproject
- **api/did/v1**: delete auth.pulsar.go file
- **.github/workflows**: add check-pr.yml to monitor PR events
- **did**: remove attestation.go from pkg/did feat(pkg/did): remove delegation.go file chore(did): remove invocation.go from pkg/did directory feat(did): remove service.go from the did package feat(did): add new verification.go file in pkg/did
- **.envrc**: add Nushell integration command to script fix(devbox.json): update 'dev' script in shell scripts section feat(middleware): add new 'request.go' file to 'x/did/internal/middleware' package
- **devbox.json**: add golangci-lint to devbox and update scripts feat(devbox.lock): add golangci-lint@latest with version 1.58.0 to lock file feat(x/did/internal/middleware): add new request.go file
- **scripts**: remove dev_deps.sh functionality
- **devbox.json**: update shell scripts for 'dev', 'testnet', and 'svc:up' commands
- Remove 'Taskfile.yml' fix(devbox.json): rename script proto to gen
- **Caddyfile**: remove deploy/etc/caddy configuration file feat(deploy/rails): remove process-compose.yml file fix(devbox.json): update shell scripts and add new service commands feat(vault): add new 'config.go' file in 'pkg/vault' feat(pkg/vault): add new ipfs client with key creation functionality feat(vault): add new 'vault.go' file with 'VaultFS' interface and 'vaultFS' struct feat: Add new 'process-compose.yml' for starting Django feat(x/did/controller): add new encrypt.go file
- **x/did/controller**: add new file 'controller.go' with keyshare operations functions feat(did/controller): add tests for controller functionality in 'controller_test.go' feat(did/controller): add new file encryptor.go feat(x/did/controller): add new file 'keyshares.go' for keyshare handling functionality feat(x/did/controller): add new network controller with keyshare generation feat(did/controller): add new property.go file with accumulator functionality feat(did/controller): add new signer.go file feat(x/did/controller): add new zk.go file with accumulator functionality refactor(x/did/keeper): change import path for controller module
- **x/did/types**: add WhitelistedOrigins to DefaultParams in genesis.go
- **x/did/types**: add new error types in 'errors.go'
- **Taskfile.yml**: add aliases and modify cmds for rails tasks feat(caddy/Caddyfile): add new configuration file feat: Add new process-compose.yml for starting django with ipfs daemon fix(devbox.json): update the path of CADDY_CONFIG in env variables
- **Taskfile.yml**: add commands for starting devbox services in background
- **Taskfile.yml**: set silent mode for task commands and run devbox services in background fix(devbox.json): update CADDY_CONFIG path and remove IPFS_PATH environment variable feat(rails/caddy): add new Caddyfile with basic server configuration
- **Taskfile.yml**: add new tasks for devbox services management feat(devbox.json): add new dependencies and environment variables feat(devbox.lock): add doppler, gum, mods, zellij to lockfile feat: Add new process-compose.yml for starting Django fix(scripts): remove version display of pkl in dev_deps.sh
- **devbox.json**: add 'caddy@latest' to project dependencies and env settings feat(devbox.lock): add caddy@latest with version 2.7.6 to the lock file feat(caddy): add new Caddyfile with configuration settings
- **x/did/types**: remove Ethereum related code from address.go

### Fix

- **network**: remove redundant error handling in 'runMpcProtocol' function
- **.gitignore**: add .DS_Store to ignored files list chore(crypto): update binary files in .DS_Store feat(process-compose.yml): Add sonrd command to processes feat(auth.templ): add createCredential and getCredential functions to `ui` package feat(auth_templ.go): add createCredential and getCredential functions feat(x/did/internal/ui): add new auth_templ.txt for tab functionality and form layout
- **devbox.json**: update environment settings and clean shell scripts
- **server**: change database table in insert method from AssertionTable to AuthenticatorTable refactor(x/did/types): change Assertion to Authenticator in state.go
- **devbox.json**: split 'gen' and 'dev' tasks into separate scripts feat(interchaintest-downloader): add new subproject with commit 74fa7144d84e33377d162e8db9b328bf5e2a3dd5 feat(ui): add new 'exports.go' file in 'ui' package feat(ui): initialize ui package in pkg/ui/ui.go
- **go.mod**: update dependencies, add direct dep on 'github.com/a-h/templ' v0.2.680 fix(go.sum): update dependencies to new versions
- **go.mod**: update dependencies versions
- **devbox.json**: remove unused dependencies and update shell scripts feat(devbox.lock): update packages and versions
- **devbox.json**: replace init_hook script and update package list fix(devbox.lock): replace docker-compose with dendrite and adjust version feat(scripts): add new dev_deps.sh to automatically install pkl

### Refactor

- **Taskfile.yml**: streamline task definitions and update naming conventions refactor(cmd/sonrd): remove unused codec import and parameter in initRootCmd function refactor(root.go): remove codec argument from initRootCmd function call in NewRootCmd feat(testnet): Add new process-compose file for starting django in test network feat(deploy/rails): add new process-compose.yml for starting django feat(x/did/controller): add new file encryptor.go
- **Taskfile.yml**: remove unused GREETING var
- **x/did/controller**: simplify controller and use types from 'types' package refactor(x/did/controller): simplify keyshare structs and methods in keyshares.go refactor(x/did/controller): update type definitions and function return types in network.go

## 0.4.0 (2024-05-08)

### Feat

- **x/did/keeper**: implement InitializeController in server.go feat(x/did): add new router.go file in module package feat(did/types): add NewMsgInitializeController function in msgs.go feat(x/did/types): add state.go with conversion functions for different lists
- **x/did/controller**: add keyshares.go for user and validator keyshare management
- **config**: add new chain/app.pkl configuration file feat(config): add new chain configuration file config.pkl feat(config): add new 'toml.pkl' configuration file feat(hway.pkl): Add new 'hway.pkl' file in 'config/pkl/rails' directory feat(ipfs.pkl): Add new 'config/pkl/rails/ipfs.pkl' file feat(config/pkl/rails): add new 'matrix.pkl' configuration file feat(config/pkl/rails): add new file 'services.pkl' with sonrEventsBridge configurations feat(did): add new 'invocation.go' file in 'pkg/did' directory feat(controller): add Controller interface and update controller struct types in did/controller/controller.go feat(x/did/controller): add new file for keyshare set functionality feat(x/did/controller): add network.go for keyshare interactions with the IPFS network feat(x/did/controller): add new file property.go feat(x/did/controller): add new signer.go file refactor(x/did/keeper): change types.KeyshareSet to controller.KeyshareSet and types.IController to controller.Controller refactor(did/types): simplify controller.go by removing unnecessary IController interface
- Remove '.zed/tasks.json' file feat: Add Taskfile.yml, enhance DIDs handling, improve code readability in CHANGELOG.md feat(devbox.json): update toolbox and shell scripts feat(devbox.lock): add go-task, ipfs and templ modules to devbox configuration feat(pkg/did): add new assertion.go file feat(did): add new attestation.go file in pkg/did directory feat(did): add new delegation.go file in pkg/did feat(did): add new file 'invokation.go' in 'pkg/did' feat(pkg/did): add new service.go file in 'did' package
- add Taskfile.yml for task management in root, x/did, x/oracle, and x/svc directories feat(did): add identifier.go and method.go for handling email and phone DIDs
- **app.go, depinject.go, keeper.go, keeper_test.go**: add AccountKeeper to DID Keeper for account management style(README.md): change list formatting from '--' to '-' for better readability style(app.go): remove unnecessary blank line for cleaner code

### Fix

- **.gitignore**: adjust file to ignore specific .env instead of all .env* files feat(Taskfile.yml): add new tasks for protobuf, templ files, devnet, and testnet fix(configs): update start_time and addresses in logs.json

### Refactor

- **x/did/keeper**: replace ConvertByteArrayTo*List methods with Get*List methods in InitializeController function refactor(x/did/types/msgs.go): rename getter functions for clarity

## 0.3.0 (2024-05-07)

### Feat

- add property type for accumulator values

### Refactor

- **identifier.go, codec.go**: move Blake3Hash function from identifier.go to codec.go for better code organization fix(identifier.go): replace local Blake3Hash function calls with types.Blake3Hash to reflect the function move
- **controller.go**: remove unused props field from controller struct feat(identifier.go): add new file for handling email and phone DIDs refactor(types/controller.go): rename Controller interface to IController for clarity refactor(types/genesis.go): remove unused DefaultCurve and ReferralRewardRate fields from DefaultParams function

## 0.2.0 (2024-05-07)

## 0.1.0 (2024-05-07)

### Feat

- **CHANGELOG.md**: add CHANGELOG.md to document project changes and updates

## 0.0.1 (2024-05-07)

### Feat

- add .cz.toml file to configure commitizen for conventional commits
- **.gitignore**: add .opencommitignore to gitignore to prevent accidental commit of opencommitignore files
- **controller**: add accumulator.go for BLS scheme operations This new file includes functions for creating and updating accumulators, creating witnesses, verifying witnesses, and converting values to elements.
- add property type for accumulator values
- **devbox.json**: add helix to the list of tools to keep dev environment up-to-date refactor(controller.go): move Controller interface to types package for better organization refactor(controller_test.go, vault.go, keeper.go): update function names and types to match new interface location feat(types/controller.go): create new file for Controller interface to improve code structure remove(crypto/dkg/benchmarks.txt): delete obsolete benchmark results file to clean up the repository
- add property type for accumulator values
- **controller**: add new controller for DID scheme with keyshare and accumulator support - Implement Controller interface with methods for linking, signing, refreshing, unlinking, and validating. - Add utility functions for running MPC protocol and converting values to elements. - Include tests for controller signing, address conversion, property linking/unlinking, and non-existent property unlinking.
- **vault**: add vaultStore for interacting with Keyshares in IPFS network
- Implement NewController method to create a new controller instance.
- Add GenerateKSS method to generate both keyshares.
- **zk**: add secret key for BLS scheme with accumulator support
- Implement methods for creating accumulator, creating witness, verifying witness, getting public key, and updating accumulator.
- zk properties for did
- **.gitignore**: add 'heighliner*' to gitignore to prevent accidental commit of heighliner related files
- zk properties for did
- **commands.go**: set coin type to 703 for better compatibility with BIP44 standard refactor(root.go): change viper configuration string to "SONR" for better project identification refactor(controller_test.go): rename TestController to TestControllerSigning for better test clarity feat(controller_test.go): add TestAddressConversion to validate address conversion and validation for Bitcoin, Ethereum, and Sonr addresses
- **controller_test.go**: add public key validation test to ensure key generation is correct refactor(controller_test.go): move public key assignment to improve test readability feat(keyshares.go): add PublicKey method to KeyshareSet to provide public key access
- **crypto**: add new iface.go file for crypto package refactor(ipfs): modify client.go to handle IPFS node reachability refactor(ipfs): remove keys.go file and move IPFSKey type to client.go refactor(did): remove Verify method from Controller interface and its implementation refactor(did): modify controller_test.go to use PublicKey's VerifySignature method refactor(did): rename GenerateKSS method to GenerateKeyshares in keeper.go refactor(did): modify PublicKey methods in user_ks.go and val_ks.go to handle errors feat(did): add new address.go and keyshares.go files in types package
- zk properties for did
- **controller.go**: add PublicKey and Unlink methods to Controller interface for additional functionality refactor(controller.go): modify Refresh and Sign methods to use StartKsProtocol for better code reuse refactor(controller.go): change deriveSecretKey to DeriveSecretKey and make it public for wider usage feat(controller_test.go): add unit tests for controller.go refactor(keeper.go): change generateKSS to GenerateKSS and make it public for wider usage refactor(protocol.go): move startKsProtocol to controller.go and make it public as StartKsProtocol refactor(vault.go): use GenerateKSS instead of generateKSS following the changes in keeper.go
- **controller.go**: add Controller interface and implement it in controller struct feat(controller.go): add CreateController function to create a controller instance feat(controller.go): add Link function to link a property to the controller feat(controller.go): add Validate function to validate that a property is linked to the controller feat(controller.go): refactor Refresh function to reset properties map feat(controller.go): refactor deriveSecretKey function to derive the secret key from the keyshares feat(vault.go): refactor NewController function to use CreateController function feat(zk.go): add Property and Witness types with associated functions feat(zk.go): refactor CreateWitness function to return Property and Witness feat(zk.go): refactor VerifyElement function to use accumulator.PublicKey feat(zk.go): add encodeProperty and encodeWitness functions to encode accumulator and witness to base58 string feat(zk.go): add decodeProperty and decodeWitness functions to decode accumulator and witness from base58 string
- add support for new libraries and refactor code for better error handling
- **ipfs/client.go**: refactor IPFSClient to struct and add NewKey method for key generation feat(ipfs/keys.go): add new file for IPFSKey and IPFSPublicKey types and related methods feat(did/keeper/controller.go): add DeriveSecretKey method to generate new secret key feat(did/keeper/genesis.go): add Logger method to return logger feat(did/keeper/keeper.go): add GenerateKSS method to generate new keyshare set refactor(did/keeper/zk.go): remove DeriveSecretKey method, moved to controller.go
- **ipfs**: add new file fs.go to define VaultFS interface for IPFS vault file system refactor(keeper.go): replace Schema with OrmDB for state management to use ORM database refactor(keeper.go): refactor NewKeeper function to initialize OrmDB and StateStore refactor(protocol.go, user_ks.go, val_ks.go): rename NewValidatorKeyshare and NewUserKeyshare to createValidatorKeyshare and createUserKeyshare for better semantics refactor(user_ks.go, val_ks.go): change PublicKey return type to *types.PublicKey for better type safety feat(zk.go): add new file zk.go to implement BLS scheme for zero-knowledge proofs feat(keys.go): add ORMModuleSchema to define ORM schema for the module
- **github-actions**: add new workflow buf-publish.yml to automatically publish proto files to buf.build/didao on push event
- **pkl**: add new configuration files for chain, rails, and matrix services
- add mpc protocol and vault system
- add mpc protocol and vault system
- **devbox**: add Dockerfile, devcontainer.json, .envrc, and devbox.json files

### Fix

- **controller.go**: improve error handling in Check function
The Check function now returns false immediately when an error occurs, instead of proceeding with the rest of the function. This prevents potential issues caused by using invalid data.
- **github workflows**: downgrade GO_VERSION from 1.22.2 to 1.21 in docker-release.yml, interchaintest-e2e.yml, and unit-test.yml for compatibility issues
- **go.mod**: update Go version from 1.22 to 1.22.2 for compatibility with latest features and security updates refactor(go.mod): remove redundant 'toolchain' line to clean up code
- **protocgen.sh**: remove unnecessary move command for crypto directory chore: add .DS_Store to .gitignore to prevent system files from being tracked
- **proto**: Add did state orm

### Refactor

- **controller.go**: rearrange Set and Check functions for better readability
The Set function is moved below the Check function. The Check function is now placed before the PublicKey function to improve the flow of the code.
- rename ControllerI interface to IController in controller.go, vault.go, keeper.go for better readability and consistency with other interface naming conventions
- **controller_test.go**: move propertyKey declaration to top of TestLinkUnlinkProperty function for better readability and code organization
- **controller.go**: remove runMpcProtocol function to improve code organization feat(vault.go): add runMpcProtocol function to better align with the responsibility of the vault module
- **did**: change KeyshareSet from struct to interface for better abstraction feat(did): add getter methods for user and validator keyshares in KeyshareSet refactor(did): update function signatures to use KeyshareSet interface instead of struct
- **did**: move controller logic from keeper to controller package for better code organization feat(did): update controller tests to use new controller package remove(did): delete unused keeper/controller.go and keeper/controller_test.go files remove(did): delete unused keeper/vault.go and keeper/zk.go files refactor(did): update keeper.go to use new controller package
- **genesis.go**: remove WhitelistedVerifications from DefaultParams function to simplify the code as it was not being used
- **did/keeper/controller.go**: change return types of Link and Unlink methods to improve type safety refactor(did/keeper/controller.go): replace string-based accumulators with actual accumulator objects for better type safety and performance refactor(did/keeper/controller.go): simplify and optimize the Link, Unlink, and Validate methods refactor(did/keeper/zk.go): replace string-based accumulators and witnesses with actual objects for better type safety and performance test(did/keeper/controller_test.go): add tests for Link, Unlink, and Validate methods to ensure correctness refactor(did/keeper/keeper.go): adjust LinkController method to new return type of Controller.Link method refactor(did/keeper/zk.go): remove unused base58 encoding/decoding functions
- **commands.go, root.go**: remove hardcoded coin type and viper string for better flexibility test(controller_test.go): add validation checks for Bitcoin and Ethereum addresses to ensure correctness refactor(address.go): update address creation methods to use updated libraries and remove unnecessary steps for cleaner code
- **controller.go**: rename DeriveSecretKey to deriveSecretKey and StartKsProtocol to runMpcProtocol for better readability and consistency refactor(controller_test.go, keeper.go, vault.go): replace GenerateKSS function parameters to use a single KeyshareSet instead of separate UserKeyshare and ValidatorKeyshare for better encapsulation refactor(genesis.go): move DefaultParams, String and Validate methods from params.go to genesis.go for better organization remove(params.go): remove redundant file after moving its methods to genesis.go
- **did**: update keyshare handling and address creation methods for better encapsulation and clarity - Replace UserKeyshare and ValidatorKeyshare structs with interfaces and private implementations to hide internal details - Replace individual keyshare creation methods with a single NewKeyshareSet method to simplify keyshare creation - Rename address creation methods to start with 'Create' for better clarity - Remove unused files (protocol.go, user_ks.go, val_ks.go) to clean up the codebase - Update controller.go and vault.go to use the new keyshare and address creation methods - These changes improve the code organization and make the keyshare handling and address creation more intuitive
- **controller.go**: remove unused proofKey from controller struct for code cleanup fix(controller.go): return error from zkVerifyElement function to improve error handling feat(controller.go): use Blake3Hash for propertyKey before signing to enhance security test(controller_test.go): remove tests for unused functions and adjust existing tests to match updated functions refactor(zk.go): modify zkVerifyElement to return error for better error handling
- Rename x/idx to x/did for clarity
- Rename x/idx to x/did for clarity
