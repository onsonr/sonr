# Changelog

## v0.0.9 (04/13/2022)

### Added

- Deactivate (bucket, channel, object) module logic in Keeper
- Return codes and messages for module CUD methods
- Implemented UpdateBucket logic

### Modified

- Renamed Delete methods to utilize Deactivate to be compliant with the W3 Spec
- Query checks status of record before returning the document
- Update methods now check if the status of the record is active before updating

## v0.0.8 (04/07/2022)

### Added

- New registry module type `Session`

  - Contains metadata for `WhoIs`, `BaseDID`, and active `Credential`.
  - Utilized to verify current user and interact with other modules
  - Interfaces over Webauthn package
  - Generates New DID's based on underlying DID from `WhoIs`

- Added Module Specific Documentation for all x/{module} folders

  - Begun Status code table
  - Provided implemented method overview and reference
  - Provide Module record description

- Implemented business logic for the following methods

  - `RegisterApplication()`
  - `AccessApplication()`
  - `CreateObject()`
  - `CreateChannel()`
  - `CreateBucket()`

- Added TaskFile commands for `Serve`, `Start`, `Build`, and `Clean`

### Removed

- Removed `x/blob` module for less complex Document structure
- Removed Peer type definition from `x/registry`
- Removed ServiceConfig type definition from `x/registry`
- Removed Buf.Build configuration from all proto

### Changed

- Updated Module record key value

  - All `x/{module}` modules now utilize DIDs as the key value
  - Querying individual records now fully utilizes DIDs

- Added **Session** to all `x/{module}` modules transaction messages
- Upgraded Object to utilize Transformation methods to allow for more complex data structures
- Modified **CUD** Messages to utilized Record for each module specific storage

