# CHANGELOG

## [v0.7.3-beta.1](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.1) - 2023-07-28 01:32:34

## Changelog
* 9b2d7e8 * fix(authr.go): fix method call to serialize credential
* 8926128 * feat(crypto/keys.go): add EncryptionKey interface * refactor(account.go): update NewAccountV1 function signature to return v1types.KeyshareSet instead of *v1types.Keyshare * refactor(account.go): remove unused import of github.com/sonrhq/core/pkg/did/types * refactor(account.go): remove unused variable privKs in NewAccountV1 function * refactor(account.go): update return statement in NewAccountV1 function to return v1types.EmptyKeyshareSet() instead of nil * refactor(account.go): update return statement in NewAccountV1 function to return kss instead of privKs * refactor(account.go): remove DIDIdentifier method from AccountV1 struct * refactor(account.go): remove DIDMethod method from AccountV1 struct * refactor(account.go): remove DIDUrl method from AccountV1 struct * refactor(account.go): remove comment explaining PublicKey method, as it is incomplete
* ca6183b * refactor(keyshare.go): remove unnecessary role check in GetAliceDKGResult() method * refactor(keyshare.go): remove unnecessary role check in GetBobDKGResult() method
* d61b5b3 * chore(application-services.mdx): delete application-services.mdx file from docs/static/whitepaper directory
* a6f2a7c * chore(account.go): remove unused code in Marshal method * chore(account.go): remove unused code in Unmarshal method * feat(keyshare.go): add UnmarshalAlice method to unmarshal keyshare for Alice
* 631dc29 * fix(openapi.yml): remove id property from objects in paths and definitions * chore(highway.go): rearrange middleware order in initGin function
* 0d3b537 * chore(app.go): import highlight-go SDK package * chore(app.go): import highlight-go middleware for Gorilla Mux * chore(app.go): set project ID for highlight-go SDK * chore(app.go): start highlight-go SDK * chore(main.go): stop highlight-go SDK before exiting

## [v0.7.3-beta.0](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.0) - 2023-07-27 18:09:16

## Changelog
* ec4e2fb * chore(root.go): update Viper configuration to use "SONR" as the config file name * chore(go.mod): add go.opentelemetry.io/otel v1.13.0 as a required module * chore(parser.go): delete internal/crypto/parser.go file

## [v0.7.2](https://github.com/sonrhq/core/releases/tag/v0.7.2) - 2023-07-27 11:48:51

## Changelog
* c6b209b * chore(go.mod): remove github.com/gin-gonic/autotls v0.0.5 dependency * chore(go.mod): remove golang.org/x/sync v0.2.0 dependency

## [v0.7.1](https://github.com/sonrhq/core/releases/tag/v0.7.1) - 2023-07-27 02:35:51

## Changelog
* c28f091 Migrate/mpc v1 (#101)
* 1c51a64 bugfix: Update scripts
* 64b00bd Update devcontainer.json
* 62ddb56 Update devcontainer.json
* eef99f9 Create snapcraft.yaml

## [v0.7.1-beta.4](https://github.com/sonrhq/core/releases/tag/v0.7.1-beta.4) - 2023-07-14 05:50:49

## Changelog
* 5ad51ae * feat(auth.go): add new file for the auth controller in the pkg/did/controller directory
* 734ec59 Refactor Type organization
* 2fd59e7 Add new controller account type
* 77fbe1a * refactor(openapi.yml): remove /core/params/vault endpoint * chore(openapi.yml): update /cosmos/auth/v1beta1/account_info/{address} endpoint summary
* 4a4532f Introduce Controller DID
* b041590 Simplify DID Method
* b76624e Add resolution
* fa401bc Add Account based did/method configurations
* 96cf5e2 add zkset
* 683ab4d Remove Vault, fix sfs (#100)
* cf98186 Strip highway sfs connection
* 3e8dde6 bugfix: Update MPC Tests
* 8e25abf bugfix: Update readme
* 9a16fb6 bugfix: Update container
* 41348bd Update Container
* bbb895b Change container start cmd
* a043bfd Update script
* eda6827 Update devcontainer.json
* bf59c2c * fix(devcontainer.json): update go feature version to 1.19
* cd146c6 * fix(devcontainer.json): update postAttachCommand to start icefiredb instead of server
* d7ed03d * feat(devcontainer.json): add support for fig feature * chore(devcontainer.json): update devcontainer base image to mcr.microsoft.com/devcontainers/universal:2
* 654ff37 * chore(devcontainer.json): update devcontainer features to use specific versions of ignite and icefire * chore(Taskfile.yml): remove unnecessary 'deps' task from cmds list
* 5656870 * chore(devcontainer.json): update devcontainer features * chore(settings.json): exclude git ignore files from explorer
* a729f20 update devcontainer
* acf33e1 * feat(devcontainer.json): add IceFire Redis Proxy port 6001 with label and onAutoForward configuration
* 02cddb9 * chore(devcontainer.json): remove unused go feature * feat(devcontainer.json): add go SDK version 1.19
* 11acc54 * chore(devcontainer.json): update postAttachCommand to start the server using 'task start' command instead of 'task chain' command
* 5201b67 * chore(devcontainer.json): update devcontainer.json * feat(devcontainer.json): add hostRequirements for cpus * feat(devcontainer.json): add postCreateCommand to install go-task * feat(devcontainer.json): add postAttachCommand to start task chain server * feat(devcontainer.json): customize codespaces to open README.md * feat(devcontainer.json): add portsAttributes for API server, RPC server, Highway server, and gRPC server
* 3d82029 * chore(.goreleaser.yaml): remove nfpm section * chore(.goreleaser.yaml): update caveats in brews section
* 819e0f0 * chore(.goreleaser.yaml): update license field to "Open GNU v3 License" for nfpms section * feat(.goreleaser.yaml): add brews section with brew formula update for Sonr version {{ .Tag }}
* b7d9572 * refactor(account.go): remove unused imports and variables * feat(account.go): add GetAccountData method * feat(account.go): add LinkController method * feat(account.go): add Type method
* c1edc51 * chore(settings.json): update explorer.excludeGitIgnore to false * chore(env.go): add IsAccountIceFireEnabled function to check if the account icefire is enabled
* c53a9b1 bugfix: Update git tag
* acc83dc * refactor(root.go): remove unused import of icefiredbcmd * chore(root.go): remove icefiredbcmd.CreateStartCommand() from command initialization
* 6125c05 feat: Updated cmd to have icefiredb start
* fb4c269 * fix(settings.json): exclude .gitignore files from the explorer view
* 40967f1 * feat(crypto/parser.go): add parser for DID components * feat(crypto/parser.go): add SplitDID function to split a DID into its components * feat(crypto/parser.go): add CombineDID function to combine DID components into a single string * refactor(mpc/mpc.go): rename ControllerV1 type to ZKSet for clarity * refactor(mpc/mpc.go): rename KeyshareV0 type to Keyshare for consistency * refactor(mpc/mpc.go): rename KeyshareSet type to KeyShareCollection for clarity
* 43c864e * refactor(types): rename zkbls.go to zkset.go in internal/mpc/v1/types directory
* 82199b3 * refactor(store): rename package sfs/store to highway/store * refactor(store): rename struct Store to IceFireStore * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package
* 043f535 Fix/jwt (#99)

## [v0.7.1-beta.3](https://github.com/sonrhq/core/releases/tag/v0.7.1-beta.3) - 2023-07-13 02:43:22

## Changelog
* 819e0f0 * chore(.goreleaser.yaml): update license field to "Open GNU v3 License" for nfpms section * feat(.goreleaser.yaml): add brews section with brew formula update for Sonr version {{ .Tag }}
* b7d9572 * refactor(account.go): remove unused imports and variables * feat(account.go): add GetAccountData method * feat(account.go): add LinkController method * feat(account.go): add Type method
* c1edc51 * chore(settings.json): update explorer.excludeGitIgnore to false * chore(env.go): add IsAccountIceFireEnabled function to check if the account icefire is enabled
* c53a9b1 bugfix: Update git tag
* acc83dc * refactor(root.go): remove unused import of icefiredbcmd * chore(root.go): remove icefiredbcmd.CreateStartCommand() from command initialization
* 6125c05 feat: Updated cmd to have icefiredb start
* fb4c269 * fix(settings.json): exclude .gitignore files from the explorer view
* 40967f1 * feat(crypto/parser.go): add parser for DID components * feat(crypto/parser.go): add SplitDID function to split a DID into its components * feat(crypto/parser.go): add CombineDID function to combine DID components into a single string * refactor(mpc/mpc.go): rename ControllerV1 type to ZKSet for clarity * refactor(mpc/mpc.go): rename KeyshareV0 type to Keyshare for consistency * refactor(mpc/mpc.go): rename KeyshareSet type to KeyShareCollection for clarity
* 43c864e * refactor(types): rename zkbls.go to zkset.go in internal/mpc/v1/types directory
* 82199b3 * refactor(store): rename package sfs/store to highway/store * refactor(store): rename struct Store to IceFireStore * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package * refactor(api): remove import of sfs/store package
* 043f535 Fix/jwt (#99)

## [v0.7.0](https://github.com/sonrhq/core/releases/tag/v0.7.0) - 2023-06-28 22:51:16

## Changelog
* 57411fc Refactor/split highway (#95)
* fc453d7 * chore(.goreleaser.yaml): remove scripts/* from archives files list

## [v0.6.29-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.29-beta.0) - 2023-06-23 21:25:15

## Changelog
* 7c3492a * chore(client.go): change broadcast mode from sync to async * feat(controller.go): add SignAndBroadcastCosmosTx method to Keeper struct

## [v0.6.29](https://github.com/sonrhq/core/releases/tag/v0.6.29) - 2023-06-23 21:28:33

## Changelog
* c9a0ae7 Remove WalletClaims from Vault Genesis, manage through distributed store
* a394937 * chore(.gitignore): add tmp directory to gitignore * feat(Taskfile.yml): add task to install and run database locally * feat(Taskfile.yml): add task to serve blockchain and database locally

## [v0.6.28-beta.16](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.16) - 2023-06-23 15:52:25

## Changelog
* 73a1fac chore(.goreleaser.yaml): remove generate.sh script from before section
* 2f8148e * feat(release.yml): add Github Actions workflow to automate release process * feat(release.yml): create new pre-release when pushing a new tag with a "v" prefix * feat(release.yml): create/update the "latest" pre-release when pushing to default branch * feat(release.yml): issue release assets for linux:amd64, darwin:amd64, and darwin:arm64 * feat(release.yml): delete the "latest" release if it already exists * feat(release.yml): publish the release with the specified tag name and files in the release directory as assets
* fa6c2f0 Migrate/v0.47 (#94)
* 86788d5 Fix/keyshare (#93)
* 5eb905e Add License (#89)
* 2b2e63e Integrate/encryption (#88)
* 75ae05b Integrate/encryption (#87)

## [v0.6.28](https://github.com/sonrhq/core/releases/tag/v0.6.28) - 2023-06-23 16:46:22

## Changelog
* 0d53d09 * chore(release.yml): remove release workflow file.

## [v0.6.28-beta.10](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.10) - 2023-06-11 00:42:34

## Changelog
* 971b6c4 * refactor(auth.go): remove unused import of vaulttypes * refactor(auth.go): change BroadcastTx signature to take sdk.Context and []byte as input and return error * fix(auth.go): handle error in BroadcastTx function and return it if any

## [v0.6.28-beta.2](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.2) - 2023-06-10 16:40:10

## Changelog
* b1fcf02 * chore(.goreleaser.yaml): update repository URL to ghcr.io/sonrhq/sonrd
* df248dd * chore(.goreleaser.yaml): change container registry from ghcr.io to registry.digitalocean.com/sonrhq/sonrd
* 55ad103 * chore(.goreleaser.yaml): change project_name from 'core' to 'sonrd' * refactor(.goreleaser.yaml): change builds id from 'core' to 'sonrd' * refactor(.goreleaser.yaml): change archives builds from 'core' to 'sonrd' * refactor(.goreleaser.yaml): change kos build from 'core' to 'sonrd'
* b876cf7 * chore(.goreleaser.yaml): remove preserve_import_paths and base_import_paths configuration options
* 74659a3 * chore(.goreleaser.yaml): add org.opencontainers.image.source metadata * chore(.goreleaser.yaml): add sbom, bare, preserve_import_paths, and base_import_paths options
* 3009dd7 * chore(changelog.yml): remove changelog generation workflow * chore(publish.yml): change branch filter from master to dev
* 8997e1f * fix(.goreleaser.yaml): change repository name from sonrd to core in kos section
* 5e6e6eb * fix(.goreleaser.yaml): change version label to use .Tag instead of .Version variable in kos section
* 7113af9 * chore(.goreleaser.yaml): update repository URL to ghcr.io/sonrhq/sonrd
* 4537fa3 * chore(.goreleaser.yaml): remove linux/arm64 platform from kos build * feat(.goreleaser.yaml): add version tag to kos build with {{ .Version }} format
* fa61874 * chore(.goreleaser.yaml): remove creation_time and ko_data_creation_time fields
* 174e4ac * refactor(.goreleaser.yaml): remove unused tag configuration * fix(.goreleaser.yaml): change snapcraft build configuration to use core instead of sonrd-id
* 25fa969 * feat(.goreleaser.yaml): add kos section to build and push sonrd image to registry * chore(Dockerfile): copy go.sum file to download dependencies and fix typo in comment
* 761918c * chore(Dockerfile): remove go.sum from COPY command * chore(Dockerfile.dev): remove go.sum from COPY command
* 1f67cc0 * chore(Dockerfile.dev): remove unnecessary newline * chore(go.mod): add github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 as indirect dependency * chore(go.mod): remove github.com/keybase/go-keychain v0.0.0-20190712205309-48d3d31d256d as indirect dependency * chore(go.mod): replace github.com/99designs/keyring with github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76 * chore(go.mod): update github.com/confio/ics23/go to github.com/cosmos/cosmos-sdk/ics23/go v0.8.0 * chore(go.mod): update github.com/gogo/protobuf to github.com/regen-network/protobuf v1.3.3
* f263cea * chore: remove unused ids field from snapcrafts in .goreleaser.yaml * refactor: change WORKDIR from /core to /sonr in Dockerfile and Dockerfile.dev
* 0c89f63 * refactor(Dockerfile): change working directory from /sonr to /core * chore(Dockerfile.dev): change working directory from /sonr to /core
* 2b2e63e Integrate/encryption (#88)
* 75ae05b Integrate/encryption (#87)

## [v0.6.26](https://github.com/sonrhq/core/releases/tag/v0.6.26) - 2023-05-28 20:30:43

## What's Changed
* Fix/claims by @prnk28 in https://github.com/sonrhq/core/pull/74
* Create devcontainer.json by @prnk28 in https://github.com/sonrhq/core/pull/75
* Migrate/pre mars by @prnk28 in https://github.com/sonrhq/core/pull/77
* Integrate Gateway, refactor module structure by @prnk28 in https://github.com/sonrhq/core/pull/79
* Remove docs by @prnk28 in https://github.com/sonrhq/core/pull/80


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25...v0.6.26

## [0.7.5](https://github.com/sonrhq/core/releases/tag/0.7.5) - 2023-08-14 23:49:32

*No description*

## [0.7.4](https://github.com/sonrhq/core/releases/tag/0.7.4) - 2023-08-13 23:55:00

*No description*

## [v0.7.3](https://github.com/sonrhq/core/releases/tag/v0.7.3) - 2023-08-07 23:55:39

## Changelog
* 58a93d3 * chore(Taskfile.yml): remove docker build command for cosmos-faucet target * chore(Taskfile.yml): remove docker tag and push commands for sonr-faucet image * chore(Taskfile.yml): remove publish task for k8s manifests
* 2b64d4f * refactor(Dockerfile): remove unused comments and empty lines * chore(Dockerfile): remove cosmos-faucet section
* 8ef09d9 * chore(docker-compose.yml): add dependencies for validator service * fix(docker-compose.yml): add dependencies on icefiredb and icefiresql services for validator service
* 23b160e * chore(Dockerfile): add toml-cli binary installation step * refactor(config/context.go): remove unnecessary config paths
* a623b82 * refactor(context.go): comment out unused code and remove unused imports * refactor(context.go): comment out unused code and remove unused variables
* b866cb4 * chore(Dockerfile): update sonr.yml file location to current directory * chore(root.go): remove srvCfg.MinGasPrices assignment
* 13154a2 * fix(Dockerfile): remove unnecessary argument --default-denom from sonrd init command
* dd01d5a * feat(Dockerfile): add environment variables CHAIN_ID, MONIKER, and KEYALGO * fix(Dockerfile): change sonrd init command to use MONIKER and CHAIN_ID variables * feat(Dockerfile): update config.toml and app.toml with new settings
* e7f20d4 * fix(geninit.go): remove default chain-id value in the genesis file flag * fix(root.go): change minimum gas prices from "0snr" to "0stake" * fix(config.go): set default values for node.p2p.host and node.p2p.port * fix(context.go): remove unnecessary call to viper.ReadInConfig()
* 42baae1 * feat(config.go): set default value for launch.chain-id to "sonr-localnet-1" * feat(config.go): set default value for launch.environment to "development" * feat(config.go): set default value for launch.moniker to "alice"
* e440542 * chore(Dockerfile): update sonr.yml file path in COPY command * refactor(context.go): remove error handling for viper.ReadInConfig()
* 027669c * refactor(root.go): remove unnecessary API and gRPC configurations * fix(context.go): set default values for highway.icefirekv.host, highway.icefirekv.port, and highway.jwt.key
* ce26d78 * feat(config.go): add IceFireSQLHost function to return the host and port of the IceFire SQL store * chore(sonr.yml): update host values for icefirekv and icefiresql in sonr.yml configuration file
* 1cd4f52 * chore(context.go): add default value for "highway.jwt.key" configuration option
* 7a78d24 * chore(context.go): set default values for highway.icefirekv.host and highway.icefirekv.port
* 89803a4 * chore(Dockerfile): copy sonr.yml file to current directory * fix(config.go): update config keys to use correct port values for highway, node, and launch services
* 88f76f3 * fix(Dockerfile): copy sonr.yml from builder stage to /root directory * fix(Dockerfile): remove MONIKER environment variable and hardcode "florence" as the value for sonrd init command
* 2fbdf4a * refactor(docker-compose.yml): remove unused environment variables in sonr service * chore(docker-compose.yml): remove init flag in icefiredb service
* 6939f51 * refactor(config.go): remove unused struct types and fields * refactor(config.go): update function names to match configuration keys * refactor(config.go): update function implementations to use configuration keys instead of environment variables
* 18458e0 * chore(Dockerfile): update sonr.yml file path in COPY command * chore(Dockerfile): remove --home flag from sonrd init command * chore(cmd/root.go): remove unnecessary gRPC configuration
* ec909bf * fix(Dockerfile): remove unused CHAIN_ID environment variable * fix(Dockerfile): remove --default-denom flag from sonrd init command
* 91473b3 * chore(Dockerfile): change working directory to /root * chore(Dockerfile): copy sonr.yml to current directory * chore(geninit.go): remove FlagDefaultBondDenom flag * chore(geninit.go): set sdk.DefaultBondDenom to "usnr" * chore(geninit.go): set flags.FlagChainID to "sonr-localnet-1"
* ca4f2ec * chore(Dockerfile): update sonr.yml file path in COPY command * chore(Dockerfile): rename config.go to defaults.go in cmd/sonrd/cmd directory
* e3bf2c6 * chore(Dockerfile): remove toml-cli installation step * chore(Dockerfile): remove config.toml update steps
* 23cd1ec * refactor(context.go): remove duplicate code for handling missing config file * chore(context.go): add missing newline before Environment() function * chore(context.go): add missing newline before NodeAPIHostAddress() function * chore(context.go): add missing newline before NodeGrpcHostAddress() function
* 683e5ae * refactor(context.go): convert variables to functions and use viper to retrieve values from environment variables * feat(context.go): add functions to retrieve chain ID, environment, JWT signing key, Highway host address, Highway request timeout, IceFire KV host address, Node API host address, Node gRPC host address, Node P2P host address, Node RPC host address, and validator address
* 2d9e9b8 * chore(Dockerfile): move sonr.yml to current directory * fix(Dockerfile): set --home flag to /root/.sonr in sonrd init command
* f1821ce * refactor(config.go): restructure the Config struct and its nested structs * feat(config.go): add LaunchConfig struct to store launch settings * feat(config.go): add HighwayConfig struct to store highway settings * feat(config.go): add HighwayAPIConfig struct to store API settings for highway * feat(config.go): add HighwayDBConfig struct to store database settings for highway * feat(config.go): add IcefireKVConfig struct to store IcefireKV settings for highway * feat(config.go): add IcefireSQLConfig struct to store IcefireSQL settings for highway * feat(config.go): add NodeConfig struct to store node settings * feat(config.go): add NodeAPIConfig struct to store API settings for node * feat(config.go): add NodeP2PConfig struct to store P2P settings for node * feat(config.go): add NodeRPCConfig struct to store RPC settings for node * feat(config.go): add NodeGRPCConfig
* c526a63 * chore(Dockerfile): update COPY command to copy sonr.yml from builder stage to /etc/sonr/sonr.yml * fix(Dockerfile): remove --home flag from sonrd init command * fix(config.go): remove unused binary fields from IcefireKV and IcefireSQL structs in Config * chore(sonr.yml): update chain-id to sonr-1
* d35b946 * chore(Dockerfile): rename docker/Dockerfile to Dockerfile * chore(Dockerfile.dev): rename docker/Dockerfile.dev to Dockerfile.dev * chore(Taskfile.yml): update docker build commands to use renamed Dockerfile paths
* 2bf701d * chore(docker-compose.yml): add container_name for icefiredb service * fix(Dockerfile): change relative path for copying sonr.yml
* cdc1962 * refactor(context.go): add additional config paths for sonr.yml file * refactor(context.go): rename HighwayHostPort to HighwayHostAddress * refactor(context.go): rename NodeAPIHost to NodeAPIHostAddress * refactor(context.go): rename NodeGrpcHost to NodeGrpcHostAddress * refactor(context.go): rename NodeP2PHost to NodeP2PHostAddress * refactor(context.go): rename NodeRPCHost to NodeRPCHostAddress
* 456bc8e * chore(Taskfile.yml): add publish task to publish k8s manifests to the cluster * feat(Taskfile.yml): add kompose convert command to convert manifests * feat(Taskfile.yml): add kubectl apply command to apply manifests to the cluster
* 622a9ee * chore(Taskfile.yml): add --no-cache flag to docker build command for building sonrd image * chore(Taskfile.yml): add --no-cache flag to docker build command for building sonr-faucet image
* 910d91a * fix(root.go): update flags.FlagNode value to "tcp://0.0.0.0:26657" * chore(root.go): remove unnecessary configuration for srvCfg.MinGasPrices * fix(Dockerfile): update COPY path for sonr.yml
* b01f2ae * chore(Taskfile.yml): add docker build command for sonr-faucet target * chore(Taskfile.yml): add docker tag and push commands for sonr-faucet image * chore(docker-compose.yml): add faucet service configuration
* f5964b5 Feat/sql (#106)
* 1951a12 docs(swimm): create doc: "README" (#104)

### Documentation

- swimm:
  - create doc: "README" (#104) ([1951a12](https://github.com/sonrhq/core/commit/1951a123a7fb321fe85015aaf3aed79a9b22c516)) ([#104](https://github.com/sonrhq/core/pull/104))

## [v0.7.3-beta.4](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.4) - 2023-08-02 03:50:58

## Changelog
* b11386b * feat(sonrd): add new file cored (#102)
* e19ae8e Add Gex explorer cmd
* ae220a3 * chore(bump.yml): disable creation of minor version branch * chore(bump.yml): disable creation of major version branch
* c054b42 * feat(bump.yml): add bump configuration file * feat(bump.yml): define release and branch configuration * feat(bump.yml): define categories for breaking changes, features, maintenance, bug fixes, documentation, and dependency updates * feat(bump.yml): define bump configuration for major, minor, and patch versions
* fe0820d Update README.md

## [v0.7.3-beta.3](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.3) - 2023-07-31 21:35:25

## Changelog
* 257d1ef * chore(go.mod): update github.com/go-webauthn/webauthn to v0.8.6 * chore(go.mod): update github.com/stretchr/testify to v1.8.4 * chore(go.mod): update github.com/go-webauthn/x to v0.1.4 * chore(go.mod): update github.com/golang-jwt/jwt/v5 to v5.0.0 * chore(go.mod): update github.com/google/go-tpm to v0.9.0
* 715ccae * fix(auth.go): remove unused isAuthenticated variable in SignInWithCredential function * fix(auth.go): remove unused addr variable in SignInWithCredential function * fix(auth.go): remove unused chal variable in RegisterEscrowIdentity function * fix(auth.go): add blank line after alias variable declaration in RegisterEscrowIdentity function * fix(auth.go): add blank line after alias variable declaration in RegisterEscrowIdentity function * fix(auth.go): add blank line after alias variable declaration in RegisterEscrowIdentity function * fix(auth.go): add blank line after alias variable declaration in RegisterEscrowIdentity function * fix(auth.go): add blank line after alias variable declaration in RegisterEscrowIdentity function * fix(auth.go): add blank line after alias variable declaration in RegisterEscrowIdentity function * fix(auth.go): add blank line after alias variable declaration in RegisterEscrowIdentity function * fix(auth.go): add blank line after alias variable declaration in RegisterEscrow

## [v0.7.3-beta.2](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.2) - 2023-07-31 15:46:47

## Changelog
* b769459 -m

## [v0.6.27](https://github.com/sonrhq/core/releases/tag/v0.6.27) - 2023-06-01 18:10:36

## What's Changed
* Remove/frontend by @prnk28 in https://github.com/sonrhq/core/pull/81


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.26...v0.6.27

## [v0.6.26-beta.9](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.9) - 2023-05-16 19:52:29

## What's Changed
* Create devcontainer.json by @prnk28 in https://github.com/sonrhq/core/pull/75


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.26-beta.6...v0.6.26-beta.9

## [v0.6.26-beta.6](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.6) - 2023-05-14 00:41:56

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.26-beta.5...v0.6.26-beta.6

### Refactor

- buckets.mdx:
  - convert protobuf messages to JSON objects and update BucketItem description ([65fcdd9](https://github.com/sonrhq/core/commit/65fcdd97f24ec349913be73304477a4f62ec5ec9))

## [v0.6.26-beta.5](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.5) - 2023-05-14 00:35:32

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.26-beta.4...v0.6.26-beta.5

## [v0.6.26-beta.4](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.4) - 2023-05-14 00:29:30

## What's Changed
* Fix/claims by @prnk28 in https://github.com/sonrhq/core/pull/74


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.26-beta.0...v0.6.26-beta.4

## [v0.6.26-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.0) - 2023-05-13 18:43:07

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25...v0.6.26-beta.0

### Feature

- about.mdx:
  - remove introduction about Sonr and add links to different sections of the documentation ([d755acd](https://github.com/sonrhq/core/commit/d755acd019b2022d7093656bce7f1d894d2c4faf))

## [v0.6.25](https://github.com/sonrhq/core/releases/tag/v0.6.25) - 2023-05-13 17:57:50

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.24...v0.6.25

## [v0.6.25-beta.5](https://github.com/sonrhq/core/releases/tag/v0.6.25-beta.5) - 2023-05-13 16:53:57

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25-beta.4...v0.6.25-beta.5

## [v0.6.25-beta.4](https://github.com/sonrhq/core/releases/tag/v0.6.25-beta.4) - 2023-05-13 16:16:42

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25-beta.3...v0.6.25-beta.4

## [v0.6.25-beta.3](https://github.com/sonrhq/core/releases/tag/v0.6.25-beta.3) - 2023-05-12 19:57:22

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25-beta.2...v0.6.25-beta.3

## [v0.6.25-beta.2](https://github.com/sonrhq/core/releases/tag/v0.6.25-beta.2) - 2023-05-12 19:36:49

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25-beta.1...v0.6.25-beta.2

### Feature

- create-turbo:
  - install dependencies ([0e97fdc](https://github.com/sonrhq/core/commit/0e97fdc7476254d6bd37e3ad2d5d8b8fa207d55e))
  - apply package-manager transform ([3555e2f](https://github.com/sonrhq/core/commit/3555e2fe50691633e06830fe81552f85a1c40e61))
  - apply official-starter transform ([5d3691b](https://github.com/sonrhq/core/commit/5d3691bb302039481f450182146b53c73fa8962e))

## [v0.6.25-beta.1](https://github.com/sonrhq/core/releases/tag/v0.6.25-beta.1) - 2023-05-10 15:28:26

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25-beta.0...v0.6.25-beta.1

## [v0.6.25-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.25-beta.0) - 2023-05-08 17:18:39

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.24...v0.6.25-beta.0

## [v0.6.24](https://github.com/sonrhq/core/releases/tag/v0.6.24) - 2023-05-08 13:42:52

## What's Changed
* Fix/session cache by @prnk28 in https://github.com/sonrhq/core/pull/73


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.23...v0.6.24

## [v0.6.24-beta.1](https://github.com/sonrhq/core/releases/tag/v0.6.24-beta.1) - 2023-05-05 02:10:41

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.24-beta.0...v0.6.24-beta.1

## [v0.6.24-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.24-beta.0) - 2023-05-03 17:42:04

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.23...v0.6.24-beta.0

### Feature

- service.go:
  - add did_document field to response of VerifyServiceAssertion function ([32cad11](https://github.com/sonrhq/core/commit/32cad117071e4aeef9d5801d0a1ce63f3fee0311))

## [v0.6.23](https://github.com/sonrhq/core/releases/tag/v0.6.23) - 2023-05-03 15:22:26

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.21...v0.6.23

## [v0.6.22-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.22-beta.0) - 2023-05-02 14:22:38

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.21...v0.6.22-beta.0

## [v0.6.21](https://github.com/sonrhq/core/releases/tag/v0.6.21) - 2023-05-01 02:35:02

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.20...v0.6.21

### Bug Fixes

- types/credential.pb.go:
  - remove unnecessary newline character ([30dceb5](https://github.com/sonrhq/core/commit/30dceb534a2ccacee65aabfb73dbcaf34becbc9b))

## [v0.6.20](https://github.com/sonrhq/core/releases/tag/v0.6.20) - 2023-04-30 18:45:17

## What's Changed
* Fix/webauthn options by @prnk28 in https://github.com/sonrhq/core/pull/71
* Fix/webauthn options by @prnk28 in https://github.com/sonrhq/core/pull/72


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.17...v0.6.20

### Refactor

- utils.go:
  - remove unused NewIDKeyValue function and its import statement ([7c94f44](https://github.com/sonrhq/core/commit/7c94f443210efd280629217b3b4bd3dba52125bc))

## [v0.6.17](https://github.com/sonrhq/core/releases/tag/v0.6.17) - 2023-04-19 01:09:00

## What's Changed
* Introduce/domain by @prnk28 in https://github.com/sonrhq/core/pull/70


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.12...v0.6.17

## [v0.6.12](https://github.com/sonrhq/core/releases/tag/v0.6.12) - 2023-04-13 15:32:29

## What's Changed
* Initial commit by @prnk28 in https://github.com/sonrhq/core/pull/69


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.11...v0.6.12

## [v0.6.11](https://github.com/sonrhq/core/releases/tag/v0.6.11) - 2023-04-12 20:40:19

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.10...v0.6.11

### Refactor

- identity:
  - move controller package to identity package and rename config.go to controller.go ([ca59411](https://github.com/sonrhq/core/commit/ca59411a9a26500c1693be0d0813840923097e2a))

## [v0.6.10](https://github.com/sonrhq/core/releases/tag/v0.6.10) - 2023-04-11 19:51:42

## What's Changed
* Introduce/login alias by @prnk28 in https://github.com/sonrhq/core/pull/68


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.9...v0.6.10

## [v0.6.9](https://github.com/sonrhq/core/releases/tag/v0.6.9) - 2023-04-05 23:55:01

## What's Changed
* Feature/service login by @prnk28 in https://github.com/sonrhq/core/pull/67


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.8...v0.6.9

## [v0.6.8](https://github.com/sonrhq/core/releases/tag/v0.6.8) - 2023-04-04 15:58:38

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.7...v0.6.8

## [v0.6.7](https://github.com/sonrhq/core/releases/tag/v0.6.7) - 2023-04-03 19:51:22

## What's Changed
* Introduce/service by @prnk28 in https://github.com/sonrhq/core/pull/66


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.6...v0.6.7

## [v0.6.6](https://github.com/sonrhq/core/releases/tag/v0.6.6) - 2023-04-02 23:41:16

## What's Changed
* Fix/tx broadcast by @prnk28 in https://github.com/sonrhq/core/pull/65


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.5...v0.6.6

## [v0.6.5](https://github.com/sonrhq/core/releases/tag/v0.6.5) - 2023-04-02 15:21:51

## What's Changed
* Introduce/vault encrypt by @prnk28 in https://github.com/sonrhq/core/pull/64


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.4...v0.6.5

## [v0.6.5-beta.3](https://github.com/sonrhq/core/releases/tag/v0.6.5-beta.3) - 2023-03-31 17:35:48

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.5-beta.2...v0.6.5-beta.3

## [v0.6.5-beta.2](https://github.com/sonrhq/core/releases/tag/v0.6.5-beta.2) - 2023-03-29 13:57:13

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.5-beta.1...v0.6.5-beta.2

## [v0.6.5-beta.1](https://github.com/sonrhq/core/releases/tag/v0.6.5-beta.1) - 2023-03-29 12:53:48

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.5-beta.0...v0.6.5-beta.1

## [v0.6.5-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.5-beta.0) - 2023-03-29 12:23:07

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.4...v0.6.5-beta.0

## [v0.6.4](https://github.com/sonrhq/core/releases/tag/v0.6.4) - 2023-03-29 12:00:32

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.3...v0.6.4

## [v0.6.3](https://github.com/sonrhq/core/releases/tag/v0.6.3) - 2023-03-28 08:59:33

## What's Changed
* Feature/inbox by @prnk28 in https://github.com/sonrhq/core/pull/63


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.2...v0.6.3

## [v0.6.2](https://github.com/sonrhq/core/releases/tag/v0.6.2) - 2023-03-25 18:10:40

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.1...v0.6.2

## [v0.6.1](https://github.com/sonrhq/core/releases/tag/v0.6.1) - 2023-03-25 00:26:30

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.0...v0.6.1

## [v0.6.0-beta.6](https://github.com/sonrhq/core/releases/tag/v0.6.0-beta.6) - 2023-03-24 01:51:08

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.0-beta.5...v0.6.0-beta.6

## [v0.6.0-beta.5](https://github.com/sonrhq/core/releases/tag/v0.6.0-beta.5) - 2023-03-24 01:33:02

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.0-beta.4...v0.6.0-beta.5

## [v0.6.0-beta.4](https://github.com/sonrhq/core/releases/tag/v0.6.0-beta.4) - 2023-03-23 22:58:38

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.0-beta.3...v0.6.0-beta.4

## [v0.6.0-beta.3](https://github.com/sonrhq/core/releases/tag/v0.6.0-beta.3) - 2023-03-23 03:57:10

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.0-beta.2...v0.6.0-beta.3

### Feature

- identity:
  - add "/mpc" suffix to GetOrbitDbStoreName() method in Params struct ([4151f7d](https://github.com/sonrhq/core/commit/4151f7d8cfa1f77b360a37f6d9a3253bca9eeaae))

## [v0.6.0-beta.2](https://github.com/sonrhq/core/releases/tag/v0.6.0-beta.2) - 2023-03-23 03:33:41

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.0-beta.1...v0.6.0-beta.2

## [v0.6.0-beta.1](https://github.com/sonrhq/core/releases/tag/v0.6.0-beta.1) - 2023-03-23 00:00:42

## What's Changed
* Feature/highway by @prnk28 in https://github.com/sonrhq/core/pull/61


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.5.1...v0.6.0-beta.1

## [v0.5.1](https://github.com/sonrhq/core/releases/tag/v0.5.1) - 2023-03-20 22:32:22

## What's Changed
* Feature/webauthn chain by @prnk28 in https://github.com/sonrhq/core/pull/59
* Merged type definitions from vault into x/identity by @prnk28 in https://github.com/sonrhq/core/pull/60


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.5.0...v0.5.1

## [v0.5.0](https://github.com/sonrhq/core/releases/tag/v0.5.0) - 2023-03-16 19:21:23

## What's Changed
* Integrate/connect by @prnk28 in https://github.com/sonrhq/core/pull/57
* Refactor/genesis by @prnk28 in https://github.com/sonrhq/core/pull/58


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.4...v0.5.0

## [v0.5.0-beta.1](https://github.com/sonrhq/core/releases/tag/v0.5.0-beta.1) - 2023-03-16 07:26:07

## What's Changed
* Integrate/connect by @prnk28 in https://github.com/sonrhq/core/pull/57


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.4...v0.5.0-beta.1

## [v0.4.4](https://github.com/sonrhq/core/releases/tag/v0.4.4) - 2023-03-14 20:13:46

## What's Changed
* * refactor(identity): change EmitTypedEvents to EmitEvent in CreateDi… by @prnk28 in https://github.com/sonrhq/core/pull/56


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.3...v0.4.4

## [v0.4.3](https://github.com/sonrhq/core/releases/tag/v0.4.3) - 2023-03-12 09:38:57

## What's Changed
* Feature/bip44 file store by @prnk28 in https://github.com/sonrhq/core/pull/55


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.2...v0.4.3

## [v0.4.3-beta.0](https://github.com/sonrhq/core/releases/tag/v0.4.3-beta.0) - 2023-03-10 20:19:14

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.2...v0.4.3-beta.0

## [v0.4.2](https://github.com/sonrhq/core/releases/tag/v0.4.2) - 2023-03-10 03:59:34

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.1...v0.4.2

## [v0.4.1](https://github.com/sonrhq/core/releases/tag/v0.4.1) - 2023-03-10 00:55:15

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.0...v0.4.1

## [v0.4.0](https://github.com/sonrhq/core/releases/tag/v0.4.0) - 2023-03-09 04:28:44

## What's Changed
* Feature/service management by @prnk28 in https://github.com/sonrhq/core/pull/53


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.3.2...v0.4.0

## [v0.3.2](https://github.com/sonrhq/core/releases/tag/v0.3.2) - 2023-03-07 00:24:31

## What's Changed
* Feature/grpc vault api by @prnk28 in https://github.com/sonrhq/core/pull/42


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.3.1...v0.3.2

## [v0.3.1](https://github.com/sonrhq/core/releases/tag/v0.3.1) - 2023-01-30 08:12:26

## What's Changed
* Feature/ucan by @prnk28 in https://github.com/sonrhq/core/pull/28


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.3.0...v0.3.1

## [v0.3.0](https://github.com/sonrhq/core/releases/tag/v0.3.0) - 2023-01-29 07:15:27

## What's Changed
* Feature/ipfs encryption by @prnk28 in https://github.com/sonrhq/core/pull/27


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.2.5...v0.3.0

## [v0.2.5](https://github.com/sonrhq/core/releases/tag/v0.2.5) - 2023-01-27 04:54:16

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.2.4...v0.2.5

## [v0.2.4](https://github.com/sonrhq/core/releases/tag/v0.2.4) - 2023-01-23 20:32:32

## What's Changed
* Implement/vault authorization by @prnk28 in https://github.com/sonrhq/core/pull/25


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.2.3...v0.2.4

## [v0.2.3](https://github.com/sonrhq/core/releases/tag/v0.2.3) - 2023-01-22 23:27:14

## What's Changed
* Improvement/merge vault identity by @prnk28 in https://github.com/sonr-hq/sonr/pull/24


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.2.2...v0.2.3

## [v0.2.2](https://github.com/sonrhq/core/releases/tag/v0.2.2) - 2023-01-22 21:37:51

## What's Changed
* Feature/webauthn identity by @prnk28 in https://github.com/sonr-hq/sonr/pull/23


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.2.1...v0.2.2

## [v0.2.1](https://github.com/sonrhq/core/releases/tag/v0.2.1) - 2023-01-18 23:53:55

## What's Changed
* Clean/api methods by @prnk28 in https://github.com/sonr-hq/sonr/pull/22


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.2.0...v0.2.1

## [v0.2.0](https://github.com/sonrhq/core/releases/tag/v0.2.0) - 2023-01-14 02:54:44

## What's Changed
* Commit Title: Update Prerequisites by @prnk28 in https://github.com/sonr-hq/sonr/pull/20
* Integrate NACL box for encryption by @prnk28 in https://github.com/sonr-hq/sonr/pull/21


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.14...v0.2.0

## [v0.1.14](https://github.com/sonrhq/core/releases/tag/v0.1.14) - 2023-01-11 05:01:03

## What's Changed
* Improvement/consolidate registration by @prnk28 in https://github.com/sonr-hq/sonr/pull/19


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.13...v0.1.14

## [v0.1.13](https://github.com/sonrhq/core/releases/tag/v0.1.13) - 2023-01-10 22:04:14

## What's Changed
* Fix/web chain integration by @prnk28 in https://github.com/sonr-hq/sonr/pull/18


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.12...v0.1.13

## [v0.1.12](https://github.com/sonrhq/core/releases/tag/v0.1.12) - 2023-01-09 18:29:41

## What's Changed
* Feature/vault by @prnk28 in https://github.com/sonr-hq/sonr/pull/17


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.11...v0.1.12

## [v0.1.11](https://github.com/sonrhq/core/releases/tag/v0.1.11) - 2023-01-08 00:26:46

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.10...v0.1.11

## [v0.1.10](https://github.com/sonrhq/core/releases/tag/v0.1.10) - 2023-01-08 00:19:38

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.9...v0.1.10

## [v0.1.9](https://github.com/sonrhq/core/releases/tag/v0.1.9) - 2023-01-07 22:56:22

## What's Changed
* Integrate/web chain by @prnk28 in https://github.com/sonr-hq/sonr/pull/15
* Add new Web options for cred creation by @prnk28 in https://github.com/sonr-hq/sonr/pull/16


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.8...v0.1.9

## [v0.1.8](https://github.com/sonrhq/core/releases/tag/v0.1.8) - 2023-01-07 22:34:54

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.7...v0.1.8

## [v0.1.7](https://github.com/sonrhq/core/releases/tag/v0.1.7) - 2023-01-07 21:25:19

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.6...v0.1.7

## [v0.1.6](https://github.com/sonrhq/core/releases/tag/v0.1.6) - 2023-01-07 21:15:28

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.5...v0.1.6

## [v0.1.5](https://github.com/sonrhq/core/releases/tag/v0.1.5) - 2023-01-07 20:41:40

## What's Changed
* Add API endpoints for IPFS and Vault modules by @prnk28 in https://github.com/sonr-hq/sonr/pull/14


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.4...v0.1.5

## [v0.1.4](https://github.com/sonrhq/core/releases/tag/v0.1.4) - 2023-01-07 20:00:55

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.3...v0.1.4

## [v0.1.3](https://github.com/sonrhq/core/releases/tag/v0.1.3) - 2023-01-07 04:48:24

## What's Changed
* Fix card margin layout by @prnk28 in https://github.com/sonr-hq/sonr/pull/13


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.2...v0.1.3

## [v0.1.2](https://github.com/sonrhq/core/releases/tag/v0.1.2) - 2023-01-07 04:04:12

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.1...v0.1.2

## [v0.1.1](https://github.com/sonrhq/core/releases/tag/v0.1.1) - 2023-01-07 02:59:05

## What's Changed
* Feature/web UI by @prnk28 in https://github.com/sonr-hq/sonr/pull/10
* Add Link Component and EmptyState to Home page by @prnk28 in https://github.com/sonr-hq/sonr/pull/11
* Polish/web by @prnk28 in https://github.com/sonr-hq/sonr/pull/12


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.0...v0.1.1

## [v0.1.0](https://github.com/sonrhq/core/releases/tag/v0.1.0) - 2023-01-04 22:21:55

## What's Changed
* Closes SNR-50 by @prnk28 in https://github.com/sonr-hq/sonr/pull/5
* Snr 71 by @prnk28 in https://github.com/sonr-hq/sonr/pull/6
* Complete/vault by @prnk28 in https://github.com/sonr-hq/sonr/pull/7
* Snr 55 by @prnk28 in https://github.com/sonr-hq/sonr/pull/8
* Auto/visualize repo by @prnk28 in https://github.com/sonr-hq/sonr/pull/9


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.0.3...v0.1.0

## [v0.0.3](https://github.com/sonrhq/core/releases/tag/v0.0.3) - 2022-12-27 21:03:35

## What's Changed
* Snr 69 build offlineonline implementations of by @prnk28 in https://github.com/sonr-hq/sonr/pull/4


**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.0.2...v0.0.3

## [v0.0.2](https://github.com/sonrhq/core/releases/tag/v0.0.2) - 2022-12-23 07:23:20

## What's Changed
* Feature/ipfs node by @prnk28 in https://github.com/sonr-hq/sonr/pull/1
* Merge SNR-53 by @prnk28 in https://github.com/sonr-hq/sonr/pull/2
* Fixes SNR-56 by @prnk28 in https://github.com/sonr-hq/sonr/pull/3

## New Contributors
* @prnk28 made their first contribution in https://github.com/sonr-hq/sonr/pull/1

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.0.1...v0.0.2

## [v0.0.1](https://github.com/sonrhq/core/releases/tag/v0.0.1) - 2022-12-13 19:22:04

**Full Changelog**: https://github.com/sonr-hq/sonr/commits/v0.0.1

\* *This CHANGELOG was automatically generated by [auto-generate-changelog](https://github.com/BobAnkh/auto-generate-changelog)*
