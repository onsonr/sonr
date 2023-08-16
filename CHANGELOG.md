# CHANGELOG

## [v0.6.29-beta.0](https://github.com/sonrhq/core/releases/tag/v0.6.29-beta.0) - 2023-06-23 21:25:15

## Changelog
* 7c3492a * chore(client.go): change broadcast mode from sync to async * feat(controller.go): add SignAndBroadcastCosmosTx method to Keeper struct

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

## [v0.6.28-beta.2](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.2) - 2023-06-10 16:40:10

## [v0.6.26](https://github.com/sonrhq/core/releases/tag/v0.6.26) - 2023-05-28 20:30:43

## [v0.7.6](https://github.com/sonrhq/core/releases/tag/v0.7.6) - 2023-08-16 19:44:33

< DESCRIPTION OF RELEASE >

## ⚡️ Binaries

Binaries for Linux and Darwin (amd64 and arm64) are available below.
Darwin users can also use the same universal binary `sonrd-0.7.6-darwin-all` for both amd64 and arm64.

#### 🔨 Build from source

If you prefer to build from source, you can use the following commands:

````bash
git clone https://github.com/sonrhq/core
cd core && git checkout v0.7.6
make build-darwin # or make build-linux
````

## 🐳 Run with Docker

As an alternative to installing and running sonrd on your system, you may run sonrd in a Docker container.
The following Docker images are available in our registry:

| Image Name                              | Base                                 | Description                       |
|-----------------------------------------|--------------------------------------|-----------------------------------|
| `sonrhq/core:0.7.6`            | `distroless/static-debian11`         | Default image based on Distroless |
| `sonrhq/core:0.7.6-distroless` | `distroless/static-debian11`         | Distroless image (same as above)  |
| `sonrhq/core:0.7.6-nonroot`    | `distroless/static-debian11:nonroot` | Distroless non-root image         |
| `sonrhq/core:0.7.6-alpine`     | `alpine`                             | Alpine image                      |

Example run:

```bash
docker run sonrhq/sonrd:0.7.6 version
# v0.7.6
````

All the images support `arm64` and `amd64` architectures.

## Changelog
* 5557674 * chore(summarizer.yml): remove unused GitHub Actions workflow file
* d6afe75 Implement/contracts (#115)
* 59b1924 * chore(bump.yml): remove changes-prefix for "Feature" category * chore(bump.yml): remove changes-prefix for "Maintenance" category * chore(bump.yml): remove changes-prefix for "Bug Fixes" category * chore(bump.yml): remove changes-prefix for "Documentation" category * chore(bump.yml): remove changes-prefix for "Dependency Updates" category * chore(bump.yml): remove skip-label for "Dependency Updates" category * fix(test.yml): add "master" branch to push event * chore(test.yml): add step to upload coverage reports to Codecov * chore(goreleaser.yaml): update pre-build hook to download libwasmvm.x86_64.so instead of libwasmvm_muslc.x86_64.a * chore(goreleaser.yaml): update ldflags to use shared linkmode instead of external linkmode
* fbe0225 * chore(build.yml): remove "master" branch from push trigger * chore(test.yml): remove "master" branch from push trigger
* 66ade0c * chore(release.yml): add condition to run the workflow only on tag starting with 'v' * chore(release.yml): update comment to specify that the workflow runs only on tag
* c1967c4 * chore(.goreleaser.yaml): add CGO_ENABLED environment variable * chore(.goreleaser.yaml): update COSMWASM_VERSION to v1.3.0
* 459d354 docs(CHANGELOG): update release notes

## [0.7.5](https://github.com/sonrhq/core/releases/tag/0.7.5) - 2023-08-14 23:49:32

*No description*

## [0.7.4](https://github.com/sonrhq/core/releases/tag/0.7.4) - 2023-08-13 23:55:00

*No description*

## [v0.7.3](https://github.com/sonrhq/core/releases/tag/v0.7.3) - 2023-08-07 23:55:39

## [v0.7.3-beta.4](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.4) - 2023-08-02 03:50:58

## [v0.7.3-beta.3](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.3) - 2023-07-31 21:35:25

## [v0.7.3-beta.2](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.2) - 2023-07-31 15:46:47

## [v0.7.3-beta.1](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.1) - 2023-07-28 01:32:34

## [v0.7.3-beta.0](https://github.com/sonrhq/core/releases/tag/v0.7.3-beta.0) - 2023-07-27 18:09:16

## [v0.7.2](https://github.com/sonrhq/core/releases/tag/v0.7.2) - 2023-07-27 11:48:51

## [v0.7.1](https://github.com/sonrhq/core/releases/tag/v0.7.1) - 2023-07-27 02:35:51

## [v0.7.1-beta.4](https://github.com/sonrhq/core/releases/tag/v0.7.1-beta.4) - 2023-07-14 05:50:49

## [v0.7.1-beta.3](https://github.com/sonrhq/core/releases/tag/v0.7.1-beta.3) - 2023-07-13 02:43:22

## [v0.7.0](https://github.com/sonrhq/core/releases/tag/v0.7.0) - 2023-06-28 22:51:16

## [v0.6.29-beta.5](https://github.com/sonrhq/core/releases/tag/v0.6.29-beta.5) - 2023-06-23 21:28:33

## Changelog
* 5edb208 * refactor(sfs.go): remove commented out code * refactor(sfs.go): move InsertPublicKeyshare and InsertEncryptedKeyshare to goroutines * fix(sfs.go): handle error when broadcasting transaction
* 44af9b5 * refactor(kss.go): use crypto.Base64Encode and crypto.Base64Decode to encode and decode data in keyshare store
* 6a2a17d Create KeyPrefix and Redis integration for db
* fa83e65 Remove account interface utilize kss, and credential for actions
* 648127b Remove highway methods from node query keeper
* ae8d3b0 * refactor(nacl.go): add context parameter to ClaimAccount function * refactor(services.go): add context parameter to ClaimAccount function call * refactor(claims.go): add context parameter to ClaimAccount function call
* c207120 feat(vscode): change explorer.excludeGitIgnore setting to false
* cf77bc2 chore(README.md): add wakatime badge to README.md

## [v0.6.28-beta.9](https://github.com/sonrhq/core/releases/tag/v0.6.28-beta.9) - 2023-06-11 00:42:34

## Changelog
* 431ce12 Add Webauthn Credential to register user response

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
