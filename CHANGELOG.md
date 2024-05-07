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
