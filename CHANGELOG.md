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

## [v0.6.29-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.29-beta.0) - 2023-06-23 21:25:15

## [v0.6.29](https://github.com/sonrhq/core/releases/tag/v0.6.29) - 2023-06-23 21:28:33

## [v0.6.28-beta.16](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.16) - 2023-06-23 15:52:25

## [v0.6.28](https://github.com/sonrhq/core/releases/tag/v0.6.28) - 2023-06-23 16:46:22

## [v0.6.28-beta.10](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.10) - 2023-06-11 00:42:34

## [v0.6.28-beta.2](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.2) - 2023-06-10 16:40:10

## [v0.6.26](https://github.com/sonrhq/core/releases/tag/v0.6.26) - 2023-05-28 20:30:43

## [0.7.5](https://github.com/sonrhq/core/releases/tag/0.7.5) - 2023-08-14 23:49:32

*No description*

## [0.7.4](https://github.com/sonrhq/core/releases/tag/0.7.4) - 2023-08-13 23:55:00

*No description*

## [v0.7.3](https://github.com/sonrhq/core/releases/tag/v0.7.3) - 2023-08-07 23:55:39

## [v0.7.3-beta.4](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.4) - 2023-08-02 03:50:58

## [v0.7.3-beta.3](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.3) - 2023-07-31 21:35:25

## [v0.7.3-beta.2](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.2) - 2023-07-31 15:46:47

## [v0.6.27](https://github.com/sonrhq/core/releases/tag/v0.6.27) - 2023-06-01 18:10:36

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
