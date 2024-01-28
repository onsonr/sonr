## Feature

- Reintroduced docs, motor, studio build.
- Frontend has been moved to internal.
- Added ECIES encryption package to crypto.
- Directory structure has been upgraded.
- Migration to nebula UI package has been undertaken.
- Added 'start' function to highway server.
- Added recovery email template.

## Improvement

- Replaced shares package with kss package.
- Added a new dependency: github.com/tink-crypto/tink-go/v2
- Renamed 'gateway' to 'highway' and 'keychain' to 'wallet'.
- Update of console and wallet handlers.

## Clean Up

- Deleted unused files and updated templates.
- Removed unused templates and handlers.
- Cleaned up Go mod using tidy.
- Unused code and redundant package names have been removed.

## Bug Fix

- Fixed issues with the studio directory.
- Fixed buf generation directory.
- Fixed Go build flags binding.
- Updated port number to 8000 and fixed cmd banner port for highway.
- Refactored encryption and decryption methods.
- Dialogue widths in swap and share modals have been updated.
