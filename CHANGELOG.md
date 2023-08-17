# CHANGELOG

## [v0.7.5](https://github.com/sonrhq/core/releases/tag/v0.7.5) - 2023-08-17 02:42:37

![Release Image](https://api.placid.app/u/cpqufnq7i?&commit[text]=024ba0ae5582ef9cd37155701eb8f720bde317dc&date[text]=2023-08-17T02:43:24Z&version[text]=0.7.5)

< DESCRIPTION OF RELEASE >

## Changelog

See the full changelog [here](https://github.com/sonrhq/core/blob/v0.7.5/CHANGELOG.md)

## ⚡️ Binaries

Binaries for Linux and Darwin (amd64 and arm64) are available below.
Darwin users can also use the same universal binary `sonrd-0.7.5-darwin-all` for both amd64 and arm64.

#### 🔨 Build from source

If you prefer to build from source, you can use the following commands:

````bash
git clone https://github.com/sonrhq/core
cd core && git checkout v0.7.5
make build-darwin # or make build-linux
````

## 🐳 Run with Docker

As an alternative to installing and running sonrd on your system, you may run sonrd in a Docker container.
The following Docker images are available in our registry:

| Image Name                              | Base                                 | Description                       |
|-----------------------------------------|--------------------------------------|-----------------------------------|
| `sonrhq/core:0.7.5`            | `distroless/static-debian11`         | Default image based on Distroless |
| `sonrhq/core:0.7.5-distroless` | `distroless/static-debian11`         | Distroless image (same as above)  |
| `sonrhq/core:0.7.5-nonroot`    | `distroless/static-debian11:nonroot` | Distroless non-root image         |
| `sonrhq/core:0.7.5-alpine`     | `alpine`                             | Alpine image                      |

Example run:

```bash
docker run sonrhq/sonrd:0.7.5 version
# v0.7.5
````

All the images support `arm64` and `amd64` architectures.

## [v0.7.5-beta.1](https://github.com/sonrhq/core/releases/tag/v0.7.5-beta.1) - 2023-08-16 23:04:23

![Release Image](https://api.placid.app/u/cpqufnq7i?&commit[text]=878ea17bba90be8b45a00308fe5a7aa4e35fc2ce&date[text]=2023-08-17T02:22:39Z&version[text]=0.7.5-beta.1)

< DESCRIPTION OF RELEASE >

## Changelog

See the full changelog [here](https://github.com/sonrhq/core/blob/v0.7.5-beta.1/CHANGELOG.md)

## ⚡️ Binaries

Binaries for Linux and Darwin (amd64 and arm64) are available below.
Darwin users can also use the same universal binary `sonrd-0.7.5-beta.1-darwin-all` for both amd64 and arm64.

#### 🔨 Build from source

If you prefer to build from source, you can use the following commands:

````bash
git clone https://github.com/sonrhq/core
cd core && git checkout v0.7.5-beta.1
make build-darwin # or make build-linux
````

## 🐳 Run with Docker

As an alternative to installing and running sonrd on your system, you may run sonrd in a Docker container.
The following Docker images are available in our registry:

| Image Name                              | Base                                 | Description                       |
|-----------------------------------------|--------------------------------------|-----------------------------------|
| `sonrhq/core:0.7.5-beta.1`            | `distroless/static-debian11`         | Default image based on Distroless |
| `sonrhq/core:0.7.5-beta.1-distroless` | `distroless/static-debian11`         | Distroless image (same as above)  |
| `sonrhq/core:0.7.5-beta.1-nonroot`    | `distroless/static-debian11:nonroot` | Distroless non-root image         |
| `sonrhq/core:0.7.5-beta.1-alpine`     | `alpine`                             | Alpine image                      |

Example run:

```bash
docker run sonrhq/sonrd:0.7.5-beta.1 version
# v0.7.5-beta.1
````

All the images support `arm64` and `amd64` architectures.

## [v0.7.4](https://github.com/sonrhq/core/releases/tag/v0.7.4) - 2023-08-13 23:55:00

This release has no changes

## What's Changed
* Fix/docker deploy [OSS-4] by @prnk28 in https://github.com/sonrhq/core/pull/107
* Setup/testnet env [OSS-6] by @prnk28 in https://github.com/sonrhq/core/pull/108


**Full Changelog**: https://github.com/sonrhq/core/compare/v0.7.3...v0.7.4

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

## [v0.7.1](https://github.com/sonrhq/core/releases/tag/v0.7.1) - 2023-07-27 02:35:51

## [v0.7.1-beta.4](https://github.com/sonrhq/core/releases/tag/v0.7.1-beta.4) - 2023-07-14 05:50:49

## [v0.7.1-beta.3](https://github.com/sonrhq/core/releases/tag/v0.7.1-beta.3) - 2023-07-13 02:43:22

## [v0.7.0](https://github.com/sonrhq/core/releases/tag/v0.7.0) - 2023-06-28 22:51:16

## [v0.6.29-beta.5](https://github.com/sonrhq/core/releases/tag/v0.6.29-beta.5) - 2023-06-23 21:28:33

## [v0.6.29-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.29-beta.0) - 2023-06-23 21:25:15

## [v0.6.29](https://github.com/sonrhq/core/releases/tag/v0.6.29) - 2023-06-23 21:28:33

## [v0.6.28-beta.16](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.16) - 2023-06-23 15:52:25

## [v0.6.28](https://github.com/sonrhq/core/releases/tag/v0.6.28) - 2023-06-23 16:46:22

## [v0.6.28-beta.9](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.9) - 2023-06-11 00:42:34

## [v0.6.28-beta.10](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.10) - 2023-06-11 00:42:34

## [v0.6.28-beta.2](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.2) - 2023-06-10 16:40:10

## [v0.6.27](https://github.com/sonrhq/core/releases/tag/v0.6.27) - 2023-06-01 18:10:36

## [v0.6.26](https://github.com/sonrhq/core/releases/tag/v0.6.26) - 2023-05-28 20:30:43

## [v0.6.26-beta.9](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.9) - 2023-05-16 19:52:29

## [v0.6.26-beta.6](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.6) - 2023-05-14 00:41:56

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.26-beta.5...v0.6.26-beta.6

## [v0.6.26-beta.5](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.5) - 2023-05-14 00:35:32

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.26-beta.4...v0.6.26-beta.5

## [v0.6.26-beta.4](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.4) - 2023-05-14 00:29:30

## [v0.6.26-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.26-beta.0) - 2023-05-13 18:43:07

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25...v0.6.26-beta.0

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

## [v0.6.25-beta.1](https://github.com/sonrhq/core/releases/tag/v0.6.25-beta.1) - 2023-05-10 15:28:26

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.25-beta.0...v0.6.25-beta.1

## [v0.6.25-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.25-beta.0) - 2023-05-08 17:18:39

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.24...v0.6.25-beta.0

## [v0.6.24](https://github.com/sonrhq/core/releases/tag/v0.6.24) - 2023-05-08 13:42:52

## [v0.6.24-beta.1](https://github.com/sonrhq/core/releases/tag/v0.6.24-beta.1) - 2023-05-05 02:10:41

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.24-beta.0...v0.6.24-beta.1

## [v0.6.24-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.24-beta.0) - 2023-05-03 17:42:04

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.23...v0.6.24-beta.0

## [v0.6.23](https://github.com/sonrhq/core/releases/tag/v0.6.23) - 2023-05-03 15:22:26

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.21...v0.6.23

## [v0.6.22-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.22-beta.0) - 2023-05-02 14:22:38

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.21...v0.6.22-beta.0

## [v0.6.21](https://github.com/sonrhq/core/releases/tag/v0.6.21) - 2023-05-01 02:35:02

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.20...v0.6.21

## [v0.6.20](https://github.com/sonrhq/core/releases/tag/v0.6.20) - 2023-04-30 18:45:17

## [v0.6.17](https://github.com/sonrhq/core/releases/tag/v0.6.17) - 2023-04-19 01:09:00

## [v0.6.12](https://github.com/sonrhq/core/releases/tag/v0.6.12) - 2023-04-13 15:32:29

## [v0.6.11](https://github.com/sonrhq/core/releases/tag/v0.6.11) - 2023-04-12 20:40:19

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.10...v0.6.11

## [v0.6.10](https://github.com/sonrhq/core/releases/tag/v0.6.10) - 2023-04-11 19:51:42

## [v0.6.9](https://github.com/sonrhq/core/releases/tag/v0.6.9) - 2023-04-05 23:55:01

## [v0.6.8](https://github.com/sonrhq/core/releases/tag/v0.6.8) - 2023-04-04 15:58:38

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.7...v0.6.8

## [v0.6.7](https://github.com/sonrhq/core/releases/tag/v0.6.7) - 2023-04-03 19:51:22

## [v0.6.6](https://github.com/sonrhq/core/releases/tag/v0.6.6) - 2023-04-02 23:41:16

## [v0.6.5](https://github.com/sonrhq/core/releases/tag/v0.6.5) - 2023-04-02 15:21:51

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

## [v0.6.0-beta.2](https://github.com/sonrhq/core/releases/tag/v0.6.0-beta.2) - 2023-03-23 03:33:41

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.6.0-beta.1...v0.6.0-beta.2

## [v0.6.0-beta.1](https://github.com/sonrhq/core/releases/tag/v0.6.0-beta.1) - 2023-03-23 00:00:42

## [v0.5.1](https://github.com/sonrhq/core/releases/tag/v0.5.1) - 2023-03-20 22:32:22

## [v0.5.0](https://github.com/sonrhq/core/releases/tag/v0.5.0) - 2023-03-16 19:21:23

## [v0.5.0-beta.1](https://github.com/sonrhq/core/releases/tag/v0.5.0-beta.1) - 2023-03-16 07:26:07

## [v0.4.4](https://github.com/sonrhq/core/releases/tag/v0.4.4) - 2023-03-14 20:13:46

## [v0.4.3](https://github.com/sonrhq/core/releases/tag/v0.4.3) - 2023-03-12 09:38:57

## [v0.4.3-beta.0](https://github.com/sonrhq/core/releases/tag/v0.4.3-beta.0) - 2023-03-10 20:19:14

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.2...v0.4.3-beta.0

## [v0.4.2](https://github.com/sonrhq/core/releases/tag/v0.4.2) - 2023-03-10 03:59:34

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.1...v0.4.2

## [v0.4.1](https://github.com/sonrhq/core/releases/tag/v0.4.1) - 2023-03-10 00:55:15

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.4.0...v0.4.1

## [v0.4.0](https://github.com/sonrhq/core/releases/tag/v0.4.0) - 2023-03-09 04:28:44

## [v0.3.2](https://github.com/sonrhq/core/releases/tag/v0.3.2) - 2023-03-07 00:24:31

## [v0.3.1](https://github.com/sonrhq/core/releases/tag/v0.3.1) - 2023-01-30 08:12:26

## [v0.3.0](https://github.com/sonrhq/core/releases/tag/v0.3.0) - 2023-01-29 07:15:27

## [v0.2.5](https://github.com/sonrhq/core/releases/tag/v0.2.5) - 2023-01-27 04:54:16

**Full Changelog**: https://github.com/sonrhq/core/compare/v0.2.4...v0.2.5

## [v0.2.4](https://github.com/sonrhq/core/releases/tag/v0.2.4) - 2023-01-23 20:32:32

## [v0.2.3](https://github.com/sonrhq/core/releases/tag/v0.2.3) - 2023-01-22 23:27:14

## [v0.2.2](https://github.com/sonrhq/core/releases/tag/v0.2.2) - 2023-01-22 21:37:51

## [v0.2.1](https://github.com/sonrhq/core/releases/tag/v0.2.1) - 2023-01-18 23:53:55

## [v0.2.0](https://github.com/sonrhq/core/releases/tag/v0.2.0) - 2023-01-14 02:54:44

## [v0.1.14](https://github.com/sonrhq/core/releases/tag/v0.1.14) - 2023-01-11 05:01:03

## [v0.1.13](https://github.com/sonrhq/core/releases/tag/v0.1.13) - 2023-01-10 22:04:14

## [v0.1.12](https://github.com/sonrhq/core/releases/tag/v0.1.12) - 2023-01-09 18:29:41

## [v0.1.11](https://github.com/sonrhq/core/releases/tag/v0.1.11) - 2023-01-08 00:26:46

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.10...v0.1.11

## [v0.1.10](https://github.com/sonrhq/core/releases/tag/v0.1.10) - 2023-01-08 00:19:38

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.9...v0.1.10

## [v0.1.9](https://github.com/sonrhq/core/releases/tag/v0.1.9) - 2023-01-07 22:56:22

## [v0.1.8](https://github.com/sonrhq/core/releases/tag/v0.1.8) - 2023-01-07 22:34:54

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.7...v0.1.8

## [v0.1.7](https://github.com/sonrhq/core/releases/tag/v0.1.7) - 2023-01-07 21:25:19

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.6...v0.1.7

## [v0.1.6](https://github.com/sonrhq/core/releases/tag/v0.1.6) - 2023-01-07 21:15:28

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.5...v0.1.6

## [v0.1.5](https://github.com/sonrhq/core/releases/tag/v0.1.5) - 2023-01-07 20:41:40

## [v0.1.4](https://github.com/sonrhq/core/releases/tag/v0.1.4) - 2023-01-07 20:00:55

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.3...v0.1.4

## [v0.1.3](https://github.com/sonrhq/core/releases/tag/v0.1.3) - 2023-01-07 04:48:24

## [v0.1.2](https://github.com/sonrhq/core/releases/tag/v0.1.2) - 2023-01-07 04:04:12

**Full Changelog**: https://github.com/sonr-hq/sonr/compare/v0.1.1...v0.1.2

## [v0.1.1](https://github.com/sonrhq/core/releases/tag/v0.1.1) - 2023-01-07 02:59:05

## [v0.1.0](https://github.com/sonrhq/core/releases/tag/v0.1.0) - 2023-01-04 22:21:55

## [v0.0.3](https://github.com/sonrhq/core/releases/tag/v0.0.3) - 2022-12-27 21:03:35

## [v0.0.2](https://github.com/sonrhq/core/releases/tag/v0.0.2) - 2022-12-23 07:23:20

## [v0.0.1](https://github.com/sonrhq/core/releases/tag/v0.0.1) - 2022-12-13 19:22:04

**Full Changelog**: https://github.com/sonr-hq/sonr/commits/v0.0.1

\* *This CHANGELOG was automatically generated by [auto-generate-changelog](https://github.com/BobAnkh/auto-generate-changelog)*
